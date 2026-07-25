<script setup lang="ts">
import { ManageAssetSettings } from "@platform/asset/components";

definePageMeta({ layout: "manage", middleware: "auth" });

const { slug, brand } = useSiteRuntime();
const { can } = useMe();
const canManageAssets = computed(() => can("docs.asset_settings.manage"));
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
      :can-manage="canManageAssets"
      permission-description="当前角色不能修改本站资源策略。请联系管理员调整角色能力。"
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
