<script setup lang="ts">
import { EditorCommandBar } from "@yueli/ui/admin";
import { createDocsNotifier } from "~/utils/feedback";
import { EditorInspector } from "@yueli/ui/admin";
import { AdminIconPicker } from "@yueli/ui/admin";
import { useActionFeedback } from "@yueli/ui/feedback";
import { ActionFeedbackButton } from "@yueli/ui/feedback/pattern";
import { AssetImageProcessor } from "@yueli/asset-nuxt/components";
import type {
  CollectionLocale,
  CollectionLocalesResponse,
  CollectionVersion,
  CollectionVersionsResponse,
  CollectionView,
  ManageDocListItem,
  ManageDocsResponse,
  DocDetail,
  DocDetailResponse,
} from "~/types";
import {
  docManageRoute,
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

function createDraftInstanceId() {
  if (globalThis.crypto?.randomUUID) return globalThis.crypto.randomUUID();
  const bytes = new Uint8Array(12);
  globalThis.crypto?.getRandomValues(bytes);
  return Array.from(bytes, (value) => value.toString(16).padStart(2, "0")).join(
    "",
  );
}

const draftInstanceId = ref(
  isNew.value ? String(route.query.draft ?? "").trim() : "",
);
if (import.meta.client && isNew.value && !draftInstanceId.value) {
  draftInstanceId.value = createDraftInstanceId();
  const legacyKey = "docs:doc:new";
  const scopedKey = `docs:doc:new:${draftInstanceId.value}`;
  const legacyDraft = localStorage.getItem(legacyKey);
  if (legacyDraft && !localStorage.getItem(scopedKey)) {
    localStorage.setItem(scopedKey, legacyDraft);
    localStorage.removeItem(legacyKey);
  }
}

function selectedValue(value: unknown) {
  if (value && typeof value === "object" && "value" in value) {
    const option = value as { value?: unknown };
    return String(option.value ?? "");
  }
  return String(value ?? "");
}

const mounted = ref(false);
const immersiveCollaboration = ref(false);
onMounted(() => {
  mounted.value = true;
  if (isNew.value && draftInstanceId.value && !route.query.draft) {
    void navigateTo(
      {
        path: route.path,
        query: { ...route.query, draft: draftInstanceId.value },
      },
      { replace: true },
    );
  }
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
  locale: "",
  versionId: "",
  translationKey: "",
  badgeText: "",
  badgeIcon: "",
  sortOrder: 0,
});
const badgeMode = ref<"none" | "text" | "icon">("none");
const badgeModeItems = [
  { label: "不显示", value: "none" },
  { label: "文字", value: "text" },
  { label: "图标", value: "icon" },
];
watch(badgeMode, (mode) => {
  if (mode !== "text") form.badgeText = "";
  if (mode !== "icon") form.badgeIcon = "";
  if (mode === "icon" && !form.badgeIcon) form.badgeIcon = "i-tabler-sparkles";
});
const slugTouched = ref(false);

type ImageProcessingRequest = {
  file: File;
  resolve: (file: File) => void;
  reject: (error: Error) => void;
};
const imageProcessingRequest = shallowRef<ImageProcessingRequest | null>(null);
const { uploadDocumentImage } = useDocsUpload();
const editorComp = ref<{ markSaved: () => void } | null>(null);

function processImage(file: File) {
  imageProcessingRequest.value?.reject(new Error("已取消上一项图片处理"));
  return new Promise<File>((resolve, reject) => {
    imageProcessingRequest.value = { file, resolve, reject };
  });
}

function finishImageProcessing(file: File) {
  const pending = imageProcessingRequest.value;
  imageProcessingRequest.value = null;
  pending?.resolve(file);
}

function cancelImageProcessing() {
  const pending = imageProcessingRequest.value;
  imageProcessingRequest.value = null;
  pending?.reject(new Error("已取消图片处理"));
}

async function uploadInlineImage(file: File) {
  const processed = await processImage(file);
  return uploadDocumentImage(processed, {
    ...(isNew.value
      ? { collectionId: form.collectionId }
      : { documentId: docId.value }),
  });
}

onBeforeUnmount(cancelImageProcessing);

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
    form.locale = "";
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
const { data: localesData } = await useAsyncData(
  () => `doc-editor-locales-${treeCollectionId.value}`,
  () =>
    treeCollectionId.value
      ? call<CollectionLocalesResponse>(
          `/api/v1/manage/collections/${treeCollectionId.value}/locales`,
        )
      : Promise.resolve({ items: [] as CollectionLocale[] }),
  {
    watch: [treeCollectionId],
    server: false,
    default: () => ({ items: [] as CollectionLocale[] }),
  },
);
const locales = computed(() => localesData.value?.items ?? []);
const defaultLocale = computed(
  () =>
    locales.value.find((item) => item.isDefault)?.locale ??
    locales.value[0]?.locale ??
    "en",
);
const localeOptions = computed(() =>
  locales.value
    .filter((item) => item.enabled || item.locale === form.locale)
    .map((item) => ({
      label: item.isDefault ? `${item.label}（默认）` : item.label,
      value: item.locale,
      description: item.locale,
    })),
);
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

watch(
  locales,
  (items) => {
    if (!items.length) return;
    if (items.some((item) => item.locale === form.locale)) return;
    form.locale =
      items.find((item) => item.isDefault)?.locale ?? items[0]!.locale;
  },
  { immediate: true },
);

// Resolve only the selected document; never fetch the collection tree.
const { data: currentRowData, error: rowError, status: rowStatus, refresh: refreshCurrentRow } = await useAsyncData(
  () => `doc-editor-location-${route.path}-${String(route.query.locale || '')}-${String(route.query.version || '')}`,
  () => isNew.value ? Promise.resolve(null) : call<ManageDocsResponse>('/api/v1/manage/docs', {
    query: isSemanticEdit.value ? {
      collectionId: treeCollectionId.value, path: semanticDocPath.value.join('/'),
      locale: String(route.query.locale || form.locale || defaultLocale.value),
      version: String(route.query.version || selectedVersionKey.value), size: 1, page: 1,
    } : { id: routeId.value, size: 1, page: 1 },
  }),
  { server: false, immediate: false },
);
const currentRow = computed(() => currentRowData.value?.items[0]);
const docId = computed(() => isNew.value ? '' : isSemanticEdit.value ? (currentRow.value?.id || '') : routeId.value);
watch(
  () => isNew.value ? '' : isSemanticEdit.value
    ? (treeCollectionId.value && locales.value.length && versions.value.length ? `${route.path}:${String(route.query.locale || defaultLocale.value)}:${String(route.query.version || selectedVersionKey.value)}` : '')
    : routeId.value,
  (key) => { if (key) void refreshCurrentRow(); }, { immediate: true },
);
// ── load existing doc (edit mode) ──────────────────────────────────────────────
const {
  data: docData,
  error: loadError,
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
const editorReady = computed(() => isNew.value || !!doc.value);

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
      form.locale = doc.value.locale || defaultLocale.value;
      form.versionId = doc.value.versionId || "";
      form.translationKey = doc.value.translationKey || "";
      form.badgeText = doc.value.badgeText || "";
      form.badgeIcon = doc.value.badgeIcon || "";
      badgeMode.value = form.badgeIcon
        ? "icon"
        : form.badgeText
          ? "text"
          : "none";
      form.sortOrder = doc.value.sortOrder ?? 0;
    }
  },
  { immediate: true },
);

