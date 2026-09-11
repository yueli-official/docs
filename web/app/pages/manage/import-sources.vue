<script setup lang="ts">
import { CollectionPaginationBar } from '@yueli/ui/collection/pattern';
import { createDocsNotifier } from '~/utils/feedback';
import type { CollectionView } from '~/types';

definePageMeta({ layout: 'manage' });
useSeoMeta({ title: '项目同步 · 控制台' });
interface Source {
  id: string; collectionTitle: string; collectionSlug: string; repository: string;
  assetName: string; defaultLocale: string; mode: string; enabled: boolean; autoCheck: boolean;
  status: string; lastTag: string; lastError: string; checkedAt: string | null;
}
interface Run { id: string; status: string; stage: string; tag: string; batchId: string; errorMessage: string; startedAt: string }
const { call } = useApi();
const { can } = useMe();
const toast = createDocsNotifier(useToast());
const items = ref<Source[]>([]), total = ref(0), page = ref(1), size = ref(20);
const available = ref(true), loading = ref(false), saving = ref(false), formOpen = ref(false);
const collections = ref<CollectionView[]>([]);
const selected = ref<Source | null>(null), runs = ref<Run[]>([]), detailOpen = ref(false);
const replacement = ref('');
const pageError = ref('');
const busy = ref<string | null>(null);
function report(error: unknown) { toast.add({ title: '操作未完成', description: docsFailureMessage(error, '请稍后重试'), color: 'error' }); }
async function action(run: () => Promise<void>) { try { await run(); } catch (error) { report(error); } }
const form = reactive({ repository: '', assetName: 'docs.zip', collection: '', defaultLocale: 'zh-CN', mode: 'upsert', personalToken: '', exclusiveMirror: false, autoCheck: true });
const modes = [{ label: '更新与新增，保留其他页面', value: 'upsert' }, { label: '完整镜像，下架已移除的页面', value: 'replace-version' }];
const labels: Record<string,string> = { idle: '尚未检查', queued: '等待检查', running: '同步中', completed: '已更新', skipped: '已是最新', paused: '已暂停', failed: '同步失败' };
const canManage = computed(() => can('docs.import.manage'));
const collectionOptions = computed(() => collections.value.map(c => ({ label: c.semanticVersion ? `${c.title} · v${c.semanticVersion}` : c.title, value: c.slug })));
function date(value: string | null) { return value ? new Date(value).toLocaleString() : '尚未检查'; }
function errorText(value: string) {
  const messages: Record<string,string> = {
    'documentation ZIP is missing from the latest release': '最新正式 Release 中没有找到文档包，请上传配置的 ZIP 附件。',
    'no eligible published release': '没有可用的正式 Release。',
    'GitHub request failed': '暂时无法连接 GitHub，请稍后检查。',
    'release asset digest mismatch': '文档包摘要不一致，本次未导入。',
    'release asset size changed': '文档包大小发生变化，本次未导入。',
    'GitHub content exceeds size limit': '文档包超过大小限制。',
  };
  return messages[value] || value;
}
async function refresh() {
  if (loading.value || !canManage.value) return;
  loading.value = true;
  try {
    const data = await call<{items:Source[];total:number;enabled:boolean}>(`/api/v1/import-sources?page=${page.value}&size=${size.value}`);
    items.value = data.items; total.value = data.total; available.value = data.enabled;
    pageError.value = '';
  } catch (error) { pageError.value = docsFailureMessage(error, '项目列表读取失败');
  } finally { loading.value = false; }
}
async function openForm() {
  await action(async () => {
  const data = await call<{items:CollectionView[]}>('/api/v1/collections');
  collections.value = data.items; formOpen.value = true;
  });
}
async function create() {
  saving.value = true;
  try {
    await call('/api/v1/import-sources', { method: 'POST', body: { ...form } });
    form.personalToken = ''; formOpen.value = false;
    toast.add({ title: form.autoCheck ? '已添加项目，后台开始检查文档包' : '已添加项目，可手动检查文档包', color: 'success' }); await refresh();
  } catch (error) { report(error); } finally { saving.value = false; }
}
async function check(source: Source) {
  busy.value = source.id;
  try { await action(async () => { await call(`/api/v1/import-sources/${source.id}/check`, { method: 'POST' }); await refresh(); }); }
  finally { busy.value = null; }
}
async function toggle(source: Source) {
  busy.value = source.id;
  try { await action(async () => { await call(`/api/v1/import-sources/${source.id}`, { method: 'PATCH', body: { enabled: !source.enabled } }); await refresh(); }); }
  finally { busy.value = null; }
}
async function setAutoCheck(source: Source, autoCheck: boolean) {
  busy.value = source.id;
  try { await action(async () => { await call(`/api/v1/import-sources/${source.id}`, { method: 'PATCH', body: { autoCheck } }); await refresh(); }); }
  finally { busy.value = null; }
}
async function detail(source: Source) {
  await action(async () => {
  const data = await call<{source:Source;runs:Run[]}>(`/api/v1/import-sources/${source.id}`);
  selected.value = data.source; runs.value = data.runs; replacement.value = ''; detailOpen.value = true;
  });
}
async function replaceToken() {
  if (!selected.value) return;
  saving.value = true;
  try {
    await call(`/api/v1/import-sources/${selected.value.id}`, { method: 'PATCH', body: { enabled: true, personalToken: replacement.value } });
    replacement.value = ''; detailOpen.value = false; await refresh();
  } catch (error) { report(error); } finally { saving.value = false; }
}
async function removeSource() {
  if (!selected.value) return;
  const id = selected.value.id;
  await action(async () => { await call(`/api/v1/import-sources/${id}`, { method: 'DELETE' }); detailOpen.value = false; await refresh(); });
}
watch(formOpen, value => { if (!value) form.personalToken = ''; });
watch(detailOpen, value => { if (!value) replacement.value = ''; });
watch([page, size], () => { void refresh(); });
let timer: ReturnType<typeof setInterval> | undefined;
onMounted(async () => { await refresh(); timer = setInterval(() => { if (!document.hidden) void refresh(); }, 10000); });
onBeforeUnmount(() => { if (timer) clearInterval(timer); });
</script>

