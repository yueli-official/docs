<script setup lang="ts">
import { createCollectionRouteQueryCodec } from "@yueli/ui/collection";
import { useVueRouterCollectionQuery } from "@yueli/ui/collection/vue-router";
import {
  CollectionDock,
  CollectionPagination,
  CollectionSortDirectionButton,
  CollectionTableToolbar,
} from "@yueli/ui/collection/pattern";
import { AssetImageCropper } from "@yueli/asset-nuxt/components";
import { assetUploadURL } from "@yueli/asset-nuxt/upload";
import { createDocsNotifier } from "~/utils/feedback";
import { AdminRowActions } from "@yueli/ui/admin";
import type { AdminRowActionItem } from "@yueli/ui/admin";
import {
  ManageEmpty,
  ManageVisualAssetField,
  SkeletonList,
} from "~/utils/manageComponents";
import { useMinimumLoading } from "@yueli/ui/feedback";
import type { CollectionView } from "~/types";

definePageMeta({ layout: "manage" });
useSeoMeta({ title: "文档集 · 控制台" });

const { call } = useApi();
const { can } = useMe();
const canManageCollections = computed(() => can("docs.collection.manage"));
const toast = createDocsNotifier(useToast());
const route = useRoute();
const router = useRouter();
type CollectionSort = "title" | "docCount";
type Direction = "asc" | "desc";
interface CollectionQuery {
  q: string;
  sort: CollectionSort;
  direction: Direction;
  page: number;
  size: number;
}
const pageSizes = [12, 24, 48] as const;
const { query, replace: replaceQuery } =
  useVueRouterCollectionQuery<CollectionQuery>({
    router,
    codec: createCollectionRouteQueryCodec({
      q: { kind: "string", default: "" },
      sort: { kind: "enum", values: ["title", "docCount"], default: "title" },
      direction: { kind: "enum", values: ["asc", "desc"], default: "asc" },
      page: { kind: "positive-integer", default: 1 },
      size: { kind: "positive-integer", values: pageSizes, default: 24 },
    }),
  });
function updateQuery(patch: Partial<CollectionQuery>, resetPage = true) {
  void replaceQuery({
    ...query.value,
    ...patch,
    ...(resetPage ? { page: 1 } : {}),
  } as CollectionQuery);
}
const q = computed(() => query.value.q);
const sort = computed({
  get: () => query.value.sort,
  set: (value: CollectionSort) => updateQuery({ sort: value }),
});
const direction = computed({
  get: () => query.value.direction,
  set: (value: Direction) => updateQuery({ direction: value }),
});
const page = computed({
  get: () => query.value.page,
  set: (value: number) => updateQuery({ page: value }, false),
});
const size = computed({
  get: () => query.value.size,
  set: (value: number) => updateQuery({ size: value }),
});
const searchInput = ref(q.value);
let searchTimer: ReturnType<typeof setTimeout> | undefined;
watch(searchInput, (value) => {
  if (searchTimer) clearTimeout(searchTimer);
  searchTimer = setTimeout(() => updateQuery({ q: value.trim() }), 300);
});
watch(q, (value) => {
  if (searchInput.value !== value) searchInput.value = value;
});
onScopeDispose(() => {
  if (searchTimer) clearTimeout(searchTimer);
});
const mounted = ref(false);
onMounted(() => {
  mounted.value = true;
});

const { data, pending, error, refresh } = await useAsyncData(
  "manage-collections",
  () => call<{ items: CollectionView[] }>("/api/v1/collections"),
  { server: false, default: () => ({ items: [] as CollectionView[] }) },
);

const items = computed(() => data.value?.items ?? []);
const loadError = computed(
  () =>
    error.value as {
      statusCode?: number;
      status?: number;
      data?: { message?: string };
      message?: string;
    } | null,
);
const isAuthError = computed(() => {
  const err = loadError.value;
  return (err?.statusCode ?? err?.status) === 401;
});
const loadErrorMessage = computed(() =>
  isAuthError.value
    ? "登录已失效，请重新登录后继续管理文档集。"
    : loadError.value?.data?.message ||
      loadError.value?.message ||
      "加载文档集失败，请稍后重试。",
);
const showSkeleton = useMinimumLoading(
  computed(() => !loadError.value && (!mounted.value || pending.value)),
);
const { login } = useAuth();

function loginAgain() {
  return login(route.fullPath);
}

