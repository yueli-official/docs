// The API remains authoritative; this gate only keeps users out of console
// surfaces for which their effective capability set has no entry point.
export default defineNuxtRouteMiddleware(async (to) => {
  if (!to.path.startsWith('/manage')) return
  const { user, refresh, login } = useAuth()
  if (!user.value) await refresh()
  if (!user.value) return login(to.fullPath)

  // Effective access is checked by the docs backend through the BFF session. On the
  // server, useApi would call the backend directly without the sealed BFF token,
  // so do the owner check only after the client session is available.
  if (import.meta.server) return

  const { me, canManage, refresh: refreshMe } = useMe()
  if (!me.value) await refreshMe()
  if (canManage.value) return
  return navigateTo('/')
})
