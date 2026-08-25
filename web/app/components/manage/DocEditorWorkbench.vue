<script setup lang="ts">
import { createDocsNotifier } from "~/utils/feedback";
import { useActionFeedback } from "@yueli/ui/feedback";
import { ActionFeedbackButton } from "@yueli/ui/feedback/pattern";
import type {
  CollectionVersion,
  CollectionVersionsResponse,
  CollectionView,
  CollectionManageTree,
  DocDetail,
  DocDetailResponse,
} from "~/types";
import {
  buildDocTree,
  docManageRoute,
  findDocBySlugPath,
  findDocSlugPathById,
  normalizeDocSlugPath,
} from "~/utils/docsManageRoutes.mjs";

// Doc editor (manage): create (/manage/docs/new) or edit by semantic path.
// Fields: collection (required, new only) + parent + slug + title + body + status.
const route = useRoute();
const { call } = useApi();
const toast = createDocsNotifier(useToast());
const ROOT = "__root__"; // USelect/Reka SelectItem cannot use an empty-string value.
const routeId = computed(() => String(route.params.id ?? ""));
const semanticCollectionSlug = computed(() =>
  String(route.params.collectionSlug ?? ""),
);
const semanticDocPath = computed(() =>
  normalizeDocSlugPath(route.params.docPath),
);
const isSemanticEdit = computed(
  () => !!semanticCollectionSlug.value && semanticDocPath.value.length > 0,
);
const isNew = computed(() => !isSemanticEdit.value && routeId.value === "new");

function selectedValue(value: unknown) {
  if (value && typeof value === "object" && "value" in value) {
    const option = value as { value?: unknown };
    return String(option.value ?? "");
  }
  return String(value ?? "");
}

const mounted = ref(false);
onMounted(() => {
  mounted.value = true;
});

// ── collections list (for dropdown) ───────────────────────────────────────────
const { data: colsData } = await useAsyncData(
  "doc-editor-collections",
  () => call<{ items: CollectionView[] }>("/api/v1/collections"),
  { server: false, default: () => ({ items: [] as CollectionView[] }) },
);
const collections = computed(() => colsData.value?.items ?? []);
const colOptions = computed(() =>
  collections.value.map((c) => ({ label: c.title, value: c.id })),
);

// ── form ───────────────────────────────────────────────────────────────────────
const form = reactive({
  collectionId: "",
  parentId: ROOT,
  slug: "",
  title: "",
  content: "",
  excerpt: "",
  seoTitle: "",
  seoDescription: "",
  status: "draft",
  locale: "en",
  versionId: "",
  translationKey: "",
  sortOrder: 0,
});
const slugTouched = ref(false);

function clientSlug(value: string) {
  return value
    .trim()
    .toLowerCase()
    .replace(/[^\p{L}\p{N}]+/gu, "-")
    .replace(/^-+|-+$/g, "");
}

// TDZ-safe: useSeoMeta AFTER form is declared — getter reads form.title, and the
// reactive effect evaluates synchronously during setup; placing it above form's
// declaration hits the temporal dead zone (Cannot access 'form' before init).
useSeoMeta({
  title: () =>
    `${form.title || (isNew.value ? "新建文档" : "编辑文档")} · 控制台`,
});

// Route query: collection = collection slug, parent = parent doc uuid (new mode)
const queryCollection = route.query.collection as string | undefined;
const queryParent = route.query.parent as string | undefined;

// ── collection helpers ─────────────────────────────────────────────────────────
const collectionSlug = computed(
  () =>
    semanticCollectionSlug.value ||
    (collections.value.find((c) => c.id === form.collectionId)?.slug ?? ""),
);
const selectedCollection = computed(
  () =>
    collections.value.find(
      (c) => c.id === form.collectionId || c.slug === collectionSlug.value,
    ) ?? null,
);
const treeCollectionId = computed(
  () => form.collectionId || selectedCollection.value?.id || "",
);
const collectionTitle = computed(() => selectedCollection.value?.title ?? "");

// Reset parentId only on user-initiated collection change (not during init)
function onCollectionChange(value: unknown) {
  const newId = selectedValue(value);
  if (newId !== form.collectionId) {
    form.parentId = ROOT;
    form.versionId = "";
  }
  form.collectionId = newId;
}

