<script setup lang="ts">
import {
  ManageEmpty,
  ManageIconPicker,
  ManageRepeaterRow,
  SkeletonList,
} from "@platform/manage/components";
import {
  platformSettingsSaveMessages,
  usePlatformSettingsProtection,
} from "@platform/manage/settings";
import { createPlatformNotifier } from "@platform/ui/feedback";
import { useActionFeedback, useMinimumLoading } from "@yueli/ui/feedback";
import {
  SettingSection,
  SettingsLayout,
  SettingsSaveDock,
} from "@yueli/ui/settings/pattern";
import { useVueSettingsWorkflow } from "@yueli/ui/settings/vue";
import type {
  CollectionList,
  HomeConfigResponse,
  HomeQuickLink,
} from "~/types";

definePageMeta({ layout: "manage" });
useSeoMeta({ title: "设置 · 控制台" });

const { call } = useApi();
const toast = createPlatformNotifier(useToast());
const route = useRoute();
const router = useRouter();
const saveError = ref("");
const section = ref<"home" | "footer" | "site">("home");
const sections = [
  {
    key: "home",
    label: "首页",
    icon: "i-tabler-home-cog",
    description: "首屏文案、快速入口和推荐文档",
  },
  {
    key: "footer",
    label: "页脚",
    icon: "i-tabler-layout-bottombar",
    description: "页脚标语、版权与联系入口",
  },
  {
    key: "site",
    label: "基础",
    icon: "i-tabler-adjustments-horizontal",
    description: "站点名称、描述与支持邮箱",
  },
] as const;
const sectionKeys = sections.map((item) => item.key);
const activeSection = computed(
  () => sections.find((item) => item.key === section.value) || sections[0],
);
const mounted = ref(false);
onMounted(() => {
  mounted.value = true;
});

const { data: collectionsData, pending: collectionsPending } =
  await useAsyncData(
    "manage-home-collections",
    () => call<CollectionList>("/api/v1/collections"),
    { server: false, default: () => ({ items: [] }) },
  );
const {
  data: homeData,
  pending: homePending,
  error: homeError,
  refresh,
} = await useAsyncData(
  "manage-home-config",
  () => call<HomeConfigResponse>("/api/v1/home"),
  { server: false },
);

const collections = computed(() => collectionsData.value?.items ?? []);
const collectionItems = computed(() =>
  collections.value.map((item) => ({ label: item.title, value: item.slug })),
);
const loading = computed(
  () => !mounted.value || collectionsPending.value || homePending.value,
);
const showSkeleton = useMinimumLoading(loading);
const quickLinks = ref<HomeQuickLink[]>([]);
const featuredCollections = ref<string[]>([]);
const homeCopy = reactive({
  eyebrow: "",
  title: "",
  subtitle: "",
});
const siteForm = reactive({ title: "", description: "", supportEmail: "" });
const footerForm = reactive({ tagline: "", copyright: "" });
const {
  status: saveStatus,
  pending: markSaving,
  success: markSaved,
  reset: resetSave,
} = useActionFeedback();
const initialized = ref(false);
const activeIconPickerLinkId = ref("");
const settingsState = useVueSettingsWorkflow({
  snapshot: () => ({
    quickLinks: quickLinks.value,
    featuredCollections: featuredCollections.value,
    homeCopy,
    site: siteForm,
    footer: footerForm,
  }),
  restore: (snapshot) => {
    quickLinks.value = snapshot.quickLinks;
    featuredCollections.value = snapshot.featuredCollections;
    Object.assign(homeCopy, snapshot.homeCopy);
    Object.assign(siteForm, snapshot.site);
    Object.assign(footerForm, snapshot.footer);
  },
});
usePlatformSettingsProtection(() => settingsState.dirty.value);

watch(
  homeData,
  (value) => {
    if (initialized.value) return;
    const config = value?.config;
    if (!config) return;
    quickLinks.value = config.quickLinks.map((link, index) => ({
      id: link.id || crypto.randomUUID(),
      title: link.title || "",
      description: link.description || "",
      icon: link.icon || "i-tabler-arrow-up-right",
      to: link.to || "",
      collectionSlug: link.collectionSlug || "",
      sortOrder: index,
      enabled: link.enabled !== false,
    }));
    featuredCollections.value = [...config.featuredCollections];
    Object.assign(homeCopy, {
      eyebrow: config.homeEyebrow,
      title: config.homeTitle,
      subtitle: config.homeSubtitle,
    });
    Object.assign(siteForm, {
      title: config.siteTitle,
      description: config.siteDescription,
      supportEmail: config.supportEmail,
    });
    Object.assign(footerForm, {
      tagline: config.footerTagline,
      copyright: config.footerCopyright,
    });
    initialized.value = true;
    nextTick(settingsState.capture);
  },
  { immediate: true },
);

