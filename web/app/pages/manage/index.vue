<script setup lang="ts">
import { ManageEmpty, ManageHeader, SkeletonList } from '@platform/manage/components'
import { useMinLoading } from '@platform/ui/use-min-loading'
import type { CollectionList, CollectionView } from '~/types'

definePageMeta({ layout: 'manage' })
useSeoMeta({ title: '控制台' })

const { call } = useApi()
const mounted = ref(false)
onMounted(() => { mounted.value = true })

const { data: collections, pending } = await useAsyncData(
  'manage-overview-collections',
  () => call<CollectionList>('/api/v1/collections'),
  { server: false, default: () => ({ items: [] }) },
)

const items = computed(() => collections.value?.items ?? [])
const showSkeleton = useMinLoading(computed(() => !mounted.value || pending.value))

function collectionIssues(col: CollectionView) {
  const issues: string[] = []
  if (!col.slug) issues.push('缺路径标识')
  if (!col.description?.trim()) issues.push('缺描述')
  if (!col.docCount) issues.push('空集合')
  if (!col.icon && !col.coverUrl) issues.push('缺视觉资产')
  return issues
}

const stats = computed(() => {
  const cols = items.value
  const totalDocs = cols.reduce((sum, col) => sum + (col.docCount || 0), 0)
  const ready = cols.filter(col => collectionIssues(col).length === 0).length
  const empty = cols.filter(col => !col.docCount).length
  const missingMeta = cols.filter(col => !col.slug || !col.description?.trim()).length
  const visualReady = cols.filter(col => col.icon || col.coverUrl).length
  return {
    collections: cols.length,
    totalDocs,
    ready,
    empty,
    missingMeta,
    visualReady,
  }
})

const opsCards = computed(() => [
  { label: '文档集资产', value: stats.value.collections, detail: `${stats.value.ready} 个就绪`, icon: 'i-tabler-stack-2', tone: 'primary' },
  { label: '文档总量', value: stats.value.totalDocs, detail: '公开站内容库存', icon: 'i-tabler-files', tone: 'success' },
  { label: '待补齐', value: stats.value.empty + stats.value.missingMeta, detail: '空集合或元信息缺口', icon: 'i-tabler-alert-triangle', tone: stats.value.empty + stats.value.missingMeta ? 'warning' : 'neutral' },
  { label: '视觉资产', value: stats.value.visualReady, detail: '图标或封面覆盖', icon: 'i-tabler-photo-check', tone: 'neutral' },
])

const readinessRows = computed(() => [
  { label: '路径标识', value: `${items.value.filter(col => col.slug).length}/${stats.value.collections}`, ok: items.value.every(col => col.slug) },
  { label: '集合说明', value: `${items.value.filter(col => col.description?.trim()).length}/${stats.value.collections}`, ok: items.value.every(col => col.description?.trim()) },
  { label: '内容容量', value: `${items.value.filter(col => col.docCount > 0).length}/${stats.value.collections}`, ok: stats.value.empty === 0 },
  { label: '视觉识别', value: `${stats.value.visualReady}/${stats.value.collections}`, ok: stats.value.visualReady === stats.value.collections },
])

const priorityCollections = computed(() =>
  [...items.value]
    .sort((a, b) => {
      const issueDelta = collectionIssues(b).length - collectionIssues(a).length
      if (issueDelta) return issueDelta
      return (b.docCount || 0) - (a.docCount || 0)
    })
    .slice(0, 5),
)

function collectionPublicPath(col: CollectionView) {
  return col.slug ? `/${col.slug}` : ''
}

function healthMeta(col: CollectionView) {
  const issues = collectionIssues(col)
  if (!issues.length) return { label: '就绪', color: 'success' as const, icon: 'i-tabler-shield-check' }
  if (issues.length > 2) return { label: `${issues.length} 项`, color: 'error' as const, icon: 'i-tabler-alert-triangle' }
  return { label: issues[0]!, color: 'warning' as const, icon: 'i-tabler-alert-circle' }
}

function statToneClass(tone: string) {
  if (tone === 'success') return 'bg-success/10 text-success ring-success/20'
  if (tone === 'warning') return 'bg-warning/10 text-warning ring-warning/20'
  if (tone === 'primary') return 'bg-primary/10 text-primary ring-primary/20'
  return 'bg-elevated text-muted ring-default'
}
</script>

