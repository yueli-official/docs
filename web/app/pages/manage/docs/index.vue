<script setup lang="ts">
import { createPlatformNotifier } from '@platform/ui/feedback'
import {
  ManageCollectionDock,
  ManageCollectionToolbar,
  ManageEmpty,
  ManageHeader,
  ManageLifecycleTabs,
  ManagePageSelection,
  ManagePagination,
  ManageRowShell,
  ManageViewToggle,
  SkeletonList
} from '@platform/manage/components'
import type { ManageCollectionDefinition } from '@platform/manage/collection'
import { useManageCollectionState } from '@platform/manage/use-manage-collection-state'
import { useManageSelection } from '@platform/manage/use-manage-selection'
import type { CollectionManageTree, CollectionVersion, CollectionVersionsResponse, CollectionView, DocDetail } from '~/types'
import { buildDocTree, docManageRoute, findDocSlugPathById } from '~/utils/docsManageRoutes.mjs'

definePageMeta({ layout: 'manage' })
useSeoMeta({ title: '文档 · 控制台' })

const { call } = useApi()
const toast = createPlatformNotifier(useToast())
const route = useRoute()
const router = useRouter()

type StatusKey = 'all' | 'draft' | 'published' | 'archived' | 'issues'
type BulkDocAction = 'publish' | 'draft' | 'archive'

interface ManagedDoc extends DocDetail {
  collectionTitle: string
  collectionSlug: string
  parentTitle: string
  path: string
  slugPath: string[]
  depth: number
}

interface MoveIntent {
  dragId: string
  targetId: string
  position: 'before' | 'after' | 'into'
}
interface PatchItem {
  id: string
  parentId: string
  sortOrder: number
}

const collectionDefinition = {
  resourceKind: 'doc',
  statuses: ['all', 'draft', 'published', 'archived', 'issues'],
  views: ['list', 'tree'],
  sortKeys: ['sortOrder'],
  pageSizes: [30, 60, 90],
  defaultStatus: 'all',
  defaultView: 'list',
  defaultSort: 'sortOrder',
  defaultDirection: 'asc',
  defaultPageSize: 30,
  pagination: 'client',
  selection: 'page',
  filters: ['collection', 'version', 'locale'],
  quickEditFields: ['title', 'slug', 'status', 'parent'],
  bulkActions: ['publish', 'draft', 'archive']
} as const satisfies ManageCollectionDefinition

const {
  status: statusFilter,
  searchInput: search,
  q: searchQuery,
  page,
  size: pageSize,
  view: viewMode,
  filterModel
} = useManageCollectionState({
  definition: collectionDefinition,
  routeQuery: computed(() => route.query),
  replaceQuery: query => router.replace({ query })
})
const collectionFilter = filterModel('collection', 'all')
const versionFilter = filterModel('version')
const localeFilter = filterModel('locale', 'en')
const bulkAction = ref<BulkDocAction | undefined>()
const bulkBusy = ref(false)
const bulkResult = ref<{ changed: number, failedIds: string[], interrupted?: boolean, message?: string }>()
const quickEditTarget = ref<ManagedDoc>()
const showQuickEdit = ref(false)

const bulkItems: Array<{ label: string; value: BulkDocAction; icon: string }> = [
  { label: '发布', value: 'publish', icon: 'i-tabler-rocket' },
  { label: '转草稿', value: 'draft', icon: 'i-tabler-pencil' },
  { label: '归档', value: 'archive', icon: 'i-tabler-archive' },
]
const pageSizeItems = [30, 60, 90].map(value => ({ label: `${value}/页`, value }))

const { data: cols, pending: collectionsPending } = await useAsyncData(
  'manage-doc-workbench-collections',
  () => call<{ items: CollectionView[] }>('/api/v1/collections'),
  { server: false, default: () => ({ items: [] as CollectionView[] }) },
)
const collections = computed(() => cols.value?.items ?? [])
const collectionById = computed(() => Object.fromEntries(collections.value.map(c => [c.id, c])))

function selectedValue(value: unknown) {
  if (value && typeof value === 'object' && 'value' in value) {
    const option = value as { value?: unknown }
    return String(option.value ?? '')
  }
  return String(value ?? '')
}

