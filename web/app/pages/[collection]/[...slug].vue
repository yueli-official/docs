<script setup lang="ts">
import type { DocDetailResponse } from '~/types'

definePageMeta({ layout: 'collection' })
const route = useRoute()
const { brand: siteBrand } = useSiteRuntime()
const collectionSlug = computed(() => route.params.collection as string)
const locale = computed(() => typeof route.query.locale === 'string' ? route.query.locale : 'en')
const version = computed(() => typeof route.query.version === 'string' ? route.query.version : '')
const routeQuery = computed(() => ({
  ...(locale.value !== 'en' ? { locale: locale.value } : {}),
  ...(version.value ? { version: version.value } : {}),
}))
function docTo(path = '') {
  return { path: path ? `/${collectionSlug.value}/${path}` : `/${collectionSlug.value}`, query: routeQuery.value }
}
const docPath = computed(() => {
  const rest = route.params.slug
  return Array.isArray(rest) ? rest.join('/') : String(rest ?? '')
})

const { ensure, byPath, flat, collection } = useCollectionTree(() => route.params.collection as string, locale, version)
await ensure()

const node = computed(() => byPath.value.get(docPath.value) ?? null)
if (!node.value) throw createError({ statusCode: 404, statusMessage: '文档不存在', fatal: true })

const { call } = useApi()
const { data: docData, pending: docPending } = await useAsyncData(
  () => `doc-${collectionSlug.value}-${locale.value}-${version.value || 'default'}-${docPath.value}`,
  () => call<DocDetailResponse>('/api/v1/docs/by-path', {
    query: {
      collection: collectionSlug.value,
      path: docPath.value,
      locale: locale.value,
      ...(version.value ? { version: version.value } : {}),
    },
  }),
  { watch: [collectionSlug, docPath, locale, version] },
)
const doc = computed(() => docData.value?.doc ?? null)

// prev/next = adjacent entries in DFS pre-order (whole flattened list).
const idx = computed(() => flat.value.findIndex(e => e.path === docPath.value))
const prev = computed(() => idx.value > 0 ? flat.value[idx.value - 1] : null)
const next = computed(() => idx.value >= 0 && idx.value < flat.value.length - 1 ? flat.value[idx.value + 1] : null)
const breadcrumbs = computed(() => {
  const parts = docPath.value.split('/').filter(Boolean)
  return parts.map((_, index) => {
    const path = parts.slice(0, index + 1).join('/')
    const entry = flat.value.find(e => e.path === path)
    return { path, title: entry?.node.title || parts[index] }
  })
})

function childPath(slug: string) {
  return [docPath.value, slug].filter(Boolean).join('/')
}

const { renderWithToc } = useMarkdown()
function stripDuplicateTitle(content: string, title: string) {
  const trimmedTitle = title.trim()
  if (!trimmedTitle) return content
  const escaped = trimmedTitle.replace(/[.*+?^${}()|[\]\\]/g, '\\$&')
  return content.replace(new RegExp(`^#\\s+${escaped}\\s*\\n+`), '')
}

const readingContent = computed(() => stripDuplicateTitle(doc.value?.content ?? '', node.value?.title ?? ''))
const toc = computed(() => renderWithToc(readingContent.value).toc)
const isLeaf = computed(() => !node.value?.children?.length)
const sectionLabel = computed(() => toc.value.length ? `${toc.value.length} 个小节` : (isLeaf.value ? '正文' : '章节'))
const readingProgress = ref(0)

let progressFrame = 0
function updateReadingProgress() {
  if (progressFrame) return
  progressFrame = requestAnimationFrame(() => {
    progressFrame = 0
    const article = document.querySelector<HTMLElement>('.reading-shell article')
    if (!article) return
    const start = article.offsetTop
    const total = Math.max(1, article.scrollHeight - window.innerHeight * 0.7)
    const scrolled = Math.min(Math.max(window.scrollY - start + 96, 0), total)
    readingProgress.value = Math.round((scrolled / total) * 100)
  })
}

onMounted(() => {
  updateReadingProgress()
  window.addEventListener('scroll', updateReadingProgress, { passive: true })
  window.addEventListener('resize', updateReadingProgress)
})
onBeforeUnmount(() => {
  window.removeEventListener('scroll', updateReadingProgress)
  window.removeEventListener('resize', updateReadingProgress)
  if (progressFrame) cancelAnimationFrame(progressFrame)
})

useSeoMeta({
  title: () => doc.value?.title || node.value?.title || docPath.value,
  description: () => doc.value?.excerpt || collection.value?.title || undefined
})
</script>

