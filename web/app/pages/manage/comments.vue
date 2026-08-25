<script setup lang="ts">
import {
  CollectionPagination,
  CollectionTableToolbar,
} from "@yueli/ui/collection/pattern";
import { createDocsNotifier } from "~/utils/feedback";
import type {
  CommentAdminView,
  CommentStatus,
  ManageCommentsResponse,
} from "~/types";

definePageMeta({ layout: "manage", middleware: "auth" });
useSeoMeta({ title: "评论 · 控制台" });

const { call } = useApi();
const toast = createDocsNotifier(useToast());
const query = ref("");
const searchInput = ref("");
const status = ref<CommentStatus | "all">("all");
const page = ref(1);
const size = ref(20);
const selected = ref<string[]>([]);
const busy = ref("");
const emphasizedDelete = ref("");
const batchAction = ref<CommentStatus | "">("");
const batchBusy = ref(false);
const deleteOpen = ref(false);
const deleteTarget = ref<CommentAdminView | null>(null);
const mounted = ref(false);
let searchTimer: ReturnType<typeof setTimeout> | undefined;

const statusItems = [
  { label: "全部状态", value: "all" },
  { label: "待审核", value: "pending" },
  { label: "已通过", value: "approved" },
  { label: "垃圾", value: "spam" },
  { label: "回收站", value: "trash" },
];
const sizeItems = [20, 50, 100].map((value) => ({
  label: `${value}/页`,
  value,
}));
const batchItems = [
  { label: "通过", value: "approved" },
  { label: "标记垃圾", value: "spam" },
  { label: "移入回收站", value: "trash" },
];
const statusMeta: Record<
  CommentStatus,
  {
    label: string;
    color: "success" | "warning" | "error" | "neutral";
    icon: string;
  }
> = {
  approved: {
    label: "已通过",
    color: "success",
    icon: "i-tabler-circle-check",
  },
  pending: { label: "待审核", color: "warning", icon: "i-tabler-clock" },
  spam: {
    label: "垃圾",
    color: "error",
    icon: "i-tabler-alert-triangle",
  },
  trash: { label: "回收站", color: "neutral", icon: "i-tabler-trash" },
};

watch(searchInput, (value) => {
  if (searchTimer) clearTimeout(searchTimer);
  searchTimer = setTimeout(() => {
    query.value = value.trim();
    page.value = 1;
    selected.value = [];
  }, 300);
});
watch([status, size], () => {
  page.value = 1;
  selected.value = [];
});
onScopeDispose(() => {
  if (searchTimer) clearTimeout(searchTimer);
});
onMounted(() => {
  mounted.value = true;
});

const { data, pending, error, refresh } = await useAsyncData(
  "manage-comments",
  () =>
    call<ManageCommentsResponse>("/api/v1/manage/comments", {
      query: {
        status: status.value === "all" ? undefined : status.value,
        q: query.value || undefined,
        page: page.value,
        size: size.value,
      },
    }),
  {
    server: false,
    watch: [query, status, page, size],
    default: () => ({ items: [], total: 0, page: 1, size: 20 }),
  },
);

const items = computed(() => data.value?.items ?? []);
const total = computed(() => data.value?.total ?? 0);
const showSkeleton = computed(() => !mounted.value || pending.value);
const totalPages = computed(() =>
  Math.max(1, Math.ceil(total.value / size.value)),
);
const allSelected = computed(
  () =>
    items.value.length > 0 &&
    items.value.every((comment) => selected.value.includes(comment.id)),
);

watch(totalPages, (value) => {
  if (page.value > value) page.value = value;
});

function toggleOne(id: string, value: boolean) {
  selected.value = value
    ? [...new Set([...selected.value, id])]
    : selected.value.filter((item) => item !== id);
}

function togglePage(value: boolean) {
  const ids = items.value.map((comment) => comment.id);
  selected.value = value
    ? [...new Set([...selected.value, ...ids])]
    : selected.value.filter((id) => !ids.includes(id));
}

async function setCommentStatus(id: string, next: CommentStatus) {
  busy.value = id;
  try {
    await call(`/api/v1/manage/comments/${id}`, {
      method: "PATCH",
      body: { status: next },
    });
    selected.value = selected.value.filter((item) => item !== id);
    await refresh();
  } catch {
    toast.add({
      title: "评论状态未更新",
      description: "请刷新评论列表后重试。",
      color: "error",
    });
  } finally {
    busy.value = "";
  }
}

