import { expect, test } from "@playwright/test";
import { capturePageFailures, loginE2E, settleNuxt } from "./runtime";

test("creates a localized release and preflights the AE bilingual package", async ({ browser }) => {
  const siteURL = process.env.DOCS_E2E_URL!;
  const context = await loginE2E(browser, {}, undefined, siteURL);
  const collectionSlug = `import-create-${Date.now()}`;
  const page = await context.newPage();
  const failures = capturePageFailures(page);

  await page.goto(new URL("/manage/collections", siteURL).toString());
  await settleNuxt(page);
  await page.getByRole("button", { name: "新建文档集" }).click();
  await expect(page.getByLabel("默认语言")).toBeVisible();
  await expect(page.getByLabel("版本号")).toHaveValue("1.0.0");
  await page.getByLabel("标题").fill("导入创建验收");
  await page.getByLabel("路径标识").fill(collectionSlug);

  const createResponsePromise = page.waitForResponse(
    (response) => response.url().endsWith("/api/v1/collections") && response.request().method() === "POST",
  );
  await page.getByRole("button", { name: "创建" }).click();
  const createResponse = await createResponsePromise;
  expect(createResponse.ok(), await createResponse.text()).toBeTruthy();
  const created = (await createResponse.json()).collection as { id: string; semanticVersion: string };
  expect(created.semanticVersion).toBe("1.0.0");
  const locales = await context.request.get(new URL(`/api/v1/manage/collections/${created.id}/locales`, siteURL).toString());
  expect(locales.ok()).toBeTruthy();
  expect(await locales.json()).toMatchObject({ items: [{ locale: "zh-CN", isDefault: true }] });
  const deleted = await context.request.delete(new URL(`/api/v1/collections/${created.id}`, siteURL).toString());
  expect(deleted.ok()).toBeTruthy();

  await page.goto(new URL("/manage/import", siteURL).toString());
  await settleNuxt(page);
  await page.locator('input[type="file"]').setInputFiles("E:/projects/yozya/docs/exports/ae-scripting-docs-v1.zip");
  await page.getByRole("button", { name: "上传并预检" }).click();
  await expect(page.getByText("可导入", { exact: true })).toBeVisible({ timeout: 90_000 });
  await expect(page.getByText("144", { exact: true })).toBeVisible();
  await expect(page.getByText("4", { exact: true })).toBeVisible();
  await expect(page.getByText("预检问题", { exact: true })).toHaveCount(0);
  const confirmResponsePromise = page.waitForResponse(
    (response) => /\/api\/v1\/imports\/docs\/[0-9a-f-]+\/confirm$/.test(response.url()),
    { timeout: 90_000 },
  );
  await page.getByRole("button", { name: "确认导入" }).click();
  const confirmResponse = await confirmResponsePromise;
  expect(confirmResponse.ok(), await confirmResponse.text()).toBeTruthy();
  expect(await confirmResponse.json()).toMatchObject({ batch: { status: "completed" } });
  await expect(page).toHaveURL(/\/manage\/import\/[0-9a-f-]+$/, { timeout: 90_000 });
  expect(failures.filter((failure) => /Cannot read properties|image_missing/.test(failure))).toEqual([]);
  await context.close();
});

test("confirm failure is shown once and remains until dismissed", async ({ browser }) => {
  const siteURL = process.env.DOCS_E2E_URL!;
  const context = await loginE2E(browser, {}, undefined, siteURL);
  const page = await context.newPage();
  await page.goto(new URL("/manage/import", siteURL).toString());
  await settleNuxt(page);
  await page.locator('input[type="file"]').setInputFiles("E:/projects/yozya/docs/exports/ae-scripting-docs-v1.zip");
  await page.getByRole("button", { name: "上传并预检" }).click();
  await expect(page.getByText("可导入", { exact: true })).toBeVisible({ timeout: 90_000 });

  await page.route("**/api/v1/imports/docs/*/confirm", async (route) => {
    await new Promise((resolve) => setTimeout(resolve, 1_000));
    await route.fulfill({ status: 500, contentType: "application/json", body: JSON.stringify({ message: "确认导入失败测试" }) });
  });
  await page.getByRole("button", { name: "确认导入" }).click();
  await expect(page.getByText("后台处理中", { exact: true })).toBeVisible();
  await expect(page.getByText("任务正在后台执行。可以离开本页，稍后从“最近导入”返回查看结果。")).toBeVisible();
  await expect(page.getByText("导入失败", { exact: true })).toHaveCount(1);
  await expect(page.getByText("确认导入失败", { exact: true })).toHaveCount(0);
  await page.waitForTimeout(5_000);
  await expect(page.getByText("导入失败", { exact: true })).toBeVisible();
  await context.close();
});

