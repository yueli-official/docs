<script setup lang="ts">
type IconOption = { label: string, value: string }

const defaultIcons: IconOption[] = [
  { label: '文档集', value: 'i-tabler-stack-2' }, { label: '指南', value: 'i-tabler-book-2' },
  { label: '手册', value: 'i-tabler-manual-gearbox' }, { label: '快速开始', value: 'i-tabler-rocket' },
  { label: '接口', value: 'i-tabler-api' }, { label: '代码', value: 'i-tabler-code' },
  { label: '设置', value: 'i-tabler-settings' }, { label: '权限', value: 'i-tabler-shield-lock' },
  { label: '数据', value: 'i-tabler-database' }, { label: '组件', value: 'i-tabler-components' },
  { label: '工具', value: 'i-tabler-tool' }, { label: '发布', value: 'i-tabler-world-upload' },
  { label: '目录', value: 'i-tabler-sitemap' }, { label: '终端', value: 'i-tabler-terminal-2' },
  { label: '安全', value: 'i-tabler-lock-check' }, { label: '星标', value: 'i-tabler-star' },
  { label: '钻石', value: 'i-tabler-diamond' }, { label: '实验', value: 'i-tabler-flask' },
]

const props = withDefaults(defineProps<{
  modelValue?: string
  disabled?: boolean
  iconOptions?: IconOption[]
  compact?: boolean
}>(), { modelValue: '', disabled: false, compact: false })

const emit = defineEmits<{ 'update:modelValue': [value: string] }>()
const currentIcon = computed(() => props.modelValue || 'i-tabler-stack-2')
const iconOptions = computed(() => props.iconOptions?.length ? props.iconOptions : defaultIcons)

function chooseIcon(value: string) {
  if (!props.disabled) emit('update:modelValue', value)
}
</script>

<template>
  <div class="grid gap-2">
    <div v-if="!compact" class="flex items-center justify-between gap-3">
      <p class="text-xs font-medium text-muted">图标</p>
      <span class="truncate font-mono text-xs text-muted">{{ currentIcon }}</span>
    </div>
    <div class="grid auto-rows-[2.25rem] justify-start gap-1.5" :class="compact ? 'max-h-48 grid-cols-[repeat(6,2.25rem)] overflow-y-auto' : 'grid-cols-[repeat(auto-fill,2.25rem)]'">
      <UTooltip v-for="item in iconOptions" :key="item.value" :text="item.label">
        <UButton
          :icon="item.value"
          :color="currentIcon === item.value ? 'primary' : 'neutral'"
          :variant="currentIcon === item.value ? 'soft' : 'outline'"
          size="sm"
          square
          class="grid size-9 place-items-center p-0"
          :ui="{ leadingIcon: 'mx-auto size-4 shrink-0' }"
          :aria-label="`选择${item.label}`"
          :aria-pressed="currentIcon === item.value"
          :disabled="disabled"
          @click="chooseIcon(item.value)"
        />
      </UTooltip>
    </div>
  </div>
</template>