const filteredItems = computed(() => {
  const keyword = q.value.trim().toLowerCase();
  const filtered = keyword
    ? items.value.filter((item) =>
        [item.title, item.slug, item.description].some((value) =>
          (value || "").toLowerCase().includes(keyword),
        ),
      )
    : [...items.value];
  const multiplier = direction.value === "asc" ? 1 : -1;
  return filtered.sort((a, b) => {
    if (sort.value === "docCount")
      return ((a.docCount || 0) - (b.docCount || 0)) * multiplier;
    return a.title.localeCompare(b.title, "zh-CN") * multiplier;
  });
});
const totalPages = computed(() =>
  Math.max(1, Math.ceil(filteredItems.value.length / size.value)),
);
const pagedItems = computed(() =>
  filteredItems.value.slice(
    (page.value - 1) * size.value,
    page.value * size.value,
  ),
);

watch(totalPages, (n) => {
  if (page.value > n) page.value = n;
});
const sortItems = [
  { label: "按标题", value: "title" },
  { label: "按文档数", value: "docCount" },
];
const pageSizeItems = [12, 24, 48].map((value) => ({
  label: `${value}/页`,
  value,
}));

const open = ref(false);
const current = ref<CollectionView | null>(null);
async function onReleaseChanged() {
  const currentID = current.value?.id;
  await refresh();
  if (currentID) {
    current.value = (data.value?.items ?? []).find((item) => item.id === currentID) ?? current.value;
  }
}
type CollectionEditorSection = "basic" | "languages" | "versions";
const editorSection = ref<CollectionEditorSection>("basic");
const editorTabs = computed(() => [
  { label: "基础", value: "basic", icon: "i-tabler-adjustments-horizontal" },
  { label: "语言", value: "languages", icon: "i-tabler-language", disabled: !current.value },
  { label: "版本", value: "versions", icon: "i-tabler-versions", disabled: !current.value },
]);
const form = reactive({
  title: "",
  slug: "",
  description: "",
  icon: "",
  cover: "",
  defaultLocale: "zh-CN",
  semanticVersion: "1.0.0",
});
const localeItems = [
  { label: "简体中文", value: "zh-CN" },
  { label: "English", value: "en-US" },
  { label: "繁體中文", value: "zh-TW" },
  { label: "日本語", value: "ja-JP" },
];
const slugTouched = ref(false);
const saving = ref(false);
const coverPct = ref(-1);
const coverCropOpen = ref(false);
const coverCropFile = ref<File | null>(null);
const pendingCoverFile = ref<File | null>(null);
const pendingCoverPreview = ref("");
const confirmingDelete = ref(false);
const deletingBusy = ref(false);
const formError = ref("");

const normalizedFormSlug = computed(() => clientSlug(form.slug));
const canSave = computed(() =>
  Boolean(
    form.title.trim() &&
      normalizedFormSlug.value &&
      (current.value || /^\d+\.\d+\.\d+$/.test(form.semanticVersion.trim())),
  ),
);

function clientSlug(value: string) {
  return value
    .trim()
    .toLowerCase()
    .replace(/[^\p{L}\p{N}]+/gu, "-")
    .replace(/^-+|-+$/g, "");
}

watch(
  () => form.title,
  (title) => {
    if (!current.value && !slugTouched.value) form.slug = clientSlug(title);
  },
);

function collectionPublicPath(col: CollectionView) {
  return col.slug ? `/${col.slug}` : "";
}

function openCreate() {
  formError.value = "";
  clearPendingCover();
  current.value = null;
  form.title = "";
  form.slug = "";
  form.description = "";
  form.icon = "";
  form.cover = "";
  form.defaultLocale = "zh-CN";
  form.semanticVersion = "1.0.0";
  slugTouched.value = false;
  confirmingDelete.value = false;
  editorSection.value = "basic";
  open.value = true;
}

function openEdit(col: CollectionView) {
  formError.value = "";
  clearPendingCover();
  current.value = col;
  form.title = col.title;
  form.slug = col.slug || "";
  form.description = col.description || "";
  form.icon = col.icon || "";
  form.cover = col.coverUrl || "";
  slugTouched.value = true;
  confirmingDelete.value = false;
  editorSection.value = "basic";
  open.value = true;
}

function collectionRowActions(col: CollectionView): AdminRowActionItem[] {
  return [
    {
      id: "view",
      label: `查看文档集：${col.title}`,
      icon: "i-tabler-external-link",
      to: collectionPublicPath(col),
      target: "_blank",
      rel: "noopener",
    },
    {
      id: "edit",
      label: `编辑文档集：${col.title}`,
      icon: "i-tabler-pencil",
      onSelect: () => openEdit(col),
    },
  ];
}

