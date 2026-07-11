<script setup lang="ts">
// Console sidebar: brand + nav (控制台/文档集/文档). Back-to-site and
// the user live in the layout's top bar. All items are always visible: the
// gate (manage-gate.global.ts) already enforces owner-only, so no isOwner branch.
const route = useRoute()
const { brand: siteBrand } = useSiteRuntime()

const nav = [
  { label: '控制台', icon: 'i-tabler-dashboard', to: '/manage' },
  { label: '文档集', icon: 'i-tabler-stack-2', to: '/manage/collections' },
  { label: '文档', icon: 'i-tabler-file-text', to: '/manage/docs' },
  { label: '批量导入', icon: 'i-tabler-file-import', to: '/manage/import' },
  { label: '设置', icon: 'i-tabler-settings', to: '/manage/home' },
]
// `控制台` is the index — exact match; the rest match their subtree.
function isActive(to: string) {
  return to === '/manage' ? route.path === '/manage' : route.path.startsWith(to)
}
</script>

<template>
  <div class="flex h-full flex-col bg-elevated/30">
    <NuxtLink to="/manage" class="font-display flex h-16 items-center gap-2 border-b border-default px-5 font-semibold text-highlighted">
      <span class="grid size-8 place-items-center rounded-lg bg-primary/10 text-primary"><UIcon name="i-tabler-book" class="size-5" /></span>
      {{ siteBrand }}
    </NuxtLink>

    <nav class="flex-1 space-y-1 p-3">
      <NuxtLink
        v-for="item in nav"
        :key="item.to"
        :to="item.to"
        class="flex items-center gap-3 rounded-lg px-3 py-2 text-sm font-medium transition"
        :class="isActive(item.to) ? 'bg-primary/10 text-primary' : 'text-muted hover:bg-elevated hover:text-default'"
      >
        <UIcon :name="item.icon" class="size-5 shrink-0" />{{ item.label }}
      </NuxtLink>
    </nav>
  </div>
</template>
