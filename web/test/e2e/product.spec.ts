import { expect, test } from "@playwright/test";
import { productSites } from "./contracts";
import { registerProductSuite } from "./product-suite";
import {
  capturePageFailures,
  expectNoHorizontalOverflow,
  loginE2E,
  settleNuxt,
} from "./runtime";

registerProductSuite("docs");

test("管理控制台在桌面宽度保持稳定双栏", async ({ browser }) => {
  const site = productSites("docs")[0]!;
  const context = await loginE2E(
    browser,
    { viewport: { width: 1060, height: 1000 } },
    "light",
    site.url,
  );
  const page = await context.newPage();
  try {
    await page.goto(new URL("/manage", site.url).toString(), {
      waitUntil: "domcontentloaded",
    });
    await settleNuxt(page);

    const sectionFor = (name: string) =>
      page.locator("section").filter({
        has: page.getByRole("heading", { name, exact: true }),
      });
    const pending = await sectionFor("待完善文档").boundingBox();
    const health = await sectionFor("工作区状态").boundingBox();
    const recent = await sectionFor("最近更新").boundingBox();
    const actions = await sectionFor("快捷操作").boundingBox();
    expect(pending).not.toBeNull();
    expect(health).not.toBeNull();
    expect(recent).not.toBeNull();
    expect(actions).not.toBeNull();
    expect(Math.abs(pending!.y - health!.y)).toBeLessThan(4);
    expect(health!.x).toBeGreaterThan(pending!.x + pending!.width);
    expect(Math.abs(recent!.x - pending!.x)).toBeLessThan(4);
    expect(Math.abs(recent!.width - pending!.width)).toBeLessThan(4);
    expect(Math.abs(actions!.x - health!.x)).toBeLessThan(4);
    expect(Math.abs(actions!.width - health!.width)).toBeLessThan(4);
    expect(recent!.y).toBeGreaterThan(pending!.y + pending!.height);
    expect(actions!.y).toBeGreaterThan(health!.y + health!.height);
  } finally {
    await context.close();
  }
});

