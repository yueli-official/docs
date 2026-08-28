<script setup lang="ts">
import { CommentModerationCollection } from "@yueli/ui/comments/admin";
import type {
  CommentModerationCollectionActions,
  CommentModerationCollectionModel,
  CommentModerationItem,
  CommentModerationLifecycle,
} from "@yueli/ui/comments/admin";
import { useMinimumLoading } from "@yueli/ui/feedback";
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
const lifecycle = computed(() => status.value as CommentModerationLifecycle);
const sortOrder = ref<"asc" | "desc">("desc");
const page = ref(1);
const size = ref(20);
const selected = ref<string[]>([]);
const busy = ref("");
const emptyingTrash = ref(false);
const batchAction = ref<CommentStatus | "">("");
const batchBusy = ref(false);
const deleteOpen = ref(false);
const deleteTarget = ref<CommentAdminView | null>(null);
const mounted = ref(false);
let searchTimer: ReturnType<typeof setTimeout> | undefined;

const pageSizes = [20, 50, 100] as const;
const batchItems = computed(() =>
  lifecycle.value === "trash"
    ? [{ label: "恢复", value: "approved" }]
    : lifecycle.value === "spam"
      ? [
          { label: "恢复", value: "approved" },
          { label: "移入回收站", value: "trash" },
        ]
      : [
          { label: "通过", value: "approved" },
          { label: "标记垃圾", value: "spam" },
          { label: "移入回收站", value: "trash" },
        ],
);
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
    label: "垃圾评论",
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
watch([status, sortOrder, size], () => {
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
        sortBy: "createdAt",
        sortOrder: sortOrder.value,
        page: page.value,
        size: size.value,
      },
    }),
  {
    server: false,
    watch: [query, status, sortOrder, page, size],
    default: () => ({ items: [], total: 0, page: 1, size: 20 }),
  },
);

const items = computed(() => data.value?.items ?? []);
const total = computed(() => data.value?.total ?? 0);
const showSkeleton = useMinimumLoading(
  computed(() => !mounted.value || pending.value),
);
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

async function emptyTrash() {
  if (emptyingTrash.value) return false;
  emptyingTrash.value = true;
  try {
    for (let batch = 0; batch < 100; batch += 1) {
      const result = await call<ManageCommentsResponse>(
        "/api/v1/manage/comments",
        {
          query: {
            status: "trash",
            sortBy: "createdAt",
            sortOrder: "asc",
            page: 1,
            size: 100,
          },
        },
      );
      if (!result.items.length) break;
      const outcomes = await Promise.allSettled(
        result.items.map((comment) =>
          call(`/api/v1/manage/comments/${comment.id}`, { method: "DELETE" }),
        ),
      );
      if (outcomes.some((outcome) => outcome.status === "rejected")) {
        throw new Error("部分评论未能永久删除");
      }
      if (result.items.length < 100) break;
    }
    selected.value = [];
    await refresh();
    return true;
  } catch (error: any) {
    toast.add({
      title: "回收站未清空",
      description: error?.message || "请稍后重试。",
      color: "error",
    });
    return false;
  } finally {
    emptyingTrash.value = false;
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
function rowActionItems(comment: CommentAdminView) {
  const disabled = busy.value === comment.id;
  if (comment.status === "trash") {
    return [
      [
        {
          id: "restore",
          label: "恢复评论",
          icon: "i-tabler-restore",
          disabled,
          onSelect: () => void setCommentStatus(comment.id, "approved"),
        },
      ],
      [
        {
          id: "delete-permanently",
          label: "永久删除",
          icon: "i-tabler-trash-x",
          tone: "danger" as const,
          disabled,
          onSelect: () => askDelete(comment),
        },
      ],
    ];
  }
  return [
    [
      comment.status === "spam"
        ? {
            id: "restore",
            label: "恢复评论",
            icon: "i-tabler-restore",
            disabled,
            onSelect: () => void setCommentStatus(comment.id, "approved"),
          }
        : {
            id: "spam",
            label: "标记为垃圾",
            icon: "i-tabler-alert-triangle",
            disabled,
            onSelect: () => void setCommentStatus(comment.id, "spam"),
          },
      {
        id: "trash",
        label: "移入回收站",
        icon: "i-tabler-trash",
        disabled,
        onSelect: () => void setCommentStatus(comment.id, "trash"),
      },
    ],
  ];
}

function changeDateSort() {
  sortOrder.value = sortOrder.value === "asc" ? "desc" : "asc";
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
      ...(comment.locale ? { locale: comment.locale } : {}),
    },
  };
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
    hour12: false,
  }).format(date);
}

