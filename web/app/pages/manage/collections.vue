<script setup lang="ts">
import { ManageCollectionCoverCrop, ManageCollectionDock, ManageCollectionToolbar, ManageEmpty, ManageHeader, ManagePagination, ManageVisualAssetField, SkeletonList } from '@platform/manage/components'
import type { ManageCollectionDefinition } from '@platform/manage/collection'
import { useManageCollectionState } from '@platform/manage/use-manage-collection-state'
import { useMinLoading } from '@platform/ui/use-min-loading'
import type { CollectionView } from '~/types'

definePageMeta({ layout: 'manage' })
useSeoMeta({ title: '文档集 · 控制台' })

const { call } = useApi()
const toast = useToast()
const route = useRoute()
const router = useRouter()
const collectionDefinition = {
  resourceKind: 'docs-collection',
  statuses: [''],
  views: ['list'],
  sortKeys: ['title', 'docCount'],
  pageSizes: [12, 24, 48],
  defaultStatus: '',
  defaultView: 'list',
  defaultSort: 'title',
  defaultDirection: 'asc',
  defaultPageSize: 24,
  pagination: 'client',
  selection: 'page',
  filters: []
} as const satisfies ManageCollectionDefinition
const { searchInput, q, sort, direction, page, size } = useManageCollectionState({
  definition: collectionDefinition,
  routeQuery: computed(() => route.query),
  replaceQuery: query => router.replace({ query })
})
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
const { login } = useAuth()

function loginAgain() {
  return login(route.fullPath)
}

const filteredItems = computed(() => {
  const keyword = q.value.trim().toLowerCase()
  const filtered = keyword
    ? items.value.filter(item => [item.title, item.slug, item.description].some(value => (value || '').toLowerCase().includes(keyword)))
    : [...items.value]
  const multiplier = direction.value === 'asc' ? 1 : -1
  return filtered.sort((a, b) => {
    if (sort.value === 'docCount') return ((a.docCount || 0) - (b.docCount || 0)) * multiplier
    return a.title.localeCompare(b.title, 'zh-CN') * multiplier
  })
})
const totalPages = computed(() => Math.max(1, Math.ceil(filteredItems.value.length / size.value)))
const pagedItems = computed(() =>
  filteredItems.value.slice((page.value - 1) * size.value, page.value * size.value),
)

watch(totalPages, (n) => { if (page.value > n) page.value = n })
const sortItems = [
  { label: '按标题', value: 'title' },
  { label: '按文档数', value: 'docCount' }
]
const pageSizeItems = [12, 24, 48].map(value => ({ label: `${value}/页`, value }))

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

function toggleDirection() {
  direction.value = direction.value === 'asc' ? 'desc' : 'asc'
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

    open.value = false
    await refresh()
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
    open.value = false
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
        <span>维护公开路径、说明和集合视觉资产。</span>
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
      <ManageCollectionToolbar v-model:search="searchInput" search-placeholder="搜索标题、路径标识或说明…" class="mb-5">
        <template #filters>
          <USelectMenu v-model="sort" :items="sortItems" value-key="value" icon="i-tabler-arrows-sort" size="sm" />
          <UButton
            :icon="direction === 'asc' ? 'i-tabler-sort-ascending' : 'i-tabler-sort-descending'"
            :label="direction === 'asc' ? '升序' : '降序'"
            color="neutral"
            variant="outline"
            size="sm"
            @click="toggleDirection"
          />
        </template>
      </ManageCollectionToolbar>

      <ManageEmpty v-if="!items.length" icon="i-tabler-stack-2" text="还没有文档集" />
      <ManageEmpty v-else-if="!filteredItems.length" icon="i-tabler-search-off" text="没有匹配的文档集" />

      <div v-else class="overflow-hidden rounded-xl border border-default bg-default">
        <div class="hidden grid-cols-[minmax(16rem,1.4fr)_minmax(10rem,.8fr)_7rem_3rem] items-center gap-3 border-b border-default bg-elevated/45 px-4 py-2.5 text-xs font-medium text-muted lg:grid">
          <span>文档集</span>
          <span>公开路径</span>
          <span>文档数</span>
          <span class="sr-only">操作</span>
        </div>

        <div class="divide-y divide-default">
          <button
            v-for="col in pagedItems"
            :key="col.id"
            type="button"
            class="grid w-full grid-cols-[minmax(0,1fr)_2.75rem] items-center gap-3 p-3 text-left transition hover:bg-elevated/55 focus-visible:outline-2 focus-visible:outline-offset-[-2px] focus-visible:outline-primary sm:p-4 lg:grid-cols-[minmax(16rem,1.4fr)_minmax(10rem,.8fr)_7rem_3rem]"
            :aria-label="`编辑文档集：${col.title}`"
            @click="openEdit(col)"
          >
            <span class="flex min-w-0 items-center gap-3">
              <img v-if="col.coverUrl" :src="col.coverUrl" :alt="col.title" class="size-11 shrink-0 rounded-lg object-cover" />
              <span v-else class="grid size-11 shrink-0 place-items-center rounded-lg bg-primary/10 text-primary">
                <UIcon :name="col.icon || 'i-tabler-stack-2'" class="size-5" />
              </span>
              <span class="min-w-0">
                <span class="line-clamp-1 text-sm font-semibold text-highlighted">{{ col.title }}</span>
                <span class="mt-1 line-clamp-1 text-xs text-muted">{{ col.description || '未填写集合说明' }}</span>
              </span>
            </span>

            <span class="col-start-1 min-w-0 pl-14 font-mono text-xs text-muted sm:text-sm lg:col-start-auto lg:pl-0">
              <span class="line-clamp-1">{{ collectionPublicPath(col) || '-' }}</span>
            </span>

            <span class="col-start-1 pl-14 text-xs text-muted sm:text-sm lg:col-start-auto lg:pl-0">
              <span class="font-semibold text-highlighted">{{ col.docCount }}</span> 篇文档
            </span>

            <span class="row-start-1 col-start-2 grid size-11 place-items-center text-muted lg:row-auto lg:col-start-auto" aria-hidden="true">
              <UIcon name="i-tabler-pencil" class="size-4" />
            </span>
          </button>
        </div>
      </div>

      <ManageCollectionDock v-if="pagedItems.length" label="文档集统计与分页">
        <template #selection>
          <span>共 {{ filteredItems.length }} 个文档集</span>
          <span v-if="q" class="text-xs text-muted">全部 {{ items.length }} 个</span>
        </template>
        <template #pagination>
          <USelect v-model="size" :items="pageSizeItems" value-key="value" size="sm" class="w-24" />
          <ManagePagination v-model="page" :total-pages="totalPages" class="!mt-0" />
        </template>
      </ManageCollectionDock>
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
