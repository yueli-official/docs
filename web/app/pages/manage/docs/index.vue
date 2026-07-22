<script setup lang="ts">
import type { TableColumn } from "@nuxt/ui";
import { createPlatformNotifier } from "@platform/ui/feedback";
import {
  createCollectionRouteQueryCodec,
  createJsonCollectionQueryPolicy,
  type CollectionControl,
  type CollectionControlValue,
  type CollectionWorkflow,
} from "@yueli/ui/collection";
import { useVueCollectionWorkflow } from "@yueli/ui/collection/vue";
import { createVueRouterCollectionQuerySync } from "@yueli/ui/collection/vue-router";
import {
  CollectionTableToolbar,
  CollectionViewToggle,
} from "@yueli/ui/collection/pattern";
import { ManageEmpty, SkeletonList } from "@platform/manage/components";
import type {
  CollectionManageTree,
  CollectionVersion,
  CollectionVersionsResponse,
  CollectionView,
  DocDetail,
  ManageDocsResponse,
} from "~/types";
import {
  docManageRoute,
  findDocSlugPathById,
} from "~/utils/docsManageRoutes.mjs";

definePageMeta({ layout: "manage" });
useSeoMeta({ title: "文档 · 控制台" });

const { call } = useApi();
const toast = createPlatformNotifier(useToast());
const router = useRouter();
const UCheckbox = resolveComponent("UCheckbox");
const UButton = resolveComponent("UButton");

type StatusKey = "all" | "draft" | "published" | "archived" | "issues";
type BulkDocAction = "publish" | "draft" | "archive";
type SortKey = "updatedAt" | "title" | "path" | "sortOrder";
type SortDirection = "asc" | "desc";

interface ManagedDoc extends DocDetail {
  collectionTitle: string;
  collectionSlug: string;
  parentTitle: string;
  path: string;
  slugPath: string[];
  depth: number;
  versionKey: string;
  versionLabel: string;
  updatedAt: string;
}

interface MoveIntent {
  dragId: string;
  targetId: string;
  position: "before" | "after" | "into";
}
interface PatchItem {
  id: string;
  parentId: string;
  sortOrder: number;
}

interface DocCollectionQuery {
  q: string;
  status: StatusKey;
  page: number;
  size: number;
  view: "list" | "tree";
  collection: string;
  version: string;
  locale: string;
  parent: string;
  sort: SortKey;
  direction: SortDirection;
}
const statusKeys = ["all", "draft", "published", "archived", "issues"] as const;
const viewKeys = ["list", "tree"] as const;
const sortKeys = ["updatedAt", "title", "path", "sortOrder"] as const;
const directionKeys = ["asc", "desc"] as const;
const pageSizes = [30, 60, 90] as const;
const defaultQuery: DocCollectionQuery = {
  q: "",
  status: "all",
  page: 1,
  size: 30,
  view: "list",
  collection: "all",
  version: "",
  locale: "en",
  parent: "all",
  sort: "updatedAt",
  direction: "desc",
};
const queryPolicy = createJsonCollectionQueryPolicy<DocCollectionQuery>();
const querySync = createVueRouterCollectionQuerySync({
  router,
  codec: createCollectionRouteQueryCodec({
    q: { kind: "string", default: defaultQuery.q, maxLength: 200 },
    status: { kind: "enum", values: statusKeys, default: defaultQuery.status },
    page: { kind: "positive-integer", default: defaultQuery.page },
    size: {
      kind: "positive-integer",
      values: pageSizes,
      default: defaultQuery.size,
    },
    view: { kind: "enum", values: viewKeys, default: defaultQuery.view },
    collection: {
      kind: "string",
      default: defaultQuery.collection,
      maxLength: 200,
    },
    version: { kind: "string", default: defaultQuery.version, maxLength: 100 },
    locale: { kind: "string", default: defaultQuery.locale, maxLength: 50 },
    parent: { kind: "string", default: defaultQuery.parent, maxLength: 100 },
    sort: { kind: "enum", values: sortKeys, default: defaultQuery.sort },
    direction: {
      kind: "enum",
      values: directionKeys,
      default: defaultQuery.direction,
    },
  }),
});
const search = ref("");
const bulkAction = ref<BulkDocAction | undefined>();
const bulkBusy = ref(false);
const bulkResult = ref<{
  changed: number;
  failedIds: string[];
  interrupted?: boolean;
  message?: string;
}>();
const quickEditTarget = ref<ManagedDoc>();
const quickEditDocs = ref<ManagedDoc[]>([]);
const showQuickEdit = ref(false);
let quickEditRequestSeq = 0;
const docsPending = ref(false);
const docsError = ref("");
const knownDocs = shallowRef(new Map<string, ManagedDoc>());
const statusCounts = ref<Record<StatusKey, number>>({
  all: 0,
  draft: 0,
  published: 0,
  archived: 0,
  issues: 0,
});

const bulkItems: Array<{ label: string; value: BulkDocAction; icon: string }> =
  [
    { label: "发布", value: "publish", icon: "i-tabler-rocket" },
    { label: "转草稿", value: "draft", icon: "i-tabler-pencil" },
    { label: "归档", value: "archive", icon: "i-tabler-archive" },
  ];

const { data: cols, pending: collectionsPending } = await useAsyncData(
  "manage-doc-workbench-collections",
  () => call<{ items: CollectionView[] }>("/api/v1/collections"),
  { server: false, default: () => ({ items: [] as CollectionView[] }) },
);
const collections = computed(() => cols.value?.items ?? []);
const collectionById = computed(() =>
  Object.fromEntries(collections.value.map((c) => [c.id, c])),
);

function selectedValue(value: unknown) {
  if (value && typeof value === "object" && "value" in value) {
    const option = value as { value?: unknown };
    return String(option.value ?? "");
  }
  return String(value ?? "");
}

