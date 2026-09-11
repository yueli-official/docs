<script setup lang="ts">
import { collectionCoverThumbUrl } from "~/utils/coverImage";
import type { Collection, CollectionList, HomeConfigResponse } from "~/types";
import {
  collectionCountLabel,
  homeFeaturedCollections,
  homeQuickLinks,
} from "~/utils/docsManual.mjs";

definePageMeta({ width: "full" });
const { call } = useApi();
const { data } = await useAsyncData("collections", () =>
  call<CollectionList>("/api/v1/collections"),
);
const { data: homeData } = await useAsyncData("docs-home-config", () =>
  call<HomeConfigResponse>("/api/v1/home"),
);
if (!homeData.value?.config)
  throw createError({
    statusCode: 500,
    statusMessage: "文档站点配置尚未初始化",
  });
const collections = computed(() => data.value?.items ?? []);
const homeConfig = computed(() => homeData.value!.config);
const taskLinks = computed(() =>
  homeQuickLinks(homeConfig.value, collections.value),
);
const failedCovers = ref<Record<string, boolean>>({});
const featuredCollections = computed<Collection[]>(() =>
  homeFeaturedCollections(collections.value, homeConfig.value, 6),
);
const browseCollections = computed(() => {
  const featuredSlugs = new Set(
    featuredCollections.value.map((collection) => collection.slug),
  );
  const remaining = collections.value.filter(
    (collection) => !featuredSlugs.has(collection.slug),
  );
  return (remaining.length ? remaining : collections.value).slice(0, 6);
});
const startPath = computed(
  () =>
    taskLinks.value[0]?.to ||
    (featuredCollections.value[0]
      ? `/${featuredCollections.value[0].slug}`
      : collections.value[0]
        ? `/${collections.value[0].slug}`
        : "/collections"),
);
function coverVisible(c: { id: string; coverUrl?: string }) {
  return Boolean(c.coverUrl && !failedCovers.value[c.id]);
}
function markCoverFailed(id: string) {
  failedCovers.value = { ...failedCovers.value, [id]: true };
}
let coversMounted = false;
function sweepBrokenCovers() {
  if (!import.meta.client || !coversMounted) return;
  nextTick(() => {
    document
      .querySelectorAll<HTMLImageElement>("img[data-cover-id]")
      .forEach((img) => {
        if (img.complete && img.naturalWidth === 0)
          markCoverFailed(img.dataset.coverId || "");
      });
  });
}
watch(collections, sweepBrokenCovers, { immediate: true });
onMounted(() => {
  coversMounted = true;
  sweepBrokenCovers();
});
useSeoMeta({
  title: () => homeConfig.value.siteTitle,
  description: () => homeConfig.value.siteDescription,
});
</script>

