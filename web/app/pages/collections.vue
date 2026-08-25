<script setup lang="ts">
import { SkeletonCards } from "~/utils/manageComponents";
import { useMinimumLoading } from "@yueli/ui/feedback";
import type { CollectionList } from "~/types";
import {
  collectionCountLabel,
  filterCollections,
  paginateItems,
} from "~/utils/docsManual.mjs";

definePageMeta({ width: "full" });

const { call } = useApi();
const { data, pending, error, refresh } = await useAsyncData(
  "collections-directory",
  () => call<CollectionList>("/api/v1/collections"),
  { default: () => ({ items: [] }) },
);

const collections = computed(() => data.value?.items ?? []);
const query = ref("");
const page = ref(1);
const pageSize = ref(18);
const pageSizeItems = [
  { label: "每页 12 个", value: 12 },
  { label: "每页 18 个", value: 18 },
  { label: "每页 24 个", value: 24 },
];
const failedCovers = ref<Record<string, boolean>>({});
const showSkeleton = useMinimumLoading(pending);

const filteredCollections = computed(() =>
  filterCollections(collections.value, query.value),
);
const pageResult = computed(() =>
  paginateItems(filteredCollections.value, page.value, pageSize.value),
);
const pagedCollections = computed(() => pageResult.value.items);
const hasQuery = computed(() => query.value.trim().length > 0);
const hasLoadError = computed(() =>
  Boolean(error.value && !collections.value.length),
);

function coverVisible(c: { id: string; coverUrl?: string }) {
  return Boolean(c.coverUrl && !failedCovers.value[c.id]);
}

function markCoverFailed(id: string) {
  failedCovers.value = { ...failedCovers.value, [id]: true };
}

function clearQuery() {
  query.value = "";
}

watch([query, pageSize], () => {
  page.value = 1;
});
watch(
  () => pageResult.value.page,
  (safePage) => {
    page.value = safePage;
  },
);

useSeoMeta({ title: "文档集", description: "浏览和搜索公开文档集" });
</script>

<template>
  <div class="min-w-0 space-y-7 overflow-x-hidden">
    <header class="border-b border-default pb-6 pt-1 sm:pt-3">
      <h1
        class="font-display text-2xl font-bold tracking-[-0.02em] text-highlighted sm:text-3xl"
      >
        文档集
      </h1>
    </header>

    <section class="space-y-6" aria-labelledby="collections-directory-title">
      <div
        v-if="hasLoadError"
        class="rounded-lg border border-error/30 bg-error/5 p-5"
      >
        <div class="flex min-w-0 items-start gap-3">
          <UIcon
            name="i-tabler-alert-triangle"
            class="mt-0.5 size-5 shrink-0 text-error"
          />
          <div class="min-w-0">
            <h2 class="text-sm font-semibold text-highlighted">
              文档集加载失败
            </h2>
            <p class="mt-1 text-sm text-muted">
              请稍后重试，或检查服务是否正常运行。
            </p>
            <UButton
              class="mt-4"
              color="neutral"
              variant="soft"
              icon="i-tabler-refresh"
              label="重新加载"
              @click="refresh()"
            />
          </div>
        </div>
      </div>

      <template v-else>
        <div class="flex min-w-0 flex-col gap-3 sm:flex-row sm:items-center">
          <UInput
            v-model="query"
            icon="i-tabler-search"
            size="lg"
            placeholder="搜索文档集"
            aria-label="搜索文档集"
            class="w-full sm:max-w-xl"
          >
            <template v-if="hasQuery" #trailing>
              <UButton
                color="neutral"
                variant="ghost"
                icon="i-tabler-x"
                square
                size="xs"
                aria-label="清除搜索"
                @click="clearQuery"
              />
            </template>
          </UInput>
          <USelect
            v-if="filteredCollections.length > 12"
            v-model="pageSize"
            :items="pageSizeItems"
            value-key="value"
            aria-label="每页数量"
            class="w-full sm:w-40"
          />
        </div>

        <div class="flex min-h-6 items-center justify-between gap-4">
          <h2 id="collections-directory-title" class="sr-only">文档集列表</h2>
          <p class="text-sm text-muted">
            {{
              pageResult.total ? `${pageResult.total} 个文档集` : "没有匹配项"
            }}
          </p>
        </div>

        <SkeletonCards v-if="showSkeleton" :count="6" />

        <div
          v-else-if="pagedCollections.length"
          class="grid gap-3 md:grid-cols-2"
        >
          <NuxtLink
            v-for="c in pagedCollections"
            :key="c.id"
            :to="`/${c.slug}`"
            class="group flex min-w-0 items-center gap-4 rounded-xl border border-default bg-default p-4 transition-colors hover:border-primary/40 hover:bg-elevated/50 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary"
          >
            <span
              class="relative grid size-12 shrink-0 place-items-center overflow-hidden rounded-lg bg-elevated text-primary"
            >
              <img
                v-if="coverVisible(c)"
                :src="collectionCoverThumbUrl(c)"
                :alt="c.title"
                class="size-full object-cover"
                @error="markCoverFailed(c.id)"
              />
              <UIcon
                v-else
                :name="c.icon || 'i-tabler-book-2'"
                class="size-5"
              />
            </span>
            <span class="min-w-0 flex-1">
              <span
                class="line-clamp-1 font-display text-base font-semibold text-highlighted transition-colors group-hover:text-primary"
              >
                {{ c.title }}
              </span>
              <span
                class="mt-1 flex min-w-0 items-center gap-2 text-xs text-muted"
              >
                <span class="truncate font-mono">/{{ c.slug }}</span>
                <span aria-hidden="true">·</span>
                <span class="shrink-0">{{
                  collectionCountLabel(c.docCount)
                }}</span>
              </span>
              <span
                v-if="c.description"
                class="mt-1.5 line-clamp-1 text-sm text-muted"
              >
                {{ c.description }}
              </span>
            </span>
            <UIcon
              name="i-tabler-chevron-right"
              class="size-4 shrink-0 text-muted transition-transform group-hover:translate-x-0.5 group-hover:text-primary"
            />
          </NuxtLink>
        </div>

        <div
          v-else
          class="rounded-2xl border border-dashed border-default p-10 text-center"
        >
          <UIcon
            name="i-tabler-search-off"
            class="mx-auto size-10 text-muted"
          />
          <h3 class="mt-4 font-display text-lg font-semibold text-highlighted">
            没有匹配的文档集
          </h3>
          <p class="mt-2 text-sm text-muted">
            换一个关键词，或清除筛选后浏览全部集合。
          </p>
          <UButton
            class="mt-5"
            color="neutral"
            variant="soft"
            icon="i-tabler-x"
            label="清除筛选"
            @click="clearQuery"
          />
        </div>

        <div
          v-if="pageResult.total > pageResult.pageSize"
          class="flex justify-center border-t border-default pt-5"
        >
          <UPagination
            v-model:page="page"
            :total="pageResult.total"
            :items-per-page="pageResult.pageSize"
          />
        </div>
      </template>
    </section>
  </div>
</template>
