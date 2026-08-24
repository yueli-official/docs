<script setup lang="ts">
const props = withDefaults(
  defineProps<{
    label: string;
    value: string | number;
    detail: string;
    icon: string;
    tone?: "primary" | "success" | "warning" | "neutral";
    to?: string;
  }>(),
  { tone: "primary", to: "" },
);

const root = computed(() =>
  props.to ? resolveComponent("NuxtLink") : "article",
);
const toneClass = computed(
  () =>
    ({
      primary: "text-primary",
      success: "text-success",
      warning: "text-warning",
      neutral: "text-muted",
    })[props.tone],
);
</script>

<template>
  <component
    :is="root"
    :to="to || undefined"
    class="relative grid min-w-0 grid-cols-[2.625rem_minmax(0,1fr)] gap-3 overflow-hidden rounded-xl bg-elevated p-4 shadow-sm transition-colors duration-150"
    :class="[toneClass, to && 'hover:bg-accented/60']"
  >
    <div class="grid size-10 place-items-center rounded-xl bg-current/10">
      <UIcon :name="icon" class="size-5" />
    </div>
    <div class="min-w-0">
      <p class="text-xs text-muted">{{ label }}</p>
      <p
        class="mt-0.5 text-2xl font-bold leading-tight tracking-[-0.03em] text-highlighted tabular-nums"
      >
        {{ value }}
      </p>
      <p class="mt-0.5 truncate text-xs text-dimmed">{{ detail }}</p>
    </div>
  </component>
</template>
