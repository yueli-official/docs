import { expect, test, type BrowserContext, type Page } from "@playwright/test";
import { productSites } from "./contracts";
import {
  capturePageFailures,
  ensureRegisteredE2EIdentity,
  expectNoHorizontalOverflow,
  loginE2E,
  settleNuxt,
} from "./runtime";

const managementRoutes = [
  "/manage",
  "/manage/docs",
  "/manage/collections",
  "/manage/import",
  "/manage/home",
  "/manage/assets",
  "/manage/comments",
  "/manage/authorization",
] as const;

async function saveDiscoverySettings(page: Page): Promise<void> {
  const response = page.waitForResponse(
    (candidate) =>
      candidate.request().method() === "PATCH" &&
      candidate.url().endsWith("/api/v1/home"),
  );
  await page
    .locator("[data-settings-header-actions]")
    .getByRole("button", { name: "保存", exact: true })
    .click();
  expect((await response).ok()).toBeTruthy();
  await expect(page.getByRole("button", { name: "已保存" })).toBeVisible();
}

async function restoreDiscoveryName(
  page: Page,
  siteURL: string,
  originalName: string,
): Promise<void> {
  await page.goto(new URL("/manage/home?section=site", siteURL).toString(), {
    waitUntil: "domcontentloaded",
  });
  await settleNuxt(page);
  const field = page
    .getByRole("textbox", { name: "站点名称", exact: false })
    .first();
  if ((await field.inputValue()) === originalName) return;
  await field.fill(originalName);
  await saveDiscoverySettings(page);
}

async function openFirstImageEditor(page: Page): Promise<void> {
  await page
    .getByRole("button", { name: "编辑图片", exact: true })
    .first()
    .click();
  await expect(page.getByRole("dialog", { name: "编辑图片" })).toBeVisible();
  await expect(
    page
      .getByRole("dialog", { name: "编辑图片" })
      .getByRole("textbox", { name: "标题", exact: false }),
  ).toBeVisible();
}

async function saveImageEditor(page: Page): Promise<void> {
  const response = page.waitForResponse(
    (candidate) =>
      candidate.request().method() === "PATCH" &&
      /\/api\/docs\/admin\/images\/[^/?]+$/.test(candidate.url()),
  );
  await page
    .getByRole("dialog", { name: "编辑图片" })
    .getByRole("button", { name: "保存更改", exact: true })
    .click();
  expect((await response).ok()).toBeTruthy();
  await expect(page.getByRole("dialog", { name: "编辑图片" })).toHaveCount(0);
}

async function searchManagedImages(page: Page, query: string): Promise<void> {
  const search = page.getByPlaceholder("搜索标题、说明或替代文本…");
  await search.fill(query);
  await search.press("Enter");
  await expect(page.getByText(query, { exact: true }).first()).toBeVisible();
}

async function restoreImageTitle(
  page: Page,
  siteURL: string,
  temporaryTitle: string,
  originalTitle: string,
): Promise<void> {
  await page.goto(new URL("/manage/images", siteURL).toString(), {
    waitUntil: "domcontentloaded",
  });
  await settleNuxt(page);
  await searchManagedImages(page, temporaryTitle);
  await openFirstImageEditor(page);
  const title = page
    .getByRole("dialog", { name: "编辑图片" })
    .getByRole("textbox", { name: "标题", exact: false });
  if ((await title.inputValue()) === originalTitle) return;
  await title.fill(originalTitle);
  await saveImageEditor(page);
}

async function saveCollectionSettings(page: Page): Promise<void> {
  const response = page.waitForResponse(
    (candidate) =>
      candidate.request().method() === "PATCH" &&
      /\/api\/docs\/admin\/collections\/[^/?]+$/.test(candidate.url()),
  );
  await page.getByRole("button", { name: "保存专题设置", exact: true }).click();
  expect((await response).ok()).toBeTruthy();
  await expect(page.getByText("专题设置已保存", { exact: true })).toBeVisible();
}