// ── versions for selected collection ─────────────────────────────────────────
const { data: versionsData } = await useAsyncData(
  () => `doc-editor-versions-${treeCollectionId.value}`,
  () =>
    treeCollectionId.value
      ? call<CollectionVersionsResponse>(
          `/api/v1/collections/${treeCollectionId.value}/versions`,
        )
      : Promise.resolve({ items: [] as CollectionVersion[] }),
  {
    watch: [treeCollectionId],
    server: false,
    default: () => ({ items: [] as CollectionVersion[] }),
  },
);
const versions = computed(() => versionsData.value?.items ?? []);
const versionOptions = computed(() =>
  versions.value.map((v) => ({
    label: v.isDefault ? `${v.label}（默认）` : v.label,
    value: v.id,
    description: v.key,
  })),
);
const localeOptions = [
  { label: "English", value: "en" },
  { label: "简体中文", value: "zh-CN" },
];
const selectedVersion = computed(
  () =>
    versions.value.find((v) => v.id === form.versionId) ??
    versions.value.find((v) => v.isDefault) ??
    versions.value[0] ??
    null,
);
const selectedVersionKey = computed(() => selectedVersion.value?.key ?? "");

watch(
  versions,
  (items) => {
    if (!items.length || form.versionId) return;
    form.versionId = items.find((v) => v.isDefault)?.id ?? items[0]!.id;
  },
  { immediate: true },
);

// ── tree for parent dropdown ───────────────────────────────────────────────────
const { data: treeDocsData, refresh: refreshTree } = await useAsyncData(
  () =>
    `doc-editor-docs-${treeCollectionId.value}-${selectedVersionKey.value || "default"}-${form.locale || "en"}`,
  () =>
    treeCollectionId.value
      ? call<{ items: DocDetail[] }>("/api/v1/docs", {
          query: {
            collectionId: treeCollectionId.value,
            locale: form.locale || "en",
            ...(selectedVersionKey.value
              ? { version: selectedVersionKey.value }
              : {}),
          },
        })
      : Promise.resolve({ items: [] as DocDetail[] }),
  {
    watch: [treeCollectionId, selectedVersionKey, () => form.locale],
    server: false,
    default: () => ({ items: [] as DocDetail[] }),
  },
);
const treeData = computed<CollectionManageTree | null>(() =>
  selectedCollection.value
    ? {
        collection: selectedCollection.value,
        tree: buildDocTree(treeDocsData.value?.items ?? []),
      }
    : null,
);

const semanticDoc = computed(() =>
  isSemanticEdit.value
    ? findDocBySlugPath(treeData.value?.tree ?? [], semanticDocPath.value)
    : null,
);
const docId = computed(() =>
  isSemanticEdit.value ? (semanticDoc.value?.id ?? "") : routeId.value,
);

// ── load existing doc (edit mode) ──────────────────────────────────────────────
const {
  data: docData,
  pending,
  refresh,
} = await useAsyncData(
  () =>
    `doc-editor-${isSemanticEdit.value ? `${semanticCollectionSlug.value}/${semanticDocPath.value.join("/")}` : routeId.value}`,
  () =>
    isNew.value || !docId.value
      ? Promise.resolve(null)
      : call<DocDetailResponse>(`/api/v1/docs/${docId.value}`),
  { server: false, watch: [docId] },
);
const doc = computed(() => docData.value?.doc ?? null);

// Init form once data arrives
watch(
  [colsData, docData],
  () => {
    const cols = colsData.value?.items ?? [];
    if (isNew.value) {
      if (!form.collectionId && cols.length) {
        const col = queryCollection
          ? cols.find((c) => c.slug === queryCollection)
          : undefined;
        form.collectionId = col?.id ?? cols[0]!.id;
      }
      if (form.parentId === ROOT && queryParent) {
        form.parentId = queryParent;
      }
    } else if (doc.value) {
      form.collectionId = doc.value.collectionId;
      form.parentId = doc.value.parentId || ROOT;
      form.slug = doc.value.slug || "";
      slugTouched.value = true;
      form.title = doc.value.title;
      form.content = doc.value.content || "";
      form.excerpt = doc.value.excerpt || "";
      form.seoTitle = doc.value.seoTitle || "";
      form.seoDescription = doc.value.seoDescription || "";
      form.status = doc.value.status || "draft";
      form.locale = doc.value.locale || "en";
      form.versionId = doc.value.versionId || "";
      form.translationKey = doc.value.translationKey || "";
      form.sortOrder = doc.value.sortOrder ?? 0;
    }
  },
  { immediate: true },
);