<template>
  <ManagePage id="project-docs" title="项目同步" icon="i-tabler-brand-github">
    <template #actions>
      <UButton v-if="canManage" icon="i-tabler-plus" :disabled="!available" @click="openForm">添加项目</UButton>
    </template>
    <ManageImportNavigation>
    <UAlert v-if="!canManage" title="没有文档导入权限" color="neutral" variant="subtle" />
    <UAlert v-else-if="!available" title="项目同步尚未启用" description="站点配置好同步凭证存储后，即可绑定 GitHub Release 文档包。" color="neutral" variant="subtle" />
    <template v-else>
      <UAlert v-if="pageError" :title="pageError" color="error" variant="subtle" class="mb-4" />
      <div class="overflow-hidden">
        <div v-if="!items.length" class="p-10 text-center text-sm text-muted">{{ loading ? '正在读取项目…' : '还没有绑定项目。在源仓库维护文档，发布后自动更新到这里。' }}</div>
        <article v-for="source in items" :key="source.id" class="flex flex-col gap-3 border-b border-default p-4 last:border-b-0 sm:flex-row sm:items-start sm:justify-between">
          <div class="min-w-0 flex-1 space-y-1">
            <div class="flex flex-wrap items-center gap-2">
              <button class="font-medium text-highlighted hover:underline" @click="detail(source)">{{ source.collectionTitle }}</button>
              <UBadge :color="source.status === 'failed' ? 'error' : 'neutral'" variant="subtle" size="sm">{{ labels[source.status] || source.status }}</UBadge>
            </div>
            <a :href="`https://github.com/${source.repository}`" target="_blank" rel="noopener noreferrer" class="block truncate text-sm text-muted hover:text-highlighted">{{ source.repository }}</a>
            <p class="text-xs text-muted">{{ source.assetName }} · {{ source.lastTag || '等待首次导入' }} · {{ date(source.checkedAt) }}</p>
            <p v-if="source.lastError" class="text-sm text-error">{{ errorText(source.lastError) }}</p>
          </div>
          <div class="flex shrink-0 flex-col gap-3 sm:items-end">
            <div class="flex flex-wrap items-center gap-2">
              <UButton color="neutral" variant="outline" size="sm" icon="i-tabler-refresh" :loading="busy === source.id" :disabled="!source.enabled || source.status === 'running'" @click="check(source)">检查</UButton>
              <UButton color="neutral" variant="outline" size="sm" :icon="source.enabled ? 'i-tabler-player-pause' : 'i-tabler-player-play'" :disabled="busy === source.id" @click="toggle(source)">{{ source.enabled ? '暂停' : '恢复' }}</UButton>
              <UButton color="neutral" variant="outline" size="sm" icon="i-tabler-history" @click="detail(source)">记录</UButton>
            </div>
            <USwitch :model-value="source.autoCheck" :disabled="busy === source.id" label="自动检查" size="sm" :aria-label="`${source.collectionTitle}：自动检查`" @update:model-value="setAutoCheck(source, $event)" />
          </div>
        </article>
        <CollectionPaginationBar :page="page" :page-size="size" :total="total" :page-sizes="[10,20,50]" :page-size-option="value => `${value} 个`" class="border-t border-default p-3" @page-change="page=$event" @page-size-change="size=$event;page=1" />
      </div>
    </template>
    <UModal v-model:open="formOpen" title="添加 GitHub 项目" description="使用公开仓库正式 Release 中的文档包。目标文档集的版本在文档集设置中管理。">
      <template #body>
        <form id="create-source" class="space-y-4" @submit.prevent="create">
          <UFormField label="GitHub 仓库" required><UInput v-model="form.repository" placeholder="https://github.com/owner/project" required class="w-full" /></UFormField>
          <div class="grid gap-4 sm:grid-cols-2">
            <UFormField label="附件名称" required><UInput v-model="form.assetName" required class="w-full" /></UFormField>
            <UFormField label="默认语言" required><UInput v-model="form.defaultLocale" required class="w-full" /></UFormField>
          </div>
          <UFormField label="目标文档集" required><USelect v-model="form.collection" :items="collectionOptions" placeholder="选择文档集" class="w-full" /></UFormField>
          <UFormField label="更新方式"><USelect v-model="form.mode" :items="modes" class="w-full" /></UFormField>
          <USwitch v-model="form.autoCheck" label="自动检查" description="每 15 分钟检查更新；关闭后由你手动触发。" />
          <UCheckbox v-if="form.mode==='replace-version'" v-model="form.exclusiveMirror" label="这个文档集只由该项目维护，允许下架包内已移除的页面。" />
          <UFormField label="开发者令牌" required description="在账户中心创建并勾选「月离文档 → 导入文档」。撤销令牌或失去权限后，同步会暂停。">
            <UInput v-model="form.personalToken" type="password" autocomplete="off" required class="w-full" />
          </UFormField>
        </form>
      </template>
      <template #footer><div class="flex w-full justify-end gap-2"><UButton color="neutral" variant="outline" @click="() => { formOpen = false; }">取消</UButton><UButton type="submit" form="create-source" :loading="saving" :disabled="!form.collection || (form.mode==='replace-version'&&!form.exclusiveMirror)">添加项目</UButton></div></template>
    </UModal>
    <USlideover v-model:open="detailOpen" :title="selected?.collectionTitle || '同步记录'">
      <template #body>
        <div class="space-y-6">
          <p class="text-sm text-muted">{{ selected?.repository }} · {{ selected?.assetName }}</p>
          <p v-if="!runs.length" class="text-sm text-muted">暂无同步记录，检查文档包后会显示在这里。</p>
          <ol class="divide-y divide-default">
            <li v-for="run in runs" :key="run.id" class="space-y-1 py-3">
              <div class="flex justify-between gap-2 text-sm"><span>{{ labels[run.status] || run.status }}</span><span class="text-muted">{{ run.tag }}</span></div>
              <p class="text-xs text-muted">{{ date(run.startedAt) }}</p><p v-if="run.errorMessage" class="text-sm text-error">{{ errorText(run.errorMessage) }}</p>
              <NuxtLink v-if="run.batchId" :to="`/manage/import/${run.batchId}`" class="text-sm text-primary">查看导入报告</NuxtLink>
            </li>
          </ol>
          <form class="space-y-3 border-t border-default pt-5" @submit.prevent="replaceToken">
            <UFormField label="更换同步令牌"><UInput v-model="replacement" type="password" autocomplete="off" class="w-full" /></UFormField>
            <UButton type="submit" :disabled="!replacement" :loading="saving" color="neutral" variant="outline">更换并恢复同步</UButton>
          </form>
          <div class="border-t border-default pt-4"><UButton color="error" variant="ghost" @click="removeSource">解除绑定</UButton><p class="mt-1 text-xs text-muted">保留已导入的文档，停止后续同步。</p></div>
        </div>
      </template>
    </USlideover>
    </ManageImportNavigation>
  </ManagePage>
</template>
