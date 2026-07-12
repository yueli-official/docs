<script setup lang="ts">
import type { FormSubmitEvent } from '@nuxt/ui'
import { z } from 'zod'
import type { DocDetail } from '~/types'

interface QuickEditDoc extends DocDetail {
  collectionSlug: string
  slugPath: string[]
}

const open = defineModel<boolean>('open', { required: true })
const { doc, docs } = defineProps<{ doc?: QuickEditDoc, docs: readonly QuickEditDoc[] }>()
const emit = defineEmits<{
  saved: [doc: DocDetail]
  openFull: [doc: QuickEditDoc]
}>()

const { call } = useApi()
const schema = z.object({
  title: z.string().trim().min(1, '标题不能为空').max(200, '标题不能超过 200 个字符'),
  slug: z.string().trim().min(1, 'Slug 不能为空').max(200, 'Slug 不能超过 200 个字符'),
  status: z.enum(['draft', 'published', 'archived']),
  parentId: z.string()
})
type Schema = z.output<typeof schema>

const state = reactive<Schema>({ title: '', slug: '', status: 'draft', parentId: '' })
const saving = ref(false)
const submitError = ref('')
const statusItems = [
  { label: '草稿', value: 'draft' },
  { label: '已发布', value: 'published' },
  { label: '归档', value: 'archived' }
]

function descendantsOf(id: string) {
  const descendants = new Set<string>()
  let changed = true
  while (changed) {
    changed = false
    for (const item of docs) {
      if (descendants.has(item.id) || (item.parentId !== id && !descendants.has(item.parentId))) continue
      descendants.add(item.id)
      changed = true
    }
  }
  return descendants
}

const parentItems = computed(() => {
  if (!doc) return [{ label: '顶级文档', value: '' }]
  const excluded = descendantsOf(doc.id)
  excluded.add(doc.id)
  return [
    { label: '顶级文档', value: '' },
    ...docs
      .filter(item => item.collectionId === doc.collectionId && item.versionId === doc.versionId && item.locale === doc.locale && !excluded.has(item.id))
      .map(item => ({ label: item.title, value: item.id }))
  ]
})

const dirty = computed(() => Boolean(doc) && (
  state.title.trim() !== doc!.title
  || state.slug.trim() !== doc!.slug
  || state.status !== doc!.status
  || state.parentId !== doc!.parentId
))

function reset() {
  if (!doc) return
  state.title = doc.title
  state.slug = doc.slug
  state.status = schema.shape.status.safeParse(doc.status).success ? doc.status as Schema['status'] : 'draft'
  state.parentId = doc.parentId || ''
  submitError.value = ''
}

function close() {
  open.value = false
}

function openFullEditor() {
  if (!doc) return
  open.value = false
  emit('openFull', doc)
}

watch([() => open.value, () => doc?.id], ([isOpen]) => {
  if (isOpen) reset()
}, { immediate: true })

async function save(event: FormSubmitEvent<Schema>) {
  if (!doc || saving.value) return
  saving.value = true
  submitError.value = ''
  try {
    const response = await call<{ doc: DocDetail }>(`/api/v1/docs/${doc.id}`, {
      method: 'PATCH',
      body: event.data
    })
    emit('saved', response.doc)
    open.value = false
  } catch (error) {
    const apiError = error as { data?: { code?: string, message?: string } }
    if (apiError.data?.code === 'docs.slug_taken') {
      submitError.value = '同一父级下已经存在这个 Slug，请换一个。'
    } else {
      submitError.value = apiError.data?.message || '保存失败，请检查标题、路径和父级后重试。'
    }
  } finally {
    saving.value = false
  }
}
</script>

<template>
  <UModal
    v-model:open="open"
    title="快速编辑文档"
    description="修改层级列表中的高频信息；正文、摘要和 SEO 请使用完整编辑器。"
    scrollable
    :ui="{ content: 'sm:max-w-2xl', footer: 'p-0' }"
  >
    <template #body>
      <UForm v-if="doc" id="docs-quick-edit-form" :schema="schema" :state="state" class="space-y-5" @submit="save">
        <UAlert v-if="submitError" title="暂时无法保存" :description="submitError" icon="i-tabler-alert-circle" color="error" variant="soft" />

        <UFormField name="title" label="标题" required>
          <UInput v-model="state.title" class="w-full" placeholder="文档标题" autofocus />
        </UFormField>

        <div class="grid items-start gap-5 sm:grid-cols-[minmax(0,2fr)_minmax(10rem,1fr)]">
          <UFormField name="slug" label="Slug" description="修改后公开路径及子文档路径会随之变化。" required :ui="{ description: 'min-h-5' }">
            <UInput v-model="state.slug" class="w-full font-mono" icon="i-tabler-link" placeholder="document-slug" />
          </UFormField>
          <UFormField name="status" label="生命周期" description="控制文档是否公开。" required :ui="{ description: 'min-h-5' }">
            <USelect v-model="state.status" :items="statusItems" value-key="value" class="w-full" />
          </UFormField>
        </div>

        <UFormField name="parentId" label="父级文档" description="父级必须属于相同文档集、版本和语言。">
          <USelectMenu v-model="state.parentId" :items="parentItems" value-key="value" :search-input="{ placeholder: '搜索父级文档' }" class="w-full" />
        </UFormField>
      </UForm>
    </template>

    <template #footer>
      <div class="flex w-full flex-col-reverse gap-2 p-4 sm:flex-row sm:items-center">
        <UButton label="打开完整编辑器" icon="i-tabler-file-pencil" color="neutral" variant="ghost" class="justify-center sm:justify-start" @click="openFullEditor" />
        <div class="flex gap-2 sm:ml-auto">
          <UButton label="取消" color="neutral" variant="outline" class="flex-1 justify-center sm:flex-none" :disabled="saving" @click="close" />
          <UButton form="docs-quick-edit-form" type="submit" label="保存更改" icon="i-tabler-check" class="flex-1 justify-center sm:flex-none" :loading="saving" :disabled="!dirty" />
        </div>
      </div>
    </template>
  </UModal>
</template>
