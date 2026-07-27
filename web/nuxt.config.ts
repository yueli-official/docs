// Nuxt 4 config for the docs-site app (consumer of the docs service).
const siteBrand = process.env.NUXT_PUBLIC_SITE_BRAND || "文档库";

export default defineNuxtConfig({
  extends: [
    "@yueli/identity-nuxt",
    "@platform/site",
    "@platform/manage",
    "@platform/asset",
    "@platform/content",
  ],
  // Public raw SFC utilities are registered by the Tailwind source import in main.css.
  modules: ["@nuxt/ui", "@yueli/ui", "@yueli/discovery-nuxt"],
  css: ["~/assets/css/main.css"],
  buildDir: process.env.NUXT_BUILD_DIR || ".nuxt",
  devServer: { port: Number(process.env.NUXT_DEV_PORT || "3003") },
  runtimeConfig: {
    apiBase: process.env.NUXT_API_BASE || "http://127.0.0.1:8086",
    sealSecret:
      process.env.NUXT_SEAL_SECRET ||
      "dev-docs-seal-secret-change-me-0123456789abc",
    downstreamBase: process.env.NUXT_DOWNSTREAM_BASE || "http://127.0.0.1:8086",
    public: {
      oidcIssuer:
        process.env.NUXT_PUBLIC_OIDC_ISSUER || "http://localhost:8081",
      oidcClientId: process.env.NUXT_PUBLIC_OIDC_CLIENT_ID || "docs-main-web",
      oidcRedirectUri:
        process.env.NUXT_PUBLIC_OIDC_REDIRECT_URI ||
        "http://localhost:3003/auth/callback",
      oidcScopes:
        process.env.NUXT_PUBLIC_OIDC_SCOPES ||
        "openid profile email roles offline_access",
      accountUrl:
        process.env.NUXT_PUBLIC_ACCOUNT_URL || "http://localhost:3000",
      siteSlug: process.env.NUXT_PUBLIC_SITE_SLUG || "docs-main",
      siteBrand,
      siteDomain: process.env.NUXT_PUBLIC_SITE_DOMAIN || "docs-main.localhost",
      assetSpace: process.env.NUXT_PUBLIC_ASSET_SPACE || "default",
      assetNamespace: process.env.NUXT_PUBLIC_ASSET_NAMESPACE || "default",
      assetProfile: process.env.NUXT_PUBLIC_ASSET_PROFILE || "docs-default",
    },
  },
  devtools: { enabled: true },
});