const parentOpen = ref(false);
const parentSearch = ref('');
const parentQuery = ref('');
const parentLoading = ref(false);
const parentFailure = ref('');
const parentRows = ref<ManageDocListItem[]>([]);
const selectedParent = shallowRef<{ id: string; title: string; slugPath: string } | null>(null);
watch(() => isNew.value ? form.parentId : ROOT, async (id) => {
  if (id === ROOT || selectedParent.value?.id === id) return;
  const result = await call<ManageDocsResponse>('/api/v1/manage/docs', { query: { id, size: 1, page: 1 } }).catch(() => null);
  if (form.parentId === id && result?.items[0]) selectedParent.value = result.items[0];
}, { immediate: true });
let parentRequest = 0;
let searchTimer: ReturnType<typeof setTimeout> | undefined;
watch(parentSearch, (value) => {
  clearTimeout(searchTimer);
  searchTimer = setTimeout(() => { parentQuery.value = value.trim(); }, 250);
});
onBeforeUnmount(() => { clearTimeout(searchTimer); parentRequest++; });
watch(currentRow, (row) => {
  if (!row) return;
  selectedParent.value = row.parentId ? { id: row.parentId, title: row.parentTitle, slugPath: row.slugPath.split('/').slice(0, -1).join('/') } : null;
}, { immediate: true });
async function loadParents() {
  const request = ++parentRequest;
  if (!parentOpen.value || !treeCollectionId.value || !selectedVersionKey.value || !form.locale) { parentLoading.value = false; return; }
  parentLoading.value = true; parentFailure.value = ''; parentRows.value = [];
  try {
    const result = await call<ManageDocsResponse>('/api/v1/manage/docs', { query: {
      collectionId: treeCollectionId.value, locale: form.locale, version: selectedVersionKey.value,
      q: parentQuery.value, excludeId: docId.value || undefined, size: 20, page: 1, sort: 'title', direction: 'asc',
    } });
    if (request === parentRequest) parentRows.value = result.items;
  } catch (error) {
    if (request === parentRequest) parentFailure.value = docsFailureMessage(error, '父文档加载失败，请重试');
  } finally { if (request === parentRequest) parentLoading.value = false; }
}
watch([parentOpen, parentQuery, treeCollectionId, selectedVersionKey, () => form.locale], loadParents);
const parentOptions = computed(() => [
  { label: '（无父文档）', value: ROOT },
  ...parentRows.value.map(row => ({ label: row.title, description: row.slugPath, value: row.id })),
]);
const parentLabel = computed(() => form.parentId === ROOT ? '（无父文档）' : selectedParent.value?.title || '当前父文档');
function selectParent(value: unknown) {
  const id = selectedValue(value); form.parentId = id;
  if (id === ROOT) selectedParent.value = null;
  else { const row = parentRows.value.find(row => row.id === id); if (row) selectedParent.value = row; }
}
const currentSlugPath = computed(() => {
  const parent = form.parentId !== ROOT ? selectedParent.value?.slugPath.split('/').filter(Boolean) || [] : [];
  return form.slug.trim() ? [...parent, form.slug.trim()] : parent;
});
async function savedDocRoute(id: string) {
  const result = await call<ManageDocsResponse>('/api/v1/manage/docs', { query: { id, size: 1, page: 1 } });
  const row = result.items[0];
  if (!row) throw new Error('Saved document location is unavailable');
  return { path: docManageRoute(row.collectionSlug, row.slugPath.split('/')), query: { locale: row.locale, version: row.versionKey } };
}

