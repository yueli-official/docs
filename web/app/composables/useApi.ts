import { useApi as useFoundationApi } from '@yueli/nuxt-runtime/runtime'
import { toDocsApiRequest, type DocsApiCallOptions } from '../utils/apiCompat'

const REAUTH_COOLDOWN_MS = 10_000
let reauthFlight: Promise<boolean> | undefined

function mayReauth(): boolean {
  try {
    const last = Number(sessionStorage.getItem('reauth-at') || 0)
    if (Date.now() - last < REAUTH_COOLDOWN_MS) return false
    sessionStorage.setItem('reauth-at', String(Date.now()))
  } catch {
    // Browser storage is optional; the single-flight guard still applies.
  }
  return true
}

async function restoreSession() {
  const reauth = getOptionalDocsReauth()
  if (!reauth) return false
  if (!reauthFlight) {
    const current = reauth({ requireLoggedIn: true }).finally(() => {
      if (reauthFlight === current) reauthFlight = undefined
    })
    reauthFlight = current
  }
  return reauthFlight
}

export function useApi(target = 'docs') {
  const request = useFoundationApi(target)

  async function call<T>(url: string, options?: DocsApiCallOptions): Promise<T> {
    const prepared = toDocsApiRequest<T>(url, options)
    try {
      return await request.request(prepared.path, prepared.options)
    } catch (error: unknown) {
      const status = (error as { statusCode?: number })?.statusCode
      if (import.meta.client && status === 401 && mayReauth()) await restoreSession()
      throw error
    }
  }

  return { call, request }
}
