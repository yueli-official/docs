<script setup lang="ts">
import { useMinLoading } from '@platform/ui/use-min-loading'
import type { CollectionTree, CollectionView, DocTreeNode } from '~/types'
import {
  searchStatusText,
  shouldDimSearchResults,
  shouldShowSearchSkeleton
} from '~/utils/docsSearchState.mjs'

interface SearchFacetsView {
  collections: Array<{ id: string; slug: string; title: string; count: number }>
}

interface SearchDocsResponse {
  items: DocTreeNode[]
  total: number
  facets: SearchFacetsView
}

interface SearchResultView extends DocTreeNode {
  collectionSlug: string
  collectionTitle: string
  path: string
}

const open = defineModel<boolean>('open', { required: true })
const query = ref('')
const activeCollection = ref('__all__')
const pending = ref(false)
const error = ref('')
const results = ref<SearchResultView[]>([])
const total = ref(0)
const facets = ref<SearchFacetsView>({ collections: [] })
const treeCache = reactive<Record<string, CollectionTree>>({})
let searchTimer: ReturnType<typeof setTimeout> | undefined
let searchSeq = 0

const route = useRoute()
const { call } = useApi()
const locale = computed(() => typeof route.query.locale === 'string' ? route.query.locale : 'en')
const version = computed(() => typeof route.query.version === 'string' ? route.query.version : '')
const routeQuery = computed(() => ({
  ...(locale.value !== 'en' ? { locale: locale.value } : {}),
  ...(version.value ? { version: version.value } : {}),
}))
const { data: collectionsData } = await useAsyncData(
  'docs-search-collections',
  () => call<{ items: CollectionView[] }>('/api/v1/collections'),
  { server: false, default: () => ({ items: [] as CollectionView[] }) },
)

const collections = computed(() => collectionsData.value?.items ?? [])
const collectionById = computed(() => Object.fromEntries(collections.value.map(item => [item.id, item])))
const collectionItems = computed(() => [
  { label: '全部文档集', value: '__all__' },
  ...collections.value.map(item => ({ label: item.title, value: item.slug })),
])
const hasQuery = computed(() => query.value.trim().length >= 2)
const showSkeleton = useMinLoading(computed(() => shouldShowSearchSkeleton({
  pending: pending.value,
  hasQuery: hasQuery.value,
  resultCount: results.value.length,
})), 450)
const dimResults = computed(() => shouldDimSearchResults({
  pending: pending.value,
  resultCount: results.value.length,
}))
const statusText = computed(() => searchStatusText({
  hasQuery: hasQuery.value,
  pending: pending.value,
  total: total.value,
}))
const emptyText = computed(() => {
  if (!query.value.trim()) return '输入关键词搜索标题、摘要和正文。'
  if (query.value.trim().length < 2) return '至少输入 2 个字符。'
  return '没有匹配的公开文档。'
})

function selectedValue(value: unknown) {
  if (value && typeof value === 'object' && 'value' in value) {
    const option = value as { value?: unknown }
    return String(option.value ?? '')
  }
  return String(value ?? '')
}

function collectionSlugFor(doc: DocTreeNode) {
  return collectionById.value[doc.collectionId]?.slug
    ?? facets.value.collections.find(item => item.id === doc.collectionId)?.slug
    ?? ''
}

function collectionTitleFor(doc: DocTreeNode) {
  return collectionById.value[doc.collectionId]?.title
    ?? facets.value.collections.find(item => item.id === doc.collectionId)?.title
    ?? '文档集'
}

function treePathIndex(tree: DocTreeNode[], collectionSlug: string) {
  const map = new Map<string, string>()
  const walk = (nodes: DocTreeNode[], base = '') => {
    for (const node of nodes) {
      const path = base ? `${base}/${node.slug}` : node.slug
      map.set(node.id, path)
      if (node.children?.length) walk(node.children, path)
    }
  }
  walk(tree)
  return { collectionSlug, map }
}

async function ensureTrees(slugs: string[]) {
  const unique = Array.from(new Set(slugs.filter(Boolean)))
  await Promise.all(unique.map(async (slug) => {
    const cacheKey = `${slug}:${locale.value}:${version.value || 'default'}`
    if (treeCache[cacheKey]) return
    treeCache[cacheKey] = await call<CollectionTree>(`/api/v1/collections/${slug}/tree`, {
      query: {
        locale: locale.value,
        ...(version.value ? { version: version.value } : {}),
      },
    })
  }))
}

function decorateResults(items: DocTreeNode[]) {
  const indexes = Object.fromEntries(
    Object.entries(treeCache).map(([cacheKey, tree]) => {
      const slug = cacheKey.split(':')[0] || ''
      return [slug, treePathIndex(tree.tree ?? [], slug)]
    }),
  )
  return items.map((doc) => {
    const collectionSlug = collectionSlugFor(doc)
    const index = collectionSlug ? indexes[collectionSlug] : undefined
    return {
      ...doc,
      collectionSlug,
      collectionTitle: collectionTitleFor(doc),
      path: index?.map.get(doc.id) ?? doc.slug,
    }
  })
}