<template>
  <div class="min-w-0 space-y-14 overflow-x-hidden">
    <header class="border-b border-default pb-10 pt-3 sm:pb-12 sm:pt-6">
      <div class="min-w-0 max-w-3xl">
        <h1
          class="font-display text-balance text-4xl font-bold leading-tight tracking-[-0.025em] text-highlighted sm:text-5xl"
        >
          {{ homeConfig.homeTitle }}
        </h1>
        <p
          v-if="homeConfig.homeSubtitle"
          class="mt-4 max-w-[65ch] break-words text-base leading-7 text-muted"
        >
          {{ homeConfig.homeSubtitle }}
        </p>

        <div class="mt-7">
          <UButton
            v-if="collections.length"
            :to="startPath"
            color="neutral"
            size="lg"
            trailing-icon="i-tabler-arrow-right"
            label="开始阅读"
          />
        </div>
      </div>
    </header>

    <section v-if="taskLinks.length" aria-labelledby="task-paths-title">
      <div class="mb-3 flex items-center justify-between gap-4">
        <h2
          id="task-paths-title"
          class="font-display text-xl font-semibold tracking-tight text-highlighted"
        >
          快速入口
        </h2>
      </div>
      <div class="grid gap-3 md:grid-cols-2">
        <NuxtLink
          v-for="task in taskLinks"
          :key="task.title"
          :to="task.to"
          class="group flex min-w-0 items-center gap-3 rounded-xl border border-default bg-default p-4 transition-colors hover:border-primary/40 hover:bg-elevated/50 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary"
        >
          <span
            class="grid size-9 shrink-0 place-items-center rounded-lg bg-primary/10 text-primary"
          >
            <UIcon :name="task.icon" class="size-5" />
          </span>
          <span class="min-w-0 flex-1">
            <span
              class="block font-medium text-highlighted transition-colors group-hover:text-primary"
              >{{ task.title }}</span
            >
            <span
              v-if="task.description"
              class="mt-0.5 line-clamp-1 text-sm text-muted"
              >{{ task.description }}</span
            >
          </span>
          <UIcon
            name="i-tabler-chevron-right"
            class="size-4 shrink-0 text-muted transition-transform group-hover:translate-x-0.5"
          />
        </NuxtLink>
      </div>
    </section>

    <section v-if="featuredCollections.length" aria-labelledby="featured-title">
      <h2
        id="featured-title"
        class="font-display mb-3 text-xl font-semibold tracking-tight text-highlighted"
      >
        推荐文档
      </h2>
      <div class="grid gap-3 md:grid-cols-2 xl:grid-cols-3">
        <NuxtLink
          v-for="collection in featuredCollections"
          :key="`featured-${collection.id}`"
          :to="`/${collection.slug}`"
          class="group flex min-w-0 items-center gap-3 rounded-xl border border-default bg-default p-4 transition-colors hover:border-primary/40 hover:bg-elevated/50 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary"
        >
          <span
            class="relative grid size-10 shrink-0 place-items-center overflow-hidden rounded-lg bg-elevated text-primary"
          >
            <img
              v-if="coverVisible(collection)"
              :src="collectionCoverThumbUrl(collection)"
              :alt="collection.title"
              :data-cover-id="collection.id"
              class="size-full object-cover"
              @error="markCoverFailed(collection.id)"
            />
            <UIcon v-else :name="collection.icon || 'i-tabler-book-2'" class="size-5" />
          </span>
          <span class="min-w-0 flex-1">
            <span
              class="line-clamp-1 font-display font-semibold text-highlighted transition-colors group-hover:text-primary"
              >{{ collection.title }}</span
            >
            <span
              v-if="collection.description"
              class="mt-1 line-clamp-1 text-sm text-muted"
              >{{ collection.description }}</span
            >
          </span>
          <UIcon
            name="i-tabler-arrow-right"
            class="size-4 shrink-0 text-muted transition-transform group-hover:translate-x-0.5"
          />
        </NuxtLink>
      </div>
    </section>

    <section v-if="collections.length" aria-labelledby="collections-title">
      <div class="mb-3 flex items-center justify-between gap-4">
        <h2
          id="collections-title"
          class="font-display text-xl font-semibold tracking-tight text-highlighted"
        >
          文档集
        </h2>
        <UButton
          to="/collections"
          color="neutral"
          variant="link"
          trailing-icon="i-tabler-arrow-right"
          label="查看全部"
          class="shrink-0"
        />
      </div>

      <div class="grid gap-3 md:grid-cols-2">
        <NuxtLink
          v-for="collection in browseCollections"
          :key="collection.id"
          :to="`/${collection.slug}`"
          class="group flex min-w-0 items-center gap-4 rounded-xl border border-default bg-default p-4 transition-colors hover:border-primary/40 hover:bg-elevated/50 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary"
        >
          <span
            class="relative grid size-12 shrink-0 place-items-center overflow-hidden rounded-lg bg-elevated text-primary"
          >
            <img
              v-if="coverVisible(collection)"
              :src="collectionCoverThumbUrl(collection)"
              :alt="collection.title"
              :data-cover-id="collection.id"
              class="size-full object-cover"
              @error="markCoverFailed(collection.id)"
            />
            <UIcon v-else :name="collection.icon || 'i-tabler-book-2'" class="size-5" />
          </span>
          <span class="min-w-0 flex-1">
            <span
              class="line-clamp-1 font-display text-base font-semibold text-highlighted transition-colors group-hover:text-primary"
              >{{ collection.title }}</span
            >
            <span
              v-if="collection.description"
              class="mt-1 line-clamp-1 text-sm text-muted"
              >{{ collection.description }}</span
            >
          </span>
          <span class="shrink-0 text-xs text-muted">{{
            collectionCountLabel(collection.docCount)
          }}</span>
          <UIcon
            name="i-tabler-chevron-right"
            class="size-4 shrink-0 text-muted transition-transform group-hover:translate-x-0.5 group-hover:text-primary"
          />
        </NuxtLink>
      </div>
    </section>

    <section v-else class="max-w-3xl border-y border-default py-6">
      <div class="flex min-w-0 items-center gap-3">
        <span
          class="grid size-9 shrink-0 place-items-center rounded-lg bg-elevated text-muted"
        >
          <UIcon name="i-tabler-file-off" class="size-5" />
        </span>
        <div class="min-w-0">
          <h2 class="font-display font-semibold text-highlighted">
            暂无公开文档
          </h2>
          <p class="mt-1 text-sm text-muted">发布文档后会显示在这里。</p>
        </div>
      </div>
    </section>
  </div>
</template>
