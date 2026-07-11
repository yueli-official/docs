// The docs console is operator-only (sub ∈ docs.operatorSubs). Login is the 'auth'
// middleware's job; mirrored here so a direct /manage URL while logged-out
// bounces to login, and a logged-in non-owner bounces to the public site.
export default defineNuxtRouteMiddleware(async (to) => {
  if (!to.path.startsWith('/manage')) return
  const { user, refresh, login } = useAuth()
  if (!user.value) await refresh()
  if (!user.value) return login(to.fullPath)

  // Owner-ness is checked by the docs backend through the BFF session. On the
  // server, useApi would call the backend directly without the sealed BFF token,
  // so do the owner check only after the client session is available.
  if (import.meta.server) return

  const { me, refresh: refreshMe } = useMe()
  if (!me.value) await refreshMe()
  if (me.value?.isOwner) return
  return navigateTo('/')
})