const activeCollection = computed(() => selectedValue(collectionFilter.value) || 'all')
const activeVersion = computed(() => selectedValue(versionFilter.value))
const activeLocale = computed(() => selectedValue(localeFilter.value) || 'en')
const activeStatus = computed<StatusKey>(() => {
  const value = selectedValue(statusFilter.value)
  return value === 'draft' || value === 'published' || value === 'archived' || value === 'issues' ? value : 'all'
})

const versionsByCollection = ref<Record<string, CollectionVersion[]>>({})
const localeItems = [
  { label: 'English', value: 'en' },
  { label: '简体中文', value: 'zh-CN' },
]

const collectionItems = computed(() => [
  { label: '全部文档集', value: 'all' },
  ...collections.value.map(c => ({ label: c.title, value: c.id })),
])
const selectedCollectionVersions = computed(() =>
  activeCollection.value === 'all' ? [] : versionsByCollection.value[activeCollection.value] ?? [],
)
const versionItems = computed(() => [
  { label: '默认版本', value: '' },
  ...selectedCollectionVersions.value.map(v => ({
    label: v.isDefault ? `${v.label}（默认）` : v.label,
    value: v.key,
    description: v.key,
  })),
])
const treeCollectionItems = computed(() =>
  collections.value.map(c => ({ label: c.title, value: c.slug })),
)

const docsByCollection = ref<Record<string, DocDetail[]>>({})
const docsPending = ref(false)
const docsError = ref('')
let loadSeq = 0

async function loadVersions() {
  const items = collections.value
  if (!items.length) {
    versionsByCollection.value = {}
    return
  }
  const entries = await Promise.all(items.map(async (col) => {
    const res = await call<CollectionVersionsResponse>(`/api/v1/collections/${col.id}/versions`)
    return [col.id, res.items] as const
  }))
  versionsByCollection.value = Object.fromEntries(entries)
}

async function loadDocs() {
  const items = collections.value
  if (!items.length) {
    docsByCollection.value = {}
    return
  }

  const seq = ++loadSeq
  docsPending.value = true
  docsError.value = ''
  try {
    const entries = await Promise.all(items.map(async (col) => {
      const version = activeCollection.value !== 'all' && col.id === activeCollection.value ? activeVersion.value : ''
      const res = await call<{ items: DocDetail[] }>('/api/v1/docs', {
        query: {
          collectionId: col.id,
          locale: activeLocale.value,
          ...(version ? { version } : {}),
        },
      })
      return [col.id, res.items] as const
    }))
    if (seq === loadSeq) docsByCollection.value = Object.fromEntries(entries)
  }
  catch (err: any) {
    docsError.value = err?.data?.message || '加载文档失败'
  }
  finally {
    if (seq === loadSeq) docsPending.value = false
  }
}

watch(collections, () => { loadVersions() }, { immediate: true })
watch([collections, activeLocale, activeVersion, activeCollection], () => { loadDocs() }, { immediate: true })

function enrichDocs(col: CollectionView, docs: DocDetail[]): ManagedDoc[] {
  const byId = new Map(docs.map(d => [d.id, d]))

  function pathFor(doc: DocDetail) {
    const parts = [doc.title]
    let cursor = byId.get(doc.parentId)
    let guard = 0
    while (cursor && guard < 20) {
      parts.unshift(cursor.title)
      cursor = byId.get(cursor.parentId)
      guard++
    }
    return parts.join(' / ')
  }

  function slugPathFor(doc: DocDetail) {
    const parts = [doc.slug]
    let cursor = byId.get(doc.parentId)
    let guard = 0
    while (cursor && guard < 20) {
      parts.unshift(cursor.slug)
      cursor = byId.get(cursor.parentId)
      guard++
    }
    return parts.filter(Boolean)
  }

  function depthFor(doc: DocDetail) {
    let depth = 0
    let cursor = byId.get(doc.parentId)
    let guard = 0
    while (cursor && guard < 20) {
      depth++
      cursor = byId.get(cursor.parentId)
      guard++
    }
    return depth
  }

  return docs.map(doc => ({
    ...doc,
    collectionTitle: col.title,
    collectionSlug: col.slug,
    parentTitle: byId.get(doc.parentId)?.title ?? '',
    path: pathFor(doc),
    slugPath: slugPathFor(doc),
    depth: depthFor(doc),
  }))
}

const allDocs = computed(() =>
  collections.value.flatMap(col => enrichDocs(col, docsByCollection.value[col.id] ?? [])),
)

