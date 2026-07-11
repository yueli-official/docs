import test from 'node:test'
import assert from 'node:assert/strict'
import {
  buildDocTree,
  docManageRoute,
  findDocBySlugPath,
  findDocSlugPathById
} from './docsManageRoutes.mjs'

const tree = [
  {
    id: 'root-id',
    slug: 'guide',
    title: 'Guide',
    children: [
      { id: 'child-id', slug: 'install', title: 'Install' }
    ]
  },
  { id: 'api-id', slug: 'api', title: 'API' }
]

test('docManageRoute builds semantic manage edit URLs from collection and slug path', () => {
  assert.equal(docManageRoute('sapphire', ['guide', 'install']), '/manage/docs/sapphire/guide/install')
})

test('findDocBySlugPath resolves nested docs from semantic route segments', () => {
  assert.equal(findDocBySlugPath(tree, ['guide', 'install'])?.id, 'child-id')
  assert.equal(findDocBySlugPath(tree, ['api'])?.id, 'api-id')
  assert.equal(findDocBySlugPath(tree, ['missing']), null)
})

test('findDocSlugPathById builds slug paths for rows and tree actions', () => {
  assert.deepEqual(findDocSlugPathById(tree, 'child-id'), ['guide', 'install'])
  assert.deepEqual(findDocSlugPathById(tree, 'api-id'), ['api'])
  assert.equal(findDocSlugPathById(tree, 'missing'), null)
})

test('buildDocTree turns flat docs into ordered nested nodes', () => {
  const nested = buildDocTree([
    { id: 'b', parentId: 'a', slug: 'child-b', title: 'B', sortOrder: 1 },
    { id: 'a', parentId: '', slug: 'root-a', title: 'A', sortOrder: 0 },
    { id: 'c', parentId: 'a', slug: 'child-c', title: 'C', sortOrder: 0 },
  ])
  assert.deepEqual(nested.map(n => n.id), ['a'])
  assert.deepEqual(nested[0].children.map(n => n.id), ['c', 'b'])
})
