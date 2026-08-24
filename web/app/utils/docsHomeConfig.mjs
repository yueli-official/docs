/**
 * @param {unknown} value
 * @returns {string[]}
 */
export function normalizeFeaturedCollections(value) {
  return Array.isArray(value)
    ? value.filter((item) => typeof item === "string")
    : [];
}
