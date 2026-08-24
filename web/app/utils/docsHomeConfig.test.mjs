import test from "node:test";
import assert from "node:assert/strict";
import { normalizeFeaturedCollections } from "./docsHomeConfig.mjs";

test("legacy null featured collections become an empty settings selection", () => {
  assert.deepEqual(normalizeFeaturedCollections(null), []);
});

test("featured collection order is preserved", () => {
  assert.deepEqual(normalizeFeaturedCollections(["guides", "reference"]), [
    "guides",
    "reference",
  ]);
});
