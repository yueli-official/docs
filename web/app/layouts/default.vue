<script setup lang="ts">
const route = useRoute()
const mainWidth = computed(() => PAGE_WIDTHS[(route.meta.width as PageWidth) ?? 'narrow'] ?? PAGE_WIDTHS.narrow)
</script>

<template>
  <div class="flex min-h-dvh flex-col bg-default text-default">
    <SiteHeader :width-class="mainWidth" />
    <main class="mx-auto w-full flex-1 px-4 py-8 sm:py-10" :class="mainWidth">
      <slot />
    </main>
    <SiteFooter :width-class="mainWidth" />
  </div>
</template>
