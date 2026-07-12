<script setup lang="ts">
import type { CollectionList, HomeConfigResponse } from '~/types'
import {
  collectionCountLabel,
  collectionStats,
  homeFeaturedCollections,
  homeQuickLinks
} from '~/utils/docsManual.mjs'

definePageMeta({ width: 'full' })
const { call } = useApi()
const { data } = await useAsyncData('collections', () => call<CollectionList>('/api/v1/collections'))
const { data: homeData } = await useAsyncData(
  'docs-home-config',
  () => call<HomeConfigResponse>('/api/v1/home'),
)
if (!homeData.value?.config) throw createError({ statusCode: 500, statusMessage: '文档站点配置尚未初始化' })
const collections = computed(() => data.value?.items ?? [])
const homeConfig = computed(() => homeData.value!.config)
const stats = computed(() => collectionStats(collections.value))
const taskLinks = computed(() => homeQuickLinks(homeConfig.value, collections.value))
const searchOpen = ref(false)
const failedCovers = ref<Record<string, boolean>>({})
const previewCollections = computed(() => collections.value.slice(0, 6))
const featuredCollections = computed(() => homeFeaturedCollections(collections.value, homeConfig.value, 6))
const startPath = computed(() => taskLinks.value[0]?.to || (featuredCollections.value[0] ? `/${featuredCollections.value[0].slug}` : '/'))
function coverVisible(c: { id: string, coverUrl?: string }) {
  return Boolean(c.coverUrl && !failedCovers.value[c.id])
}
function markCoverFailed(id: string) {
  failedCovers.value = { ...failedCovers.value, [id]: true }
}
function sweepBrokenCovers() {
  if (!import.meta.client) return
  nextTick(() => {
    document.querySelectorAll<HTMLImageElement>('img[data-cover-id]').forEach((img) => {
      if (img.complete && img.naturalWidth === 0) markCoverFailed(img.dataset.coverId || '')
    })
  })
}
watch(collections, sweepBrokenCovers, { immediate: true })
onMounted(sweepBrokenCovers)
useSeoMeta({
  title: () => homeConfig.value.siteTitle,
  description: () => homeConfig.value.siteDescription,
})
</script>