function clearPendingCover() {
  if (pendingCoverPreview.value) URL.revokeObjectURL(pendingCoverPreview.value);
  pendingCoverPreview.value = "";
  pendingCoverFile.value = null;
  coverCropFile.value = null;
  coverCropOpen.value = false;
  coverPct.value = -1;
}

function setCoverUrl(value: string) {
  clearPendingCover();
  form.cover = value;
}

function onPickCover(file: File) {
  if (!file.type.startsWith("image/")) {
    formError.value = "请选择图片文件";
    return;
  }
  if (file.size > 10 * 1024 * 1024) {
    formError.value = "图片不能超过 10MB";
    return;
  }
  formError.value = "";
  if (pendingCoverPreview.value) URL.revokeObjectURL(pendingCoverPreview.value);
  pendingCoverPreview.value = "";
  pendingCoverFile.value = null;
  coverPct.value = -1;
  coverCropFile.value = file;
  coverCropOpen.value = true;
}

function onCroppedCover(payload: { file: File }) {
  const { file } = payload;
  if (pendingCoverPreview.value) URL.revokeObjectURL(pendingCoverPreview.value);
  pendingCoverFile.value = file;
  pendingCoverPreview.value = URL.createObjectURL(file);
  coverCropFile.value = null;
}

function putWithProgress(
  url: string,
  file: File,
  onProgress?: (pct: number) => void,
  headers?: Record<string, string>,
): Promise<void> {
  return new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest();
    xhr.open("PUT", assetUploadURL(url));
    for (const [key, value] of Object.entries(headers ?? {}))
      xhr.setRequestHeader(key, value);
    xhr.upload.onprogress = (event) => {
      if (event.lengthComputable && onProgress)
        onProgress(Math.round((event.loaded / event.total) * 100));
    };
    xhr.onload = () =>
      xhr.status >= 200 && xhr.status < 300
        ? resolve()
        : reject(new Error(`上传失败 (HTTP ${xhr.status})`));
    xhr.onerror = () => reject(new Error("上传网络错误，请确认素材服务在线"));
    xhr.send(file);
  });
}

async function uploadCollectionCover(collectionId: string, file: File) {
  coverPct.value = 0;
  const init = await call<{
    uploadUrl: string;
    uploadToken: string;
    uploadHeaders?: Record<string, string>;
  }>(`/api/v1/collections/${collectionId}/cover`, {
    method: "POST",
    body: { filename: file.name, mime: file.type, size: file.size },
  });
  await putWithProgress(
    init.uploadUrl,
    file,
    (pct) => {
      coverPct.value = pct;
    },
    init.uploadHeaders,
  );
  const res = await call<{ collection: CollectionView; coverUrl: string }>(
    `/api/v1/collections/${collectionId}/cover/finalize`,
    { method: "POST", body: { uploadToken: init.uploadToken } },
  );
  return res;
}

async function save() {
  formError.value = "";
  if (!form.title.trim()) {
    formError.value = "请填写标题";
    return;
  }
  if (!normalizedFormSlug.value) {
    formError.value = "请填写路径标识";
    return;
  }
  if (!current.value && !/^\d+\.\d+\.\d+$/.test(form.semanticVersion.trim())) {
    formError.value = "版本号必须使用 major.minor.patch，例如 1.0.0";
    return;
  }

  saving.value = true;
  try {
    const file = pendingCoverFile.value;
    const body = {
      title: form.title.trim(),
      slug: form.slug.trim(),
      description: form.description,
      icon: form.icon,
      cover: form.cover,
      ...(!current.value
        ? {
            defaultLocale: form.defaultLocale,
            semanticVersion: form.semanticVersion.trim(),
          }
        : {}),
    };
    const res = current.value
      ? await call<{ collection: CollectionView }>(
          `/api/v1/collections/${current.value.id}`,
          { method: "PATCH", body },
        )
      : await call<{ collection: CollectionView }>("/api/v1/collections", {
          method: "POST",
          body,
        });

    let saved = res.collection;
    if (file && saved?.id) {
      const coverRes = await uploadCollectionCover(saved.id, file);
      saved = coverRes.collection;
      form.cover = coverRes.coverUrl;
      clearPendingCover();
    }

    open.value = false;
    await refresh();
  } catch (e: any) {
    toast.add({
      title: current.value ? "更新失败" : "创建失败",
      description: e?.data?.message || "请重试",
      color: "error",
    });
  } finally {
    saving.value = false;
    coverPct.value = -1;
  }
}

