import type { MaybeRefOrGetter } from 'vue'
import type { CollectionTree, DocTreeNode } from '~/types'

// Module-level in-flight dedup: concurrent ensure() calls for the same
// (slug, locale, version) key share ONE fetch and all await it, instead of the loser
// early-returning before the fetch resolves (which left byPath empty → spurious
// 404 on cross-collection nav, e.g. layout watch + page await firing together).
// Safe to share across requests: tree data is public/read-only and each caller
// still writes its own per-request treeState().
const inFlight = new Map<string, Promise<CollectionTree>>()

// Fetch a collection's full tree ONCE per (slug, locale, version) and share it across the
// layout + all doc pages under it via useState, so client-side nav between docs
// in the same collection does not refetch. The slug is reactive: when the active
// collection changes, the composable re-keys to that collection's own state.
export function useCollectionTree(
  slugInput: MaybeRefOrGetter<string>,
  localeInput: MaybeRefOrGetter<string> = 'en',
  versionInput: MaybeRefOrGetter<string> = '',
) {
  const { call } = useApi()
  const slug = toRef(slugInput)
  const locale = computed(() => toValue(localeInput) || 'en')
  const version = computed(() => toValue(versionInput) || '')
  const key = computed(() => `tree-${slug.value}-${locale.value}-${version.value || 'default'}`)
  const cache = useState<Record<string, CollectionTree | null>>('docs-collection-tree-cache', () => ({}))
  const pendingByKey = useState<Record<string, boolean>>('docs-collection-tree-pending', () => ({}))

  async function ensure() {
    const k = key.value
    if (cache.value[k]) return
    let p = inFlight.get(k)
    if (!p) {
      pendingByKey.value = { ...pendingByKey.value, [k]: true }
      p = call<CollectionTree>(`/api/v1/collections/${slug.value}/tree`, {
        query: {
          locale: locale.value,
          ...(version.value ? { version: version.value } : {}),
        },
      })
        .finally(() => {
          pendingByKey.value = { ...pendingByKey.value, [k]: false }
          inFlight.delete(k)
        })
      inFlight.set(k, p)
    }
    cache.value = { ...cache.value, [k]: await p } // all concurrent callers await the SAME in-flight fetch
  }

  const current = computed(() => cache.value[key.value] ?? null)
  const tree = computed(() => current.value?.tree ?? [])
  const collection = computed(() => current.value?.collection ?? null)
  const pending = computed(() => pendingByKey.value[key.value] ?? false)

  // DFS pre-order: each node carries its full URL path (ancestor slugs joined).
  const flat = computed(() => {
    const out: { node: DocTreeNode; path: string }[] = []
    const walk = (nodes: DocTreeNode[], base: string) => {
      for (const n of nodes) {
        const p = base ? `${base}/${n.slug}` : n.slug
        out.push({ node: n, path: p })
        if (n.children?.length) walk(n.children, p)
      }
    }
    walk(tree.value, '')
    return out
  })
  const byPath = computed(() => new Map(flat.value.map(e => [e.path, e.node])))

  return { ensure, tree, collection, flat, byPath, pending }
}
