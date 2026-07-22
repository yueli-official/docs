<script setup lang="ts">
import { createPlatformNotifier } from "@platform/ui/feedback";
import { ManageEmpty } from "@platform/manage/components";
import type {
  DocsImportBatch,
  DocsImportSummary,
  DocsImportUploadResponse,
} from "~/types";

definePageMeta({ layout: "manage" });
useSeoMeta({ title: "批量导入 · 控制台" });

const { call } = useApi();
const toast = createPlatformNotifier(useToast());

const file = ref<File | null>(null);
const uploading = ref(false);
const confirming = ref(false);
const errorMessage = ref("");
const batch = ref<DocsImportBatch | null>(null);
const summary = ref<DocsImportSummary | null>(null);

const canUpload = computed(() => Boolean(file.value && !uploading.value));
const canConfirm = computed(() =>
  Boolean(
    batch.value?.id &&
    batch.value.status === "checked" &&
    !summary.value?.blocking,
  ),
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
  if (!file.value) return;
  if (!file.value.name.toLowerCase().endsWith(".zip")) {
    errorMessage.value = "请选择 ZIP 文件";
    return;
  }
  uploading.value = true;
  errorMessage.value = "";
  try {
    const body = new FormData();
    body.append("file", file.value);
    const res = await call<DocsImportUploadResponse>("/api/v1/imports/docs", {
      method: "POST",
      body,
    });
    batch.value = res.batch;
    summary.value = res.summary;
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
  if (!batch.value) return;
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
      description: err?.data?.message || "请检查预检结果后重试",
      color: "error",
    });
  } finally {
    confirming.value = false;
  }
}
</script>

<template>
  <YAdminPage
    id="import"
    title="批量导入"
    icon="i-tabler-file-import"
    main-id="manage-main"
    body-class="mx-auto w-full max-w-screen-2xl"
  >
    <div class="grid gap-4 xl:grid-cols-[minmax(0,1fr)_360px]">
      <section class="min-w-0 space-y-4">
        <div class="rounded-lg border border-default bg-default p-5">
          <div class="mb-4 flex flex-wrap items-start justify-between gap-3">
            <div>
              <h2 class="text-base font-semibold text-highlighted">
                选择导入包
              </h2>
              <p class="mt-1 text-sm text-muted">
                仅支持标准 ZIP 包，导入前会检查
                manifest、Markdown、图片与内部链接。
              </p>
            </div>
            <UBadge
              label="ZIP only"
              icon="i-tabler-file-zip"
              variant="subtle"
            />
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
      </section>

      <aside
        class="rounded-lg border border-default bg-default p-4 xl:sticky xl:top-24 xl:self-start"
      >
        <div class="mb-4 flex items-start justify-between gap-3">
          <div>
            <p class="text-xs font-medium uppercase tracking-wide text-muted">
              导入状态
            </p>
            <h2 class="mt-1 text-base font-semibold text-highlighted">
              {{ batch?.status || "等待上传" }}
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
              batch?.mode || "-"
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
  </YAdminPage>
</template>
