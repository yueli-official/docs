<script setup lang="ts">
import { createDocsNotifier } from "~/utils/feedback";
import type {
  CollectionReleasesResponse,
  CollectionLocale,
  CollectionLocalesResponse,
  CollectionView,
} from "~/types";

const props = defineProps<{
  collection: CollectionView | null;
  section: "languages" | "versions";
}>();
const emit = defineEmits<{ changed: [collection: CollectionView] }>();
const { call } = useApi();
const toast = createDocsNotifier(useToast());
const loading = ref(false);
const locales = ref<CollectionLocale[]>([]);
const releases = ref<CollectionView[]>([]);
const localeEditorOpen = ref(false);
const localeCloneOpen = ref(false);
const releaseEditorOpen = ref(false);
const releaseInitializeOpen = ref(false);
const editingLocale = ref(false);
const saving = ref(false);
const deletingLocale = ref(false);
const releaseInitializeVersion = ref("1.0.0");

const localePresets = [
  { label: "简体中文", locale: "zh-CN", htmlLang: "zh-CN", direction: "ltr" as const },
  { label: "繁體中文", locale: "zh-TW", htmlLang: "zh-TW", direction: "ltr" as const },
  { label: "English", locale: "en", htmlLang: "en", direction: "ltr" as const },
  { label: "日本語", locale: "ja", htmlLang: "ja", direction: "ltr" as const },
  { label: "한국어", locale: "ko", htmlLang: "ko", direction: "ltr" as const },
];
const localePresetItems = localePresets.map((item) => ({
  label: `${item.label} · ${item.locale}`,
  value: item.locale,
}));
const directionItems = [
  { label: "从左到右", value: "ltr" },
  { label: "从右到左", value: "rtl" },
];
const localeForm = reactive({
  locale: "",
  label: "",
  htmlLang: "",
  direction: "ltr" as "ltr" | "rtl",
  isDefault: false,
  enabled: true,
  sortOrder: 0,
});
const releaseForm = reactive({
  sourceSemanticVersion: "1.0.0",
  targetSemanticVersion: "2.0.0",
  title: "",
  slug: "",
});
const localeCloneForm = reactive({
  sourceLocale: "",
  targetLocale: "zh-CN",
});
const sourceLocaleItems = computed(() => locales.value
  .filter((item) => item.enabled && item.docCount > 0)
  .map((item) => ({ label: `${item.label} · ${item.docCount} 篇`, value: item.locale })));
const targetLocaleItems = computed(() => {
  const values = new Map<string, { label: string; value: string; disabled?: boolean }>();
  for (const preset of localePresets) {
    const existing = locales.value.find((item) => item.locale === preset.locale);
    values.set(preset.locale, {
      label: existing ? `${existing.label}${existing.docCount ? ` · 已有 ${existing.docCount} 篇` : ""}` : preset.label,
      value: preset.locale,
      disabled: Boolean(existing?.docCount),
    });
  }
  for (const item of locales.value) {
    if (!values.has(item.locale)) {
      values.set(item.locale, {
        label: `${item.label}${item.docCount ? ` · 已有 ${item.docCount} 篇` : ""}`,
        value: item.locale,
        disabled: Boolean(item.docCount),
      });
    }
  }
  return [...values.values()].filter((item) => item.value !== localeCloneForm.sourceLocale);
});

watch(
  () => [props.collection?.id, props.section] as const,
  ([value]) => {
    if (value) void load();
  },
  { immediate: true },
);

async function load() {
  if (!props.collection) return;
  loading.value = true;
  try {
    if (props.section === "languages") {
      const response = await call<CollectionLocalesResponse>(`/api/v1/manage/collections/${props.collection.id}/locales`);
      locales.value = response.items;
    } else {
      const response = await call<CollectionReleasesResponse>(`/api/v1/collections/${props.collection.slug}/releases`);
      releases.value = response.items;
    }
  } finally {
    loading.value = false;
  }
}

function resetLocale() {
  Object.assign(localeForm, {
    locale: "",
    label: "",
    htmlLang: "",
    direction: "ltr",
    isDefault: false,
    enabled: true,
    sortOrder: locales.value.length * 10,
  });
  editingLocale.value = false;
}

function addLocale() {
  resetLocale();
  localeEditorOpen.value = true;
}

function cloneLocale() {
  const source = locales.value.find((item) => item.isDefault && item.docCount > 0)
    ?? locales.value.find((item) => item.docCount > 0);
  if (!source) return;
  localeCloneForm.sourceLocale = source.locale;
  const preferredTargets = ["zh-CN", "en", "zh-TW", "ja", "ko"];
  localeCloneForm.targetLocale = preferredTargets.find((locale) =>
    locale !== source.locale && !(locales.value.find((item) => item.locale === locale)?.docCount),
  ) ?? "";
  localeCloneOpen.value = true;
}