async function loadDocPage(
  nextQuery: Readonly<DocCollectionQuery>,
  activeWorkflow: CollectionWorkflow<ManagedDoc, string, DocCollectionQuery>,
) {
  const token = activeWorkflow.beginLoad();
  docsPending.value = true;
  docsError.value = "";
  try {
    const response = await call<ManageDocsResponse>("/api/v1/manage/docs", {
      query: {
        ...(nextQuery.q ? { q: nextQuery.q } : {}),
        status: nextQuery.status === "issues" ? "all" : nextQuery.status,
        quality: nextQuery.status === "issues" ? "issues" : "all",
        ...(nextQuery.collection !== "all"
          ? { collectionId: nextQuery.collection }
          : {}),
        ...(nextQuery.version ? { version: nextQuery.version } : {}),
        ...(nextQuery.locale !== "all" ? { locale: nextQuery.locale } : {}),
        ...(nextQuery.parent !== "all" ? { parentId: nextQuery.parent } : {}),
        sort: nextQuery.sort,
        direction: nextQuery.direction,
        page: nextQuery.page,
        size: nextQuery.size,
      },
    });
    const items = response.items.map(normalizeManagedDoc);
    const nextKnownDocs = new Map(knownDocs.value);
    for (const item of items) nextKnownDocs.set(item.id, item);
    knownDocs.value = nextKnownDocs;
    statusCounts.value = response.counts;
    const lastPage = Math.max(1, Math.ceil(response.total / nextQuery.size));
    if (nextQuery.page > lastPage) {
      activeWorkflow.setQuery({ ...nextQuery, page: lastPage });
      return;
    }
    activeWorkflow.resolveLoad(token, { items, total: response.total });
  } catch (error) {
    const apiError = error as { data?: { message?: string } };
    docsError.value = apiError.data?.message || "加载文档失败";
    activeWorkflow.rejectLoad(token, { key: "docs.manage.load_failed" });
  } finally {
    docsPending.value = false;
  }
}

const {
  snapshot: docCollection,
  workflow: docWorkflow,
  reload: reloadDocPage,
} = useVueCollectionWorkflow({
  initialQuery: defaultQuery,
  queryPolicy,
  keyOf: (doc: ManagedDoc) => doc.id,
  querySync,
  dataQueryKey: (query) => JSON.stringify(query),
  load: loadDocPage,
});
const collectionQuery = computed(() => docCollection.value.query);
function updateCollectionQuery(
  patch: Partial<DocCollectionQuery>,
  resetPage = true,
) {
  docWorkflow.setQuery({
    ...collectionQuery.value,
    ...patch,
    ...(resetPage ? { page: 1 } : {}),
  });
}
const activeCollection = computed({
  get: () => collectionQuery.value.collection,
  set: (value: string) =>
    updateCollectionQuery({ collection: value, version: "", parent: "all" }),
});
const activeVersion = computed({
  get: () => collectionQuery.value.version,
  set: (value: string) => updateCollectionQuery({ version: value }),
});
const activeLocale = computed({
  get: () => collectionQuery.value.locale,
  set: (value: string) => updateCollectionQuery({ locale: value || "all" }),
});
const activeStatus = computed<StatusKey>({
  get: () => collectionQuery.value.status,
  set: (value) => updateCollectionQuery({ status: value }),
});
const activeParent = computed({
  get: () => collectionQuery.value.parent,
  set: (value: string) => updateCollectionQuery({ parent: value }),
});
const activeSort = computed<SortKey>({
  get: () => collectionQuery.value.sort,
  set: (value) => updateCollectionQuery({ sort: value }),
});
const activeDirection = computed<SortDirection>({
  get: () => collectionQuery.value.direction,
  set: (value) => updateCollectionQuery({ direction: value }),
});
const page = computed({
  get: () => collectionQuery.value.page,
  set: (value: number) => updateCollectionQuery({ page: value }, false),
});
const pageSize = computed<(typeof pageSizes)[number]>({
  get: () => collectionQuery.value.size as (typeof pageSizes)[number],
  set: (value) => updateCollectionQuery({ size: value }),
});
function setPageSize(value: unknown) {
  const next = Number(value);
  if (pageSizes.includes(next as (typeof pageSizes)[number])) {
    pageSize.value = next as (typeof pageSizes)[number];
  }
}
const viewMode = computed({
  get: () => collectionQuery.value.view,
  set: (value: "list" | "tree") =>
    updateCollectionQuery({ view: value }, false),
});
let searchTimer: ReturnType<typeof setTimeout> | undefined;
search.value = collectionQuery.value.q;
watch(search, (value) => {
  if (searchTimer) clearTimeout(searchTimer);
  searchTimer = setTimeout(
    () => updateCollectionQuery({ q: value.trim() }),
    300,
  );
});
watch(
  () => collectionQuery.value.q,
  (value) => {
    if (search.value !== value) search.value = value;
  },
);
onScopeDispose(() => {
  if (searchTimer) clearTimeout(searchTimer);
});
function submitSearch(value: string) {
  if (searchTimer) clearTimeout(searchTimer);
  search.value = value;
  updateCollectionQuery({ q: value.trim() });
}

const versionsByCollection = ref<Record<string, CollectionVersion[]>>({});
const localeItems = [
  { label: "全部语言", value: "all" },
  { label: "English", value: "en" },
  { label: "简体中文", value: "zh-CN" },
];

const collectionItems = computed(() => [
  { label: "全部文档集", value: "all" },
  ...collections.value.map((c) => ({ label: c.title, value: c.id })),
]);
const selectedCollectionVersions = computed(() =>
  activeCollection.value === "all"
    ? []
    : (versionsByCollection.value[activeCollection.value] ?? []),
);
const versionItems = computed(() => [
  { label: "默认版本", value: "" },
  ...selectedCollectionVersions.value.map((v) => ({
    label: v.isDefault ? `${v.label}（默认）` : v.label,
    value: v.key,
    description: v.key,
  })),
]);
const treeCollectionItems = computed(() =>
  collections.value.map((c) => ({ label: c.title, value: c.slug })),
);

async function loadVersions() {
  const items = collections.value;
  if (!items.length) {
    versionsByCollection.value = {};
    return;
  }
  const entries = await Promise.all(
    items.map(async (col) => {
      const res = await call<CollectionVersionsResponse>(
        `/api/v1/collections/${col.id}/versions`,
      );
      return [col.id, res.items] as const;
    }),
  );
  versionsByCollection.value = Object.fromEntries(entries);
}

watch(
  collections,
  () => {
    loadVersions();
  },
  { immediate: true },
);

