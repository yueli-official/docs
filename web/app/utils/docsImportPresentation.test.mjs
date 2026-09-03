import test from "node:test";
import assert from "node:assert/strict";
import {
  formatImportDate,
  importModeLabel,
  importStatusBadgeUI,
  importStatusMeta,
} from "./docsImportPresentation.mjs";

test("import presentation translates stable states and modes", () => {
  assert.equal(importStatusMeta("checked").label, "已预检");
  assert.equal(importStatusMeta("completed").color, "success");
  assert.equal(importStatusMeta("rolled_back").label, "已回滚");
  assert.equal(importModeLabel("create-only"), "仅创建");
  assert.equal(importModeLabel("upsert"), "创建或更新");
  assert.equal(importModeLabel("replace-version"), "替换版本");
});

test("import dates fail closed to a readable placeholder", () => {
  assert.equal(formatImportDate(""), "未记录");
  assert.equal(formatImportDate("not-a-date"), "未记录");
  assert.notEqual(formatImportDate("2026-07-25T08:30:00Z"), "未记录");
});

test("running import badges spin unless reduced motion is requested", () => {
  assert.deepEqual(importStatusBadgeUI("running"), {
    leadingIcon: "animate-spin motion-reduce:animate-none",
  });
  assert.equal(importStatusBadgeUI("completed"), undefined);
});
