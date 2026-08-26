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

test("文档集编辑器使用紧凑底栏与方形封面合同", async ({ browser }) => {
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
    await page.getByRole("button", { name: /^编辑文档集：/u }).first().click();

    const editor = page.getByRole("dialog", { name: "编辑文档集" });
    await expect(editor).toBeVisible();
    await expect(editor.getByText("路径预览", { exact: true })).toHaveCount(0);
    await expect(editor.getByText("危险操作", { exact: true })).toHaveCount(0);
    await expect(editor.getByPlaceholder("搜索图标")).toHaveCount(0);
    await editor.getByRole("button", { name: "查看封面规格", exact: true }).hover();
    await expect(
      page.getByText("1:1 · 256 × 256 · WebP", { exact: true }).last(),
    ).toBeVisible();

    const coverPreview = editor.locator("[data-cover-preview]");
    const coverActions = editor.locator("[data-cover-actions]");
    const coverLink = editor.locator("[data-cover-link]");
    const iconColumn = editor.locator("[data-icon-column]");
    const [previewBox, actionsBox, linkBox, iconColumnBox] = await Promise.all([
      coverPreview.boundingBox(),
      coverActions.boundingBox(),
      coverLink.boundingBox(),
      iconColumn.boundingBox(),
    ]);
    expect(previewBox).not.toBeNull();
    expect(actionsBox).not.toBeNull();
    expect(linkBox).not.toBeNull();
    expect(iconColumnBox).not.toBeNull();
    expect(actionsBox!.y).toBeGreaterThan(previewBox!.y + previewBox!.height);
    expect(linkBox!.y).toBeGreaterThan(actionsBox!.y + actionsBox!.height);
    expect(iconColumnBox!.x).toBeGreaterThan(previewBox!.x + previewBox!.width);

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
