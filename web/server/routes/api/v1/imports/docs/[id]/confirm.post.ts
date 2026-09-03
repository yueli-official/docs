import {
  createError,
  defineEventHandler,
  getRouterParam,
  sendProxy,
} from "h3";

// Confirmation can upload hundreds of deduplicated package assets and may
// legitimately wait for an upstream Retry-After window. Keep that long-running
// operation outside the generic API proxy's 10 second deadline.
export default defineEventHandler(async (event) => {
  const id = getRouterParam(event, "id");
  if (!id || !/^[0-9a-f-]+$/i.test(id)) {
    throw createError({ statusCode: 400, statusMessage: "Invalid import batch ID" });
  }

  const config = oidcConfig(event);
  const target = identityBffTarget(config.downstreamBase);
  let authHeaders = await sessionAuthHeaders(event);
  if (!authHeaders.authorization) {
    authHeaders = await guestSessionAuthHeaders(event, config.clientId);
  }
  const credential = identityBffCredential(authHeaders);
  const headers = new Headers();
  if (credential.kind === "bearer") headers.set("authorization", `Bearer ${credential.token}`);

  const targetURL = new URL(
    `${target.pathPrefix || ""}/imports/docs/${encodeURIComponent(id)}/confirm`,
    target.origin,
  );
  return await sendProxy(event, targetURL.toString(), {
    fetchOptions: {
      headers,
      method: "POST",
      signal: AbortSignal.timeout(30 * 60_000),
    },
    sendStream: true,
  });
});