async function submitLocaleClone() {
  if (!props.collection || !localeCloneForm.sourceLocale || !localeCloneForm.targetLocale) return;
  const preset = localePresets.find((item) => item.locale === localeCloneForm.targetLocale);
  const existing = locales.value.find((item) => item.locale === localeCloneForm.targetLocale);
  if (!preset && !existing) return;
  saving.value = true;
  try {
    await call(`/api/v1/manage/collections/${props.collection.id}/locales/clone`, {
      method: "POST",
      body: {
        sourceLocale: localeCloneForm.sourceLocale,
        targetLocale: localeCloneForm.targetLocale,
        targetLabel: existing?.label || preset?.label,
        targetHtmlLang: existing?.htmlLang || preset?.htmlLang,
        targetDirection: existing?.direction || preset?.direction || "ltr",
        targetSortOrder: existing?.sortOrder ?? locales.value.length * 10,
      },
    });
    localeCloneOpen.value = false;
    await load();
  } catch (error: any) {
    toast.add({ title: "语言复制失败", description: error?.data?.message || "目标语言已有文档，或源语言没有可复制内容", color: "error" });
  } finally {
    saving.value = false;
  }
}

function editLocale(value: CollectionLocale) {
  Object.assign(localeForm, {
    locale: value.locale,
    label: value.label,
    htmlLang: value.htmlLang,
    direction: value.direction,
    isDefault: value.isDefault,
    enabled: value.enabled,
    sortOrder: value.sortOrder,
  });
  editingLocale.value = true;
  localeEditorOpen.value = true;
}

function applyLocalePreset(value: string) {
  const preset = localePresets.find((item) => item.locale === value);
  if (!preset) return;
  Object.assign(localeForm, preset);
}

async function saveLocale() {
  if (!props.collection || !localeForm.locale.trim() || !localeForm.label.trim()) return;
  saving.value = true;
  try {
    await call(`/api/v1/manage/collections/${props.collection.id}/locales`, {
      method: "POST",
      body: localeForm,
    });
    localeEditorOpen.value = false;
    await load();
  } catch (error: any) {
    toast.add({ title: "语言设置保存失败", description: error?.data?.message || "请检查设置后重试", color: "error" });
  } finally {
    saving.value = false;
  }
}

async function deleteLocale() {
  if (!props.collection || !editingLocale.value) return;
  deletingLocale.value = true;
  try {
    await call(`/api/v1/manage/collections/${props.collection.id}/locales/${encodeURIComponent(localeForm.locale)}`, { method: "DELETE" });
    localeEditorOpen.value = false;
    await load();
  } catch (error: any) {
    toast.add({ title: "无法删除语言", description: error?.data?.message || "默认语言或已有文档的语言不能删除", color: "error" });
  } finally {
    deletingLocale.value = false;
  }
}

function nextReleaseVersion(value: string | undefined) {
  const match = /^(\d+)\.(\d+)\.(\d+)$/.exec(value || "");
  return match ? `${Number(match[1]) + 1}.0.0` : "2.0.0";
}

function addRelease() {
  if (!props.collection) return;
  const target = nextReleaseVersion(props.collection.semanticVersion);
  Object.assign(releaseForm, {
    sourceSemanticVersion: props.collection.semanticVersion || "1.0.0",
    targetSemanticVersion: target,
    title: `${props.collection.title} ${target}`,
    slug: `${props.collection.slug}-${target.replaceAll(".", "-")}`,
  });
  releaseEditorOpen.value = true;
}

async function cloneRelease() {
  if (!props.collection || !releaseForm.targetSemanticVersion.trim() || !releaseForm.title.trim()) return;
  saving.value = true;
  try {
    const response = await call<{ collection: CollectionView }>(`/api/v1/manage/collections/${props.collection.id}/clone-release`, {
      method: "POST",
      body: releaseForm,
    });
    releaseEditorOpen.value = false;
    await load();
    emit("changed", response.collection);
  } catch (error: any) {
    toast.add({ title: "新版本创建失败", description: error?.data?.message || "请检查版本号和路径后重试", color: "error" });
  } finally {
    saving.value = false;
  }
}

