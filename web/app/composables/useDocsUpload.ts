import { assetUploadURL, createAssetUploadMemo } from "@yueli/asset-nuxt/upload";

type UploadInit = {
  uploadUrl: string;
  uploadToken: string;
  uploadHeaders?: Record<string, string>;
};

type DocumentImageScope = {
  collectionId?: string;
  documentId?: string;
};

export function useDocsUpload() {
  const { call } = useApi();
  const scopes = new Map<string, ReturnType<typeof createAssetUploadMemo<{ result?: Promise<string> }>>>();

  function put(
    url: string,
    file: File,
    headers?: Record<string, string>,
  ): Promise<void> {
    return new Promise((resolve, reject) => {
      const transferURL = assetUploadURL(url);
      const xhr = new XMLHttpRequest();
      xhr.open("PUT", transferURL);
      for (const [key, value] of Object.entries(headers ?? {})) {
        xhr.setRequestHeader(key, value);
      }
      xhr.onload = () =>
        xhr.status >= 200 && xhr.status < 300
          ? resolve()
          : reject(new Error(`上传失败（HTTP ${xhr.status}）`));
      xhr.onerror = () =>
        reject(
          new Error(
            transferURL.startsWith("/asset-api/")
              ? "站点素材代理连接失败，请刷新后重试"
              : "对象存储直传失败，请检查网络或存储 CORS",
          ),
        );
      xhr.send(file);
    });
  }

  async function sendDocumentImage(
    file: File,
    scope: DocumentImageScope,
  ): Promise<string> {
    const init = await call<UploadInit>("/api/v1/images", {
      method: "POST",
      body: {
        ...scope,
        filename: file.name,
        mime: file.type,
        size: file.size,
      },
    });
    await put(init.uploadUrl, file, init.uploadHeaders);
    const result = await call<{ url: string }>("/api/v1/images/finalize", {
      method: "POST",
      body: { ...scope, uploadToken: init.uploadToken },
    });
    return result.url;
  }

  async function uploadDocumentImage(file: File, scope: DocumentImageScope): Promise<string> {
    const key = JSON.stringify([scope.collectionId || '', scope.documentId || '']);
    let memo = scopes.get(key);
    if (!memo) { memo = createAssetUploadMemo(); scopes.set(key, memo); }
    const attempt = await memo.get(file, () => ({}));
    return attempt.result ||= sendDocumentImage(file, scope).catch(error => { delete attempt.result; throw error; });
  }
  return { uploadDocumentImage };
}
