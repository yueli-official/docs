<script setup lang="ts">
import { createDocsNotifier } from "~/utils/feedback";
import { ManageEmpty, SkeletonList } from "~/utils/manageComponents";
import { useMinimumLoading } from "@yueli/ui/feedback";
import type {
  DocsImportBatch,
  DocsImportDetailResponse,
  DocsImportItem,
} from "~/types";
import {
  importModeLabel,
  importStatusMeta,
} from "~/utils/docsImportPresentation.mjs";

definePageMeta({ layout: "manage" });
useSeoMeta({ title: "导入详情 · 控制台" });

const route = useRoute("/manage/import/[importId]");
const { call } = useApi();
const { can } = useMe();
const canManageImports = computed(() => can("docs.import.manage"));
const toast = createDocsNotifier(useToast());

const importId = computed(() => String(route.params.importId || ""));
const mounted = ref(false);
const rollbackConfirm = ref(false);
const rollbackBusy = ref(false);
onMounted(() => {
  mounted.value = true;
});

const { data, pending, error, refresh } = await useAsyncData(
  () => `docs-import-${importId.value}`,
  () =>
    canManageImports.value
      ? call<DocsImportDetailResponse>(`/api/v1/imports/docs/${importId.value}`)
      : Promise.resolve({
          batch: null as unknown as DocsImportBatch,
          items: [] as DocsImportItem[],
        }),
  {
    server: false,
    default: () => ({
      batch: null as unknown as DocsImportBatch,
      items: [] as DocsImportItem[],
    }),
  },
);

const batch = computed(() => data.value?.batch ?? null);
const items = computed(() => data.value?.items ?? []);
const summary = computed(() => batch.value?.summary ?? null);
const showSkeleton = useMinimumLoading(
  computed(() => !mounted.value || pending.value),
);
const loadError = computed(
  () => error.value as { data?: { message?: string }; message?: string } | null,
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
  ];
});

function toneClass(tone: string) {
  if (tone === "primary") return "bg-primary/10 text-primary ring-primary/20";
  if (tone === "success") return "bg-success/10 text-success ring-success/20";
  if (tone === "warning") return "bg-warning/10 text-warning ring-warning/20";
  return "bg-elevated text-muted ring-default";
}

function actionMeta(action: string) {
  const map: Record<
    string,
    {
      label: string;
      color: "primary" | "neutral" | "warning" | "error";
      icon: string;
    }
  > = {
    create: { label: "创建", color: "primary", icon: "i-tabler-file-plus" },
    update: { label: "更新", color: "neutral", icon: "i-tabler-refresh" },
    archive: { label: "归档", color: "warning", icon: "i-tabler-archive" },
    skip: {
      label: "跳过",
      color: "neutral",
      icon: "i-tabler-player-skip-forward",
    },
    conflict: {
      label: "冲突",
      color: "error",
      icon: "i-tabler-alert-triangle",
    },
    error: { label: "错误", color: "error", icon: "i-tabler-alert-circle" },
  };
  return (
    map[action] ?? { label: action, color: "neutral", icon: "i-tabler-circle" }
  );
}

async function rollback() {
  if (!canManageImports.value || !batch.value?.id) return;
  rollbackBusy.value = true;
  try {
    await call(`/api/v1/imports/docs/${batch.value.id}/rollback`, {
      method: "POST",
    });
    rollbackConfirm.value = false;
    await refresh();
  } catch (err: any) {
    toast.add({
      title: "回滚失败",
      description: err?.data?.message || "请稍后重试",
      color: "error",
    });
  } finally {
    rollbackBusy.value = false;
  }
}
</script>

