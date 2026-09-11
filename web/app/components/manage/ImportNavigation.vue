<script setup lang="ts">
import { TabbedSurface } from '@yueli/ui/admin';
const route = useRoute();
const current = computed(() => route.path.startsWith('/manage/import-sources') ? 'sync' : 'import');
const tabs = [
  { label: '导入', value: 'import', icon: 'i-tabler-file-import' },
  { label: '项目同步', value: 'sync', icon: 'i-tabler-brand-github' },
];
function select(value: string | number) {
  return navigateTo(value === 'sync' ? '/manage/import-sources' : '/manage/import');
}
</script>

<template>
  <TabbedSurface :model-value="current" :items="tabs" navigation-label="文档导入与同步" @update:model-value="select">
    <div class="min-w-0 p-4 sm:p-5"><slot /></div>
  </TabbedSurface>
</template>
