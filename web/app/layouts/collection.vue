<script setup lang="ts">
const route = useRoute()
const slug = computed(() => route.params.collection as string)
const locale = computed(() => typeof route.query.locale === 'string' ? route.query.locale : 'en')
const version = computed(() => typeof route.query.version === 'string' ? route.query.version : '')
const routeQuery = computed(() => ({
  ...(locale.value !== 'en' ? { locale: locale.value } : {}),
  ...(version.value ? { version: version.value } : {}),
}))
const collectionTo = computed(() => ({ path: `/${slug.value}`, query: routeQuery.value }))
const { ensure, tree, collection, pending } = useCollectionTree(() => route.params.collection as string, locale, version)
await ensure()
watch([() => route.params.collection, locale, version], () => ensure())
const navOpen = ref(false)

// current doc path = everything after /{collection}/ (the [...slug] catch-all).
const currentPath = computed(() => {
  const rest = route.params.slug
  return Array.isArray(rest) ? rest.join('/') : (rest ?? '')
})
</script>

<template>
  <div class="flex min-h-dvh flex-col bg-default text-default">
    <SiteHeader />
    <div class="mx-auto flex w-full max-w-screen-xl flex-1 gap-8 px-4 py-8">
      <aside class="hidden w-64 shrink-0 lg:block">
        <div class="sticky top-24 max-h-[calc(100dvh-7rem)] overflow-y-auto pb-8">
          <NuxtLink :to="collectionTo" class="mb-4 flex min-w-0 items-center gap-3 rounded-lg border border-default bg-default p-2 text-highlighted transition hover:border-primary/40 hover:text-primary">
            <span class="relative size-10 shrink-0 overflow-hidden rounded-md bg-elevated">
              <img v-if="collection?.coverUrl" :src="collection.coverUrl" :alt="collection.title" class="size-full object-cover" >
              <span v-else class="grid size-full place-items-center text-primary">
                <UIcon :name="collection?.icon || 'i-tabler-book-2'" class="size-5" />
              </span>
            </span>
            <span class="min-w-0">
              <span class="line-clamp-2 font-display text-sm font-semibold leading-snug">{{ collection?.title || slug }}</span>
              <span class="mt-0.5 block text-xs font-normal text-muted">{{ tree.length }} 个顶级章节</span>
            </span>
          </NuxtLink>
          <div v-if="pending && !tree.length" class="space-y-2">
            <USkeleton v-for="i in 6" :key="i" class="h-4 rounded" :class="i % 3 === 0 ? 'ml-4 w-32' : 'w-44'" />
          </div>
          <DocTree v-else :nodes="tree" :current-path="currentPath" />
        </div>
      </aside>
      <div class="min-w-0 flex-1">
        <div v-if="currentPath" class="mb-5 flex items-center justify-between gap-3 rounded-xl border border-default bg-elevated/35 px-3 py-2 lg:hidden">
          <div class="flex min-w-0 items-center gap-3">
            <span class="relative size-10 shrink-0 overflow-hidden rounded-md bg-elevated">
              <img v-if="collection?.coverUrl" :src="collection.coverUrl" :alt="collection.title" class="size-full object-cover" >
              <span v-else class="grid size-full place-items-center text-primary">
                <UIcon :name="collection?.icon || 'i-tabler-book-2'" class="size-5" />
              </span>
            </span>
            <span class="min-w-0">
              <span class="line-clamp-1 text-sm font-medium text-highlighted">{{ collection?.title || slug }}</span>
              <span class="block text-xs text-muted">{{ tree.length }} 个顶级章节</span>
            </span>
          </div>
          <UButton
            color="neutral"
            variant="soft"
            icon="i-tabler-list-tree"
            label="目录"
            :disabled="pending && !tree.length"
            @click="() => { navOpen = true }"
          />
        </div>
        <slot />
      </div>
    </div>
    <UDrawer v-model:open="navOpen" title="文档目录" handle>
      <template #body>
        <div class="space-y-4">
          <NuxtLink
            :to="collectionTo"
            class="flex items-center gap-2 rounded-lg border border-default px-3 py-2 text-sm font-semibold text-highlighted"
            @click="navOpen = false"
          >
            <UIcon name="i-tabler-book-2" class="size-4 text-primary" />
            <span class="min-w-0 line-clamp-1">{{ collection?.title || slug }}</span>
          </NuxtLink>
          <div v-if="pending && !tree.length" class="space-y-2">
            <USkeleton v-for="i in 6" :key="i" class="h-4 rounded" :class="i % 3 === 0 ? 'ml-4 w-32' : 'w-44'" />
          </div>
          <DocTree v-else :nodes="tree" :current-path="currentPath" @click="navOpen = false" />
        </div>
      </template>
    </UDrawer>
    <SiteFooter />
  </div>
</template>
