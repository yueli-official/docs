<script setup lang="ts">
type IconOption = { label: string, value: string }

const props = withDefaults(defineProps<{
  icon?: string
  coverUrl?: string
  previewUrl?: string
  title?: string
  description?: string
  accept?: string
  progress?: number
  uploading?: boolean
  disabled?: boolean
  iconOptions?: IconOption[]
}>(), {
  icon: '', coverUrl: '', previewUrl: '', title: '视觉资产',
  description: '选择一个 Tabler 图标，或上传封面图用于公开页展示。',
  accept: 'image/*', progress: -1, uploading: false, disabled: false,
})

const emit = defineEmits<{
  'update:icon': [value: string]
  'update:coverUrl': [value: string]
  'pick-cover': [file: File]
  'clear-cover': []
}>()

const fileInput = ref<HTMLInputElement>()
const currentIcon = computed(() => props.icon || 'i-tabler-stack-2')
const visibleCoverUrl = computed(() => props.previewUrl || props.coverUrl)
const hasCover = computed(() => Boolean(visibleCoverUrl.value))
const isUploading = computed(() => props.uploading || props.progress >= 0)
const coverModel = computed({
  get: () => props.coverUrl,
  set: value => emit('update:coverUrl', value),
})

function openFilePicker() {
  if (!props.disabled && !isUploading.value) fileInput.value?.click()
}

function onFileChange(event: Event) {
  const input = event.target as HTMLInputElement
  const file = input.files?.[0]
  input.value = ''
  if (file) emit('pick-cover', file)
}

function clearCover() {
  if (props.disabled || isUploading.value) return
  emit('update:coverUrl', '')
  emit('clear-cover')
}
</script>

<template>
  <section class="overflow-hidden rounded-lg border border-default bg-default">
    <div class="border-b border-default bg-elevated/35 px-3 py-2.5">
      <div class="flex items-start justify-between gap-3">
        <div class="min-w-0">
          <p class="text-sm font-semibold text-highlighted">{{ title }}</p>
          <p class="mt-0.5 text-xs text-muted">{{ description }}</p>
        </div>
        <span class="grid size-9 shrink-0 place-items-center rounded-lg bg-primary/10 text-primary"><UIcon :name="currentIcon" class="size-5" /></span>
      </div>
    </div>
    <div class="grid gap-4 p-3">
      <div class="relative aspect-[16/9] overflow-hidden rounded-lg border border-default bg-elevated">
        <img v-if="hasCover" :src="visibleCoverUrl" alt="" class="size-full object-cover">
        <div v-else class="grid size-full place-items-center text-muted">
          <div class="text-center">
            <span class="mx-auto grid size-12 place-items-center rounded-lg bg-default text-primary ring-1 ring-default"><UIcon :name="currentIcon" class="size-6" /></span>
            <p class="mt-2 text-xs text-muted">未设置封面时使用图标展示</p>
          </div>
        </div>
        <div v-if="isUploading" class="absolute inset-0 grid place-items-center bg-default/75 backdrop-blur-sm" role="status" aria-live="polite">
          <div class="w-3/4">
            <UProgress :model-value="Math.max(0, progress)" />
            <p class="mt-1 text-center text-xs text-muted">{{ Math.max(0, progress) }}%</p>
          </div>
        </div>
      </div>
      <div class="grid gap-3">
        <ManageIconPicker :model-value="currentIcon" :disabled="disabled" :icon-options="iconOptions" @update:model-value="emit('update:icon', $event)" />
        <UFormField label="封面">
          <div class="grid gap-2 sm:grid-cols-[minmax(0,1fr)_auto]">
            <UInput v-model="coverModel" icon="i-tabler-link" placeholder="https://..." class="min-w-0" :disabled="disabled || isUploading" />
            <div class="grid grid-cols-2 gap-2 sm:flex">
              <UButton icon="i-tabler-upload" :label="hasCover ? '更换' : '上传'" color="neutral" variant="outline" :disabled="disabled || isUploading" @click="openFilePicker" />
              <UButton icon="i-tabler-trash" label="移除" color="neutral" variant="ghost" :disabled="disabled || isUploading || !hasCover" @click="clearCover" />
            </div>
          </div>
        </UFormField>
        <input ref="fileInput" type="file" :accept="accept" class="hidden" @change="onFileChange">
      </div>
    </div>
  </section>
</template>