async function doDelete() {
  if (!current.value) return;
  deletingBusy.value = true;
  try {
    await call(`/api/v1/collections/${current.value.id}`, { method: "DELETE" });
    open.value = false;
    await refresh();
  } catch (e: any) {
    toast.add({
      title: "删除失败",
      description: e?.data?.message || "请重试",
      color: "error",
    });
  } finally {
    deletingBusy.value = false;
  }
}
</script>

<template>
  <ManagePage
    id="collections"
    title="文档集"
    icon="i-tabler-stack-2"
    main-id="manage-main"
    body-class="w-full"
  >
    <template #actions>
      <UButton
        v-if="canManageCollections"
        icon="i-tabler-plus"
        label="新建文档集"
        @click="openCreate"
      />
    </template>

    <div
      v-if="!canManageCollections"
      class="rounded-lg border border-default bg-default p-8"
    >
      <div class="mx-auto max-w-md text-center">
        <span
          class="mx-auto grid size-12 place-items-center rounded-lg bg-warning/10 text-warning"
        >
          <UIcon name="i-tabler-lock" class="size-6" />
        </span>
        <h2 class="mt-4 text-lg font-semibold text-highlighted">
          没有文档集管理权限
        </h2>
        <p class="mt-2 text-sm leading-6 text-muted">
          当前角色不能创建、修改或删除文档集。请联系管理员调整角色能力。
        </p>
      </div>
    </div>

    <div
      v-else-if="loadError"
      class="rounded-lg border border-default bg-default p-8"
    >
      <div class="mx-auto max-w-md text-center">
        <span
          class="mx-auto grid size-12 place-items-center rounded-lg bg-warning/10 text-warning"
        >
          <UIcon
            :name="isAuthError ? 'i-tabler-lock' : 'i-tabler-alert-triangle'"
            class="size-6"
          />
        </span>
        <h2 class="mt-4 text-lg font-semibold text-highlighted">
          {{ isAuthError ? "需要重新登录" : "加载失败" }}
        </h2>
        <p class="mt-2 text-sm leading-6 text-muted">{{ loadErrorMessage }}</p>
        <div class="mt-5 flex justify-center gap-2">
          <UButton
            v-if="isAuthError"
            icon="i-tabler-login-2"
            label="重新登录"
            @click="
              () => {
                loginAgain();
              }
            "
          />
          <UButton
            v-else
            icon="i-tabler-refresh"
            label="重试"
            color="neutral"
            variant="soft"
            @click="
              () => {
                refresh();
              }
            "
          />
        </div>
      </div>
    </div>

    <SkeletonList v-else-if="showSkeleton" :rows="8" />

    <template v-else>
      <section class="overflow-hidden rounded-xl border border-default bg-default">
        <CollectionTableToolbar
          v-model:search="searchInput"
          label="文档集工具栏"
          search-placeholder="搜索标题、路径标识或说明…"
          filter-label="筛选"
        >
          <template #utilities>
            <USelectMenu
              v-model="sort"
              :items="sortItems"
              value-key="value"
              icon="i-tabler-arrows-sort"
              aria-label="排序方式"
              size="sm"
              class="w-36"
            />
            <CollectionSortDirectionButton v-model="direction" />
          </template>
        </CollectionTableToolbar>

        <ManageEmpty
          v-if="!items.length"
          icon="i-tabler-stack-2"
          text="还没有文档集"
        />
        <ManageEmpty
          v-else-if="!filteredItems.length"
          icon="i-tabler-search-off"
          text="没有匹配的文档集"
        />

        <template v-else>
        <div
          class="hidden grid-cols-[minmax(16rem,1.4fr)_minmax(10rem,.8fr)_7rem_8rem] items-center gap-3 border-b border-default bg-elevated/45 px-4 py-2.5 text-xs font-medium text-muted lg:grid"
        >
          <span>文档集</span>
          <span>公开路径</span>
          <span>文档数</span>
          <span class="sr-only">操作</span>
        </div>

        <div class="divide-y divide-default">
          <div
            v-for="col in pagedItems"
            :key="col.id"
            class="grid w-full grid-cols-[minmax(0,1fr)_8rem] items-center gap-3 p-3 text-left transition hover:bg-elevated/55 sm:p-4 lg:grid-cols-[minmax(16rem,1.4fr)_minmax(10rem,.8fr)_7rem_8rem]"
          >
            <button
              type="button"
              class="flex min-w-0 items-center gap-3 rounded-lg text-left focus-visible:outline-2 focus-visible:outline-offset-4 focus-visible:outline-primary"
              :aria-label="`编辑文档集：${col.title}`"
              @click="openEdit(col)"
            >
              <img
                v-if="col.coverUrl"
                :src="col.coverUrl"
                :alt="col.title"
                class="size-11 shrink-0 rounded-lg object-cover"
              />
              <span
                v-else
                class="grid size-11 shrink-0 place-items-center rounded-lg bg-primary/10 text-primary"
              >
                <UIcon :name="col.icon || 'i-tabler-stack-2'" class="size-5" />
              </span>
              <span class="min-w-0">
                <span class="flex min-w-0 items-center gap-2">
                  <span class="line-clamp-1 text-sm font-semibold text-highlighted">{{ col.title }}</span>
                  <UBadge v-if="col.semanticVersion" :label="col.semanticVersion" color="neutral" variant="soft" size="xs" />
                </span>
                <span class="mt-1 line-clamp-1 text-xs text-muted">{{
                  col.description || "未填写集合说明"
                }}</span>
              </span>
            </button>

            <span
              class="col-start-1 min-w-0 pl-14 font-mono text-xs text-muted sm:text-sm lg:col-start-auto lg:pl-0"
            >
              <span class="line-clamp-1">{{
                collectionPublicPath(col) || "-"
              }}</span>
            </span>

            <span
              class="col-start-1 pl-14 text-xs text-muted sm:text-sm lg:col-start-auto lg:pl-0"
            >
              <span class="font-semibold text-highlighted">{{
                col.docCount
              }}</span>
              篇文档
            </span>

            <AdminRowActions
              class="row-start-1 col-start-2 lg:row-auto lg:col-start-auto"
              :label="`${col.title} 的操作`"
              :items="collectionRowActions(col)"
            />
          </div>
        </div>
        </template>
      </section>

      <CollectionDock v-if="pagedItems.length" label="文档集统计与分页">
        <template #selection>
          <span>共 {{ filteredItems.length }} 个文档集</span>
          <span v-if="q" class="text-xs text-muted"
            >全部 {{ items.length }} 个</span
          >
        </template>
        <template #pagination>
          <USelect
            v-model="size"
            :items="pageSizeItems"
            value-key="value"
            size="sm"
            class="w-24"
          />
          <CollectionPagination v-model="page" :total-pages="totalPages" />
        </template>
      </CollectionDock>
    </template>

    <USlideover
      v-model:open="open"
      :title="current ? '编辑文档集' : '新建文档集'"
    >
      <template #body>
        <div class="space-y-5">
          <UAlert
            v-if="formError"
            color="warning"
            variant="subtle"
            icon="i-tabler-alert-triangle"
            title="请完善文档集信息"
            :description="formError"
            role="alert"
          />
          <UTabs
            v-model="editorSection"
            :items="editorTabs"
            :content="false"
            value-key="value"
            variant="pill"
            color="neutral"
            class="w-full"
            :ui="{
              list: 'w-full rounded-xl bg-elevated/70 p-1',
              indicator: 'rounded-lg bg-default ring-1 ring-default shadow-xs',
              trigger: 'min-h-9 flex-1 justify-center gap-2 rounded-lg data-[state=active]:text-highlighted',
              leadingIcon: 'size-4.5 shrink-0',
            }"
            aria-label="文档集设置分区"
            data-collection-settings-tabs
          />
          <div v-if="editorSection === 'basic'" class="space-y-4">
            <UFormField label="标题" required>
              <UInput
                v-model="form.title"
                placeholder="文档集标题"
                class="w-full"
                autofocus
              />
            </UFormField>

            <UFormField
              label="路径标识"
              required
              help="集合公开路径；保存后该集合下文档路径会同步使用新标识。"
            >
              <UInput
                v-model="form.slug"
                icon="i-tabler-link"
                placeholder="quickstart"
                class="w-full"
                @input="slugTouched = true"
              />
            </UFormField>

            <UFormField label="描述">
              <UTextarea v-model="form.description" :rows="3" class="w-full" />
            </UFormField>

            <div v-if="!current" class="grid gap-4 sm:grid-cols-2">
              <UFormField
                label="默认语言"
                required
                help="创建后仍可添加其他语言；默认语言必须始终保留。"
              >
                <USelect
                  v-model="form.defaultLocale"
                  :items="localeItems"
                  value-key="value"
                  class="w-full"
                />
              </UFormField>
              <UFormField
                label="版本号"
                required
                help="使用 major.minor.patch；创建后不可修改。"
              >
                <UInput
                  v-model="form.semanticVersion"
                  inputmode="decimal"
                  placeholder="1.0.0"
                  class="w-full"
                />
              </UFormField>
            </div>

            <ManageVisualAssetField
              :icon="form.icon"
              :cover-url="form.cover"
              :preview-url="pendingCoverPreview"
              :progress="coverPct"
              :uploading="coverPct >= 0"
              @update:icon="form.icon = $event"
              @update:cover-url="setCoverUrl"
              @pick-cover="onPickCover"
              @clear-cover="form.cover = ''"
            />
          </div>
          <ManageCollectionVariantsManager
            v-else-if="current"
            :collection="current"
            :section="editorSection === 'languages' ? 'languages' : 'versions'"
            @changed="onReleaseChanged"
          />
        </div>
      </template>

      <template #footer>
        <div class="flex w-full items-center justify-between gap-4">
          <UPopover
            v-if="current"
            v-model:open="confirmingDelete"
            :content="{ side: 'top', align: 'start', sideOffset: 10 }"
            :ui="{ content: 'w-80 p-4' }"
          >
            <UButton
              label="删除文档集"
              icon="i-tabler-trash"
              color="neutral"
              variant="ghost"
              class="text-muted hover:text-error"
            />
            <template #content>
              <p class="text-sm font-semibold text-highlighted">
                删除「{{ current.title }}」？
              </p>
              <p class="mt-1 text-xs leading-5 text-muted">
                该文档集及其全部文档会被永久删除。
              </p>
              <div class="mt-4 flex justify-end gap-2">
                <UButton
                  label="取消"
                  color="neutral"
                  variant="ghost"
                  size="sm"
                  @click="
                    () => {
                      confirmingDelete = false;
                    }
                  "
                />
                <UButton
                  label="确认删除"
                  icon="i-tabler-trash"
                  color="error"
                  size="sm"
                  :loading="deletingBusy"
                  @click="doDelete"
                />
              </div>
            </template>
          </UPopover>
          <span v-else aria-hidden="true" />

          <div class="flex shrink-0 justify-end gap-2">
            <UButton
              label="取消"
              color="neutral"
              variant="outline"
              @click="
                () => {
                  open = false;
                }
              "
            />
            <UButton
              :label="current ? '保存' : '创建'"
              icon="i-tabler-check"
              :loading="saving"
              :disabled="!canSave"
              @click="save"
            />
          </div>
        </div>
      </template>
    </USlideover>
    <AssetImageCropper
      v-model:open="coverCropOpen"
      :file="coverCropFile"
      title="裁剪文档集封面"
      :aspect-ratio="1"
      :output-width="256"
      :output-height="256"
      output-type="image/webp"
      :quality="0.85"
      @cropped="onCroppedCover"
    >
      <template #controls="{ sourceSize, outputSize }">
        <section class="rounded-xl border border-default bg-default p-4">
          <h3 class="text-sm font-semibold text-highlighted">输出规格</h3>
          <dl class="mt-3 divide-y divide-default text-sm">
            <div class="flex items-center justify-between gap-4 py-2">
              <dt class="text-muted">使用位置</dt>
              <dd class="text-right text-default">首页与文档集列表</dd>
            </div>
            <div class="flex items-center justify-between gap-4 py-2">
              <dt class="text-muted">裁剪比例</dt>
              <dd class="tabular-nums text-default">1:1</dd>
            </div>
            <div class="flex items-center justify-between gap-4 py-2">
              <dt class="text-muted">输出尺寸</dt>
              <dd class="tabular-nums text-default">
                {{ outputSize?.width || 256 }} × {{ outputSize?.height || 256 }} px
              </dd>
            </div>
            <div class="flex items-center justify-between gap-4 py-2">
              <dt class="text-muted">输出格式</dt>
              <dd class="text-default">WebP</dd>
            </div>
            <div
              v-if="sourceSize"
              class="flex items-center justify-between gap-4 pt-2"
            >
              <dt class="text-muted">原图尺寸</dt>
              <dd class="tabular-nums text-default">
                {{ sourceSize.width }} × {{ sourceSize.height }} px
              </dd>
            </div>
          </dl>
        </section>
      </template>
    </AssetImageCropper>
  </ManagePage>
</template>
