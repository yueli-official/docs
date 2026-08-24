/** Reduces a referrer to a privacy-safe attribution host. */
export function trafficSource(referrer, currentURL) {
  if (!referrer) return "direct";
  try {
    const previous = new URL(referrer);
    const current = new URL(currentURL);
    if (!["http:", "https:"].includes(previous.protocol)) return "direct";
    const previousHost = normalizeHost(previous.hostname);
    const currentHost = normalizeHost(current.hostname);
    if (!previousHost) return "direct";
    return previousHost === currentHost ? "internal" : previousHost;
  } catch {
    return "direct";
  }
}

function normalizeHost(value) {
  return value
    .trim()
    .toLowerCase()
    .replace(/^www\./, "")
    .replace(/\.$/, "")
    .slice(0, 200);
}
