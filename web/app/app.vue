<script setup lang="ts">
import { FeedbackToastRegion } from "@yueli/ui/feedback/pattern";

const { brand: siteBrand } = useSiteRuntime();
useHead({
  titleTemplate: (t) => (t ? `${t} · ${siteBrand.value}` : siteBrand.value),
});
</script>

<template>
  <UApp :toaster="null">
    <NuxtLoadingIndicator color="var(--ui-primary)" :height="3" />
    <NuxtLayout>
      <NuxtPage />
    </NuxtLayout>
    <FeedbackToastRegion />
  </UApp>
</template>
