import { expect, test } from "@playwright/test";
import { settleNuxt } from "./runtime";

test("code block copy works on the LAN HTTP preview and confirms success", async ({ browser }) => {
  const siteURL = process.env.DOCS_E2E_URL!;
  const context = await browser.newContext({ permissions: ["clipboard-read", "clipboard-write"] });
  const collectionsResponse = await context.request.get(new URL("/api/v1/collections", siteURL).toString());
  const collections = (await collectionsResponse.json()).items as Array<{ slug: string; title: string }>;
  const collection = collections.find((item) => item.title === "ae脚本文档");
  expect(collection).toBeTruthy();

  const page = await context.newPage();
  await page.goto(new URL(`/${collection!.slug}/general/globals`, siteURL).toString());
  await settleNuxt(page);
  const copy = page.getByRole("button", { name: /复制/ }).first();
  await expect(copy).toBeVisible();
  await copy.click();
  await expect(page.getByText(/已复制|复制成功/).first()).toBeVisible();
  const clipboard = await page.evaluate(() => navigator.clipboard?.readText());
  if (clipboard !== undefined) expect(clipboard.trim().length).toBeGreaterThan(0);
  await context.close();
});