async function runSearch() {
  const q = query.value.trim()
  const seq = ++searchSeq
  error.value = ''
  if (q.length < 2) {
    results.value = []
    total.value = 0
    facets.value = { collections: [] }
    pending.value = false
    return
  }

  pending.value = true
  try {
    const collection = activeCollection.value === '__all__' ? '' : activeCollection.value
    const res = await call<SearchDocsResponse>('/api/v1/docs/search', {
      query: {
        q,
        collection,
        locale: locale.value,
        ...(collection && version.value ? { version: version.value } : {}),
      },
    })
    if (seq !== searchSeq) return
    facets.value = res.facets ?? { collections: [] }
    total.value = res.total ?? res.items.length
    const slugs = res.facets?.collections?.map(item => item.slug) ?? []
    await ensureTrees(slugs)
    if (seq !== searchSeq) return
    results.value = decorateResults(res.items ?? [])
  }
  catch (err: any) {
    if (seq !== searchSeq) return
    error.value = err?.data?.message || '搜索失败，请稍后重试'
    results.value = []
    total.value = 0
  }
  finally {
    if (seq === searchSeq) pending.value = false
  }
}

function scheduleSearch() {
  if (searchTimer) clearTimeout(searchTimer)
  searchTimer = setTimeout(() => { runSearch() }, 220)
}

function resultTo(result: SearchResultView) {
  return result.collectionSlug && result.path
    ? { path: `/${result.collectionSlug}/${result.path}`, query: routeQuery.value }
    : { path: '/', query: routeQuery.value }
}

watch([query, activeCollection, locale, version], scheduleSearch)
watch(open, (value) => {
  if (value) nextTick(() => scheduleSearch())
})

defineShortcuts({
  meta_k: () => { open.value = true },
  ctrl_k: () => { open.value = true },
})
</script>

<template>
  <UModal
    v-model:open="open"
    title="搜索文档"
    description="搜索所有公开文档。"
    :ui="{ content: 'sm:max-w-3xl lg:max-w-4xl' }"
  >
    <template #body>
      <div class="space-y-4">
        <div class="grid gap-3 sm:grid-cols-[minmax(0,1fr)_220px]">
          <UInput
            v-model="query"
            icon="i-tabler-search"
            size="lg"
            autofocus
            placeholder="搜索标题、摘要或正文"
            class="min-w-0"
          />
          <USelectMenu
            :model-value="activeCollection"
            :items="collectionItems"
            value-key="value"
            :search-input="{ placeholder: '筛选文档集…' }"
            class="w-full"
            @update:model-value="activeCollection = selectedValue($event)"
          />
        </div>

        <div class="flex flex-wrap items-center justify-between gap-3 text-sm">
          <p class="flex items-center gap-2 text-muted">
            <UIcon v-if="pending" name="i-tabler-loader-2" class="size-4 animate-spin text-primary" />
            <span>{{ statusText }}</span>
          </p>
          <div class="hidden items-center gap-1 text-xs text-muted sm:flex">
            <UKbd value="Ctrl" />
            <UKbd value="K" />
          </div>
        </div>

        <UAlert
          v-if="error"
          color="error"
          variant="soft"
          icon="i-tabler-alert-circle"
          :title="error"
        />

        <div class="min-h-[360px] md:min-h-[440px]">
          <div v-if="showSkeleton" class="space-y-2">
            <div v-for="i in 7" :key="i" class="rounded-lg border border-default p-3">
              <div class="flex items-start gap-3">
                <USkeleton class="size-10 shrink-0 rounded-lg" />
                <div class="min-w-0 flex-1 space-y-2">
                  <USkeleton class="h-4 w-2/3" />
                  <USkeleton class="h-3 w-full" />
                  <USkeleton class="h-3 w-1/2" />
                </div>
              </div>
            </div>
          </div>

          <div v-else-if="results.length" class="relative max-h-[58dvh] space-y-2 overflow-y-auto pr-1 transition-opacity" :class="dimResults ? 'opacity-70' : 'opacity-100'">
            <NuxtLink
              v-for="item in results"
              :key="item.id"
              :to="resultTo(item)"
              class="group block rounded-lg border border-default p-3 transition hover:border-primary/40 hover:bg-elevated/60 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary"
              @click="open = false"
            >
              <div class="flex items-start gap-3">
                <span class="mt-0.5 grid size-10 shrink-0 place-items-center rounded-lg bg-primary/10 text-primary">
                  <UIcon name="i-tabler-file-text" class="size-5" />
                </span>
                <span class="min-w-0 flex-1">
                  <span class="flex min-w-0 flex-wrap items-center gap-2">
                    <span class="line-clamp-1 font-medium text-highlighted group-hover:text-primary">{{ item.title }}</span>
                    <UBadge :label="item.collectionTitle" color="neutral" variant="subtle" size="sm" />
                  </span>
                  <span v-if="item.excerpt" class="mt-1 line-clamp-2 text-sm leading-6 text-muted">{{ item.excerpt }}</span>
                  <span class="mt-2 block truncate font-mono text-xs text-muted">/{{ item.collectionSlug }}/{{ item.path }}</span>
                </span>
                <UIcon name="i-tabler-arrow-right" class="mt-2 size-4 shrink-0 text-muted transition group-hover:translate-x-0.5 group-hover:text-primary" />
              </div>
            </NuxtLink>
          </div>

          <div v-else class="grid min-h-[360px] place-items-center rounded-lg border border-dashed border-default px-4 py-8 text-center md:min-h-[440px]">
            <div>
              <UIcon name="i-tabler-search" class="mx-auto size-8 text-muted" />
              <p class="mt-3 text-sm font-medium text-highlighted">{{ emptyText }}</p>
            </div>
          </div>
        </div>
      </div>
    </template>
  </UModal>
</template>