function submitSearch(value: string) {
  if (searchTimer) clearTimeout(searchTimer);
  searchInput.value = value;
  query.value = value.trim();
  page.value = 1;
  selected.value = [];
}
function changeLifecycle(value: CommentModerationLifecycle) {
  status.value = value;
}
function moderationItem(comment: CommentAdminView): CommentModerationItem {
  return {
    id: comment.id,
    content: comment.content,
    createdAt: comment.createdAt,
    authorName: comment.authorName || "匿名用户",
    avatarUrl: comment.avatarUrl,
    anonymous: !comment.userSub,
    reply: Boolean(comment.parentId),
    approve: comment.status === "pending",
    approving: busy.value === comment.id,
    actions: rowActionItems(comment),
    ...(comment.status === "approved"
      ? {}
      : { status: statusMeta[comment.status] }),
    source: {
      label: comment.documentTitle || "文档已删除",
      to: publicDocumentLink(comment),
      icon: "i-tabler-file-text",
    },
  };
}
const moderationModel = computed<CommentModerationCollectionModel>(() => ({
  search: searchInput.value,
  searchPlaceholder: "搜索评论、评论者或文档…",
  items: items.value.map(moderationItem),
  state: error.value ? "error" : showSkeleton.value ? "loading" : "ready",
  ...(error.value ? { errorMessage: "评论列表暂时无法加载。" } : {}),
  total: total.value,
  page: page.value,
  pageSize: size.value,
  pageSizes,
  activeFilterCount: 0,
  controls: [],
  sortOrder: sortOrder.value,
  lifecycle: lifecycle.value,
  emptyingTrash: emptyingTrash.value,
  selection: {
    enabled: true,
    count: selected.value.length,
    pageSelected: allSelected.value,
    pageIndeterminate:
      selected.value.length > 0 && !allSelected.value,
    isSelected: (id) => selected.value.includes(id),
  },
}));
const moderationActions: CommentModerationCollectionActions = {
  updateSearch: (value) => {
    searchInput.value = value;
  },
  search: submitSearch,
  controlChange: () => undefined,
  clearFilters: () => changeLifecycle("all"),
  retry: refresh,
  sort: changeDateSort,
  lifecycleChange: changeLifecycle,
  emptyTrash,
  approve: (id) => setCommentStatus(id, "approved"),
  pageChange: (value) => {
    page.value = value;
  },
  pageSizeChange: (value) => {
    size.value = value;
  },
  togglePage,
  toggleItem: toggleOne,
  clearSelection: () => {
    selected.value = [];
  },
};
</script>

<template>
  <ManagePage
    id="comments"
    title="评论"
    icon="i-tabler-messages"
    main-id="manage-main"
    body-class="w-full"
  >
    <CommentModerationCollection
      :model="moderationModel"
      :actions="moderationActions"
      :format-date="formatDate"
    >
      <template #bulk-actions>
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
      </template>
    </CommentModerationCollection>

    <UModal
      v-model:open="deleteOpen"
      title="永久删除评论"
      :description="`永久删除「${deleteTarget?.authorName || '用户'}」的评论及其回复？此操作不可撤销。`"
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
          label="永久删除"
          icon="i-tabler-trash-x"
          color="error"
          :loading="busy === deleteTarget?.id"
          @click="confirmDelete"
        />
      </template>
    </UModal>
  </ManagePage>
</template>