test("文档集编辑器使用紧凑封面与共享图标管理器", async ({ browser }, testInfo) => {
  const site = productSites("docs")[0]!;
  const context = await loginE2E(
    browser,
    { viewport: { width: 1440, height: 960 } },
    "light",
    site.url,
  );
  const page = await context.newPage();
  const failures = capturePageFailures(page);
  try {
    await page.goto(new URL("/manage/collections", site.url).toString(), {
      waitUntil: "domcontentloaded",
    });
    await settleNuxt(page);
    const publicCollection = page.getByRole("link", { name: /^查看文档集：/u }).first();
    const editCollectionAction = page.getByRole("button", { name: /^编辑文档集：/u }).first();
    await expect(publicCollection).toBeVisible();
    await expect(publicCollection).toHaveAttribute("href", /^\/[a-z0-9-]+$/u);
    const [publicActionBox, publicIconBox, editActionBox, editIconBox] = await Promise.all([
      publicCollection.boundingBox(),
      publicCollection.locator(".iconify").boundingBox(),
      editCollectionAction.boundingBox(),
      editCollectionAction.locator(".iconify").boundingBox(),
    ]);
    for (const box of [publicActionBox, editActionBox]) {
      expect(box).not.toBeNull();
      expect(box!.width).toBeCloseTo(24, 0);
      expect(box!.height).toBeCloseTo(24, 0);
    }
    for (const box of [publicIconBox, editIconBox]) {
      expect(box).not.toBeNull();
      expect(box!.width).toBeCloseTo(16, 0);
      expect(box!.height).toBeCloseTo(16, 0);
    }
    for (const [buttonBox, iconBox] of [[publicActionBox, publicIconBox], [editActionBox, editIconBox]]) {
      expect(Math.abs(
        buttonBox!.x + buttonBox!.width / 2 -
        (iconBox!.x + iconBox!.width / 2),
      )).toBeLessThan(1);
      expect(Math.abs(
        buttonBox!.y + buttonBox!.height / 2 -
        (iconBox!.y + iconBox!.height / 2),
      )).toBeLessThan(1);
    }
    await publicCollection.hover();
    await expect(page.getByText(await publicCollection.getAttribute("aria-label") || "", { exact: true }).last()).toBeVisible();
    await expect(
      page.getByRole("button", { name: /^管理语言与版本：/u }),
    ).toHaveCount(0);
    await editCollectionAction.click();

    const editor = page.getByRole("dialog", { name: "编辑文档集" });
    await expect(editor).toBeVisible();
    const basicTab = editor.getByRole("tab", { name: "基础" });
    const languageTab = editor.getByRole("tab", { name: "语言" });
    const versionTab = editor.getByRole("tab", { name: "版本" });
    await expect(basicTab).toBeVisible();
    await languageTab.click();
    await expect(editor.getByText("启用语言", { exact: true })).toBeVisible();
    await versionTab.click();
    await expect(editor.getByText("关联版本", { exact: true })).toBeVisible();
    await expect(
      editor.getByRole("button", { name: "设置当前版本", exact: true })
        .or(editor.getByText(/^\d+\.\d+\.\d+$/u).first()),
    ).toBeVisible();
    await expect(editor.getByRole("button", { name: "克隆为新版本", exact: true })).toBeVisible();
    await basicTab.click();
    await expect(editor.getByText("路径预览", { exact: true })).toHaveCount(0);
    await expect(editor.getByText("危险操作", { exact: true })).toHaveCount(0);
    const iconSearch = editor.getByPlaceholder("搜索全部 Tabler 图标");
    await expect(iconSearch).toBeVisible();
    await editor.getByRole("button", { name: "查看封面规格", exact: true }).hover();
    await expect(
      page.getByText("1:1 · 256 × 256 · WebP", { exact: true }).last(),
    ).toBeVisible();

    const coverPreview = editor.locator("[data-cover-preview]");
    const iconColumn = editor.locator("[data-icon-column]");
    await expect(editor.getByText("封面链接", { exact: true })).toHaveCount(0);
    await expect(editor.getByRole("button", { name: "更换封面" })).toHaveCount(0);
    const removeCover = editor.getByRole("button", { name: "移除封面" });
    await expect(removeCover).toBeVisible();
    const [initialPreviewBox, removeBox] = await Promise.all([
      coverPreview.boundingBox(),
      removeCover.boundingBox(),
    ]);
    expect(initialPreviewBox).not.toBeNull();
    expect(removeBox).not.toBeNull();
    expect(removeBox!.x).toBeGreaterThan(
      initialPreviewBox!.x + initialPreviewBox!.width - 42,
    );
    expect(removeBox!.y).toBeLessThan(initialPreviewBox!.y + 42);
    await removeCover.click();
    await expect(editor.getByRole("button", { name: "上传封面" })).toBeVisible();
    const [previewBox, iconColumnBox] = await Promise.all([
      coverPreview.boundingBox(),
      iconColumn.boundingBox(),
    ]);
    expect(previewBox).not.toBeNull();
    expect(iconColumnBox).not.toBeNull();
    expect(iconColumnBox!.x).toBeGreaterThan(previewBox!.x + previewBox!.width);

    await iconSearch.fill("terminal");
    await expect(editor.getByRole("button", { name: "选择终端" })).toBeVisible();
    await expect(editor.getByRole("button", { name: "选择应用" })).toHaveCount(0);
    await iconSearch.fill("fish");
    await expect(
      editor.getByRole("button", { name: "选择fish", exact: true }),
    ).toBeVisible();
    await iconSearch.clear();
    await page.screenshot({
      path: testInfo.outputPath("collection-visual-assets.png"),
      fullPage: false,
    });

    const dataIconButton = editor.getByRole("button", { name: "选择数据", exact: true });
    const [iconButtonBox, iconGlyphBox] = await Promise.all([
      dataIconButton.boundingBox(),
      dataIconButton.locator(".iconify").boundingBox(),
    ]);
    expect(iconButtonBox).not.toBeNull();
    expect(iconGlyphBox).not.toBeNull();
    expect(
      Math.abs(
        iconButtonBox!.x + iconButtonBox!.width / 2 -
          (iconGlyphBox!.x + iconGlyphBox!.width / 2),
      ),
    ).toBeLessThan(1);
    expect(
      Math.abs(
        iconButtonBox!.y + iconButtonBox!.height / 2 -
          (iconGlyphBox!.y + iconGlyphBox!.height / 2),
      ),
    ).toBeLessThan(1);
    await dataIconButton.click();
    await expect(editor.getByText("i-tabler-database", { exact: true })).toBeVisible();

    const deleteButton = editor.getByRole("button", { name: "删除文档集", exact: true });
    const cancelButton = editor.getByRole("button", { name: "取消", exact: true });
    const saveButton = editor.getByRole("button", { name: "保存", exact: true });
    await page.waitForTimeout(250);
    const [deleteBox, cancelBox, saveBox] = await Promise.all([
      deleteButton.boundingBox(),
      cancelButton.boundingBox(),
      saveButton.boundingBox(),
    ]);
    expect(deleteBox).not.toBeNull();
    expect(cancelBox).not.toBeNull();
    expect(saveBox).not.toBeNull();
    expect(Math.abs(deleteBox!.y - saveBox!.y)).toBeLessThan(4);
    expect(deleteBox!.x).toBeLessThan(cancelBox!.x);
    expect(cancelBox!.x).toBeLessThan(saveBox!.x);

    await editor.locator('input[type="file"]').setInputFiles({
      name: "collection-cover.svg",
      mimeType: "image/svg+xml",
      buffer: Buffer.from(
        '<svg xmlns="http://www.w3.org/2000/svg" width="512" height="512"><rect width="512" height="512" fill="#2563eb"/></svg>',
      ),
    });
    const cropper = page.getByRole("dialog", { name: "裁剪文档集封面" });
    await expect(cropper).toBeVisible();
    await expect(cropper.getByText("裁剪比例", { exact: true })).toBeVisible();
    await expect(cropper.getByText("1:1", { exact: true })).toBeVisible();
    await expect(cropper.getByText("256 × 256 px", { exact: true })).toBeVisible();
    await expect(cropper.getByText("WebP", { exact: true })).toBeVisible();
    await expectNoHorizontalOverflow(page, "文档集封面裁剪器");
    await page.setViewportSize({ width: 390, height: 844 });
    await expectNoHorizontalOverflow(page, "移动端文档集封面裁剪器");

    await cropper.getByRole("button", { name: "取消", exact: true }).click();
    await page.setViewportSize({ width: 1440, height: 960 });
    await page.goto(new URL("/manage/assets", site.url).toString(), {
      waitUntil: "domcontentloaded",
    });
    await settleNuxt(page);
    await expect(
      page.locator("[data-asset-variant-summary]").filter({
        hasText: "256 × 256",
      }),
    ).toBeVisible();
    expect(
      failures.filter(
        (failure) =>
          !failure.includes("/_nuxt/builds/meta/dev.json") &&
          failure !==
            "console: Failed to load resource: the server responded with a status of 404 (Not Found)",
      ),
    ).toEqual([]);
  } finally {
    await context.close();
  }
});

