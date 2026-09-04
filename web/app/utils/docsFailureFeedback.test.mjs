import assert from "node:assert/strict";
import { readFile } from "node:fs/promises";
import test from "node:test";

test("Docs feedback resolves stable codes without raw exception messages", async () => {
  const source = await readFile(new URL("./docsFailureFeedback.ts", import.meta.url), "utf8");
  assert.match(source, /resolveFailureFeedback\(error/u);
  assert.match(source, /"docs\.import\.compression_unsupported"/u);
  assert.doesNotMatch(source, /error\.message|data\?\.message/u);
});
