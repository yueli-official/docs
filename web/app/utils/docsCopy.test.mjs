import test from "node:test";
import assert from "node:assert/strict";
import { readFileSync } from "node:fs";
import { fileURLToPath } from "node:url";
import { dirname, resolve } from "node:path";

const root = resolve(dirname(fileURLToPath(import.meta.url)), "..");

function readApp(path) {
  return readFileSync(resolve(root, path), "utf8");
}

test("homepage separates recommended documents from quick links", () => {
  const page = readApp("pages/index.vue");
  assert.match(page, /快速入口/);
  assert.match(page, /推荐文档/);
  assert.doesNotMatch(page, /推荐入口/);
  assert.doesNotMatch(page, /搜索全部文档/);
  assert.doesNotMatch(page, /homeConfig\.homeEyebrow/);
  assert.doesNotMatch(
    page,
    /collectionStats|stats\.collectionCount|stats\.docCount/,
  );
  assert.doesNotMatch(page, /搜索标题、摘要或正文/);
});

test("document variants live in the header and use transient shared feedback", () => {
  const switcher = readApp("components/DocumentVariantSwitcher.vue");
  const collection = readApp("pages/[collection]/index.vue");
  const reader = readApp("pages/[collection]/[...slug].vue");

  assert.match(switcher, /createDocsNotifier\(useToast\(\)\)/);
  assert.match(switcher, /duration:\s*3_000/);
  assert.match(switcher, /data-document-variant-switcher/);
  assert.match(
    switcher,
    /resolution\.fallback === "collection" && props\.translationKey/,
  );
  assert.doesNotMatch(switcher, /docs-variant-notice|<UAlert/);
  assert.match(collection, /sm:justify-between[\s\S]*DocumentVariantSwitcher/);
  assert.match(reader, /sm:justify-between[\s\S]*DocumentVariantSwitcher/);
  assert.doesNotMatch(reader, /分享文档|ContentShareActions/);
});

