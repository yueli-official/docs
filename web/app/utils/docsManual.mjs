function normalize(value) {
  return String(value || '').trim().toLocaleLowerCase()
}

export function filterCollections(collections, query) {
  const q = normalize(query)
  if (!q) return collections
  return collections.filter((collection) => {
    return normalize(collection.title).includes(q)
      || normalize(collection.description).includes(q)
      || normalize(collection.slug).includes(q)
  })
}

export function paginateItems(items, page = 1, pageSize = 20) {
  const safeSize = Math.max(1, Number(pageSize) || 20)
  const total = items.length
  const totalPages = Math.max(1, Math.ceil(total / safeSize))
  const safePage = Math.min(Math.max(1, Number(page) || 1), totalPages)
  const start = total ? (safePage - 1) * safeSize + 1 : 0
  const end = Math.min(safePage * safeSize, total)
  return {
    items: items.slice((safePage - 1) * safeSize, safePage * safeSize),
    page: safePage,
    pageSize: safeSize,
    total,
    totalPages,
    start,
    end
  }
}

export function collectionCountLabel(count) {
  return count > 0 ? `${count} 篇文档` : '文档集'
}

export function rootSectionCountLabel(count) {
  return count > 0 ? `${count} 个章节` : '暂无公开章节'
}

export function collectionStats(collections) {
  return {
    collectionCount: collections.length,
    docCount: collections.reduce((sum, collection) => sum + Number(collection.docCount || 0), 0)
  }
}

export function homeQuickLinks(config = {}, collections = []) {
  const bySlug = new Map(collections.map(collection => [collection.slug, collection]))
  return (config.quickLinks || [])
    .filter(link => link && link.enabled !== false)
    .slice()
    .sort((a, b) => Number(a.sortOrder || 0) - Number(b.sortOrder || 0))
    .map((link) => {
      const collection = bySlug.get(link.collectionSlug)
      const to = link.to || (collection ? `/${collection.slug}` : '/')
      const title = link.title || collection?.title || '快速入口'
      const description = link.description || collection?.description || ''
      const icon = link.icon || collection?.icon || 'i-tabler-arrow-up-right'
      return { ...link, title, description, icon, to }
    })
}

export function homeFeaturedCollections(collections, config = {}, limit = 6) {
  const bySlug = new Map(collections.map(collection => [collection.slug, collection]))
  return (config.featuredCollections || [])
    .map(slug => bySlug.get(slug))
    .filter(Boolean)
    .slice(0, limit)
}