async function applyBatch() {
  if (!batchAction.value || !selected.value.length || batchBusy.value) return;
  batchBusy.value = true;
  const ids = [...selected.value];
  try {
    const results = await Promise.allSettled(
      ids.map((id) =>
        call(`/api/v1/manage/comments/${id}`, {
          method: "PATCH",
          body: { status: batchAction.value },
        }),
      ),
    );
    const failed = ids.filter(
      (_, index) => results[index]?.status === "rejected",
    );
    selected.value = failed;
    batchAction.value = "";
    await refresh();
    toast.add({
      title: failed.length ? "部分评论未处理" : "评论已批量处理",
      description: failed.length
        ? `${failed.length} 条失败项已保留选中。`
        : undefined,
      color: failed.length ? "warning" : "success",
    });
  } finally {
    batchBusy.value = false;
  }
}

function askDelete(comment: CommentAdminView) {
  deleteTarget.value = comment;
  deleteOpen.value = true;
}

function reloadComments() {
  void refresh();
}

function closeDelete() {
  deleteOpen.value = false;
}

async function confirmDelete() {
  if (!deleteTarget.value) return;
  const id = deleteTarget.value.id;
  busy.value = id;
  try {
    await call(`/api/v1/manage/comments/${id}`, { method: "DELETE" });
    deleteOpen.value = false;
    deleteTarget.value = null;
    selected.value = selected.value.filter((item) => item !== id);
    await refresh();
  } catch {
    toast.add({
      title: "评论未删除",
      description: "请刷新评论列表后重试。",
      color: "error",
    });
  } finally {
    busy.value = "";
  }
}

function publicDocumentLink(comment: CommentAdminView) {
  return {
    path: `/${comment.collectionSlug}/${comment.slugPath}`,
    query: {
      ...(comment.versionKey ? { version: comment.versionKey } : {}),
      ...(comment.locale !== "en" ? { locale: comment.locale } : {}),
    },
  };
}

function authorInitial(name: string) {
  return (name || "?").charAt(0).toUpperCase();
}

function emphasizeDelete(id: string) {
  emphasizedDelete.value = id;
}

function clearDeleteEmphasis(id: string) {
  if (emphasizedDelete.value === id) emphasizedDelete.value = "";
}

function formatDate(value: string) {
  const date = new Date(value);
  if (Number.isNaN(date.getTime())) return "未记录";
  return new Intl.DateTimeFormat("zh-CN", {
    year: "numeric",
    month: "2-digit",
    day: "2-digit",
    hour: "2-digit",
    minute: "2-digit",
  }).format(date);
}
</script>

