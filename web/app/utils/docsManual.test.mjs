import test from 'node:test'
import assert from 'node:assert/strict'
import {
  collectionCountLabel,
  collectionStats,
  filterCollections,
  homeFeaturedCollections,
  homeQuickLinks,
  paginateItems,
  rootSectionCountLabel
} from './docsManual.mjs'

const collections = [
  { title: 'Sapphire Guide', slug: 'sapphire', description: 'API and setup manuals' },
  { title: 'Billing Reference', slug: 'billing', description: 'Payments and invoices' },
  { title: 'Account Admin', slug: 'account-admin', description: 'User roles and SSO' }
]

test('filterCollections matches title and description case-insensitively', () => {
  assert.deepEqual(filterCollections(collections, 'api'), [collections[0]])
  assert.deepEqual(filterCollections(collections, 'BILL'), [collections[1]])
  assert.deepEqual(filterCollections(collections, 'sso'), [collections[2]])
})

test('filterCollections returns all collections for blank query', () => {
  assert.deepEqual(filterCollections(collections, '   '), collections)
})

test('filterCollections matches collection slug', () => {
  assert.deepEqual(filterCollections(collections, 'account-admin'), [collections[2]])
})

test('paginateItems returns safe page metadata', () => {
  const result = paginateItems([1, 2, 3, 4, 5], 3, 2)
  assert.deepEqual(result, {
    items: [5],
    page: 3,
    pageSize: 2,
    total: 5,
    totalPages: 3,
    start: 5,
    end: 5
  })
  assert.equal(paginateItems([1, 2], 99, 10).page, 1)
})

test('collectionCountLabel avoids emphasizing unknown or zero counts', () => {
  assert.equal(collectionCountLabel(12), '12 篇文档')
  assert.equal(collectionCountLabel(0), '文档集')
  assert.equal(collectionCountLabel(undefined), '文档集')
})

test('rootSectionCountLabel describes visible root sections', () => {
  assert.equal(rootSectionCountLabel(3), '3 个章节')
  assert.equal(rootSectionCountLabel(0), '暂无公开章节')
})

test('collectionStats summarizes visible collections and published docs', () => {
  assert.deepEqual(collectionStats([
    { docCount: 5 },
    { docCount: 0 },
    { docCount: 3 }
  ]), { collectionCount: 3, docCount: 8 })
})

test('homeQuickLinks uses backend configured entries instead of keyword matching', () => {
  const links = homeQuickLinks({
    quickLinks: [
      { id: 'ops', title: '运营手册', description: '日常运营入口', icon: 'i-tabler-compass', collectionSlug: 'admin' },
      { id: 'external', title: '状态页', description: '查看服务状态', icon: 'i-tabler-activity', to: 'https://status.example.com' },
      { id: 'hidden', title: '隐藏', description: '不展示', icon: 'i-tabler-eye-off', collectionSlug: 'quickstart', enabled: false }
    ]
  }, [
    { title: '快速开始', slug: 'quickstart', description: 'Start here' },
    { title: '后台指南', slug: 'admin', description: 'Admin operations' }
  ])

  assert.deepEqual(links.map(link => [link.title, link.to]), [
    ['运营手册', '/admin'],
    ['状态页', 'https://status.example.com']
  ])
})

test('homeFeaturedCollections follows backend configured order', () => {
  const ranked = homeFeaturedCollections([
    { title: 'Smoke Test', slug: 'smoke', docCount: 1 },
    { title: 'Quickstart', slug: 'quickstart', docCount: 5 },
    { title: 'Reference', slug: 'reference', docCount: 3 }
  ], { featuredCollections: ['reference', 'quickstart'] })

  assert.deepEqual(ranked.map(collection => collection.slug), ['reference', 'quickstart'])
})

test('homeFeaturedCollections stays empty when backend has no recommendations configured', () => {
  const ranked = homeFeaturedCollections([
    { title: 'Smoke Test', slug: 'smoke', docCount: 1 },
    { title: 'Quickstart', slug: 'quickstart', docCount: 5 }
  ], { featuredCollections: [] })

  assert.deepEqual(ranked, [])
})
