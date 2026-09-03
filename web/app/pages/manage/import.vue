<script setup lang="ts">
import { createDocsNotifier } from "~/utils/feedback";
import { ManageEmpty } from "~/utils/manageComponents";
import { CollectionPagination } from "@yueli/ui/collection/pattern";
import type {
  DocsImportBatch,
  DocsImportListResponse,
  DocsImportSummary,
  DocsImportUploadResponse,
  CollectionView,
} from "~/types";
import {
  formatImportDate,
  importModeLabel,
  importStatusBadgeUI,
  importStatusMeta,
} from "~/utils/docsImportPresentation.mjs";

definePageMeta({ layout: "manage" });
useSeoMeta({ title: "批量导入 · 控制台" });

const { call } = useApi();
const { can } = useMe();
const canManageImports = computed(() => can("docs.import.manage"));
const canManageCollections = computed(() => can("docs.collection.manage"));
const toast = createDocsNotifier(useToast());

const file = ref<File | null>(null);
const uploading = ref(false);
const confirming = ref(false);
const errorMessage = ref("");
const batch = ref<DocsImportBatch | null>(null);
const summary = ref<DocsImportSummary | null>(null);
const collectionSlug = ref("");
const defaultLocale = ref("zh-CN");
const importMode = ref("upsert");
const historyPage = ref(1);
const historySize = 8;
const creatingCollection = ref(false);
const showCreateCollection = ref(false);
const createCollectionError = ref("");
const newCollection = reactive({
  title: "",
  slug: "",
  defaultLocale: "zh-CN",
  semanticVersion: "1.0.0",
});
const newCollectionSlugTouched = ref(false);
const mounted = ref(false);
onMounted(() => {
  mounted.value = true;
});

const canUpload = computed(() =>
  Boolean(
    canManageImports.value &&
    file.value &&
    collectionSlug.value &&
    !uploading.value &&
    !confirming.value,
  ),
);

const { data: collectionsData, refresh: refreshCollections } = await useAsyncData(
  "docs-import-collections",
  () => call<{ items: CollectionView[] }>("/api/v1/collections"),
  { server: false, default: () => ({ items: [] as CollectionView[] }) },
);
const collectionOptions = computed(() =>
  (collectionsData.value?.items ?? []).map((item) => ({ label: item.title, value: item.slug })),
);
const collectionTitleByID = computed(() => new Map(
  (collectionsData.value?.items ?? []).map((item) => [item.id, item.title] as const),
));
function importCollectionTitle(item: DocsImportBatch) {
  return collectionTitleByID.value.get(item.collectionId) || "未知文档集";
}
watch(collectionOptions, (items) => {
  if (!collectionSlug.value && items[0]) collectionSlug.value = items[0].value;
}, { immediate: true });
function normalizeSlug(value: string) {
  return value.trim().toLowerCase().replace(/[^\p{L}\p{N}]+/gu, "-").replace(/^-+|-+$/g, "");
}
watch(() => newCollection.title, (title) => {
  if (!newCollectionSlugTouched.value) newCollection.slug = normalizeSlug(title);
});
const canCreateCollection = computed(() => Boolean(
  canManageCollections.value &&
  newCollection.title.trim() &&
  normalizeSlug(newCollection.slug) &&
  /^\d+\.\d+\.\d+$/.test(newCollection.semanticVersion.trim()) &&
  !creatingCollection.value,
));

function openCreateCollection() {
  createCollectionError.value = "";
  newCollection.title = "";
  newCollection.slug = "";
  newCollection.defaultLocale = defaultLocale.value.trim() || "zh-CN";
  newCollection.semanticVersion = "1.0.0";
  newCollectionSlugTouched.value = false;
  showCreateCollection.value = true;
}

function closeCreateCollection() {
  showCreateCollection.value = false;
  createCollectionError.value = "";
}

async function createCollection() {
  if (!canCreateCollection.value) return;
  creatingCollection.value = true;
  createCollectionError.value = "";
  try {
    const res = await call<{ collection: CollectionView }>("/api/v1/collections", {
      method: "POST",
      body: {
        title: newCollection.title.trim(),
        slug: normalizeSlug(newCollection.slug),
        description: "",
        icon: "",
        cover: "",
        defaultLocale: newCollection.defaultLocale,
        semanticVersion: newCollection.semanticVersion.trim(),
      },
    });
    await refreshCollections();
    collectionSlug.value = res.collection.slug;
    defaultLocale.value = newCollection.defaultLocale;
    batch.value = null;
    summary.value = null;
    showCreateCollection.value = false;
    toast.add({
      title: "文档集已创建",
      description: `${res.collection.title} 已设为本次导入目标`,
      color: "success",
      icon: "i-tabler-circle-check",
    });
  } catch (err: any) {
    createCollectionError.value = err?.data?.message || err?.message || "创建文档集失败";
  } finally {
    creatingCollection.value = false;
  }
}
const modeOptions = [
  { label: "更新同路径文档", value: "upsert" },
  { label: "仅创建，不覆盖", value: "create-only" },
  { label: "同步并归档缺失文档", value: "replace-version" },
];
const canConfirm = computed(() =>
  Boolean(
    canManageImports.value &&
    batch.value?.id &&
    batch.value.status === "checked" &&
    !summary.value?.blocking &&
    !confirming.value,
  ),
);

