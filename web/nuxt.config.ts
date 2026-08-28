const siteBrand = process.env.NUXT_PUBLIC_SITE_BRAND || '月离文档'
const cookieSecure = process.env.NUXT_COOKIE_SECURE === undefined
  ? process.env.NODE_ENV === 'production'
  : process.env.NUXT_COOKIE_SECURE === 'true'

export default defineNuxtConfig({
  extends: [
    '@yueli/identity-nuxt',
    '@yueli/asset-nuxt',
    '@yueli/content-nuxt',
  ],
  modules: ['@nuxt/ui', '@yueli/ui', '@yueli/nuxt-runtime', '@yueli/discovery-nuxt'],
  icon: {
    provider: 'server',
    fallbackToApi: false,
    serverBundle: { collections: ['tabler'] },
    clientBundle: {
      scan: {
        globInclude: [
          'app/**/*.{vue,js,mjs,ts,jsx,tsx}',
          'node_modules/@yueli/**/*.{vue,js,mjs,ts,jsx,tsx}',
        ],
        globExclude: [
          'test/**',
          'tests/**',
          'coverage/**',
          'dist/**',
          '.nuxt/**',
          '.output/**',
          '.*',
        ],
      },
      sizeLimitKb: 256,
    },
  },
  yueliRuntime: {
    defaultTarget: 'docs',
    targets: {
      docs: {
        path: '/',
        ssr: {
          cookies: ['rs_session', 'yueli_guest', '__Host-yueli_guest'],
          headers: ['accept-language', 'user-agent'],
        },
      },
      asset: {
        path: '/asset-api',
        ssr: {
          cookies: ['rs_session', 'yueli_guest', '__Host-yueli_guest'],
          headers: ['accept-language', 'user-agent'],
        },
      },
      identity: {
        path: '/identity-api',
        ssr: {
          cookies: ['rs_session'],
          headers: ['accept-language', 'user-agent'],
        },
      },
    },
  },
  css: ['~/assets/css/main.css'],
  app: {
    head: {
      htmlAttrs: { lang: 'zh-CN' },
      meta: [
        { property: 'og:site_name', content: siteBrand },
        { property: 'og:type', content: 'website' },
      ],
    },
  },
  buildDir: process.env.NUXT_BUILD_DIR || '.nuxt',
  devServer: {
    host: '127.0.0.1',
    port: Number(process.env.NUXT_DEV_PORT || '3003'),
  },
  fonts: {
    providers: {
      google: false,
      googleicons: false,
      bunny: false,
      fontshare: false,
      fontsource: false,
    },
  },
  nitro: {
    esbuild: {
      options: {
        exclude: /node_modules(?!.*(?:@yueli\+|@yueli[\\/]))/,
      },
    },
  },
  runtimeConfig: {
    apiBase: process.env.NUXT_API_BASE || 'http://127.0.0.1:8086',
    assetBase: process.env.NUXT_ASSET_BASE || 'http://127.0.0.1:8082',
    identityBase: process.env.NUXT_IDENTITY_BASE || 'http://127.0.0.1:8081',
    downstreamBase: process.env.NUXT_DOWNSTREAM_BASE || 'http://127.0.0.1:8086',
    cookieSecure,
    authCookieSecure: cookieSecure,
    assetAudience: 'asset-api',
    sealSecret: process.env.NUXT_SEAL_SECRET || '',
    public: {
      oidcIssuer: process.env.NUXT_PUBLIC_OIDC_ISSUER || 'http://localhost:8081',
      oidcClientId: process.env.NUXT_PUBLIC_OIDC_CLIENT_ID || 'docs-main-web',
      oidcRedirectUri:
        process.env.NUXT_PUBLIC_OIDC_REDIRECT_URI || 'http://localhost:3003/auth/callback',
      oidcPostLogoutRedirectUri:
        process.env.NUXT_PUBLIC_OIDC_POST_LOGOUT_REDIRECT_URI || 'http://localhost:3003/',
      oidcScopes:
        process.env.NUXT_PUBLIC_OIDC_SCOPES || 'openid profile email roles offline_access',
      accountUrl: process.env.NUXT_PUBLIC_ACCOUNT_URL || 'http://localhost:3000',
      siteSlug: process.env.NUXT_PUBLIC_SITE_SLUG || 'docs-main',
      siteBrand,
      siteDomain: process.env.NUXT_PUBLIC_SITE_DOMAIN || 'docs-main.localhost',
      assetSpace: process.env.NUXT_PUBLIC_ASSET_SPACE || 'docs',
      assetNamespace: process.env.NUXT_PUBLIC_ASSET_NAMESPACE || 'docs-main',
      assetProfile: process.env.NUXT_PUBLIC_ASSET_PROFILE || 'docs-default',
    },
  },
  devtools: { enabled: true },
})
