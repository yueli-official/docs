<script setup lang="ts">
import { BackToTop } from "@yueli/ui/navigation/back-to-top";

const route = useRoute();
const mainWidth = computed(
  () =>
    PAGE_WIDTHS[(route.meta.width as PageWidth) ?? "narrow"] ??
    PAGE_WIDTHS.narrow,
);
</script>

<template>
  <div class="flex min-h-dvh flex-col bg-default text-default">
    <SiteHeader :width-class="mainWidth" />
    <main
      id="public-main"
      tabindex="-1"
      class="mx-auto w-full flex-1 px-4 py-8 outline-none sm:py-10"
      :class="mainWidth"
    >
      <slot />
    </main>
    <SiteFooter :width-class="mainWidth" />
    <BackToTop target-id="public-main" label="返回顶部" />
  </div>
</template>