watch(
  () => route.query.section,
  (value) => {
    section.value =
      typeof value === "string" &&
      sectionKeys.includes(value as typeof section.value)
        ? (value as typeof section.value)
        : "home";
  },
  { immediate: true },
);

watch(section, (value) => {
  if (route.query.section === value) return;
  router.replace({ query: { ...route.query, section: value } });
});

function addQuickLink() {
  quickLinks.value.push({
    id: crypto.randomUUID(),
    title: "",
    description: "",
    icon: "i-tabler-arrow-up-right",
    to: "",
    collectionSlug: collections.value[0]?.slug || "",
    sortOrder: quickLinks.value.length,
    enabled: true,
  });
}

function removeQuickLink(id: string) {
  quickLinks.value = quickLinks.value.filter((link) => link.id !== id);
  if (activeIconPickerLinkId.value === id) activeIconPickerLinkId.value = "";
}

function moveQuickLink(index: number, dir: -1 | 1) {
  const next = index + dir;
  if (next < 0 || next >= quickLinks.value.length) return;
  const copy = [...quickLinks.value];
  const current = copy[index];
  copy[index] = copy[next]!;
  copy[next] = current!;
  quickLinks.value = copy;
}

function isIconPickerOpen(id: string) {
  return activeIconPickerLinkId.value === id;
}

function setIconPickerOpen(id: string, open: boolean) {
  activeIconPickerLinkId.value = open ? id : "";
}

function chooseQuickLinkIcon(link: HomeQuickLink, value: string) {
  link.icon = value;
  activeIconPickerLinkId.value = "";
}

function quickLinkActionItems(link: HomeQuickLink, index: number) {
  return [
    [
      {
        label: "上移",
        icon: "i-tabler-arrow-up",
        disabled: index === 0,
        onSelect: () => moveQuickLink(index, -1),
      },
      {
        label: "下移",
        icon: "i-tabler-arrow-down",
        disabled: index === quickLinks.value.length - 1,
        onSelect: () => moveQuickLink(index, 1),
      },
    ],
    [
      {
        label: "删除",
        icon: "i-tabler-trash",
        color: "error" as const,
        onSelect: () => removeQuickLink(link.id),
      },
    ],
  ];
}

function toggleFeatured(slug: string) {
  featuredCollections.value = featuredCollections.value.includes(slug)
    ? featuredCollections.value.filter((item) => item !== slug)
    : [...featuredCollections.value, slug];
}

function collectionTitle(slug: string) {
  return collections.value.find((item) => item.slug === slug)?.title || slug;
}

function quickLinkTarget(link: HomeQuickLink) {
  if (link.to) return link.to;
  if (link.collectionSlug) return `/${link.collectionSlug}`;
  return "/";
}

async function save() {
  markSaving();
  saveError.value = "";
  try {
    const body = {
      quickLinks: quickLinks.value.map((link, index) => ({
        ...link,
        sortOrder: index,
      })),
      featuredCollections: featuredCollections.value,
      homeEyebrow: homeCopy.eyebrow,
      homeTitle: homeCopy.title,
      homeSubtitle: homeCopy.subtitle,
      siteTitle: siteForm.title,
      siteDescription: siteForm.description,
      supportEmail: siteForm.supportEmail,
      footerTagline: footerForm.tagline,
      footerCopyright: footerForm.copyright,
    };
    await call<HomeConfigResponse>("/api/v1/home", { method: "PATCH", body });
    initialized.value = false;
    await refresh();
    settingsState.capture();
    markSaved();
  } catch (e: any) {
    resetSave();
    saveError.value = e?.data?.message || "请稍后重试";
    toast.add({
      title: "设置保存失败",
      description: saveError.value,
      color: "error",
    });
  }
}

function discardChanges() {
  settingsState.discard();
  saveError.value = "";
  resetSave();
}
</script>

