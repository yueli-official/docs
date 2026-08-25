<script setup lang="ts">
import type { CommentListResponse, CommentView } from "~/types";

const props = defineProps<{ documentId: string }>();
const { call } = useApi();
const items = ref<CommentView[]>([]);
const total = ref(0);
const loading = ref(true);
const errorMessage = ref("");
const replyTo = ref("");

async function load() {
  loading.value = true;
  errorMessage.value = "";
  try {
    const response = await call<CommentListResponse>(
      `/api/v1/docs/${props.documentId}/comments`,
      { query: { page: 1, size: 100 } },
    );
    items.value = response.items;
    total.value = response.total;
  } catch {
    errorMessage.value = "评论暂时无法加载。";
  } finally {
    loading.value = false;
  }
}

function onSubmitted() {
  replyTo.value = "";
  return load();
}

function toggleReply(id: string) {
  replyTo.value = replyTo.value === id ? "" : id;
}

function authorInitial(name: string) {
  return (name || "?").trim().charAt(0).toUpperCase();
}

function formatDate(value: string) {
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
onMounted(load);
</script>

<template>
  <section
    class="mt-14 border-t border-default pt-9"
    aria-labelledby="document-comments-title"
    data-document-comments
  >
    <div class="mb-6 flex items-center justify-between gap-3">
      <h2
        id="document-comments-title"
        class="font-display flex items-center gap-2 text-xl font-semibold text-highlighted"
      >
        <UIcon name="i-tabler-messages" class="size-5 text-primary" />
        评论
        <span v-if="total" class="text-base font-normal text-muted">
          {{ total }}
        </span>
      </h2>
    </div>

    <ClientOnly>
      <DocumentCommentForm
        :document-id="documentId"
        class="mb-9"
        @submitted="onSubmitted"
      />
      <template #fallback>
        <div class="mb-9 space-y-3" aria-label="正在准备评论表单">
          <USkeleton class="h-24 rounded-lg" />
          <USkeleton class="ml-auto h-8 w-24 rounded-md" />
        </div>
      </template>
    </ClientOnly>

    <div v-if="loading" class="space-y-5" aria-label="正在加载评论">
      <div v-for="index in 3" :key="index" class="flex gap-3">
        <USkeleton class="size-9 shrink-0 rounded-full" />
        <div class="min-w-0 flex-1 space-y-2">
          <USkeleton class="h-4 w-32" />
          <USkeleton class="h-4 w-full" />
          <USkeleton class="h-4 w-2/3" />
        </div>
      </div>
    </div>

    <UAlert
      v-else-if="errorMessage"
      color="error"
      variant="soft"
      icon="i-tabler-alert-circle"
      title="评论加载失败"
      :description="errorMessage"
      :actions="[{ label: '重新加载', onClick: load }]"
    />

    <div v-else-if="!items.length" class="py-8 text-center text-muted">
      <UIcon name="i-tabler-message-2" class="mx-auto size-7" />
      <p class="mt-2 text-sm">还没有评论，来写下第一条讨论吧。</p>
    </div>

    <ul v-else class="space-y-7">
      <li v-for="comment in items" :key="comment.id" data-comment-thread>
        <div class="flex gap-3">
          <UAvatar
            :src="comment.avatarUrl"
            :alt="comment.authorName"
            :text="authorInitial(comment.authorName)"
            size="sm"
            class="shrink-0"
          />
          <div class="min-w-0 flex-1">
            <div class="flex flex-wrap items-center gap-2 text-sm">
              <span class="font-medium text-highlighted">
                {{ comment.authorName }}
              </span>
              <span class="text-dimmed">·</span>
              <time class="text-muted" :datetime="comment.createdAt">
                {{ formatDate(comment.createdAt) }}
              </time>
            </div>
            <p class="mt-2 whitespace-pre-wrap text-sm leading-7 text-default">
              {{ comment.content }}
            </p>
            <UButton
              :label="replyTo === comment.id ? '取消回复' : '回复'"
              :icon="
                replyTo === comment.id
                  ? 'i-tabler-x'
                  : 'i-tabler-corner-down-right'
              "
              color="neutral"
              variant="ghost"
              size="xs"
              class="-ml-2 mt-1"
              @click="toggleReply(comment.id)"
            />

            <DocumentCommentForm
              v-if="replyTo === comment.id"
              :document-id="documentId"
              :parent-id="comment.id"
              compact
              class="mt-3 rounded-lg bg-elevated/40 p-3"
              @submitted="onSubmitted"
            />

            <ul
              v-if="comment.replies?.length"
              class="mt-4 space-y-4 border-l border-default pl-4"
            >
              <li
                v-for="reply in comment.replies"
                :key="reply.id"
                class="flex gap-2.5"
              >
                <UAvatar
                  :src="reply.avatarUrl"
                  :alt="reply.authorName"
                  :text="authorInitial(reply.authorName)"
                  size="2xs"
                  class="mt-0.5 shrink-0"
                />
                <div class="min-w-0 flex-1">
                  <div class="flex flex-wrap items-center gap-2 text-sm">
                    <span class="font-medium text-highlighted">
                      {{ reply.authorName }}
                    </span>
                    <span class="text-dimmed">·</span>
                    <time class="text-muted" :datetime="reply.createdAt">
                      {{ formatDate(reply.createdAt) }}
                    </time>
                  </div>
                  <p
                    class="mt-1 whitespace-pre-wrap text-sm leading-7 text-default"
                  >
                    {{ reply.content }}
                  </p>
                </div>
              </li>
            </ul>
          </div>
        </div>
      </li>
    </ul>
  </section>
</template>
