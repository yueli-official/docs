import { expect, test } from "@playwright/test";
import AxeBuilder from "@axe-core/playwright";
import { productSites } from "./contracts";
import {
  capturePageFailures,
  expectNoHorizontalOverflow,
  loginE2E,
  requiredEnv,
  settleNuxt,
} from "./runtime";

const accountURL = requiredEnv("DOCS_E2E_ACCOUNT_URL");
const missing = "__docs_e2e_missing__";

export function registerJourneySuite(product: string) {
  for (const site of productSites(product)) {
    const contract = site.contract;
    test.describe(`${site.slug} (${site.product})`, () => {
      test("公开入口完成渲染且没有浏览器错误", async ({ page }) => {
        const errors = capturePageFailures(page);
        const response = await page.goto(
          new URL(contract.public.path, site.url).toString(),
          { waitUntil: "domcontentloaded" },
        );
        expect(response?.ok()).toBeTruthy();
        await expect(
          page.locator(contract.public.readySelector).first(),
        ).toBeVisible();
        await settleNuxt(page);
        expect(errors).toEqual([]);
      });

      test("文档阅读页使用共享评论线程与服务端排序", async ({
        page,
      }, testInfo) => {
        const requestedOrders: string[] = [];
        await page.route("**/api/v1/docs/*/comments**", async (route) => {
          const requestURL = new URL(route.request().url());
          const order = requestURL.searchParams.get("sortOrder") || "asc";
          requestedOrders.push(order);
          const older = {
            id: "comment-old",
            authorName: "测试读者",
            content: "这篇文档把步骤说清楚了。",
            createdAt: "2026-08-28T08:00:00Z",
            replies: [
              {
                id: "comment-reply",
                parentId: "comment-old",
                authorName: "维护者",
                content: "谢谢反馈。",
                createdAt: "2026-08-28T08:05:00Z",
              },
            ],
          };
          const newer = {
            id: "comment-new",
            authorName: "另一位读者",
            content: "排序切换也保持当前线程结构。",
            createdAt: "2026-08-28T09:00:00Z",
          };
          await route.fulfill({
            status: 200,
            contentType: "application/json",
            body: JSON.stringify({
              total: 2,
              page: 1,
              size: 100,
              items: order === "desc" ? [newer, older] : [older, newer],
            }),
          });
        });

        await page.setViewportSize({ width: 1440, height: 960 });
        await page.goto(
          new URL("/getting-started/overview?locale=en", site.url).toString(),
          { waitUntil: "domcontentloaded" },
        );
        await settleNuxt(page);
        const comments = page.locator("[data-document-comments]");
        await comments.scrollIntoViewIfNeeded();
        await expect(
          comments.getByRole("heading", { name: "2 条评论", exact: true }),
        ).toBeVisible();
        await expect
          .poll(async () => {
            const targets = [
              comments,
              page
                .getByRole("link", { name: /下一篇/u })
                .last()
                .locator(".."),
            ];
            return Promise.all(
              targets.map((target) =>
                target.evaluate((element) => {
                  const style = getComputedStyle(element);
                  return [
                    style.borderTopWidth,
                    style.borderRightWidth,
                    style.borderBottomWidth,
                    style.borderLeftWidth,
                  ];
                }),
              ),
            );
          })
          .toEqual([
            ["0px", "0px", "0px", "0px"],
            ["0px", "0px", "0px", "0px"],
          ]);
        await expect
          .poll(async () => {
            const targets = [
              comments.locator("[data-public-comment]").first(),
              comments.locator("[data-public-comment-composer]").last(),
              page.getByRole("link", { name: /下一篇/u }).last(),
            ];
            return Promise.all(
              targets.map((target) =>
                target.evaluate((element) => {
                  const style = getComputedStyle(element);
                  return [
                    style.borderTopWidth,
                    style.borderRightWidth,
                    style.borderBottomWidth,
                    style.borderLeftWidth,
                  ];
                }),
              ),
            );
          })
          .toEqual([
            ["1px", "1px", "1px", "1px"],
            ["1px", "1px", "1px", "1px"],
            ["1px", "1px", "1px", "1px"],
          ]);
        await expect
          .poll(async () => {
            const [navigation, discussion] = await Promise.all([
              page
                .getByRole("link", { name: /下一篇/u })
                .last()
                .boundingBox(),
              comments.boundingBox(),
            ]);
            return Boolean(
              navigation &&
              discussion &&
              navigation.y + navigation.height < discussion.y,
            );
          })
          .toBe(true);
        await expect(
          comments.locator("[data-comment-id]").first(),
        ).toHaveAttribute("data-comment-id", "comment-old");
        await expect(
          comments.getByText("谢谢反馈。", { exact: true }),
        ).toBeVisible();
        await expect(
          comments.getByRole("button", { name: "登录后评论", exact: true }),
        ).toBeVisible();

        await comments
          .getByRole("button", { name: "最新", exact: true })
          .click();
        await expect.poll(() => requestedOrders.at(-1)).toBe("desc");
        await expect(
          comments.locator("[data-comment-id]").first(),
        ).toHaveAttribute("data-comment-id", "comment-new");
        await comments
          .getByRole("button", { name: "回复", exact: true })
          .first()
          .click();
        await expect(
          comments.getByRole("textbox", { name: "写下回复…", exact: true }),
        ).toBeFocused();
        await page.screenshot({
          path: testInfo.outputPath("docs-comments-desktop.png"),
          fullPage: true,
        });

        const accessibility = await new AxeBuilder({ page })
          .exclude("nuxt-devtools-frame")
          .analyze();
        expect(
          accessibility.violations.filter((violation) =>
            ["serious", "critical"].includes(violation.impact || ""),
          ),
        ).toEqual([]);

        await page.setViewportSize({ width: 390, height: 844 });
        await expectNoHorizontalOverflow(page, "Docs 评论移动端");
        await page.screenshot({
          path: testInfo.outputPath("docs-comments-mobile.png"),
          fullPage: true,
        });
      });

      test("匿名访问管理入口进入账户登录流程", async ({ page }) => {
        const errors = capturePageFailures(page);
        await page.goto(new URL(contract.manage.path, site.url).toString(), {
          waitUntil: "domcontentloaded",
        });
        await expect(page).toHaveURL(
          new RegExp(
            `^${accountURL.replace(/[.*+?^${}()|[\]\\]/g, "\\$&")}/login(?:\\?|$)`,
          ),
        );
        await expect(
          page.getByRole("heading", { name: "欢迎回来" }),
        ).toBeVisible();
        await settleNuxt(page);
        expect(errors).toEqual([]);
      });

      test("已登录运营者可以进入管理界面", async ({ browser }) => {
        const context = await loginE2E(browser);
        const page = await context.newPage();
        const errors = capturePageFailures(page);
        try {
          const manageURL = new URL(contract.manage.path, site.url).toString();
          await page.goto(manageURL, { waitUntil: "domcontentloaded" });
          await expect(page).toHaveURL(manageURL);
          await expect(
            page.locator(contract.manage.readySelector).first(),
          ).toBeVisible();
          await expect(
            page
              .getByRole("heading", {
                name: new RegExp(contract.manage.heading),
              })
              .first(),
          ).toBeVisible();
          await settleNuxt(page);
          expect(errors).toEqual([]);
        } finally {
          await context.close();
        }
      });

      test("设置页脏数据保护阻止意外离开", async ({ browser }) => {
        const settings = contract.manage.settings;
        test.skip(!settings, "该产品没有可编辑的通用设置页");
        if (!settings) return;
        const context = await loginE2E(browser);
        const page = await context.newPage();
        const errors = capturePageFailures(page);
        try {
          await page.goto(new URL(settings.path, site.url).toString(), {
            waitUntil: "domcontentloaded",
          });
          await settleNuxt(page);
          const field = page
            .getByRole("textbox", { name: settings.fieldLabel, exact: false })
            .first();
          await expect(field).toBeVisible();
          await expect(page.locator('[data-manage-dock="save"]')).toHaveCount(
            0,
          );
          await field.fill(`${await field.inputValue()} · 未保存`);
          await expect(page.locator('[data-manage-dock="save"]')).toBeVisible();

          const dialogHandled = new Promise<void>((resolve) => {
            page.once("dialog", async (dialog) => {
              expect(dialog.type()).toBe("confirm");
              expect(dialog.message()).toContain("未保存");
              await dialog.dismiss();
              resolve();
            });
          });
          await page
            .locator(`a[href="${contract.manage.path}"]`)
            .first()
            .click();
          await dialogHandled;
          await expect(page).toHaveURL(
            new RegExp(
              `${settings.path.replace(/[.*+?^${}()|[\]\\]/g, "\\$&")}(?:\\?|$)`,
            ),
          );
          expect(errors).toEqual([]);
        } finally {
          await context.close();
        }
      });

      test("空结果状态有明确反馈", async ({ browser, page }) => {
        const context = contract.empty.authenticated
          ? await loginE2E(browser)
          : undefined;
        const targetPage = context ? await context.newPage() : page;
        const errors = capturePageFailures(targetPage);
        try {
          await targetPage.goto(
            new URL(contract.empty.path, site.url).toString(),
            { waitUntil: "domcontentloaded" },
          );
          const emptyState = targetPage
            .getByText(contract.empty.text, { exact: true })
            .first();
          if (contract.empty.inputPlaceholder) {
            const input = targetPage.getByPlaceholder(
              contract.empty.inputPlaceholder,
            );
            await expect(async () => {
              await input.fill(missing);
              await expect(input).toHaveValue(missing);
              await expect(emptyState).toBeVisible({ timeout: 2_000 });
            }).toPass({ timeout: 15_000, intervals: [250, 500, 1_000] });
          }
          await expect(emptyState).toBeVisible();
          await settleNuxt(targetPage);
          expect(errors).toEqual([]);
        } finally {
          await context?.close();
        }
      });

      test("缺失实体返回产品错误状态", async ({ page }) => {
        const errors = capturePageFailures(page);
        const errorURL = new URL(contract.error.path, site.url).toString();
        const response = await page.goto(errorURL, {
          waitUntil: "domcontentloaded",
        });
        expect(response?.status()).toBe(contract.error.status);
        await expect(
          page.getByText(contract.error.text, { exact: false }).first(),
        ).toBeVisible();
        await settleNuxt(page);
        expect(
          errors.filter(
            (error) =>
              error !== `http ${contract.error.status}: ${errorURL}` &&
              !error.includes(`status of ${contract.error.status}`),
          ),
        ).toEqual([]);
      });
    });
  }
}
