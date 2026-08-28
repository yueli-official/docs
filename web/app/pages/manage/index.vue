<script setup lang="ts">
import DashboardTrendChart from "~/components/DashboardTrendChart.vue";
import ManageMetricCard from "~/components/ManageMetricCard.vue";
import type { DashboardOverview } from "~/types";

definePageMeta({ layout: "manage" });
useSeoMeta({ title: "控制台" });

const { call } = useApi();
const period = ref(14);
const periodItems = [
  { label: "7 天", value: 7 },
  { label: "14 天", value: 14 },
  { label: "30 天", value: 30 },
];
const dashboardCardClass = "bg-elevated shadow-sm";
const mounted = ref(false);
onMounted(() => {
  mounted.value = true;
});

const {
  data: analytics,
  pending: analyticsPending,
  error: analyticsError,
  refresh,
} = await useAsyncData(
  "docs-dashboard-analytics",
  () =>
    call<DashboardOverview>("/api/v1/dashboard/overview", {
      query: { days: period.value },
    }),
  {
    server: false,
    watch: [period],
    default: () => ({
      days: period.value,
      allTimeViews: 0,
      allTimeUniqueVisitorDays: 0,
      periodViews: 0,
      periodUniqueVisitorDays: 0,
      previousPeriodViews: 0,
      previousUniqueVisitorDays: 0,
      periodSearches: 0,
      previousPeriodSearches: 0,
      zeroResultSearches: 0,
      series: [],
      topDocuments: [],
      topSources: [],
      topSearches: [],
    }),
  },
);

const formatter = new Intl.NumberFormat("zh-CN");
function formatNumber(value: number) {
  return formatter.format(value);
}
function formatMetricValue(value: number) {
  return Number.isInteger(value) ? formatNumber(value) : value.toFixed(1);
}
function comparison(current: number, previous: number, unit = "") {
  const suffix = unit ? ` ${unit}` : "";
  if (!previous)
    return current
      ? `本期新增 ${formatNumber(current)}${suffix}`
      : "与前期持平";
  const percent = Math.round(((current - previous) / previous) * 100);
  return `${percent >= 0 ? "+" : ""}${percent}% 较前 ${period.value} 天`;
}
const readingDepth = computed(() => {
  const visitors = analytics.value?.periodUniqueVisitorDays ?? 0;
  return visitors ? (analytics.value?.periodViews ?? 0) / visitors : 0;
});
const zeroResultRate = computed(() => {
  const searches = analytics.value?.periodSearches ?? 0;
  return searches
    ? ((analytics.value?.zeroResultSearches ?? 0) / searches) * 100
    : 0;
});
const visitorAverage = computed(
  () => (analytics.value?.periodUniqueVisitorDays ?? 0) / period.value,
);
const visitorPeakPoint = computed(() =>
  (analytics.value?.series ?? []).reduce<
    DashboardOverview["series"][number] | undefined
  >(
    (peak, point) =>
      !peak || point.uniqueVisitorDays > peak.uniqueVisitorDays ? point : peak,
    undefined,
  ),
);
const visitorComparison = computed(() =>
  comparison(
    analytics.value?.periodUniqueVisitorDays ?? 0,
    analytics.value?.previousUniqueVisitorDays ?? 0,
    "访客日",
  ),
);
const visitorComparisonColor = computed(() =>
  (analytics.value?.periodUniqueVisitorDays ?? 0) >=
  (analytics.value?.previousUniqueVisitorDays ?? 0)
    ? ("success" as const)
    : ("warning" as const),
);
const topPeak = computed(() =>
  Math.max(
    1,
    ...(analytics.value?.topDocuments ?? []).map((document) => document.views),
  ),
);
const sourcePeak = computed(() =>
  Math.max(1, ...(analytics.value?.topSources ?? []).map((source) => source.views)),
);
const searchPeak = computed(() =>
  Math.max(1, ...(analytics.value?.topSearches ?? []).map((search) => search.searches)),
);
const metricCards = computed(() => [
  {
    label: "累计浏览",
    value: analytics.value?.allTimeViews ?? 0,
    detail: `${formatNumber(analytics.value?.allTimeUniqueVisitorDays ?? 0)} 访客日`,
    icon: "i-tabler-eye",
    tone: "primary" as const,
  },
  {
    label: `近 ${period.value} 天浏览`,
    value: analytics.value?.periodViews ?? 0,
    detail: comparison(
      analytics.value?.periodViews ?? 0,
      analytics.value?.previousPeriodViews ?? 0,
    ),
    icon: "i-tabler-chart-line",
    tone: "primary" as const,
  },
  {
    label: "浏览深度",
    value: readingDepth.value,
    detail: "平均每访客日浏览",
    icon: "i-tabler-chart-dots-3",
    tone: "neutral" as const,
  },
  {
    label: "搜索次数",
    value: analytics.value?.periodSearches ?? 0,
    detail: analytics.value?.periodSearches
      ? `${zeroResultRate.value.toFixed(0)}% 未找到结果`
      : "当前周期暂无搜索",
    icon: "i-tabler-search",
    tone: zeroResultRate.value > 20 ? ("warning" as const) : ("success" as const),
  },
]);

