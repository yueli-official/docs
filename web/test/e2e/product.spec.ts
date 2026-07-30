import { expect, test } from "@playwright/test";
import { productSites } from "../../../../../tests/e2e/contracts";
import { registerProductSuite } from "../../../../../tests/e2e/product-suite";
import { loginE2E, settleNuxt } from "../../../../../tests/e2e/runtime";

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