const statusTabs = computed(() => {
  return [
    { key: 'all', label: '全部', count: allDocs.value.length },
    { key: 'draft', label: '草稿', count: allDocs.value.filter(doc => doc.status === 'draft').length },
    { key: 'published', label: '已发布', count: allDocs.value.filter(doc => doc.status === 'published').length },
    { key: 'archived', label: '归档', count: allDocs.value.filter(doc => doc.status === 'archived').length },
    { key: 'issues', label: '待完善', count: allDocs.value.filter(doc => qualityIssues(doc).length > 0).length },
  ]
})

function qualityIssues(doc: ManagedDoc) {
  const issues: string[] = []
  if (!doc.excerpt?.trim()) issues.push('缺摘要')
  if (!doc.slugPath.length) issues.push('缺路径')
  if (!doc.collectionSlug) issues.push('缺文档集')
  return issues
}

const filteredDocs = computed(() => {
  const q = searchQuery.value.trim().toLowerCase()
  return allDocs.value.filter((doc) => {
    if (activeCollection.value !== 'all' && doc.collectionId !== activeCollection.value) return false
    if (activeStatus.value === 'issues' && !qualityIssues(doc).length) return false
    if (activeStatus.value !== 'all' && activeStatus.value !== 'issues' && doc.status !== activeStatus.value) return false
    if (!q) return true
    return [
      doc.title,
      doc.slug,
      doc.path,
      doc.collectionTitle,
      doc.excerpt,
    ].some(v => (v || '').toLowerCase().includes(q))
  })
})

const totalPages = computed(() => Math.max(1, Math.ceil(filteredDocs.value.length / pageSize.value)))
const pagedDocs = computed(() => {
  const start = (page.value - 1) * pageSize.value
  return filteredDocs.value.slice(start, start + pageSize.value)
})
const {
  selectedIds: selectedDocIds,
  isPageSelected,
  isPageIndeterminate,
  toggleOne: toggleDocSelection,
  togglePage: togglePageSelection,
  clear: clearDocSelection,
  keepOnly: keepValidSelection,
  replace: replaceSelection
} = useManageSelection({
  visibleIds: computed(() => pagedDocs.value.map(doc => doc.id)),
  filteredTotal: computed(() => filteredDocs.value.length),
  resetKey: computed(() => route.fullPath)
})
const selectedDocs = computed(() =>
  allDocs.value.filter(doc => selectedDocIds.value.includes(doc.id)),
)

watch([activeCollection, activeStatus, activeLocale, activeVersion], () => {
  selectedDocIds.value = []
  bulkAction.value = undefined
  bulkResult.value = undefined
})
watch(activeCollection, () => { versionFilter.value = '' })
watch(totalPages, (n) => { if (page.value > n) page.value = n })
watch(allDocs, (docs) => {
  keepValidSelection(docs.map(doc => doc.id))
})

const selectedCollection = computed(() =>
  activeCollection.value === 'all' ? null : collectionById.value[activeCollection.value],
)

const createTarget = computed(() => {
  const slug = selectedCollection.value?.slug
  return slug ? `/manage/docs/new?collection=${encodeURIComponent(slug)}` : '/manage/docs/new'
})

function openDoc(docOrId: ManagedDoc | string) {
  const doc = typeof docOrId === 'string'
    ? allDocs.value.find(item => item.id === docOrId)
    : docOrId
  if (!doc && typeof docOrId === 'string') {
    const slugPath = findDocSlugPathById(tree.value?.tree ?? [], docOrId)
    if (slugPath && activeTreeSlug.value) {
      navigateTo(docManageRoute(activeTreeSlug.value, slugPath))
      return
    }
  }
  if (!doc) {
    navigateTo(`/manage/docs/${docOrId}`)
    return
  }
  navigateTo(docManageRoute(doc.collectionSlug, doc.slugPath))
}

function addChild(doc: ManagedDoc) {
  navigateTo(`/manage/docs/new?collection=${encodeURIComponent(doc.collectionSlug)}&parent=${encodeURIComponent(doc.id)}`)
}

function openQuickEdit(doc: ManagedDoc) {
  quickEditTarget.value = doc
  showQuickEdit.value = true
}

async function onQuickEditSaved() {
  await Promise.all([loadDocs(), refreshTree()])
}

function clearBulkSelection() {
  clearDocSelection()
  bulkAction.value = undefined
}

