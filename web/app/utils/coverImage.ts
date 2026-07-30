interface CollectionCoverSource {
  coverAssetId?: string
  coverUrl?: string
}

const COLLECTION_COVER_TRANSFORM = '@300x200_mode=fill_type=webp_q=72.webp'

export function collectionCoverThumbUrl(source: CollectionCoverSource): string {
  const coverUrl = source.coverUrl || ''
  const assetId = source.coverAssetId || ''
  if (!assetId) return coverUrl

  const path = `/api/v1/assets/${encodeURIComponent(assetId)}/image/${COLLECTION_COVER_TRANSFORM}`
  if (!coverUrl) return path

  try {
    const url = new URL(coverUrl)
    return `${url.origin}${path}`
  } catch {
    return path
  }
}
