<script setup lang="ts">
import { ReadingTableOfContents } from "@yueli/ui/navigation/table-of-contents";
import type { CollectionVariantsResponse, DocDetailResponse } from "~/types";
import { createTrafficReplayKey } from "~/utils/traffic-replay-key.mjs";
import { trafficSource } from "~/utils/traffic-source.mjs";

definePageMeta({ layout: "collection", middleware: "url-lifecycle" });
const route = useRoute();
const { call } = useApi();
const collectionSlug = computed(() => route.params.collection as string);
const { data: variantData } = await useAsyncData(
  () => `collection-variants-${collectionSlug.value}`,
  () => call<CollectionVariantsResponse>(`/api/v1/collections/${collectionSlug.value}/variants`),
  { watch: [collectionSlug] },
);
const defaultLocale = computed(() =>
  variantData.value?.locales.find((item) => item.isDefault)?.locale || "en",
);
const locale = computed(() =>
  typeof route.query.locale === "string" ? route.query.locale : defaultLocale.value,
);
const currentLocale = computed(() =>
  variantData.value?.locales.find((item) => item.locale === locale.value),
);
const version = computed(() =>
  typeof route.query.version === "string" ? route.query.version : "",
);
const routeQuery = computed(() => ({
  ...(locale.value !== defaultLocale.value ? { locale: locale.value } : {}),
  ...(version.value ? { version: version.value } : {}),
}));
function docTo(path = "") {
  return {
    path: path
      ? `/${collectionSlug.value}/${path}`
      : `/${collectionSlug.value}`,
    query: routeQuery.value,
  };
}
const docPath = computed(() => {
  const rest = route.params.slug;
  return Array.isArray(rest) ? rest.join("/") : String(rest ?? "");
});

const { ensure, byPath, flat, collection } = useCollectionTree(
  () => route.params.collection as string,
  locale,
  version,
);
await ensure();

const node = computed(() => byPath.value.get(docPath.value) ?? null);
if (!node.value)
  throw createError({
    statusCode: 404,
    statusMessage: "文档不存在",
    fatal: true,
  });

const { data: docData, pending: docPending } = await useAsyncData(
  () =>
    `doc-${collectionSlug.value}-${locale.value}-${version.value || "default"}-${docPath.value}`,
  () =>
    call<DocDetailResponse>("/api/v1/docs/by-path", {
      query: {
        collection: collectionSlug.value,
        path: docPath.value,
        locale: locale.value,
        ...(version.value ? { version: version.value } : {}),
      },
    }),
  { watch: [collectionSlug, docPath, locale, version] },
);
const doc = computed(() => docData.value?.doc ?? null);
const recordedDocument = ref("");
watch(
  () => doc.value?.id,
  (documentID) => {
    if (
      !import.meta.client ||
      !documentID ||
      documentID === recordedDocument.value
    )
      return;
    recordedDocument.value = documentID;
    const viewEvent = {
      eventId: createTrafficReplayKey(),
      occurredAt: new Date().toISOString(),
      source: trafficSource(document.referrer, window.location.href),
    };
    const recordView = () =>
      call(`/api/v1/docs/${documentID}/view`, {
        method: "POST",
        body: viewEvent,
      });
    recordView().catch(() => recordView().catch(() => {}));
  },
  { immediate: true },
);

// prev/next = adjacent entries in DFS pre-order (whole flattened list).
const idx = computed(() =>
  flat.value.findIndex((e) => e.path === docPath.value),
);
const prev = computed(() => (idx.value > 0 ? flat.value[idx.value - 1] : null));
const next = computed(() =>
  idx.value >= 0 && idx.value < flat.value.length - 1
    ? flat.value[idx.value + 1]
    : null,
);
function childPath(slug: string) {
  return [docPath.value, slug].filter(Boolean).join("/");
}

const { renderWithToc } = useMarkdown();
function stripDuplicateTitle(content: string, title: string) {
  const trimmedTitle = title.trim();
  if (!trimmedTitle) return content;
  const escaped = trimmedTitle.replace(/[.*+?^${}()|[\]\\]/g, "\\$&");
  return content.replace(new RegExp(`^#\\s+${escaped}\\s*\\n+`), "");
}

const readingContent = computed(() =>
  stripDuplicateTitle(doc.value?.content ?? "", node.value?.title ?? ""),
);
const toc = computed(() => renderWithToc(readingContent.value).toc);
const isLeaf = computed(() => !node.value?.children?.length);
useDiscoveryPage(() => docData.value?.discovery);
useHead(() => ({
  htmlAttrs: {
    lang: currentLocale.value?.htmlLang || locale.value,
    dir: currentLocale.value?.direction || "ltr",
  },
}));
</script>

