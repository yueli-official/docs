<script setup lang="ts">
withDefaults(defineProps<{ widthClass?: string }>(), { widthClass: 'max-w-screen-xl' })
const { brand: siteBrand } = useSiteRuntime()
</script>

<template>
  <footer class="border-t border-default">
    <div class="mx-auto w-full px-4 py-8 text-center text-xs text-muted" :class="widthClass">{{ siteBrand }}</div>
  </footer>
</template>
