import { expect, test } from "@playwright/test";
import { loginE2E, settleNuxt } from "./runtime";

test("long reading TOC keeps only meaningful headings in a compact scroll rail", async ({ browser }) => {
  const siteURL = process.env.DOCS_E2E_URL!;
  const context = await loginE2E(browser, { viewport: { width: 1440, height: 900 } }, undefined, siteURL);
  const collectionsResponse = await context.request.get(new URL("/api/v1/collections", siteURL).toString());
  const collections = (await collectionsResponse.json()).items as Array<{ slug: string; title: string }>;
  const collection = collections.find((item) => item.title === "ae脚本文档");
  expect(collection).toBeTruthy();

  const page = await context.newPage();
  await page.goto(new URL(`/${collection!.slug}/text/textdocument`, siteURL).toString());
  await settleNuxt(page);
  const toc = page.locator("[data-y-reading-toc]");
  await expect(toc).toBeVisible();
  const metrics = await toc.locator("ul").evaluate((list) => {
    const style = getComputedStyle(list);
    const links = [...list.querySelectorAll("a")];
    return {
      overflowX: style.overflowX,
      overflowY: style.overflowY,
      maxHeight: style.maxHeight,
      scrollsVertically: list.scrollHeight > list.clientHeight,
      linksTruncate: links.every((link) => {
        const linkStyle = getComputedStyle(link);
        return linkStyle.overflowX === "hidden" && linkStyle.textOverflow === "ellipsis" && linkStyle.whiteSpace === "nowrap";
      }),
      titlesComplete: links.every((link) => link.title === link.textContent?.trim()),
      excludesFieldHeadings: !links.some((link) => ["描述", "参数", "返回", "示例", "类型"].includes(link.textContent?.trim() || "")),
    };
  });
  expect(metrics).toMatchObject({ overflowX: "hidden", overflowY: "auto", scrollsVertically: true, linksTruncate: true, titlesComplete: true, excludesFieldHeadings: true });
  const editLink = page.getByRole("link", { name: "编辑" });
  await expect(editLink).toBeVisible();
  await expect(editLink).toHaveAttribute("href", new RegExp(`/manage/docs/${collection!.slug}/text/textdocument$`));

  const mobile = await context.newPage();
  await mobile.setViewportSize({ width: 390, height: 844 });
  await mobile.goto(new URL(`/${collection!.slug}/text/textdocument`, siteURL).toString());
  await settleNuxt(mobile);
  await expect(mobile.locator("[data-y-reading-toc]")).toBeHidden();

  const anonymousContext = await browser.newContext({ viewport: { width: 1440, height: 900 } });
  const anonymous = await anonymousContext.newPage();
  await anonymous.goto(new URL(`/${collection!.slug}/text/textdocument`, siteURL).toString());
  await settleNuxt(anonymous);
  await expect(anonymous.getByRole("link", { name: "编辑" })).toHaveCount(0);
  await anonymousContext.close();
  await context.close();
});
