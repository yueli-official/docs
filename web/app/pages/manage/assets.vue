<script setup lang="ts">
import { AssetRegistrationSummary } from "@yueli/asset-nuxt/components";

definePageMeta({ layout: "manage", middleware: "auth" });

const { can } = useMe();
const canEditAssets = computed(() => can("docs.asset_settings.manage"));
useSeoMeta({ title: "资源策略 · 控制台" });
</script>

<template>
  <ManagePage
    id="assets"
    title="资源策略"
    icon="i-tabler-database-cog"
    main-id="manage-main"
    body-class="w-full"
  >
    <ClientOnly>
      <AssetRegistrationSummary expected-namespace="docs" :profile-order="['docs-collection-cover']" :can-edit="canEditAssets" />
      <template #fallback><USkeleton class="h-56 w-full rounded-xl" /></template>
    </ClientOnly>
  </ManagePage>
</template>
