<script setup lang="ts">
import { ManageAssetSettings } from "@platform/asset/components";

definePageMeta({ layout: "manage", middleware: "auth" });

const { slug, brand } = useSiteRuntime();
useSeoMeta({ title: "资源配置 · 控制台" });
</script>

<template>
  <YAdminPage
    id="assets"
    title="资源配置"
    icon="i-tabler-database-cog"
    main-id="manage-main"
    body-class="mx-auto w-full max-w-screen-2xl"
  >
    <ManageAssetSettings
      :site-key="slug"
      :site-name="brand"
      :show-header="false"
      description="配置文档站使用资源中心时的默认存储、文档集封面、正文图片与附件规则。"
      :profile-order="[
        'docs-collection-cover',
        'docs-import-image',
        'attachment',
      ]"
    />
  </YAdminPage>
</template>
