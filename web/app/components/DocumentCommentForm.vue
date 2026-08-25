<script setup lang="ts">
import type { CommentView } from "~/types";

const props = withDefaults(
  defineProps<{
    documentId: string;
    parentId?: string;
    compact?: boolean;
  }>(),
  { parentId: "", compact: false },
);
const emit = defineEmits<{ submitted: [comment: CommentView] }>();

const route = useRoute();
const { call } = useApi();
const { loggedIn, login } = useAuth();
const content = ref("");
const busy = ref(false);
const errorMessage = ref("");
const submitted = ref(false);

function commentError(error: unknown) {
  const code = (error as { data?: { failure?: { code?: string } } } | null)?.data
    ?.failure?.code;
  if (code === "common.unauthorized") return "登录状态已失效，请重新登录后评论。";
  if (code === "docs.not_found") return "这篇文档已不可评论，请刷新页面。";
  if (code === "docs.invalid_input") return "评论需填写 1–2000 个字符。";
  return "评论暂时没有发布，请稍后重试。";
}

function openLogin() {
  void login(route.fullPath);
}

async function submit() {
  const value = content.value.trim();
  if (!value || busy.value) return;
  busy.value = true;
  errorMessage.value = "";
  submitted.value = false;
  try {
    const response = await call<{ comment: CommentView }>(
      `/api/v1/docs/${props.documentId}/comments`,
      {
        method: "POST",
        body: { content: value, parentId: props.parentId || undefined },
      },
    );
    content.value = "";
    submitted.value = true;
    emit("submitted", response.comment);
  } catch (error) {
    errorMessage.value = commentError(error);
  } finally {
    busy.value = false;
  }
}
</script>

<template>
  <div>
    <div
      v-if="!loggedIn"
      class="flex flex-wrap items-center justify-between gap-3 rounded-lg bg-elevated/45 px-4 py-3"
    >
      <p class="text-sm text-muted">登录后参与文档讨论。</p>
      <UButton
        label="登录后评论"
        icon="i-tabler-login-2"
        color="neutral"
        variant="outline"
        size="sm"
        @click="openLogin"
      />
    </div>

    <form v-else class="space-y-3" @submit.prevent="submit">
      <UAlert
        v-if="errorMessage"
        color="error"
        variant="soft"
        icon="i-tabler-alert-circle"
        title="评论未发布"
        :description="errorMessage"
      />
      <UAlert
        v-else-if="submitted"
        color="success"
        variant="soft"
        icon="i-tabler-circle-check"
        title="评论已发布"
      />
      <UTextarea
        v-model="content"
        :rows="compact ? 3 : 4"
        maxlength="2000"
        :placeholder="parentId ? '写下回复…' : '写下你的评论…'"
        class="w-full"
      />
      <div class="flex items-center justify-between gap-3">
        <span class="text-xs tabular-nums text-muted">
          {{ content.length }}/2000
        </span>
        <UButton
          type="submit"
          :label="parentId ? '发布回复' : '发表评论'"
          icon="i-tabler-send"
          size="sm"
          :loading="busy"
          :disabled="!content.trim()"
        />
      </div>
    </form>
  </div>
</template>
