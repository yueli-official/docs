export function shouldShowSearchSkeleton({ pending, hasQuery, resultCount }) {
  return Boolean(pending && hasQuery && Number(resultCount || 0) === 0)
}

export function shouldDimSearchResults({ pending, resultCount }) {
  return Boolean(pending && Number(resultCount || 0) > 0)
}

export function searchStatusText({ hasQuery, pending, total }) {
  if (hasQuery && pending) return '正在更新结果'
  if (hasQuery) return `找到 ${Number(total || 0)} 个结果`
  return '按 Enter 或点击结果打开文档'
}
