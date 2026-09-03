import {
  createError,
  defineEventHandler,
  getRequestHeader,
  getRequestWebStream,
  sendProxy,
} from "h3";

const MAX_IMPORT_PACKAGE_BYTES = 100 * 1024 * 1024;

// Documentation packages are intentionally larger and slower to preflight than
// ordinary JSON API calls. This exact route streams the multipart body while
// retaining Identity's shared session and guest-token adapters.
export default defineEventHandler(async (event) => {
  const config = oidcConfig(event);
  const target = identityBffTarget(config.downstreamBase);
  let authHeaders = await sessionAuthHeaders(event);
  if (!authHeaders.authorization) {
    authHeaders = await guestSessionAuthHeaders(event, config.clientId);
  }
  const credential = identityBffCredential(authHeaders);
  const headers = new Headers();
  const contentType = getRequestHeader(event, "content-type");
  const contentLength = getRequestHeader(event, "content-length");
  const declaredBytes = Number(contentLength);
  if (!contentLength || !Number.isSafeInteger(declaredBytes) || declaredBytes < 0) {
    throw createError({ statusCode: 411, statusMessage: "Import package size is required" });
  }
  if (declaredBytes > MAX_IMPORT_PACKAGE_BYTES) {
    throw createError({ statusCode: 413, statusMessage: "Import package exceeds 100 MiB" });
  }
  if (contentType) headers.set("content-type", contentType);
  if (contentLength) headers.set("content-length", contentLength);
  if (credential.kind === "bearer") headers.set("authorization", `Bearer ${credential.token}`);

  const targetURL = new URL(
    `${target.pathPrefix || ""}/imports/docs`,
    target.origin,
  );
  return await sendProxy(event, targetURL.toString(), {
    fetchOptions: {
      body: getRequestWebStream(event),
      duplex: "half",
      headers,
      method: "POST",
      signal: AbortSignal.timeout(5 * 60_000),
    },
    sendStream: true,
  });
});
