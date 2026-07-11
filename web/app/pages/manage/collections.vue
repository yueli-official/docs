<script setup lang="ts">
import { ManageCollectionCoverCrop, ManageEmpty, ManageHeader, ManagePageFooter, ManagePagination, ManageVisualAssetField, SkeletonList } from '@platform/manage/components'
import { useMinLoading } from '@platform/ui/use-min-loading'
import type { CollectionView } from '~/types'

definePageMeta({ layout: 'manage' })
useSeoMeta({ title: '文档集 · 控制台' })

type HealthTone = 'success' | 'warning' | 'error'

const { call } = useApi()
const toast = useToast()
const route = useRoute()
const search = ref('')
const page = ref(1)
const pageSize = 24
const mounted = ref(false)
onMounted(() => { mounted.value = true })

const { data, pending, error, refresh } = await useAsyncData(
  'manage-collections',
  () => call<{ items: CollectionView[] }>('/api/v1/collections'),
  { server: false, default: () => ({ items: [] as CollectionView[] }) },
)

const items = computed(() => data.value?.items ?? [])
const loadError = computed(() => error.value as { statusCode?: number, status?: number, data?: { message?: string }, message?: string } | null)
const isAuthError = computed(() => {
  const err = loadError.value
  return (err?.statusCode ?? err?.status) === 401
})
const loadErrorMessage = computed(() =>
  isAuthError.value ? '登录已失效，请重新登录后继续管理文档集。' : (loadError.value?.data?.message || loadError.value?.message || '加载文档集失败，请稍后重试。'),
)
const showSkeleton = useMinLoading(computed(() => !loadError.value && (!mounted.value || pending.value)))
const selectedId = ref('')
const { login } = useAuth()

function loginAgain() {
  return login(route.fullPath)
}

const opsStats = computed(() => {
  const collections = items.value
  const empty = collections.filter(item => !item.docCount).length
  const withCover = collections.filter(item => Boolean(item.coverUrl)).length
  const ready = collections.filter(item => collectionIssues(item).length === 0).length
  return { collections: collections.length, empty, withCover, ready }
})

const opsCards = computed(() => [
  { label: '文档集', value: opsStats.value.collections, detail: `${opsStats.value.ready} 个就绪`, icon: 'i-tabler-stack-2', tone: 'primary' },
  { label: '空集合', value: opsStats.value.empty, detail: '需要补内容', icon: 'i-tabler-folder-off', tone: opsStats.value.empty ? 'warning' : 'neutral' },
  { label: '封面覆盖', value: opsStats.value.withCover, detail: '集合页视觉资产', icon: 'i-tabler-photo-check', tone: 'neutral' },
])

const filteredItems = computed(() => {
  const q = search.value.trim().toLowerCase()
  if (!q) return items.value
  return items.value.filter(item =>
    [item.title, item.slug, item.description].some(v => (v || '').toLowerCase().includes(q)),
  )
})
const totalPages = computed(() => Math.max(1, Math.ceil(filteredItems.value.length / pageSize)))
const pagedItems = computed(() =>
  filteredItems.value.slice((page.value - 1) * pageSize, page.value * pageSize),
)
const selectedItem = computed(() =>
  items.value.find(item => item.id === selectedId.value) ?? pagedItems.value[0] ?? null,
)

watch(search, () => { page.value = 1 })
watch(totalPages, (n) => { if (page.value > n) page.value = n })
watch(pagedItems, (pageItems) => {
  if (!pageItems.length) {
    selectedId.value = ''
    return
  }
  if (!pageItems.some(item => item.id === selectedId.value)) selectedId.value = pageItems[0]!.id
}, { immediate: true })

const open = ref(false)
const current = ref<CollectionView | null>(null)
const form = reactive({ title: '', slug: '', description: '', icon: '', cover: '' })
const slugTouched = ref(false)
const saving = ref(false)
const coverPct = ref(-1)
const coverCropOpen = ref(false)
const coverCropFile = ref<File | null>(null)
const pendingCoverFile = ref<File | null>(null)
const pendingCoverPreview = ref('')
const confirmingDelete = ref(false)
const deletingBusy = ref(false)

const normalizedFormSlug = computed(() => clientSlug(form.slug))
const canSave = computed(() => Boolean(form.title.trim() && normalizedFormSlug.value))
const formPublicPath = computed(() => normalizedFormSlug.value ? `/${normalizedFormSlug.value}` : '')

