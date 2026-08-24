import assert from "node:assert/strict";
import test from "node:test";

import { createTrafficReplayKey } from "./traffic-replay-key.mjs";

test("document traffic replay keys prefer browser randomUUID", () => {
  const expected = "019c0000-0000-4000-8000-000000000001";
  assert.equal(createTrafficReplayKey({ randomUUID: () => expected }), expected);
});

test("document traffic replay keys remain UUID-shaped on local HTTP", () => {
  const cryptoSource = {
    getRandomValues(target) {
      target.fill(0);
      return target;
    },
  };
  assert.equal(
    createTrafficReplayKey(cryptoSource),
    "00000000-0000-4000-8000-000000000000",
  );
});
