<script setup lang="ts">
import { createDocsNotifier } from "~/utils/feedback";
import { ManageEmpty } from "~/utils/manageComponents";
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
  importStatusMeta,
} from "~/utils/docsImportPresentation.mjs";

definePageMeta({ layout: "manage" });
useSeoMeta({ title: "批量导入 · 控制台" });

const { call } = useApi();
const { can } = useMe();
const canManageImports = computed(() => can("docs.import.manage"));
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

const { data: collectionsData } = await useAsyncData(
  "docs-import-collections",
  () => call<{ items: CollectionView[] }>("/api/v1/collections"),
  { server: false, default: () => ({ items: [] as CollectionView[] }) },
);
const collectionOptions = computed(() =>
  (collectionsData.value?.items ?? []).map((item) => ({ label: item.title, value: item.slug })),
);
watch(collectionOptions, (items) => {
  if (!collectionSlug.value && items[0]) collectionSlug.value = items[0].value;
}, { immediate: true });
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
  "docs-import-history",
  () =>
    canManageImports.value
      ? call<DocsImportListResponse>("/api/v1/imports/docs", {
          query: { limit: 12 },
        })
      : Promise.resolve({ items: [] }),
  {
    server: false,
    default: () => ({ items: [] as DocsImportBatch[] }),
  },
);
const recentImports = computed(() => historyData.value?.items ?? []);
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
    await refreshHistory();
  } catch (err: any) {
    errorMessage.value = err?.data?.message || err?.message || "上传失败";
    toast.add({
      title: "上传失败",
      description: errorMessage.value,
      color: "error",
      icon: "i-tabler-alert-circle",
    });
  } finally {
    uploading.value = false;
  }
}

async function confirmImport() {
  if (!canManageImports.value || !batch.value) return;
  confirming.value = true;
  try {
    const res = await call<DocsImportUploadResponse>(
      `/api/v1/imports/docs/${batch.value.id}/confirm`,
      { method: "POST" },
    );
    await navigateTo(`/manage/import/${res.batch.id}`);
  } catch (err: any) {
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
            <UBadge
              label="ZIP 格式"
              icon="i-tabler-file-zip"
              variant="subtle"
            />
          </div>

          <div class="mb-3 grid gap-3 md:grid-cols-3">
            <UFormField label="目标文档集" required>
              <USelect v-model="collectionSlug" :items="collectionOptions" value-key="value" class="w-full" />
            </UFormField>
            <UFormField label="默认语言" required>
              <UInput v-model="defaultLocale" placeholder="zh-CN" class="w-full" />
            </UFormField>
            <UFormField label="导入模式" required>
              <USelect v-model="importMode" :items="modeOptions" value-key="value" class="w-full" />
            </UFormField>
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

        <ManageEmpty
          v-if="!summary && !errorMessage"
          icon="i-tabler-file-import"
          text="选择 ZIP 后开始预检"
        />

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
                查看预检、执行、失败与回滚记录。
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
                <span class="block truncate font-mono text-xs text-highlighted">
                  {{ item.id }}
                </span>
                <span class="mt-1 block text-xs text-muted">
                  {{ importModeLabel(item.mode) }} ·
                  {{ item.summary.creates + item.summary.updates }} 篇变更
                </span>
              </span>
              <UBadge
                :label="importStatusMeta(item.status).label"
                :color="importStatusMeta(item.status).color"
                :icon="importStatusMeta(item.status).icon"
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
              {{ batch ? importStatusMeta(batch.status).label : "等待上传" }}
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
