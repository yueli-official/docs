import type { MeView } from '~/types'
// Effective access comes from the Docs-local authorization instance, never
// from Identity role claims or client-side subject lists.
export function useMe() {
  const { call } = useApi()
  const me = useState<MeView | null>('docs-me', () => null)
  async function refresh() {
    try { me.value = (await call<{ me: MeView }>('/api/v1/me')).me }
    catch { me.value = null }
  }
  const isAdministrator = computed(() => me.value?.isAdministrator ?? false)
  const can = (capability: string) => me.value?.capabilities.includes(capability) ?? false
  const canManage = computed(() => [
    'docs.document.read',
    'docs.document.create',
    'docs.collection.manage',
    'docs.import.manage',
    'docs.site_settings.manage',
    'docs.asset_settings.manage',
  ].some(can))
  return { me, isAdministrator, can, canManage, refresh }
}