const {
  data: historyData,
  pending: historyPending,
  error: historyError,
  refresh: refreshHistory,
} = await useAsyncData(
  () => `docs-import-history-${historyPage.value}`,
  () =>
    canManageImports.value
      ? call<DocsImportListResponse>("/api/v1/imports/docs", {
          query: { page: historyPage.value, size: historySize },
        })
      : Promise.resolve({ items: [], total: 0, page: 1, size: historySize }),
  {
    server: false,
    default: () => ({ items: [] as DocsImportBatch[], total: 0, page: 1, size: historySize }),
    watch: [historyPage],
  },
);
const recentImports = computed(() => historyData.value?.items ?? []);
const historyTotal = computed(() => historyData.value?.total ?? 0);
const historyTotalPages = computed(() => Math.max(1, Math.ceil(historyTotal.value / historySize)));
const showHistoryLoading = computed(
  () => !mounted.value || historyPending.value,
);

const summaryCards = computed(() => {
  const s = summary.value;
  return [
    {
      label: "创建",
      value: s?.creates ?? 0,
      icon: "i-tabler-file-plus",
      tone: "primary",
    },
    {
      label: "更新",
      value: s?.updates ?? 0,
      icon: "i-tabler-refresh",
      tone: "neutral",
    },
    {
      label: "归档",
      value: s?.archives ?? 0,
      icon: "i-tabler-archive",
      tone: "warning",
    },
    {
      label: "图片",
      value: s?.images ?? 0,
      icon: "i-tabler-photo",
      tone: "success",
    },
    {
      label: "问题",
      value: (s?.errors ?? 0) + (s?.conflicts ?? 0),
      icon: "i-tabler-alert-triangle",
      tone: s?.blocking ? "error" : "neutral",
    },
  ];
});

function toneClass(tone: string) {
  if (tone === "primary") return "bg-primary/10 text-primary ring-primary/20";
  if (tone === "success") return "bg-success/10 text-success ring-success/20";
  if (tone === "warning") return "bg-warning/10 text-warning ring-warning/20";
  if (tone === "error") return "bg-error/10 text-error ring-error/20";
  return "bg-elevated text-muted ring-default";
}

function onFileChange(event: Event) {
  const target = event.target as HTMLInputElement;
  file.value = target.files?.[0] ?? null;
  batch.value = null;
  summary.value = null;
  errorMessage.value = "";
}

async function upload() {
  if (!canManageImports.value || !file.value) return;
  if (!file.value.name.toLowerCase().endsWith(".zip")) {
    errorMessage.value = "请选择 ZIP 文件";
    return;
  }
  uploading.value = true;
  errorMessage.value = "";
  try {
    const body = new FormData();
    body.append("file", file.value);
    body.append("collection", collectionSlug.value);
    body.append("defaultLocale", defaultLocale.value.trim());
    body.append("mode", importMode.value);
    const res = await call<DocsImportUploadResponse>("/api/v1/imports/docs", {
      method: "POST",
      body,
    });
    batch.value = res.batch;
    summary.value = res.summary;
    historyPage.value = 1;
    await refreshHistory();
  } catch (err: any) {
    errorMessage.value = err?.data?.message || err?.message || "上传失败";
  } finally {
    uploading.value = false;
  }
}

async function confirmImport() {
  if (!canManageImports.value || !batch.value) return;
  confirming.value = true;
  batch.value = { ...batch.value, status: "running" };
  try {
    const res = await call<DocsImportUploadResponse>(
      `/api/v1/imports/docs/${batch.value.id}/confirm`,
      { method: "POST" },
    );
    await navigateTo(`/manage/import/${res.batch.id}`);
  } catch (err: any) {
    if (batch.value) batch.value = { ...batch.value, status: "failed" };
    toast.add({
      title: "导入失败",
      description:
        err?.data?.message || err?.message || "请检查预检结果后重试",
      color: "error",
      icon: "i-tabler-alert-circle",
      duration: 0,
    });
  } finally {
    confirming.value = false;
  }
}
</script>