test("文档集语言与版本保持同一逻辑页面上下文", async ({ browser }) => {
  const site = productSites("docs")[0]!;
  const context = await loginE2E(browser, {}, undefined, site.url);
  const page = await context.newPage();
  await page.setViewportSize({ width: 1440, height: 900 });
  const slug = `e2e-variants-${Date.now()}`;
  const nextSlug = `${slug}-2-0-0`;
  let collectionId = "";
  let nextCollectionId = "";
  try {
    const created = await context.request.post(new URL("/api/v1/collections", site.url).toString(), {
      data: { title: "E2E 多语言版本", slug, description: "variant acceptance" },
    });
    expect(created.ok()).toBeTruthy();
    collectionId = (await created.json() as { id: string }).id;

    const initialVersions = await context.request.get(new URL(`/api/v1/collections/${collectionId}/versions`, site.url).toString());
    const defaultVersion = (await initialVersions.json() as { items: Array<{ id: string; key: string }> }).items[0]!;

    async function createPublishedDoc(data: Record<string, unknown>) {
      const response = await context.request.post(new URL("/api/v1/docs", site.url).toString(), { data });
      expect(response.ok()).toBeTruthy();
      const doc = (await response.json() as { doc: { id: string } }).doc;
      const published = await context.request.post(new URL(`/api/v1/docs/${doc.id}/publish`, site.url).toString());
      expect(published.ok()).toBeTruthy();
      return doc;
    }

    await createPublishedDoc({ collectionId, versionId: defaultVersion.id, locale: "en", slug: "overview", title: "Overview", content: "# Overview", translationKey: "overview" });
    const copiedLocale = await context.request.post(new URL(`/api/v1/manage/collections/${collectionId}/locales/clone`, site.url).toString(), {
      data: {
        sourceLocale: "en", targetLocale: "zh-CN", targetLabel: "简体中文",
        targetHtmlLang: "zh-CN", targetDirection: "ltr", targetSortOrder: 10,
      },
    });
    expect(copiedLocale.ok()).toBeTruthy();
    expect(await copiedLocale.json()).toMatchObject({ created: 1 });
    const chineseDrafts = await context.request.get(new URL(`/api/v1/docs?collectionId=${collectionId}&locale=zh-CN`, site.url).toString());
    const chineseDraft = (await chineseDrafts.json() as { items: Array<{ id: string; status: string; translationKey: string }> }).items[0]!;
    expect(chineseDraft).toMatchObject({ status: "draft", translationKey: "overview" });
    expect((await context.request.patch(new URL(`/api/v1/docs/${chineseDraft.id}`, site.url).toString(), {
      data: { slug: "zonglan", title: "总览", content: "# 总览" },
    })).ok()).toBeTruthy();
    expect((await context.request.post(new URL(`/api/v1/docs/${chineseDraft.id}/publish`, site.url).toString())).ok()).toBeTruthy();

    const initializedRelease = await context.request.post(new URL(`/api/v1/manage/collections/${collectionId}/release`, site.url).toString(), {
      data: { semanticVersion: "1.0.0" },
    });
    expect(initializedRelease.ok()).toBeTruthy();
    expect(await initializedRelease.json()).toMatchObject({ collection: { semanticVersion: "1.0.0" } });
    const cloned = await context.request.post(new URL(`/api/v1/manage/collections/${collectionId}/clone-release`, site.url).toString(), {
      data: {
        targetSemanticVersion: "2.0.0",
        title: "E2E 多语言版本 2.0.0",
        slug: nextSlug,
      },
    });
    expect(cloned.ok()).toBeTruthy();
    nextCollectionId = (await cloned.json() as { collection: { id: string } }).collection.id;

    const clonedEnglish = await context.request.get(new URL(`/api/v1/docs?collectionId=${nextCollectionId}&locale=en`, site.url).toString());
    const clonedEnglishDoc = (await clonedEnglish.json() as { items: Array<{ id: string }> }).items[0]!;
    const updateCloned = await context.request.patch(new URL(`/api/v1/docs/${clonedEnglishDoc.id}`, site.url).toString(), {
      data: { slug: "overview-v2", title: "Overview 2.0", content: "# Overview 2.0" },
    });
    expect(updateCloned.ok()).toBeTruthy();
    expect((await context.request.post(new URL(`/api/v1/docs/${clonedEnglishDoc.id}/publish`, site.url).toString())).ok()).toBeTruthy();

    const releases = await context.request.get(new URL(`/api/v1/collections/${slug}/releases`, site.url).toString());
    expect(releases.ok()).toBeTruthy();
    expect(await releases.json()).toMatchObject({
      items: [
        expect.objectContaining({ slug: nextSlug, semanticVersion: "2.0.0" }),
        expect.objectContaining({ slug, semanticVersion: "1.0.0" }),
      ],
    });

    await page.goto(new URL(`/${nextSlug}`, site.url).toString(), { waitUntil: "networkidle" });
    const collectionVersionSelect = page.getByRole("combobox", { name: "切换文档版本" });
    await expect(collectionVersionSelect).toBeVisible();
    await collectionVersionSelect.press("ArrowDown");
    await page.getByRole("option", { name: "1.0.0", exact: true }).click();
    await expect(page).toHaveURL(new RegExp(`/${slug}$`, "u"));
    await expect(
      page.getByText("已打开文档集首页", { exact: true }),
    ).toHaveCount(0, { timeout: 500 });

    const translated = await context.request.get(new URL(`/api/v1/collections/${slug}/variant?translationKey=overview&locale=zh-CN`, site.url).toString());
    expect(await translated.json()).toMatchObject({ path: "zonglan", locale: "zh-CN", version: "default" });
    const fallback = await context.request.get(new URL(`/api/v1/collections/${nextSlug}/variant?translationKey=overview&locale=zh-CN`, site.url).toString());
    expect(await fallback.json()).toMatchObject({ path: "overview-v2", locale: "en", fallback: "default_locale" });

    await page.goto(new URL(`/${slug}/overview`, site.url).toString(), { waitUntil: "networkidle" });
    const languageSelect = page.getByRole("combobox", { name: "切换文档语言" });
    await expect(languageSelect).toBeVisible();
    await expect(page.getByRole("button", { name: "分享文档" })).toHaveCount(0);
    const [headerBox, switcherBox] = await Promise.all([
      page.locator("article > header").first().boundingBox(),
      page.locator("[data-document-variant-switcher]").boundingBox(),
    ]);
    expect(headerBox).not.toBeNull();
    expect(switcherBox).not.toBeNull();
    expect(switcherBox!.x).toBeGreaterThan(headerBox!.x + headerBox!.width / 2);
    await languageSelect.press("ArrowDown");
    await page.getByRole("option", { name: "简体中文", exact: true }).click();
    await expect(page).toHaveURL(new RegExp(`/${slug}/zonglan\\?locale=zh-CN$`, "u"));
    await expect(
      page.locator("article > header").getByRole("heading", { name: "总览", level: 1 }),
    ).toBeVisible();
    await expect(page.getByText("正在显示默认语言", { exact: true })).toHaveCount(0);

    const versionSelect = page.getByRole("combobox", { name: "切换文档版本" });
    await expect(versionSelect).toBeVisible();
    await versionSelect.press("ArrowDown");
    await page.getByRole("option", { name: "2.0.0", exact: true }).click();
    await expect(page).toHaveURL(new RegExp(`/${nextSlug}/overview-v2$`, "u"));
    const fallbackToast = page.getByText("正在显示默认语言", { exact: true });
    await expect(fallbackToast).toBeVisible();
    await expect(page.getByText("此页暂无简体中文版本。", { exact: true })).toBeVisible();
    await expect(fallbackToast).toHaveCount(0, { timeout: 5_000 });
  } finally {
    if (nextCollectionId) {
      await context.request.delete(new URL(`/api/v1/collections/${nextCollectionId}`, site.url).toString());
    }
    if (collectionId) {
      await context.request.delete(new URL(`/api/v1/collections/${collectionId}`, site.url).toString());
    }
    await context.close();
  }
});

