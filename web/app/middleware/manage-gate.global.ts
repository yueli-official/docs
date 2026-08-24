// The API remains authoritative; this gate only keeps users out of console
// surfaces for which their effective capability set has no entry point.
export default defineNuxtRouteMiddleware(async (to) => {
  if (!to.path.startsWith('/manage')) return
  const { user, refresh, login } = useAuth()
  if (!user.value) await refresh()
  if (!user.value) return login(to.fullPath)

  const { me, canManage, refresh: refreshMe } = useMe()
  if (!me.value) await refreshMe()
  // Resolve the product-local capability set during SSR so the navigation tree is
  // identical before and after hydration. Redirect decisions remain client-side;
  // a transient SSR access lookup must not turn into an incorrect public redirect.
  if (import.meta.server) return
  if (canManage.value) return
  return navigateTo('/')
})