test("converted AE package preflights with alerts and images", async ({ browser }) => {
  const siteURL = process.env.DOCS_E2E_URL!;
  const context = await loginE2E(browser, {}, undefined, siteURL);
  const page = await context.newPage();
  await page.goto(new URL("/manage/import", siteURL).toString());
  await settleNuxt(page);
  await page.locator('input[type="file"]').setInputFiles("E:/projects/yozya/docs/exports/ae-scripting-docs-v1.zip");
  await page.getByRole("button", { name: "上传并预检" }).click();
  await expect(page.getByText("可导入", { exact: true })).toBeVisible({ timeout: 90_000 });
  await expect(page.getByText("144", { exact: true })).toBeVisible();
  await expect(page.getByText("4", { exact: true })).toBeVisible();
  await expect(page.getByText("预检问题", { exact: true })).toHaveCount(0);
  await context.close();
});

test("large Sapphire package reaches a completed preflight", async ({ browser }) => {
  test.slow();
  const siteURL = process.env.DOCS_E2E_URL!;
  const context = await loginE2E(browser, {}, undefined, siteURL);
  const collectionSlug = `sapphire-import-${Date.now()}`;
  const createResponse = await context.request.post(new URL("/api/v1/collections", siteURL).toString(), {
    data: { title: "Sapphire 导入验收", slug: collectionSlug, defaultLocale: "zh-CN", semanticVersion: "1.0.0" },
  });
  expect(createResponse.ok(), await createResponse.text()).toBeTruthy();
  const collection = (await createResponse.json()).collection as { id: string };
  const page = await context.newPage();
  try {
    await page.goto(new URL("/manage/import", siteURL).toString());
    await settleNuxt(page);
    await page.getByLabel("目标文档集").click();
    await page.locator('[data-slot="itemLabel"]').getByText("Sapphire 导入验收", { exact: true }).click();
    await page.locator('input[type="file"]').setInputFiles("E:/projects/yozya/docs/exports/sapphire-docs-v1.zip");
    await page.getByRole("button", { name: "上传并预检" }).click();
    await expect(page.getByText("可导入", { exact: true })).toBeVisible({ timeout: 180_000 });
    await expect(page.getByText("预检问题", { exact: true })).toHaveCount(0);
    const confirmResponsePromise = page.waitForResponse(
      (response) => /\/api\/v1\/imports\/docs\/[0-9a-f-]+\/confirm$/.test(response.url()),
      { timeout: 180_000 },
    );
    await page.getByRole("button", { name: "确认导入" }).click();
    const confirmResponse = await confirmResponsePromise;
    expect(confirmResponse.ok(), await confirmResponse.text()).toBeTruthy();
  } finally {
    await context.request.delete(new URL(`/api/v1/collections/${collection.id}`, siteURL).toString());
    await context.close();
  }
});

