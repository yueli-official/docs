<script setup lang="ts">
const { brand: siteBrand } = useSiteRuntime()
useHead({ titleTemplate: t => (t ? `${t} · ${siteBrand.value}` : siteBrand.value) })
</script>

<template>
  <UApp>
    <NuxtLoadingIndicator color="var(--ui-primary)" :height="3" />
    <NuxtLayout>
      <NuxtPage />
    </NuxtLayout>
  </UApp>
</template>