function normalizeManagedDoc(
  doc: ManageDocsResponse["items"][number],
): ManagedDoc {
  const slugPath = doc.slugPath.split("/").filter(Boolean);
  return {
    ...doc,
    content: "",
    translationKey: "",
    slugPath,
    path: slugPath.join(" / "),
    depth: Math.max(0, slugPath.length - 1),
  };
}

const statusItems = computed(() => {
  return [
    { value: "all", label: `全部 · ${statusCounts.value.all}` },
    { value: "draft", label: `草稿 · ${statusCounts.value.draft}` },
    { value: "published", label: `已发布 · ${statusCounts.value.published}` },
    { value: "archived", label: `归档 · ${statusCounts.value.archived}` },
    { value: "issues", label: `待完善 · ${statusCounts.value.issues}` },
  ];
});

function qualityIssues(doc: ManagedDoc) {
  const issues: string[] = [];
  if (!doc.excerpt?.trim()) issues.push("缺摘要");
  if (!doc.slugPath.length) issues.push("缺路径");
  if (!doc.collectionSlug) issues.push("缺文档集");
  return issues;
}

const pagedDocs = computed(() => docCollection.value.items);
const selectedDocIds = computed<readonly string[]>(() =>
  docCollection.value.selection.mode === "keys"
    ? docCollection.value.selection.keys
    : [],
);
function clearDocSelection() {
  docWorkflow.clearSelection();
}
function replaceSelection(ids: readonly string[]) {
  docWorkflow.clearSelection();
  for (const id of ids) docWorkflow.toggleKey(id);
}
const rowSelection = computed<Record<string, boolean>>({
  get: () =>
    Object.fromEntries(selectedDocIds.value.map((id) => [id, true] as const)),
  set: (selection) =>
    replaceSelection(
      Object.entries(selection)
        .filter(([, selected]) => selected)
        .map(([id]) => id),
    ),
});
const selectedDocs = computed(() =>
  selectedDocIds.value.flatMap((id) => {
    const doc = knownDocs.value.get(id);
    return doc ? [doc] : [];
  }),
);
const parentItems = computed(() => {
  const items = [
    { label: "全部层级", value: "all" },
    { label: "仅顶级文档", value: "root" },
  ];
  if (
    activeCollection.value === "all" ||
    tree.value?.collection.id !== activeCollection.value
  ) {
    return items;
  }
  return [
    ...items,
    ...flattenTree(tree.value.tree).map(({ doc, depth }) => ({
      label: `${"　".repeat(depth)}${doc.title}`,
      value: doc.id,
    })),
  ];
});
const collectionFilterControls = computed<CollectionControl[]>(() => [
  {
    kind: "select",
    id: "status",
    label: "状态",
    value: activeStatus.value,
    options: statusItems.value,
    class: "w-32",
  },
  {
    kind: "select",
    id: "collection",
    label: "文档集",
    value: activeCollection.value,
    options: collectionItems.value,
    searchPlaceholder: "搜索文档集…",
    class: "w-40",
  },
  {
    kind: "select",
    id: "locale",
    label: "语言",
    value: activeLocale.value,
    options: localeItems,
    class: "w-32",
  },
  ...(activeCollection.value === "all"
    ? []
    : [
        {
          kind: "select" as const,
          id: "parent",
          label: "父级",
          value: activeParent.value,
          options: parentItems.value,
          searchPlaceholder: "搜索父级文档…",
          class: "w-40",
        },
      ]),
  ...(activeCollection.value === "all"
    ? []
    : [
        {
          kind: "select" as const,
          id: "version",
          label: "版本",
          value: activeVersion.value,
          options: versionItems.value,
          class: "w-36",
        },
      ]),
]);
const activeFilterCount = computed(
  () =>
    [
      activeStatus.value !== "all",
      activeCollection.value !== "all",
      activeLocale.value !== "en",
      Boolean(activeVersion.value),
      activeParent.value !== "all",
    ].filter(Boolean).length,
);
const activeFilterItems = computed(() => {
  const items: Array<{ id: string; label: string }> = [];
  if (activeStatus.value !== "all") {
    const label = statusItems.value
      .find((item) => item.value === activeStatus.value)
      ?.label.split(" · ")[0];
    items.push({ id: "status", label: `状态：${label ?? activeStatus.value}` });
  }
  if (activeCollection.value !== "all") {
    const label = collectionItems.value.find(
      (item) => item.value === activeCollection.value,
    )?.label;
    items.push({
      id: "collection",
      label: `文档集：${label ?? activeCollection.value}`,
    });
  }
  if (activeLocale.value !== "en") {
    const label = localeItems.find(
      (item) => item.value === activeLocale.value,
    )?.label;
    items.push({ id: "locale", label: `语言：${label ?? activeLocale.value}` });
  }
  if (activeVersion.value) {
    const label = versionItems.value.find(
      (item) => item.value === activeVersion.value,
    )?.label;
    items.push({
      id: "version",
      label: `版本：${label ?? activeVersion.value}`,
    });
  }
  if (activeParent.value !== "all") {
    const label = parentItems.value.find(
      (item) => item.value === activeParent.value,
    )?.label;
    items.push({ id: "parent", label: `父级：${label ?? activeParent.value}` });
  }
  return items;
});
const collectionState = computed(() => {
  if (docsError.value) return "error";
  if (collectionsPending.value || docsPending.value) return "loading";
  return "ready";
});
function changeCollectionControl(id: string, value: CollectionControlValue) {
  if (typeof value !== "string") return;
  if (id === "status" && statusKeys.includes(value as StatusKey))
    activeStatus.value = value as StatusKey;
  if (id === "collection") activeCollection.value = value;
  if (id === "locale") activeLocale.value = value;
  if (id === "version") activeVersion.value = value;
  if (id === "parent") activeParent.value = value;
}
function clearCollectionFilter(id: string) {
  if (id === "status") activeStatus.value = "all";
  if (id === "collection") activeCollection.value = "all";
  if (id === "locale") activeLocale.value = "en";
  if (id === "version") activeVersion.value = "";
  if (id === "parent") activeParent.value = "all";
}
function clearCollectionFilters() {
  updateCollectionQuery({
    status: "all",
    collection: "all",
    version: "",
    locale: "en",
    parent: "all",
  });
}
const pageSizeItems = pageSizes.map((value) => ({
  label: `${value} 篇`,
  value,
}));
const firstVisibleDoc = computed(() =>
  docCollection.value.total === 0 ? 0 : (page.value - 1) * pageSize.value + 1,
);
const lastVisibleDoc = computed(() =>
  Math.min(docCollection.value.total, page.value * pageSize.value),
);
const getDocRowId = (doc: ManagedDoc) => doc.id;