<template>
  <YAdminPage
    id="site-settings"
    title="站点设置"
    icon="i-tabler-settings"
    main-id="manage-main"
    body-class="mx-auto w-full max-w-screen-2xl"
  >
    <SettingsLayout
      v-model:active-section="section"
      :title="activeSection.label"
      :description="activeSection.description"
      :sections="sections"
      :show-section-navigation="false"
      :show-header="false"
      navigation-label="设置分区"
    >
      <SkeletonList v-if="showSkeleton" :rows="8" />

      <UAlert
        v-else-if="homeError"
        color="error"
        variant="subtle"
        icon="i-tabler-alert-circle"
        title="设置加载失败"
        description="站点配置尚未初始化或服务不可用，请先运行开发环境 provision。"
      />

      <template v-else-if="section === 'home' && !homeError">
        <SettingSection
          title="首页首屏"
          description="控制公开首页的眉标、标题和任务导向说明。"
          class="mb-5"
        >
          <div class="grid gap-4">
            <div class="grid gap-3 sm:grid-cols-[180px_minmax(0,1fr)]">
              <UFormField label="眉标"
                ><UInput v-model="homeCopy.eyebrow" class="w-full"
              /></UFormField>
              <UFormField label="首页标题"
                ><UInput v-model="homeCopy.title" class="w-full"
              /></UFormField>
            </div>
            <UFormField label="首页介绍"
              ><UTextarea v-model="homeCopy.subtitle" :rows="3" class="w-full"
            /></UFormField>
          </div>
        </SettingSection>

        <div class="grid gap-5 xl:grid-cols-[minmax(0,1fr)_380px]">
          <section class="min-w-0 rounded-lg border border-default bg-default">
            <div
              class="flex flex-wrap items-center justify-between gap-3 border-b border-default bg-elevated/35 px-4 py-3"
            >
              <div>
                <h2 class="text-sm font-semibold text-highlighted">
                  快速入口配置
                </h2>
                <p class="mt-0.5 text-xs text-muted">
                  配置首页入口卡片的标题、路径和展示状态。
                </p>
              </div>
              <UButton
                icon="i-tabler-plus"
                label="添加入口"
                color="neutral"
                variant="soft"
                @click="addQuickLink"
              />
            </div>

            <div class="p-4">
              <ManageEmpty
                v-if="!quickLinks.length"
                icon="i-tabler-route"
                text="还没有快速入口"
              />

              <div v-else class="grid gap-3">
                <ManageRepeaterRow
                  v-for="(link, index) in quickLinks"
                  :key="link.id"
                  :label="`快速入口 ${index + 1}`"
                >
                  <div class="flex flex-wrap items-start justify-between gap-3">
                    <div
                      class="grid min-w-0 grid-cols-[2.75rem_minmax(0,1fr)] items-center gap-3"
                    >
                      <UPopover
                        :open="isIconPickerOpen(link.id)"
                        :content="{ align: 'start', side: 'bottom' }"
                        :ui="{ content: 'w-auto p-3' }"
                        @update:open="setIconPickerOpen(link.id, $event)"
                      >
                        <UButton
                          :icon="link.icon || 'i-tabler-arrow-up-right'"
                          color="primary"
                          variant="soft"
                          square
                          class="size-11 shrink-0 justify-center"
                          aria-label="选择入口图标"
                        />
                        <template #content>
                          <ManageIconPicker
                            :model-value="link.icon"
                            compact
                            @update:model-value="
                              chooseQuickLinkIcon(link, $event)
                            "
                          />
                        </template>
                      </UPopover>
                      <div class="min-w-0">
                        <p
                          class="truncate text-sm font-semibold text-highlighted"
                        >
                          {{ link.title || `入口 ${index + 1}` }}
                        </p>
                        <p class="mt-0.5 truncate text-xs text-muted">
                          {{ quickLinkTarget(link) }}
                        </p>
                      </div>
                    </div>
                    <UDropdownMenu
                      :items="quickLinkActionItems(link, index)"
                      :ui="{ content: 'w-36' }"
                    >
                      <UButton
                        icon="i-tabler-dots-vertical"
                        color="neutral"
                        variant="ghost"
                        square
                        aria-label="入口操作"
                      />
                    </UDropdownMenu>
                  </div>

                  <div class="grid gap-3">
                    <div class="grid gap-3 md:grid-cols-2">
                      <UFormField label="标题">
                        <UInput
                          v-model="link.title"
                          placeholder="快速开始"
                          class="w-full"
                        />
                      </UFormField>
                      <UFormField label="关联文档集">
                        <USelectMenu
                          v-model="link.collectionSlug"
                          :items="collectionItems"
                          value-key="value"
                          placeholder="选择文档集"
                          :search-input="{ placeholder: '搜索文档集' }"
                          class="w-full"
                        />
                      </UFormField>
                    </div>

                    <UFormField label="说明">
                      <UTextarea
                        v-model="link.description"
                        :rows="2"
                        class="w-full"
                      />
                    </UFormField>

                    <div
                      class="grid gap-3 md:grid-cols-[minmax(0,1fr)_auto] md:items-end"
                    >
                      <UFormField label="自定义链接">
                        <UInput
                          v-model="link.to"
                          placeholder="留空时使用文档集路径"
                          class="w-full"
                        />
                      </UFormField>
                      <UCheckbox
                        v-model="link.enabled"
                        label="公开展示"
                        class="pb-2"
                      />
                    </div>
                  </div>
                </ManageRepeaterRow>
              </div>
            </div>
          </section>

          <aside class="space-y-5 xl:sticky xl:top-20 xl:self-start">
            <section class="rounded-lg border border-default bg-default">
              <div class="border-b border-default bg-elevated/35 px-4 py-3">
                <h2 class="text-sm font-semibold text-highlighted">推荐文档</h2>
                <p class="mt-0.5 text-xs text-muted">
                  选择首页展示的文档集，右侧顺序按选择顺序生成。
                </p>
              </div>
              <div class="grid gap-2 p-4">
                <button
                  v-for="collection in collections"
                  :key="collection.id"
                  type="button"
                  class="flex items-center justify-between gap-3 rounded-lg border px-3 py-2 text-left transition"
                  :class="
                    featuredCollections.includes(collection.slug)
                      ? 'border-primary/40 bg-primary/10 text-primary'
                      : 'border-default bg-elevated/35 text-default hover:border-primary/30'
                  "
                  @click="toggleFeatured(collection.slug)"
                >
                  <span class="flex min-w-0 items-center gap-3">
                    <span
                      class="grid size-8 shrink-0 place-items-center rounded-md bg-default text-primary ring-1 ring-default"
                    >
                      <UIcon
                        :name="collection.icon || 'i-tabler-stack-2'"
                        class="size-4"
                      />
                    </span>
                    <span class="min-w-0">
                      <span class="block truncate text-sm font-medium">{{
                        collection.title
                      }}</span>
                      <span class="block truncate font-mono text-xs opacity-70"
                        >/{{ collection.slug }}</span
                      >
                    </span>
                  </span>
                  <UIcon
                    :name="
                      featuredCollections.includes(collection.slug)
                        ? 'i-tabler-circle-check'
                        : 'i-tabler-circle'
                    "
                    class="size-5 shrink-0"
                  />
                </button>
              </div>
            </section>

            <section class="rounded-lg border border-default bg-default p-4">
              <div class="flex items-center justify-between gap-3">
                <h2 class="text-sm font-semibold text-highlighted">展示顺序</h2>
                <UIcon name="i-tabler-list-numbers" class="size-5 text-muted" />
              </div>
              <div v-if="featuredCollections.length" class="mt-3 space-y-2">
                <div
                  v-for="(slug, index) in featuredCollections"
                  :key="slug"
                  class="flex items-center justify-between gap-3 rounded-md bg-elevated px-3 py-2 text-sm"
                >
                  <span class="truncate"
                    >{{ index + 1 }}. {{ collectionTitle(slug) }}</span
                  >
                  <span class="font-mono text-xs text-muted">/{{ slug }}</span>
                </div>
              </div>
              <p v-else class="mt-3 text-sm text-muted">还没有选择推荐文档。</p>
            </section>
          </aside>
        </div>
      </template>

      <SettingSection
        v-else-if="section === 'footer' && !homeError"
        title="页脚内容"
        description="用于所有文档页面底部的品牌说明与联系信息。"
      >
        <div class="grid gap-4">
          <UFormField label="页脚标语"
            ><UInput v-model="footerForm.tagline" class="w-full"
          /></UFormField>
          <UFormField label="版权信息"
            ><UInput
              v-model="footerForm.copyright"
              placeholder="© 2026 Yueli"
              class="w-full"
          /></UFormField>
        </div>
      </SettingSection>

      <SettingSection
        v-else-if="!homeError"
        title="站点基础"
        description="这些字段用于导航品牌、SEO 描述和支持入口。"
      >
        <div class="grid gap-4 sm:grid-cols-2">
          <UFormField label="站点名称" required
            ><UInput v-model="siteForm.title" class="w-full"
          /></UFormField>
          <UFormField label="支持邮箱"
            ><UInput
              v-model="siteForm.supportEmail"
              type="email"
              class="w-full"
          /></UFormField>
          <UFormField label="站点描述" class="sm:col-span-2"
            ><UTextarea v-model="siteForm.description" :rows="3" class="w-full"
          /></UFormField>
        </div>
      </SettingSection>
      <SettingsSaveDock
        :dirty="settingsState.dirty.value"
        :status="saveStatus"
        :error="saveError"
        :messages="platformSettingsSaveMessages"
        dock-class="lg:left-60"
        @discard="discardChanges"
        @save="save"
      />
    </SettingsLayout>
  </YAdminPage>
</template>
