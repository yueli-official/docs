<script setup lang="ts">
import type { HomeConfigResponse } from '~/types'

withDefaults(defineProps<{ widthClass?: string }>(), { widthClass: 'max-w-screen-xl' })
const { brand: siteBrand } = useSiteRuntime()
const { call } = useApi()
const { data: siteConfigData } = await useAsyncData(
  'docs-public-site-config',
  () => call<HomeConfigResponse>('/api/v1/home'),
  { default: () => ({ config: {} as HomeConfigResponse['config'] }) },
)
const config = computed(() => siteConfigData.value?.config)
const title = computed(() => config.value?.siteTitle || siteBrand.value)
const tagline = computed(() => config.value?.footerTagline || config.value?.siteDescription || '产品手册、集成说明和操作指南')
const copyright = computed(() => config.value?.footerCopyright || title.value)
const supportEmail = computed(() => config.value?.supportEmail || '')
</script>

<template>
  <footer class="border-t border-default">
    <div class="mx-auto grid w-full gap-1 px-4 py-8 text-center text-xs text-muted" :class="widthClass">
      <p>{{ tagline }}</p>
      <p>{{ copyright }}</p>
      <a v-if="supportEmail" :href="`mailto:${supportEmail}`" class="text-primary hover:underline">{{ supportEmail }}</a>
    </div>
  </footer>
</template>