<template>
  <ManagePage
    id="comments"
    title="评论"
    icon="i-tabler-messages"
    main-id="manage-main"
    body-class="w-full"
  >
    <section
      class="overflow-hidden rounded-xl border border-default bg-default"
      aria-label="评论治理列表"
      data-manage-comments
    >
      <CollectionTableToolbar
        v-model:search="searchInput"
        label="评论工具栏"
        search-placeholder="搜索评论、评论者或文档…"
        filter-label="筛选"
        :selection-count="selected.length"
      >
        <template #utilities>
          <USelect
            v-model="status"
            :items="statusItems"
            value-key="value"
            aria-label="评论状态"
            size="sm"
            class="w-32"
          />
        </template>
        <template #selection>
          <div class="flex min-w-0 items-center justify-between gap-2">
            <span class="text-sm font-medium text-highlighted">
              已选择 {{ selected.length }} 条评论
            </span>
            <div class="flex items-center gap-2">
              <USelect
                v-model="batchAction"
                :items="batchItems"
                value-key="value"
                placeholder="批量操作"
                size="xs"
                class="w-28"
              />
              <UButton
                label="应用"
                size="xs"
                :loading="batchBusy"
                :disabled="!batchAction"
                @click="applyBatch"
              />
              <UButton
                icon="i-tabler-x"
                aria-label="取消选择"
                color="neutral"
                variant="ghost"
                size="xs"
                square
                @click="selected = []"
              />
            </div>
          </div>
        </template>
      </CollectionTableToolbar>

      <div
        v-if="showSkeleton"
        class="grid gap-3 p-4 sm:p-5"
        aria-label="正在加载评论"
      >
        <USkeleton v-for="index in 7" :key="index" class="h-20 rounded-lg" />
      </div>

      <div v-else-if="error" class="px-5 py-10 text-center">
        <UIcon name="i-tabler-alert-circle" class="mx-auto size-7 text-error" />
        <p class="mt-2 text-sm text-muted">评论列表暂时无法加载。</p>
        <UButton
          label="重新加载"
          icon="i-tabler-refresh"
          color="neutral"
          variant="outline"
          size="sm"
          class="mt-4"
          @click="reloadComments"
        />
      </div>

      <div v-else-if="!items.length" class="px-5 py-12 text-center text-muted">
        <UIcon name="i-tabler-message-off" class="mx-auto size-7" />
        <p class="mt-2 text-sm">当前没有匹配的评论。</p>
      </div>

      <template v-else>
        <div class="border-b border-default px-4 py-2.5 sm:px-5">
          <div
            class="grid grid-cols-[1.5rem_minmax(0,1fr)_6.5rem] items-center gap-3 text-xs font-medium text-muted lg:grid-cols-[1.5rem_minmax(14rem,1.3fr)_minmax(9rem,0.8fr)_9rem_5.5rem_8.5rem_8.5rem]"
          >
            <UCheckbox
              :model-value="allSelected"
              aria-label="选择当前页评论"
              @update:model-value="togglePage(Boolean($event))"
            />
            <span>评论</span>
            <span class="hidden lg:block">来源</span>
            <span class="hidden lg:block">用户</span>
            <span class="hidden lg:block">状态</span>
            <span class="hidden lg:block">评论日期</span>
            <span class="text-right">操作</span>
          </div>
        </div>
        <div class="divide-y divide-default">
          <article
            v-for="comment in items"
            :key="comment.id"
            class="grid min-w-0 grid-cols-[1.5rem_minmax(0,1fr)_6.5rem] items-start gap-3 px-4 py-3 sm:px-5 lg:grid-cols-[1.5rem_minmax(14rem,1.3fr)_minmax(9rem,0.8fr)_9rem_5.5rem_8.5rem_8.5rem] lg:items-center"
            data-manage-comment-row
          >
            <UCheckbox
              :model-value="selected.includes(comment.id)"
              :aria-label="`选择 ${comment.authorName} 的评论`"
              @update:model-value="toggleOne(comment.id, Boolean($event))"
            />
            <div class="min-w-0">
              <p
                class="line-clamp-3 whitespace-pre-wrap text-sm leading-6 text-default"
              >
                {{ comment.content }}
              </p>
              <div
                class="mt-1.5 flex min-w-0 flex-wrap items-center gap-1.5 text-xs text-muted lg:hidden"
              >
                <UAvatar
                  :src="comment.avatarUrl"
                  :text="authorInitial(comment.authorName)"
                  alt=""
                  size="2xs"
                  class="shrink-0"
                />
                <span class="truncate font-medium text-highlighted">
                  {{ comment.authorName }}
                </span>
                <UBadge
                  :label="statusMeta[comment.status].label"
                  :color="statusMeta[comment.status].color"
                  :icon="statusMeta[comment.status].icon"
                  variant="subtle"
                  size="sm"
                />
                <span v-if="comment.parentId" class="text-dimmed">回复</span>
                <span class="text-dimmed">·</span>
                <time :datetime="comment.createdAt">
                  {{ formatDate(comment.createdAt) }}
                </time>
                <NuxtLink
                  :to="publicDocumentLink(comment)"
                  class="mt-1 flex basis-full items-center gap-1.5 truncate hover:text-primary"
                >
                  <UIcon name="i-tabler-file-text" class="size-3.5 shrink-0" />
                  <span class="truncate">{{ comment.documentTitle }}</span>
                </NuxtLink>
              </div>
            </div>
            <div class="hidden min-w-0 lg:block">
              <NuxtLink
                :to="publicDocumentLink(comment)"
                class="flex min-w-0 items-center gap-1.5 text-xs leading-5 text-muted hover:text-primary"
              >
                <UIcon name="i-tabler-file-text" class="size-3.5 shrink-0" />
                <span class="line-clamp-2">{{ comment.documentTitle }}</span>
              </NuxtLink>
            </div>
            <div class="hidden min-w-0 items-center gap-2 lg:flex">
              <UAvatar
                :src="comment.avatarUrl"
                :text="authorInitial(comment.authorName)"
                alt=""
                size="2xs"
                class="shrink-0"
              />
              <div class="min-w-0">
                <p class="truncate text-sm font-medium text-highlighted">
                  {{ comment.authorName }}
                </p>
                <p v-if="comment.parentId" class="text-xs text-dimmed">回复</p>
              </div>
            </div>
            <div class="hidden lg:flex lg:items-center">
              <UBadge
                :label="statusMeta[comment.status].label"
                :color="statusMeta[comment.status].color"
                :icon="statusMeta[comment.status].icon"
                variant="subtle"
              />
            </div>
            <div class="hidden text-xs text-muted lg:flex lg:items-center">
              <time :datetime="comment.createdAt">
                {{ formatDate(comment.createdAt) }}
              </time>
            </div>
            <div class="flex flex-nowrap items-center justify-end gap-1">
              <UButton
                v-if="comment.status === 'pending'"
                class="sm:hidden"
                icon="i-tabler-check"
                aria-label="通过"
                color="success"
                variant="soft"
                size="xs"
                square
                :loading="busy === comment.id"
                @click="setCommentStatus(comment.id, 'approved')"
              />
              <UButton
                v-if="comment.status === 'pending'"
                class="hidden sm:inline-flex"
                label="通过"
                icon="i-tabler-check"
                color="success"
                variant="soft"
                size="xs"
                :loading="busy === comment.id"
                @click="setCommentStatus(comment.id, 'approved')"
              />
              <UButton
                v-if="comment.status !== 'spam'"
                class="sm:hidden"
                icon="i-tabler-alert-triangle"
                aria-label="标记为垃圾"
                color="warning"
                variant="soft"
                size="xs"
                square
                :loading="busy === comment.id"
                @click="setCommentStatus(comment.id, 'spam')"
              />
              <UButton
                v-if="comment.status !== 'spam'"
                class="hidden sm:inline-flex"
                label="标记为垃圾"
                icon="i-tabler-alert-triangle"
                color="warning"
                variant="soft"
                size="xs"
                :loading="busy === comment.id"
                @click="setCommentStatus(comment.id, 'spam')"
              />
              <UButton
                :icon="
                  busy === comment.id ? 'i-tabler-loader-2' : 'i-tabler-trash'
                "
                :aria-label="`删除 ${comment.authorName} 的评论`"
                color="neutral"
                variant="ghost"
                size="xs"
                square
                :loading="busy === comment.id"
                class="transition-colors"
                :style="{
                  color:
                    emphasizedDelete === comment.id
                      ? 'var(--ui-error)'
                      : 'var(--ui-text-dimmed)',
                  backgroundColor:
                    emphasizedDelete === comment.id
                      ? 'color-mix(in oklab, var(--ui-error) 10%, transparent)'
                      : undefined,
                }"
                @mouseenter="emphasizeDelete(comment.id)"
                @mouseleave="clearDeleteEmphasis(comment.id)"
                @focus="emphasizeDelete(comment.id)"
                @blur="clearDeleteEmphasis(comment.id)"
                @click="askDelete(comment)"
              />
            </div>
          </article>
        </div>
      </template>

      <footer
        class="flex flex-col gap-3 border-t border-default px-4 py-3 sm:flex-row sm:items-center sm:justify-between sm:px-5"
      >
        <span class="text-xs text-muted">共 {{ total }} 条评论</span>
        <div class="flex items-center gap-2">
          <USelect
            v-model="size"
            :items="sizeItems"
            value-key="value"
            aria-label="每页评论数量"
            size="sm"
            class="w-24"
          />
          <CollectionPagination v-model="page" :total-pages="totalPages" />
        </div>
      </footer>
    </section>

    <UModal
      v-model:open="deleteOpen"
      title="删除评论"
      :description="`删除「${deleteTarget?.authorName || '用户'}」的评论及其回复。此操作不可撤销。`"
      :ui="{ footer: 'justify-end' }"
    >
      <template #footer>
        <UButton
          label="取消"
          color="neutral"
          variant="outline"
          :disabled="busy === deleteTarget?.id"
          @click="closeDelete"
        />
        <UButton
          label="确认删除"
          icon="i-tabler-trash"
          color="error"
          :loading="busy === deleteTarget?.id"
          @click="confirmDelete"
        />
      </template>
    </UModal>
  </ManagePage>
</template>
