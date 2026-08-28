<script setup lang="ts">
import {
  ManageEmpty,
  ManageRepeaterRow,
  SkeletonList,
} from "~/utils/manageComponents";
import {
  docsSettingsSaveMessages,
  useDocsSettingsProtection,
} from "~/utils/settings";
import { createDocsNotifier } from "~/utils/feedback";
import { normalizeFeaturedCollections } from "~/utils/docsHomeConfig.mjs";
import { useActionFeedback, useMinimumLoading } from "@yueli/ui/feedback";
import { ActionFeedbackButton } from "@yueli/ui/feedback/pattern";
import { AdminIconPicker, AdminRowActions, TabbedSurface } from "@yueli/ui/admin";
import type { AdminRowActionItem } from "@yueli/ui/admin";
import { SettingSection } from "@yueli/ui/settings/pattern";
import { useVueSettingsWorkflow } from "@yueli/ui/settings/vue";
import type {
  CollectionList,
  HomeConfigResponse,
  HomeQuickLink,
} from "~/types";

definePageMeta({ layout: "manage" });
useSeoMeta({ title: "设置 · 控制台" });

const { call } = useApi();
const { can } = useMe();
const canManageSiteSettings = computed(() => can("docs.site_settings.manage"));
const toast = createDocsNotifier(useToast());
const route = useRoute();
const router = useRouter();
const saveError = ref("");
const section = ref<"home" | "footer" | "site">("home");
const sectionTabs = [
  {
    value: "home",
    label: "首页",
    icon: "i-tabler-home-cog",
  },
  {
    value: "footer",
    label: "页脚",
    icon: "i-tabler-layout-bottombar",
  },
  {
    value: "site",
    label: "基础",
    icon: "i-tabler-adjustments-horizontal",
  },
] as const;
const sectionKeys = sectionTabs.map((item) => item.value);
const mounted = ref(false);
onMounted(() => {
  mounted.value = true;
});

const { data: collectionsData, pending: collectionsPending } =
  await useAsyncData(
    "manage-home-collections",
    () =>
      canManageSiteSettings.value
        ? call<CollectionList>("/api/v1/collections")
        : Promise.resolve({ items: [] }),
    { server: false, default: () => ({ items: [] }) },
  );
const {
  data: homeData,
  pending: homePending,
  error: homeError,
  refresh,
} = await useAsyncData(
  "manage-home-config",
  () =>
    canManageSiteSettings.value
      ? call<HomeConfigResponse>("/api/v1/home")
      : Promise.resolve(null as unknown as HomeConfigResponse),
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
const featuredPickerOpen = ref(false);
const featuredSearch = ref("");
let quickLinkDraftSequence = 0;
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
useDocsSettingsProtection(() => settingsState.dirty.value);

watch(
  homeData,
  (value) => {
    if (initialized.value) return;
    const config = value?.config;
    if (!config) return;
    quickLinks.value = config.quickLinks.map((link, index) => ({
      id: link.id || newQuickLinkDraftKey(),
      title: link.title || "",
      description: link.description || "",
      icon: link.icon || "i-tabler-arrow-up-right",
      to: link.to || "",
      collectionSlug: link.collectionSlug || "",
      sortOrder: index,
      enabled: link.enabled !== false,
    }));
    featuredCollections.value = normalizeFeaturedCollections(
      config.featuredCollections,
    );
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
    id: newQuickLinkDraftKey(),
    title: "",
    description: "",
    icon: "i-tabler-arrow-up-right",
    to: "",
    collectionSlug: collections.value[0]?.slug || "",
    sortOrder: quickLinks.value.length,
    enabled: true,
  });
}