<template>
  <ManagePage
    id="import-detail"
    title="导入详情"
    icon="i-tabler-list-details"
    main-id="manage-main"
    body-class="w-full"
  >
    <template #trailing>
      <span class="truncate font-mono text-xs text-muted">{{ importId }}</span>
    </template>
    <template #actions>
      <UButton
        v-if="canManageImports"
        icon="i-tabler-file-import"
        label="继续导入"
        color="neutral"
        variant="outline"
        to="/manage/import"
      />
    </template>

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
          当前角色不能查看或回滚导入批次。请联系管理员调整角色能力。
        </p>
      </div>
    </div>

    <SkeletonList v-else-if="showSkeleton" :rows="8" />

    <div
      v-else-if="loadError"
      class="rounded-lg border border-default bg-default p-8"
    >
      <div class="mx-auto max-w-md text-center">
        <span
          class="mx-auto grid size-12 place-items-center rounded-lg bg-error/10 text-error"
        >
          <UIcon name="i-tabler-alert-circle" class="size-6" />
        </span>
        <h2 class="mt-4 text-lg font-semibold text-highlighted">
          导入批次加载失败
        </h2>
        <p class="mt-2 text-sm leading-6 text-muted">
          {{ loadError.data?.message || loadError.message || "请稍后重试" }}
        </p>
        <UButton
          class="mt-5"
          icon="i-tabler-refresh"
          label="重试"
          color="neutral"
          variant="soft"
          @click="() => refresh()"
        />
      </div>
    </div>

    <template v-else-if="batch">
      <section
        class="mb-4 overflow-hidden rounded-lg border border-default bg-default"
      >
        <div class="border-b border-default bg-elevated/35 px-5 py-4">
          <div class="flex flex-wrap items-center justify-between gap-3">
            <div>
              <h2 class="text-base font-semibold text-highlighted">批次状态</h2>
              <p class="mt-1 text-sm text-muted">
                模式 {{ importModeLabel(batch.mode) }} · 默认语言
                {{ batch.defaultLocale }}
              </p>
            </div>
            <UBadge
              :label="importStatusMeta(batch.status).label"
              :color="importStatusMeta(batch.status).color"
              :icon="importStatusMeta(batch.status).icon"
              variant="subtle"
            />
          </div>
        </div>

        <div class="grid sm:grid-cols-2 xl:grid-cols-4">
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
      </section>

      <div class="grid gap-4 xl:grid-cols-[minmax(0,1fr)_340px]">
        <section
          class="min-w-0 overflow-hidden rounded-lg border border-default bg-default"
        >
          <div
            class="hidden grid-cols-[110px_minmax(220px,1fr)_150px_180px] gap-3 border-b border-default bg-elevated/45 px-4 py-2.5 text-xs font-medium text-muted lg:grid"
          >
            <span>动作</span>
            <span>文档</span>
            <span>语言 / 版本</span>
            <span>源文件</span>
          </div>

          <ManageEmpty
            v-if="!items.length"
            icon="i-tabler-file-search"
            text="没有导入条目"
          />

          <div v-else class="divide-y divide-default">
            <div
              v-for="item in items"
              :key="item.id"
              class="grid gap-3 px-4 py-3 lg:grid-cols-[110px_minmax(220px,1fr)_150px_180px] lg:items-center"
            >
              <UBadge
                :label="actionMeta(item.action).label"
                :color="actionMeta(item.action).color"
                :icon="actionMeta(item.action).icon"
                variant="subtle"
                size="sm"
              />
              <div class="min-w-0">
                <p class="truncate text-sm font-medium text-highlighted">
                  {{ item.title || item.path }}
                </p>
                <p class="truncate font-mono text-xs text-muted">
                  {{ item.path }}
                </p>
                <div v-if="item.issues?.length" class="mt-2 space-y-1">
                  <p
                    v-for="(issue, index) in item.issues"
                    :key="`${issue.code}-${index}`"
                    class="line-clamp-1 text-xs text-warning"
                  >
                    {{ issue.message }}
                  </p>
                </div>
              </div>
              <p class="font-mono text-xs text-muted">
                {{ item.locale }} / {{ item.versionKey }}
              </p>
              <p class="truncate font-mono text-xs text-muted">
                {{ item.sourceMarkdownPath }}
              </p>
            </div>
          </div>
        </section>

        <aside
          class="rounded-lg border border-default bg-default p-4 xl:sticky xl:top-24 xl:self-start"
        >
          <div class="mb-4 flex items-start justify-between gap-3">
            <div>
              <p class="text-xs font-medium uppercase tracking-wide text-muted">
                批次操作
              </p>
              <h2 class="mt-1 text-base font-semibold text-highlighted">
                {{ importStatusMeta(batch.status).label }}
              </h2>
            </div>
            <span
              class="grid size-10 place-items-center rounded-lg bg-primary/10 text-primary"
            >
              <UIcon name="i-tabler-history" class="size-5" />
            </span>
          </div>

          <div
            class="space-y-3 rounded-lg border border-default bg-elevated/35 p-3 text-sm"
          >
            <div class="flex items-center justify-between gap-3">
              <span class="text-muted">条目</span>
              <span class="font-mono text-xs text-default">{{
                items.length
              }}</span>
            </div>
            <div class="flex items-center justify-between gap-3">
              <span class="text-muted">阻塞</span>
              <span class="font-mono text-xs text-default">{{
                summary?.blocking ? "是" : "否"
              }}</span>
            </div>
            <div class="flex items-center justify-between gap-3">
              <span class="text-muted">警告</span>
              <span class="font-mono text-xs text-default">{{
                summary?.warnings ?? 0
              }}</span>
            </div>
          </div>

          <USeparator class="my-4" />

          <UAlert
            v-if="batch.status === 'failed' && batch.errorMessage"
            class="mb-4"
            color="error"
            variant="soft"
            icon="i-tabler-alert-circle"
            title="导入执行失败"
            :description="batch.errorMessage"
          />

          <div
            v-if="batch.status === 'completed'"
            class="rounded-lg border border-warning/30 bg-warning/5 p-3"
          >
            <div v-if="!rollbackConfirm" class="space-y-3">
              <div>
                <p class="text-sm font-medium text-highlighted">回滚这次导入</p>
                <p class="mt-1 text-xs leading-5 text-muted">
                  创建的文档会被删除，更新和归档会尽量恢复到导入前快照。
                </p>
              </div>
              <UButton
                label="回滚导入"
                icon="i-tabler-history"
                color="warning"
                variant="soft"
                block
                @click="
                  () => {
                    rollbackConfirm = true;
                  }
                "
              />
            </div>
            <div v-else class="space-y-3">
              <p class="text-sm text-highlighted">确定回滚这个导入批次？</p>
              <div class="flex justify-end gap-2">
                <UButton
                  label="取消"
                  color="neutral"
                  variant="ghost"
                  size="sm"
                  @click="
                    () => {
                      rollbackConfirm = false;
                    }
                  "
                />
                <UButton
                  label="确认回滚"
                  icon="i-tabler-history"
                  color="warning"
                  size="sm"
                  :loading="rollbackBusy"
                  @click="rollback"
                />
              </div>
            </div>
          </div>

          <UAlert
            v-else
            color="neutral"
            variant="soft"
            icon="i-tabler-info-circle"
            title="当前状态不可回滚"
            :description="
              batch.status === 'rolled_back'
                ? '这个批次已经回滚。'
                : '只有已完成的导入批次可以回滚。'
            "
          />
        </aside>
      </div>
    </template>
  </ManagePage>
</template>