// Flatten a doc tree into path-style select options; exclude editing node + descendants.
function flattenTree(
  nodes: DocDetail[],
  excludeId: string,
  prefix = "",
): { label: string; value: string }[] {
  const out: { label: string; value: string }[] = [];
  for (const node of nodes) {
    if (node.id === excludeId) continue; // also skips descendants
    const label = prefix ? `${prefix} / ${node.title}` : node.title;
    out.push({ label, value: node.id });
    if (node.children?.length)
      out.push(...flattenTree(node.children, excludeId, label));
  }
  return out;
}

const parentOptions = computed(() => [
  { label: "（无父文档）", value: ROOT },
  ...flattenTree(treeData.value?.tree ?? [], isNew.value ? "" : docId.value),
]);

const currentSlugPath = computed(() => {
  const parentPath =
    form.parentId !== ROOT
      ? (findDocSlugPathById(treeData.value?.tree ?? [], form.parentId) ?? [])
      : [];
  const editableSlug = form.slug.trim();
  if (editableSlug) return [...parentPath, editableSlug];
  if (isNew.value) return parentPath;
  if (isSemanticEdit.value) return semanticDocPath.value;
  return docId.value
    ? (findDocSlugPathById(treeData.value?.tree ?? [], docId.value) ??
        (doc.value?.slug ? [doc.value.slug] : []))
    : [];
});
const publicDocUrl = computed(() =>
  collectionSlug.value && currentSlugPath.value.length
    ? [
        `/${collectionSlug.value}/${currentSlugPath.value.map(encodeURIComponent).join("/")}`,
        new URLSearchParams({
          ...(form.locale && form.locale !== "en"
            ? { locale: form.locale }
            : {}),
          ...(selectedVersion.value && !selectedVersion.value.isDefault
            ? { version: selectedVersion.value.key }
            : {}),
        }).toString(),
      ]
        .filter(Boolean)
        .join("?")
    : "",
);

// ── status ─────────────────────────────────────────────────────────────────────
const statusMeta: Record<
  string,
  { label: string; color: "neutral" | "success" | "warning"; icon: string }
> = {
  draft: { label: "草稿", color: "neutral", icon: "i-tabler-pencil" },
  published: {
    label: "已发布",
    color: "success",
    icon: "i-tabler-circle-check",
  },
  archived: { label: "已归档", color: "warning", icon: "i-tabler-archive" },
};
const sm = computed(
  () => statusMeta[form.status || "draft"] ?? statusMeta.draft!,
);

const contentText = computed(() =>
  form.content
    .replace(/<[^>]*>/g, " ")
    .replace(/\s+/g, " ")
    .trim(),
);
// ── quick publish / unpublish (edit mode) ──────────────────────────────────────
const busy = ref("");
async function setStatus(status: "draft" | "published" | "archived") {
  if (!docId.value) return;
  busy.value = status;
  try {
    if (status === "published") {
      await call(`/api/v1/docs/${docId.value}/publish`, { method: "POST" });
    } else if (status === "archived") {
      await call(`/api/v1/docs/${docId.value}/archive`, { method: "POST" });
    } else {
      await call(`/api/v1/docs/${docId.value}`, {
        method: "PATCH",
        body: { status },
      });
    }
    await refresh();
    form.status = doc.value?.status || status;
  } catch (e: any) {
    toast.add({
      title: "操作失败",
      description: e?.data?.message || "请重试",
      color: "error",
    });
  } finally {
    busy.value = "";
  }
}

