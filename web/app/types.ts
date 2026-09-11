import type { DiscoveryProjection } from "@yueli/discovery-nuxt/types";
export interface Collection {
  id: string;
  slug: string;
  title: string;
  description: string;
  coverAssetId?: string;
  coverUrl: string;
  icon: string;
  sortOrder: number;
  releaseFamilyId?: string;
  releaseFamilyName?: string;
  semanticVersion?: string;
  derivedFromCollectionId?: string;
  docCount: number;
}
// CollectionView is the API response shape for collection list/detail endpoints.
export type CollectionView = Collection;
export interface DocBase {
  id: string;
  collectionId: string;
  versionId: string;
  parentId: string;
  slug: string;
  title: string;
  excerpt: string;
  status: string;
  seoTitle?: string;
  seoDescription?: string;
  locale: string;
  translationKey: string;
  badgeText: string;
  badgeIcon: string;
  sortOrder: number;
}
export interface DocTreeNode extends DocBase {
  children?: DocTreeNode[];
}
export interface DocDetail extends DocBase {
  content: string;
  children?: DocDetail[];
}
export type DocNode = DocTreeNode;
export interface CollectionTree {
  collection: Collection;
  tree: DocTreeNode[];
}
export interface CollectionManageTree {
  collection: Collection;
  tree: DocDetail[];
}
export interface PublicDocDetail {
  doc: DocDetail;
}
export interface DocDetailResponse {
  doc: DocDetail;
  discovery?: DiscoveryProjection;
}

export interface URLResolution {
  kind: "canonical" | "alias" | "redirect" | "gone" | "unknown";
  location?: string;
  statusCode?: number;
}

export interface URLResolutionResponse {
  resolution: URLResolution;
}
export interface CollectionList {
  items: Collection[];
}
export interface CollectionReleasesResponse {
  items: Collection[];
}
export interface MeView {
  sub: string;
  authenticated: boolean;
  isAdministrator: boolean;
  roles: string[];
  capabilities: string[];
}
export interface HomeQuickLink {
  id: string;
  title: string;
  description: string;
  icon: string;
  to: string;
  collectionSlug: string;
  sortOrder: number;
  enabled: boolean;
}
export interface HomeConfig {
  quickLinks: HomeQuickLink[];
  featuredCollections: string[];
  homeEyebrow: string;
  homeTitle: string;
  homeSubtitle: string;
  siteTitle: string;
  siteDescription: string;
  supportEmail: string;
  footerTagline: string;
  footerCopyright: string;
}
export interface HomeConfigResponse {
  config: HomeConfig;
}
export interface CollectionVersion {
  id: string;
  collectionId: string;
  key: string;
  label: string;
  status: string;
  isDefault: boolean;
  sortOrder: number;
  sourceVersionId?: string;
}
export interface CollectionVersionsResponse {
  items: CollectionVersion[];
}
export interface CollectionLocale {
  title: string;
  collectionId: string;
  locale: string;
  label: string;
  htmlLang: string;
  direction: "ltr" | "rtl";
  isDefault: boolean;
  enabled: boolean;
  sortOrder: number;
  docCount: number;
}
export interface CollectionLocalesResponse {
  items: CollectionLocale[];
}
export interface CollectionVariantsResponse {
  locales: CollectionLocale[];
  versions: CollectionVersion[];
}
export interface DocumentVariantResolution {
  path: string;
  locale: string;
  version: string;
  translationKey?: string;
  fallback?: "default_locale" | "collection";
}

export type ManageDocStatusKey =
  "all" | "draft" | "published" | "archived" | "issues";
export interface ManageDocListItem {
  id: string;
  collectionId: string;
  collectionSlug: string;
  collectionTitle: string;
  versionId: string;
  versionKey: string;
  versionLabel: string;
  parentId: string;
  parentTitle: string;
  slug: string;
  slugPath: string;
  title: string;
  excerpt: string;
  status: string;
  locale: string;
  badgeText: string;
  badgeIcon: string;
  sortOrder: number;
  updatedAt: string;
}
export interface ManageDocsResponse {
  items: ManageDocListItem[];
  total: number;
  page: number;
  size: number;
  counts: Record<ManageDocStatusKey, number>;
}

export interface DocsImportIssue {
  severity: string;
  code: string;
  message: string;
  path?: string;
}
export interface DocsImportSummary {
  creates: number;
  updates: number;
  archives: number;
  skips: number;
  conflicts: number;
  errors: number;
  images: number;
  warnings: number;
  blocking: boolean;
  issues: DocsImportIssue[];
}
export interface DocsImportBatch {
  id: string;
  collectionId: string;
  versionId: string;
  defaultLocale: string;
  mode: string;
  status: string;
  errorMessage: string;
  summary: DocsImportSummary;
  createdAt: string;
  updatedAt: string;
  completedAt?: string;
}
export interface DocsImportItem {
  id: string;
  locale: string;
  versionKey: string;
  path: string;
  sourceMarkdownPath: string;
  title: string;
  slug: string;
  translationKey: string;
  action: string;
  targetDocId: string;
  issues: DocsImportIssue[];
}
export interface DocsImportUploadResponse {
  batch: DocsImportBatch;
  summary: DocsImportSummary;
}
export interface DocsImportDetailResponse {
  batch: DocsImportBatch;
  items: DocsImportItem[];
}
export interface DocsImportListResponse {
  items: DocsImportBatch[];
  total: number;
  page: number;
  size: number;
}

export interface DashboardTrafficPoint {
  day: string;
  views: number;
  uniqueVisitorDays: number;
}

export interface DashboardTopDocument {
  id: string;
  title: string;
  collectionSlug: string;
  versionKey: string;
  slugPath: string;
  locale: string;
  views: number;
  uniqueVisitorDays: number;
}

export interface DashboardTrafficSource {
  source: string;
  views: number;
}

export interface DashboardSearchQuery {
  query: string;
  searches: number;
  zeroResults: number;
}

export interface DashboardOverview {
  days: number;
  allTimeViews: number;
  allTimeUniqueVisitorDays: number;
  periodViews: number;
  periodUniqueVisitorDays: number;
  previousPeriodViews: number;
  previousUniqueVisitorDays: number;
  periodSearches: number;
  previousPeriodSearches: number;
  zeroResultSearches: number;
  series: DashboardTrafficPoint[];
  topDocuments: DashboardTopDocument[];
  topSources: DashboardTrafficSource[];
  topSearches: DashboardSearchQuery[];
}

export interface CommentView {
  id: string;
  parentId?: string;
  authorName: string;
  avatarUrl?: string;
  content: string;
  createdAt: string;
  replies?: CommentView[];
}

export interface CommentListResponse {
  items: CommentView[];
  total: number;
  page: number;
  size: number;
}

export type CommentStatus = "approved" | "pending" | "spam" | "trash";

export interface CommentAdminView {
  id: string;
  documentId: string;
  documentTitle: string;
  collectionSlug: string;
  versionKey: string;
  slugPath: string;
  locale: string;
  parentId?: string;
  authorName: string;
  avatarUrl?: string;
  userSub: string;
  content: string;
  status: CommentStatus;
  createdAt: string;
}

export interface ManageCommentsResponse {
  items: CommentAdminView[];
  total: number;
  page: number;
  size: number;
}
