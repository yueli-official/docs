<script setup lang="ts">
import { PageHeader } from "@yueli/ui/dashboard/pattern";
import { createPlatformNotifier } from "@platform/ui/feedback";
import {
  createCollectionRouteQueryCodec,
  createJsonCollectionQueryPolicy,
  type CollectionControl,
  type CollectionControlValue,
  type CollectionPanelMessages,
  type CollectionPanelState,
  type CollectionWorkflow,
} from "@yueli/ui/collection";
import { useVueCollectionWorkflow } from "@yueli/ui/collection/vue";
import { createVueRouterCollectionQuerySync } from "@yueli/ui/collection/vue-router";
import {
  CollectionPanel,
  CollectionViewToggle,
} from "@yueli/ui/collection/pattern";
import { ManageEmpty, SkeletonList } from "@platform/manage/components";
import type {
  CollectionManageTree,
  CollectionVersion,
  CollectionVersionsResponse,
  CollectionView,
  DocDetail,
} from "~/types";
import {
  buildDocTree,
  docManageRoute,
  findDocSlugPathById,
} from "~/utils/docsManageRoutes.mjs";

definePageMeta({ layout: "manage" });
useSeoMeta({ title: "文档 · 控制台" });

const { call } = useApi();
const toast = createPlatformNotifier(useToast());
const router = useRouter();

type StatusKey = "all" | "draft" | "published" | "archived" | "issues";
type BulkDocAction = "publish" | "draft" | "archive";

interface ManagedDoc extends DocDetail {
  collectionTitle: string;
  collectionSlug: string;
  parentTitle: string;
  path: string;
  slugPath: string[];
  depth: number;
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
}
const statusKeys = ["all", "draft", "published", "archived", "issues"] as const;
const viewKeys = ["list", "tree"] as const;
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
const showQuickEdit = ref(false);

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
  const matching = filterManagedDocs(nextQuery);
  const lastPage = Math.max(1, Math.ceil(matching.length / nextQuery.size));
  if (nextQuery.page > lastPage) {
    activeWorkflow.setQuery({ ...nextQuery, page: lastPage });
    return;
  }
  const start = (nextQuery.page - 1) * nextQuery.size;
  activeWorkflow.resolveLoad(token, {
    items: matching.slice(start, start + nextQuery.size),
    total: matching.length,
  });
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
    updateCollectionQuery({ collection: value, version: "" }),
});
const activeVersion = computed({
  get: () => collectionQuery.value.version,
  set: (value: string) => updateCollectionQuery({ version: value }),
});
const activeLocale = computed({
  get: () => collectionQuery.value.locale,
  set: (value: string) => updateCollectionQuery({ locale: value || "en" }),
});
const activeStatus = computed<StatusKey>({
  get: () => collectionQuery.value.status,
  set: (value) => updateCollectionQuery({ status: value }),
});
const page = computed({
  get: () => collectionQuery.value.page,
  set: (value: number) => updateCollectionQuery({ page: value }, false),
});
const pageSize = computed({
  get: () => collectionQuery.value.size,
  set: (value: number) => updateCollectionQuery({ size: value }),
});
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

const docsByCollection = ref<Record<string, DocDetail[]>>({});
const docsPending = ref(false);
const docsError = ref("");
let loadSeq = 0;

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

async function loadDocs() {
  const items = collections.value;
  if (!items.length) {
    docsByCollection.value = {};
    return;
  }

  const seq = ++loadSeq;
  docsPending.value = true;
  docsError.value = "";
  try {
    const entries = await Promise.all(
      items.map(async (col) => {
        const version =
          activeCollection.value !== "all" && col.id === activeCollection.value
            ? activeVersion.value
            : "";
        const res = await call<{ items: DocDetail[] }>("/api/v1/docs", {
          query: {
            collectionId: col.id,
            locale: activeLocale.value,
            ...(version ? { version } : {}),
          },
        });
        return [col.id, res.items] as const;
      }),
    );
    if (seq === loadSeq) docsByCollection.value = Object.fromEntries(entries);
  } catch (err: any) {
    docsError.value = err?.data?.message || "加载文档失败";
  } finally {
    if (seq === loadSeq) docsPending.value = false;
  }
}

watch(
  collections,
  () => {
    loadVersions();
  },
  { immediate: true },
);
watch(
  [collections, activeLocale, activeVersion, activeCollection],
  () => {
    loadDocs();
  },
  { immediate: true },
);

