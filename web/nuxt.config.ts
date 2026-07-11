// Nuxt 4 config for the docs-site app (consumer of the docs service).
const siteBrand = process.env.NUXT_PUBLIC_SITE_BRAND || '文档库'

export default defineNuxtConfig({
  extends: ['@platform/auth', '@platform/content'],
  modules: ['@nuxt/ui'],
  css: ['~/assets/css/main.css'],
  app: {
    head: {
      meta: [
        { property: 'og:site_name', content: siteBrand },
        { property: 'og:type', content: 'website' },
        { name: 'twitter:card', content: 'summary_large_image' }
      ]
    }
  },
  devServer: { port: Number(process.env.NUXT_DEV_PORT || '3003') },
  fonts: { providers: { google: false, googleicons: false, bunny: false, fontshare: false } },
  runtimeConfig: {
    apiBase: process.env.NUXT_API_BASE || 'http://127.0.0.1:8086',
    sealSecret: process.env.NUXT_SEAL_SECRET || 'dev-docs-seal-secret-change-me-0123456789abc',
    downstreamBase: process.env.NUXT_DOWNSTREAM_BASE || 'http://127.0.0.1:8086',
    public: {
      oidcIssuer: process.env.NUXT_PUBLIC_OIDC_ISSUER || 'http://localhost:8081',
      oidcClientId: process.env.NUXT_PUBLIC_OIDC_CLIENT_ID || 'docs-ae-web',
      oidcRedirectUri: process.env.NUXT_PUBLIC_OIDC_REDIRECT_URI || 'http://localhost:3003/auth/callback',
      oidcScopes: process.env.NUXT_PUBLIC_OIDC_SCOPES || 'openid profile email roles offline_access',
      accountUrl: process.env.NUXT_PUBLIC_ACCOUNT_URL || 'http://localhost:3000',
      siteSlug: process.env.NUXT_PUBLIC_SITE_SLUG || 'docs-ae',
      siteBrand,
      siteDomain: process.env.NUXT_PUBLIC_SITE_DOMAIN || 'docs-ae.localhost',
      assetSpace: process.env.NUXT_PUBLIC_ASSET_SPACE || 'default',
      assetNamespace: process.env.NUXT_PUBLIC_ASSET_NAMESPACE || 'default',
      assetProfile: process.env.NUXT_PUBLIC_ASSET_PROFILE || 'docs-default'
    }
  },
  devtools: { enabled: true }
})