async function restoreCollectionName(
  page: Page,
  collectionURL: string,
  originalName: string,
): Promise<void> {
  await page.goto(collectionURL, { waitUntil: "domcontentloaded" });
  await settleNuxt(page);
  const name = page
    .getByRole("textbox", { name: "名称", exact: false })
    .first();
  if ((await name.inputValue()) === originalName) return;
  await name.fill(originalName);
  await saveCollectionSettings(page);
}

async function authenticatedPage(
  context: BrowserContext,
  siteURL: string,
): Promise<Page> {
  const page = await context.newPage();
  await page.goto(new URL("/manage", siteURL).toString(), {
    waitUntil: "domcontentloaded",
  });
  await expect(page).toHaveURL(new URL("/manage", siteURL).toString());
  await settleNuxt(page);
  return page;
}

export function registerManagementSuite(product: string) {
  for (const site of productSites(product)) {
    test.describe(`${site.slug} (${site.product}) management contract`, () => {
      test("匿名调用 Docs 管理 API 被拒绝", async ({ request }) => {
        for (const path of [
          "/api/docs/admin/overview",
          "/api/docs/admin/site-settings",
          "/api/docs/admin/images?size=1",
          "/api/docs/admin/collections",
        ]) {
          const response = await request.get(
            new URL(path, site.url).toString(),
          );
          expect(
            [401, 403],
            `${path} unexpectedly returned HTTP ${response.status()}`,
          ).toContain(response.status());
        }
      });

      test("普通会员无法进入管理页面或调用管理 API", async ({ browser }) => {
        const context = await ensureRegisteredE2EIdentity(
          browser,
          {
            email: "docs-member-acceptance@example.test",
            password: "Docs-member-acceptance-2026!",
          },
          "Docs Acceptance Member",
          site.url,
        );
        const page = await context.newPage();
        try {
          await page.goto(new URL("/manage", site.url).toString(), {
            waitUntil: "domcontentloaded",
          });
          await expect(page).toHaveURL(new URL("/", site.url).toString());
          await settleNuxt(page);

          for (const path of [
            "/api/docs/admin/overview",
            "/api/docs/admin/site-settings",
            "/api/docs/admin/images?size=1",
            "/api/docs/admin/collections",
          ]) {
            const response = await context.request.get(
              new URL(path, site.url).toString(),
            );
            expect(
              response.status(),
              `${path} unexpectedly returned HTTP ${response.status()}`,
            ).toBe(403);
          }
        } finally {
          await context.close();
        }
      });

      test("管理员可无 5xx 与浏览器错误遍历全部管理入口", async ({
        browser,
      }) => {
        const context = await loginE2E(browser);
        const page = await authenticatedPage(context, site.url);
        const failures = capturePageFailures(page);
        try {
          for (const path of managementRoutes) {
            const failureCount = failures.length;
            const url = new URL(path, site.url).toString();
            const response = await page.goto(url, {
              waitUntil: "domcontentloaded",
            });
            expect(
              response?.ok(),
              `${path} returned ${response?.status()}`,
            ).toBeTruthy();
            await expect(page).toHaveURL(url);
            await expect(page.locator("main h1").first()).toBeVisible();
            await settleNuxt(page);
            expect(
              failures.slice(failureCount),
              `${path} 不应产生浏览器错误`,
            ).toEqual([]);
          }
          expect(failures).toEqual([]);
        } finally {
          await context.close();
        }
      });

      test("文档列表提供明确的编辑入口", async ({ browser }, testInfo) => {
        const context = await loginE2E(
          browser,
          { viewport: { width: 1440, height: 900 } },
          undefined,
          site.url,
        );
        const page = await context.newPage();
        const failures = capturePageFailures(page);
        try {
          await page.goto(new URL("/manage/docs", site.url).toString(), {
            waitUntil: "domcontentloaded",
          });
          await page.waitForLoadState("networkidle");
          await expect(
            page.getByRole("link", { name: "新建文档", exact: true }),
          ).toBeVisible();
          const quickEdit = page
            .getByRole("button", { name: /^快速编辑：/u })
            .first();
          const fullEdit = page
            .getByRole("link", { name: /^编辑文档：/u })
            .first();
          await expect(quickEdit).toBeVisible();
          await expect(fullEdit).toBeVisible();
          await page.screenshot({
            path: testInfo.outputPath("documents-desktop.png"),
            fullPage: false,
          });
          await fullEdit.click();
          await expect(page).toHaveURL(/\/manage\/docs\/.+/u);
          await page.waitForLoadState("networkidle");
          await expect(
            page.getByRole("button", { name: "保存", exact: true }),
          ).toBeVisible();
          await expect(
            page.locator("[data-admin-console-breadcrumb]"),
          ).toHaveCount(0);
          const [canvasBox, commandbarBox] = await Promise.all([
            page.locator("[data-admin-console-canvas]").boundingBox(),
            page.locator("[data-docs-editor-commandbar]").boundingBox(),
          ]);
          expect(commandbarBox?.y).toBe(canvasBox?.y);
          expect(commandbarBox?.x).toBe(canvasBox?.x);
          await page.screenshot({
            path: testInfo.outputPath("document-editor-desktop.png"),
            fullPage: false,
          });
          await page.setViewportSize({ width: 390, height: 844 });
          await expectNoHorizontalOverflow(page, "Docs 文档编辑器");
          await page.screenshot({
            path: testInfo.outputPath("document-editor-mobile.png"),
            fullPage: false,
          });
          await page.evaluate(() => {
            window.localStorage.setItem("nuxt-color-mode", "dark");
          });
          await page.reload({ waitUntil: "networkidle" });
          await expect(page.locator("html")).toHaveClass(/\bdark\b/u);
          await expectNoHorizontalOverflow(page, "Docs 深色文档编辑器");
          await page.screenshot({
            path: testInfo.outputPath("document-editor-mobile-dark.png"),
            fullPage: false,
          });
          expect(failures).toEqual([]);
        } finally {
          await context.close();
        }
      });

      test("文档设置收纳辅助编辑信息", async ({ browser }, testInfo) => {
        const context = await loginE2E(
          browser,
          { viewport: { width: 1440, height: 900 } },
          undefined,
          site.url,
        );
        const page = await context.newPage();
        const failures = capturePageFailures(page);
        try {
          await page.goto(
            new URL(
              "/manage/docs/operations/troubleshooting",
              site.url,
            ).toString(),
            { waitUntil: "networkidle" },
          );
          await page.getByRole("button", { name: "文档设置" }).click();
          const settings = page.getByRole("dialog", { name: "文档设置" });
          await expect(settings).toBeVisible();
          await expect(
            settings.getByRole("heading", { name: "摘要与链接" }),
          ).toBeVisible();
          await expect(
            settings.getByRole("heading", { name: "归属与路径" }),
          ).toBeVisible();
          await expect(
            settings.getByRole("heading", { name: "搜索优化" }),
          ).toBeVisible();
          await page.screenshot({
            path: testInfo.outputPath("document-settings-desktop.png"),
            fullPage: false,
          });
          expect(failures).toEqual([]);
        } finally {
          await context.close();
        }
      });

      test("评论后台使用公开用户资料与标准列布局", async ({
        browser,
      }, testInfo) => {
        const context = await loginE2E(
          browser,
          { viewport: { width: 1440, height: 900 } },
          undefined,
          site.url,
        );
        const page = await context.newPage();
        try {
          await page.goto(new URL("/manage/comments", site.url).toString(), {
            waitUntil: "networkidle",
          });
          const comments = page.locator("[data-manage-comments]");
          await expect(comments).toBeVisible();
          for (const heading of [
            "评论",
            "来源",
            "用户",
            "状态",
            "评论日期",
            "操作",
          ]) {
            await expect(
              comments.getByText(heading, { exact: true }).first(),
            ).toBeVisible();
          }
          const firstComment = comments.locator("article").first();
          await expect(firstComment).toContainText("测试管理员");
          await expect(
            firstComment.locator("img:visible").first(),
          ).toHaveAttribute("src", /\S+/u);
          await expect(
            firstComment.getByText("成员", { exact: true }),
          ).toHaveCount(0);
          await page.screenshot({
            path: testInfo.outputPath("comments-desktop.png"),
            fullPage: false,
          });
          await page.setViewportSize({ width: 390, height: 844 });
          await expectNoHorizontalOverflow(page, "Docs 评论后台");
          await page.screenshot({
            path: testInfo.outputPath("comments-mobile.png"),
            fullPage: false,
          });
        } finally {
          await context.close();
        }
      });

      test("站点设置把保存动作放在页头", async ({ browser }, testInfo) => {
        const context = await loginE2E(
          browser,
          { viewport: { width: 1440, height: 900 } },
          undefined,
          site.url,
        );
        const page = await context.newPage();
        try {
          await page.goto(
            new URL("/manage/home?section=site", site.url).toString(),
            { waitUntil: "domcontentloaded" },
          );
          await page.waitForLoadState("networkidle");
          const save = page.getByRole("button", { name: /^保存/u }).first();
          const firstField = page.getByRole("textbox").first();
          await expect(save).toBeVisible();
          await expect(firstField).toBeVisible();
          const saveBox = await save.boundingBox();
          const fieldBox = await firstField.boundingBox();
          expect(saveBox?.y).toBeLessThan(fieldBox!.y);
          await expect(page.locator('[data-manage-dock="save"]')).toHaveCount(
            0,
          );
          await firstField.fill(`${await firstField.inputValue()} `);
          await expect(save).toBeEnabled();
          await page.screenshot({
            path: testInfo.outputPath("site-settings-header-save.png"),
            fullPage: false,
          });
        } finally {
          await context.close();
        }
      });

      test("设置失败 Toast 保持右下角紧凑宽度", async ({
        browser,
      }, testInfo) => {
        const context = await loginE2E(browser, {}, undefined, site.url);
        const page = await context.newPage();
        try {
          await page.route("**/api/**", async (route) => {
            if (route.request().method() === "PATCH") {
              await route.fulfill({
                status: 500,
                contentType: "application/json",
                body: JSON.stringify({ message: "请稍后重试" }),
              });
              return;
            }
            await route.continue();
          });
          await page.goto(
            new URL("/manage/home?section=site", site.url).toString(),
            { waitUntil: "domcontentloaded" },
          );
          await page.waitForLoadState("networkidle");
          const siteName = page.getByRole("textbox", { name: "站点名称" });
          await siteName.fill(`${await siteName.inputValue()} `);
          await page.getByRole("button", { name: /^保存/u }).first().click();
          const message = page.getByText("设置保存失败", { exact: true });
          await expect(message).toBeVisible();
          const toast = message.locator(
            "xpath=ancestor::*[@role='status' or @role='alert' or self::li][1]",
          );
          await expect(toast).toBeVisible();
          const box = await toast.boundingBox();
          const viewportWidth = await page.evaluate(() => window.innerWidth);
          expect(box?.width).toBeLessThanOrEqual(420);
          expect(box!.x + box!.width).toBeGreaterThanOrEqual(
            viewportWidth - 24,
          );
          await page.screenshot({
            path: testInfo.outputPath("settings-error-toast.png"),
            fullPage: false,
          });
        } finally {
          await context.close();
        }
      });

      test("站点设置保存后持久化且测试结束恢复原值", async ({ browser }) => {
        const context = await loginE2E(browser);
        const page = await authenticatedPage(context, site.url);
        let originalName = "";
        try {
          await page.goto(
            new URL("/manage/home?section=site", site.url).toString(),
            {
              waitUntil: "domcontentloaded",
            },
          );
          await settleNuxt(page);
          const field = page
            .getByRole("textbox", { name: "站点名称", exact: false })
            .first();
          originalName = await field.inputValue();
          const temporaryName = `${originalName.slice(0, 65)} · 验收`;
          await field.fill(temporaryName);
          await saveDiscoverySettings(page);

          await page.reload({ waitUntil: "domcontentloaded" });
          await settleNuxt(page);
          await expect(
            page
              .getByRole("textbox", { name: "站点名称", exact: false })
              .first(),
          ).toHaveValue(temporaryName);

          await page
            .getByRole("textbox", { name: "站点名称", exact: false })
            .first()
            .fill(originalName);
          await saveDiscoverySettings(page);
          await page.reload({ waitUntil: "domcontentloaded" });
          await settleNuxt(page);
          await expect(
            page
              .getByRole("textbox", { name: "站点名称", exact: false })
              .first(),
          ).toHaveValue(originalName);
        } finally {
          if (originalName)
            await restoreDiscoveryName(page, site.url, originalName).catch(
              () => undefined,
            );
          await context.close();
        }
      });

      test("图片编辑器使用可读分类标签且选择器有明确名称", async ({
        browser,
      }) => {
        const context = await loginE2E(browser);
        const page = await authenticatedPage(context, site.url);
        try {
          await page.goto(new URL("/manage/images", site.url).toString(), {
            waitUntil: "domcontentloaded",
          });
          await settleNuxt(page);
          await openFirstImageEditor(page);

          const dialog = page.getByRole("dialog", { name: "编辑图片" });
          const category = dialog.getByRole("button", {
            name: "主分类",
            exact: true,
          });
          const tags = dialog.getByRole("button", {
            name: "标签",
            exact: true,
          });
          await expect(category).toBeVisible();
          await expect(tags).toBeVisible();

          const rawID =
            /(?:^[A-Za-z0-9_-]{22}$)|(?:^[0-9a-f]{8}-[0-9a-f-]{27}$)/iu;
          expect((await category.textContent())?.trim()).not.toMatch(rawID);
          expect((await tags.textContent())?.trim()).not.toMatch(rawID);

          await category.click();
          const categoryLabels = (
            await page.getByRole("option").allTextContents()
          ).map((value) => value.trim());
          expect(categoryLabels).toEqual(
            expect.arrayContaining(["壁纸", "插画", "摄影"]),
          );
        } finally {
          await context.close();
        }
      });

      test("图片标题保存后可检索且测试结束恢复原值", async ({ browser }) => {
        const context = await loginE2E(browser);
        const page = await authenticatedPage(context, site.url);
        let originalTitle = "";
        let temporaryTitle = "";
        try {
          await page.goto(new URL("/manage/images", site.url).toString(), {
            waitUntil: "domcontentloaded",
          });
          await settleNuxt(page);
          await openFirstImageEditor(page);
          const title = page
            .getByRole("dialog", { name: "编辑图片" })
            .getByRole("textbox", { name: "标题", exact: false });
          originalTitle = await title.inputValue();
          temporaryTitle = `${originalTitle.slice(0, 140)} · 验收`;
          await title.fill(temporaryTitle);
          await saveImageEditor(page);

          await searchManagedImages(page, temporaryTitle);
          await openFirstImageEditor(page);
          await expect(
            page
              .getByRole("dialog", { name: "编辑图片" })
              .getByRole("textbox", { name: "标题", exact: false }),
          ).toHaveValue(temporaryTitle);
          await page
            .getByRole("dialog", { name: "编辑图片" })
            .getByRole("textbox", { name: "标题", exact: false })
            .fill(originalTitle);
          await saveImageEditor(page);

          await searchManagedImages(page, originalTitle);
          await expect(
            page.getByText(originalTitle, { exact: true }).first(),
          ).toBeVisible();
        } finally {
          if (originalTitle && temporaryTitle)
            await restoreImageTitle(
              page,
              site.url,
              temporaryTitle,
              originalTitle,
            ).catch(() => undefined);
          await context.close();
        }
      });

      test("专题名称保存后刷新持久化且测试结束恢复原值", async ({
        browser,
      }) => {
        const context = await loginE2E(browser);
        const page = await authenticatedPage(context, site.url);
        let originalName = "";
        let collectionURL = "";
        try {
          await page.goto(new URL("/manage/collections", site.url).toString(), {
            waitUntil: "domcontentloaded",
          });
          await settleNuxt(page);
          const firstCollection = page
            .getByRole("link", { name: /^编辑专题：/u })
            .first();
          await expect(
            firstCollection,
            "验收夹具必须至少包含一个可编辑专题",
          ).toBeVisible();
          const href = await firstCollection.getAttribute("href");
          expect(href).toBeTruthy();
          collectionURL = new URL(href!, site.url).toString();

          await page.goto(collectionURL, { waitUntil: "domcontentloaded" });
          await settleNuxt(page);
          const name = page
            .getByRole("textbox", { name: "名称", exact: false })
            .first();
          originalName = await name.inputValue();
          const temporaryName = `${originalName.slice(0, 90)} · 验收`;
          await name.fill(temporaryName);
          await saveCollectionSettings(page);

          await page.reload({ waitUntil: "domcontentloaded" });
          await settleNuxt(page);
          await expect(
            page.getByRole("textbox", { name: "名称", exact: false }).first(),
          ).toHaveValue(temporaryName);

          await page
            .getByRole("textbox", { name: "名称", exact: false })
            .first()
            .fill(originalName);
          await saveCollectionSettings(page);
          await page.reload({ waitUntil: "domcontentloaded" });
          await settleNuxt(page);
          await expect(
            page.getByRole("textbox", { name: "名称", exact: false }).first(),
          ).toHaveValue(originalName);
        } finally {
          if (collectionURL && originalName)
            await restoreCollectionName(
              page,
              collectionURL,
              originalName,
            ).catch(() => undefined);
          await context.close();
        }
      });

      test("资源策略展示消费者注册且不提供第二套配置入口", async ({
        browser,
      }) => {
        const context = await loginE2E(browser, {}, undefined, site.url);
        const page = await context.newPage();
        try {
          await page.goto(new URL("/manage/assets", site.url).toString(), {
            waitUntil: "networkidle",
          });
          await expect(
            page.getByRole("heading", { name: "资源策略", exact: true }),
          ).toBeVisible();
          await expect(
            page.locator("[data-asset-registration-summary]"),
          ).toBeVisible();
          await expect(
            page.getByText("docs-collection-cover", { exact: true }),
          ).toBeVisible();
          await expect(page.getByRole("button", { name: "保存" })).toHaveCount(
            0,
          );
          await expect(page.getByLabel("站点名称")).toHaveCount(0);
        } finally {
          await context.close();
        }
      });

      test("分类状态变更必须先生成影响预览且不会直接改写数据", async ({
        browser,
      }) => {
        const context = await loginE2E(browser);
        const page = await authenticatedPage(context, site.url);
        try {
          await page.goto(
            new URL("/manage/classification", site.url).toString(),
            { waitUntil: "domcontentloaded" },
          );
          await settleNuxt(page);
          const firstActions = page
            .getByRole("button", { name: /^更多分类操作：/u })
            .first();
          await expect(
            firstActions,
            "验收夹具必须至少包含一个可治理分类",
          ).toBeVisible();
          await firstActions.click();

          const previewResponse = page.waitForResponse(
            (candidate) =>
              candidate.request().method() === "POST" &&
              /\/api\/docs\/admin\/classification\/governance\/preview$/.test(
                candidate.url(),
              ),
          );
          await page.getByRole("menuitem", { name: "停用分类" }).click();
          expect((await previewResponse).ok()).toBeTruthy();

          const dialog = page.getByRole("dialog", { name: "治理影响预览" });
          await expect(dialog).toBeVisible();
          await expect(
            dialog.getByText("planned", { exact: true }),
          ).toBeVisible();
          await expect(
            dialog.getByRole("button", { name: "执行计划", exact: true }),
          ).toBeVisible();
          await dialog
            .getByRole("button", { name: "关闭", exact: true })
            .click();
          await expect(dialog).toHaveCount(0);
        } finally {
          await context.close();
        }
      });
    });
  }
}
