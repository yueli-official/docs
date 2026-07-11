<script setup lang="ts">
import { ManageEmpty, ManageHeader, ManagePageFooter, ManagePagination, ManageTabs, SkeletonList } from '@platform/ui/components'
import type { CollectionManageTree, CollectionVersion, CollectionVersionsResponse, CollectionView, DocDetail } from '~/types'
import { buildDocTree, docManageRoute, findDocSlugPathById } from '~/utils/docsManageRoutes.mjs'

definePageMeta({ layout: 'manage' })
useSeoMeta({ title: '文档 · 控制台' })

const { call } = useApi()
const toast = useToast()

type StatusKey = 'all' | 'draft' | 'published' | 'archived' | 'issues'
type ViewMode = 'list' | 'tree'
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

const search = ref('')
const collectionFilter = ref('all')
const versionFilter = ref('')
const localeFilter = ref('en')
const statusFilter = ref<StatusKey>('all')
const viewMode = ref<ViewMode>('list')
const page = ref(1)
const pageSize = 30
const selectedDocId = ref('')
const selectedDocIds = ref<string[]>([])
const bulkAction = ref<BulkDocAction | undefined>()
const bulkBusy = ref(false)

const bulkItems: Array<{ label: string; value: BulkDocAction; icon: string }> = [
  { label: '发布', value: 'publish', icon: 'i-tabler-rocket' },
  { label: '转草稿', value: 'draft', icon: 'i-tabler-pencil' },
  { label: '归档', value: 'archive', icon: 'i-tabler-archive' },
]

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
    { key: 'all', label: '全部' },
    { key: 'draft', label: '草稿' },
    { key: 'published', label: '已发布' },
    { key: 'archived', label: '归档' },
    { key: 'issues', label: '待完善' },
  ]
})

function qualityIssues(doc: ManagedDoc) {
  const issues: string[] = []
  if (doc.status === 'draft') issues.push('未发布')
  if (doc.status === 'archived') issues.push('已归档')
  if (!doc.excerpt?.trim()) issues.push('缺摘要')
  if (!doc.slugPath.length) issues.push('缺路径')
  if (!doc.collectionSlug) issues.push('缺文档集')
  return issues
}

function qualityMeta(doc: ManagedDoc) {
  const issues = qualityIssues(doc)
  if (!issues.length) return { label: '就绪', color: 'success' as const, icon: 'i-tabler-shield-check' }
  if (issues.length > 1) return { label: `${issues.length} 项`, color: 'error' as const, icon: 'i-tabler-alert-triangle' }
  return { label: issues[0], color: 'warning' as const, icon: 'i-tabler-alert-circle' }
}

function readinessItems(doc: ManagedDoc) {
  return [
    { label: '已发布', ok: doc.status === 'published' },
    { label: '摘要', ok: Boolean(doc.excerpt?.trim()) },
    { label: '语义路径', ok: doc.slugPath.length > 0 },
    { label: '文档集', ok: Boolean(doc.collectionSlug) },
  ]
}