async function initializeRelease() {
  if (!props.collection || !releaseInitializeVersion.value.trim()) return;
  saving.value = true;
  try {
    const response = await call<{ collection: CollectionView }>(`/api/v1/manage/collections/${props.collection.id}/release`, {
      method: "POST",
      body: { semanticVersion: releaseInitializeVersion.value.trim() },
    });
    releaseInitializeOpen.value = false;
    await load();
    emit("changed", response.collection);
  } catch (error: any) {
    toast.add({ title: "版本设置失败", description: error?.data?.message || "版本号必须使用 x.y.z 格式", color: "error" });
  } finally {
    saving.value = false;
  }
}
</script>

<template>
  <div class="min-w-0" data-collection-variants-panel>
      <div v-if="loading" class="space-y-3">
        <USkeleton v-for="index in 4" :key="index" class="h-16 rounded-xl" />
      </div>

      <section v-else-if="section === 'languages'">
        <div class="mb-4 flex items-center justify-between gap-4">
          <h3 class="text-sm font-semibold text-highlighted">启用语言</h3>
          <div class="flex items-center gap-2">
            <UButton label="复制语言" icon="i-tabler-copy" size="sm" color="neutral" variant="outline" :disabled="!sourceLocaleItems.length" @click="cloneLocale" />
            <UButton label="添加语言" icon="i-tabler-plus" size="sm" color="neutral" variant="outline" @click="addLocale" />
          </div>
        </div>
        <div class="divide-y divide-default rounded-xl border border-default bg-default">
          <button
            v-for="item in locales"
            :key="item.locale"
            type="button"
            class="flex w-full items-center gap-3 px-4 py-3 text-left hover:bg-elevated"
            @click="editLocale(item)"
          >
            <span class="grid size-9 shrink-0 place-items-center rounded-lg bg-muted text-xs font-semibold text-toned">{{ item.locale.split('-')[0]?.toUpperCase() }}</span>
            <span class="min-w-0 flex-1">
              <span class="block truncate text-sm font-semibold text-highlighted">{{ item.label }}</span>
              <span class="block text-xs text-muted">{{ item.locale }} · {{ item.direction.toUpperCase() }} · {{ item.docCount }} 篇</span>
            </span>
            <UBadge v-if="item.isDefault" label="默认" color="primary" variant="subtle" />
            <UBadge v-else-if="!item.enabled" label="已停用" color="neutral" variant="subtle" />
            <UIcon name="i-tabler-chevron-right" class="size-4 text-dimmed" />
          </button>
        </div>
      </section>

      <section v-else>
        <div class="mb-4 flex items-center justify-between gap-4">
          <div>
            <h3 class="text-sm font-semibold text-highlighted">关联版本</h3>
            <p class="mt-1 text-xs text-muted">每个版本都是可独立维护的文档集。</p>
          </div>
          <div class="flex items-center gap-2">
            <UButton v-if="!collection?.semanticVersion" label="设置当前版本" icon="i-tabler-tag" size="sm" color="neutral" variant="outline" @click="() => { releaseInitializeOpen = true }" />
            <UButton label="克隆为新版本" icon="i-tabler-copy-plus" size="sm" color="neutral" variant="outline" @click="addRelease" />
          </div>
        </div>
        <div class="divide-y divide-default rounded-xl border border-default bg-default">
          <NuxtLink
            v-for="item in releases"
            :key="item.id"
            :to="`/${item.slug}`"
            target="_blank"
            class="flex w-full items-center gap-3 px-4 py-3 text-left hover:bg-elevated"
          >
            <span class="grid size-9 shrink-0 place-items-center rounded-lg bg-muted text-xs font-semibold text-toned">V</span>
            <span class="min-w-0 flex-1">
              <span class="block truncate text-sm font-semibold text-highlighted">{{ item.semanticVersion || '未设置版本' }}</span>
              <span class="block truncate text-xs text-muted">{{ item.title }} · /{{ item.slug }}</span>
            </span>
            <UBadge v-if="item.id === collection?.id" label="当前" color="primary" variant="subtle" />
            <UIcon name="i-tabler-external-link" class="size-4 text-dimmed" />
          </NuxtLink>
        </div>
      </section>

  <UModal v-model:open="localeEditorOpen" :title="editingLocale ? '编辑语言' : '添加语言'">
    <template #body>
      <div class="space-y-4">
        <UFormField v-if="!editingLocale" label="语言预设">
          <USelectMenu :items="localePresetItems" value-key="value" class="w-full" placeholder="选择常用语言" @update:model-value="applyLocalePreset(String($event || ''))" />
        </UFormField>
        <UFormField label="语言代码" required help="创建后保持稳定">
          <UInput v-model="localeForm.locale" :disabled="editingLocale" class="w-full" placeholder="zh-CN" />
        </UFormField>
        <UFormField label="显示名称" required><UInput v-model="localeForm.label" class="w-full" /></UFormField>
        <UFormField label="HTML lang"><UInput v-model="localeForm.htmlLang" class="w-full" /></UFormField>
        <UFormField label="文字方向"><USelect v-model="localeForm.direction" :items="directionItems" value-key="value" class="w-full" /></UFormField>
        <UFormField label="展示顺序"><UInputNumber v-model="localeForm.sortOrder" class="w-full" /></UFormField>
        <div class="flex items-center justify-between"><span class="text-sm text-highlighted">启用语言</span><USwitch v-model="localeForm.enabled" /></div>
        <div class="flex items-center justify-between"><span class="text-sm text-highlighted">设为默认语言</span><USwitch v-model="localeForm.isDefault" /></div>
      </div>
    </template>
    <template #footer>
      <div class="flex w-full items-center justify-between gap-4">
        <UButton v-if="editingLocale && !localeForm.isDefault" label="删除" color="neutral" variant="ghost" class="text-muted hover:text-error" :loading="deletingLocale" @click="deleteLocale" />
        <span v-else />
        <div class="flex gap-2"><UButton label="取消" color="neutral" variant="outline" @click="() => { localeEditorOpen = false }" /><UButton label="保存" :loading="saving" @click="saveLocale" /></div>
      </div>
    </template>
  </UModal>

  <UModal v-model:open="releaseEditorOpen" title="克隆为新版本">
    <template #body>
      <div class="space-y-4">
        <UFormField v-if="!collection?.semanticVersion" label="当前版本" required help="首次建立版本关系时，为现有文档集补充版本号。">
          <UInput v-model="releaseForm.sourceSemanticVersion" class="w-full" placeholder="1.0.0" />
        </UFormField>
        <UFormField label="新版本" required help="固定使用 major.minor.patch，创建后不可修改。">
          <UInput v-model="releaseForm.targetSemanticVersion" class="w-full" placeholder="2.0.0" />
        </UFormField>
        <UFormField label="文档集名称" required><UInput v-model="releaseForm.title" class="w-full" /></UFormField>
        <UFormField label="公开路径" required><UInput v-model="releaseForm.slug" class="w-full" /></UFormField>
        <div class="rounded-xl border border-default bg-muted/60 px-4 py-3 text-sm text-toned">
          文档、语言、层级和视觉资产会复制到新文档集；复制后的文档保持草稿，两个版本此后独立维护。
        </div>
      </div>
    </template>
    <template #footer>
      <div class="flex w-full justify-end gap-2"><UButton label="取消" color="neutral" variant="outline" @click="() => { releaseEditorOpen = false }" /><UButton label="创建版本" icon="i-tabler-copy-plus" :loading="saving" @click="cloneRelease" /></div>
    </template>
  </UModal>

  <UModal v-model:open="releaseInitializeOpen" title="设置当前版本">
    <template #body>
      <UFormField label="版本号" required help="固定使用 major.minor.patch，保存后不可修改。">
        <UInput v-model="releaseInitializeVersion" class="w-full" placeholder="1.0.0" />
      </UFormField>
    </template>
    <template #footer>
      <div class="flex w-full justify-end gap-2"><UButton label="取消" color="neutral" variant="outline" @click="() => { releaseInitializeOpen = false }" /><UButton label="保存版本" :loading="saving" @click="initializeRelease" /></div>
    </template>
  </UModal>

  <UModal v-model:open="localeCloneOpen" title="复制语言草稿">
    <template #body>
      <div class="space-y-4">
        <UFormField label="源语言" required>
          <USelect v-model="localeCloneForm.sourceLocale" :items="sourceLocaleItems" value-key="value" class="w-full" />
        </UFormField>
        <UFormField label="目标语言" required>
          <USelect v-model="localeCloneForm.targetLocale" :items="targetLocaleItems" value-key="value" class="w-full" />
        </UFormField>
        <div class="rounded-lg border border-default bg-muted/50 px-3 py-2.5 text-sm text-toned">
          复制后的文档全部为草稿，保留原有层级与关联关系。
        </div>
      </div>
    </template>
    <template #footer>
      <div class="flex w-full justify-end gap-2"><UButton label="取消" color="neutral" variant="outline" @click="() => { localeCloneOpen = false }" /><UButton label="复制为草稿" icon="i-tabler-copy" :loading="saving" @click="submitLocaleClone" /></div>
    </template>
  </UModal>
  </div>
</template>
