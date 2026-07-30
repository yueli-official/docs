<script setup lang="ts">
import { SkeletonCards } from '~/utils/manageComponents'
import { useMinimumLoading } from '@yueli/ui/feedback'
import type { CollectionList } from '~/types'
import { collectionCountLabel, collectionStats, filterCollections, paginateItems } from '~/utils/docsManual.mjs'

definePageMeta({ width: 'full' })

const { call } = useApi()
const { brand: siteBrand } = useSiteRuntime()
const { data, pending, error, refresh } = await useAsyncData(
  'collections-directory',
  () => call<CollectionList>('/api/v1/collections'),
  { default: () => ({ items: [] }) },
)

const collections = computed(() => data.value?.items ?? [])
const stats = computed(() => collectionStats(collections.value))
const query = ref('')
const page = ref(1)
const pageSize = ref(18)
const pageSizeItems = [
  { label: '每页 12 个', value: 12 },
  { label: '每页 18 个', value: 18 },
  { label: '每页 24 个', value: 24 },
]
const failedCovers = ref<Record<string, boolean>>({})
const showSkeleton = useMinimumLoading(pending)

const filteredCollections = computed(() => filterCollections(collections.value, query.value))
const pageResult = computed(() => paginateItems(filteredCollections.value, page.value, pageSize.value))
const pagedCollections = computed(() => pageResult.value.items)
const hasQuery = computed(() => query.value.trim().length > 0)
const hasLoadError = computed(() => Boolean(error.value && !collections.value.length))

function coverVisible(c: { id: string, coverUrl?: string }) {
  return Boolean(c.coverUrl && !failedCovers.value[c.id])
}

function markCoverFailed(id: string) {
  failedCovers.value = { ...failedCovers.value, [id]: true }
}

function clearQuery() {
  query.value = ''
}

watch([query, pageSize], () => { page.value = 1 })
watch(() => pageResult.value.page, (safePage) => { page.value = safePage })

useSeoMeta({ title: '全部文档集', description: '按标题、路径和说明查找全部公开文档集' })
</script>