watch(
  [
    activeCollection,
    activeStatus,
    activeLocale,
    activeVersion,
    activeParent,
    activeSort,
    activeDirection,
  ],
  () => {
    bulkAction.value = undefined;
    bulkResult.value = undefined;
    clearDocSelection();
  },
);
watch(activeCollection, (value) => {
  if (value === "all" && activeVersion.value) activeVersion.value = "";
});

const selectedCollection = computed(() =>
  activeCollection.value === "all"
    ? null
    : collectionById.value[activeCollection.value],
);

const createTarget = computed(() => {
  const slug = selectedCollection.value?.slug;
  return slug
    ? `/manage/docs/new?collection=${encodeURIComponent(slug)}`
    : "/manage/docs/new";
});

function openDoc(docOrId: ManagedDoc | string) {
  const doc =
    typeof docOrId === "string" ? knownDocs.value.get(docOrId) : docOrId;
  if (!doc && typeof docOrId === "string") {
    const slugPath = findDocSlugPathById(tree.value?.tree ?? [], docOrId);
    if (slugPath && activeTreeSlug.value) {
      navigateTo(docManageRoute(activeTreeSlug.value, slugPath));
      return;
    }
  }
  if (!doc) {
    navigateTo(`/manage/docs/${docOrId}`);
    return;
  }
  navigateTo(docManageRoute(doc.collectionSlug, doc.slugPath));
}

function addChild(doc: ManagedDoc) {
  navigateTo(
    `/manage/docs/new?collection=${encodeURIComponent(doc.collectionSlug)}&parent=${encodeURIComponent(doc.id)}`,
  );
}

async function openQuickEdit(doc: ManagedDoc) {
  const seq = ++quickEditRequestSeq;
  quickEditTarget.value = doc;
  quickEditDocs.value = [doc];
  showQuickEdit.value = true;
  try {
    const source = await fetchManageTree(
      doc.collectionSlug,
      doc.locale,
      doc.versionKey,
    );
    if (seq === quickEditRequestSeq && quickEditTarget.value?.id === doc.id) {
      quickEditDocs.value = flattenManagedTree(source, doc.versionKey);
    }
  } catch {
    if (seq === quickEditRequestSeq) quickEditDocs.value = [doc];
  }
}

async function onQuickEditSaved() {
  await Promise.all([reloadDocPage(), refreshTree()]);
}

function clearBulkSelection() {
  clearDocSelection();
  bulkAction.value = undefined;
}

async function updateDocStatus(
  id: string,
  status: "draft" | "published" | "archived",
) {
  if (status === "published") {
    await call(`/api/v1/docs/${id}/publish`, { method: "POST" });
    return;
  }
  if (status === "archived") {
    await call(`/api/v1/docs/${id}/archive`, { method: "POST" });
    return;
  }
  await call(`/api/v1/docs/${id}`, { method: "PATCH", body: { status } });
}

async function applyBulkAction() {
  if (!bulkAction.value || !selectedDocs.value.length) return;

  const status =
    bulkAction.value === "publish"
      ? "published"
      : bulkAction.value === "archive"
        ? "archived"
        : "draft";
  const docs = [...selectedDocs.value];
  const requestedIds = docs.map((doc) => doc.id);
  bulkBusy.value = true;
  bulkResult.value = undefined;
  try {
    const results = await Promise.allSettled(
      docs.map((doc) => updateDocStatus(doc.id, status)),
    );
    const changed = results.filter(
      (result) => result.status === "fulfilled",
    ).length;
    const failedIds = results.flatMap((result, index) =>
      result.status === "rejected" ? [requestedIds[index]!] : [],
    );
    bulkResult.value = { changed, failedIds };
    if (failedIds.length) replaceSelection(failedIds);
    else clearBulkSelection();
    bulkAction.value = undefined;
    await Promise.all([reloadDocPage(), refreshTree()]);
  } catch (error) {
    const apiError = error as { data?: { message?: string } };
    replaceSelection(requestedIds);
    bulkResult.value = {
      changed: 0,
      failedIds: requestedIds,
      interrupted: true,
      message:
        apiError.data?.message ||
        "批量请求中断，已保留选择，请核对当前状态后重试。",
    };
    await Promise.all([reloadDocPage(), refreshTree()]);
  } finally {
    bulkBusy.value = false;
  }
}

// ─── Structure view ──────────────────────────────────────────────────────────

const treeSlug = ref("");
const activeTreeSlug = computed(() => selectedValue(treeSlug.value));
watch(
  collections,
  (items) => {
    if (!activeTreeSlug.value && items.length) treeSlug.value = items[0]!.slug;
  },
  { immediate: true },
);
watch(activeCollection, (id) => {
  if (id !== "all")
    treeSlug.value = collectionById.value[id]?.slug ?? treeSlug.value;
});

const activeTreeCollection = computed(
  () => collections.value.find((c) => c.slug === activeTreeSlug.value) ?? null,
);
const treeOverride = ref<CollectionManageTree | null>(null);
const loadedTree = ref<CollectionManageTree | null>(null);
const treePending = ref(false);
let treeRequestSeq = 0;
const tree = computed<CollectionManageTree | null>(
  () => treeOverride.value ?? loadedTree.value,
);

async function fetchManageTree(
  slug: string,
  locale: string,
  version = "",
): Promise<CollectionManageTree> {
  return call<CollectionManageTree>(
    `/api/v1/manage/collections/${encodeURIComponent(slug)}/tree`,
    {
      query: {
        locale: locale === "all" ? "en" : locale,
        ...(version ? { version } : {}),
      },
    },
  );
}

async function refreshTree() {
  const slug = activeTreeSlug.value;
  if (!slug) {
    loadedTree.value = null;
    return;
  }
  const seq = ++treeRequestSeq;
  treeOverride.value = null;
  treePending.value = true;
  try {
    const version =
      activeTreeCollection.value?.id === activeCollection.value
        ? activeVersion.value
        : "";
    const result = await fetchManageTree(slug, activeLocale.value, version);
    if (seq === treeRequestSeq) loadedTree.value = result;
  } finally {
    if (seq === treeRequestSeq) treePending.value = false;
  }
}