test("manage navigation keeps one settings entry and moves its workflows into tabs", () => {
  const layout = readApp("layouts/manage.vue");
  const page = readApp("pages/manage/home.vue");
  assert.match(layout, /YAdminConsoleLayout/);
  assert.match(layout, /<template>\s*<YAdminConsoleLayout/);
  assert.doesNotMatch(
    layout,
    /<template>\s*<ClientOnly>[\s\S]*?<YAdminConsoleLayout/,
  );
  assert.doesNotMatch(layout, /正在打开[^\n]{0,16}控制台/);
  assert.match(layout, /search: "搜索控制台"/);
  assert.match(layout, /label: "站点设置"/);
  assert.match(layout, /i-tabler-settings/);
  assert.ok(
    layout.indexOf('label: "站点设置"') > layout.indexOf('label: "文档"'),
  );
  const settingsNavigation = layout.slice(
    layout.indexOf('label: "站点设置"'),
    layout.indexOf('can("docs.asset_settings.manage")'),
  );
  assert.match(settingsNavigation, /to: "\/manage\/home"/);
  assert.doesNotMatch(settingsNavigation, /children:|type: "trigger"/);
  assert.doesNotMatch(layout, /站点设置 · (?:首页|页脚|基础)/);
  assert.equal(layout.match(/label: "站点设置"/g)?.length, 2);
  assert.match(
    page,
    /import \{[\s\S]{0,200}TabbedSurface[\s\S]{0,100}from "@yueli\/ui\/admin"/,
  );
  assert.match(page, /<TabbedSurface/);
  assert.match(page, /navigation-label="站点设置"/);
  assert.match(page, /value: "home"/);
  assert.match(page, /value: "footer"/);
  assert.match(page, /value: "site"/);
  assert.doesNotMatch(page, /SettingsLayout/);
  assert.ok(
    layout.indexOf('label: "文档"') < layout.indexOf('label: "评论"') &&
      layout.indexOf('label: "评论"') < layout.indexOf('label: "文档集"'),
    "评论入口应紧跟文档入口",
  );
});

test("Identity BFF downstream configuration stays origin-only", () => {
  const config = readApp("../nuxt.config.ts");
  assert.match(
    config,
    /downstreamBase:\s*process\.env\.NUXT_DOWNSTREAM_BASE\s*\|\|\s*['"]http:\/\/127\.0\.0\.1:8086['"]/,
  );
  assert.doesNotMatch(
    config,
    /NUXT_DOWNSTREAM_BASE\s*\|\|\s*['"][^'"]+\/api\/v1['"]/,
  );
});

test("manage sidebar owns one direct brand link and one account footer", () => {
  const layout = readApp("layouts/manage.vue");
  assert.match(layout, /brand-to="\/"/);
  assert.match(layout, /brand-icon="i-tabler-book"/);
  assert.match(layout, /#account="\{ collapsed \}"/);
  assert.match(layout, /show-appearance/);
  assert.match(layout, /trigger-mode/);
  assert.match(layout, /:current-label="currentLabel"/);
  assert.match(layout, /:immersive="isDocumentEditor"/);
  assert.match(layout, /route\.path\.startsWith\("\/manage\/docs\/"\)/);
  assert.doesNotMatch(layout, /workspaceMenuItems|UDropdownMenu/);
  assert.doesNotMatch(layout, /secondaryNavigation/);
  assert.doesNotMatch(layout, /UColorModeButton/);
});

test("document editor uses the shared immersive console seam", () => {
  const editor = readApp("components/manage/DocEditorWorkbench.vue");
  assert.match(editor, /data-docs-editor-commandbar/);
  assert.match(editor, /<EditorCommandBar/);
  assert.match(editor, /v-model:immersive="immersiveCollaboration"/);
  assert.doesNotMatch(editor, /-mt-(?:5|8|10)|-mx-(?:5|8|10)/);
  assert.match(editor, /showValidationError\("请填写标题"\)/);
  assert.match(editor, /title: "操作失败"/);
  assert.doesNotMatch(editor, /validationError|请完善文档信息/);
});

test("manage gate resolves Docs capabilities before the server render", () => {
  const gate = readApp("middleware/manage-gate.global.ts");
  const refresh = gate.indexOf("await refreshMe()");
  const serverReturn = gate.indexOf("if (import.meta.server) return");
  assert.ok(refresh >= 0, "manage gate must load Docs effective access");
  assert.ok(
    serverReturn > refresh,
    "server middleware must populate Docs access before rendering navigation",
  );
});

test("manage dashboard uses Blog-style reading analytics instead of maintenance activity", () => {
  const layout = readApp("layouts/manage.vue");
  const page = readApp("pages/manage/index.vue");

  assert.match(layout, /const searchGroups = computed/);
  assert.match(layout, /can\("docs\.document\.create"\)/);
  assert.match(layout, /can\("docs\.import\.manage"\)/);
  assert.match(layout, /isAdministrator\.value/);
  assert.match(page, /DashboardTrendChart/);
  assert.match(page, /\/api\/v1\/dashboard\/overview/);
  assert.match(page, /累计浏览/);
  assert.match(page, /近.*天浏览/);
  assert.match(page, /浏览深度/);
  assert.match(page, /搜索次数/);
  assert.match(page, /浏览趋势/);
  assert.match(page, /访客趋势/);
  assert.match(page, /热门文档/);
  assert.match(page, /流量来源/);
  assert.match(page, /热门搜索/);
  assert.match(page, /publicDocumentLink/);
  assert.match(page, /locale:\s*document\.locale/);
  assert.match(page, /version:\s*document\.versionKey/);
  assert.doesNotMatch(page, /工作区状态|可用操作|快捷操作/);
  assert.doesNotMatch(page, /最近更新|文档集|最近导入|草稿|待完善/);
  assert.doesNotMatch(page, /aria-label="关键指标"/);
  assert.doesNotMatch(page, /服务状态/);
});

test("manage homepage layout uses the shared icon picker without preview chrome", () => {
  const page = readApp("pages/manage/home.vue");
  assert.doesNotMatch(page, /首页预览|页脚预览|公开首页预览/);
  assert.match(page, /快速入口配置/);
  assert.match(page, /AdminIconPicker/);
  assert.match(page, /<SettingSection\s+title="首页文案"/);
  assert.match(page, /<SettingSection\s+title="快速入口配置"/);
  assert.match(page, /<SettingSection\s+title="推荐文档"/);
  assert.match(page, /label="添加文档"/);
  assert.match(page, /v-model:open="featuredPickerOpen"/);
  assert.match(page, /placeholder="搜索文档集"/);
  assert.match(page, /featuredCollectionActionItems/);
  assert.doesNotMatch(page, /<SettingSection\s+title="展示顺序"|toggleFeatured/);
  assert.match(page, /<SettingSection\s+title="页脚内容"/);
  assert.match(page, /<SettingSection\s+title="站点基础"/);
  assert.doesNotMatch(page, /配置首页入口卡片的标题、路径和展示状态/);
  assert.doesNotMatch(page, /选择首页展示的文档集，右侧顺序按选择顺序生成/);
  assert.doesNotMatch(page, /用于所有文档页面底部的品牌说明与联系信息/);
  assert.doesNotMatch(page, /这些字段用于导航品牌、SEO 描述和支持入口/);
  assert.match(page, /UPopover/);
  assert.match(page, /compact/);
  assert.match(page, /grid-cols-\[2\.75rem_minmax\(0,1fr\)\]/);
  assert.match(page, /quickLinkActionItems/);
  assert.doesNotMatch(page, /group-hover/);
  assert.doesNotMatch(page, /opacity-0/);
  assert.doesNotMatch(page, /<UInput v-model="link\.icon"/);
});

test("collection settings initialize a single release and derive translation drafts", () => {
  const panel = readApp("components/manage/CollectionVariantsManager.vue");
  assert.match(panel, /设置当前版本/);
  assert.match(
    panel,
    /\/manage\/collections\/\$\{props\.collection\.id\}\/release/,
  );
  assert.match(panel, /复制语言/);
  assert.match(panel, /\/locales\/clone/);
  assert.match(panel, /sourceLocale/);
  assert.match(panel, /targetLocale/);
  assert.match(panel, /复制后的文档全部为草稿/);
});

test("collection creation captures its initial language and semantic version", () => {
  const page = readApp("pages/manage/collections.vue");
  assert.match(page, /label="默认语言"/);
  assert.match(page, /label="版本号"/);
  assert.match(page, /defaultLocale/);
  assert.match(page, /semanticVersion/);
});

test("docs site exposes a dedicated collections directory", () => {
  const page = readApp("pages/collections.vue");
  const home = readApp("pages/index.vue");
  const header = readApp("components/SiteHeader.vue");

  assert.match(page, />\s*文档集\s*</);
  assert.match(page, /filterCollections/);
  assert.match(page, /paginateItems/);
  assert.match(page, /UPagination/);
  assert.match(page, /pageSizeItems/);
  assert.match(page, /placeholder="搜索文档集"/);
  assert.doesNotMatch(page, /面包屑|Collection directory/);
  assert.doesNotMatch(page, /按标题、路径或说明查找文档集/);
  assert.doesNotMatch(
    page,
    /collectionStats|stats\.collectionCount|stats\.docCount/,
  );
  assert.match(home, /to="\/collections"/);
  assert.match(header, /to="\/collections"/);
});

test("document editor exposes frequent properties and lifecycle actions", () => {
  const editor = readApp("components/manage/DocEditorWorkbench.vue");

  assert.match(editor, /import \{ EditorInspector \} from "@yueli\/ui\/admin"/);
  assert.match(editor, /<EditorInspector/);
  assert.match(editor, /data-docs-editor-workspace/);
  assert.match(editor, /data-docs-inspector-content/);
  assert.match(editor, /data-docs-inspector-organization/);
  assert.match(editor, /data-docs-inspector-seo/);
  assert.match(editor, /xl:pr-\[27rem\]/);
  assert.match(
    editor,
    /label: "内容", value: "content", icon: "i-tabler-stack-2"/,
  );
  assert.match(
    editor,
    /label: "组织", value: "organization", icon: "i-tabler-arrows-sort"/,
  );
  assert.match(editor, /label: "SEO", value: "seo", icon: "i-tabler-search"/);
  const inspector = editor.slice(
    editor.indexOf("<EditorInspector"),
    editor.indexOf("</EditorInspector>"),
  );
  const inspectorFields = [
    ...inspector.matchAll(/<U(?:Input|Textarea|Select|SelectMenu)\b[\s\S]*?>/g),
  ].map((match) => match[0]);
  assert.ok(inspectorFields.length > 0);
  for (const field of inspectorFields)
    assert.match(field, /class="[^"]*w-full/);
  assert.match(editor, /data-docs-lifecycle-actions/);
  assert.match(editor, /label="摘要"/);
  assert.match(editor, /label="文档集"/);
  assert.match(editor, /label="父文档"/);
  assert.doesNotMatch(editor, /label="版本"/);
  assert.match(editor, /label="语言"/);
  assert.match(editor, /label: "下架"/);
  assert.match(editor, /label: "归档"/);
  assert.doesNotMatch(editor, /data-docs-editor-properties/);
  assert.doesNotMatch(editor, /<USlideover/);
  assert.doesNotMatch(editor, /状态操作|低频生命周期动作/);
  assert.doesNotMatch(editor, /<h2[^>]*>摘要与链接<\/h2>/);
  assert.doesNotMatch(editor, /<h2[^>]*>归属与路径<\/h2>/);

  const docsIndex = readApp("pages/manage/docs/index.vue");
  assert.doesNotMatch(docsIndex, /selectedDocMoreItems/);
  assert.doesNotMatch(docsIndex, /label="更多设置"/);
});

test("document editor wires image upload and clears the exact local draft after save", () => {
  const editor = readApp("components/manage/DocEditorWorkbench.vue");

  assert.match(editor, /:image-uploader="uploadInlineImage"/);
  assert.match(editor, /ref="editorComp"/);
  assert.match(editor, /editorComp\.value\?\.markSaved\(\)/);
  assert.match(editor, /draftInstanceId/);
  assert.doesNotMatch(editor, /:draft-entity-id="isNew \? 'new' : docId"/);
});

test("document inspector uses SEO and full-width fields", () => {
  const editor = readApp("components/manage/DocEditorWorkbench.vue");
  assert.match(editor, /label: "SEO", value: "seo", icon: "i-tabler-search"/);
  const inspector = editor.slice(
    editor.indexOf("<EditorInspector"),
    editor.indexOf("</EditorInspector>"),
  );
  const fields = [
    ...inspector.matchAll(/<U(?:Input|Textarea|Select|SelectMenu)\b[\s\S]*?>/g),
  ].map((match) => match[0]);
  assert.ok(fields.length > 0);
  for (const field of fields) assert.match(field, /class="[^"]*w-full/);
});

test("quick edit uses a non-empty root parent sentinel for Nuxt UI comboboxes", () => {
  const modal = readApp("components/manage/DocQuickEditModal.vue");
  assert.match(modal, /const ROOT_PARENT_VALUE = "__root__"/);
  assert.match(modal, /v-model="parentSelection"/);
  assert.doesNotMatch(modal, /label: "顶级文档", value: ""/);
});

test("document management uses Nuxt UI Table and keeps bulk actions in the toolbar", () => {
  const page = readApp("pages/manage/docs/index.vue");
  assert.match(page, /import type \{ TableColumn \} from "@nuxt\/ui"/);
  assert.match(page, /<UTable/);
  assert.match(page, /v-model:row-selection="rowSelection"/);
  assert.match(page, /data-docs-bulk-actions/);
  assert.match(page, /<CollectionTableToolbar/);
  assert.match(page, /<template #filters>/);
  assert.match(page, /<template #utilities>/);
  assert.match(page, /<template #selection>/);
  assert.match(page, /v-for="control in collectionFilterControls"/);
  assert.match(page, /aria-label="文档筛选条件"/);
  assert.doesNotMatch(page, /常用筛选|更多条件/);
  assert.match(page, /header: sortableHeader\("标题", "title"\)/);
  assert.match(page, /header: sortableHeader\("路径 \/ 父级", "path"\)/);
  assert.match(page, /header: sortableHeader\("更新", "updatedAt"\)/);
  assert.doesNotMatch(page, /<template #context>/);
  assert.doesNotMatch(page, /row-layout-class=/);
  assert.doesNotMatch(page, /data-collection-modebar/);
});

test("document and tree actions follow effective capabilities", () => {
  const page = readApp("pages/manage/docs/index.vue");
  const tree = readApp("components/DocTreeAdmin.vue");
  const collections = readApp("pages/manage/collections.vue");

  assert.match(page, /can\("docs\.document\.read"\)/);
  assert.match(page, /can\("docs\.document\.create"\)/);
  assert.match(page, /can\("docs\.document\.update"\)/);
  assert.match(page, /can\("docs\.document\.publish"\)/);
  assert.match(page, /can\("docs\.document\.archive"\)/);
  assert.match(page, /can\("docs\.document\.delete_permanently"\)/);
  assert.match(page, /const bulkItems = computed/);
  assert.match(page, /label: "修改语言"/);
  assert.match(page, /v-model:open="bulkLocaleOpen"/);
  assert.match(page, /confirmBulkLocale/);
  assert.match(page, /const docColumns = computed/);
  assert.match(page, /没有文档读取权限/);
  assert.match(tree, /canEdit\?: boolean/);
  assert.match(tree, /:draggable="canMove"/);
  assert.match(tree, /v-if="canMove \|\| canDelete"/);
  assert.match(collections, /can\("docs\.collection\.manage"\)/);
  assert.match(collections, /没有文档集管理权限/);
});

test("docs import is a recoverable capability-backed workflow", () => {
  const start = readApp("pages/manage/import.vue");
  const detail = readApp("pages/manage/import/[importId].vue");

  assert.match(start, /can\("docs\.import\.manage"\)/);
  assert.match(start, /DocsImportListResponse/);
  assert.match(start, /\/api\/v1\/imports\/docs/);
  assert.match(start, /最近导入/);
  assert.doesNotMatch(start, /confirmError|title="确认导入失败"/);
  assert.match(start, /title: "导入失败"[\s\S]*duration: 0/);
  assert.match(start, /importStatusMeta/);
  assert.match(start, /importModeLabel/);
  assert.match(start, /docs\.json/);
  assert.match(start, /body\.append\("collection", collectionSlug\.value\)/);
  assert.match(start, /body\.append\("defaultLocale", defaultLocale\.value\.trim\(\)\)/);
  assert.match(start, /body\.append\("mode", importMode\.value\)/);
  assert.match(start, /label="目标文档集"/);
  assert.match(start, /class="grid items-start gap-4/);
  assert.doesNotMatch(start, /xl:top-24/);
  assert.doesNotMatch(start, /ZIP only/);
  assert.match(detail, /can\("docs\.import\.manage"\)/);
  assert.match(detail, /batch\.errorMessage/);
  assert.match(detail, /importStatusMeta/);
  assert.match(detail, /importModeLabel/);
  assert.doesNotMatch(detail, /"yes"\s*:\s*"no"/);
});

test("docs settings and authorization consume effective capabilities", () => {
  const home = readApp("pages/manage/home.vue");
  const assets = readApp("pages/manage/assets.vue");
  const authorization = readApp("pages/manage/authorization.vue");

  assert.match(home, /can\("docs\.site_settings\.manage"\)/);
  assert.match(home, /没有站点设置权限/);
  assert.match(assets, /AssetPolicyPage/);
  assert.match(assets, /expected-namespace="docs"/);
  assert.match(assets, /:can-edit="canEditAssets"/);
  assert.doesNotMatch(assets, /ManageAssetSettings/);
  assert.match(authorization, /<ManagePage/);
  assert.match(authorization, /<TabbedSurface/);
  assert.match(authorization, /用户管理/);
  assert.match(authorization, /\/api\/v1\/authorization\/manage\/console/);
  assert.match(authorization, /\/api\/v1\/authorization\/manage\/grants/);
  assert.match(authorization, /data-authorization-role-row/);
  assert.match(authorization, /title="撤销角色"/);
  assert.match(authorization, /label="确认撤销"/);
  assert.match(authorization, /docs\.administrator_grant_protected/);
  assert.doesNotMatch(authorization, /window\.confirm/);
  assert.match(authorization, /isAdministrator/);
  assert.match(authorization, /自定义角色/);
  assert.match(authorization, /create-role/);
  assert.match(authorization, /retire-role/);
  assert.match(authorization, /\/preview/);
  assert.match(authorization, /确认后才会创建服务端 draft/);
  assert.match(authorization, /管理员能力受保护/);
  assert.match(authorization, /拒绝申请时请填写原因/);
  assert.doesNotMatch(authorization, /<main class=/);
});

test("docs compiles Tailwind utilities from the public UI package", () => {
  const css = readApp("assets/css/main.css");
  assert.match(css, /@import "@yueli\/ui\/tailwind\.css";/);
});

test("docs form fields use one border instead of outline and ring focus layers", () => {
  const config = readApp("app.config.ts");
  const css = readApp("assets/css/main.css");

  assert.match(config, /createUiPreset/);
  assert.doesNotMatch(config, /docs-field-border/);
  assert.doesNotMatch(css, /\.docs-field-border/);
});

test("docs site exposes color mode controls in public and manage chrome", () => {
  const header = readApp("components/SiteHeader.vue");
  const manage = readApp("layouts/manage.vue");
  const css = readApp("assets/css/main.css");

  assert.match(header, /UColorModeButton/);
  assert.match(manage, /show-appearance/);
  assert.doesNotMatch(manage, /UColorModeButton/);
  assert.match(css, /color-scheme: light/);
  assert.match(css, /color-scheme: dark/);
});

test("collection cover upload uses the shared crop dialog before upload", () => {
  const page = readApp("pages/manage/collections.vue");
  const visualAsset = readApp("components/ManageVisualAssetField.vue");
  const assetDeclaration = JSON.parse(
    readFileSync(resolve(root, "../../deploy/asset.consumer.json"), "utf8"),
  );
  const coverProfile = assetDeclaration.profiles.find(
    (profile) => profile.key === "docs-collection-cover",
  );

  assert.match(page, /AssetImageCropper/);
  assert.match(page, /onCroppedCover/);
  assert.match(page, /CollectionTableToolbar/);
  assert.doesNotMatch(page, /CollectionToolbar,/);
  assert.doesNotMatch(page, /路径预览|危险操作/);
  assert.match(page, /:aspect-ratio="1"/);
  assert.match(page, /:output-width="256"/);
  assert.match(page, /:output-height="256"/);
  assert.match(page, /output-type="image\/webp"/);
  assert.match(visualAsset, /aspect-square/);
  assert.match(visualAsset, /1:1 · 256 × 256 · WebP/);
  assert.match(visualAsset, /data-cover-column/);
  assert.match(visualAsset, /data-icon-column/);
  assert.match(visualAsset, /AdminIconPicker/);
  assert.match(visualAsset, /i-tabler-info-circle/);
  assert.ok(
    visualAsset.indexOf("data-cover-preview") <
      visualAsset.indexOf("data-icon-column"),
  );
  assert.doesNotMatch(visualAsset, /data-cover-actions|data-cover-link|封面链接/);
  assert.doesNotMatch(visualAsset, /aspect-\[16\/9\]/);
  assert.doesNotMatch(visualAsset, /:help="coverSpec"/);
  assert.deepEqual(coverProfile.variants, [
    {
      key: "cover",
      width: 256,
      height: 256,
      mode: "fill",
      format: "webp",
      quality: 85,
      visibility: "public",
      metadataPolicy: "strip",
    },
  ]);

  const footer = page.slice(page.indexOf("<template #footer>"));
  assert.ok(
    footer.indexOf("删除文档集") < footer.indexOf('label="取消"') &&
      footer.indexOf('label="取消"') <
        footer.indexOf(":label=\"current ? '保存' : '创建'\""),
  );
});

test("document reader has a fuller reading surface", () => {
  const page = readApp("pages/[collection]/[...slug].vue");
  const layout = readApp("layouts/collection.vue");
  assert.match(page, /ReadingTableOfContents/);
  assert.match(page, /can\("docs\.document\.update"\)/);
  assert.match(page, /docManageRoute\(collectionSlug\.value, docPath\.value\)/);
  assert.match(page, /label="编辑"/);
  assert.match(page, /:max-level="3"/);
  assert.match(page, /xl:grid-cols-\[minmax\(0,3fr\)_minmax\(0,1fr\)\]/);
  assert.doesNotMatch(page, /_240px/);
  assert.match(page, /本页目录/);
  assert.match(layout, /max-w-\[1400px\]/);
  assert.doesNotMatch(
    page,
    /阅读进度|reading-progress-card|sectionLabel|个小节|面包屑|breadcrumbs|siteBrand/,
  );
  assert.doesNotMatch(page, /下一篇：/);
  assert.match(page, /md:text-\[2\.3125rem\]/);
  assert.doesNotMatch(page, /lg:text-5xl/);
  assert.match(page, /reading-shell/);
  assert.doesNotMatch(page, /ContentShareActions|aria-label="分享文档"/);
});

test("docs comments provide a reader thread and an administrator moderation queue", () => {
  const reader = readApp("pages/[collection]/[...slug].vue");
  const comments = readApp("components/DocumentComments.vue");
  const chapterNavigation = readApp("components/DocNav.vue");
  const manage = readApp("pages/manage/comments.vue");
  const layout = readApp("layouts/manage.vue");

  assert.match(reader, /<DocumentComments/);
  assert.match(reader, /<DocNav[\s\S]*<DocumentComments/);
  assert.match(comments, /data-document-comments/);
  assert.match(comments, /PublicCommentThread/);
  assert.match(comments, /sortOrder: order\.value/);
  assert.match(comments, /登录后评论/);
  assert.match(comments, /\/api\/v1\/docs\/\$\{props\.documentId\}\/comments/);
  assert.doesNotMatch(comments, /border-t border-default/);
  assert.doesNotMatch(chapterNavigation, /border-t border-default/);
  assert.doesNotMatch(comments, /data-comment-thread|DocumentCommentForm/);
  assert.doesNotMatch(comments, /isMember|label="成员"/);
  assert.match(manage, /<ManagePage/);
  assert.match(manage, /CommentModerationCollection/);
  assert.match(manage, /批量操作/);
  assert.match(manage, /icon: "i-tabler-file-text"/);
  assert.match(manage, /actions: rowActionItems\(comment\)/);
  assert.doesNotMatch(manage, /<CollectionTableToolbar|<CollectionSortHeader/);
  assert.match(manage, /sortBy:\s*"createdAt"/);
  assert.match(manage, /sortOrder:\s*sortOrder\.value/);
  assert.match(manage, /useMinimumLoading/);
  assert.doesNotMatch(manage, /<UButton[\s\S]{0,160}label="标记为垃圾"/);
  assert.match(manage, /title="永久删除评论"/);
  assert.match(manage, /label: "移入回收站"/);
  assert.match(manage, /label: "永久删除"/);
  assert.match(manage, /lifecycleChange: changeLifecycle/);
  assert.match(manage, /emptyTrash/);
  assert.match(layout, /to: "\/manage\/comments"/);
  assert.match(layout, /label: "评论"/);
});