test("文档列表支持把所选层级文档批量改为共同语言", async ({ browser }) => {
  const site = productSites("docs")[0]!;
  const context = await loginE2E(browser, {}, undefined, site.url);
  const page = await context.newPage();
  await page.setViewportSize({ width: 1440, height: 900 });
  const slug = `e2e-bulk-locale-${Date.now()}`;
  let collectionId = "";
  try {
    const created = await context.request.post(
      new URL("/api/v1/collections", site.url).toString(),
      { data: { title: "E2E 批量语言", slug, description: "bulk locale" } },
    );
    expect(created.ok()).toBeTruthy();
    collectionId = (await created.json() as { id: string }).id;
    const localeCreated = await context.request.post(
      new URL(`/api/v1/manage/collections/${collectionId}/locales`, site.url).toString(),
      {
        data: {
          locale: "zh-CN",
          label: "简体中文",
          htmlLang: "zh-CN",
          direction: "ltr",
          isDefault: false,
          enabled: true,
          sortOrder: 10,
        },
      },
    );
    expect(localeCreated.ok()).toBeTruthy();
    const rootCreated = await context.request.post(
      new URL("/api/v1/docs", site.url).toString(),
      { data: { collectionId, title: "批量语言父文档", slug: "parent", locale: "en" } },
    );
    expect(rootCreated.ok()).toBeTruthy();
    const rootDoc = (await rootCreated.json() as { doc: { id: string } }).doc;
    const childCreated = await context.request.post(
      new URL("/api/v1/docs", site.url).toString(),
      {
        data: {
          collectionId,
          parentId: rootDoc.id,
          title: "批量语言子文档",
          slug: "child",
          locale: "en",
        },
      },
    );
    expect(childCreated.ok()).toBeTruthy();
    const childDoc = (await childCreated.json() as { doc: { id: string } }).doc;

    await page.goto(
      new URL(`/manage/docs?collection=${collectionId}`, site.url).toString(),
      { waitUntil: "domcontentloaded" },
    );
    await settleNuxt(page);
    await page
      .getByRole("checkbox", { name: "选择文档：批量语言父文档" })
      .click();
    await page
      .getByRole("checkbox", { name: "选择文档：批量语言子文档" })
      .click();
    const bulkActions = page.locator("[data-docs-bulk-actions]");
    await expect(bulkActions).toContainText("已选择 2 篇");
    await bulkActions.getByRole("combobox").click();
    await page.getByRole("option", { name: "修改语言", exact: true }).click();
    await bulkActions.getByRole("button", { name: "应用", exact: true }).click();
    const dialog = page.getByRole("dialog", { name: "批量修改语言" });
    await expect(dialog).toBeVisible();
    await page.setViewportSize({ width: 390, height: 844 });
    await expectNoHorizontalOverflow(page, "Docs 批量修改语言");
    await page.setViewportSize({ width: 1440, height: 900 });
    await dialog.getByRole("combobox", { name: "目标语言" }).click();
    await page.getByRole("option", { name: "简体中文", exact: true }).click();
    await dialog.getByRole("button", { name: "修改语言", exact: true }).click();
    await expect(dialog).toHaveCount(0);

    for (const id of [rootDoc.id, childDoc.id]) {
      const response = await context.request.get(
        new URL(`/api/v1/docs/${id}`, site.url).toString(),
      );
      expect(response.ok()).toBeTruthy();
      expect(await response.json()).toMatchObject({ doc: { id, locale: "zh-CN" } });
    }
  } finally {
    if (collectionId) {
      await context.request.delete(
        new URL(`/api/v1/collections/${collectionId}`, site.url).toString(),
      );
    }
    await context.close();
  }
});