watch([activeTreeSlug, activeLocale, activeVersion], () => void refreshTree(), {
  immediate: true,
});

function flattenTree(
  nodes: readonly DocDetail[],
  depth = 0,
): Array<{ doc: DocDetail; depth: number }> {
  return nodes.flatMap((doc) => [
    { doc, depth },
    ...flattenTree(doc.children ?? [], depth + 1),
  ]);
}

function flattenManagedTree(
  source: CollectionManageTree,
  versionKey = "",
): ManagedDoc[] {
  function visit(
    nodes: readonly DocDetail[],
    parentTitle = "",
    parentSlugs: string[] = [],
    depth = 0,
  ): ManagedDoc[] {
    return nodes.flatMap((doc) => {
      const slugPath = [...parentSlugs, doc.slug].filter(Boolean);
      const managed: ManagedDoc = {
        ...doc,
        collectionTitle: source.collection.title,
        collectionSlug: source.collection.slug,
        parentTitle,
        path: slugPath.join(" / "),
        slugPath,
        depth,
        versionKey,
        versionLabel: versionKey,
        updatedAt: "",
      };
      return [
        managed,
        ...visit(doc.children ?? [], doc.title, slugPath, depth + 1),
      ];
    });
  }
  return visit(source.tree);
}

function findInTree(
  nodes: DocDetail[],
  id: string,
  parent: DocDetail | null = null,
): {
  node: DocDetail;
  parent: DocDetail | null;
  siblings: DocDetail[];
  index: number;
} | null {
  for (let i = 0; i < nodes.length; i++) {
    const n = nodes[i]!;
    if (n.id === id) return { node: n, parent, siblings: nodes, index: i };
    if (n.children?.length) {
      const r = findInTree(n.children, id, n);
      if (r) return r;
    }
  }
  return null;
}

function isDescendant(node: DocDetail, targetId: string): boolean {
  return (
    node.children?.some(
      (c) => c.id === targetId || isDescendant(c, targetId),
    ) ?? false
  );
}

function computePatches(
  treeNodes: DocDetail[],
  intent: MoveIntent,
): PatchItem[] | null {
  const dragCtx = findInTree(treeNodes, intent.dragId);
  const targetCtx = findInTree(treeNodes, intent.targetId);
  if (!dragCtx || !targetCtx) return null;

  const {
    node: dragNode,
    parent: dragParent,
    siblings: dragSiblings,
    index: dragIdx,
  } = dragCtx;
  if (isDescendant(dragNode, intent.targetId)) return null;

  const oldParentId = dragParent?.id ?? "";
  let newParentId: string;
  let newSiblings: DocDetail[];
  let insertIdx: number;

  if (intent.position === "into") {
    newParentId = intent.targetId;
    newSiblings = targetCtx.node.children ?? [];
    insertIdx = newSiblings.length;
  } else {
    newParentId = targetCtx.parent?.id ?? "";
    newSiblings = targetCtx.siblings;
    insertIdx =
      intent.position === "before" ? targetCtx.index : targetCtx.index + 1;
  }

  const patches: PatchItem[] = [];
  const isSameParent = oldParentId === newParentId;

  if (isSameParent) {
    const reordered = [...dragSiblings];
    reordered.splice(dragIdx, 1);
    const adj = insertIdx > dragIdx ? insertIdx - 1 : insertIdx;
    reordered.splice(adj, 0, dragNode);
    reordered.forEach((n, i) => {
      if (n.sortOrder !== i)
        patches.push({ id: n.id, parentId: newParentId, sortOrder: i });
    });
  } else {
    dragSiblings
      .filter((_, i) => i !== dragIdx)
      .forEach((n, i) => {
        if (n.sortOrder !== i)
          patches.push({ id: n.id, parentId: oldParentId, sortOrder: i });
      });
    const newOrder = [...newSiblings];
    newOrder.splice(insertIdx, 0, dragNode);
    newOrder.forEach((n, i) => {
      if (n.sortOrder !== i || n.id === dragNode.id)
        patches.push({ id: n.id, parentId: newParentId, sortOrder: i });
    });
  }

  return patches.length > 0 ? patches : null;
}

function applyMoveToTree(nodes: DocDetail[], intent: MoveIntent) {
  const dragCtx = findInTree(nodes, intent.dragId);
  if (!dragCtx) return;
  const { node: dragNode, siblings: oldSiblings, index: dragIdx } = dragCtx;

  const targetBefore = findInTree(nodes, intent.targetId);
  const newParentId =
    intent.position === "into"
      ? intent.targetId
      : (targetBefore?.parent?.id ?? "");

  oldSiblings.splice(dragIdx, 1);
  oldSiblings.forEach((n, i) => {
    n.sortOrder = i;
  });
  dragNode.parentId = newParentId;

  const targetAfter = findInTree(nodes, intent.targetId);
  if (!targetAfter) return;

  if (intent.position === "into") {
    if (!targetAfter.node.children) targetAfter.node.children = [];
    targetAfter.node.children.push(dragNode);
    targetAfter.node.children.forEach((n, i) => {
      n.sortOrder = i;
    });
  } else {
    const siblings = targetAfter.parent?.children ?? nodes;
    const idx =
      intent.position === "before" ? targetAfter.index : targetAfter.index + 1;
    siblings.splice(idx, 0, dragNode);
    siblings.forEach((n, i) => {
      n.sortOrder = i;
    });
  }
}

async function onMove(intent: MoveIntent) {
  if (!tree.value) return;
  const patches = computePatches(tree.value.tree, intent);
  if (!patches?.length) return;

  const snapshot = JSON.parse(
    JSON.stringify(tree.value),
  ) as CollectionManageTree;
  const optimistic = JSON.parse(
    JSON.stringify(tree.value),
  ) as CollectionManageTree;
  applyMoveToTree(optimistic.tree, intent);
  treeOverride.value = optimistic;

  try {
    await Promise.all(
      patches.map((p) =>
        call(`/api/v1/docs/${p.id}`, {
          method: "PATCH",
          body: {
            parentId: p.parentId,
            sortOrder: p.sortOrder,
          },
        }),
      ),
    );
    await Promise.all([reloadDocPage(), refreshTree()]);
  } catch {
    treeOverride.value = snapshot;
    toast.add({ title: "移动失败，已还原", color: "error" });
    await refreshTree();
  }
}

