import test from 'node:test'
import assert from 'node:assert/strict'
import {
  searchStatusText,
  shouldDimSearchResults,
  shouldShowSearchSkeleton
} from './docsSearchState.mjs'

test('shouldShowSearchSkeleton only replaces the panel when no previous results exist', () => {
  assert.equal(shouldShowSearchSkeleton({ pending: true, hasQuery: true, resultCount: 0 }), true)
  assert.equal(shouldShowSearchSkeleton({ pending: true, hasQuery: true, resultCount: 3 }), false)
  assert.equal(shouldShowSearchSkeleton({ pending: false, hasQuery: true, resultCount: 0 }), false)
})

test('shouldDimSearchResults keeps existing results visible while a new query loads', () => {
  assert.equal(shouldDimSearchResults({ pending: true, resultCount: 2 }), true)
  assert.equal(shouldDimSearchResults({ pending: true, resultCount: 0 }), false)
  assert.equal(shouldDimSearchResults({ pending: false, resultCount: 2 }), false)
})

test('searchStatusText reports loading without hiding the previous total', () => {
  assert.equal(searchStatusText({ hasQuery: true, pending: true, total: 8 }), '正在更新结果')
  assert.equal(searchStatusText({ hasQuery: true, pending: false, total: 8 }), '找到 8 个结果')
  assert.equal(searchStatusText({ hasQuery: false, pending: false, total: 0 }), '按 Enter 或点击结果打开文档')
})