// ── save ───────────────────────────────────────────────────────────────────────
const {
  status: saveStatus,
  pending: markSaving,
  success: markSaved,
  reset: resetSave,
} = useActionFeedback();
function showValidationError(description: string) {
  toast.add({
    title: "操作失败",
    description,
    color: "error",
    icon: "i-tabler-alert-circle",
  });
}
async function save() {
  if (!form.title.trim()) {
    showValidationError("请填写标题");
    return;
  }
  if (!form.collectionId) {
    showValidationError("请选择文档集");
    return;
  }
  if (!isNew.value && !form.slug.trim()) {
    showValidationError("请填写 URL slug");
    return;
  }
  markSaving();
  try {
    if (isNew.value) {
      const res = await call<{ doc: DocDetail }>("/api/v1/docs", {
        method: "POST",
        body: {
          collectionId: form.collectionId,
          ...(form.parentId !== ROOT ? { parentId: form.parentId } : {}),
          ...(form.slug.trim() ? { slug: form.slug.trim() } : {}),
          title: form.title,
          content: form.content,
          seoTitle: form.seoTitle,
          seoDescription: form.seoDescription,
          locale: form.locale || "en",
          versionId: form.versionId,
          translationKey: form.translationKey,
          sortOrder: Number(form.sortOrder || 0),
        },
      });
      if (form.excerpt.trim() && res.doc?.id) {
        await call(`/api/v1/docs/${res.doc.id}`, {
          method: "PATCH",
          body: { excerpt: form.excerpt },
        });
      }
      markSaved();
      const parentSlugPath =
        form.parentId !== ROOT
          ? (findDocSlugPathById(treeData.value?.tree ?? [], form.parentId) ??
            [])
          : [];
      await navigateTo(
        docManageRoute(collectionSlug.value, [
          ...parentSlugPath,
          res.doc!.slug,
        ]),
      );
    } else {
      if (!docId.value) {
        toast.add({
          title: "找不到文档",
          color: "error",
          icon: "i-tabler-alert-circle",
        });
        return;
      }
      const res = await call<{ doc: DocDetail }>(
        `/api/v1/docs/${docId.value}`,
        {
          method: "PATCH",
          body: {
            title: form.title,
            slug: form.slug.trim(),
            content: form.content,
            excerpt: form.excerpt,
            seoTitle: form.seoTitle,
            seoDescription: form.seoDescription,
            status: form.status,
            locale: form.locale || "en",
            versionId: form.versionId,
            translationKey: form.translationKey,
            sortOrder: Number(form.sortOrder || 0),
            parentId: form.parentId === ROOT ? "" : form.parentId,
          },
        },
      );
      markSaved();
      await refreshTree();
      await refresh();
      const savedSlug = res.doc?.slug || form.slug.trim();
      const parentSlugPath =
        form.parentId !== ROOT
          ? (findDocSlugPathById(treeData.value?.tree ?? [], form.parentId) ??
            [])
          : [];
      if (collectionSlug.value && savedSlug) {
        await navigateTo(
          docManageRoute(collectionSlug.value, [...parentSlugPath, savedSlug]),
          { replace: true },
        );
      }
    }
  } catch (e: any) {
    resetSave();
    toast.add({
      title: isNew.value ? "创建失败" : "保存失败",
      description: e?.data?.message || "请重试",
      color: "error",
    });
  }
}

// ── settings drawer ────────────────────────────────────────────────────────────
const settingsOpen = ref(false);
const {
  status: copyStatus,
  success: markCopied,
  error: markCopyFailed,
} = useActionFeedback();

async function copyText(value: string) {
  if (!value) return;
  try {
    await navigator.clipboard.writeText(value);
    markCopied();
    // feedback-contract: clipboard writes have no persistent visible result outside the current control
    toast.add({
      title: "公开链接已复制",
      color: "success",
      icon: "i-tabler-copy-check",
    });
  } catch {
    markCopyFailed();
    toast.add({
      title: "复制失败",
      color: "error",
      icon: "i-tabler-alert-circle",
    });
  }
}

function previewDoc() {
  if (publicDocUrl.value) window.open(publicDocUrl.value, "_blank");
}

// ── keyboard shortcuts ─────────────────────────────────────────────────────────
defineShortcuts({
  meta_s: { usingInput: true, handler: () => save() },
  ctrl_s: { usingInput: true, handler: () => save() },
});

// ── title auto-grow ────────────────────────────────────────────────────────────
const titleEl = ref<HTMLTextAreaElement>();
function autoGrowTitle() {
  const el = titleEl.value;
  if (!el) return;
  el.style.height = "0px";
  el.style.height = `${el.scrollHeight}px`;
}
watch(
  () => form.title,
  () => nextTick(autoGrowTitle),
);
watch(
  () => form.title,
  (title) => {
    if (isNew.value && !slugTouched.value) form.slug = clientSlug(title);
  },
);
onMounted(() => nextTick(autoGrowTitle));
</script>

