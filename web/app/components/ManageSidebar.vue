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
  {
    label: '设置',
    icon: 'i-tabler-settings',
    to: '/manage/home',
    children: [
      { label: '首页', icon: 'i-tabler-home-cog', to: '/manage/home?section=home', section: 'home' },
      { label: '页脚', icon: 'i-tabler-layout-bottombar', to: '/manage/home?section=footer', section: 'footer' },
      { label: '基础', icon: 'i-tabler-adjustments-horizontal', to: '/manage/home?section=site', section: 'site' },
    ],
  },
]
// `控制台` is the index — exact match; the rest match their subtree.
function isActive(to: string) {
  const path = to.split('?')[0] || to
  return path === '/manage' ? route.path === '/manage' : route.path.startsWith(path)
}

function isChildActive(section?: string) {
  return route.path === '/manage/home' && (route.query.section || 'home') === section
}
</script>

<template>
  <div class="flex h-full flex-col bg-elevated/30">
    <NuxtLink to="/manage" class="font-display flex h-16 items-center gap-2 border-b border-default px-5 font-semibold text-highlighted">
      <span class="grid size-8 place-items-center rounded-lg bg-primary/10 text-primary"><UIcon name="i-tabler-book" class="size-5" /></span>
      {{ siteBrand }}
    </NuxtLink>

    <nav aria-label="文档后台" class="flex-1 space-y-1 p-3">
      <div v-for="item in nav" :key="item.to" class="space-y-1">
        <NuxtLink
          :to="item.to"
          class="flex items-center gap-3 rounded-lg px-3 py-2 text-sm font-medium transition"
          :class="isActive(item.to) ? 'bg-primary/10 text-primary' : 'text-muted hover:bg-elevated hover:text-default'"
        >
          <UIcon :name="item.icon" class="size-5 shrink-0" />
          <span class="min-w-0 flex-1 truncate">{{ item.label }}</span>
          <UIcon v-if="item.children?.length" name="i-tabler-chevron-down" class="size-4 shrink-0 opacity-70" />
        </NuxtLink>
        <div v-if="item.children?.length && isActive(item.to)" class="ml-4 space-y-1 border-l border-default pl-3">
          <NuxtLink
            v-for="child in item.children"
            :key="child.to"
            :to="child.to"
            class="flex items-center gap-2 rounded-md px-2 py-1.5 text-xs font-medium transition"
            :class="isChildActive(child.section) ? 'bg-primary/10 text-primary' : 'text-muted hover:bg-elevated hover:text-default'"
          >
            <UIcon :name="child.icon" class="size-4 shrink-0" />
            <span class="truncate">{{ child.label }}</span>
          </NuxtLink>
        </div>
      </div>
    </nav>
  </div>
</template>