<template>
  <div class="min-w-0 space-y-10 overflow-x-hidden">
    <header class="grid min-w-0 gap-6 border-b border-default pb-8 lg:grid-cols-[minmax(0,1fr)_360px] lg:items-end">
      <div class="min-w-0 max-w-3xl">
        <p class="mb-3 inline-flex items-center gap-2 rounded-full bg-primary/10 px-3 py-1 text-xs font-medium text-primary">
          <UIcon name="i-tabler-book-2" class="size-4" />
          {{ homeConfig.homeEyebrow }}
        </p>
        <h1 class="font-display text-balance text-4xl font-bold leading-tight tracking-tight text-highlighted sm:text-5xl">{{ homeConfig.homeTitle }}</h1>
        <p class="mt-4 max-w-[66ch] break-words text-base leading-7 text-muted">
          {{ homeConfig.homeSubtitle }}
        </p>

        <div class="mt-6 flex min-w-0 flex-col gap-3 sm:flex-row">
          <UButton
            color="neutral"
            variant="outline"
            size="xl"
            icon="i-tabler-search"
            class="min-h-12 min-w-0 flex-1 justify-start rounded-lg bg-default text-left text-muted"
            @click="() => { searchOpen = true }"
          >
            <span class="min-w-0 flex-1 truncate text-left">搜索标题、摘要或正文</span>
            <span class="hidden rounded bg-elevated px-1.5 py-0.5 text-xs text-muted sm:inline-flex">Ctrl K</span>
          </UButton>
          <UButton
            v-if="collections.length"
            :to="startPath"
            size="xl"
            icon="i-tabler-player-play"
            label="开始阅读"
          />
        </div>

        <dl class="mt-5 grid max-w-sm grid-cols-2 gap-3">
          <div class="rounded-lg border border-default bg-elevated/30 px-3 py-2">
            <dt class="text-xs text-muted">文档集</dt>
            <dd class="font-display mt-1 text-xl font-semibold text-highlighted">{{ stats.collectionCount }}</dd>
          </div>
          <div class="rounded-lg border border-default bg-elevated/30 px-3 py-2">
            <dt class="text-xs text-muted">公开文档</dt>
            <dd class="font-display mt-1 text-xl font-semibold text-highlighted">{{ stats.docCount }}</dd>
          </div>
        </dl>
      </div>

      <section v-if="taskLinks.length" class="min-w-0 rounded-xl border border-default bg-elevated/40 p-4" aria-labelledby="task-paths-title">
        <div class="mb-4 flex items-center gap-2">
          <UIcon name="i-tabler-route" class="size-5 text-primary" />
          <h2 id="task-paths-title" class="font-display text-lg font-semibold text-highlighted">快速入口</h2>
        </div>
        <div class="space-y-2">
          <NuxtLink
            v-for="task in taskLinks"
            :key="task.title"
            :to="task.to"
            class="group flex min-w-0 items-start gap-3 rounded-lg border border-default bg-default px-3 py-3 transition duration-200 hover:border-primary/40 hover:bg-elevated/60 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary"
          >
            <span class="mt-0.5 grid size-9 shrink-0 place-items-center rounded-lg bg-primary/10 text-primary">
              <UIcon :name="task.icon" class="size-5" />
            </span>
            <span class="min-w-0">
              <span class="block font-medium text-highlighted transition group-hover:text-primary">{{ task.title }}</span>
              <span class="mt-0.5 line-clamp-2 text-sm leading-5 text-muted">{{ task.description }}</span>
            </span>
            <UIcon name="i-tabler-chevron-right" class="ml-auto mt-2 size-4 shrink-0 text-muted transition group-hover:translate-x-0.5 group-hover:text-primary" />
          </NuxtLink>
        </div>
      </section>
    </header>

    <section v-if="featuredCollections.length" class="space-y-4" aria-labelledby="featured-title">
      <div class="flex flex-col gap-2 sm:flex-row sm:items-end sm:justify-between">
        <div>
          <h2 id="featured-title" class="font-display text-2xl font-semibold tracking-tight text-highlighted">推荐文档</h2>
          <p class="mt-1 text-sm text-muted">优先阅读这些文档集，适合连续阅读和快速定位。</p>
        </div>
      </div>

      <div class="grid gap-3 md:grid-cols-2 xl:grid-cols-3">
        <NuxtLink
          v-for="c in featuredCollections"
          :key="`featured-${c.id}`"
          :to="`/${c.slug}`"
          class="group flex min-w-0 items-start gap-3 rounded-xl border border-default bg-default p-4 transition duration-200 hover:border-primary/40 hover:bg-elevated/50 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary"
        >
          <span class="relative size-11 shrink-0 overflow-hidden rounded-lg bg-elevated ring-1 ring-default">
            <img v-if="coverVisible(c)" :src="collectionCoverThumbUrl(c)" :alt="c.title" :data-cover-id="c.id" class="size-full object-cover" @error="markCoverFailed(c.id)" >
            <span v-else class="grid size-full place-items-center text-primary">
              <UIcon :name="c.icon || 'i-tabler-book-2'" class="size-5" />
            </span>
          </span>
          <span class="min-w-0 flex-1">
            <span class="flex min-w-0 flex-wrap items-center gap-2">
              <span class="line-clamp-1 font-display font-semibold text-highlighted transition group-hover:text-primary">{{ c.title }}</span>
              <UBadge :label="collectionCountLabel(c.docCount)" color="neutral" variant="subtle" />
            </span>
            <span v-if="c.description" class="mt-1 line-clamp-2 text-sm leading-6 text-muted">{{ c.description }}</span>
          </span>
          <UIcon name="i-tabler-arrow-right" class="mt-1 size-4 shrink-0 text-muted transition group-hover:translate-x-0.5 group-hover:text-primary" />
        </NuxtLink>
      </div>
    </section>

    <section v-if="collections.length" class="space-y-6" aria-labelledby="collections-title">
      <div class="grid gap-4 border-t border-default pt-8 lg:grid-cols-[minmax(0,1fr)_320px] lg:items-end">
        <div>
          <h2 id="collections-title" class="font-display text-2xl font-semibold tracking-tight text-highlighted">全部文档集</h2>
          <p class="mt-1 text-sm text-muted">
            按产品、能力域或教程集合浏览，完整查找和分页在目录页完成。
          </p>
        </div>
        <div class="flex justify-start lg:justify-end">
          <UButton to="/collections" color="neutral" variant="soft" icon="i-tabler-list-search" label="查看全部文档集" />
        </div>
      </div>

      <div class="grid gap-4 md:grid-cols-2 xl:grid-cols-3">
        <NuxtLink
          v-for="c in previewCollections"
          :key="c.id"
          :to="`/${c.slug}`"
          class="group flex min-w-0 flex-col overflow-hidden rounded-xl border border-default bg-default transition duration-200 hover:border-primary/40 hover:shadow-sm focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary"
        >
          <div class="relative aspect-[16/7] border-b border-default bg-elevated">
            <img v-if="coverVisible(c)" :src="collectionCoverThumbUrl(c)" :alt="c.title" :data-cover-id="c.id" class="size-full object-cover" @error="markCoverFailed(c.id)" >
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
                <p v-if="c.description" class="mt-1 line-clamp-2 text-sm leading-6 text-muted">{{ c.description }}</p>
              </div>
            </div>
            <span class="mt-4 inline-flex items-center gap-1 text-sm font-medium text-primary">
              查看目录
              <UIcon name="i-tabler-arrow-right" class="size-4 transition group-hover:translate-x-0.5" />
            </span>
          </div>
        </NuxtLink>
      </div>
    </section>

    <section v-else class="max-w-3xl py-6">
      <div class="flex min-w-0 items-start gap-4 border-b border-default pb-6">
        <span class="mt-1 grid size-10 shrink-0 place-items-center rounded-lg bg-muted text-muted">
          <UIcon name="i-tabler-file-off" class="size-5" />
        </span>
        <div class="min-w-0">
          <h2 class="font-display text-xl font-semibold tracking-tight text-highlighted">暂无公开文档</h2>
          <p class="mt-2 max-w-[58ch] text-sm leading-6 text-muted">当前没有可浏览的公开文档集。</p>
        </div>
      </div>
    </section>

    <DocsSearchDialog v-model:open="searchOpen" />
  </div>
</template>