<template>
  <div v-if="node" class="reading-shell" data-reading-shell>
    <div
      class="grid gap-12"
      :class="toc.length ? 'xl:grid-cols-[minmax(0,1fr)_240px]' : ''"
    >
      <article class="min-w-0">
        <header class="border-b border-default pb-8">
          <div
            class="flex flex-col gap-3 sm:flex-row sm:items-start sm:justify-between"
          >
            <div
              class="flex flex-wrap items-center gap-2 text-xs font-medium text-muted"
            >
              <span
                class="inline-flex items-center gap-1 rounded-md bg-primary/10 px-2 py-1 text-primary"
              >
                <UIcon
                  :name="collection?.icon || 'i-tabler-book-2'"
                  class="size-3.5"
                />
                {{ collection?.title }}
              </span>
              <span
                class="inline-flex items-center gap-1 rounded-md bg-elevated px-2 py-1"
              >
                <UIcon
                  :name="isLeaf ? 'i-tabler-file-text' : 'i-tabler-folder'"
                  class="size-3.5"
                />
                {{ isLeaf ? "文档" : "章节" }}
              </span>
            </div>
            <DocumentVariantSwitcher
              class="sm:justify-end"
              :collection-slug="collectionSlug"
              :translation-key="doc?.translationKey || node.translationKey"
            />
          </div>
          <h1
            class="font-display mt-4 flex flex-wrap items-center gap-3 text-balance text-[2rem] font-semibold leading-[1.18] text-highlighted md:text-[2.3125rem]"
          >
            <span>{{ node.title }}</span>
            <UIcon v-if="node.badgeIcon" :name="node.badgeIcon" class="size-7 shrink-0 text-primary" />
            <UBadge v-else-if="node.badgeText" :label="node.badgeText" color="neutral" variant="soft" size="md" />
          </h1>
          <p
            v-if="node.excerpt"
            class="mt-5 max-w-3xl text-lg leading-8 text-muted"
          >
            {{ node.excerpt }}
          </p>
        </header>

        <div class="mt-8">
          <ContentProse v-if="readingContent" :content="readingContent" />
          <div v-else-if="docPending" class="space-y-3">
            <USkeleton class="h-6 w-2/3" />
            <USkeleton class="h-4 w-full" />
            <USkeleton class="h-4 w-11/12" />
            <USkeleton class="h-4 w-4/5" />
            <USkeleton class="h-4 w-10/12" />
          </div>
        </div>

        <section v-if="!isLeaf && node.children" class="mt-10">
          <div class="mb-4 flex items-center justify-between gap-3">
            <h2 class="text-lg font-semibold text-highlighted">章节内容</h2>
            <span class="text-sm text-muted"
              >{{ node.children.length }} 篇</span
            >
          </div>
          <div class="grid gap-3 sm:grid-cols-2">
            <NuxtLink
              v-for="c in node.children"
              :key="c.id"
              :to="docTo(childPath(c.slug))"
              class="group rounded-lg border border-default bg-default p-4 transition hover:border-primary/40 hover:bg-elevated/50"
            >
              <span
                class="flex items-center gap-2 text-sm font-semibold text-highlighted transition group-hover:text-primary"
              >
                <UIcon name="i-tabler-file-text" class="size-4" />
                {{ c.title }}
                <UIcon v-if="c.badgeIcon" :name="c.badgeIcon" class="size-4 text-primary" />
                <UBadge v-else-if="c.badgeText" :label="c.badgeText" color="neutral" variant="soft" size="xs" />
              </span>
              <span
                v-if="c.excerpt"
                class="mt-2 line-clamp-2 text-sm leading-6 text-muted"
                >{{ c.excerpt }}</span
              >
            </NuxtLink>
          </div>
        </section>

        <DocNav
          :prev="prev ? { path: prev.path, title: prev.node.title } : undefined"
          :next="next ? { path: next.path, title: next.node.title } : undefined"
          :collection="collectionSlug"
          :query="routeQuery"
        />

        <DocumentComments v-if="doc" :document-id="doc.id" />
      </article>

      <aside v-if="toc.length" class="hidden xl:block">
        <div class="sticky top-24">
          <ReadingTableOfContents
            :items="toc"
            title="本页目录"
            :min-level="2"
            :max-level="4"
          />
        </div>
      </aside>
    </div>
  </div>
</template>
