<script setup lang="ts">
import type { Ref } from 'vue'
import type { DocNode } from '~/types'

interface MoveIntent {
  dragId: string
  targetId: string
  position: 'before' | 'after' | 'into'
}

interface DragState {
  node: DocNode | null
}

defineProps<{ nodes: DocNode[]; collectionSlug: string }>()
const emit = defineEmits<{
  edit: [id: string]
  addChild: [parentId: string]
  delete: [id: string]
  move: [intent: MoveIntent]
}>()

const confirmingId = ref<string | null>(null)

// ─── Shared drag state across recursive instances via provide/inject ───────────
const DRAG_KEY = 'docTreeAdminDrag'
const parentState = inject<Ref<DragState> | null>(DRAG_KEY, null)
const dragState: Ref<DragState> = parentState ?? ref<DragState>({ node: null })
if (!parentState) provide(DRAG_KEY, dragState)

// ─── Per-instance drop indicator state ───────────────────────────────────────
const dropInfo = ref<{ nodeId: string; position: 'before' | 'after' | 'into' } | null>(null)

// Returns true if targetId is a descendant of node (cycle guard)
function isDescendantOf(node: DocNode, targetId: string): boolean {
  return node.children?.some(c => c.id === targetId || isDescendantOf(c, targetId)) ?? false
}

function getDropPosition(e: DragEvent): 'before' | 'after' | 'into' {
  const el = e.currentTarget as HTMLElement
  const rect = el.getBoundingClientRect()
  const ratio = (e.clientY - rect.top) / rect.height
  if (ratio < 0.28) return 'before'
  if (ratio > 0.72) return 'after'
  return 'into'
}

function onDragStart(e: DragEvent, node: DocNode) {
  // Don't initiate drag when user clicks action buttons
  if ((e.target as HTMLElement).closest('button')) {
    e.preventDefault()
    return
  }
  dragState.value.node = node
  e.dataTransfer!.effectAllowed = 'move'
  e.dataTransfer!.setData('text/plain', node.id) // required for Firefox
}

function onDragEnd() {
  dragState.value.node = null
  dropInfo.value = null
}

function onDragOver(e: DragEvent, node: DocNode) {
  const dragging = dragState.value.node
  if (!dragging || dragging.id === node.id) return
  if (isDescendantOf(dragging, node.id)) return // refuse drop into own subtree
  e.preventDefault()
  e.dataTransfer!.dropEffect = 'move'
  dropInfo.value = { nodeId: node.id, position: getDropPosition(e) }
}

function onDragLeave(e: DragEvent, node: DocNode) {
  if (dropInfo.value?.nodeId !== node.id) return
  const el = e.currentTarget as HTMLElement
  const rel = e.relatedTarget as Node | null
  if (rel && el.contains(rel)) return // pointer moved to a child element, still inside
  dropInfo.value = null
}

function onDrop(e: DragEvent, node: DocNode) {
  e.preventDefault()
  const info = dropInfo.value
  dropInfo.value = null
  const dragging = dragState.value.node
  dragState.value.node = null
  if (!dragging || !info || info.nodeId !== node.id) return
  if (dragging.id === node.id || isDescendantOf(dragging, node.id)) return
  emit('move', { dragId: dragging.id, targetId: node.id, position: info.position })
}
</script>

<template>
  <ul class="space-y-0.5">
    <li v-for="n in nodes" :key="n.id" class="relative">
      <!-- drop-before line -->
      <div
        v-if="dropInfo?.nodeId === n.id && dropInfo.position === 'before'"
        class="pointer-events-none absolute inset-x-0 top-0 z-10 h-0.5 rounded bg-primary"
      />

      <!-- node row -->
      <div
        class="group flex cursor-grab items-center gap-2 rounded-md px-2 py-1.5 hover:bg-elevated"
        :class="{
          'opacity-50': dragState.node?.id === n.id,
          'ring-2 ring-inset ring-primary': dropInfo?.nodeId === n.id && dropInfo.position === 'into',
        }"
        draggable="true"
        @dragstart="onDragStart($event, n)"
        @dragend="onDragEnd"
        @dragover="onDragOver($event, n)"
        @dragleave="onDragLeave($event, n)"
        @drop="onDrop($event, n)"
      >
        <UIcon
          name="i-tabler-grip-vertical"
          class="h-4 w-4 shrink-0 text-muted opacity-0 transition group-hover:opacity-60"
        />

        <span class="flex min-w-0 flex-1 items-center gap-2">
          <span class="truncate text-sm text-default">{{ n.title }}</span>
          <UBadge v-if="n.status === 'draft'" size="xs" variant="subtle" color="neutral">草稿</UBadge>
        </span>

        <!-- normal actions (visible on hover) -->
        <span v-if="confirmingId !== n.id" class="flex items-center gap-0.5 opacity-0 transition group-hover:opacity-100">
          <UButton
            icon="i-tabler-pencil"
            variant="ghost"
            size="xs"
            color="neutral"
            aria-label="编辑"
            @click="emit('edit', n.id)"
          />
          <UButton
            icon="i-tabler-plus"
            variant="ghost"
            size="xs"
            color="neutral"
            aria-label="加子文档"
            @click="emit('addChild', n.id)"
          />
          <UButton
            icon="i-tabler-trash"
            variant="ghost"
            size="xs"
            color="error"
            aria-label="删除"
            @click="() => { confirmingId = n.id }"
          />
        </span>

        <!-- inline delete confirm (replaces action icons, no modal) -->
        <span v-else class="flex items-center gap-1">
          <UButton
            label="取消"
            variant="ghost"
            size="xs"
            color="neutral"
            @click="() => { confirmingId = null }"
          />
          <UButton
            label="确认删除"
            size="xs"
            color="error"
            @click="emit('delete', n.id); confirmingId = null"
          />
        </span>
      </div>

      <!-- drop-after line -->
      <div
        v-if="dropInfo?.nodeId === n.id && dropInfo.position === 'after'"
        class="pointer-events-none absolute inset-x-0 bottom-0 z-10 h-0.5 rounded bg-primary"
      />

      <!-- children: recurse with all emits forwarded -->
      <div v-if="n.children?.length" class="ml-3 mt-0.5 border-l border-default pl-2">
        <DocTreeAdmin
          :nodes="n.children"
          :collection-slug="collectionSlug"
          @edit="emit('edit', $event)"
          @add-child="emit('addChild', $event)"
          @delete="emit('delete', $event)"
          @move="emit('move', $event)"
        />
      </div>
    </li>
  </ul>
</template>
