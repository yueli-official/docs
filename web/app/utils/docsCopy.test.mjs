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
});

test("manage navigation puts homepage configuration in settings at the end", () => {
  const layout = readApp("layouts/manage.vue");
  const page = readApp("pages/manage/home.vue");
  assert.match(layout, /YAdminShell/);
  assert.match(layout, /label: "站点设置"/);
  assert.match(layout, /i-tabler-settings/);
  assert.ok(
    layout.indexOf('label: "站点设置"') > layout.indexOf('label: "文档"'),
  );
  assert.doesNotMatch(layout, /首页布局/);
  assert.match(page, /SettingsLayout/);
  assert.match(page, /:show-header="false"/);
  assert.match(page, /首屏文案、快速入口和推荐文档/);
  assert.doesNotMatch(page, /推荐入口/);
});

test("manage sidebar owns one context menu and one account footer", () => {
  const layout = readApp("layouts/manage.vue");
  assert.match(layout, /workspaceMenuItems/);
  assert.match(layout, /打开文档站/);
  assert.match(layout, /#brand="\{ collapsed \}"/);
  assert.match(layout, /#sidebar-footer="\{ collapsed \}"/);
  assert.match(layout, /show-appearance/);
  assert.match(layout, /trigger-mode/);
  assert.doesNotMatch(layout, /secondaryNavigation/);
  assert.doesNotMatch(layout, /UColorModeButton/);
  assert.doesNotMatch(layout, /:collapsible="false"/);
});

test("manage homepage layout uses the shared icon picker without preview chrome", () => {
  const page = readApp("pages/manage/home.vue");
  assert.doesNotMatch(page, /首页预览|页脚预览|公开首页预览/);
  assert.match(page, /快速入口配置/);
  assert.match(page, /ManageIconPicker/);
  assert.match(page, /UPopover/);
  assert.match(page, /compact/);
  assert.match(page, /grid-cols-\[2\.75rem_minmax\(0,1fr\)\]/);
  assert.match(page, /quickLinkActionItems/);
  assert.doesNotMatch(page, /group-hover/);
  assert.doesNotMatch(page, /opacity-0/);
  assert.doesNotMatch(page, /<UInput v-model="link\.icon"/);
});

test("docs site exposes a dedicated collections directory", () => {
  const page = readApp("pages/collections.vue");
  const home = readApp("pages/index.vue");
  const header = readApp("components/SiteHeader.vue");

  assert.match(page, /全部文档集/);
  assert.match(page, /filterCollections/);
  assert.match(page, /paginateItems/);
  assert.match(page, /UPagination/);
  assert.match(page, /pageSizeItems/);
  assert.match(home, /to="\/collections"/);
  assert.match(header, /to="\/collections"/);
});

test("docs manage keeps low frequency actions out of the primary editor chrome", () => {
  const editor = readApp("components/manage/DocEditorWorkbench.vue");
  const publishControl = editor.slice(
    editor.indexOf("发布控制"),
    editor.indexOf("公开链接"),
  );
  const settingsStatus = editor.slice(
    editor.indexOf("状态操作"),
    editor.indexOf("SEO 标题"),
  );

  assert.doesNotMatch(publishControl, /label="转回草稿"/);
  assert.doesNotMatch(publishControl, /label="归档"/);
  assert.match(settingsStatus, /label="转回草稿"/);
  assert.match(settingsStatus, /label="归档"/);

  const docsIndex = readApp("pages/manage/docs/index.vue");
  assert.doesNotMatch(docsIndex, /selectedDocMoreItems/);
  assert.doesNotMatch(docsIndex, /label="更多设置"/);
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

test("docs compiles Tailwind utilities from the public UI package", () => {
  const css = readApp("assets/css/main.css");
  assert.match(css, /@import "@yueli\/ui\/tailwind\.css";/);
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
  assert.match(page, /PlatformImageCropper/);
  assert.match(page, /onCroppedCover/);
});

test("document reader has a fuller reading surface", () => {
  const page = readApp("pages/[collection]/[...slug].vue");
  assert.match(page, /阅读进度/);
  assert.match(page, /本页目录/);
  assert.match(page, /下一篇/);
  assert.match(page, /reading-progress-card/);
  assert.doesNotMatch(page, /下一篇：/);
  assert.match(page, /md:text-\[2\.3125rem\]/);
  assert.doesNotMatch(page, /lg:text-5xl/);
  assert.match(page, /reading-shell/);
});
