<script setup lang="ts">
import { platformDashboardMessages } from '@platform/manage/dashboard'
import { SkeletonList } from '@platform/manage/components'
import { useMinimumLoading } from '@yueli/ui/feedback'
import { DashboardLayout } from '@yueli/ui/dashboard/pattern'
import type { CollectionList, CollectionView } from '~/types'

definePageMeta({ layout: 'manage' })
useSeoMeta({ title: '控制台' })

const { call } = useApi()
const mounted = ref(false)
onMounted(() => { mounted.value = true })

const { data: collections, pending, error } = await useAsyncData(
  'manage-overview-collections',
  () => call<CollectionList>('/api/v1/collections'),
  { server: false, default: () => ({ items: [] }) },
)

const items = computed(() => collections.value?.items ?? [])
const showSkeleton = useMinimumLoading(computed(() => !mounted.value || pending.value))

function collectionIssues(collection: CollectionView) {
  const issues: string[] = []
  if (!collection.description?.trim()) issues.push('缺少说明')
  if (!collection.docCount) issues.push('还没有文档')
  if (!collection.icon && !collection.coverUrl) issues.push('缺少识别图标或封面')
  return issues
}

const totalDocs = computed(() => items.value.reduce((sum, collection) => sum + (collection.docCount || 0), 0))
const attention = computed(() => items.value.filter(collection => collectionIssues(collection).length > 0))
const emptyCollections = computed(() => items.value.filter(collection => !collection.docCount).length)
const metrics = computed(() => [
  { label: '文档集', value: items.value.length, icon: 'i-tabler-stack-2', to: '/manage/collections' },
  { label: '文档总数', value: totalDocs.value, icon: 'i-tabler-files', to: '/manage/docs' },
  { label: '空文档集', value: emptyCollections.value, icon: 'i-tabler-file-off', to: '/manage/collections' },
  { label: '待完善', value: attention.value.length, icon: 'i-tabler-alert-circle', to: '/manage/collections' },
])
const continueCollections = computed(() => [...items.value].sort((a, b) => (b.docCount || 0) - (a.docCount || 0)).slice(0, 6))
</script>

<template>
  <DashboardLayout
    title="控制台"
    description="查看文档运营队列，继续内容工作并确认当前站点可用。"
    :messages="platformDashboardMessages"
    recent-title="继续工作"
    recent-description="按内容量列出常用文档集；进入后继续维护文档。"
  >
    <template #actions>
      <UButton to="/manage/docs" icon="i-tabler-file-text" label="管理文档" />
    </template>

    <template #metrics>
      <div v-if="showSkeleton" class="grid grid-cols-2 gap-3 lg:grid-cols-4">
        <USkeleton v-for="item in 4" :key="item" class="h-24 rounded-xl" />
      </div>
      <div v-else class="grid grid-cols-2 gap-3 lg:grid-cols-4">
        <NuxtLink v-for="metric in metrics" :key="metric.label" :to="metric.to" class="rounded-xl border border-default bg-default p-4 transition hover:border-primary/40 hover:bg-elevated/30">
          <div class="flex items-center gap-3">
            <span class="grid size-10 shrink-0 place-items-center rounded-lg bg-primary/10 text-primary"><UIcon :name="metric.icon" class="size-5" /></span>
            <div class="min-w-0"><p class="text-xl font-semibold text-highlighted tabular-nums sm:text-2xl">{{ metric.value }}</p><p class="truncate text-xs text-muted">{{ metric.label }}</p></div>
          </div>
        </NuxtLink>
      </div>
    </template>

    <template #pending>
      <div v-if="showSkeleton" class="grid gap-2"><USkeleton v-for="item in 3" :key="item" class="h-14 rounded-lg" /></div>
      <div v-else-if="attention.length" class="divide-y divide-default">
        <NuxtLink
          v-for="collection in attention.slice(0, 5)"
          :key="collection.id"
          :to="{ path: '/manage/docs', query: { collection: collection.id } }"
          class="group flex min-h-14 items-center gap-3 py-2.5 first:pt-0 last:pb-0"
        >
          <span class="grid size-9 shrink-0 place-items-center rounded-lg bg-warning/10 text-warning"><UIcon name="i-tabler-alert-circle" class="size-4" /></span>
          <div class="min-w-0 flex-1"><p class="truncate text-sm font-medium text-highlighted">{{ collection.title }}</p><p class="truncate text-xs text-muted">{{ collectionIssues(collection).join(' · ') }}</p></div>
          <UIcon name="i-tabler-chevron-right" class="size-4 shrink-0 text-dimmed transition group-hover:translate-x-0.5" />
        </NuxtLink>
      </div>
      <UAlert v-else color="success" variant="subtle" icon="i-tabler-circle-check" title="当前没有待完善的文档集" />
    </template>

    <template #recent>
      <SkeletonList v-if="showSkeleton" :rows="5" class="p-4" />
      <div v-else-if="continueCollections.length" class="divide-y divide-default">
        <NuxtLink
          v-for="collection in continueCollections"
          :key="collection.id"
          :to="{ path: '/manage/docs', query: { collection: collection.id } }"
          class="group grid min-h-14 grid-cols-[auto_minmax(0,1fr)_auto] items-center gap-3 px-4 py-3 transition hover:bg-elevated/50"
        >
          <span class="grid size-9 place-items-center overflow-hidden rounded-lg bg-primary/10 text-primary">
            <img v-if="collection.coverUrl" :src="collection.coverUrl" :alt="collection.title" class="size-full object-cover" />
            <UIcon v-else :name="collection.icon || 'i-tabler-stack-2'" class="size-4" />
          </span>
          <div class="min-w-0"><p class="truncate text-sm font-medium text-highlighted group-hover:text-primary">{{ collection.title }}</p><p class="truncate text-xs text-muted">{{ collection.description || '未填写文档集说明' }}</p></div>
          <span class="text-xs text-muted">{{ collection.docCount }} 篇</span>
        </NuxtLink>
      </div>
      <div v-else class="p-8 text-center text-sm text-muted">还没有文档集，先创建内容容器。</div>
    </template>

    <template #health>
      <UAlert v-if="error" color="error" variant="subtle" icon="i-tabler-alert-circle" title="文档服务暂时不可用" description="刷新后仍失败时，请到平台状态检查服务。" />
      <div v-else class="space-y-3">
        <div class="flex items-center justify-between gap-3 rounded-lg bg-success/10 px-3 py-2.5 text-sm"><span class="flex items-center gap-2 text-success"><UIcon name="i-tabler-circle-check" class="size-4" />文档服务可用</span><span class="text-xs text-muted">正常</span></div>
        <p class="text-xs leading-5 text-muted">内容缺口统一进入待完善队列，不展示常驻完整度面板或正常状态噪声。</p>
      </div>
    </template>

    <template #quickActions>
      <div class="grid gap-2">
        <UButton to="/manage/docs" icon="i-tabler-file-text" label="管理文档" color="neutral" variant="soft" block />
        <UButton to="/manage/collections" icon="i-tabler-stack-2" label="管理文档集" color="neutral" variant="soft" block />
        <UButton to="/manage/import" icon="i-tabler-file-import" label="批量导入" color="neutral" variant="soft" block />
        <UButton to="/manage/home" icon="i-tabler-settings" label="站点设置" color="neutral" variant="ghost" block />
      </div>
    </template>
  </DashboardLayout>
</template>
