<script setup lang="ts">
import { PageHeader } from "@yueli/ui/admin";

defineOptions({ inheritAttrs: false });

const props = withDefaults(
  defineProps<{
    id: string;
    title: string;
    description?: string;
    icon?: string;
    bodyClass?: string;
    mainId?: string;
  }>(),
  { bodyClass: "", mainId: "manage-main" },
);
const slots = useSlots();
const headingId = computed(() => `${props.id}-title`);
const pageDescription = computed(() => props.description || ({
  documents: "组织文档、查看发布状态与维护目录",
  collections: "按主题组织知识，管理文档集与公开路径",
  comments: "查看阅读反馈，处理待审核评论",
  "project-docs": "管理项目文档来源与同步任务",
} as Record<string, string>)[props.id]);
</script>

<template>
  <section
    v-bind="$attrs"
    :id="id"
    :aria-labelledby="headingId"
    class="min-w-0 space-y-5"
    data-manage-page
  >
    <PageHeader :title="title" :description="pageDescription" :icon="icon" :heading-id="headingId">
      <template v-if="slots.tools" #tools><slot name="tools" /></template>
      <template v-if="slots.actions" #actions><slot name="actions" /></template>
    </PageHeader>

    <div
      v-if="slots.toolbar || slots['toolbar-left'] || slots['toolbar-right']"
      class="yueli-card flex flex-wrap items-center justify-between gap-3 px-4 py-3 sm:px-5"
      data-manage-page-toolbar
    >
      <div class="flex min-w-0 flex-1 flex-wrap items-center gap-2">
        <slot name="toolbar-left" />
        <slot name="toolbar" />
      </div>
      <div v-if="slots['toolbar-right']" class="flex items-center gap-2">
        <slot name="toolbar-right" />
      </div>
    </div>

    <div class="min-w-0" :class="bodyClass"><slot /></div>
    <footer v-if="slots.footer"><slot name="footer" /></footer>
  </section>
</template>