test("large Houdini package reaches a completed preflight", async ({ browser }) => {
  test.slow();
  const siteURL = process.env.DOCS_E2E_URL!;
  const context = await loginE2E(browser, {}, undefined, siteURL);
  const collectionSlug = `houdini-import-${Date.now()}`;
  const createResponse = await context.request.post(new URL("/api/v1/collections", siteURL).toString(), {
    data: { title: "Houdini 导入验收", slug: collectionSlug, defaultLocale: "zh-CN", semanticVersion: "1.0.0" },
  });
  expect(createResponse.ok(), await createResponse.text()).toBeTruthy();
  const collection = (await createResponse.json()).collection as { id: string };
  const page = await context.newPage();
  try {
    await page.goto(new URL("/manage/import", siteURL).toString());
    await settleNuxt(page);
    await page.getByLabel("目标文档集").click();
    await page.locator('[data-slot="itemLabel"]').getByText("Houdini 导入验收", { exact: true }).click();
    await page.locator('input[type="file"]').setInputFiles("E:/projects/yozya/docs/exports/houdini-vex-docs-v1.zip");
    const responsePromise = page.waitForResponse(
      (response) => response.url().endsWith("/api/v1/imports/docs") && response.request().method() === "POST",
      { timeout: 240_000 },
    );
    await page.getByRole("button", { name: "上传并预检" }).click();
    const response = await responsePromise;
    expect(response.ok(), await response.text()).toBeTruthy();
    await expect(page.getByText("可导入", { exact: true })).toBeVisible({ timeout: 240_000 });
    const confirmResponsePromise = page.waitForResponse(
      (item) => /\/api\/v1\/imports\/docs\/[0-9a-f-]+\/confirm$/.test(item.url()),
      { timeout: 240_000 },
    );
    await page.getByRole("button", { name: "确认导入" }).click();
    const confirmResponse = await confirmResponsePromise;
    expect(confirmResponse.ok(), await confirmResponse.text()).toBeTruthy();
  } finally {
    await context.request.delete(new URL(`/api/v1/collections/${collection.id}`, siteURL).toString());
    await context.close();
  }
});

test("creates and selects a document collection without leaving import", async ({ browser }) => {
  const siteURL = process.env.DOCS_E2E_URL!;
  const context = await loginE2E(browser, {}, undefined, siteURL);
  const page = await context.newPage();
  const failures = capturePageFailures(page);
  const collectionSlug = `inline-import-${Date.now()}`;
  let createdID = "";

  await page.goto(new URL("/manage/import", siteURL).toString());
  await settleNuxt(page);
  await page.getByRole("button", { name: "新建", exact: true }).click();
  const creator = page.locator("[data-import-collection-creator]");
  await creator.getByLabel("标题").fill("导入页即时创建验收");
  await creator.getByLabel("路径标识").fill(collectionSlug);
  await expect(creator.getByLabel("首个版本")).toHaveValue("1.0.0");

  const responsePromise = page.waitForResponse(
    (response) => response.url().endsWith("/api/v1/collections") && response.request().method() === "POST",
  );
  await creator.getByRole("button", { name: "创建并选中" }).click();
  const response = await responsePromise;
  expect(response.ok(), await response.text()).toBeTruthy();
  const created = (await response.json()).collection as { id: string; slug: string };
  createdID = created.id;
  expect(created.slug).toBe(collectionSlug);
  await expect(page.getByText("新建并选作文档集")).toHaveCount(0);
  await expect(page.getByText("文档集已创建", { exact: true })).toBeVisible();
  await expect(page.getByLabel("目标文档集")).toContainText("导入页即时创建验收");
  expect(failures).toEqual([]);

  const mobile = await context.newPage();
  await mobile.setViewportSize({ width: 390, height: 844 });
  await mobile.goto(new URL("/manage/import", siteURL).toString());
  await settleNuxt(mobile);
  await mobile.getByRole("button", { name: "新建", exact: true }).click();
  const mobileCreator = mobile.locator("[data-import-collection-creator]");
  await expect(mobileCreator.getByLabel("标题")).toBeVisible();
  await expect(mobileCreator.getByLabel("路径标识")).toBeVisible();
  await expect(mobileCreator.getByLabel("默认语言")).toBeVisible();
  await expect(mobileCreator.getByLabel("首个版本")).toBeVisible();
  await mobile.getByRole("button", { name: "取消新建文档集" }).click();
  await expect(mobileCreator).toHaveCount(0);

  if (createdID) {
    const deleted = await context.request.delete(new URL(`/api/v1/collections/${createdID}`, siteURL).toString());
    expect(deleted.ok()).toBeTruthy();
  }
  await context.close();
});