function enrichDocs(col: CollectionView, docs: DocDetail[]): ManagedDoc[] {
  const byId = new Map(docs.map((d) => [d.id, d]));

  function pathFor(doc: DocDetail) {
    const parts = [doc.title];
    let cursor = byId.get(doc.parentId);
    let guard = 0;
    while (cursor && guard < 20) {
      parts.unshift(cursor.title);
      cursor = byId.get(cursor.parentId);
      guard++;
    }
    return parts.join(" / ");
  }

  function slugPathFor(doc: DocDetail) {
    const parts = [doc.slug];
    let cursor = byId.get(doc.parentId);
    let guard = 0;
    while (cursor && guard < 20) {
      parts.unshift(cursor.slug);
      cursor = byId.get(cursor.parentId);
      guard++;
    }
    return parts.filter(Boolean);
  }

  function depthFor(doc: DocDetail) {
    let depth = 0;
    let cursor = byId.get(doc.parentId);
    let guard = 0;
    while (cursor && guard < 20) {
      depth++;
      cursor = byId.get(cursor.parentId);
      guard++;
    }
    return depth;
  }

  return docs.map((doc) => ({
    ...doc,
    collectionTitle: col.title,
    collectionSlug: col.slug,
    parentTitle: byId.get(doc.parentId)?.title ?? "",
    path: pathFor(doc),
    slugPath: slugPathFor(doc),
    depth: depthFor(doc),
  }));
}

const allDocs = computed(() =>
  collections.value.flatMap((col) =>
    enrichDocs(col, docsByCollection.value[col.id] ?? []),
  ),
);

const statusItems = computed(() => {
  return [
    { value: "all", label: `全部 · ${allDocs.value.length}` },
    {
      value: "draft",
      label: `草稿 · ${allDocs.value.filter((doc) => doc.status === "draft").length}`,
    },
    {
      value: "published",
      label: `已发布 · ${allDocs.value.filter((doc) => doc.status === "published").length}`,
    },
    {
      value: "archived",
      label: `归档 · ${allDocs.value.filter((doc) => doc.status === "archived").length}`,
    },
    {
      value: "issues",
      label: `待完善 · ${allDocs.value.filter((doc) => qualityIssues(doc).length > 0).length}`,
    },
  ];
});

function qualityIssues(doc: ManagedDoc) {
  const issues: string[] = [];
  if (!doc.excerpt?.trim()) issues.push("缺摘要");
  if (!doc.slugPath.length) issues.push("缺路径");
  if (!doc.collectionSlug) issues.push("缺文档集");
  return issues;
}

function filterManagedDocs(query: Readonly<DocCollectionQuery>) {
  const q = query.q.trim().toLowerCase();
  return allDocs.value.filter((doc) => {
    if (query.collection !== "all" && doc.collectionId !== query.collection)
      return false;
    if (query.status === "issues" && !qualityIssues(doc).length) return false;
    if (
      query.status !== "all" &&
      query.status !== "issues" &&
      doc.status !== query.status
    )
      return false;
    if (!q) return true;
    return [
      doc.title,
      doc.slug,
      doc.path,
      doc.collectionTitle,
      doc.excerpt,
    ].some((v) => (v || "").toLowerCase().includes(q));
  });
}

