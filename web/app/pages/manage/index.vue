<script setup lang="ts">
import { SkeletonList } from "~/utils/manageComponents";
import { useMinimumLoading } from "@yueli/ui/feedback";
import type {
  CollectionList,
  ManageDocListItem,
  ManageDocsResponse,
} from "~/types";
import { docManageRoute } from "~/utils/docsManageRoutes.mjs";

definePageMeta({ layout: "manage" });
useSeoMeta({ title: "控制台" });

interface DashboardData {
  collections: CollectionList | null;
  recent: ManageDocsResponse | null;
  issues: ManageDocsResponse | null;
}

const { call } = useApi();
const { can, isAdministrator } = useMe();
const canReadDocs = computed(() => can("docs.document.read"));
const canCreateDocs = computed(() => can("docs.document.create"));
const canManageCollections = computed(() => can("docs.collection.manage"));
const canManageImports = computed(() => can("docs.import.manage"));
const canManageSiteSettings = computed(() => can("docs.site_settings.manage"));
const canManageAssetSettings = computed(() =>
  can("docs.asset_settings.manage"),
);
const mounted = ref(false);
onMounted(() => {
  mounted.value = true;
});

async function loadOptional<T>(request: Promise<T>): Promise<T | null> {
  try {
    return await request;
  } catch {
    return null;
  }
}

const { data, pending, refresh } = await useAsyncData(
  "manage-overview",
  async (): Promise<DashboardData> => {
    const [collections, recent, issues] = await Promise.all([
      canManageCollections.value
        ? loadOptional(call<CollectionList>("/api/v1/collections"))
        : null,
      canReadDocs.value
        ? loadOptional(
            call<ManageDocsResponse>("/api/v1/manage/docs", {
              query: {
                status: "all",
                quality: "all",
                sort: "updatedAt",
                direction: "desc",
                page: 1,
                size: 6,
              },
            }),
          )
        : null,
      canReadDocs.value
        ? loadOptional(
            call<ManageDocsResponse>("/api/v1/manage/docs", {
              query: {
                status: "all",
                quality: "issues",
                sort: "updatedAt",
                direction: "desc",
                page: 1,
                size: 5,
              },
            }),
          )
        : null,
    ]);
    return { collections, recent, issues };
  },
  {
    server: false,
    default: () => ({
      collections: null,
      recent: null,
      issues: null,
    }),
  },
);

const showSkeleton = useMinimumLoading(
  computed(() => !mounted.value || pending.value),
);

const recentDocs = computed(() => data.value?.recent?.items ?? []);
const issueDocs = computed(() => data.value?.issues?.items ?? []);
const collectionCount = computed(
  () => data.value?.collections?.items.length ?? null,
);
const documentDataUnavailable = computed(
  () => canReadDocs.value && (!data.value?.recent || !data.value?.issues),
);
const collectionDataUnavailable = computed(
  () => canManageCollections.value && !data.value?.collections,
);
const workspaceDegraded = computed(
  () => documentDataUnavailable.value || collectionDataUnavailable.value,
);
const quickActions = computed(() => [
  ...(canReadDocs.value
    ? [
        {
          label: "管理文档",
          description: "查找、筛选与维护内容",
          icon: "i-tabler-files",
          to: "/manage/docs",
        },
      ]
    : []),
  ...(canManageCollections.value
    ? [
        {
          label: "管理文档集",
          description: "调整内容结构与分组",
          icon: "i-tabler-stack-2",
          to: "/manage/collections",
        },
      ]
    : []),
  ...(canManageImports.value
    ? [
        {
          label: "批量导入",
          description: "一次导入多篇文档",
          icon: "i-tabler-file-import",
          to: "/manage/import",
        },
      ]
    : []),
  ...(canManageSiteSettings.value
    ? [
        {
          label: "站点设置",
          description: "维护站点展示信息",
          icon: "i-tabler-settings",
          to: "/manage/home",
        },
      ]
    : []),
  ...(canManageAssetSettings.value
    ? [
        {
          label: "资源配置",
          description: "检查存储与用途规则",
          icon: "i-tabler-database-cog",
          to: "/manage/assets",
        },
      ]
    : []),
  ...(isAdministrator.value
    ? [
        {
          label: "权限与申请",
          description: "管理作者能力与申请",
          icon: "i-tabler-shield-lock",
          to: "/manage/authorization",
        },
      ]
    : []),
]);

function docLink(doc: ManageDocListItem) {
  return docManageRoute(
    doc.collectionSlug,
    doc.slugPath.split("/").filter(Boolean),
  );
}

function issueLabel(doc: ManageDocListItem) {
  const issues: string[] = [];
  if (!doc.title.trim()) issues.push("缺标题");
  if (!doc.slug.trim()) issues.push("缺路径");
  if (!doc.excerpt.trim()) issues.push("缺摘要");
  return issues.join(" · ") || "需要复核内容信息";
}

