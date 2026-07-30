export interface DocsReauthOptions {
  requireLoggedIn?: boolean
}

export type DocsReauthHandler = (options?: DocsReauthOptions) => Promise<boolean>

export function getOptionalDocsReauth(): DocsReauthHandler | undefined {
  const nuxtApp = tryUseNuxtApp() as
    | (ReturnType<typeof useNuxtApp> & { $platformReauth?: DocsReauthHandler })
    | null
  return nuxtApp?.$platformReauth
}
