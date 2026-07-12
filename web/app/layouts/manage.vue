<script setup lang="ts">
import { ManageShell, ManageUserMenu } from '@platform/manage/components'

const route = useRoute()
const { user, logout } = useAuth()
const accountUrl = computed(() => useRuntimeConfig().public.accountUrl || 'http://localhost:3000')
const { brand: siteBrand } = useSiteRuntime()

const contextLabel = computed(() => {
  if (route.path === '/manage') return '控制台'
  if (route.path === '/manage/collections') return '文档集'
  if (route.path === '/manage/docs') return '文档'
  if (route.path.startsWith('/manage/docs/')) return '编辑文档'
  if (route.path === '/manage/import') return '批量导入'
  if (route.path.startsWith('/manage/import/')) return '导入详情'
  if (route.path === '/manage/home') return '站点设置'
  return '控制台'
})
const showBackToTop = computed(() => ['/manage', '/manage/home'].includes(route.path))
</script>

<template>
  <ManageShell
    :site-name="siteBrand"
    :context-label="contextLabel"
    storage-key="docs-manage"
    content-class="max-w-screen-2xl"
    :show-back-to-top="showBackToTop"
  >
    <template #sidebar><ManageSidebar /></template>
    <template #user>
      <ManageUserMenu
        :name="user?.name"
        :email="user?.email"
        :settings-to="accountUrl"
        :logout
      />
    </template>
    <slot />
  </ManageShell>
</template>