const publicDocUrl = computed(() =>
  collectionSlug.value && currentSlugPath.value.length
    ? [
        `/${collectionSlug.value}/${currentSlugPath.value.map(encodeURIComponent).join("/")}`,
        new URLSearchParams({
          ...(form.locale && form.locale !== defaultLocale.value
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

// ── quick publish / unpublish (edit mode) ──────────────────────────────────────
const busy = ref("");
async function setStatus(status: "draft" | "published" | "archived") {
  if ((!docId.value && !isNew.value) || busy.value) return;
  busy.value = status;
  try {
    if (status === "published") {
      await save(true);
      return;
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
      description: docsFailureMessage(e, "请重试"),
      color: "error",
    });
  } finally {
    busy.value = "";
  }
}
const lifecycleMenuItems = computed(() => {
  const items = [];
  if (form.status === "archived") {
    items.push({
      label: "发布",
      icon: "i-tabler-rocket",
      disabled: !editorReady.value || Boolean(busy.value) || saveStatus.value === "pending",
      onSelect: () => void setStatus("published"),
    });
  }
  if (form.status !== "draft") {
    items.push({
      label: "下架",
      icon: "i-tabler-pencil",
      disabled: !editorReady.value || Boolean(busy.value) || saveStatus.value === "pending",
      onSelect: () => void setStatus("draft"),
    });
  }
  if (!isNew.value && form.status !== "archived") {
    items.push({
      label: "归档",
      icon: "i-tabler-archive",
      disabled: !editorReady.value || Boolean(busy.value) || saveStatus.value === "pending",
      onSelect: () => void setStatus("archived"),
    });
  }
  return [items];
});

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
async function save(publish = false) {
  if (!editorReady.value || saveStatus.value === "pending") return;
  if (!form.title.trim()) {
    showValidationError("请填写标题");
    return;
  }
  if (!form.collectionId) {
    settingsSection.value = "content";
    settingsOpen.value = true;
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
          locale: form.locale || defaultLocale.value,
          versionId: form.versionId,
          translationKey: form.translationKey,
          badgeText: form.badgeText,
          badgeIcon: form.badgeIcon,
          sortOrder: Number(form.sortOrder || 0),
        },
      });
      if (form.excerpt.trim() && res.doc?.id) {
        await call(`/api/v1/docs/${res.doc.id}`, {
          method: "PATCH",
          body: { excerpt: form.excerpt },
        });
      }
      if (publish) {
        await call(`/api/v1/docs/${res.doc.id}/publish`, { method: "POST" });
        form.status = "published";
      }
      markSaved();
      editorComp.value?.markSaved();
      await navigateTo(await savedDocRoute(res.doc.id));
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
            locale: form.locale || defaultLocale.value,
            versionId: form.versionId,
            translationKey: form.translationKey,
            badgeText: form.badgeText,
            badgeIcon: form.badgeIcon,
            sortOrder: Number(form.sortOrder || 0),
            parentId: form.parentId === ROOT ? "" : form.parentId,
          },
        },
      );
      if (publish) {
        await call(`/api/v1/docs/${res.doc.id}/publish`, { method: "POST" });
        form.status = "published";
      }
      markSaved();
      editorComp.value?.markSaved();
      await navigateTo(await savedDocRoute(res.doc.id), { replace: true });
      await refresh();
    }
  } catch (e: any) {
    resetSave();
    toast.add({
      title: isNew.value ? "创建失败" : "保存失败",
      description: docsFailureMessage(e, "请重试"),
      color: "error",
    });
  }
}

