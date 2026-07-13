<script setup lang="ts">
import { ManageAssetSettings } from '@platform/asset/components'

definePageMeta({ layout: 'manage', middleware: 'auth' })

const { slug, brand } = useSiteRuntime()
useSeoMeta({ title: '资源配置 · 控制台' })
</script>

<template>
  <ManageAssetSettings
    :site-key="slug"
    :site-name="brand"
    description="配置文档站使用资源中心时的默认存储、文档集封面、正文图片与附件规则。"
    :profile-order="['docs-collection-cover', 'docs-import-image', 'attachment']"
  />
</template>
