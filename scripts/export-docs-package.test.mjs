import assert from "node:assert/strict";
import test from "node:test";
import fs from "node:fs";
import { dirname, resolve } from "node:path";
import { fileURLToPath } from "node:url";
import {
  convertDirectiveContainers,
  discoverDocumentKeys,
  exportAllPackages,
  exportPackage,
} from "./export-docs-package.mjs";

const repoRoot = resolve(dirname(fileURLToPath(import.meta.url)), "..");

test("converts supported directive containers to portable GFM alerts", () => {
  const source = `前文

:::note
第一行

第二行
:::

:::tip 单行提示 :::#### 类型

:::warning[兼容性]
请谨慎升级。
:::
`;
  assert.equal(
    convertDirectiveContainers(source),
    `前文

> [!NOTE]
> 第一行
>
> 第二行

> [!TIP]
> 单行提示

#### 类型

> [!WARNING] 兼容性
> 请谨慎升级。
`,
  );
});

test("leaves directive-looking text inside fenced code untouched", () => {
  const source = "```md\n:::note\nexample\n:::\n```\n";
  assert.equal(convertDirectiveContainers(source), source);
});

test("converts legacy containers whose type follows the opening marker", () => {
  assert.equal(convertDirectiveContainers(":::\ntip legacy content\n:::\n"), "> [!TIP]\n> legacy content\n");
});

test("rejects unknown and unclosed directives instead of losing content", () => {
  assert.throws(() => convertDirectiveContainers(":::custom\ntext\n:::\n"), /unsupported directive custom/);
  assert.throws(() => convertDirectiveContainers(":::note\ntext\n"), /unclosed directive note/);
});

test("exports the bilingual ae-scripting package without changing its sources", () => {
  const result = exportPackage("ae-scripting", repoRoot);
  assert.equal(result.report.navigationPages, 60);
  assert.deepEqual(result.report.locales["zh-CN"], { documents: 60, directives: 308, assets: 2 });
  assert.deepEqual(result.report.locales["en-US"], { documents: 60, directives: 308, assets: 2 });
  assert.equal(fs.statSync(result.zipPath).size > 0, true);
});

test("discovers only bilingual document sets with navigation", () => {
  assert.deepEqual(discoverDocumentKeys(repoRoot), [
    "ae-expression",
    "ae-plugin",
    "ae-scripting",
    "ai-scripting",
    "houdini-vex",
    "javascript-tools",
    "pr-plugin",
    "pr-scripting",
    "sapphire",
  ]);
});

test("exports every discovered document set and writes a batch report", () => {
  const result = exportAllPackages(repoRoot);
  assert.equal(result.succeeded.length, 9);
  assert.deepEqual(result.failed, []);
  assert.equal(fs.existsSync(result.reportPath), true);
  for (const item of result.succeeded) assert.equal(fs.statSync(item.zipPath).size > 0, true);
});