async function onDelete(id: string) {
  try {
    await call(`/api/v1/docs/${id}`, { method: "DELETE" });
    await Promise.all([refreshTree(), reloadDocPage()]);
  } catch (err: any) {
    toast.add({
      title: "删除失败",
      description: err?.data?.message || "请重试",
      color: "error",
    });
  }
}

function formatUpdatedAt(value: string) {
  if (!value) return "—";
  return new Intl.DateTimeFormat("zh-CN", {
    year: "numeric",
    month: "2-digit",
    day: "2-digit",
  }).format(new Date(value));
}

function statusLabel(status: string) {
  return (
    { draft: "草稿", published: "已发布", archived: "已归档" }[status] ?? status
  );
}

function changeColumnSort(key: SortKey) {
  if (activeSort.value === key) {
    updateCollectionQuery({
      direction: activeDirection.value === "asc" ? "desc" : "asc",
    });
    return;
  }
  updateCollectionQuery({
    sort: key,
    direction: key === "updatedAt" ? "desc" : "asc",
  });
}

function sortableHeader(label: string, key: SortKey) {
  return () => {
    const active = activeSort.value === key;
    const icon = active
      ? activeDirection.value === "asc"
        ? "i-tabler-arrow-up"
        : "i-tabler-arrow-down"
      : "i-tabler-arrows-sort";
    return h(UButton, {
      label,
      icon,
      color: "neutral",
      variant: "ghost",
      size: "xs",
      class: "-mx-2",
      "aria-pressed": active,
      onClick: () => changeColumnSort(key),
    });
  };
}

const docColumns: TableColumn<ManagedDoc>[] = [
  {
    id: "select",
    header: ({ table }) =>
      h(UCheckbox, {
        modelValue: table.getIsSomePageRowsSelected()
          ? "indeterminate"
          : table.getIsAllPageRowsSelected(),
        disabled: bulkBusy.value,
        ariaLabel: "选择当前页文档",
        "onUpdate:modelValue": (value: boolean | "indeterminate") =>
          table.toggleAllPageRowsSelected(value === true),
      }),
    cell: ({ row }) =>
      h(UCheckbox, {
        modelValue: row.getIsSelected(),
        disabled: bulkBusy.value,
        ariaLabel: `选择文档：${row.original.title}`,
        "onUpdate:modelValue": (value: boolean | "indeterminate") =>
          row.toggleSelected(value === true),
      }),
    enableSorting: false,
    enableHiding: false,
    meta: {
      class: {
        th: "w-11 px-4",
        td: "w-11 px-4",
      },
    },
  },
  {
    accessorKey: "title",
    header: sortableHeader("标题", "title"),
    meta: {
      class: {
        th: "w-[40%] min-w-60",
        td: "w-[40%] min-w-60",
      },
    },
  },
  {
    accessorKey: "path",
    header: sortableHeader("路径 / 父级", "path"),
    meta: {
      class: {
        th: "hidden w-[22%] lg:table-cell",
        td: "hidden w-[22%] lg:table-cell",
      },
    },
  },
  {
    id: "collection",
    accessorFn: (doc) => doc.collectionTitle,
    header: "文档集 / 版本",
    meta: {
      class: {
        th: "hidden w-[18%] xl:table-cell",
        td: "hidden w-[18%] xl:table-cell",
      },
    },
  },
  {
    accessorKey: "status",
    header: "状态",
    meta: {
      class: {
        th: "hidden w-24 md:table-cell",
        td: "hidden w-24 md:table-cell",
      },
    },
  },
  {
    id: "updated",
    accessorFn: (doc) => doc.updatedAt,
    header: sortableHeader("更新", "updatedAt"),
    meta: {
      class: {
        th: "hidden w-28 lg:table-cell",
        td: "hidden w-28 lg:table-cell",
      },
    },
  },
  {
    id: "actions",
    header: "操作",
    enableSorting: false,
    enableHiding: false,
    meta: {
      class: {
        th: "w-28 text-right",
        td: "w-28 text-right",
      },
    },
  },
];
</script>

