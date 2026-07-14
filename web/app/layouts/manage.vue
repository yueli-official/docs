<script setup lang="ts">
import { ManageShell } from "@platform/manage/components";

const route = useRoute();
const { brand: siteBrand } = useSiteRuntime();

const contextLabel = computed(() => {
  if (route.path === "/manage") return "控制台";
  if (route.path === "/manage/collections") return "文档集";
  if (route.path === "/manage/docs") return "文档";
  if (route.path.startsWith("/manage/docs/")) return "编辑文档";
  if (route.path === "/manage/import") return "批量导入";
  if (route.path.startsWith("/manage/import/")) return "导入详情";
  if (route.path === "/manage/home") return "站点设置";
  if (route.path === "/manage/assets") return "资源配置";
  return "控制台";
});
const showBackToTop = computed(() =>
  ["/manage", "/manage/home", "/manage/assets"].includes(route.path),
);
</script>

<template>
  <ManageShell
    :site-name="siteBrand"
    :context-label="contextLabel"
    storage-key="docs-manage"
    :show-back-to-top="showBackToTop"
  >
    <template #sidebar><ManageSidebar /></template>
    <template #user>
      <ConsumerManageAccountControl home-to="/" />
    </template>
    <slot />
  </ManageShell>
</template>
