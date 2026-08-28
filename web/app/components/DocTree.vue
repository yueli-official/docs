<script setup lang="ts">
import type { DocNode } from '~/types'
const props = withDefaults(defineProps<{ nodes: DocNode[]; basePath?: string; currentPath: string; defaultLocale?: string }>(), {
  defaultLocale: 'en',
})
function pathOf(n: DocNode) {
  return props.basePath ? `${props.basePath}/${n.slug}` : n.slug
}
const route = useRoute()
const collectionSlug = computed(() => route.params.collection as string)
const routeQuery = computed(() => ({
  ...(typeof route.query.locale === 'string' && route.query.locale !== props.defaultLocale ? { locale: route.query.locale } : {}),
  ...(typeof route.query.version === 'string' && route.query.version ? { version: route.query.version } : {}),
}))
</script>

<template>
  <ul class="space-y-0.5">
    <li v-for="n in nodes" :key="n.id">
      <NuxtLink
        :to="{ path: `/${collectionSlug}/${pathOf(n)}`, query: routeQuery }"
        class="block rounded-md px-2 py-1 text-sm leading-snug transition"
        :class="currentPath === pathOf(n)
          ? 'bg-primary/10 font-medium text-primary'
          : 'text-muted hover:bg-elevated hover:text-default'"
      >
        <span class="inline-flex min-w-0 items-center gap-1.5">
          <span class="truncate">{{ n.title }}</span>
          <UIcon v-if="n.badgeIcon" :name="n.badgeIcon" class="size-3.5 shrink-0" />
          <span v-else-if="n.badgeText" class="shrink-0 rounded bg-elevated px-1.5 py-0.5 text-[10px] leading-none text-toned">{{ n.badgeText }}</span>
        </span>
      </NuxtLink>
      <div v-if="n.children?.length" class="ml-3 mt-0.5 border-l border-default pl-2">
        <DocTree :nodes="n.children" :base-path="pathOf(n)" :current-path="currentPath" :default-locale="defaultLocale" />
      </div>
    </li>
  </ul>
</template>
