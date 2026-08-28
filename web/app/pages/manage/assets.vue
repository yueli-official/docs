<script setup lang="ts">
import { AssetPolicyPage } from "@yueli/asset-nuxt/components";

definePageMeta({ layout: "manage", middleware: "auth" });

const { can } = useMe();
const canEditAssets = computed(() => can("docs.asset_settings.manage"));
useSeoMeta({ title: "资源策略 · 控制台" });
</script>

<template>
  <AssetPolicyPage
    expected-namespace="docs"
    :profile-order="['docs-collection-cover']"
    :can-edit="canEditAssets"
  />
</template>
