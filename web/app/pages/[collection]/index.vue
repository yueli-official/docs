<script setup lang="ts">
import { rootSectionCountLabel } from '~/utils/docsManual.mjs'

definePageMeta({ layout: 'collection' })
const route = useRoute()
const { brand: siteBrand } = useSiteRuntime()
const slug = computed(() => route.params.collection as string)
const locale = computed(() => typeof route.query.locale === 'string' ? route.query.locale : 'en')
const version = computed(() => typeof route.query.version === 'string' ? route.query.version : '')
const routeQuery = computed(() => ({
  ...(locale.value !== 'en' ? { locale: locale.value } : {}),
  ...(version.value ? { version: version.value } : {}),
}))
function docTo(path: string) {
  return { path: `/${slug.value}/${path}`, query: routeQuery.value }
}
// NOTE (Task 3 fix): useCollectionTree takes a GETTER so it re-keys reactively
// across collections. Pass `() => route.params.collection as string`, not a snapshot.
const { ensure, tree, collection } = useCollectionTree(() => route.params.collection as string, locale, version)
await ensure()
if (!collection.value) throw createError({ statusCode: 404, statusMessage: '文档集不存在', fatal: true })
const startNodes = computed(() => tree.value.slice(0, 4))
const showStartNodes = computed(() => tree.value.length > startNodes.value.length)
useSeoMeta({ title: () => collection.value?.title || slug.value, description: () => collection.value?.description || undefined })
</script>

<template>
  <div class="min-w-0 space-y-10 overflow-x-hidden">
    <header class="border-b border-default pb-8">
      <nav class="mb-5 flex items-center gap-2 text-sm text-muted" aria-label="面包屑">
        <NuxtLink to="/" class="inline-flex items-center gap-1 hover:text-primary">
          <UIcon name="i-tabler-home" class="size-4" />
          {{ siteBrand }}
        </NuxtLink>
        <UIcon name="i-tabler-chevron-right" class="size-4" />
        <span class="line-clamp-1 text-default">{{ collection?.title }}</span>
      </nav>

      <div class="min-w-0 max-w-[72ch]">
        <p class="mb-3 inline-flex items-center gap-2 rounded-full bg-primary/10 px-3 py-1 text-xs font-medium text-primary">
          <UIcon :name="collection?.icon || 'i-tabler-book-2'" class="size-4" />
          文档集
        </p>
        <h1 class="font-display text-balance text-3xl font-bold leading-tight tracking-tight text-highlighted sm:text-[2.25rem]">{{ collection?.title }}</h1>
        <p v-if="collection?.description" class="mt-4 max-w-2xl break-words text-base leading-7 text-muted">{{ collection.description }}</p>

        <div class="mt-5 flex min-w-0 flex-wrap items-center gap-x-4 gap-y-2 text-sm text-muted">
          <span class="inline-flex min-w-0 items-center gap-2">
            <span class="relative size-7 shrink-0 overflow-hidden rounded-md bg-elevated ring-1 ring-default">
              <img v-if="collection?.coverUrl" :src="collection.coverUrl" :alt="collection.title" class="size-full object-cover" >
              <span v-else class="grid size-full place-items-center text-primary">
                <UIcon :name="collection?.icon || 'i-tabler-book-2'" class="size-4" />
              </span>
            </span>
            <span class="font-medium text-default">{{ rootSectionCountLabel(tree.length) }}</span>
          </span>
          <span class="hidden h-4 w-px bg-border sm:block" />
          <span>按章节组织，适合连续阅读和快速定位。</span>
        </div>
      </div>
    </header>

    <section v-if="showStartNodes" class="space-y-4">
      <div class="flex items-center gap-2">
        <UIcon name="i-tabler-player-play" class="size-5 text-primary" />
        <h2 class="font-display text-xl font-semibold text-highlighted">从这里开始</h2>
      </div>
      <div class="grid gap-3 sm:grid-cols-2">
        <NuxtLink
          v-for="n in startNodes"
          :key="`start-${n.id}`"
          :to="docTo(n.slug)"
          class="group flex min-w-0 items-start gap-3 rounded-lg border border-default bg-default p-4 transition duration-200 hover:border-primary/40 hover:bg-elevated/40"
        >
          <span class="mt-0.5 grid size-8 shrink-0 place-items-center rounded-lg bg-primary/10 text-primary">
            <UIcon name="i-tabler-file-text" class="size-4" />
          </span>
          <span class="min-w-0">
            <span class="line-clamp-2 font-medium text-highlighted transition group-hover:text-primary">{{ n.title }}</span>
            <span v-if="n.excerpt" class="mt-1 line-clamp-2 text-sm text-muted">{{ n.excerpt }}</span>
            <span v-else-if="n.children?.length" class="mt-1 block text-sm text-muted">{{ n.children.length }} 个子章节</span>
          </span>
        </NuxtLink>
      </div>
    </section>

    <section class="space-y-4">
      <div class="flex items-center gap-2">
        <UIcon name="i-tabler-list-tree" class="size-5 text-primary" />
        <h2 class="font-display text-xl font-semibold text-highlighted">章节目录</h2>
      </div>

      <div v-if="tree.length" class="grid gap-3 sm:grid-cols-2">
        <NuxtLink
          v-for="n in tree"
          :key="n.id"
          :to="docTo(n.slug)"
          class="group rounded-lg border border-default p-4 transition duration-200 hover:border-primary/40 hover:bg-elevated/40"
        >
          <div class="flex min-w-0 items-start justify-between gap-3">
            <div class="min-w-0">
              <h3 class="line-clamp-2 font-medium text-highlighted transition group-hover:text-primary">{{ n.title }}</h3>
              <p v-if="n.excerpt" class="mt-1 line-clamp-2 text-sm leading-6 text-muted">{{ n.excerpt }}</p>
            </div>
            <UIcon name="i-tabler-chevron-right" class="mt-1 size-4 shrink-0 text-muted transition group-hover:translate-x-0.5 group-hover:text-primary" />
          </div>
          <p v-if="n.children?.length" class="mt-3 text-xs text-muted">{{ n.children.length }} 个子章节</p>
        </NuxtLink>
      </div>

      <div v-else class="rounded-2xl border border-dashed border-default px-4 py-10 text-center sm:p-10">
        <UIcon name="i-tabler-file-off" class="mx-auto size-10 text-muted" />
        <h2 class="mt-4 font-display text-lg font-semibold text-highlighted">这个文档集还没有公开章节</h2>
        <p class="mt-2 text-sm text-muted">内容发布后会显示在这里。</p>
        <UButton class="mt-5" to="/" color="neutral" variant="soft" icon="i-tabler-arrow-left" :label="`返回${siteBrand}`" />
      </div>
    </section>
  </div>
</template>