<template>
  <YAdminPage
    id="documents"
    title="文档"
    icon="i-tabler-file-text"
    main-id="manage-main"
    body-class="mx-auto w-full max-w-screen-2xl"
  >
    <template #actions>
      <UButton icon="i-tabler-plus" label="新建文档" :to="createTarget" />
    </template>

    <ClientOnly>
      <section
        v-if="viewMode === 'list'"
        class="overflow-hidden rounded-xl border border-default bg-default shadow-sm"
        :inert="bulkBusy"
        :aria-busy="bulkBusy"
        aria-label="文档列表"
      >
        <CollectionTableToolbar
          v-model:search="search"
          label="文档列表工具栏"
          search-placeholder="搜索标题、路径、slug 或文档集…"
          search-action="搜索"
          filter-label="筛选"
          :filter-count="activeFilterCount"
          :selection-count="selectedDocIds.length"
          @search="submitSearch"
        >
          <template #filters>
            <div
              class="grid w-80 max-w-[calc(100vw-2rem)] gap-3"
              aria-label="文档筛选条件"
            >
              <UFormField
                v-for="control in collectionFilterControls"
                :key="control.id"
                :label="control.label"
              >
                <USelectMenu
                  v-if="control.kind === 'select' && control.searchPlaceholder"
                  :model-value="control.value"
                  :items="control.options.slice()"
                  value-key="value"
                  :search-input="{
                    placeholder: control.searchPlaceholder,
                  }"
                  :aria-label="control.label"
                  size="sm"
                  class="w-full"
                  @update:model-value="
                    changeCollectionControl(control.id, selectedValue($event))
                  "
                />
                <USelect
                  v-else-if="control.kind === 'select'"
                  :model-value="control.value"
                  :items="control.options.slice()"
                  value-key="value"
                  :aria-label="control.label"
                  size="sm"
                  class="w-full"
                  @update:model-value="
                    changeCollectionControl(control.id, selectedValue($event))
                  "
                />
              </UFormField>

              <div
                v-if="activeFilterCount"
                class="flex justify-end border-t border-default pt-3"
              >
                <UButton
                  label="清除全部筛选"
                  color="neutral"
                  variant="ghost"
                  size="xs"
                  @click="clearCollectionFilters"
                />
              </div>
            </div>
          </template>

          <template #utilities>
            <CollectionViewToggle
              v-model="viewMode"
              :items="[
                { key: 'list', label: '列表', icon: 'i-tabler-list' },
                { key: 'tree', label: '树状', icon: 'i-tabler-sitemap' },
              ]"
            />
          </template>

          <template #active-filters>
            <span class="mr-1 text-xs text-muted">已应用</span>
            <UButton
              v-for="filter in activeFilterItems"
              :key="filter.id"
              :label="filter.label"
              trailing-icon="i-tabler-x"
              color="neutral"
              variant="soft"
              size="xs"
              @click="clearCollectionFilter(filter.id)"
            />
            <UButton
              label="清除全部"
              color="neutral"
              variant="ghost"
              size="xs"
              @click="clearCollectionFilters"
            />
          </template>

          <template #selection>
            <div
              data-docs-bulk-actions
              class="flex min-w-0 items-center justify-between gap-2"
            >
              <span class="shrink-0 text-xs font-medium text-highlighted">
                已选择 {{ selectedDocIds.length }} 篇
              </span>
              <div class="flex min-w-0 items-center justify-end gap-1.5">
                <USelect
                  v-model="bulkAction"
                  :items="bulkItems"
                  value-key="value"
                  placeholder="批量操作"
                  size="xs"
                  class="w-28 min-w-0"
                />
                <UButton
                  class="min-[24rem]:hidden"
                  icon="i-tabler-check"
                  aria-label="应用批量操作"
                  color="primary"
                  variant="soft"
                  size="xs"
                  square
                  :disabled="!bulkAction"
                  :loading="bulkBusy"
                  @click="applyBulkAction"
                />
                <UButton
                  class="hidden min-[24rem]:inline-flex"
                  label="应用"
                  icon="i-tabler-check"
                  color="primary"
                  variant="soft"
                  size="xs"
                  :disabled="!bulkAction"
                  :loading="bulkBusy"
                  @click="applyBulkAction"
                />
                <UButton
                  icon="i-tabler-x"
                  aria-label="取消选择"
                  color="neutral"
                  variant="ghost"
                  size="xs"
                  square
                  @click="clearBulkSelection"
                />
              </div>
            </div>
          </template>
        </CollectionTableToolbar>

        <div
          v-if="collectionState === 'error'"
          class="grid min-h-56 place-items-center px-6 py-12 text-center"
          role="alert"
        >
          <div>
            <span
              class="mx-auto grid size-10 place-items-center rounded-full bg-error/10 text-error"
            >
              <UIcon name="i-tabler-alert-circle" class="size-5" />
            </span>
            <p class="mt-3 text-sm font-medium text-highlighted">
              文档加载失败
            </p>
            <p class="mt-1 max-w-md text-xs leading-5 text-muted">
              {{ docsError }}
            </p>
            <UButton
              class="mt-4"
              label="重新加载"
              color="neutral"
              variant="outline"
              size="xs"
              @click="reloadDocPage"
            />
          </div>
        </div>

        <UTable
          v-else
          v-model:row-selection="rowSelection"
          :data="pagedDocs.slice()"
          :columns="docColumns"
          :get-row-id="getDocRowId"
          :loading="collectionState === 'loading'"
          class="shrink-0"
          :ui="{
            root: 'overflow-x-auto',
            base: 'table-fixed border-separate border-spacing-0',
            thead: '[&>tr]:bg-elevated/50 [&>tr]:after:content-none',
            tbody: '[&>tr]:last:[&>td]:border-b-0',
            th: 'border-b border-default px-4 py-2 text-xs font-medium text-muted',
            td: 'border-b border-default px-4 py-3 align-middle',
            separator: 'h-0',
          }"
        >
          <template #title-cell="{ row }">
            <div class="min-w-0">
              <button
                type="button"
                class="block max-w-full truncate text-left text-sm font-medium text-highlighted hover:text-primary"
                @click="openQuickEdit(row.original)"
              >
                {{ row.original.title }}
              </button>
              <p class="mt-1 truncate text-xs text-muted">
                {{ row.original.excerpt || "暂无摘要" }}
              </p>
              <p class="mt-1 truncate font-mono text-xs text-dimmed lg:hidden">
                /{{ row.original.slugPath.join("/") }} ·
                {{ row.original.collectionTitle }}
              </p>
              <p
                v-if="qualityIssues(row.original).length"
                class="mt-1 inline-flex max-w-full items-center gap-1 truncate text-xs text-warning"
              >
                <UIcon name="i-tabler-alert-circle" class="size-3.5 shrink-0" />
                <span class="truncate">{{
                  qualityIssues(row.original).slice(0, 2).join(" · ")
                }}</span>
              </p>
            </div>
          </template>

          <template #path-cell="{ row }">
            <div class="min-w-0 text-xs">
              <p class="truncate font-mono text-default">
                /{{ row.original.slugPath.join("/") }}
              </p>
              <p class="mt-1 truncate text-muted">
                {{ row.original.parentTitle || "顶级文档" }}
              </p>
            </div>
          </template>

          <template #collection-cell="{ row }">
            <div class="min-w-0 text-xs">
              <p class="truncate text-default">
                {{ row.original.collectionTitle }}
              </p>
              <p class="mt-1 truncate text-muted">
                {{ row.original.versionLabel || "默认版本" }} ·
                {{ row.original.locale }}
              </p>
            </div>
          </template>

          <template #status-cell="{ row }">
            <span
              class="inline-flex rounded-md px-2 py-1 text-xs font-medium"
              :class="{
                'bg-success/10 text-success':
                  row.original.status === 'published',
                'bg-warning/10 text-warning': row.original.status === 'draft',
                'bg-elevated text-muted': row.original.status === 'archived',
              }"
            >
              {{ statusLabel(row.original.status) }}
            </span>
          </template>

          <template #updated-cell="{ row }">
            <time class="text-xs text-muted" :datetime="row.original.updatedAt">
              {{ formatUpdatedAt(row.original.updatedAt) }}
            </time>
          </template>

          <template #actions-cell="{ row }">
            <div class="flex justify-end gap-1">
              <UTooltip text="添加子文档">
                <UButton
                  icon="i-tabler-file-plus"
                  color="neutral"
                  variant="ghost"
                  size="xs"
                  square
                  :aria-label="`添加子文档：${row.original.title}`"
                  @click="addChild(row.original)"
                />
              </UTooltip>
              <UTooltip text="快速编辑">
                <UButton
                  icon="i-tabler-pencil"
                  color="neutral"
                  variant="ghost"
                  size="xs"
                  square
                  :aria-label="`快速编辑：${row.original.title}`"
                  @click="openQuickEdit(row.original)"
                />
              </UTooltip>
              <UTooltip text="完整编辑">
                <UButton
                  icon="i-tabler-file-pencil"
                  color="neutral"
                  variant="ghost"
                  size="xs"
                  square
                  :aria-label="`完整编辑：${row.original.title}`"
                  @click="openDoc(row.original)"
                />
              </UTooltip>
            </div>
          </template>

          <template #empty>
            <div
              class="grid min-h-56 place-items-center px-6 py-12 text-center"
            >
              <div>
                <span
                  class="mx-auto grid size-10 place-items-center rounded-full bg-elevated text-muted"
                >
                  <UIcon
                    :name="
                      collections.length
                        ? 'i-tabler-file-search'
                        : 'i-tabler-stack-2'
                    "
                    class="size-5"
                  />
                </span>
                <p class="mt-3 text-sm font-medium text-highlighted">
                  {{ collections.length ? "没有匹配的文档" : "还没有文档集" }}
                </p>
                <p class="mt-1 text-xs text-muted">
                  {{
                    collections.length
                      ? "请调整搜索或筛选条件后重试。"
                      : "请先创建文档集，再添加文档。"
                  }}
                </p>
              </div>
            </div>
          </template>
        </UTable>

        <footer
          class="flex flex-col gap-3 border-t border-default bg-muted/20 px-3 py-3 text-xs sm:flex-row sm:items-center sm:justify-between sm:px-4"
        >
          <p class="text-muted">
            <template v-if="selectedDocIds.length">
              已选择 {{ selectedDocIds.length }} 篇文档
            </template>
            <template v-else>
              显示 {{ firstVisibleDoc }}–{{ lastVisibleDoc }}，共
              {{ docCollection.total }} 篇
            </template>
          </p>
          <div class="flex flex-wrap items-center gap-2">
            <UPagination
              :page="page"
              :total="docCollection.total"
              :items-per-page="pageSize"
              :show-edges="false"
              :sibling-count="1"
              size="xs"
              @update:page="page = $event"
            />
            <span class="text-muted">每页</span>
            <USelect
              :model-value="pageSize"
              :items="pageSizeItems"
              value-key="value"
              size="xs"
              class="w-24"
              aria-label="每页文档数量"
              @update:model-value="setPageSize"
            />
          </div>
        </footer>
      </section>

      <div
        v-if="viewMode === 'list' && bulkResult"
        class="mt-3 flex min-w-0 flex-wrap items-center gap-2 rounded-lg border border-default bg-elevated px-3 py-2.5"
        role="status"
      >
        <UIcon
          :name="
            bulkResult.interrupted || bulkResult.failedIds.length
              ? 'i-tabler-alert-triangle'
              : 'i-tabler-circle-check'
          "
          :class="
            bulkResult.interrupted || bulkResult.failedIds.length
              ? 'text-warning'
              : 'text-success'
          "
        />
        <span class="min-w-0 flex-1 text-xs text-default">
          <template v-if="bulkResult.interrupted">{{
            bulkResult.message
          }}</template>
          <template v-else
            >已处理 {{ bulkResult.changed }} 篇<span
              v-if="bulkResult.failedIds.length"
              >，{{ bulkResult.failedIds.length }} 篇未完成</span
            ></template
          >
        </span>
        <UButton
          v-if="bulkResult.failedIds[0]"
          label="查看首个失败项"
          color="warning"
          variant="link"
          size="xs"
          @click="openDoc(bulkResult.failedIds[0])"
        />
        <UButton
          icon="i-tabler-x"
          color="neutral"
          variant="ghost"
          size="xs"
          square
          aria-label="关闭批量结果"
          @click="bulkResult = undefined"
        />
      </div>

      <template v-if="viewMode === 'tree'">
        <div
          class="mb-4 flex flex-wrap items-center justify-between gap-3 rounded-lg border border-default bg-elevated/35 px-4 py-3"
        >
          <div class="min-w-0">
            <p class="text-sm font-medium text-highlighted">树状结构</p>
            <p class="text-xs text-muted">
              {{ tree?.tree?.length ?? 0 }} 个根节点 ·
              {{ activeTreeSlug || "未选择文档集" }}
            </p>
          </div>
          <div class="flex items-center gap-2">
            <USelectMenu
              :model-value="treeSlug"
              :items="treeCollectionItems"
              value-key="value"
              placeholder="选择文档集"
              :search-input="{ placeholder: '搜索文档集…' }"
              class="w-56"
              @update:model-value="treeSlug = selectedValue($event)"
            />
            <CollectionViewToggle
              v-model="viewMode"
              :items="[
                { key: 'list', label: '列表', icon: 'i-tabler-list' },
                { key: 'tree', label: '树状', icon: 'i-tabler-sitemap' },
              ]"
            />
          </div>
        </div>

        <SkeletonList v-if="treePending" :rows="8" />
        <ManageEmpty
          v-else-if="!tree?.tree?.length"
          icon="i-tabler-sitemap"
          text="这个文档集还没有树状结构"
        />
        <div v-else class="rounded-lg border border-default bg-default p-3">
          <DocTreeAdmin
            :nodes="tree.tree"
            :collection-slug="activeTreeSlug"
            @edit="openDoc"
            @add-child="
              (parentId) =>
                navigateTo(
                  `/manage/docs/new?collection=${activeTreeSlug}&parent=${parentId}`,
                )
            "
            @delete="onDelete"
            @move="onMove"
          />
        </div>
      </template>
    </ClientOnly>

    <ManageDocQuickEditModal
      v-model:open="showQuickEdit"
      :doc="quickEditTarget"
      :docs="quickEditDocs"
      @saved="onQuickEditSaved"
      @open-full="() => quickEditTarget && openDoc(quickEditTarget)"
    />
  </YAdminPage>
</template>
