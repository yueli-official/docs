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

  await page.route("**/api/v1/imports/docs/*/confirm", (route) =>
    route.fulfill({ status: 500, contentType: "application/json", body: JSON.stringify({ message: "确认导入失败测试" }) }),
  );
  await page.getByRole("button", { name: "确认导入" }).click();
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