<template>
  <div class="yueli-admin-canvas min-h-full min-w-0" data-docs-editor>
    <div
      class="sticky top-0 z-30 flex min-h-16 items-center justify-between gap-2 border-b border-default bg-default px-3 py-1.5 sm:gap-4 sm:px-4 sm:py-2 lg:px-8"
      data-docs-editor-commandbar
    >
      <div class="flex min-w-0 items-center gap-2">
        <UDashboardSidebarToggle class="size-11 sm:size-8 lg:hidden" />
        <UTooltip text="返回文档列表">
          <UButton
            to="/manage/docs"
            icon="i-tabler-arrow-left"
            color="neutral"
            variant="ghost"
            square
            class="size-11 sm:size-8"
            aria-label="返回文档列表"
          />
        </UTooltip>
        <span
          class="hidden max-w-[min(32vw,28rem)] truncate text-sm font-semibold text-toned md:block"
        >
          {{ isNew ? "新建文档" : "文档编辑" }}
        </span>
        <template v-if="!isNew && doc">
          <span class="hidden h-5 w-px bg-accented sm:block" />
          <UBadge
            class="hidden sm:inline-flex"
            :color="sm.color"
            :icon="sm.icon"
            :label="sm.label"
            variant="subtle"
          />
        </template>
      </div>

      <div class="flex shrink-0 items-center gap-1.5">
        <UTooltip v-if="!isNew && doc" text="预览公开页">
          <UButton
            icon="i-tabler-eye"
            color="neutral"
            variant="ghost"
            square
            class="size-11 sm:size-8"
            aria-label="预览公开页"
            :disabled="!publicDocUrl"
            @click="previewDoc"
          />
        </UTooltip>
        <UTooltip text="文档设置">
          <UButton
            icon="i-tabler-adjustments-horizontal"
            color="neutral"
            variant="ghost"
            square
            class="size-11 sm:size-8"
            aria-label="文档设置"
            @click="void (settingsOpen = true)"
          />
        </UTooltip>
        <UButton
          v-if="!isNew && doc && form.status !== 'published'"
          label="发布"
          icon="i-tabler-rocket"
          color="primary"
          variant="soft"
          class="min-h-11 sm:min-h-8"
          :loading="busy === 'published'"
          @click="setStatus('published')"
        />
        <ActionFeedbackButton
          :status="saveStatus"
          :idle-label="isNew ? '创建' : '保存'"
          :pending-label="isNew ? '创建中' : '保存中'"
          :success-label="isNew ? '已创建' : '已保存'"
          :idle-icon="isNew ? 'i-tabler-check' : undefined"
          class="min-h-11 sm:min-h-8"
          @click="save"
        />
      </div>
    </div>

    <div
      v-if="!mounted || (!isNew && pending && !doc)"
      class="mx-auto max-w-4xl space-y-5 px-4 py-8 sm:px-6 lg:px-8"
      aria-label="正在加载文档编辑器"
    >
      <USkeleton class="h-16 w-4/5 rounded-xl" />
      <USkeleton class="h-[34rem] rounded-xl" />
    </div>

    <main
      v-else
      class="px-4 pb-12 pt-6 sm:px-6 sm:pb-16 sm:pt-8 lg:px-8 lg:pt-10"
    >
      <section
        class="mx-auto w-full max-w-6xl rounded-xl bg-default p-3 shadow-sm sm:rounded-2xl sm:p-4 lg:p-6"
        data-docs-editor-document
        aria-label="文档正文编辑"
      >
        <header class="mb-5 px-1">
          <textarea
            ref="titleEl"
            v-model="form.title"
            rows="1"
            placeholder="未命名文档"
            class="block w-full resize-none overflow-hidden border-0 bg-transparent font-display text-[1.75rem] font-bold leading-[1.12] tracking-[-0.04em] text-highlighted outline-none placeholder:text-dimmed sm:text-[clamp(2rem,3vw,2.25rem)]"
            aria-label="文档标题"
            @input="autoGrowTitle"
          />
          <div
            class="mt-3 flex min-h-9 items-center gap-1.5 rounded-xl border border-default bg-default/80 px-2.5 py-1.5 text-xs"
          >
            <UIcon name="i-tabler-link" class="size-4 shrink-0 text-primary" />
            <span class="shrink-0 text-dimmed">
              /{{ collectionSlug || "collection" }}/<template
                v-if="currentSlugPath.length > 1"
                >{{ currentSlugPath.slice(0, -1).join("/") }}/</template
              >
            </span>
            <input
              v-model="form.slug"
              placeholder="url-slug"
              class="min-w-0 flex-1 bg-transparent text-toned outline-none transition placeholder:text-dimmed focus:text-highlighted"
              aria-label="文档永久链接"
              @input="slugTouched = true"
            />
            <span class="hidden shrink-0 text-dimmed sm:inline">
              {{ contentText.length }} 字符
            </span>
          </div>
        </header>

        <ContentEditor
          v-model="form.content"
          class="[&>div>.rounded-xl]:border-default [&>div>.rounded-xl]:bg-muted [&_[data-slot=content]]:mx-auto [&_[data-slot=content]]:min-h-[28rem] [&_[data-slot=content]]:w-full [&_[data-slot=content]]:px-[1.125rem] [&_[data-slot=content]]:py-6 sm:[&_[data-slot=content]]:min-h-[max(40rem,calc(100svh-19rem))] sm:[&_[data-slot=content]]:px-[clamp(2rem,4vw,3rem)] sm:[&_[data-slot=content]]:py-9"
          draft-key-prefix="docs:doc"
          :draft-entity-id="isNew ? 'new' : docId"
          :draft-mode="isNew ? 'create' : 'edit'"
          :has-initial-content="!isNew && !!doc?.content"
        />
      </section>
    </main>

    <!-- settings slideover: collection · parent · status -->
    <USlideover
      v-model:open="settingsOpen"
      title="文档设置"
      description="归属、发布与搜索设置"
      :ui="{
        content: 'w-full max-w-2xl bg-default',
        header: 'bg-default',
        body: 'bg-muted p-4 sm:p-5',
        footer: 'bg-default',
      }"
    >
      <template #body>
        <div class="space-y-5">
          <section
            class="space-y-4 rounded-lg bg-default p-4 ring-1 ring-default"
          >
            <h2 class="text-sm font-semibold text-highlighted">摘要与链接</h2>
            <UFormField label="摘要">
              <UTextarea
                v-model="form.excerpt"
                autoresize
                :rows="3"
                placeholder="用于搜索结果、集合页和分享预览"
                class="w-full"
              />
            </UFormField>
            <div v-if="!isNew" class="space-y-2">
              <div class="flex items-center justify-between gap-3">
                <p class="text-xs font-medium text-muted">公开链接</p>
                <UTooltip
                  :text="copyStatus === 'success' ? '已复制' : '复制公开链接'"
                >
                  <UButton
                    :icon="
                      copyStatus === 'success'
                        ? 'i-tabler-check'
                        : 'i-tabler-copy'
                    "
                    color="neutral"
                    variant="ghost"
                    size="xs"
                    square
                    :aria-label="
                      copyStatus === 'success'
                        ? '公开链接已复制'
                        : '复制公开链接'
                    "
                    :disabled="!publicDocUrl"
                    @click="copyText(publicDocUrl)"
                  />
                </UTooltip>
              </div>
              <p
                class="truncate rounded-md bg-elevated px-3 py-2 font-mono text-xs text-default"
              >
                {{ publicDocUrl || "发布后生成" }}
              </p>
            </div>
          </section>

          <section
            class="space-y-5 rounded-lg bg-default p-4 ring-1 ring-default"
          >
            <h2 class="text-sm font-semibold text-highlighted">归属与路径</h2>
            <UFormField label="文档集" required>
              <USelectMenu
                v-if="isNew"
                :model-value="form.collectionId"
                :items="colOptions"
                value-key="value"
                placeholder="选择文档集"
                :search-input="{ placeholder: '搜索文档集…' }"
                class="w-full"
                @update:model-value="onCollectionChange"
              />
              <p
                v-else
                class="rounded-lg bg-elevated px-3 py-2 text-sm text-default"
              >
                {{ collectionTitle || form.collectionId }}
              </p>
            </UFormField>

            <!-- parent doc (optional) -->
            <UFormField label="父文档" help="留空则为顶级文档">
              <USelectMenu
                v-model="form.parentId"
                :items="parentOptions"
                value-key="value"
                placeholder="选择父文档"
                :search-input="{ placeholder: '搜索标题或路径…' }"
                class="w-full"
              />
            </UFormField>

            <UFormField label="版本">
              <USelectMenu
                :model-value="form.versionId"
                :items="versionOptions"
                value-key="value"
                placeholder="默认版本"
                :search-input="{ placeholder: '搜索版本…' }"
                class="w-full"
                @update:model-value="form.versionId = selectedValue($event)"
              />
            </UFormField>

            <div class="grid gap-4 sm:grid-cols-2">
              <UFormField label="语言">
                <USelectMenu
                  :model-value="form.locale"
                  :items="localeOptions"
                  value-key="value"
                  class="w-full"
                  @update:model-value="
                    form.locale = selectedValue($event) || 'en'
                  "
                />
              </UFormField>

              <UFormField label="排序" help="同级文档按数值从小到大排列">
                <UInput
                  v-model.number="form.sortOrder"
                  type="number"
                  min="0"
                  class="w-full"
                />
              </UFormField>
            </div>

            <UFormField
              label="翻译关联键"
              help="不同语言的同一篇文档使用同一个 key；留空则系统生成"
            >
              <UInput
                v-model="form.translationKey"
                placeholder="install-guide"
                class="w-full"
              />
            </UFormField>

            <UFormField
              label="URL slug"
              required
              help="同一父文档下唯一；保存后公开链接同步更新"
            >
              <UInput
                v-model="form.slug"
                icon="i-tabler-link"
                placeholder="overview"
                class="w-full"
                @input="slugTouched = true"
              />
            </UFormField>
          </section>

          <!-- status (edit mode only; new docs start as draft) -->
          <section
            v-if="!isNew"
            class="space-y-3 rounded-lg bg-default p-4 ring-1 ring-default"
          >
            <div class="flex items-center justify-between gap-3">
              <div>
                <p class="text-sm font-semibold text-highlighted">状态操作</p>
                <p class="text-xs text-muted">低频生命周期动作</p>
              </div>
              <UBadge
                :color="sm.color"
                :icon="sm.icon"
                :label="sm.label"
                variant="subtle"
              />
            </div>
            <div class="grid gap-2">
              <UButton
                v-if="form.status !== 'published'"
                label="发布"
                icon="i-tabler-rocket"
                color="primary"
                variant="soft"
                block
                :loading="busy === 'published'"
                @click="setStatus('published')"
              />
              <UButton
                v-if="form.status !== 'draft'"
                label="转回草稿"
                icon="i-tabler-pencil"
                color="neutral"
                variant="outline"
                block
                :loading="busy === 'draft'"
                @click="setStatus('draft')"
              />
              <UButton
                v-if="form.status !== 'archived'"
                label="归档"
                icon="i-tabler-archive"
                color="warning"
                variant="soft"
                block
                :loading="busy === 'archived'"
                @click="setStatus('archived')"
              />
            </div>
          </section>

          <section
            class="space-y-4 rounded-lg bg-default p-4 ring-1 ring-default"
          >
            <h2 class="text-sm font-semibold text-highlighted">搜索优化</h2>
            <UFormField label="SEO 标题">
              <UInput
                v-model="form.seoTitle"
                placeholder="默认使用文档标题"
                class="w-full"
              />
            </UFormField>

            <UFormField label="SEO 描述">
              <UTextarea
                v-model="form.seoDescription"
                :rows="3"
                placeholder="用于搜索引擎和分享摘要"
                class="w-full"
              />
            </UFormField>
          </section>
        </div>
      </template>

      <template #footer>
        <div class="flex w-full justify-end gap-2">
          <UButton
            label="完成"
            color="neutral"
            variant="outline"
            @click="void (settingsOpen = false)"
          />
          <ActionFeedbackButton
            :status="saveStatus"
            idle-label="保存"
            pending-label="保存中"
            success-label="已保存"
            @click="save"
          />
        </div>
      </template>
    </USlideover>
  </div>
</template>
