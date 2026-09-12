import { execFileSync } from "node:child_process";
import { expect, request, test, type APIRequestContext } from "@playwright/test";
import { capturePageFailures, expectNoHorizontalOverflow, loginE2E, settleNuxt } from "./runtime";

// Credentials stay in memory; a failure trace must not capture PAT plaintext.
test.use({ trace: "off", video: "off" });

test("personal token completes collection, media, document and import publishing", async ({ browser }, testInfo) => {
  test.setTimeout(300_000);
  const site = process.env.DOCS_E2E_URL!;
  const account = process.env.DOCS_E2E_ACCOUNT_URL!;
  const session = await loginE2E(browser, {}, undefined, site);
  const page = await session.newPage();
  const failures = capturePageFailures(page);
  const tokenIDs: number[] = [];
  const collections: string[] = [];
  const clients: APIRequestContext[] = [];
  const suffix = Date.now().toString();

  async function json(client: APIRequestContext, method: string, path: string, status: number, data?: unknown) {
    const response = await client.fetch(path, { method, data });
    expect(response.status(), `${method} ${path}: ${response.status() === status ? "" : await response.text()}`).toBe(status);
    return status === 204 ? undefined : await response.json();
  }
  async function clientFor(token: string) {
    const client = await request.newContext({ baseURL: site, extraHTTPHeaders: { authorization: `Bearer ${token}` } });
    clients.push(client);
    return client;
  }
  async function tokenFor(...capabilities: string[]) {
    const scopes = capabilities.map((key) => `site:${Buffer.from("docs-main-web").toString("base64url")}:${key}`);
    const result = await json(session.request, "POST", `${account}/api/v1/pat`, 201, { name: `Docs 权限回归 ${suffix}`, scopes, expiresInDays: 1 });
    tokenIDs.push(result.id);
    return clientFor(result.token);
  }

  // A small valid PNG is sufficient to exercise real Asset storage and delivery.
  const png = Buffer.from("iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAIAAAD8GO2jAAAAKklEQVR4nGPQaFpAU8QwasGoBaMWjFowasGoBaMWjFowasGoBaMWDBULAID+KExJqnjYAAAAAElFTkSuQmCC", "base64");
  const gif = Buffer.from("R0lGODlhAQABAIAAAAUEBAAAACwAAAAAAQABAAACAkQBADs=", "base64");

  try {
    const catalog = await json(session.request, "GET", `${account}/api/v1/pat/scopes`, 200);
    const scopes = catalog.items.filter((item: { site: string }) => item.site === "docs-main-web");
    expect(scopes).toHaveLength(9);
    expect(catalog.unavailableSites).not.toContain("月离文档");

    await page.goto(`${account}/developer-tokens`);
    await settleNuxt(page);
    await page.getByRole("button", { name: "创建令牌", exact: true }).click();
    const dialog = page.getByRole("dialog");
    await dialog.getByLabel("名称", { exact: true }).fill(`Docs 完整投稿 ${suffix}`);
    for (const scope of scopes) {
      await dialog.getByRole("checkbox", { name: scope.label, exact: true }).check();
    }
    await page.screenshot({ path: testInfo.outputPath("developer-token-desktop.png"), fullPage: true });
    await page.setViewportSize({ width: 390, height: 844 });
    await expectNoHorizontalOverflow(page, "开发者令牌");
    await page.screenshot({ path: testInfo.outputPath("developer-token-mobile.png"), fullPage: true });
    const created = page.waitForResponse((response) => response.url().endsWith("/api/v1/pat") && response.request().method() === "POST");
    await dialog.getByRole("button", { name: "创建令牌", exact: true }).click();
    const tokenResponse = await created;
    expect(tokenResponse.status()).toBe(201);
    const token = await tokenResponse.json();
    tokenIDs.push(token.id);
    const pat = await clientFor(token.token);

    const reader = await tokenFor("docs.document.read");
    await json(reader, "POST", "/api/v1/collections", 403, { title: "不应创建", defaultLocale: "zh-CN", semanticVersion: "1.0.0" });
    const collection = await json(pat, "POST", "/api/v1/collections", 201, {
      title: "开发者令牌投稿验收", slug: `pat-publishing-${suffix}`, description: "由本地自动化创建", defaultLocale: "zh-CN", semanticVersion: "1.0.0",
    });
    collections.push(collection.id);
    const collectionPath = `/api/v1/collections/${collection.id}`;
    const managePath = `/api/v1/manage/collections/${collection.id}`;
    await json(pat, "PATCH", collectionPath, 200, { title: "开发者令牌投稿验收", description: "已通过令牌更新", slug: collection.slug });
    const versions = await json(reader, "GET", `${collectionPath}/versions`, 200);
    expect(versions.items).toHaveLength(1);
    await json(pat, "POST", `${managePath}/locales`, 200, { locale: "en-US", label: "English", enabled: true });
    await json(pat, "POST", `${managePath}/locales`, 200, { locale: "fr", label: "Français", enabled: true });
    await json(pat, "DELETE", `${managePath}/locales/fr`, 204);
    expect((await json(pat, "GET", `${managePath}/locales`, 200)).items.map((item: { locale: string }) => item.locale)).toContain("en-US");

    const version = (await json(pat, "POST", `${collectionPath}/versions`, 201, { key: "next", label: "预览版", status: "draft" })).version;
    await json(pat, "PATCH", `${managePath}/versions/${version.id}`, 200, { label: "预览版", status: "published", isDefault: false });
    await json(reader, "POST", `${collectionPath}/versions`, 403, { key: "v2" });

    async function upload(initPath: string, finalizePath: string, target: object, bytes = png, mime = "image/png") {
      const upload = await json(pat, "POST", initPath, 201, { ...target, filename: `pat-${suffix}.${mime.split("/")[1]}`, mime, size: bytes.length });
      const response = await session.request.put(new URL(upload.uploadUrl, site).toString(), {
        headers: { "content-type": mime, ...upload.uploadHeaders }, data: bytes,
      });
      expect(response.ok(), `Asset PUT: ${response.status()}`).toBeTruthy();
      return json(pat, "POST", finalizePath, 200, { ...target, uploadToken: upload.uploadToken });
    }

    const cover = await upload(`${collectionPath}/cover`, `${collectionPath}/cover/finalize`, {});
    expect(cover.coverUrl).toContain("/media/");
    expect((await session.request.get(new URL(cover.coverUrl, site).toString())).ok()).toBeTruthy();
    const parent = (await json(pat, "POST", "/api/v1/docs", 201, { collectionId: collection.id, title: "本地投稿", slug: "guide", locale: "zh-CN" })).doc;
    const doc = (await json(pat, "POST", "/api/v1/docs", 201, { collectionId: collection.id, parentId: parent.id, title: "从令牌发布", slug: "start", content: "<p>草稿</p>", locale: "zh-CN" })).doc;
    await json(reader, "POST", "/api/v1/images", 403, { documentId: doc.id, filename: "image.png" });
    const writer = await tokenFor("docs.document.update");
    await json(writer, "PATCH", `/api/v1/docs/${doc.id}`, 403, { status: "published" });
    const publisher = await tokenFor("docs.document.publish");
    await json(publisher, "PATCH", `/api/v1/docs/${doc.id}`, 403, { content: "不能编辑" });
    const image = await upload("/api/v1/images", "/api/v1/images/finalize", { documentId: doc.id }, gif, "image/gif");
    await json(pat, "PATCH", `/api/v1/docs/${doc.id}`, 200, { content: `<h2>脚本完成的投稿</h2><p>文档集、语言、版本、图片和正文均由开发者令牌提交。</p><img src="${image.url}" alt="令牌上传图片">`, sortOrder: 2 });
    await json(pat, "POST", `/api/v1/docs/${parent.id}/publish`, 200);
    await json(pat, "POST", `/api/v1/docs/${doc.id}/publish`, 200);
    expect((await json(reader, "GET", `/api/v1/docs/${doc.id}`, 200)).doc.status).toBe("published");
    expect((await json(pat, "GET", `/api/v1/manage/collections/${collection.slug}/tree?locale=zh-CN`, 200)).tree.length).toBeGreaterThan(0);

    await page.setViewportSize({ width: 1440, height: 1000 });
    await page.goto(`${site}/${collection.slug}/guide/start`);
    await settleNuxt(page);
    await expect(page.getByRole("heading", { name: "脚本完成的投稿" })).toBeVisible();
    const renderedImage = page.getByRole("img", { name: "令牌上传图片" });
    await expect(renderedImage).toBeVisible();
    await expect.poll(() => renderedImage.evaluate((img: HTMLImageElement) => img.naturalWidth)).toBeGreaterThan(0);
    await page.screenshot({ path: testInfo.outputPath("published-document.png"), fullPage: true });
    await page.setViewportSize({ width: 390, height: 844 });
    await expectNoHorizontalOverflow(page, "投稿文档");
    await page.screenshot({ path: testInfo.outputPath("published-document-mobile.png"), fullPage: true });

    const translated = await json(pat, "POST", `${managePath}/locales/clone`, 200, { sourceLocale: "zh-CN", targetLocale: "en-US", targetLabel: "English" });
    expect(translated.created).toBe(2);
    const release = (await json(pat, "POST", `${managePath}/clone-release`, 201, { targetSemanticVersion: "2.0.0", title: "令牌发布第二版", slug: `pat-release-${suffix}` })).collection;
    collections.push(release.id);

    // Exercise the existing package path as well as individual document APIs.
    for (const revision of [1, 2]) {
      const zip = execFileSync("python", ["-c", "import base64,io,json,sys,zipfile; b=io.BytesIO(); z=zipfile.ZipFile(b,'w'); z.writestr('docs.json',json.dumps({'schemaVersion':1,'defaultLocale':'zh-CN','locales':{'zh-CN':'.'}})); z.writestr('imported.md','---\\nid: pat-import\\ntitle: 令牌导入\\n---\\n# 令牌导入\\n第'+sys.argv[1]+'版\\n![图片](image.png)\\n'); z.writestr('image.png',base64.b64decode(sys.argv[2])); z.close(); sys.stdout.buffer.write(b.getvalue())", String(revision), png.toString("base64")]);
      const upload = await pat.post("/api/v1/imports/docs", { multipart: { file: { name: "docs.zip", mimeType: "application/zip", buffer: zip }, collection: collection.slug, defaultLocale: "zh-CN", mode: "upsert" } });
      expect(upload.status(), await upload.text()).toBe(201);
      const batch = (await upload.json()).batch;
      await json(pat, "POST", `/api/v1/imports/docs/${batch.id}/confirm`, 202);
      await expect.poll(async () => (await json(pat, "GET", `/api/v1/imports/docs/${batch.id}`, 200)).batch.status, { timeout: 60_000, intervals: [1_000] }).toBe("completed");
    }

    for (const path of ["/api/v1/import-sources", "/api/v1/authorization/manage/roles"]) {
      await json(pat, "GET", path, 403);
    }
    await json(reader, "DELETE", `/api/v1/docs/${doc.id}`, 403);
    await json(pat, "POST", `/api/v1/docs/${doc.id}/archive`, 200);
    expect((await json(reader, "GET", `/api/v1/docs/${doc.id}`, 200)).doc.status).toBe("archived");
    await json(pat, "DELETE", `/api/v1/docs/${doc.id}`, 204);

    await json(session.request, "DELETE", `${account}/api/v1/pat/${token.id}`, 204);
    tokenIDs.splice(tokenIDs.indexOf(token.id), 1);
    const revoked = await pat.post("/api/v1/docs", { data: { collectionId: collection.id, title: "不应创建" } });
    expect([401, 403]).toContain(revoked.status());
    const noFallback = await session.request.post(`${site}/api/v1/docs`, { headers: { authorization: `Bearer ${token.token}` }, data: { collectionId: collection.id, title: "不得回退浏览器会话" } });
    expect([401, 403]).toContain(noFallback.status());
    expect(failures).toEqual([]);
  } finally {
    for (const id of collections.reverse()) {
      await json(session.request, "DELETE", `${site}/api/v1/collections/${id}`, 204);
    }
    for (const id of tokenIDs) {
      await json(session.request, "DELETE", `${account}/api/v1/pat/${id}`, 204);
    }
    for (const client of clients) await client.dispose();
    await session.close();
  }
});
