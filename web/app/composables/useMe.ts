import type { MeView } from '~/types'
// Docs operator status comes from the docs backend (sub ∈ docs.operatorSubs), not from
// the IdP roles claim — so we must ask /api/v1/me rather than read the token.
export function useMe() {
  const { call } = useApi()
  const me = useState<MeView | null>('docs-me', () => null)
  async function refresh() {
    try { me.value = (await call<{ me: MeView }>('/api/v1/me')).me }
    catch { me.value = null }
  }
  const isOwner = computed(() => me.value?.isOwner ?? false)
  return { me, isOwner, refresh }
}
