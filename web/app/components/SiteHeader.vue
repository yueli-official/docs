<script setup lang="ts">
// Shared public site header (logo home-link + account menu). Used by the
// default layout and the collection (reading) layout so no public page is
// stranded without top chrome.
import type { DropdownMenuItem } from '@nuxt/ui'
import type { HomeConfigResponse } from '~/types'

withDefaults(defineProps<{ widthClass?: string }>(), { widthClass: 'max-w-screen-xl' })
const { user, loggedIn, login, logout } = useAuth()
const { isOwner, refresh: refreshMe } = useMe()
const accountUrl = computed(() => useRuntimeConfig().public.accountUrl || 'http://localhost:3000')
const initial = computed(() => (user.value?.name || user.value?.email || '?').charAt(0).toUpperCase())
const searchOpen = ref(false)
const { brand: siteBrand } = useSiteRuntime()
const { call } = useApi()
const { data: siteConfigData } = await useAsyncData(
  'docs-public-site-config',
  () => call<HomeConfigResponse>('/api/v1/home'),
  { default: () => ({ config: {} as HomeConfigResponse['config'] }) },
)
const siteTitle = computed(() => siteConfigData.value?.config?.siteTitle || siteBrand.value)

async function handleLogin() {
  await login()
}

watch(loggedIn, async (value) => {
  if (value) await refreshMe()
}, { immediate: true })

const userMenuItems = computed<DropdownMenuItem[][]>(() => {
  const primary = [{ label: user.value?.name || user.value?.email || '', type: 'label' as const }]
  const nav = [
    ...(isOwner.value ? [{ label: '控制台', icon: 'i-tabler-layout-dashboard', to: '/manage' }] : []),
    { label: '用户设置', icon: 'i-tabler-user-cog', onSelect: () => navigateTo(accountUrl.value, { external: true }) },
  ]
  const session = [{ label: '退出登录', icon: 'i-tabler-logout', onSelect: () => logout() }]
  return [primary, nav, session]
})
</script>

<template>
  <header class="sticky top-0 z-20 border-b border-default bg-default/75 backdrop-blur">
    <div class="mx-auto flex h-16 w-full items-center justify-between gap-4 px-4" :class="widthClass">
      <NuxtLink to="/" class="font-display flex items-center gap-2 text-base font-semibold text-highlighted">
        <span class="grid size-8 place-items-center rounded-lg bg-primary/10 text-primary">
          <UIcon name="i-tabler-book-2" class="size-5" />
        </span>
        {{ siteTitle }}
      </NuxtLink>
      <div class="flex items-center gap-1.5">
        <UButton
          to="/collections"
          color="neutral"
          variant="ghost"
          icon="i-tabler-stack-2"
          label="文档集"
          class="hidden md:inline-flex"
        />
        <UButton
          color="neutral"
          variant="ghost"
          icon="i-tabler-search"
          class="hidden sm:inline-flex"
          @click="() => { searchOpen = true }"
        >
          <span>搜索</span>
          <span class="ml-1 hidden items-center gap-1 rounded bg-elevated px-1.5 py-0.5 text-[10px] text-muted lg:inline-flex">Ctrl K</span>
        </UButton>
        <UTooltip text="搜索文档">
          <UButton
            color="neutral"
            variant="ghost"
            icon="i-tabler-search"
            square
            aria-label="搜索文档"
            class="sm:hidden"
            @click="() => { searchOpen = true }"
          />
        </UTooltip>
        <UColorModeButton aria-label="切换夜间模式" />
        <template v-if="loggedIn">
          <UDropdownMenu
            :items="userMenuItems"
            :ui="{ content: 'w-48' }"
          >
            <UButton variant="ghost" color="neutral" class="gap-2 px-1.5">
              <UAvatar :text="initial" size="xs" />
              <span class="hidden max-w-32 truncate text-sm sm:block">{{ user?.name || user?.email }}</span>
            </UButton>
          </UDropdownMenu>
        </template>
        <UButton v-else variant="ghost" color="neutral" icon="i-tabler-login-2" label="登录" @click="handleLogin" />
      </div>
    </div>
    <DocsSearchDialog v-model:open="searchOpen" />
  </header>
</template>
