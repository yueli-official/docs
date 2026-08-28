<script setup lang="ts">
import type {
  CollectionReleasesResponse,
  CollectionVariantsResponse,
  DocumentVariantResolution,
} from "~/types";
import { createDocsNotifier } from "~/utils/feedback";

const props = withDefaults(defineProps<{
  collectionSlug: string;
  translationKey?: string;
}>(), {
  translationKey: "",
});
const route = useRoute();
const { call } = useApi();
const toast = createDocsNotifier(useToast());
const { data } = await useAsyncData(
  () => `document-variant-switcher-${props.collectionSlug}`,
  async () => {
    const [variants, releases] = await Promise.all([
      call<CollectionVariantsResponse>(`/api/v1/collections/${props.collectionSlug}/variants`),
      call<CollectionReleasesResponse>(`/api/v1/collections/${props.collectionSlug}/releases`),
    ]);
    return { variants, releases };
  },
  { watch: [() => props.collectionSlug] },
);
const locales = computed(() => data.value?.variants?.locales ?? []);
const releases = computed(() => data.value?.releases?.items ?? []);
const defaultLocale = computed(() => locales.value.find((item) => item.isDefault) ?? locales.value[0]);
const selectedLocale = computed(() =>
  typeof route.query.locale === "string"
    ? route.query.locale
    : defaultLocale.value?.locale || "en",
);
const localeItems = computed(() => locales.value.map((item) => ({
  label: item.label,
  value: item.locale,
})),
);
const releaseItems = computed(() => releases.value.map((item) => ({
  label: item.semanticVersion || item.title,
  value: item.slug,
})),
);

async function switchVariant(next: { locale?: string; release?: string }) {
  const requestedLocale = next.locale || selectedLocale.value;
  const targetCollection = next.release || props.collectionSlug;
  const targetVariants = targetCollection === props.collectionSlug
    ? data.value?.variants
    : await call<CollectionVariantsResponse>(`/api/v1/collections/${targetCollection}/variants`);
  const targetDefaultLocale = targetVariants?.locales.find((item) => item.isDefault)?.locale
    || targetVariants?.locales[0]?.locale
    || "en";
  const resolution = await call<DocumentVariantResolution>(
    `/api/v1/collections/${targetCollection}/variant`,
    {
      query: {
        translationKey: props.translationKey,
        locale: requestedLocale,
      },
    },
  );
  if (resolution.fallback === "default_locale") {
    const requested = locales.value.find((item) => item.locale === requestedLocale)?.label || requestedLocale;
    toast.add({
      title: "正在显示默认语言",
      description: `此页暂无${requested}版本。`,
      color: "info",
      duration: 3_000,
    });
  } else if (resolution.fallback === "collection" && props.translationKey) {
    toast.add({
      title: "已打开文档集首页",
      description: "目标版本没有对应页面。",
      color: "info",
      duration: 3_000,
    });
  }
  const query: Record<string, string> = {};
  if (resolution.locale && resolution.locale !== targetDefaultLocale) query.locale = resolution.locale;
  await navigateTo({
    path: resolution.path
      ? `/${targetCollection}/${resolution.path}`
      : `/${targetCollection}`,
    query,
  });
}
</script>

<template>
  <div
    v-if="locales.length > 1 || releases.length > 1"
    class="flex flex-wrap items-center gap-2"
    data-document-variant-switcher
  >
    <USelect
      v-if="locales.length > 1"
      :model-value="selectedLocale"
      :items="localeItems"
      value-key="value"
      icon="i-tabler-language"
      size="sm"
      aria-label="切换文档语言"
      class="min-w-36"
      @update:model-value="switchVariant({ locale: String($event) })"
    />
    <USelect
      v-if="releases.length > 1"
      :model-value="collectionSlug"
      :items="releaseItems"
      value-key="value"
      icon="i-tabler-versions"
      size="sm"
      aria-label="切换文档版本"
      class="min-w-36"
      @update:model-value="switchVariant({ release: String($event) })"
    />
  </div>
</template>
