export function useSiteRuntime() {
  const config = useRuntimeConfig()
  return {
    slug: computed(() => config.public.siteSlug || 'docs-ae'),
    brand: computed(() => config.public.siteBrand || '文档库'),
    domain: computed(() => config.public.siteDomain || 'docs-ae.localhost'),
    assetSpace: computed(() => config.public.assetSpace || 'default'),
    assetNamespace: computed(() => config.public.assetNamespace || 'default'),
    assetProfile: computed(() => config.public.assetProfile || 'docs-default')
  }
}