function setPeriod(value: string | number) {
  period.value = Number(value);
}

function sourceLabel(value: string) {
  const labels: Record<string, string> = {
    direct: "直接访问",
    "google.com": "Google",
    "bing.com": "Bing",
    "baidu.com": "百度",
    "zhihu.com": "知乎",
    "x.com": "X",
    "weibo.com": "微博",
  };
  return labels[value] ?? value;
}

function publicDocumentLink(document: DashboardOverview["topDocuments"][number]) {
  return {
    path: `/${document.collectionSlug}/${document.slugPath}`,
    query: {
      ...(document.locale ? { locale: document.locale } : {}),
      ...(document.versionKey ? { version: document.versionKey } : {}),
    },
  };
}
</script>

<template>
  <ManagePage
    id="dashboard"
    title="控制台"
    icon="i-tabler-dashboard"
    main-id="manage-main"
    body-class="w-full space-y-5"
    data-docs-dashboard-analytics
  >
    <UAlert
      v-if="analyticsError"
      color="error"
      variant="subtle"
      icon="i-tabler-alert-circle"
      title="统计暂时不可用"
      description="文档阅读不受影响；重试后仍失败时再检查 Docs API。"
    >
      <template #actions>
        <UButton
          color="error"
          variant="soft"
          icon="i-tabler-refresh"
          label="重试"
          @click="() => refresh()"
        />
      </template>
    </UAlert>

    <div v-if="!mounted || (analyticsPending && !analytics?.series.length)" class="grid gap-3 sm:grid-cols-2 lg:grid-cols-4">
      <USkeleton v-for="item in 4" :key="item" class="h-28 rounded-xl" />
    </div>
    <div v-else class="grid gap-3 sm:grid-cols-2 lg:grid-cols-4">
      <ManageMetricCard
        v-for="card in metricCards"
        :key="card.label"
        :label="card.label"
        :value="formatMetricValue(card.value)"
        :detail="card.detail"
        :icon="card.icon"
        :tone="card.tone"
        data-docs-dashboard-metric
      />
    </div>

    <div class="grid items-start gap-4 xl:grid-cols-[minmax(0,2fr)_minmax(18rem,1fr)]">
      <UCard
        variant="soft"
        :class="[dashboardCardClass, 'divide-y-0']"
        data-docs-dashboard-trend
        :aria-busy="analyticsPending"
        :ui="{ header: 'p-5 sm:p-6', body: 'px-5 pb-5 pt-0 sm:px-6 sm:pb-6 sm:pt-0' }"
      >
        <template #header>
          <div class="flex flex-col gap-4 sm:flex-row sm:items-start sm:justify-between">
            <h2 class="text-base font-semibold text-highlighted">浏览趋势</h2>
            <UTabs
              :model-value="period"
              :items="periodItems"
              :content="false"
              :disabled="!mounted || analyticsPending"
              size="sm"
              class="w-fit"
              aria-label="统计时间范围"
              @update:model-value="setPeriod"
            />
          </div>
        </template>
        <DashboardTrendChart :points="analytics?.series ?? []" />
      </UCard>

      <UCard
        title="访客趋势"
        variant="soft"
        :class="[dashboardCardClass, 'divide-y-0']"
        data-docs-dashboard-audience
        :ui="{ header: 'p-5 sm:p-6', body: 'space-y-5 px-5 pb-5 pt-0 sm:px-6 sm:pb-6 sm:pt-0' }"
      >
        <div class="flex items-end justify-between gap-4">
          <div>
            <p class="text-3xl font-semibold tabular-nums tracking-tight text-highlighted">
              {{ formatNumber(analytics?.periodUniqueVisitorDays ?? 0) }}
            </p>
            <p class="mt-1 text-xs text-muted">访客日</p>
          </div>
          <UBadge
            :color="visitorComparisonColor"
            variant="soft"
            :label="visitorComparison"
          />
        </div>
        <DashboardTrendChart
          :points="analytics?.series ?? []"
          metric="uniqueVisitorDays"
          compact
        />
        <dl class="grid grid-cols-2 gap-4">
          <div>
            <dt class="text-xs text-muted">日均访客日</dt>
            <dd class="mt-1 text-lg font-semibold tabular-nums text-highlighted">
              {{ visitorAverage.toFixed(1) }}
            </dd>
          </div>
          <div>
            <dt class="text-xs text-muted">单日峰值</dt>
            <dd class="mt-1 text-lg font-semibold tabular-nums text-highlighted">
              {{ formatNumber(visitorPeakPoint?.uniqueVisitorDays ?? 0) }}
            </dd>
          </div>
        </dl>
      </UCard>
    </div>

    <div class="grid items-start gap-4 xl:grid-cols-[minmax(0,2fr)_minmax(18rem,1fr)]">
      <UCard
        title="热门文档"
        variant="soft"
        :class="[dashboardCardClass, 'divide-y-0']"
        data-docs-dashboard-top-documents
        :ui="{ header: 'p-5 sm:p-6', body: 'p-0 sm:p-0' }"
      >
        <div v-if="analytics?.topDocuments.length" class="divide-y divide-default">
          <NuxtLink
            v-for="(document, index) in analytics.topDocuments"
            :key="document.id"
            :to="publicDocumentLink(document)"
            class="group grid grid-cols-[auto_minmax(0,1fr)_auto] items-center gap-3 px-5 py-4 transition hover:bg-accented sm:px-6"
          >
            <UBadge color="neutral" variant="soft" :label="String(index + 1)" />
            <div class="min-w-0 space-y-2">
              <p class="truncate text-sm font-medium text-highlighted group-hover:text-primary">
                {{ document.title }}
              </p>
              <UProgress :model-value="document.views" :max="topPeak" size="2xs" />
              <p class="text-xs text-muted">{{ document.uniqueVisitorDays }} 访客日</p>
            </div>
            <div class="text-end">
              <p class="text-lg font-semibold tabular-nums text-highlighted">
                {{ formatNumber(document.views) }}
              </p>
              <p class="text-xs text-dimmed">次浏览</p>
            </div>
          </NuxtLink>
        </div>
        <div v-else class="p-8 text-center text-sm text-muted">
          当前周期还没有文档浏览记录。
        </div>
      </UCard>

      <div class="space-y-4">
        <UCard
          title="流量来源"
          variant="soft"
          :class="[dashboardCardClass, 'divide-y-0']"
          data-docs-dashboard-sources
          :ui="{ header: 'p-5 sm:p-6', body: 'space-y-4 px-5 pb-5 pt-0 sm:px-6 sm:pb-6 sm:pt-0' }"
        >
          <div v-if="analytics?.topSources.length" class="space-y-4">
            <div v-for="source in analytics.topSources" :key="source.source" class="space-y-2">
              <div class="flex items-center justify-between gap-3 text-sm">
                <span class="min-w-0 truncate text-toned">{{ sourceLabel(source.source) }}</span>
                <span class="shrink-0 font-medium tabular-nums text-highlighted">
                  {{ formatNumber(source.views) }}
                </span>
              </div>
              <UProgress :model-value="source.views" :max="sourcePeak" size="2xs" />
            </div>
          </div>
          <p v-else class="text-sm text-muted">当前周期还没有可归因的来源。</p>
        </UCard>

        <UCard
          title="热门搜索"
          variant="soft"
          :class="[dashboardCardClass, 'divide-y-0']"
          data-docs-dashboard-searches
          :ui="{ header: 'p-5 sm:p-6', body: 'space-y-4 px-5 pb-5 pt-0 sm:px-6 sm:pb-6 sm:pt-0' }"
        >
          <div v-if="analytics?.topSearches.length" class="space-y-4">
            <div v-for="search in analytics.topSearches" :key="search.query" class="space-y-2">
              <div class="flex items-center justify-between gap-3 text-sm">
                <span class="min-w-0 truncate text-toned">{{ search.query }}</span>
                <span class="shrink-0 font-medium tabular-nums text-highlighted">
                  {{ formatNumber(search.searches) }}
                </span>
              </div>
              <UProgress :model-value="search.searches" :max="searchPeak" size="2xs" />
              <p v-if="search.zeroResults" class="text-xs text-warning">
                {{ search.zeroResults }} 次没有结果
              </p>
            </div>
          </div>
          <p v-else class="text-sm text-muted">当前周期还没有搜索记录。</p>
        </UCard>
      </div>
    </div>
  </ManagePage>
</template>
