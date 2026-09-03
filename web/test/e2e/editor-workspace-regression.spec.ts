import { expect, test } from "@playwright/test";
import { loginE2E, settleNuxt } from "./runtime";

const editorPath = "/manage/docs/ae-scripting/general/application";

test("editor path copy works on the LAN HTTP preview", async ({ browser }) => {
  const siteURL = process.env.DOCS_E2E_URL!;
  const context = await loginE2E(browser, { permissions: ["clipboard-read", "clipboard-write"] }, undefined, siteURL);
  const page = await context.newPage();
  await page.goto(new URL(editorPath, siteURL).toString());
  await settleNuxt(page);
  await page.getByRole("button", { name: "文档设置" }).click();
  const copy = page.getByRole("button", { name: /复制/ }).first();
  await expect(copy).toBeVisible();
  await copy.click();
  await expect(page.getByRole("button", { name: "公开链接已复制" })).toBeVisible();
  await context.close();
});

test("long editor keeps its writing controls available and supports immersive collaboration", async ({ browser }) => {
  const siteURL = process.env.DOCS_E2E_URL!;
  const context = await loginE2E(browser, { viewport: { width: 1440, height: 900 } }, undefined, siteURL);
  const page = await context.newPage();
  await page.goto(new URL(editorPath, siteURL).toString());
  await settleNuxt(page);
  await expect(page.getByRole("button", { name: "沉浸式协作" })).toBeVisible();
  await expect(page.getByLabel("摘要")).toHaveCount(0);
  await expect(page.getByLabel("路径标识")).toHaveCount(0);
  await page.getByRole("button", { name: "文档设置" }).click();
  await page.getByRole("tab", { name: "组织" }).click();
  const iconResults = page.locator("[data-admin-icon-results]");
  await expect(iconResults).toBeVisible();
  const iconColumns = await iconResults.evaluate((element) => getComputedStyle(element).gridTemplateColumns.split(" ").length);
  expect(iconColumns).toBeGreaterThan(5);
  await page.getByRole("button", { name: "文档设置" }).click();
  await page.locator("#manage-main").evaluate((element) => { element.scrollTop = element.scrollHeight; });
  await expect(page.getByLabel("文档标题")).toBeInViewport();
  await expect(page.getByRole("toolbar").first()).toBeInViewport();
  await page.getByRole("button", { name: "沉浸式协作" }).click();
  await expect(page.locator("[data-docs-editor-workspace]")).toHaveAttribute("data-collaboration-mode", "immersive");
  await expect(page.getByRole("toolbar").first()).toBeInViewport();
  await expect(page.locator("[data-docs-editor-commandbar]")).toHaveCount(0);
  await expect(page.getByLabel("文档标题")).toHaveCount(0);
  await expect(page.getByRole("button", { name: "退出沉浸式协作" })).toBeVisible();

  const mobile = await context.newPage();
  await mobile.setViewportSize({ width: 390, height: 844 });
  await mobile.goto(new URL(editorPath, siteURL).toString());
  await settleNuxt(mobile);
  await mobile.locator("#manage-main").evaluate((element) => { element.scrollTop = element.scrollHeight; });
  await expect(mobile.getByLabel("文档标题")).toBeInViewport();
  await expect(mobile.getByRole("toolbar").first()).toBeInViewport();
  await mobile.getByRole("button", { name: "沉浸式协作" }).click();
  await expect(mobile.locator("[data-docs-editor-workspace]")).toHaveAttribute("data-collaboration-mode", "immersive");
  await expect(mobile.locator("[data-docs-editor-commandbar]")).toHaveCount(0);
  await expect(mobile.getByRole("button", { name: "退出沉浸式协作" })).toBeVisible();
  await context.close();
});
