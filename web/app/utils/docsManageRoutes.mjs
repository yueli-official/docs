function cleanSegment(value) {
  return String(value || '').trim().replace(/^\/+|\/+$/g, '')
}

export function normalizeDocSlugPath(path) {
  if (Array.isArray(path)) return path.map(cleanSegment).filter(Boolean)
  return cleanSegment(path).split('/').map(cleanSegment).filter(Boolean)
}

export function docManageRoute(collectionSlug, slugPath) {
  const collection = encodeURIComponent(cleanSegment(collectionSlug))
  const path = normalizeDocSlugPath(slugPath).map(encodeURIComponent).join('/')
  return path ? `/manage/docs/${collection}/${path}` : `/manage/docs/${collection}`
}

export function findDocBySlugPath(nodes, slugPath) {
  const parts = normalizeDocSlugPath(slugPath)
  if (!parts.length) return null
  let currentNodes = nodes || []
  let current = null
  for (const part of parts) {
    current = currentNodes.find(node => node?.slug === part) || null
    if (!current) return null
    currentNodes = current.children || []
  }
  return current
}

export function findDocSlugPathById(nodes, id, prefix = []) {
  for (const node of nodes || []) {
    const path = [...prefix, node.slug].filter(Boolean)
    if (node.id === id) return path
    const childPath = findDocSlugPathById(node.children || [], id, path)
    if (childPath) return childPath
  }
  return null
}

export function buildDocTree(items) {
  const byId = new Map()
  const roots = []

  for (const item of items || []) {
    byId.set(item.id, { ...item, children: [] })
  }

  for (const node of byId.values()) {
    const parent = node.parentId ? byId.get(node.parentId) : null
    if (parent) parent.children.push(node)
    else roots.push(node)
  }

  function sortNodes(nodes) {
    nodes.sort((a, b) => (a.sortOrder ?? 0) - (b.sortOrder ?? 0) || String(a.title || '').localeCompare(String(b.title || '')))
    for (const node of nodes) sortNodes(node.children || [])
    return nodes
  }

  return sortNodes(roots)
}
