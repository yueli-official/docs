import { expect, test } from "@playwright/test";
import { capturePageFailures, loginE2E, settleNuxt } from "./runtime";

test("seeded Docs administrator retains access to the management console", async ({ browser }) => {
  const siteURL = process.env.DOCS_E2E_URL!;
  const context = await loginE2E(browser, {}, undefined, siteURL);
  const me = await context.request.get(new URL("/api/v1/me", siteURL).toString());
  expect(me.ok()).toBeTruthy();
  expect(await me.json()).toMatchObject({ me: { isAdministrator: true } });
  const page = await context.newPage();
  const failures = capturePageFailures(page);
  await page.goto(siteURL);
  await settleNuxt(page);
  await page.getByRole("button", { name: /测试管理员/ }).click();
  await expect(page.getByRole("menuitem", { name: "控制台" })).toBeVisible();
  await expect(page.getByRole("menuitem", { name: "申请成为作者" })).toHaveCount(0);
  await page.goto(new URL("/manage", siteURL).toString());
  await settleNuxt(page);
  await expect(page.getByRole("heading", { name: "控制台" })).toBeVisible();
  expect(failures.filter((failure) => failure.includes("Cannot read properties of null"))).toEqual([]);
  await context.close();
});
