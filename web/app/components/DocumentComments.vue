<script setup lang="ts">
import { PublicCommentThread } from "@yueli/ui/comments";
import type {
  PublicCommentDraft,
  PublicCommentMessages,
  PublicCommentOrder,
  PublicCommentState,
} from "@yueli/ui/comments";
import type { CommentListResponse, CommentView } from "~/types";

const props = defineProps<{ documentId: string }>();
const route = useRoute();
const { call } = useApi();
const { loggedIn, user, login } = useAuth();

const items = ref<CommentView[]>([]);
const total = ref(0);
const state = ref<PublicCommentState>("loading");
const order = ref<PublicCommentOrder>("asc");
const viewer = computed(() => ({
  authenticated: loggedIn.value,
  name: user.value?.name || user.value?.email || "",
  avatarUrl: user.value?.avatar || "",
}));
const messages: PublicCommentMessages = {
  count: (count) => `${count} 条评论`,
  replies: (count) => `${count} 条回复`,
  sort: "评论排序",
  loading: "正在加载评论",
  oldest: "最早",
  newest: "最新",
  reply: "回复",
  cancelReply: "取消回复",
  anonymous: "匿名用户",
  empty: "还没有评论，来写下第一条讨论吧",
  closed: "本文档已关闭评论",
  loadError: "评论加载失败",
  retry: "重新加载",
  writeComment: "写下你的评论…",
  writeReply: "写下回复…",
  authorName: "昵称 *",
  authorEmail: "邮箱（选填，不公开）",
  anonymousHint: "匿名评论需要审核",
  login: "登录后评论",
  submit: "发表评论",
  submitReply: "回复",
  submitted: "评论已发布",
  pending: "评论已提交，待审核后显示",
  submitError: "评论暂时没有发布，请稍后重试",
  nameRequired: "请填写昵称",
};

async function load() {
  state.value = "loading";
  try {
    const response = await call<CommentListResponse>(
      `/api/v1/docs/${props.documentId}/comments`,
      { query: { page: 1, size: 100, sortOrder: order.value } },
    );
    items.value = response.items;
    total.value = response.total;
    state.value = "ready";
  } catch {
    state.value = "error";
  }
}

function commentError(error: unknown) {
  const code = (error as { data?: { failure?: { code?: string } } } | null)
    ?.data?.failure?.code;
  if (code === "common.unauthorized")
    return "登录状态已失效，请重新登录后评论。";
  if (code === "docs.not_found") return "这篇文档已不可评论，请刷新页面。";
  if (code === "docs.invalid_input") return "评论需填写 1–2000 个字符。";
  return messages.submitError;
}

async function submitComment(draft: PublicCommentDraft) {
  try {
    await call<{ comment: CommentView }>(
      `/api/v1/docs/${props.documentId}/comments`,
      {
        method: "POST",
        body: { content: draft.content, parentId: draft.parentId },
      },
    );
    await load();
    return { pending: false };
  } catch (error) {
    throw new Error(commentError(error));
  }
}

function openLogin() {
  void login(route.fullPath);
}

function formatCommentTime(value: string) {
  const date = new Date(value);
  if (Number.isNaN(date.getTime())) return "刚刚";
  return new Intl.DateTimeFormat("zh-CN", {
    year: "numeric",
    month: "2-digit",
    day: "2-digit",
    hour: "2-digit",
    minute: "2-digit",
  }).format(date);
}

watch(() => props.documentId, load);
watch(order, load);
onMounted(load);
</script>

<template>
  <section class="mt-14" data-document-comments>
    <PublicCommentThread
      v-model:order="order"
      :comments="items"
      :total="total"
      :state="state"
      :viewer="viewer"
      :messages="messages"
      :format-time="formatCommentTime"
      :submit="submitComment"
      :login="openLogin"
      :retry="load"
      :max-length="2000"
      input-position="bottom"
    />
  </section>
</template>
