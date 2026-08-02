import { publicAssetMediaUrl } from './asset-media.mjs'

interface CollectionCoverSource {
  coverAssetId?: string
  coverUrl?: string
}

export function collectionCoverThumbUrl(source: CollectionCoverSource): string {
  const coverUrl = source.coverUrl || ''
  const assetId = source.coverAssetId || ''
  if (!assetId) return coverUrl

  if (!coverUrl) return publicAssetMediaUrl(assetId, 'cover', '')

  try {
    const url = new URL(coverUrl)
    return publicAssetMediaUrl(assetId, 'cover', url.origin)
  } catch {
    return publicAssetMediaUrl(assetId, 'cover', '')
  }
}