function clientSlug(value: string) {
  return value
    .trim()
    .toLowerCase()
    .replace(/[^\p{L}\p{N}]+/gu, '-')
    .replace(/^-+|-+$/g, '')
}

watch(() => form.title, (title) => {
  if (!current.value && !slugTouched.value) form.slug = clientSlug(title)
})

function collectionPublicPath(col: CollectionView) {
  return col.slug ? `/${col.slug}` : ''
}

function collectionIssues(col: CollectionView) {
  const issues: string[] = []
  if (!col.slug) issues.push('缺路径标识')
  if (!col.description?.trim()) issues.push('缺描述')
  if (!col.docCount) issues.push('空集合')
  if (!col.icon && !col.coverUrl) issues.push('缺视觉资产')
  return issues
}

function collectionHealth(col: CollectionView): { label: string; color: HealthTone; icon: string } {
  const issues = collectionIssues(col)
  if (!issues.length) return { label: '就绪', color: 'success', icon: 'i-tabler-shield-check' }
  if (issues.includes('缺路径标识') || issues.length > 2) return { label: `${issues.length} 项`, color: 'error', icon: 'i-tabler-alert-triangle' }
  return { label: issues[0]!, color: 'warning', icon: 'i-tabler-alert-circle' }
}

function statToneClass(tone: string) {
  if (tone === 'success') return 'bg-success/10 text-success ring-success/20'
  if (tone === 'warning') return 'bg-warning/10 text-warning ring-warning/20'
  if (tone === 'primary') return 'bg-primary/10 text-primary ring-primary/20'
  return 'bg-elevated text-muted ring-default'
}

function selectCollection(col: CollectionView) {
  selectedId.value = col.id
}

function openCreate() {
  clearPendingCover()
  current.value = null
  form.title = ''
  form.slug = ''
  form.description = ''
  form.icon = ''
  form.cover = ''
  slugTouched.value = false
  confirmingDelete.value = false
  open.value = true
}

function openEdit(col: CollectionView) {
  clearPendingCover()
  current.value = col
  selectedId.value = col.id
  form.title = col.title
  form.slug = col.slug || ''
  form.description = col.description || ''
  form.icon = col.icon || ''
  form.cover = col.coverUrl || ''
  slugTouched.value = true
  confirmingDelete.value = false
  open.value = true
}

function clearPendingCover() {
  if (pendingCoverPreview.value) URL.revokeObjectURL(pendingCoverPreview.value)
  pendingCoverPreview.value = ''
  pendingCoverFile.value = null
  coverCropFile.value = null
  coverCropOpen.value = false
  coverPct.value = -1
}

function setCoverUrl(value: string) {
  clearPendingCover()
  form.cover = value
}

function onPickCover(file: File) {
  if (!file.type.startsWith('image/')) {
    toast.add({ title: '请选择图片文件', color: 'warning', icon: 'i-tabler-alert-triangle' })
    return
  }
  if (file.size > 10 * 1024 * 1024) {
    toast.add({ title: '图片不能超过 10MB', color: 'warning', icon: 'i-tabler-alert-triangle' })
    return
  }
  if (pendingCoverPreview.value) URL.revokeObjectURL(pendingCoverPreview.value)
  pendingCoverPreview.value = ''
  pendingCoverFile.value = null
  coverPct.value = -1
  coverCropFile.value = file
  coverCropOpen.value = true
}

function onCroppedCover(file: File) {
  if (pendingCoverPreview.value) URL.revokeObjectURL(pendingCoverPreview.value)
  pendingCoverFile.value = file
  pendingCoverPreview.value = URL.createObjectURL(file)
  coverCropFile.value = null
}

function putWithProgress(url: string, file: File, onProgress?: (pct: number) => void, headers?: Record<string, string>): Promise<void> {
  return new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest()
    xhr.open('PUT', url)
    for (const [key, value] of Object.entries(headers ?? {})) xhr.setRequestHeader(key, value)
    xhr.upload.onprogress = (event) => {
      if (event.lengthComputable && onProgress) onProgress(Math.round((event.loaded / event.total) * 100))
    }
    xhr.onload = () => xhr.status >= 200 && xhr.status < 300
      ? resolve()
      : reject(new Error(`上传失败 (HTTP ${xhr.status})`))
    xhr.onerror = () => reject(new Error('上传网络错误，请确认素材服务在线'))
    xhr.send(file)
  })
}

