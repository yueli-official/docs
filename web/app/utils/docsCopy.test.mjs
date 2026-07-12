import test from 'node:test'
import assert from 'node:assert/strict'
import { readFileSync } from 'node:fs'
import { fileURLToPath } from 'node:url'
import { dirname, resolve } from 'node:path'

const root = resolve(dirname(fileURLToPath(import.meta.url)), '..')

function readApp(path) {
  return readFileSync(resolve(root, path), 'utf8')
}

test('homepage separates recommended documents from quick links', () => {
  const page = readApp('pages/index.vue')
  assert.match(page, /快速入口/)
  assert.match(page, /推荐文档/)
  assert.doesNotMatch(page, /推荐入口/)
  assert.doesNotMatch(page, /搜索全部文档/)
})

test('manage navigation puts homepage configuration in settings at the end', () => {
  const sidebar = readApp('components/ManageSidebar.vue')
  const page = readApp('pages/manage/home.vue')
  assert.match(sidebar, /label: '设置'/)
  assert.match(sidebar, /i-tabler-settings/)
  assert.ok(sidebar.indexOf("label: '设置'") > sidebar.indexOf("label: '文档'"))
  assert.doesNotMatch(sidebar, /首页布局/)
  assert.match(page, /ManageSettingsLayout/)
  assert.match(page, /首屏文案、快速入口和推荐文档/)
  assert.doesNotMatch(page, /推荐入口/)
})

test('manage homepage layout uses the shared icon picker without preview chrome', () => {
  const page = readApp('pages/manage/home.vue')
  assert.doesNotMatch(page, /首页预览|页脚预览|公开首页预览/)
  assert.match(page, /快速入口配置/)
  assert.match(page, /ManageIconPicker/)
  assert.match(page, /UPopover/)
  assert.match(page, /compact/)
  assert.match(page, /grid-cols-\[2\.75rem_minmax\(0,1fr\)\]/)
  assert.match(page, /quickLinkActionItems/)
  assert.doesNotMatch(page, /group-hover/)
  assert.doesNotMatch(page, /opacity-0/)
  assert.doesNotMatch(page, /<UInput v-model="link\.icon"/)
})

test('docs site exposes a dedicated collections directory', () => {
  const page = readApp('pages/collections.vue')
  const home = readApp('pages/index.vue')
  const header = readApp('components/SiteHeader.vue')

  assert.match(page, /全部文档集/)
  assert.match(page, /filterCollections/)
  assert.match(page, /paginateItems/)
  assert.match(page, /UPagination/)
  assert.match(page, /pageSizeItems/)
  assert.match(home, /to="\/collections"/)
  assert.match(header, /to="\/collections"/)
})

test('docs manage keeps low frequency actions out of the primary editor chrome', () => {
  const editor = readApp('components/manage/DocEditorWorkbench.vue')
  const publishControl = editor.slice(
    editor.indexOf('发布控制'),
    editor.indexOf('公开链接'),
  )
  const settingsStatus = editor.slice(
    editor.indexOf('状态操作'),
    editor.indexOf('SEO 标题'),
  )

  assert.doesNotMatch(publishControl, /label="转回草稿"/)
  assert.doesNotMatch(publishControl, /label="归档"/)
  assert.match(settingsStatus, /label="转回草稿"/)
  assert.match(settingsStatus, /label="归档"/)

  const docsIndex = readApp('pages/manage/docs/index.vue')
  assert.doesNotMatch(docsIndex, /selectedDocMoreItems/)
  assert.doesNotMatch(docsIndex, /label="更多设置"/)
})

test('docs site exposes color mode controls in public and manage chrome', () => {
  const header = readApp('components/SiteHeader.vue')
  const manage = readApp('layouts/manage.vue')
  const css = readApp('assets/css/main.css')

  assert.match(header, /UColorModeButton/)
  assert.match(manage, /UColorModeButton/)
  assert.match(css, /color-scheme: light/)
  assert.match(css, /color-scheme: dark/)
})

test('collection cover upload uses the shared crop dialog before upload', () => {
  const page = readApp('pages/manage/collections.vue')
  assert.match(page, /PlatformImageCropper/)
  assert.match(page, /onCroppedCover/)
})

test('document reader has a fuller reading surface', () => {
  const page = readApp('pages/[collection]/[...slug].vue')
  assert.match(page, /阅读进度/)
  assert.match(page, /本页目录/)
  assert.match(page, /下一篇/)
  assert.match(page, /reading-progress-card/)
  assert.doesNotMatch(page, /下一篇：/)
  assert.match(page, /md:text-\[2\.3125rem\]/)
  assert.doesNotMatch(page, /lg:text-5xl/)
  assert.match(page, /reading-shell/)
})