<template>
  <ManagePage
    id="import"
    title="批量导入"
    icon="i-tabler-file-import"
    main-id="manage-main"
    body-class="w-full"
  >
    <div
      v-if="!canManageImports"
      class="rounded-lg border border-default bg-default p-8"
    >
      <div class="mx-auto max-w-md text-center">
        <span
          class="mx-auto grid size-12 place-items-center rounded-lg bg-warning/10 text-warning"
        >
          <UIcon name="i-tabler-lock" class="size-6" />
        </span>
        <h2 class="mt-4 text-lg font-semibold text-highlighted">
          没有批量导入权限
        </h2>
        <p class="mt-2 text-sm leading-6 text-muted">
          当前角色不能上传、确认或回滚导入批次。请联系管理员调整角色能力。
        </p>
      </div>
    </div>

    <div
      v-else
      class="grid items-start gap-4 xl:grid-cols-[minmax(0,1fr)_360px]"
    >
      <section class="min-w-0 space-y-4">
        <div class="rounded-lg border border-default bg-default p-5">
          <div class="mb-4 flex flex-wrap items-start justify-between gap-3">
            <div>
              <h2 class="text-base font-semibold text-highlighted">
                选择导入包
              </h2>
              <p class="mt-1 text-sm text-muted">
                上传 Markdown ZIP；多语言或显式导航包可附 docs.json。
                导入前会检查正文、图片与内部链接。
              </p>
            </div>
          </div>

          <div class="mb-4 flex gap-2 text-sm text-muted">
            <UIcon name="i-tabler-clock-play" class="mt-0.5 size-4 shrink-0 text-primary" />
            <p>预检会保存为导入批次；确认后由后台继续执行，可以离开页面并从最近导入查看进度。</p>
          </div>

          <div class="mb-3 grid gap-3 md:grid-cols-3">
            <UFormField label="目标文档集" required>
              <div class="flex gap-2">
                <USelect v-model="collectionSlug" :items="collectionOptions" value-key="value" class="min-w-0 flex-1" />
                <UButton
                  v-if="canManageCollections"
                  label="新建"
                  icon="i-tabler-plus"
                  color="neutral"
                  variant="outline"
                  @click="openCreateCollection"
                />
              </div>
            </UFormField>
            <UFormField label="默认语言" required>
              <UInput v-model="defaultLocale" placeholder="zh-CN" class="w-full" />
            </UFormField>
            <UFormField label="导入模式" required>
              <USelect v-model="importMode" :items="modeOptions" value-key="value" class="w-full" />
            </UFormField>
          </div>

          <div
            v-if="showCreateCollection"
            class="mb-4 border-t border-default pt-4"
            data-import-collection-creator
          >
            <div class="mb-3 flex items-start justify-between gap-3">
              <div>
                <h3 class="text-sm font-semibold text-highlighted">新建并选作文档集</h3>
                <p class="mt-1 text-xs text-muted">创建默认语言和首个版本后，继续使用当前 ZIP 预检。</p>
              </div>
              <UButton
                icon="i-tabler-x"
                aria-label="取消新建文档集"
                color="neutral"
                variant="ghost"
                size="xs"
                @click="closeCreateCollection"
              />
            </div>
            <div class="grid gap-3 md:grid-cols-2 xl:grid-cols-4">
              <UFormField label="标题" required>
                <UInput v-model="newCollection.title" autofocus placeholder="文档集标题" class="w-full" />
              </UFormField>
              <UFormField label="路径标识" required>
                <UInput
                  v-model="newCollection.slug"
                  placeholder="quickstart"
                  class="w-full"
                  @input="newCollectionSlugTouched = true"
                />
              </UFormField>
              <UFormField label="默认语言" required>
                <USelect
                  v-model="newCollection.defaultLocale"
                  :items="[
                    { label: '简体中文', value: 'zh-CN' },
                    { label: 'English', value: 'en-US' },
                    { label: '繁體中文', value: 'zh-TW' },
                    { label: '日本語', value: 'ja-JP' },
                  ]"
                  value-key="value"
                  class="w-full"
                />
              </UFormField>
              <UFormField label="首个版本" required>
                <UInput v-model="newCollection.semanticVersion" placeholder="1.0.0" class="w-full" />
              </UFormField>
            </div>
            <p v-if="createCollectionError" class="mt-3 text-sm text-error" role="alert">
              {{ createCollectionError }}
            </p>
            <div class="mt-3 flex justify-end">
              <UButton
                label="创建并选中"
                icon="i-tabler-folder-plus"
                :disabled="!canCreateCollection"
                :loading="creatingCollection"
                @click="createCollection"
              />
            </div>
          </div>

          <div class="grid gap-3 lg:grid-cols-[minmax(0,1fr)_auto]">
            <UInput
              type="file"
              accept=".zip,application/zip"
              icon="i-tabler-file-zip"
              class="min-w-0"
              @change="onFileChange"
            />
            <UButton
              icon="i-tabler-cloud-upload"
              label="上传并预检"
              :disabled="!canUpload"
              :loading="uploading"
              @click="upload"
            />
          </div>

          <UAlert
            v-if="errorMessage"
            class="mt-4"
            color="error"
            variant="soft"
            icon="i-tabler-alert-circle"
            title="上传失败"
            :description="errorMessage"
          />
        </div>

        <div
          v-if="summary"
          class="overflow-hidden rounded-lg border border-default bg-default"
        >
          <div class="border-b border-default bg-elevated/35 px-5 py-4">
            <div class="flex flex-wrap items-center justify-between gap-3">
              <div>
                <h2 class="text-base font-semibold text-highlighted">
                  预检摘要
                </h2>
                <p class="mt-1 text-sm text-muted">
                  {{
                    summary.blocking
                      ? "存在阻塞问题，修复 ZIP 后重新上传。"
                      : "没有阻塞问题，可以确认导入。"
                  }}
                </p>
              </div>
              <UBadge
                :label="summary.blocking ? '不可导入' : '可导入'"
                :color="summary.blocking ? 'warning' : 'success'"
                :icon="
                  summary.blocking
                    ? 'i-tabler-alert-triangle'
                    : 'i-tabler-shield-check'
                "
                variant="subtle"
              />
            </div>
          </div>

          <div class="grid sm:grid-cols-2 xl:grid-cols-5">
            <div
              v-for="card in summaryCards"
              :key="card.label"
              class="border-b border-default p-4 sm:border-r xl:border-b-0"
            >
              <div
                class="mb-3 grid size-10 place-items-center rounded-lg ring-1"
                :class="toneClass(card.tone)"
              >
                <UIcon :name="card.icon" class="size-5" />
              </div>
              <p class="text-xs font-medium text-muted">{{ card.label }}</p>
              <p class="mt-1 text-2xl font-semibold text-highlighted">
                {{ card.value }}
              </p>
            </div>
          </div>
        </div>

        <div
          v-if="summary?.issues?.length"
          class="rounded-lg border border-default bg-default"
        >
          <div class="border-b border-default bg-elevated/35 px-5 py-3">
            <h2 class="text-sm font-semibold text-highlighted">预检问题</h2>
          </div>
          <div class="divide-y divide-default">
            <div
              v-for="(issue, index) in summary.issues"
              :key="`${issue.code}-${issue.path}-${index}`"
              class="grid gap-2 px-5 py-3 lg:grid-cols-[140px_minmax(0,1fr)_220px]"
            >
              <UBadge
                :label="issue.severity === 'error' ? '错误' : '警告'"
                :color="issue.severity === 'error' ? 'error' : 'warning'"
                variant="subtle"
                size="sm"
              />
              <div class="min-w-0">
                <p class="text-sm font-medium text-highlighted">
                  {{ issue.message }}
                </p>
                <p class="mt-1 font-mono text-xs text-muted">
                  {{ issue.code }}
                </p>
              </div>
              <p class="truncate font-mono text-xs text-muted">
                {{ issue.path || "-" }}
              </p>
            </div>
          </div>
        </div>

        <div
          v-if="!summary && !errorMessage"
          class="flex items-center gap-3 rounded-lg border border-dashed border-default px-4 py-3 text-sm text-muted"
        >
          <span class="grid size-9 shrink-0 place-items-center rounded-lg bg-primary/10 text-primary">
            <UIcon name="i-tabler-file-search" class="size-4.5" />
          </span>
          <p>选择 ZIP 并预检后，这里会显示变更数量和需要处理的问题。</p>
        </div>

        <section
          aria-labelledby="import-history-title"
          class="overflow-hidden rounded-lg border border-default bg-default"
        >
          <div
            class="flex items-center justify-between gap-3 border-b border-default bg-elevated/35 px-5 py-3"
          >
            <div>
              <h2
                id="import-history-title"
                class="text-sm font-semibold text-highlighted"
              >
                最近导入
              </h2>
              <p class="mt-0.5 text-xs text-muted">
                批次由后台持续执行，离开此页不会隐藏记录。
              </p>
            </div>
            <UButton
              v-if="historyError"
              label="重试"
              icon="i-tabler-refresh"
              color="neutral"
              variant="ghost"
              size="xs"
              @click="() => refreshHistory()"
            />
          </div>
          <div v-if="showHistoryLoading" class="grid gap-2 p-4">
            <USkeleton v-for="item in 4" :key="item" class="h-14 rounded-lg" />
          </div>
          <div
            v-else-if="historyError"
            class="px-5 py-8 text-center text-sm text-error"
          >
            导入历史暂时无法加载
          </div>
          <div v-else-if="recentImports.length" class="divide-y divide-default">
            <NuxtLink
              v-for="item in recentImports"
              :key="item.id"
              :to="`/manage/import/${item.id}`"
              class="group grid gap-2 px-5 py-3 transition hover:bg-elevated/40 sm:grid-cols-[minmax(0,1fr)_auto_auto] sm:items-center"
            >
              <span class="min-w-0">
                <span class="block truncate text-sm font-medium text-highlighted">
                  {{ importCollectionTitle(item) }}
                </span>
                <span class="mt-1 block text-xs text-muted">
                  {{ importModeLabel(item.mode) }} ·
                  {{ item.summary.creates + item.summary.updates }} 篇变更 ·
                  <span class="font-mono">{{ item.id }}</span>
                </span>
              </span>
              <UBadge
                :label="importStatusMeta(item.status).label"
                :color="importStatusMeta(item.status).color"
                :icon="importStatusMeta(item.status).icon"
                :ui="importStatusBadgeUI(item.status)"
                :data-import-status="item.status"
                variant="subtle"
                size="sm"
              />
              <time
                class="text-xs tabular-nums text-muted"
                :datetime="item.createdAt"
              >
                {{ formatImportDate(item.createdAt) }}
              </time>
            </NuxtLink>
          </div>
          <ManageEmpty v-else icon="i-tabler-history" text="还没有导入记录" />
          <div
            v-if="historyTotalPages > 1"
            class="flex items-center justify-between gap-3 border-t border-default px-5 py-3"
          >
            <p class="text-xs text-muted">共 {{ historyTotal }} 个批次</p>
            <CollectionPagination v-model="historyPage" :total-pages="historyTotalPages" />
          </div>
        </section>
      </section>

      <aside
        class="rounded-lg border border-default bg-default p-4 xl:self-start"
      >
        <div class="mb-4 flex items-start justify-between gap-3">
          <div>
            <p class="text-xs font-medium uppercase tracking-wide text-muted">
              导入状态
            </p>
            <h2 class="mt-1 text-base font-semibold text-highlighted">
              {{ confirming ? "后台处理中" : batch ? importStatusMeta(batch.status).label : "等待上传" }}
            </h2>
          </div>
          <span
            class="grid size-10 place-items-center rounded-lg bg-primary/10 text-primary"
          >
            <UIcon name="i-tabler-file-import" class="size-5" />
          </span>
        </div>

        <div
          class="space-y-3 rounded-lg border border-default bg-elevated/35 p-3 text-sm"
        >
          <div class="flex items-center justify-between gap-3">
            <span class="text-muted">批次</span>
            <span class="max-w-44 truncate font-mono text-xs text-default">{{
              batch?.id || "-"
            }}</span>
          </div>
          <div class="flex items-center justify-between gap-3">
            <span class="text-muted">模式</span>
            <span class="font-mono text-xs text-default">{{
              batch ? importModeLabel(batch.mode) : "-"
            }}</span>
          </div>
          <div class="flex items-center justify-between gap-3">
            <span class="text-muted">默认语言</span>
            <span class="font-mono text-xs text-default">{{
              batch?.defaultLocale || "-"
            }}</span>
          </div>
        </div>

        <USeparator class="my-4" />

        <div
          v-if="confirming"
          class="mb-4 flex gap-2 rounded-lg bg-primary/5 p-3 text-sm text-default"
        >
          <UIcon name="i-tabler-loader-2" class="mt-0.5 size-4 shrink-0 animate-spin text-primary" />
          <p>任务正在后台执行。可以离开本页，稍后从“最近导入”返回查看结果。</p>
        </div>

        <div class="grid gap-2">
          <UButton
            label="确认导入"
            icon="i-tabler-check"
            block
            :disabled="!canConfirm"
            :loading="confirming"
            @click="confirmImport"
          />
          <UButton
            v-if="batch?.id"
            label="查看详情"
            icon="i-tabler-list-details"
            color="neutral"
            variant="outline"
            block
            :to="`/manage/import/${batch.id}`"
          />
        </div>
      </aside>
    </div>
  </ManagePage>
</template>