function newQuickLinkDraftKey() {
  quickLinkDraftSequence += 1;
  return `draft-quick-link-${quickLinkDraftSequence}`;
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

const collectionBySlug = computed(
  () => new Map(collections.value.map((item) => [item.slug, item] as const)),
);
const selectedFeaturedCollections = computed(() =>
  featuredCollections.value.map((slug) => {
    const collection = collectionBySlug.value.get(slug);
    return {
      slug,
      title: collection?.title || slug,
      icon: collection?.icon || "i-tabler-stack-2",
    };
  }),
);
const availableFeaturedCollections = computed(() => {
  const selected = new Set(featuredCollections.value);
  const keyword = featuredSearch.value.trim().toLocaleLowerCase();
  return collections.value.filter((collection) => {
    if (selected.has(collection.slug)) return false;
    if (!keyword) return true;
    return [collection.title, collection.slug, collection.description]
      .join(" ")
      .toLocaleLowerCase()
      .includes(keyword);
  });
});

function openFeaturedPicker() {
  featuredSearch.value = "";
  featuredPickerOpen.value = true;
}

function closeFeaturedPicker() {
  featuredPickerOpen.value = false;
}

function addFeatured(slug: string) {
  if (!featuredCollections.value.includes(slug)) {
    featuredCollections.value = [...featuredCollections.value, slug];
  }
}

function removeFeatured(slug: string) {
  featuredCollections.value = featuredCollections.value.filter(
    (item) => item !== slug,
  );
}

function moveFeatured(index: number, direction: -1 | 1) {
  const target = index + direction;
  if (target < 0 || target >= featuredCollections.value.length) return;
  const next = [...featuredCollections.value];
  [next[index], next[target]] = [next[target]!, next[index]!];
  featuredCollections.value = next;
}

function featuredCollectionActionItems(
  slug: string,
  index: number,
): AdminRowActionItem[] {
  return [
    {
      id: "move-up",
      label: "上移",
      icon: "i-tabler-arrow-up",
      disabled: index === 0,
      onSelect: () => moveFeatured(index, -1),
    },
    {
      id: "move-down",
      label: "下移",
      icon: "i-tabler-arrow-down",
      disabled: index === featuredCollections.value.length - 1,
      onSelect: () => moveFeatured(index, 1),
    },
    {
      id: "remove",
      label: "移除",
      icon: "i-tabler-trash",
      tone: "danger",
      onSelect: () => removeFeatured(slug),
    },
  ];
}

function quickLinkTarget(link: HomeQuickLink) {
  if (link.to) return link.to;
  if (link.collectionSlug) return `/${link.collectionSlug}`;
  return "/";
}

async function save() {
  if (!canManageSiteSettings.value) return;
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
  <ManagePage
    id="site-settings"
    title="站点设置"
    icon="i-tabler-settings"
    main-id="manage-main"
    body-class="w-full"
  >
    <template #actions>
      <div
        v-if="canManageSiteSettings"
        class="flex items-center gap-2"
        data-settings-header-actions
      >
        <UButton
          v-if="settingsState.dirty.value"
          :label="docsSettingsSaveMessages.discard"
          color="neutral"
          variant="ghost"
          :disabled="saveStatus === 'pending'"
          @click="discardChanges"
        />
        <ActionFeedbackButton
          :status="saveStatus"
          :idle-label="docsSettingsSaveMessages.save"
          :pending-label="docsSettingsSaveMessages.savePending"
          :success-label="docsSettingsSaveMessages.saveSuccess"
          :disabled="
            !settingsState.dirty.value ||
            saveStatus === 'pending' ||
            showSkeleton ||
            Boolean(homeError)
          "
          @click="save"
        />
      </div>
    </template>

    <div
      v-if="!canManageSiteSettings"
      class="rounded-lg border border-default bg-default p-8"
    >
      <div class="mx-auto max-w-md text-center">
        <span
          class="mx-auto grid size-12 place-items-center rounded-lg bg-warning/10 text-warning"
        >
          <UIcon name="i-tabler-lock" class="size-6" />
        </span>
        <h2 class="mt-4 text-lg font-semibold text-highlighted">
          没有站点设置权限
        </h2>
        <p class="mt-2 text-sm leading-6 text-muted">
          当前角色不能修改首页、页脚或站点基础资料。请联系管理员调整角色能力。
        </p>
      </div>
    </div>

    <TabbedSurface
      v-else
      v-model="section"
      :items="sectionTabs"
      navigation-label="站点设置"
      data-manage-surface="site-settings"
    >
      <div v-if="showSkeleton" class="p-4 sm:p-5">
        <SkeletonList :rows="8" />
      </div>

      <UAlert
        v-else-if="homeError"
        class="m-4 sm:m-5"
        color="error"
        variant="subtle"
        icon="i-tabler-alert-circle"
        title="设置加载失败"
        description="站点配置尚未初始化或服务不可用，请先运行开发环境 provision。"
      />

      <div
        v-else-if="section === 'home' && !homeError"
        class="min-w-0 p-4 sm:p-5"
      >
        <SettingSection
          title="首页文案"
          class="mb-5"
        >
          <div class="grid gap-4 md:grid-cols-2">
            <UFormField label="标题">
              <UInput v-model="homeCopy.title" class="w-full" />
            </UFormField>
            <UFormField label="简介">
              <UInput v-model="homeCopy.subtitle" class="w-full" />
            </UFormField>
          </div>
        </SettingSection>

        <div class="grid gap-5 xl:grid-cols-[minmax(0,1fr)_380px]">
          <SettingSection title="快速入口配置">
            <template #actions>
              <UButton
                icon="i-tabler-plus"
                label="添加入口"
                color="neutral"
                variant="soft"
                @click="addQuickLink"
              />
            </template>

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
                          <AdminIconPicker
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
          </SettingSection>

          <aside class="space-y-5 xl:sticky xl:top-20 xl:self-start">
            <SettingSection title="推荐文档">
              <template #actions>
                <UButton
                  label="添加文档"
                  icon="i-tabler-plus"
                  color="neutral"
                  variant="soft"
                  size="sm"
                  @click="openFeaturedPicker"
                />
              </template>

              <ManageEmpty
                v-if="!selectedFeaturedCollections.length"
                icon="i-tabler-star"
                text="还没有推荐文档"
              />
              <div v-else class="space-y-2">
                <div
                  v-for="(collection, index) in selectedFeaturedCollections"
                  :key="collection.slug"
                  :data-featured-collection="collection.slug"
                  class="flex min-w-0 items-center gap-3 rounded-lg border border-default bg-default px-3 py-2.5"
                >
                  <span
                    class="grid size-9 shrink-0 place-items-center rounded-md bg-primary/10 text-primary"
                  >
                    <UIcon :name="collection.icon" class="size-4" />
                  </span>
                  <span class="min-w-0 flex-1">
                    <span class="block truncate text-sm font-medium text-highlighted">
                      {{ collection.title }}
                    </span>
                    <span class="block truncate text-xs text-muted">
                      /{{ collection.slug }}
                    </span>
                  </span>
                  <AdminRowActions
                    :label="`${collection.title} 的排序与移除操作`"
                    :items="featuredCollectionActionItems(collection.slug, index)"
                  />
                </div>
              </div>
            </SettingSection>
          </aside>
        </div>
      </div>

      <div v-else-if="section === 'footer' && !homeError" class="p-4 sm:p-5">
        <SettingSection title="页脚内容">
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
      </div>

      <div v-else-if="!homeError" class="p-4 sm:p-5">
        <SettingSection title="站点基础">
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
      </div>
    </TabbedSurface>

    <UModal v-model:open="featuredPickerOpen" title="添加推荐文档">
      <template #body>
        <div class="space-y-4">
          <UInput
            v-model="featuredSearch"
            icon="i-tabler-search"
            placeholder="搜索文档集"
            autofocus
            class="w-full"
          />

          <div
            v-if="availableFeaturedCollections.length"
            class="max-h-80 divide-y divide-default overflow-y-auto rounded-lg border border-default"
          >
            <div
              v-for="collection in availableFeaturedCollections"
              :key="collection.id"
              :data-featured-candidate="collection.slug"
              class="flex min-w-0 items-center gap-3 px-3 py-2.5"
            >
              <span
                class="grid size-9 shrink-0 place-items-center rounded-md bg-primary/10 text-primary"
              >
                <UIcon
                  :name="collection.icon || 'i-tabler-stack-2'"
                  class="size-4"
                />
              </span>
              <span class="min-w-0 flex-1">
                <span class="block truncate text-sm font-medium text-highlighted">
                  {{ collection.title }}
                </span>
                <span class="block truncate text-xs text-muted">
                  /{{ collection.slug }}
                </span>
              </span>
              <UButton
                label="添加"
                color="neutral"
                variant="soft"
                size="xs"
                @click="addFeatured(collection.slug)"
              />
            </div>
          </div>
          <ManageEmpty
            v-else
            icon="i-tabler-search-off"
            :text="featuredSearch ? '没有匹配的文档集' : '所有文档集都已添加'"
          />
        </div>
      </template>
      <template #footer>
        <div class="flex w-full justify-end">
          <UButton
            label="完成"
            color="neutral"
            variant="outline"
            @click="closeFeaturedPicker"
          />
        </div>
      </template>
    </UModal>
  </ManagePage>
</template>