// ── settings inspector ─────────────────────────────────────────────────────────
type SettingsSection = "content" | "organization" | "seo";
const settingsOpen = ref(false);
const settingsSection = ref<SettingsSection>("content");
const settingsTabs = [
  { label: "内容", value: "content", icon: "i-tabler-stack-2" },
  { label: "组织", value: "organization", icon: "i-tabler-arrows-sort" },
  { label: "SEO", value: "seo", icon: "i-tabler-search" },
];
function toggleSettings() {
  settingsOpen.value = !settingsOpen.value;
}
const {
  status: copyStatus,
  success: markCopied,
  error: markCopyFailed,
} = useActionFeedback();

async function copyText(value: string) {
  if (!value) return;
  try {
    await writeClipboardText(value);
    markCopied();
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

function toggleImmersiveCollaboration() {
  immersiveCollaboration.value = !immersiveCollaboration.value;
}

// ── keyboard shortcuts ─────────────────────────────────────────────────────────
defineShortcuts({
  meta_s: { usingInput: true, handler: () => save() },
  ctrl_s: { usingInput: true, handler: () => save() },
  "meta_,": {
    usingInput: true,
    handler: () => {
      settingsOpen.value = !settingsOpen.value;
    },
  },
  "ctrl_,": {
    usingInput: true,
    handler: () => {
      settingsOpen.value = !settingsOpen.value;
    },
  },
});

watch(
  () => form.title,
  (title) => {
    if (isNew.value && !slugTouched.value) form.slug = clientSlug(title);
  },
);
</script>

<template>
  <div
    class="yueli-admin-canvas min-w-0"
    :class="immersiveCollaboration ? 'fixed inset-0 z-50 h-svh overflow-hidden bg-default' : 'min-h-full'"
    data-docs-editor
    :data-collaboration-mode="immersiveCollaboration ? 'immersive' : 'standard'"
  >
    <EditorCommandBar v-if="!immersiveCollaboration"
      v-model:immersive="immersiveCollaboration" v-model:settings-open="settingsOpen"
      back-to="/manage/docs" back-label="返回文档列表" settings-label="文档设置"
      data-docs-editor-commandbar>
      <template #title><input
          v-model="form.title"
          class="min-w-0 flex-1 border-0 bg-transparent text-sm font-semibold text-highlighted outline-none placeholder:text-dimmed md:text-base"
          placeholder="未命名文档"
          aria-label="文档标题"
        /></template>
      <template #preview><UTooltip v-if="!isNew && doc?.status === 'published'" text="查看公开页">
          <UButton
            icon="i-tabler-eye"
            color="neutral"
            variant="ghost"
            square
            class="size-8"
            aria-label="查看公开页"
            :disabled="!publicDocUrl"
            @click="previewDoc"
          />
        </UTooltip></template>
      <template #lifecycle><div data-docs-lifecycle-actions>
          <UFieldGroup v-if="form.status !== 'published'" size="sm">
            <UButton
              label="发布"
              icon="i-tabler-rocket"
              color="primary"
              variant="soft"
              class="h-8"
              :loading="busy === 'published'"
              :disabled="!editorReady || saveStatus === 'pending'"
              @click="setStatus('published')"
            />
            <UDropdownMenu v-if="lifecycleMenuItems[0]?.length" :items="lifecycleMenuItems">
              <UButton
                icon="i-tabler-chevron-down"
                color="primary"
                variant="soft"
                class="h-8"
                aria-label="更多发布操作"
                :disabled="!editorReady || Boolean(busy) || saveStatus === 'pending'"
              />
            </UDropdownMenu>
          </UFieldGroup>
          <UDropdownMenu v-else :items="lifecycleMenuItems">
            <UButton
              :label="sm.label"
              :icon="sm.icon"
              trailing-icon="i-tabler-chevron-down"
              color="neutral"
              variant="soft"
              class="h-8"
              :loading="Boolean(busy)"
              :disabled="!editorReady || saveStatus === 'pending'"
            />
          </UDropdownMenu>
        </div></template>
      <template #actions><ActionFeedbackButton
          :status="saveStatus"
          :idle-label="isNew ? '创建' : '保存'"
          :pending-label="isNew ? '创建中' : '保存中'"
          :success-label="isNew ? '已创建' : '已保存'"
          :idle-icon="isNew ? 'i-tabler-check' : undefined"
          class="h-8"
          :disabled="!editorReady || Boolean(busy)"
          @click="save()"
        /></template>
    </EditorCommandBar>

    <UAlert v-if="loadError || rowError" color="error" title="文档加载失败" description="请刷新页面重试。加载成功后才能保存或发布。" class="mx-auto my-6 max-w-4xl" />
    <UAlert v-else-if="isSemanticEdit && rowStatus === 'success' && !currentRow" color="error" title="找不到文档" description="请检查文档地址，或返回列表重新打开。" class="mx-auto my-6 max-w-4xl" />
    <div
      v-else-if="!mounted || (!isNew && !doc)"
      class="mx-auto max-w-4xl space-y-5 px-4 py-8 sm:px-6 lg:px-8"
      aria-label="正在加载文档编辑器"
    >
      <USkeleton class="h-16 w-4/5 rounded-xl" />
      <USkeleton class="h-[34rem] rounded-xl" />
    </div>

    <main
      v-else
      class="transition-[padding] duration-200 ease-out"
      :class="[
        immersiveCollaboration
          ? 'h-svh overflow-y-auto p-0'
          : 'px-4 pb-12 pt-6 sm:px-6 sm:pb-16 sm:pt-8 lg:px-8 lg:pt-10',
        settingsOpen && !immersiveCollaboration ? 'xl:pr-[27rem]' : '',
      ]"
      data-docs-editor-workspace
      :data-collaboration-mode="immersiveCollaboration ? 'immersive' : 'standard'"
    >
      <section
        class="mx-auto w-full bg-default"
        :class="
          immersiveCollaboration
            ? 'min-h-full max-w-none'
            : 'max-w-6xl rounded-xl p-3 shadow-sm sm:rounded-2xl sm:p-4 lg:p-6'
        "
        data-docs-editor-document
        aria-label="文档正文编辑"
      >
        <ContentEditor
          ref="editorComp"
          v-model="form.content"
          :class="[
            '[&>div>.rounded-xl]:border-default [&>div>.rounded-xl]:bg-muted [&_[data-slot=content]]:mx-auto [&_[data-slot=content]]:min-h-[28rem] [&_[data-slot=content]]:w-full [&_[data-slot=content]]:px-[1.125rem] [&_[data-slot=content]]:py-6 sm:[&_[data-slot=content]]:px-[clamp(2rem,4vw,3rem)] sm:[&_[data-slot=content]]:py-9',
            immersiveCollaboration
              ? '[&>div>.rounded-xl]:rounded-none [&>div>.rounded-xl]:border-0 [&_[data-slot=content]]:min-h-[calc(100svh-3.5rem)]'
              : 'sm:[&_[data-slot=content]]:min-h-[max(40rem,calc(100svh-19rem))]',
          ]"
          draft-key-prefix="docs:doc"
          :draft-entity-id="isNew ? draftInstanceId : docId"
          :draft-mode="isNew ? 'create' : 'edit'"
          :has-initial-content="!isNew && !!doc?.content"
          :image-uploader="uploadInlineImage"
          :style="{
            '--content-editor-toolbar-top': immersiveCollaboration
              ? '0px'
              : '4rem',
          }"
        >
          <template v-if="immersiveCollaboration" #toolbar-actions>
            <UTooltip text="退出沉浸式协作">
              <UButton
                icon="i-tabler-minimize"
                color="neutral"
                variant="ghost"
                size="sm"
                square
                aria-label="退出沉浸式协作"
                @click="toggleImmersiveCollaboration"
              />
            </UTooltip>
          </template>
        </ContentEditor>
      </section>
    </main>

    <EditorInspector v-model:open="settingsOpen" title="文档设置">
      <template #default="{ docked }">
        <div
          class="min-w-0"
          data-docs-editor-inspector
          :data-inspector-mode="docked ? 'docked' : 'overlay'"
        >
          <UTabs
            v-model="settingsSection"
            :items="settingsTabs"
            :content="false"
            value-key="value"
            variant="pill"
            color="neutral"
            class="w-full"
            :ui="{
              list: 'w-full rounded-xl bg-elevated/70 p-1',
              indicator: 'rounded-lg bg-default ring-1 ring-default shadow-xs',
              trigger:
                'min-h-9 flex-1 justify-center gap-2 rounded-lg data-[state=active]:text-highlighted',
              leadingIcon: 'size-4.5 shrink-0',
            }"
            data-docs-inspector-tabs
          />

          <section
            v-if="settingsSection === 'content'"
            class="mt-5 space-y-5"
            data-docs-inspector-content
          >
            <UFormField label="路径标识" required>
              <UFieldGroup class="w-full">
                <UInput
                  v-model="form.slug"
                  placeholder="url-slug"
                  class="min-w-0 w-full flex-1"
                  @input="slugTouched = true"
                />
                <UButton
                  v-if="!isNew"
                  :icon="copyStatus === 'success' ? 'i-tabler-check' : 'i-tabler-copy'"
                  color="neutral"
                  variant="outline"
                  :aria-label="copyStatus === 'success' ? '公开链接已复制' : '复制公开链接'"
                  :disabled="!publicDocUrl"
                  @click="copyText(publicDocUrl)"
                />
              </UFieldGroup>
            </UFormField>

            <UFormField label="摘要" help="用于搜索结果、集合页和分享预览">
              <UTextarea
                v-model="form.excerpt"
                autoresize
                :rows="3"
                :maxrows="6"
                class="w-full"
              />
            </UFormField>

            <UFormField label="文档集" required>
              <USelectMenu
                v-if="isNew"
                :model-value="form.collectionId"
                :items="colOptions"
                value-key="value"
                placeholder="选择文档集"
                aria-label="选择文档集"
                :search-input="{ placeholder: '搜索文档集…' }"
                class="w-full"
                @update:model-value="onCollectionChange"
              />
              <UInput
                v-else
                :model-value="collectionTitle || form.collectionId"
                disabled
                class="w-full"
              />
            </UFormField>

            <UFormField label="父文档">
              <USelectMenu
                v-model:open="parentOpen"
                v-model:search-term="parentSearch"
                :model-value="form.parentId"
                :items="parentOptions"
                :loading="parentLoading"
                ignore-filter
                @update:model-value="selectParent"
                value-key="value"
                placeholder="选择父文档"
                aria-label="选择父文档"
                :search-input="{ placeholder: '搜索标题或路径…' }"
                class="w-full"
              ><template #default><span class="truncate">{{ parentLabel }}</span></template></USelectMenu>
              <p class="mt-1 text-xs text-muted">最多显示 20 条，输入标题或路径搜索更多文档。</p>
              <div v-if="parentFailure" class="mt-2 text-sm text-error" role="alert">{{ parentFailure }} <UButton label="重试" variant="link" size="xs" @click="loadParents" /></div>
            </UFormField>

            <UFormField label="语言">
              <USelectMenu
                :model-value="form.locale"
                :items="localeOptions"
                value-key="value"
                aria-label="选择文档语言"
                class="w-full"
                @update:model-value="
                  form.locale = selectedValue($event) || defaultLocale
                "
              />
            </UFormField>
          </section>

          <section
            v-else-if="settingsSection === 'organization'"
            class="mt-5 space-y-5"
            data-docs-inspector-organization
          >
            <UFormField label="排序" help="同级文档按数值从小到大排列">
              <UInput
                v-model.number="form.sortOrder"
                type="number"
                min="0"
                class="w-full"
              />
            </UFormField>

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

            <div
              class="space-y-3 rounded-xl border border-default bg-default p-3"
            >
              <UFormField
                label="标题标记"
                help="在目录和标题旁显示一项简短状态，例如 Beta 或已废弃。"
              >
                <USelect
                  v-model="badgeMode"
                  :items="badgeModeItems"
                  value-key="value"
                  class="w-full"
                />
              </UFormField>
              <UFormField v-if="badgeMode === 'text'" label="标记文字">
                <UInput
                  v-model="form.badgeText"
                  maxlength="24"
                  placeholder="Beta"
                  class="w-full"
                />
              </UFormField>
              <AdminIconPicker
                v-else-if="badgeMode === 'icon'"
                :model-value="form.badgeIcon || 'i-tabler-sparkles'"
                @update:model-value="form.badgeIcon = $event"
              />
            </div>
          </section>

          <section v-else class="mt-5 space-y-5" data-docs-inspector-seo>
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

      <template #footer="{ close }">
        <div class="flex w-full justify-end gap-2">
          <UButton
            label="完成"
            color="neutral"
            variant="outline"
            @click="close"
          />
          <ActionFeedbackButton
            :status="saveStatus"
            idle-label="保存"
            pending-label="保存中"
            success-label="已保存"
            :disabled="!editorReady || Boolean(busy)"
          @click="save()"
          />
        </div>
      </template>
    </EditorInspector>

    <AssetImageProcessor
      :open="!!imageProcessingRequest"
      :file="imageProcessingRequest?.file"
      purpose="content"
      title="处理文档图片"
      :fixed-max-output-width="1200"
      fixed-output-type="image/webp"
      @update:open="(value) => !value && cancelImageProcessing()"
      @processed="finishImageProcessing($event.file)"
      @cancel="cancelImageProcessing"
      @error="
        toast.add({
          title: '图片处理失败',
          description: $event,
          color: 'error',
        })
      "
    />
  </div>
</template>