async function updateDocStatus(id: string, status: 'draft' | 'published' | 'archived') {
  if (status === 'published') {
    await call(`/api/v1/docs/${id}/publish`, { method: 'POST' })
    return
  }
  if (status === 'archived') {
    await call(`/api/v1/docs/${id}/archive`, { method: 'POST' })
    return
  }
  await call(`/api/v1/docs/${id}`, { method: 'PATCH', body: { status } })
}

async function applyBulkAction() {
  if (!bulkAction.value || !selectedDocs.value.length) return

  const status = bulkAction.value === 'publish'
    ? 'published'
    : bulkAction.value === 'archive'
      ? 'archived'
      : 'draft'
  const docs = [...selectedDocs.value]
  const requestedIds = docs.map(doc => doc.id)
  bulkBusy.value = true
  bulkResult.value = undefined
  try {
    const results = await Promise.allSettled(docs.map(doc => updateDocStatus(doc.id, status)))
    const changed = results.filter(result => result.status === 'fulfilled').length
    const failedIds = results.flatMap((result, index) => result.status === 'rejected' ? [requestedIds[index]!] : [])
    bulkResult.value = { changed, failedIds }
    if (failedIds.length) replaceSelection(failedIds)
    else clearBulkSelection()
    bulkAction.value = undefined
    await refreshTree()
  }
  catch (error) {
    const apiError = error as { data?: { message?: string } }
    replaceSelection(requestedIds)
    bulkResult.value = {
      changed: 0,
      failedIds: requestedIds,
      interrupted: true,
      message: apiError.data?.message || '批量请求中断，已保留选择，请核对当前状态后重试。'
    }
    await refreshTree()
  }
  finally {
    bulkBusy.value = false
  }
}

// ─── Structure view ──────────────────────────────────────────────────────────

const treeSlug = ref('')
const activeTreeSlug = computed(() => selectedValue(treeSlug.value))
watch(collections, (items) => {
  if (!activeTreeSlug.value && items.length) treeSlug.value = items[0]!.slug
}, { immediate: true })
watch(activeCollection, (id) => {
  if (id !== 'all') treeSlug.value = collectionById.value[id]?.slug ?? treeSlug.value
})

const activeTreeCollection = computed(() =>
  collections.value.find(c => c.slug === activeTreeSlug.value) ?? null,
)
const treeOverride = ref<CollectionManageTree | null>(null)
const treePending = computed(() => docsPending.value)
const tree = computed<CollectionManageTree | null>(() =>
  treeOverride.value ?? (activeTreeCollection.value
    ? {
        collection: activeTreeCollection.value,
        tree: buildDocTree(docsByCollection.value[activeTreeCollection.value.id] ?? []),
      }
    : null),
)
watch([activeTreeSlug, docsByCollection], () => { treeOverride.value = null }, { deep: true })

async function refreshTree() {
  treeOverride.value = null
  await loadDocs()
}

function findInTree(
  nodes: DocDetail[],
  id: string,
  parent: DocDetail | null = null,
): { node: DocDetail; parent: DocDetail | null; siblings: DocDetail[]; index: number } | null {
  for (let i = 0; i < nodes.length; i++) {
    const n = nodes[i]!
    if (n.id === id) return { node: n, parent, siblings: nodes, index: i }
    if (n.children?.length) {
      const r = findInTree(n.children, id, n)
      if (r) return r
    }
  }
  return null
}

function isDescendant(node: DocDetail, targetId: string): boolean {
  return node.children?.some(c => c.id === targetId || isDescendant(c, targetId)) ?? false
}