<template>
  <div>
    <ManageHeader title="控制台">
      <template #subtitle>
        <span>文档站运营总览</span>
      </template>
    </ManageHeader>

    <SkeletonList v-if="showSkeleton" :rows="8" />

    <template v-else>
      <section class="mb-5 overflow-hidden rounded-lg border border-default bg-default">
        <div class="grid 2xl:grid-cols-[minmax(360px,0.85fr)_minmax(0,1.15fr)]">
          <div class="border-b border-default bg-elevated/35 p-5 2xl:border-b-0 2xl:border-r">
            <div class="flex items-center gap-2 text-xs font-semibold uppercase tracking-wide text-muted">
              <UIcon name="i-tabler-radar-2" class="size-4 text-primary" />
              文档运营
            </div>
            <div class="mt-3 flex flex-wrap items-end justify-between gap-3">
              <div class="min-w-0">
                <h2 class="text-2xl font-semibold text-highlighted">内容运营总览</h2>
                <p class="mt-1 text-sm text-muted">从集合、内容容量和发布准备度判断今天该处理什么。</p>
              </div>
              <UBadge
                :label="pending ? '同步中' : '实时数据'"
                :color="pending ? 'warning' : 'success'"
                :icon="pending ? 'i-tabler-loader-2' : 'i-tabler-broadcast'"
                variant="subtle"
              />
            </div>
          </div>

          <div class="grid sm:grid-cols-2 xl:grid-cols-4">
            <div
              v-for="card in opsCards"
              :key="card.label"
              class="border-b border-default p-4 sm:border-r xl:border-b-0"
            >
              <div class="mb-4 grid size-10 place-items-center rounded-lg ring-1" :class="statToneClass(card.tone)">
                <UIcon :name="card.icon" class="size-5" />
              </div>
              <p class="text-xs font-medium text-muted">{{ card.label }}</p>
              <div class="mt-1 flex items-baseline gap-2">
                <span class="text-2xl font-semibold text-highlighted">{{ card.value }}</span>
                <span class="truncate text-xs text-muted">{{ card.detail }}</span>
              </div>
            </div>
          </div>
        </div>
      </section>

      <ManageEmpty
        v-if="!items.length"
        icon="i-tabler-stack-2"
        text="还没有文档集"
      />

      <div v-else class="grid gap-4 xl:grid-cols-[minmax(0,1fr)_360px]">
        <section class="min-w-0 overflow-hidden rounded-lg border border-default bg-default">
          <div class="flex flex-wrap items-center justify-between gap-3 border-b border-default bg-elevated/35 px-4 py-3">
            <div>
              <p class="text-sm font-semibold text-highlighted">重点集合</p>
              <p class="text-xs text-muted">按风险缺口和内容量排序</p>
            </div>
          </div>

          <div class="divide-y divide-default">
            <div
              v-for="col in priorityCollections"
              :key="col.id"
              class="grid gap-3 px-4 py-3 lg:grid-cols-[minmax(240px,1.5fr)_120px_130px_108px] lg:items-center"
            >
              <div class="min-w-0">
                <div class="flex min-w-0 items-center gap-3">
                  <img
                    v-if="col.coverUrl"
                    :src="col.coverUrl"
                    :alt="col.title"
                    class="size-10 shrink-0 rounded-lg object-cover"
                  />
                  <span v-else class="grid size-10 shrink-0 place-items-center rounded-lg bg-primary/10 text-primary">
                    <UIcon :name="col.icon || 'i-tabler-stack-2'" class="size-5" />
                  </span>
                  <div class="min-w-0">
                    <p class="truncate text-sm font-semibold text-highlighted">{{ col.title }}</p>
                    <p class="truncate text-xs text-muted">{{ col.description || '未填写集合说明' }}</p>
                  </div>
                </div>
              </div>

              <p class="truncate font-mono text-sm text-default">{{ collectionPublicPath(col) || '-' }}</p>
              <div class="flex items-center gap-2 text-sm text-default">
                <span class="font-semibold text-highlighted">{{ col.docCount }}</span>
                <span class="text-muted">篇文档</span>
              </div>
              <div class="flex items-center justify-between gap-2 lg:justify-end">
                <UBadge
                  :label="healthMeta(col).label"
                  :color="healthMeta(col).color"
                  :icon="healthMeta(col).icon"
                  variant="subtle"
                  size="sm"
                />
              </div>
            </div>
          </div>
        </section>

        <aside class="space-y-4">
          <section class="rounded-lg border border-default bg-default p-4">
            <div class="mb-3 flex items-center justify-between gap-3">
              <div>
                <p class="text-sm font-semibold text-highlighted">发布准备度</p>
                <p class="text-xs text-muted">集合级基础门禁</p>
              </div>
              <UBadge
                :label="`${stats.ready}/${stats.collections}`"
                :color="stats.ready === stats.collections ? 'success' : 'warning'"
                icon="i-tabler-shield-check"
                variant="subtle"
              />
            </div>

            <div class="grid gap-2">
              <div
                v-for="row in readinessRows"
                :key="row.label"
                class="flex items-center justify-between gap-3 rounded-md bg-elevated px-2.5 py-2 text-sm"
              >
                <span class="text-default">{{ row.label }}</span>
                <span class="flex items-center gap-2">
                  <span class="font-mono text-xs text-muted">{{ row.value }}</span>
                  <UIcon
                    :name="row.ok ? 'i-tabler-circle-check' : 'i-tabler-circle-x'"
                    class="size-4"
                    :class="row.ok ? 'text-success' : 'text-warning'"
                  />
                </span>
              </div>
            </div>
          </section>
        </aside>
      </div>
    </template>
  </div>
</template>