function statusLabel(status: string) {
  return (
    { draft: "草稿", published: "已发布", archived: "已归档" }[status] ?? status
  );
}

function formatUpdatedAt(value: string) {
  if (!value) return "未记录";
  return new Intl.DateTimeFormat("zh-CN", {
    month: "2-digit",
    day: "2-digit",
    hour: "2-digit",
    minute: "2-digit",
  }).format(new Date(value));
}
</script>

<template>
  <YAdminPage
    id="dashboard"
    title="控制台"
    icon="i-tabler-dashboard"
    main-id="manage-main"
    body-class="mx-auto w-full max-w-screen-2xl space-y-4"
  >
    <template #actions>
      <UButton
        v-if="canCreateDocs"
        to="/manage/docs/new"
        icon="i-tabler-plus"
        label="新建文档"
      />
    </template>

    <div
      :class="[
        'grid min-w-0 gap-4',
        canReadDocs && 'lg:grid-cols-[minmax(0,1fr)_20rem]',
      ]"
    >
      <div v-if="canReadDocs" class="min-w-0 space-y-4">
        <section
          aria-labelledby="pending-docs-title"
          class="overflow-hidden rounded-xl border border-default bg-default"
        >
          <div class="border-b border-default px-4 py-3 sm:px-5">
            <h2
              id="pending-docs-title"
              class="text-sm font-semibold text-highlighted"
            >
              待完善文档
            </h2>
            <p class="mt-0.5 text-xs text-muted">
              先处理缺少摘要、标题或路径的内容。
            </p>
          </div>
          <div class="p-4 sm:p-5">
            <div v-if="showSkeleton" class="grid gap-2">
              <USkeleton
                v-for="item in 3"
                :key="item"
                class="h-14 rounded-lg"
              />
            </div>
            <div
              v-else-if="!data?.issues"
              class="flex flex-col items-start gap-3 rounded-lg bg-error/10 px-3 py-3 text-sm text-error sm:flex-row sm:items-center sm:justify-between"
            >
              <span class="inline-flex items-center gap-2">
                <UIcon name="i-tabler-alert-circle" class="size-4" />
                待完善文档暂时无法加载
              </span>
              <UButton
                label="重试"
                icon="i-tabler-refresh"
                color="error"
                variant="soft"
                size="xs"
                @click="() => refresh()"
              />
            </div>
            <div v-else-if="issueDocs.length" class="divide-y divide-default">
              <NuxtLink
                v-for="doc in issueDocs"
                :key="doc.id"
                :to="docLink(doc)"
                class="group grid min-h-14 grid-cols-[auto_minmax(0,1fr)_auto] items-center gap-3 py-2.5 first:pt-0 last:pb-0"
              >
                <span
                  class="grid size-8 place-items-center rounded-lg bg-warning/10 text-warning"
                >
                  <UIcon name="i-tabler-alert-circle" class="size-4" />
                </span>
                <div class="min-w-0">
                  <p class="truncate text-sm font-medium text-highlighted">
                    {{ doc.title || "未命名文档" }}
                  </p>
                  <p class="truncate text-xs text-muted">
                    {{ doc.collectionTitle }} · {{ issueLabel(doc) }}
                  </p>
                </div>
                <UIcon
                  name="i-tabler-chevron-right"
                  class="size-4 text-dimmed transition group-hover:translate-x-0.5"
                />
              </NuxtLink>
            </div>
            <div
              v-else
              class="flex items-center gap-2 rounded-lg bg-success/10 px-3 py-2.5 text-sm text-success"
            >
              <UIcon name="i-tabler-circle-check" class="size-4" />
              当前没有待完善的文档
            </div>
          </div>
        </section>

        <section
          aria-labelledby="recent-docs-title"
          class="overflow-hidden rounded-xl border border-default bg-default"
        >
          <div class="border-b border-default px-4 py-3 sm:px-5">
            <h2
              id="recent-docs-title"
              class="text-sm font-semibold text-highlighted"
            >
              最近更新
            </h2>
            <p class="mt-0.5 text-xs text-muted">继续处理近期有变更的文档。</p>
          </div>
          <SkeletonList v-if="showSkeleton" :rows="5" class="p-4" />
          <div
            v-else-if="!data?.recent"
            class="flex flex-col items-start gap-3 p-4 text-sm text-error sm:flex-row sm:items-center sm:justify-between sm:px-5"
          >
            <span class="inline-flex items-center gap-2">
              <UIcon name="i-tabler-alert-circle" class="size-4" />
              最近更新暂时无法加载
            </span>
            <UButton
              label="重试"
              icon="i-tabler-refresh"
              color="error"
              variant="soft"
              size="xs"
              @click="() => refresh()"
            />
          </div>
          <div v-else-if="recentDocs.length" class="divide-y divide-default">
            <NuxtLink
              v-for="doc in recentDocs"
              :key="doc.id"
              :to="docLink(doc)"
              class="group grid min-h-14 grid-cols-[minmax(0,1fr)_auto_auto] items-center gap-4 px-4 py-3 transition hover:bg-elevated/40 sm:px-5"
            >
              <div class="min-w-0">
                <p
                  class="truncate text-sm font-medium text-highlighted group-hover:text-primary"
                >
                  {{ doc.title }}
                </p>
                <p class="truncate font-mono text-xs text-muted">
                  /{{ doc.slugPath }}
                </p>
              </div>
              <span class="hidden text-xs text-muted sm:inline">{{
                statusLabel(doc.status)
              }}</span>
              <time
                class="text-xs tabular-nums text-muted"
                :datetime="doc.updatedAt"
                >{{ formatUpdatedAt(doc.updatedAt) }}</time
              >
            </NuxtLink>
          </div>
          <div v-else class="p-8 text-center text-sm text-muted">
            还没有文档，先创建第一篇内容。
          </div>
        </section>
      </div>

      <aside class="min-w-0 space-y-4">
        <section
          aria-labelledby="workspace-status-title"
          class="rounded-xl border border-default bg-default p-4"
        >
          <div class="mb-4 flex items-start justify-between gap-3">
            <div>
              <h2
                id="workspace-status-title"
                class="text-sm font-semibold text-highlighted"
              >
                工作区状态
              </h2>
              <p class="mt-0.5 text-xs text-muted">
                仅反映当前账号可访问的管理数据。
              </p>
            </div>
            <UBadge
              v-if="!showSkeleton"
              :color="workspaceDegraded ? 'warning' : 'success'"
              variant="soft"
              :label="workspaceDegraded ? '部分不可用' : '可工作'"
            />
          </div>
          <div v-if="showSkeleton" class="grid gap-3">
            <USkeleton v-for="item in 3" :key="item" class="h-5 rounded-md" />
          </div>
          <dl v-else class="grid gap-3 text-sm">
            <div
              v-if="canReadDocs"
              class="flex items-center justify-between gap-3"
            >
              <dt class="text-muted">文档管理</dt>
              <dd
                class="inline-flex items-center gap-1.5"
                :class="
                  documentDataUnavailable ? 'text-warning' : 'text-success'
                "
              >
                <span
                  class="size-1.5 rounded-full"
                  :class="documentDataUnavailable ? 'bg-warning' : 'bg-success'"
                />
                {{ documentDataUnavailable ? "不可用" : "正常" }}
              </dd>
            </div>
            <div
              v-if="canManageCollections"
              class="flex items-center justify-between gap-3"
            >
              <dt class="text-muted">文档集</dt>
              <dd
                class="font-medium tabular-nums"
                :class="
                  collectionDataUnavailable
                    ? 'text-warning'
                    : 'text-highlighted'
                "
              >
                {{ collectionDataUnavailable ? "不可用" : collectionCount }}
              </dd>
            </div>
            <div class="flex items-center justify-between gap-3">
              <dt class="text-muted">可用操作</dt>
              <dd class="font-medium tabular-nums text-highlighted">
                {{ quickActions.length }}
              </dd>
            </div>
          </dl>
          <UButton
            v-if="!showSkeleton && workspaceDegraded"
            label="重新检查"
            icon="i-tabler-refresh"
            color="neutral"
            variant="outline"
            size="xs"
            class="mt-4"
            @click="() => refresh()"
          />
        </section>

        <section
          aria-labelledby="quick-actions-title"
          class="rounded-xl border border-default bg-default p-4"
        >
          <div class="mb-2">
            <h2
              id="quick-actions-title"
              class="text-sm font-semibold text-highlighted"
            >
              快捷操作
            </h2>
            <p class="mt-0.5 text-xs text-muted">进入常用的内容管理任务。</p>
          </div>
          <div class="divide-y divide-default">
            <NuxtLink
              v-for="action in quickActions"
              :key="action.to"
              :to="action.to"
              class="group flex items-center gap-3 py-3 first:pt-2 last:pb-0"
            >
              <span
                class="grid size-8 shrink-0 place-items-center rounded-lg bg-elevated text-muted transition group-hover:bg-primary/10 group-hover:text-primary"
              >
                <UIcon :name="action.icon" class="size-4" />
              </span>
              <span class="min-w-0 flex-1">
                <span
                  class="block text-sm font-medium text-highlighted group-hover:text-primary"
                  >{{ action.label }}</span
                >
                <span class="block truncate text-xs text-muted">{{
                  action.description
                }}</span>
              </span>
              <UIcon
                name="i-tabler-chevron-right"
                class="size-4 shrink-0 text-dimmed transition group-hover:translate-x-0.5 group-hover:text-primary"
              />
            </NuxtLink>
          </div>
        </section>
      </aside>
    </div>
  </YAdminPage>
</template>