function computePatches(treeNodes: DocDetail[], intent: MoveIntent): PatchItem[] | null {
  const dragCtx = findInTree(treeNodes, intent.dragId)
  const targetCtx = findInTree(treeNodes, intent.targetId)
  if (!dragCtx || !targetCtx) return null

  const { node: dragNode, parent: dragParent, siblings: dragSiblings, index: dragIdx } = dragCtx
  if (isDescendant(dragNode, intent.targetId)) return null

  const oldParentId = dragParent?.id ?? ''
  let newParentId: string
  let newSiblings: DocDetail[]
  let insertIdx: number

  if (intent.position === 'into') {
    newParentId = intent.targetId
    newSiblings = targetCtx.node.children ?? []
    insertIdx = newSiblings.length
  }
  else {
    newParentId = targetCtx.parent?.id ?? ''
    newSiblings = targetCtx.siblings
    insertIdx = intent.position === 'before' ? targetCtx.index : targetCtx.index + 1
  }

  const patches: PatchItem[] = []
  const isSameParent = oldParentId === newParentId

  if (isSameParent) {
    const reordered = [...dragSiblings]
    reordered.splice(dragIdx, 1)
    const adj = insertIdx > dragIdx ? insertIdx - 1 : insertIdx
    reordered.splice(adj, 0, dragNode)
    reordered.forEach((n, i) => {
      if (n.sortOrder !== i)
        patches.push({ id: n.id, parentId: newParentId, sortOrder: i })
    })
  }
  else {
    dragSiblings.filter((_, i) => i !== dragIdx).forEach((n, i) => {
      if (n.sortOrder !== i)
        patches.push({ id: n.id, parentId: oldParentId, sortOrder: i })
    })
    const newOrder = [...newSiblings]
    newOrder.splice(insertIdx, 0, dragNode)
    newOrder.forEach((n, i) => {
      if (n.sortOrder !== i || n.id === dragNode.id)
        patches.push({ id: n.id, parentId: newParentId, sortOrder: i })
    })
  }

  return patches.length > 0 ? patches : null
}

function applyMoveToTree(nodes: DocDetail[], intent: MoveIntent) {
  const dragCtx = findInTree(nodes, intent.dragId)
  if (!dragCtx) return
  const { node: dragNode, siblings: oldSiblings, index: dragIdx } = dragCtx

  const targetBefore = findInTree(nodes, intent.targetId)
  const newParentId = intent.position === 'into' ? intent.targetId : targetBefore?.parent?.id ?? ''

  oldSiblings.splice(dragIdx, 1)
  oldSiblings.forEach((n, i) => { n.sortOrder = i })
  dragNode.parentId = newParentId

  const targetAfter = findInTree(nodes, intent.targetId)
  if (!targetAfter) return

  if (intent.position === 'into') {
    if (!targetAfter.node.children) targetAfter.node.children = []
    targetAfter.node.children.push(dragNode)
    targetAfter.node.children.forEach((n, i) => { n.sortOrder = i })
  }
  else {
    const siblings = targetAfter.parent?.children ?? nodes
    const idx = intent.position === 'before' ? targetAfter.index : targetAfter.index + 1
    siblings.splice(idx, 0, dragNode)
    siblings.forEach((n, i) => { n.sortOrder = i })
  }
}

async function onMove(intent: MoveIntent) {
  if (!tree.value) return
  const patches = computePatches(tree.value.tree, intent)
  if (!patches?.length) return

  const snapshot = JSON.parse(JSON.stringify(tree.value)) as CollectionManageTree
  const optimistic = JSON.parse(JSON.stringify(tree.value)) as CollectionManageTree
  applyMoveToTree(optimistic.tree, intent)
  treeOverride.value = optimistic

  try {
    await Promise.all(patches.map(p =>
      call(`/api/v1/docs/${p.id}`, {
        method: 'PATCH',
        body: {
          parentId: p.parentId,
          sortOrder: p.sortOrder,
        },
      }),
    ))
    await loadDocs()
  }
  catch {
    treeOverride.value = snapshot
    toast.add({ title: '移动失败，已还原', color: 'error' })
    await refreshTree()
  }
}

async function onDelete(id: string) {
  try {
    await call(`/api/v1/docs/${id}`, { method: 'DELETE' })
    await Promise.all([refreshTree(), loadDocs()])
  }
  catch (err: any) {
    toast.add({ title: '删除失败', description: err?.data?.message || '请重试', color: 'error' })
  }
}
</script>