async function uploadCollectionCover(collectionId: string, file: File) {
  coverPct.value = 0
  const init = await call<{ uploadUrl: string; uploadToken: string; uploadHeaders?: Record<string, string> }>(
    `/api/v1/collections/${collectionId}/cover`,
    { method: 'POST', body: { filename: file.name, mime: file.type, size: file.size } },
  )
  await putWithProgress(init.uploadUrl, file, pct => { coverPct.value = pct }, init.uploadHeaders)
  const res = await call<{ collection: CollectionView; coverUrl: string }>(
    `/api/v1/collections/${collectionId}/cover/finalize`,
    { method: 'POST', body: { uploadToken: init.uploadToken } },
  )
  return res
}

async function copyText(value: string, title: string) {
  if (!value) return
  try {
    await navigator.clipboard.writeText(value)
    toast.add({ title, color: 'success', icon: 'i-tabler-check' })
  }
  catch {
    toast.add({ title: '复制失败', color: 'error', icon: 'i-tabler-alert-circle' })
  }
}

async function save() {
  if (!form.title.trim()) {
    toast.add({ title: '请填写标题', color: 'warning', icon: 'i-tabler-alert-triangle' })
    return
  }
  if (!normalizedFormSlug.value) {
    toast.add({ title: '请填写路径标识', color: 'warning', icon: 'i-tabler-alert-triangle' })
    return
  }

  saving.value = true
  try {
    const file = pendingCoverFile.value
    const body = {
      title: form.title.trim(),
      slug: form.slug.trim(),
      description: form.description,
      icon: form.icon,
      cover: form.cover,
    }
    const res = current.value
      ? await call<{ collection: CollectionView }>(`/api/v1/collections/${current.value.id}`, { method: 'PATCH', body })
      : await call<{ collection: CollectionView }>('/api/v1/collections', { method: 'POST', body })

    let saved = res.collection
    if (file && saved?.id) {
      const coverRes = await uploadCollectionCover(saved.id, file)
      saved = coverRes.collection
      form.cover = coverRes.coverUrl
      clearPendingCover()
    }

    toast.add({
      title: current.value ? '已更新文档集' : '已创建',
      color: 'success',
      icon: 'i-tabler-check',
    })
    open.value = false
    await refresh()
    if (saved?.id) selectedId.value = saved.id
  }
  catch (e: any) {
    toast.add({ title: current.value ? '更新失败' : '创建失败', description: e?.data?.message || '请重试', color: 'error' })
  }
  finally {
    saving.value = false
    coverPct.value = -1
  }
}

async function doDelete() {
  if (!current.value) return
  deletingBusy.value = true
  try {
    await call(`/api/v1/collections/${current.value.id}`, { method: 'DELETE' })
    toast.add({ title: `已删除「${current.value.title}」`, color: 'success', icon: 'i-tabler-check' })
    open.value = false
    selectedId.value = ''
    await refresh()
  }
  catch (e: any) {
    toast.add({ title: '删除失败', description: e?.data?.message || '请重试', color: 'error' })
  }
  finally {
    deletingBusy.value = false
  }
}
</script>

