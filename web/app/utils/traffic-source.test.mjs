import assert from "node:assert/strict";
import test from "node:test";

import { trafficSource } from "./traffic-source.mjs";

test("document traffic sources keep only an external host", () => {
  assert.equal(trafficSource("", "https://docs.example/guide/a"), "direct");
  assert.equal(
    trafficSource("https://docs.example/search?q=private", "https://docs.example/guide/a"),
    "internal",
  );
  assert.equal(
    trafficSource("https://www.Google.COM/search?q=private", "https://docs.example/guide/a"),
    "google.com",
  );
});