<template>
  <div v-if="node" class="reading-shell">
    <nav class="mb-6 flex flex-wrap items-center gap-2 text-sm text-muted" aria-label="面包屑">
      <NuxtLink to="/" class="hover:text-primary">{{ siteBrand }}</NuxtLink>
      <UIcon name="i-tabler-chevron-right" class="size-4" />
      <NuxtLink :to="docTo()" class="hover:text-primary">{{ collection?.title }}</NuxtLink>
      <template v-for="crumb in breadcrumbs.slice(0, -1)" :key="crumb.path">
        <UIcon name="i-tabler-chevron-right" class="size-4" />
        <NuxtLink :to="docTo(crumb.path)" class="line-clamp-1 max-w-48 hover:text-primary">{{ crumb.title }}</NuxtLink>
      </template>
    </nav>

    <div class="grid gap-10 xl:grid-cols-[minmax(0,1fr)_280px]">
      <article class="min-w-0">
        <header class="border-b border-default pb-8">
          <div class="flex flex-wrap items-center gap-2 text-xs font-medium text-muted">
            <span class="inline-flex items-center gap-1 rounded-md bg-primary/10 px-2 py-1 text-primary">
              <UIcon :name="collection?.icon || 'i-tabler-book-2'" class="size-3.5" />
              {{ collection?.title }}
            </span>
            <span class="inline-flex items-center gap-1 rounded-md bg-elevated px-2 py-1">
              <UIcon :name="isLeaf ? 'i-tabler-file-text' : 'i-tabler-folder'" class="size-3.5" />
              {{ isLeaf ? '文档' : '章节' }}
            </span>
          </div>
          <h1 class="font-display mt-4 text-balance text-[2rem] font-semibold leading-[1.18] text-highlighted md:text-[2.3125rem]">{{ node.title }}</h1>
          <p v-if="node.excerpt" class="mt-5 max-w-3xl text-lg leading-8 text-muted">{{ node.excerpt }}</p>
          <div class="mt-6 flex flex-wrap items-center gap-3 text-sm text-muted">
            <span class="inline-flex items-center gap-1.5">
              <UIcon name="i-tabler-list-details" class="size-4 text-primary" />
              {{ sectionLabel }}
            </span>
          </div>
        </header>

        <div class="mt-8">
          <ContentProse v-if="readingContent" :content="readingContent" />
          <div v-else-if="docPending" class="space-y-3">
            <USkeleton class="h-6 w-2/3" />
            <USkeleton class="h-4 w-full" />
            <USkeleton class="h-4 w-11/12" />
            <USkeleton class="h-4 w-4/5" />
            <USkeleton class="h-4 w-10/12" />
          </div>
        </div>

        <section v-if="!isLeaf && node.children" class="mt-10">
          <div class="mb-4 flex items-center justify-between gap-3">
            <h2 class="text-lg font-semibold text-highlighted">章节内容</h2>
            <span class="text-sm text-muted">{{ node.children.length }} 篇</span>
          </div>
          <div class="grid gap-3 sm:grid-cols-2">
            <NuxtLink
              v-for="c in node.children"
              :key="c.id"
              :to="docTo(childPath(c.slug))"
              class="group rounded-lg border border-default bg-default p-4 transition hover:border-primary/40 hover:bg-elevated/50"
            >
              <span class="flex items-center gap-2 text-sm font-semibold text-highlighted transition group-hover:text-primary">
                <UIcon name="i-tabler-file-text" class="size-4" />
                {{ c.title }}
              </span>
              <span v-if="c.excerpt" class="mt-2 line-clamp-2 text-sm leading-6 text-muted">{{ c.excerpt }}</span>
            </NuxtLink>
          </div>
        </section>

        <DocNav
          :prev="prev ? { path: prev.path, title: prev.node.title } : undefined"
          :next="next ? { path: next.path, title: next.node.title } : undefined"
          :collection="collectionSlug"
          :query="routeQuery"
        />
      </article>

      <aside class="hidden xl:block">
        <div class="sticky top-24 space-y-4">
          <section class="rounded-lg border border-default bg-default p-4">
            <div class="mb-3 flex items-center justify-between gap-3">
              <h2 class="text-sm font-semibold text-highlighted">本页目录</h2>
              <UIcon name="i-tabler-list-details" class="size-4 text-muted" />
            </div>
            <TableOfContents v-if="toc.length" :items="toc" :show-title="false" />
            <p v-else class="text-sm leading-6 text-muted">这页没有可跳转的小节。</p>
          </section>

          <section class="reading-progress-card rounded-lg border border-default bg-default p-4">
            <div class="flex items-center justify-between gap-3">
              <span class="flex items-center gap-2 text-sm font-semibold text-highlighted">
                <UIcon name="i-tabler-progress" class="size-4 text-primary" />
                阅读进度
              </span>
              <span class="font-mono text-sm text-primary">{{ readingProgress }}%</span>
            </div>
            <div class="mt-3 h-1.5 overflow-hidden rounded-full bg-elevated">
              <div class="h-full rounded-full bg-primary transition-[width]" :style="{ width: `${readingProgress}%` }" />
            </div>
            <p class="mt-2 text-xs text-muted">{{ sectionLabel }}</p>
          </section>

          <NuxtLink
            v-if="next"
            :to="docTo(next.path)"
            class="group block rounded-lg border border-primary/20 bg-primary/5 p-4 transition hover:border-primary/40 hover:bg-primary/10"
          >
            <span class="flex items-center gap-2 text-xs font-medium text-primary">
              <UIcon name="i-tabler-arrow-right" class="size-4" />
              下一篇
            </span>
            <span class="mt-2 line-clamp-2 text-sm font-semibold text-highlighted transition group-hover:text-primary">{{ next.node.title }}</span>
          </NuxtLink>
        </div>
      </aside>
    </div>
  </div>
</template>