<template>
  <div>
    <ManageHeader title="文档">
      <template #subtitle>
        <span>搜索、筛选、批量处理与层级调整</span>
      </template>
      <template #actions>
        <UButton icon="i-tabler-plus" label="新建文档" :to="createTarget" />
      </template>
    </ManageHeader>

    <ClientOnly>
      <ManageLifecycleTabs v-model="statusFilter" :items="statusTabs" class="mb-4" />

      <ManageCollectionToolbar v-model:search="search" search-placeholder="搜索标题、路径、slug 或文档集…" class="mb-4">
        <template #filters>
          <USelectMenu
            :model-value="collectionFilter"
            :items="collectionItems"
            value-key="value"
            placeholder="筛选文档集"
            :search-input="{ placeholder: '搜索文档集…' }"
            class="w-full"
            @update:model-value="collectionFilter = selectedValue($event)"
          />
          <USelectMenu
            :model-value="localeFilter"
            :items="localeItems"
            value-key="value"
            placeholder="语言"
            class="w-full"
            @update:model-value="localeFilter = selectedValue($event) || 'en'"
          />
          <USelectMenu
            :model-value="versionFilter"
            :items="versionItems"
            value-key="value"
            placeholder="版本"
            :disabled="activeCollection === 'all'"
            class="w-full"
            @update:model-value="versionFilter = selectedValue($event)"
          />
        </template>
        <template #actions>
          <ManageViewToggle v-model="viewMode" :items="[
            { key: 'list', label: '列表', icon: 'i-tabler-list' },
            { key: 'tree', label: '树状', icon: 'i-tabler-sitemap' }
          ]" />
        </template>
      </ManageCollectionToolbar>

      <UAlert
        v-if="docsError"
        color="error"
        variant="soft"
        icon="i-tabler-alert-circle"
        :title="docsError"
        class="mb-4"
      />

      <template v-if="viewMode === 'list'">
        <SkeletonList v-if="collectionsPending || docsPending" :rows="8" />

        <ManageEmpty
          v-else-if="!collections.length"
          icon="i-tabler-stack-2"
          text="还没有文档集"
        />

        <ManageEmpty
          v-else-if="!filteredDocs.length"
          icon="i-tabler-file-search"
          text="没有匹配的文档"
        />

        <template v-else>
          <div class="overflow-hidden rounded-lg border border-default bg-default" :inert="bulkBusy" :aria-busy="bulkBusy">
            <ManageRowShell
              v-for="doc in pagedDocs"
              :key="doc.id"
              :selected="selectedDocIds.includes(doc.id)"
              :selection-disabled="bulkBusy"
              :selection-label="`选择文档：${doc.title}`"
              @select="toggleDocSelection(doc.id)"
            >
              <template #media>
                <span class="grid size-10 shrink-0 place-items-center rounded-lg bg-primary/10 text-primary">
                  <UIcon name="i-tabler-file-text" class="size-5" />
                </span>
              </template>

              <div class="min-w-0">
                <button type="button" class="block max-w-full truncate text-left text-sm font-semibold text-highlighted hover:text-primary" @click="openQuickEdit(doc)">
                  {{ doc.title }}
                </button>
                <p class="mt-0.5 truncate font-mono text-xs text-muted">{{ doc.slugPath.join(' / ') }}</p>
                <p class="mt-1 truncate text-xs text-dimmed">{{ doc.collectionTitle }} · {{ doc.parentTitle || '顶级文档' }}</p>
                <p v-if="qualityIssues(doc).length" class="mt-1 inline-flex max-w-full items-center gap-1 truncate text-xs text-warning">
                  <UIcon name="i-tabler-alert-circle" class="size-3.5 shrink-0" />
                  <span class="truncate">{{ qualityIssues(doc).slice(0, 2).join(' · ') }}</span>
                </p>
              </div>

              <template #meta>
                <div class="min-w-0 text-xs md:w-36 md:text-right">
                  <p class="truncate text-default">{{ doc.locale }}</p>
                  <p class="mt-0.5 truncate text-muted">{{ doc.parentTitle || '顶级文档' }}</p>
                </div>
              </template>

              <template #actions>
                <UTooltip text="添加子文档">
                  <UButton icon="i-tabler-file-plus" color="neutral" variant="ghost" size="sm" square :aria-label="`添加子文档：${doc.title}`" @click="addChild(doc)" />
                </UTooltip>
                <UTooltip text="快速编辑">
                  <UButton icon="i-tabler-pencil" color="neutral" variant="ghost" size="sm" square :aria-label="`快速编辑：${doc.title}`" @click="openQuickEdit(doc)" />
                </UTooltip>
                <UTooltip text="完整编辑">
                  <UButton icon="i-tabler-file-pencil" color="neutral" variant="ghost" size="sm" square :aria-label="`完整编辑：${doc.title}`" @click="openDoc(doc)" />
                </UTooltip>
              </template>
            </ManageRowShell>
          </div>

          <ManageCollectionDock label="文档批量操作与分页">
            <template #selection>
              <div class="flex flex-wrap items-center gap-2">
                <ManagePageSelection
                  :model-value="isPageSelected"
                  :indeterminate="isPageIndeterminate"
                  label="选择当前页文档"
                  @update:model-value="togglePageSelection"
                />
                <div v-if="bulkResult" class="flex min-w-0 flex-wrap items-center gap-2 rounded-lg bg-elevated px-2.5 py-1.5">
                  <UIcon
                    :name="bulkResult.interrupted || bulkResult.failedIds.length ? 'i-tabler-alert-triangle' : 'i-tabler-circle-check'"
                    :class="bulkResult.interrupted || bulkResult.failedIds.length ? 'text-warning' : 'text-success'"
                  />
                  <span class="text-xs text-default">
                    <template v-if="bulkResult.interrupted">{{ bulkResult.message }}</template>
                    <template v-else>
                      已处理 {{ bulkResult.changed }} 篇<span v-if="bulkResult.failedIds.length">，{{ bulkResult.failedIds.length }} 篇未完成</span>
                    </template>
                  </span>
                  <UButton
                    v-if="bulkResult.failedIds[0]"
                    label="查看首个失败项"
                    color="warning"
                    variant="link"
                    size="xs"
                    @click="openDoc(bulkResult.failedIds[0])"
                  />
                  <UButton icon="i-tabler-x" color="neutral" variant="ghost" size="xs" square aria-label="关闭批量结果" @click="bulkResult = undefined" />
                </div>
                <template v-if="selectedDocIds.length">
                  <span class="text-sm text-default">已选 {{ selectedDocIds.length }}</span>
                  <USeparator orientation="vertical" class="hidden h-4 sm:block" />
                  <USelect
                    v-model="bulkAction"
                    :items="bulkItems"
                    value-key="value"
                    placeholder="批量操作"
                    size="sm"
                    class="w-32"
                  />
                  <UButton
                    label="应用"
                    icon="i-tabler-check"
                    color="primary"
                    variant="soft"
                    size="sm"
                    :disabled="!bulkAction"
                    :loading="bulkBusy"
                    @click="applyBulkAction"
                  />
                  <UButton
                    label="清空"
                    icon="i-tabler-x"
                    color="neutral"
                    variant="ghost"
                    size="sm"
                    @click="clearBulkSelection"
                  />
                </template>
                <template v-else>
                  <span>每页 {{ pageSize }} 篇</span>
                  <span class="hidden sm:inline">当前 {{ (page - 1) * pageSize + 1 }}-{{ Math.min(page * pageSize, filteredDocs.length) }}</span>
                </template>
              </div>
            </template>
            <template #pagination>
              <USelect v-model="pageSize" :items="pageSizeItems" value-key="value" size="sm" class="w-20" :disabled="bulkBusy" />
              <ManagePagination v-model="page" :total-pages="totalPages" class="!mt-0" />
            </template>
          </ManageCollectionDock>
        </template>
      </template>

      <template v-else>
        <div class="mb-4 flex flex-wrap items-center justify-between gap-3 rounded-lg border border-default bg-elevated/35 px-4 py-3">
          <div class="min-w-0">
            <p class="text-sm font-medium text-highlighted">树状结构</p>
            <p class="text-xs text-muted">{{ tree?.tree?.length ?? 0 }} 个根节点 · {{ activeTreeSlug || '未选择文档集' }}</p>
          </div>
          <USelectMenu
            :model-value="treeSlug"
            :items="treeCollectionItems"
            value-key="value"
            placeholder="选择文档集"
            :search-input="{ placeholder: '搜索文档集…' }"
            class="w-56"
            @update:model-value="treeSlug = selectedValue($event)"
          />
        </div>

        <SkeletonList v-if="treePending" :rows="8" />
        <ManageEmpty v-else-if="!tree?.tree?.length" icon="i-tabler-sitemap" text="这个文档集还没有树状结构" />
        <div v-else class="rounded-lg border border-default bg-default p-3">
          <DocTreeAdmin
            :nodes="tree.tree"
            :collection-slug="activeTreeSlug"
            @edit="openDoc"
            @add-child="(parentId) => navigateTo(`/manage/docs/new?collection=${activeTreeSlug}&parent=${parentId}`)"
            @delete="onDelete"
            @move="onMove"
          />
        </div>
      </template>
    </ClientOnly>

    <ManageDocQuickEditModal
      v-model:open="showQuickEdit"
      :doc="quickEditTarget"
      :docs="allDocs"
      @saved="onQuickEditSaved"
      @open-full="() => quickEditTarget && openDoc(quickEditTarget)"
    />
  </div>
</template>
