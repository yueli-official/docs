<script setup lang="ts">
import type { HomeConfigResponse } from '~/types'

withDefaults(defineProps<{ widthClass?: string }>(), { widthClass: 'max-w-screen-xl' })
const { call } = useApi()
const { data: siteConfigData } = await useAsyncData(
  'docs-public-site-config',
  () => call<HomeConfigResponse>('/api/v1/home'),
)
if (!siteConfigData.value?.config) throw createError({ statusCode: 500, statusMessage: '文档站点配置尚未初始化' })
const config = computed(() => siteConfigData.value!.config)
const tagline = computed(() => config.value.footerTagline)
const copyright = computed(() => config.value.footerCopyright)
const supportEmail = computed(() => config.value.supportEmail)
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