const filteredDocs = computed(() => {
  const q = search.value.trim().toLowerCase()
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

const totalPages = computed(() => Math.max(1, Math.ceil(filteredDocs.value.length / pageSize)))
const pagedDocs = computed(() => {
  const start = (page.value - 1) * pageSize
  return filteredDocs.value.slice(start, start + pageSize)
})
const selectedDocs = computed(() =>
  allDocs.value.filter(doc => selectedDocIds.value.includes(doc.id)),
)
const selectedDoc = computed(() =>
  allDocs.value.find(doc => doc.id === selectedDocId.value) ?? pagedDocs.value[0] ?? null,
)
const pageSelectionCount = computed(() =>
  pagedDocs.value.filter(doc => selectedDocIds.value.includes(doc.id)).length,
)
const isPageSelected = computed(() =>
  pagedDocs.value.length > 0 && pageSelectionCount.value === pagedDocs.value.length,
)
const isPageIndeterminate = computed(() =>
  pageSelectionCount.value > 0 && pageSelectionCount.value < pagedDocs.value.length,
)

watch([search, activeCollection, activeStatus, activeLocale, activeVersion], () => {
  page.value = 1
  selectedDocIds.value = []
  bulkAction.value = undefined
})
watch(activeCollection, () => { versionFilter.value = '' })
watch(totalPages, (n) => { if (page.value > n) page.value = n })
watch(pagedDocs, (docs) => {
  if (!docs.length) {
    selectedDocId.value = ''
    return
  }
  if (!docs.some(doc => doc.id === selectedDocId.value)) selectedDocId.value = docs[0]!.id
}, { immediate: true })
watch(allDocs, (docs) => {
  const valid = new Set(docs.map(doc => doc.id))
  selectedDocIds.value = selectedDocIds.value.filter(id => valid.has(id))
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

function selectDoc(doc: ManagedDoc) {
  selectedDocId.value = doc.id
}

function toggleDocSelection(id: string) {
  selectedDocIds.value = selectedDocIds.value.includes(id)
    ? selectedDocIds.value.filter(item => item !== id)
    : [...selectedDocIds.value, id]
}

function togglePageSelection(value?: boolean | 'indeterminate') {
  const pageIds = pagedDocs.value.map(doc => doc.id)
  if (value === false || isPageSelected.value) {
    selectedDocIds.value = selectedDocIds.value.filter(id => !pageIds.includes(id))
    return
  }
  selectedDocIds.value = Array.from(new Set([...selectedDocIds.value, ...pageIds]))
}

function clearBulkSelection() {
  selectedDocIds.value = []
  bulkAction.value = undefined
}

function docPublicPath(doc: ManagedDoc) {
  return `/${doc.collectionSlug}/${doc.slugPath.map(encodeURIComponent).join('/')}`
}

async function copyText(value: string, title: string) {
  try {
    await navigator.clipboard.writeText(value)
    toast.add({ title, color: 'success', icon: 'i-tabler-check' })
  }
  catch {
    toast.add({ title: '复制失败', color: 'error', icon: 'i-tabler-alert-circle' })
  }
}

async function setDocStatus(doc: ManagedDoc, status: 'draft' | 'published' | 'archived') {
  try {
    await updateDocStatus(doc.id, status)
    toast.add({
      title: status === 'published' ? '已发布' : status === 'archived' ? '已归档' : '已转为草稿',
      color: 'success',
      icon: 'i-tabler-check',
    })
    await Promise.all([loadDocs(), refreshTree()])
  }
  catch (err: any) {
    toast.add({ title: '状态更新失败', description: err?.data?.message || '请重试', color: 'error' })
  }
}

function selectedDocMoreItems(doc: ManagedDoc) {
  const lifecycle: Array<{ label: string; icon: string; onSelect: () => void }> = []
  if (doc.status !== 'draft') {
    lifecycle.push({
      label: '转回草稿',
      icon: 'i-tabler-pencil',
      onSelect: () => setDocStatus(doc, 'draft'),
    })
  }
  if (doc.status !== 'archived') {
    lifecycle.push({
      label: '归档',
      icon: 'i-tabler-archive',
      onSelect: () => setDocStatus(doc, 'archived'),
    })
  }
  return [lifecycle]
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
  bulkBusy.value = true
  try {
    const results = await Promise.allSettled(docs.map(doc => updateDocStatus(doc.id, status)))
    const changed = results.filter(result => result.status === 'fulfilled').length
    const failed = results.length - changed
    if (changed) {
      toast.add({
        title: `已处理 ${changed} 篇文档`,
        description: failed ? `${failed} 篇失败，请刷新后重试` : undefined,
        color: failed ? 'warning' : 'success',
        icon: failed ? 'i-tabler-alert-circle' : 'i-tabler-check',
      })
    }
    if (!changed && failed) {
      toast.add({ title: '批量操作失败', description: '没有文档被更新', color: 'error' })
    }
    if (changed) {
      clearBulkSelection()
      await refreshTree()
    }
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
    toast.add({ title: '已删除文档', color: 'success', icon: 'i-tabler-check' })
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
      <ManageTabs v-model="statusFilter" :items="statusTabs" class="mb-4" />

      <section class="mb-4 rounded-lg border border-default bg-elevated/30 p-3">
        <div class="grid gap-3 lg:grid-cols-[minmax(0,1fr)_220px_150px_170px_auto]">
          <UInput
            v-model="search"
            icon="i-tabler-search"
            placeholder="搜索标题、路径、slug 或文档集"
            class="min-w-0"
          />
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
          <div class="inline-flex rounded-lg border border-default bg-default p-0.5">
            <button
              type="button"
              class="flex min-h-9 items-center gap-1.5 rounded-md px-3 text-sm transition"
              :class="viewMode === 'list' ? 'bg-elevated text-highlighted shadow-sm' : 'text-muted hover:text-default'"
              @click="viewMode = 'list'"
            >
              <UIcon name="i-tabler-list" class="size-4" />列表
            </button>
            <button
              type="button"
              class="flex min-h-9 items-center gap-1.5 rounded-md px-3 text-sm transition"
              :class="viewMode === 'tree' ? 'bg-elevated text-highlighted shadow-sm' : 'text-muted hover:text-default'"
              @click="viewMode = 'tree'"
            >
              <UIcon name="i-tabler-sitemap" class="size-4" />树状
            </button>
          </div>
        </div>
      </section>

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
          <div class="grid gap-4 xl:grid-cols-[minmax(0,1fr)_340px]">
            <div class="min-w-0 overflow-hidden rounded-lg border border-default bg-default">
              <div class="hidden grid-cols-[42px_minmax(200px,2fr)_140px_92px_108px_76px] gap-3 border-b border-default bg-elevated/45 px-4 py-2.5 text-xs font-medium text-muted lg:grid">
                <span class="flex items-center justify-center">
                  <UCheckbox
                    :model-value="isPageSelected"
                    :indeterminate="isPageIndeterminate"
                    aria-label="选择当前页文档"
                    @update:model-value="togglePageSelection"
                  />
                </span>
                <span>标题</span>
                <span>文档集</span>
                <span>质量</span>
                <span>父级</span>
                <span class="text-right">操作</span>
              </div>

              <div class="divide-y divide-default">
                <div
                  v-for="doc in pagedDocs"
                  :key="doc.id"
                  data-doc-row
                  class="grid cursor-pointer gap-3 px-4 py-3 transition lg:grid-cols-[42px_minmax(200px,2fr)_140px_92px_108px_76px] lg:items-center"
                  :class="selectedDoc?.id === doc.id ? 'bg-primary/5 ring-1 ring-inset ring-primary/20' : selectedDocIds.includes(doc.id) ? 'bg-primary/5' : 'hover:bg-elevated/55'"
                  @click="selectDoc(doc)"
                >
                  <div class="flex items-center lg:justify-center">
                    <UCheckbox
                      :model-value="selectedDocIds.includes(doc.id)"
                      :aria-label="`选择 ${doc.title}`"
                      @click.stop
                      @update:model-value="toggleDocSelection(doc.id)"
                    />
                  </div>

                  <div class="min-w-0 text-left">
                    <div class="flex min-w-0 items-center gap-2">
                      <span class="grid size-8 shrink-0 place-items-center rounded-lg bg-primary/10 text-primary">
                        <UIcon name="i-tabler-file-text" class="size-4" />
                      </span>
                      <div class="min-w-0">
                        <p class="truncate text-sm font-medium text-highlighted">{{ doc.title }}</p>
                        <p class="truncate text-xs text-muted">{{ doc.path }}</p>
                      </div>
                    </div>
                  </div>

                  <div class="min-w-0 text-sm text-default">
                    <p class="truncate">{{ doc.collectionTitle }}</p>
                    <p class="truncate text-xs text-muted">{{ doc.collectionSlug }}</p>
                  </div>

                  <div>
                    <UBadge
                      :label="qualityMeta(doc).label"
                      :color="qualityMeta(doc).color"
                      :icon="qualityMeta(doc).icon"
                      variant="subtle"
                      size="sm"
                    />
                  </div>

                  <div class="min-w-0 text-sm text-muted">
                    <span v-if="doc.parentTitle" class="truncate">{{ doc.parentTitle }}</span>
                    <span v-else class="text-dimmed">顶级</span>
                  </div>

                  <div class="flex justify-end gap-1">
                    <UTooltip text="添加子文档">
                      <UButton icon="i-tabler-file-plus" color="neutral" variant="ghost" size="sm" square @click.stop="addChild(doc)" />
                    </UTooltip>
                    <UTooltip text="编辑">
                      <UButton icon="i-tabler-pencil" color="neutral" variant="ghost" size="sm" square @click.stop="openDoc(doc)" />
                    </UTooltip>
                  </div>
                </div>
              </div>
            </div>

            <aside class="rounded-lg border border-default bg-default p-4 xl:sticky xl:top-24 xl:self-start">
              <template v-if="selectedDoc">
                <div class="mb-4 flex items-start justify-between gap-3">
                  <div class="min-w-0">
                    <p class="text-xs font-medium uppercase tracking-wide text-muted">文档详情</p>
                    <h2 class="mt-1 truncate text-base font-semibold text-highlighted">{{ selectedDoc.title }}</h2>
                  </div>
                </div>

                <div class="rounded-lg border border-default bg-elevated/35 p-3">
                  <div class="mb-3 flex items-center justify-between gap-3">
                    <p class="text-xs font-semibold uppercase tracking-wide text-muted">发布检查</p>
                    <UBadge
                      :label="qualityMeta(selectedDoc).label"
                      :color="qualityMeta(selectedDoc).color"
                      :icon="qualityMeta(selectedDoc).icon"
                      variant="subtle"
                      size="sm"
                    />
                  </div>
                  <div class="grid gap-2">
                    <div
                      v-for="item in readinessItems(selectedDoc)"
                      :key="item.label"
                      class="flex items-center justify-between gap-3 rounded-md bg-default px-2.5 py-2 text-sm"
                    >
                      <span class="text-default">{{ item.label }}</span>
                      <UIcon
                        :name="item.ok ? 'i-tabler-circle-check' : 'i-tabler-circle-x'"
                        class="size-4"
                        :class="item.ok ? 'text-success' : 'text-warning'"
                      />
                    </div>
                  </div>
                </div>

                <USeparator class="my-4" />

                <div class="space-y-3 text-sm">
                  <div>
                    <div class="flex items-center justify-between gap-3">
                      <p class="text-xs font-medium text-muted">公开路径</p>
                      <div class="flex gap-1">
                        <UTooltip text="打开公开页">
                          <UButton
                            icon="i-tabler-external-link"
                            color="neutral"
                            variant="ghost"
                            size="sm"
                            square
                            class="size-8"
                            aria-label="打开公开页"
                            :to="docPublicPath(selectedDoc)"
                            target="_blank"
                          />
                        </UTooltip>
                        <UTooltip text="复制路径">
                          <UButton
                            icon="i-tabler-copy"
                            color="neutral"
                            variant="ghost"
                            size="sm"
                            square
                            class="size-8"
                            aria-label="复制路径"
                            @click="copyText(docPublicPath(selectedDoc), '已复制路径')"
                          />
                        </UTooltip>
                      </div>
                    </div>
                    <p class="mt-1 truncate font-mono text-default">{{ docPublicPath(selectedDoc) }}</p>
                  </div>
                  <div class="grid grid-cols-2 gap-3">
                    <div>
                      <p class="text-xs font-medium text-muted">文档集</p>
                      <p class="mt-1 truncate text-default">{{ selectedDoc.collectionTitle }}</p>
                    </div>
                    <div>
                      <p class="text-xs font-medium text-muted">父级</p>
                      <p class="mt-1 truncate text-default">{{ selectedDoc.parentTitle || '顶级' }}</p>
                    </div>
                  </div>
                  <div v-if="selectedDoc.excerpt">
                    <p class="text-xs font-medium text-muted">摘要</p>
                    <p class="mt-1 line-clamp-3 text-default">{{ selectedDoc.excerpt }}</p>
                  </div>
                </div>

                <USeparator class="my-4" />

                <div class="grid gap-2">
                  <div class="mb-1 flex items-center justify-between gap-3">
                    <p class="text-xs font-semibold uppercase tracking-wide text-muted">文档操作</p>
                    <span class="font-mono text-xs text-muted">{{ selectedDoc.collectionSlug }}</span>
                  </div>
                  <UButton label="编辑文档" icon="i-tabler-pencil" block @click="openDoc(selectedDoc)" />
                  <UButton
                    v-if="selectedDoc.status !== 'published'"
                    label="发布"
                    icon="i-tabler-rocket"
                    color="primary"
                    variant="soft"
                    block
                    @click="setDocStatus(selectedDoc, 'published')"
                  />
                  <div class="grid grid-cols-2 gap-2">
                    <UButton label="子文档" icon="i-tabler-file-plus" color="neutral" variant="outline" @click="addChild(selectedDoc)" />
                    <UDropdownMenu :items="selectedDocMoreItems(selectedDoc)" :ui="{ content: 'w-40' }">
                      <UButton label="更多设置" icon="i-tabler-dots-vertical" color="neutral" variant="outline" block />
                    </UDropdownMenu>
                  </div>
                </div>
              </template>
              <ManageEmpty v-else icon="i-tabler-file-search" text="选择一篇文档查看操作" />
            </aside>
          </div>

          <ManagePageFooter>
            <template #left>
              <div class="flex flex-wrap items-center gap-2">
                <UCheckbox
                  :model-value="isPageSelected"
                  :indeterminate="isPageIndeterminate"
                  aria-label="选择当前页文档"
                  @update:model-value="togglePageSelection"
                />
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
            <template #right>
              <ManagePagination v-model="page" :total-pages="totalPages" />
            </template>
          </ManagePageFooter>
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
  </div>
</template>