test("import history is server paginated", async ({ browser }) => {
  const siteURL = process.env.DOCS_E2E_URL!;
  const context = await loginE2E(browser, {}, undefined, siteURL);
  const first = await context.request.get(new URL("/api/v1/imports/docs?page=1&size=2", siteURL).toString());
  expect(first.ok(), await first.text()).toBeTruthy();
  const firstPage = await first.json() as { items: Array<{ id: string }>; total: number; page: number; size: number };
  expect(firstPage).toMatchObject({ page: 1, size: 2 });
  expect(firstPage.items.length).toBeLessThanOrEqual(2);
  expect(firstPage.total).toBeGreaterThanOrEqual(firstPage.items.length);
  if (firstPage.total > 2) {
    const second = await context.request.get(new URL("/api/v1/imports/docs?page=2&size=2", siteURL).toString());
    expect(second.ok(), await second.text()).toBeTruthy();
    const secondPage = await second.json() as { items: Array<{ id: string }>; page: number; size: number };
    expect(secondPage).toMatchObject({ page: 2, size: 2 });
    expect(secondPage.items.some((item) => firstPage.items.some((firstItem) => firstItem.id === item.id))).toBeFalsy();
  }
  const page = await context.newPage();
  await page.goto(new URL("/manage/import", siteURL).toString());
  await settleNuxt(page);
  await expect(page.getByText("预检会保存为导入批次；确认后由后台继续执行，可以离开页面并从最近导入查看进度。")).toBeVisible();
  await expect(page.getByText(`共 ${firstPage.total} 个批次`)).toBeVisible();
  expect(await page.evaluate(() => document.documentElement.scrollWidth <= document.documentElement.clientWidth)).toBeTruthy();
  await page.screenshot({ path: "test-results/import-task-center-desktop.png", fullPage: true });

  const mobile = await context.newPage();
  await mobile.setViewportSize({ width: 390, height: 844 });
  await mobile.goto(new URL("/manage/import", siteURL).toString());
  await settleNuxt(mobile);
  await expect(mobile.getByText(`共 ${firstPage.total} 个批次`)).toBeVisible();
  expect(await mobile.evaluate(() => document.documentElement.scrollWidth <= document.documentElement.clientWidth)).toBeTruthy();
  await mobile.screenshot({ path: "test-results/import-task-center-mobile.png", fullPage: true });
  await context.close();
});

test("running import history uses a reduced-motion-safe spinner", async ({ browser }) => {
  const siteURL = process.env.DOCS_E2E_URL!;
  const context = await loginE2E(browser, {}, undefined, siteURL);
  const page = await context.newPage();
  await page.route("**/api/v1/imports/docs?page=1&size=8", (route) => route.fulfill({
    status: 200,
    contentType: "application/json",
    body: JSON.stringify({
      items: [{
        id: "01a-running-test",
        collectionId: "collection-test",
        versionId: "version-test",
        defaultLocale: "zh-CN",
        mode: "upsert",
        status: "running",
        errorMessage: "",
        summary: { creates: 20, updates: 0, archives: 0, skips: 0, conflicts: 0, errors: 0, images: 2, warnings: 0, blocking: false, issues: [] },
        createdAt: "2026-09-04T02:27:00+08:00",
        updatedAt: "2026-09-04T02:27:00+08:00",
      }],
      total: 1,
      page: 1,
      size: 8,
    }),
  }));
  await page.goto(new URL("/manage/import", siteURL).toString());
  await settleNuxt(page);
  const icon = page.locator('[data-import-status="running"] [data-slot="leadingIcon"]');
  await expect(icon).toHaveClass(/animate-spin/);
  expect(await icon.evaluate((element) => getComputedStyle(element).animationName)).not.toBe("none");
  await context.close();
});
