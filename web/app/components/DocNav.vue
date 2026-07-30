<script setup lang="ts">
defineProps<{
  prev?: { path: string; title: string }
  next?: { path: string; title: string }
  collection: string
  query?: Record<string, string>
}>()
</script>

<template>
  <nav v-if="prev || next" class="mt-14 grid gap-4 border-t border-default pt-8 sm:grid-cols-2">
    <NuxtLink
      v-if="prev"
      :to="{ path: `/${collection}/${prev.path}`, query }"
      class="group flex flex-col rounded-2xl border border-default p-4 transition hover:border-primary/40 hover:shadow-sm"
    >
      <span class="flex items-center gap-1 text-xs text-muted"><UIcon name="i-tabler-arrow-left" class="size-3.5" />上一篇</span>
      <span class="mt-1 line-clamp-2 font-medium text-highlighted transition group-hover:text-primary">{{ prev.title }}</span>
    </NuxtLink>
    <span v-else class="hidden sm:block" />
    <NuxtLink
      v-if="next"
      :to="{ path: `/${collection}/${next.path}`, query }"
      class="group flex flex-col rounded-2xl border border-default p-4 text-right transition hover:border-primary/40 hover:shadow-sm sm:col-start-2"
    >
      <span class="flex items-center justify-end gap-1 text-xs text-muted">下一篇<UIcon name="i-tabler-arrow-right" class="size-3.5" /></span>
      <span class="mt-1 line-clamp-2 font-medium text-highlighted transition group-hover:text-primary">{{ next.title }}</span>
    </NuxtLink>
  </nav>
</template>