<template>
  <div class="min-w-0 space-y-8 overflow-x-hidden">
    <header class="grid gap-6 border-b border-default pb-8 lg:grid-cols-[minmax(0,1fr)_360px] lg:items-end">
      <div class="min-w-0">
        <nav class="mb-5 flex items-center gap-2 text-sm text-muted" aria-label="面包屑">
          <NuxtLink to="/" class="inline-flex items-center gap-1 hover:text-primary">
            <UIcon name="i-tabler-home" class="size-4" />
            {{ siteBrand }}
          </NuxtLink>
          <UIcon name="i-tabler-chevron-right" class="size-4" />
          <span class="text-default">全部文档集</span>
        </nav>
        <p class="mb-3 inline-flex items-center gap-2 rounded-full bg-primary/10 px-3 py-1 text-xs font-medium text-primary">
          <UIcon name="i-tabler-list-search" class="size-4" />
          Collection directory
        </p>
        <h1 class="font-display text-balance text-4xl font-bold leading-tight tracking-tight text-highlighted sm:text-5xl">全部文档集</h1>
        <p class="mt-4 max-w-[66ch] break-words text-base leading-7 text-muted">
          按标题、路径或说明查找文档集。适合文档数量变多后快速定位入口。
        </p>
      </div>

      <dl class="grid grid-cols-2 gap-3">
        <div class="rounded-lg border border-default bg-elevated/30 px-3 py-2">
          <dt class="text-xs text-muted">文档集</dt>
          <dd class="font-display mt-1 text-xl font-semibold text-highlighted">{{ hasLoadError ? '-' : stats.collectionCount }}</dd>
        </div>
        <div class="rounded-lg border border-default bg-elevated/30 px-3 py-2">
          <dt class="text-xs text-muted">公开文档</dt>
          <dd class="font-display mt-1 text-xl font-semibold text-highlighted">{{ hasLoadError ? '-' : stats.docCount }}</dd>
        </div>
      </dl>
    </header>

    <section class="space-y-5" aria-labelledby="collections-directory-title">
      <div v-if="hasLoadError" class="rounded-lg border border-error/30 bg-error/5 p-5">
        <div class="flex min-w-0 items-start gap-3">
          <UIcon name="i-tabler-alert-triangle" class="mt-0.5 size-5 shrink-0 text-error" />
          <div class="min-w-0">
            <h2 class="text-sm font-semibold text-highlighted">文档集加载失败</h2>
            <p class="mt-1 text-sm text-muted">请稍后重试，或检查服务是否正常运行。</p>
            <UButton class="mt-4" color="neutral" variant="soft" icon="i-tabler-refresh" label="重新加载" @click="refresh()" />
          </div>
        </div>
      </div>

      <template v-else>
        <div class="grid gap-3 rounded-lg border border-default bg-default p-3 lg:grid-cols-[minmax(0,1fr)_180px_auto] lg:items-end">
          <UFormField label="查找文档集">
            <UInput
              v-model="query"
              icon="i-tabler-search"
              size="lg"
              placeholder="输入标题、路径或说明"
              class="w-full"
            />
          </UFormField>
          <UFormField label="分页">
            <USelect v-model="pageSize" :items="pageSizeItems" value-key="value" class="w-full" />
          </UFormField>
          <UButton
            color="neutral"
            variant="ghost"
            icon="i-tabler-x"
            label="清除"
            :disabled="!hasQuery"
            @click="clearQuery"
          />
        </div>

        <div class="flex min-h-9 flex-col gap-2 sm:flex-row sm:items-center sm:justify-between">
          <div>
            <h2 id="collections-directory-title" class="font-display text-xl font-semibold text-highlighted">文档集目录</h2>
            <p class="mt-1 text-sm text-muted">
              <span v-if="pageResult.total">显示 {{ pageResult.start }}-{{ pageResult.end }} / {{ pageResult.total }}</span>
              <span v-else>没有匹配项</span>
            </p>
          </div>
          <UPagination
            v-if="pageResult.total > pageResult.pageSize"
            v-model:page="page"
            :total="pageResult.total"
            :items-per-page="pageResult.pageSize"
          />
        </div>

        <SkeletonCards v-if="showSkeleton" :count="6" />

        <div v-else-if="pagedCollections.length" class="grid gap-4 md:grid-cols-2 xl:grid-cols-3">
          <NuxtLink
            v-for="c in pagedCollections"
            :key="c.id"
            :to="`/${c.slug}`"
            class="group flex min-w-0 flex-col overflow-hidden rounded-xl border border-default bg-default transition duration-200 hover:border-primary/40 hover:shadow-sm focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary"
          >
            <div class="relative aspect-[16/7] border-b border-default bg-elevated">
              <img v-if="coverVisible(c)" :src="collectionCoverThumbUrl(c)" :alt="c.title" class="size-full object-cover" @error="markCoverFailed(c.id)" >
              <div v-else class="grid size-full place-items-center bg-elevated">
                <span class="grid size-11 place-items-center rounded-lg bg-default text-primary ring-1 ring-default">
                  <UIcon :name="c.icon || 'i-tabler-book-2'" class="size-6" />
                </span>
              </div>
              <span class="absolute right-3 top-3 rounded-full bg-default/90 px-2.5 py-1 text-xs font-medium text-muted ring-1 ring-default backdrop-blur">
                {{ collectionCountLabel(c.docCount) }}
              </span>
            </div>
            <div class="flex flex-1 flex-col p-4">
              <div class="flex min-w-0 items-start gap-3">
                <span class="mt-0.5 grid size-8 shrink-0 place-items-center rounded-md bg-primary/10 text-primary">
                  <UIcon :name="c.icon || 'i-tabler-book-2'" class="size-4" />
                </span>
                <div class="min-w-0">
                  <h3 class="font-display text-base font-semibold leading-snug text-highlighted transition group-hover:text-primary">{{ c.title }}</h3>
                  <p class="mt-1 font-mono text-xs text-muted">/{{ c.slug }}</p>
                  <p v-if="c.description" class="mt-2 line-clamp-2 text-sm leading-6 text-muted">{{ c.description }}</p>
                </div>
              </div>
              <span class="mt-4 inline-flex items-center gap-1 text-sm font-medium text-primary">
                查看目录
                <UIcon name="i-tabler-arrow-right" class="size-4 transition group-hover:translate-x-0.5" />
              </span>
            </div>
          </NuxtLink>
        </div>

        <div v-else class="rounded-2xl border border-dashed border-default p-10 text-center">
          <UIcon name="i-tabler-search-off" class="mx-auto size-10 text-muted" />
          <h3 class="mt-4 font-display text-lg font-semibold text-highlighted">没有匹配的文档集</h3>
          <p class="mt-2 text-sm text-muted">换一个关键词，或清除筛选后浏览全部集合。</p>
          <UButton class="mt-5" color="neutral" variant="soft" icon="i-tabler-x" label="清除筛选" @click="clearQuery" />
        </div>

        <div v-if="pageResult.total > pageResult.pageSize" class="flex flex-col gap-3 border-t border-default pt-5 sm:flex-row sm:items-center sm:justify-between">
          <p class="text-sm text-muted">第 {{ pageResult.page }} / {{ pageResult.totalPages }} 页</p>
          <UPagination v-model:page="page" :total="pageResult.total" :items-per-page="pageResult.pageSize" />
        </div>
      </template>
    </section>
  </div>
</template>