const pagedDocs = computed(() => docCollection.value.items);
const selectedDocIds = computed<readonly string[]>(() =>
  docCollection.value.selection.mode === "keys"
    ? docCollection.value.selection.keys
    : [],
);
const isPageSelected = computed(() => docCollection.value.isPageSelected);
const isPageIndeterminate = computed(
  () => docCollection.value.isPageIndeterminate,
);
function toggleDocSelection(id: string) {
  docWorkflow.toggleKey(id);
}
function togglePageSelection(selected: boolean) {
  docWorkflow.togglePage(selected);
}
function clearDocSelection() {
  docWorkflow.clearSelection();
}
function replaceSelection(ids: readonly string[]) {
  docWorkflow.clearSelection();
  for (const id of ids) docWorkflow.toggleKey(id);
}
const selectedDocs = computed(() =>
  allDocs.value.filter((doc) => selectedDocIds.value.includes(doc.id)),
);
const collectionControls = computed<CollectionControl[]>(() => [
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
    ].filter(Boolean).length,
);
const collectionMessages: CollectionPanelMessages = {
  searchPlaceholder: "搜索标题、路径、slug 或文档集…",
  searchAction: "搜索",
  filtersAction: "筛选",
  activeFilters: (count) => `筛选（${count}）`,
  clearFilters: "清除筛选",
  selectPage: "选择当前页文档",
  selectItem: (label) => `选择文档：${label}`,
  bulkRegion: "文档批量操作",
  selected: (count) => `已选择 ${count} 篇文档`,
  selectAllResults: "选择全部结果",
  clearSelection: "取消选择",
  emptyTitle: "没有匹配的文档",
  emptyDescription: "请调整搜索或筛选条件后重试。",
  errorTitle: "文档加载失败",
  retry: "重新加载",
  showing: (first, last, total) => `显示 ${first}–${last}，共 ${total} 篇`,
  pageSize: "每页",
  pageSizeControl: "每页文档数量",
  pageSizeOption: (value) => `${value} 篇`,
};
const collectionState = computed<CollectionPanelState>(() => {
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
}
function clearCollectionFilters() {
  updateCollectionQuery({
    status: "all",
    collection: "all",
    version: "",
    locale: "en",
  });
}
const docKey = (doc: ManagedDoc) => doc.id;
const docLabel = (doc: ManagedDoc) => doc.title;

watch([activeCollection, activeStatus, activeLocale, activeVersion], () => {
  bulkAction.value = undefined;
  bulkResult.value = undefined;
});
watch(activeCollection, (value) => {
  if (value === "all" && activeVersion.value) activeVersion.value = "";
});
watch(allDocs, () => {
  reloadDocPage();
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
    typeof docOrId === "string"
      ? allDocs.value.find((item) => item.id === docOrId)
      : docOrId;
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

function openQuickEdit(doc: ManagedDoc) {
  quickEditTarget.value = doc;
  showQuickEdit.value = true;
}

async function onQuickEditSaved() {
  await Promise.all([loadDocs(), refreshTree()]);
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
    await refreshTree();
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
    await refreshTree();
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
const treePending = computed(() => docsPending.value);
const tree = computed<CollectionManageTree | null>(
  () =>
    treeOverride.value ??
    (activeTreeCollection.value
      ? {
          collection: activeTreeCollection.value,
          tree: buildDocTree(
            docsByCollection.value[activeTreeCollection.value.id] ?? [],
          ),
        }
      : null),
);
watch(
  [activeTreeSlug, docsByCollection],
  () => {
    treeOverride.value = null;
  },
  { deep: true },
);

async function refreshTree() {
  treeOverride.value = null;
  await loadDocs();
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
    await loadDocs();
  } catch {
    treeOverride.value = snapshot;
    toast.add({ title: "移动失败，已还原", color: "error" });
    await refreshTree();
  }
}

async function onDelete(id: string) {
  try {
    await call(`/api/v1/docs/${id}`, { method: "DELETE" });
    await Promise.all([refreshTree(), loadDocs()]);
  } catch (err: any) {
    toast.add({
      title: "删除失败",
      description: err?.data?.message || "请重试",
      color: "error",
    });
  }
}
</script>

<template>
  <div>
    <PageHeader title="文档">
      <template #subtitle>
        <span>搜索、筛选、批量处理与层级调整</span>
      </template>
      <template #actions>
        <UButton icon="i-tabler-plus" label="新建文档" :to="createTarget" />
      </template>
    </PageHeader>

    <ClientOnly>
      <CollectionPanel
        v-if="viewMode === 'list'"
        v-model:search="search"
        :items="pagedDocs"
        :item-key="docKey"
        :item-label="docLabel"
        :controls="collectionControls"
        :messages="collectionMessages"
        :state="collectionState"
        :error-message="docsError"
        :total="docCollection.total"
        :page="page"
        :page-size="pageSize"
        :page-sizes="pageSizes"
        :active-filter-count="activeFilterCount"
        :selection-count="selectedDocIds.length"
        :page-selected="isPageSelected"
        :page-indeterminate="isPageIndeterminate"
        :is-selected="docWorkflow.isSelected"
        :is-item-selectable="() => !bulkBusy"
        :inert="bulkBusy"
        :aria-busy="bulkBusy"
        label="文档列表"
        selectable
        @search="submitSearch"
        @control-change="changeCollectionControl"
        @clear-filters="clearCollectionFilters"
        @retry="loadDocs"
        @toggle-page="togglePageSelection"
        @toggle-item="toggleDocSelection"
        @clear-selection="clearBulkSelection"
        @page-change="page = $event"
        @page-size-change="pageSize = $event"
      >
        <template #view>
          <CollectionViewToggle
            v-model="viewMode"
            :items="[
              { key: 'list', label: '列表', icon: 'i-tabler-list' },
              { key: 'tree', label: '树状', icon: 'i-tabler-sitemap' },
            ]"
          />
        </template>

        <template #columns>
          <div class="grid grid-cols-[minmax(0,1fr)_auto] items-center gap-3">
            <span>文档、路径与文档集</span>
            <span class="hidden w-40 text-right md:block"
              >语言、父级与操作</span
            >
          </div>
        </template>

        <template #empty>
          <div class="grid min-h-56 place-items-center px-6 py-12 text-center">
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

        <template #bulk-actions>
          <USelect
            v-model="bulkAction"
            :items="bulkItems"
            value-key="value"
            placeholder="批量操作"
            size="xs"
            class="w-28"
          />
          <UButton
            label="应用"
            icon="i-tabler-check"
            color="primary"
            variant="soft"
            size="xs"
            :disabled="!bulkAction"
            :loading="bulkBusy"
            @click="applyBulkAction"
          />
        </template>

        <template #item="{ item: doc }">
          <div
            class="grid min-w-0 gap-3 sm:grid-cols-[2.5rem_minmax(0,1fr)] md:grid-cols-[2.5rem_minmax(0,1fr)_10rem_auto] md:items-center"
          >
            <span
              class="hidden size-10 place-items-center rounded-lg bg-primary/10 text-primary sm:grid"
            >
              <UIcon name="i-tabler-file-text" class="size-5" />
            </span>
            <div class="min-w-0">
              <button
                type="button"
                class="block max-w-full truncate text-left text-sm font-semibold text-highlighted hover:text-primary"
                @click="openQuickEdit(doc)"
              >
                {{ doc.title }}
              </button>
              <p class="mt-0.5 truncate font-mono text-xs text-muted">
                {{ doc.slugPath.join(" / ") }}
              </p>
              <p class="mt-1 truncate text-xs text-dimmed">
                {{ doc.collectionTitle }} · {{ doc.parentTitle || "顶级文档" }}
              </p>
              <p
                v-if="qualityIssues(doc).length"
                class="mt-1 inline-flex max-w-full items-center gap-1 truncate text-xs text-warning"
              >
                <UIcon name="i-tabler-alert-circle" class="size-3.5 shrink-0" />
                <span class="truncate">{{
                  qualityIssues(doc).slice(0, 2).join(" · ")
                }}</span>
              </p>
            </div>
            <div class="min-w-0 text-xs md:text-right">
              <p class="truncate text-default">{{ doc.locale }}</p>
              <p class="mt-0.5 truncate text-muted">
                {{ doc.parentTitle || "顶级文档" }}
              </p>
            </div>
            <div class="flex justify-end gap-1">
              <UTooltip text="添加子文档"
                ><UButton
                  icon="i-tabler-file-plus"
                  color="neutral"
                  variant="ghost"
                  size="xs"
                  square
                  :aria-label="`添加子文档：${doc.title}`"
                  @click="addChild(doc)"
              /></UTooltip>
              <UTooltip text="快速编辑"
                ><UButton
                  icon="i-tabler-pencil"
                  color="neutral"
                  variant="ghost"
                  size="xs"
                  square
                  :aria-label="`快速编辑：${doc.title}`"
                  @click="openQuickEdit(doc)"
              /></UTooltip>
              <UTooltip text="完整编辑"
                ><UButton
                  icon="i-tabler-file-pencil"
                  color="neutral"
                  variant="ghost"
                  size="xs"
                  square
                  :aria-label="`完整编辑：${doc.title}`"
                  @click="openDoc(doc)"
              /></UTooltip>
            </div>
          </div>
        </template>
      </CollectionPanel>

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
      :docs="allDocs"
      @saved="onQuickEditSaved"
      @open-full="() => quickEditTarget && openDoc(quickEditTarget)"
    />
  </div>
</template>