<template>
  <div>
    <ManageHeader title="文档集">
      <template #subtitle>
        <span>{{ filteredItems.length }} / {{ items.length }} 个集合</span>
      </template>
      <template #actions>
        <UButton icon="i-tabler-plus" label="新建文档集" @click="openCreate" />
      </template>
    </ManageHeader>

    <div v-if="loadError" class="rounded-lg border border-default bg-default p-8">
      <div class="mx-auto max-w-md text-center">
        <span class="mx-auto grid size-12 place-items-center rounded-lg bg-warning/10 text-warning">
          <UIcon :name="isAuthError ? 'i-tabler-lock' : 'i-tabler-alert-triangle'" class="size-6" />
        </span>
        <h2 class="mt-4 text-lg font-semibold text-highlighted">{{ isAuthError ? '需要重新登录' : '加载失败' }}</h2>
        <p class="mt-2 text-sm leading-6 text-muted">{{ loadErrorMessage }}</p>
        <div class="mt-5 flex justify-center gap-2">
          <UButton v-if="isAuthError" icon="i-tabler-login-2" label="重新登录" @click="() => { loginAgain() }" />
          <UButton v-else icon="i-tabler-refresh" label="重试" color="neutral" variant="soft" @click="() => { refresh() }" />
        </div>
      </div>
    </div>

    <SkeletonList v-else-if="showSkeleton" :rows="8" />

    <template v-else>
      <section class="mb-5 overflow-hidden rounded-lg border border-default bg-default">
        <div class="grid 2xl:grid-cols-[minmax(360px,0.85fr)_minmax(0,1.15fr)]">
          <div class="border-b border-default bg-elevated/35 p-5 2xl:border-b-0 2xl:border-r">
            <div class="flex items-center gap-2 text-xs font-semibold uppercase tracking-wide text-muted">
              <UIcon name="i-tabler-building-bank" class="size-4 text-primary" />
              文档集运营
            </div>
            <div class="mt-3 flex flex-wrap items-end justify-between gap-3">
              <div class="min-w-0">
                <h2 class="text-2xl font-semibold text-highlighted">文档集管理</h2>
                <p class="mt-1 text-sm text-muted">维护公开路径、说明文案和集合页视觉资产</p>
              </div>
              <UBadge
                :label="pending ? '同步中' : '实时数据'"
                :color="pending ? 'warning' : 'success'"
                :icon="pending ? 'i-tabler-loader-2' : 'i-tabler-radar-2'"
                variant="subtle"
              />
            </div>
          </div>

          <div class="grid sm:grid-cols-3">
            <div
              v-for="card in opsCards"
              :key="card.label"
              class="border-b border-default p-4 sm:border-r xl:border-b-0"
            >
              <div class="mb-4 grid size-10 place-items-center rounded-lg ring-1" :class="statToneClass(card.tone)">
                <UIcon :name="card.icon" class="size-5" />
              </div>
              <p class="text-xs font-medium text-muted">{{ card.label }}</p>
              <div class="mt-1 flex items-baseline gap-2">
                <span class="text-2xl font-semibold text-highlighted">{{ card.value }}</span>
                <span class="truncate text-xs text-muted">{{ card.detail }}</span>
              </div>
            </div>
          </div>
        </div>
      </section>

      <div class="mb-4 rounded-lg border border-default bg-elevated/30 p-3">
        <div class="grid gap-3 lg:grid-cols-[minmax(0,1fr)_auto]">
          <UInput
            v-model="search"
            icon="i-tabler-search"
            placeholder="搜索标题、路径标识或说明"
            class="min-w-0"
          />
          <div class="flex items-center justify-between gap-3 rounded-lg border border-default bg-default px-3 py-2 text-sm lg:min-w-64">
            <span class="text-muted">当前结果</span>
            <span class="font-mono text-highlighted">{{ filteredItems.length }} / {{ items.length }}</span>
          </div>
        </div>
      </div>

      <ManageEmpty
        v-if="!items.length"
        icon="i-tabler-stack-2"
        text="还没有文档集"
      />

      <ManageEmpty
        v-else-if="!filteredItems.length"
        icon="i-tabler-search-off"
        text="没有匹配的文档集"
      />

      <template v-else>
        <div class="grid gap-4 xl:grid-cols-[minmax(0,1fr)_360px]">
          <div class="min-w-0 overflow-hidden rounded-lg border border-default bg-default">
            <div class="hidden grid-cols-[minmax(260px,1.8fr)_160px_120px_130px] gap-3 border-b border-default bg-elevated/45 px-4 py-2.5 text-xs font-medium text-muted lg:grid">
              <span>文档集</span>
              <span>公开路径</span>
              <span>文档数</span>
              <span>检查</span>
            </div>

            <div class="divide-y divide-default">
              <button
                v-for="col in pagedItems"
                :key="col.id"
                type="button"
                class="grid w-full gap-3 px-4 py-3 text-left transition lg:grid-cols-[minmax(260px,1.8fr)_160px_120px_130px] lg:items-center"
                :class="selectedItem?.id === col.id ? 'bg-primary/5 ring-1 ring-inset ring-primary/20' : 'hover:bg-elevated/55'"
                @click="selectCollection(col)"
              >
                <div class="min-w-0">
                  <div class="flex min-w-0 items-center gap-3">
                    <img
                      v-if="col.coverUrl"
                      :src="col.coverUrl"
                      :alt="col.title"
                      class="size-10 shrink-0 rounded-lg object-cover"
                    />
                    <span v-else class="grid size-10 shrink-0 place-items-center rounded-lg bg-primary/10 text-primary">
                      <UIcon :name="col.icon || 'i-tabler-stack-2'" class="size-5" />
                    </span>
                    <div class="min-w-0">
                      <p class="truncate text-sm font-semibold text-highlighted">{{ col.title }}</p>
                      <p class="truncate text-xs text-muted">{{ col.description || '未填写集合说明' }}</p>
                    </div>
                  </div>
                </div>

                <div class="min-w-0 font-mono text-sm text-default">
                  <p class="truncate">{{ collectionPublicPath(col) || '-' }}</p>
                </div>

                <div class="text-sm text-default">
                  <span class="font-semibold text-highlighted">{{ col.docCount }}</span>
                  <span class="ml-1 text-muted">篇</span>
                </div>

                <div>
                  <UBadge
                    :label="collectionHealth(col).label"
                    :color="collectionHealth(col).color"
                    :icon="collectionHealth(col).icon"
                    variant="subtle"
                    size="sm"
                  />
                </div>

              </button>
            </div>
          </div>

          <aside class="rounded-lg border border-default bg-default p-4 xl:sticky xl:top-24 xl:self-start">
            <template v-if="selectedItem">
              <div class="mb-4 flex items-start justify-between gap-3">
                <div class="min-w-0">
                  <p class="text-xs font-medium uppercase tracking-wide text-muted">文档集详情</p>
                  <h2 class="mt-1 truncate text-base font-semibold text-highlighted">{{ selectedItem.title }}</h2>
                </div>
                <UBadge
                  :label="collectionHealth(selectedItem).label"
                  :color="collectionHealth(selectedItem).color"
                  :icon="collectionHealth(selectedItem).icon"
                  variant="subtle"
                />
              </div>

              <div class="rounded-lg border border-default bg-elevated/35 p-3">
                <div class="mb-3 flex items-center justify-between gap-3">
                  <p class="text-xs font-semibold uppercase tracking-wide text-muted">发布检查</p>
                  <span class="font-mono text-xs text-muted">{{ selectedItem.slug }}</span>
                </div>
                <div class="grid gap-2">
                  <div
                    v-for="issue in ['路径标识', '描述', '文档内容', '视觉资产']"
                    :key="issue"
                    class="flex items-center justify-between gap-3 rounded-md bg-default px-2.5 py-2 text-sm"
                  >
                    <span class="text-default">{{ issue }}</span>
                    <UIcon
                      :name="
                        (issue === '路径标识' && selectedItem.slug)
                          || (issue === '描述' && selectedItem.description)
                          || (issue === '文档内容' && selectedItem.docCount)
                          || (issue === '视觉资产' && (selectedItem.icon || selectedItem.coverUrl))
                          ? 'i-tabler-circle-check'
                          : 'i-tabler-circle-x'
                      "
                      class="size-4"
                      :class="
                        (issue === '路径标识' && selectedItem.slug)
                          || (issue === '描述' && selectedItem.description)
                          || (issue === '文档内容' && selectedItem.docCount)
                          || (issue === '视觉资产' && (selectedItem.icon || selectedItem.coverUrl))
                          ? 'text-success'
                          : 'text-warning'
                      "
                    />
                  </div>
                </div>
              </div>

              <USeparator class="my-4" />

              <div class="space-y-3 text-sm">
                <div>
                  <div class="flex items-center justify-between gap-3">
                    <p class="text-xs font-medium text-muted">线上入口</p>
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
                          :to="collectionPublicPath(selectedItem)"
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
                          @click="copyText(collectionPublicPath(selectedItem), '已复制路径')"
                        />
                      </UTooltip>
                    </div>
                  </div>
                  <p class="mt-1 truncate font-mono text-default">{{ collectionPublicPath(selectedItem) }}</p>
                </div>
                <div>
                  <p class="text-xs font-medium text-muted">文档容量</p>
                  <p class="mt-1 text-default">{{ selectedItem.docCount }} 篇文档</p>
                </div>
                <div v-if="selectedItem.description">
                  <p class="text-xs font-medium text-muted">说明</p>
                  <p class="mt-1 line-clamp-3 text-default">{{ selectedItem.description }}</p>
                </div>
              </div>

              <USeparator class="my-4" />

              <div class="grid gap-2">
                <UButton label="编辑文档集" icon="i-tabler-pencil" block @click="openEdit(selectedItem)" />
              </div>
            </template>
            <ManageEmpty v-else icon="i-tabler-stack-2" text="选择一个文档集查看操作" />
          </aside>
        </div>

        <ManagePageFooter>
          <template #left>
            <span>每页 {{ pageSize }} 个</span>
            <span class="hidden sm:inline">当前 {{ (page - 1) * pageSize + 1 }}-{{ Math.min(page * pageSize, filteredItems.length) }}</span>
          </template>
          <template #right>
            <ManagePagination v-model="page" :total-pages="totalPages" />
          </template>
        </ManagePageFooter>
      </template>
    </template>

    <USlideover v-model:open="open" :title="current ? '编辑文档集' : '新建文档集'">
      <template #body>
        <div class="space-y-5">
          <div class="space-y-4">
            <UFormField label="标题" required>
              <UInput v-model="form.title" placeholder="文档集标题" class="w-full" autofocus />
            </UFormField>

            <UFormField label="路径标识" required help="集合公开路径；保存后该集合下文档路径会同步使用新标识。">
              <UInput
                v-model="form.slug"
                icon="i-tabler-link"
                placeholder="quickstart"
                class="w-full"
                @input="slugTouched = true"
              />
            </UFormField>

            <div class="rounded-lg border border-default bg-elevated/35 px-3 py-2">
              <div class="flex items-center justify-between gap-3">
                <span class="text-xs font-medium text-muted">路径预览</span>
                <span class="truncate font-mono text-sm text-default">{{ formPublicPath || '/...' }}</span>
              </div>
            </div>

            <UFormField label="描述">
              <UTextarea v-model="form.description" :rows="3" class="w-full" />
            </UFormField>

            <ManageVisualAssetField
              :icon="form.icon"
              :cover-url="form.cover"
              :preview-url="pendingCoverPreview"
              :progress="coverPct"
              :uploading="coverPct >= 0"
              @update:icon="form.icon = $event"
              @update:cover-url="setCoverUrl"
              @pick-cover="onPickCover"
              @clear-cover="form.cover = ''"
            />
          </div>

          <template v-if="current">
            <USeparator />
            <div class="space-y-3">
              <p class="text-xs font-semibold uppercase tracking-wide text-muted">危险操作</p>
              <div class="rounded-lg border border-error/30 bg-error/5 p-3">
                <div v-if="!confirmingDelete" class="flex items-center justify-between gap-3">
                  <div class="min-w-0">
                    <p class="text-sm font-medium text-highlighted">删除文档集</p>
                    <p class="mt-0.5 text-xs text-muted">会连同该集下所有文档一并删除，不可恢复。</p>
                  </div>
                  <UButton label="删除" icon="i-tabler-trash" color="error" variant="soft" class="shrink-0" @click="() => { confirmingDelete = true }" />
                </div>
                <div v-else>
                  <p class="text-sm text-highlighted">确定删除「{{ current.title }}」？此集下所有文档将一并删除，不可恢复。</p>
                  <div class="mt-3 flex justify-end gap-2">
                    <UButton label="取消" color="neutral" variant="ghost" size="sm" @click="() => { confirmingDelete = false }" />
                    <UButton label="确认删除" icon="i-tabler-trash" color="error" size="sm" :loading="deletingBusy" @click="doDelete" />
                  </div>
                </div>
              </div>
            </div>
          </template>
        </div>
      </template>

      <template #footer>
        <div class="flex w-full justify-end gap-2">
          <UButton label="取消" color="neutral" variant="outline" @click="() => { open = false }" />
          <UButton
            :label="current ? '保存' : '创建'"
            icon="i-tabler-check"
            :loading="saving"
            :disabled="!canSave"
            @click="save"
          />
        </div>
      </template>
    </USlideover>
    <ManageCollectionCoverCrop
      v-model:open="coverCropOpen"
      :file="coverCropFile"
      @cropped="onCroppedCover"
    />
  </div>
</template>
