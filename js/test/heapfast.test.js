import assert from "node:assert";

import { startWasmModule } from "../heapfast.js";

before(function () {
  const go = new Go();
  startWasmModule(go);
});

describe("HeapsortExported", function () {
  it("Should sort integer array", function () {
    const intarr = new BigInt64Array([20n, 0n, 30n, 0n, 10n, 0n, 40n, 0n]);
    const ascresult = [10n, 20n, 30n, 40n];
    const descresult = [40n, 30n, 20n, 10n];
    let len = HeapsortAscInt(new Uint8Array(intarr.buffer));
    assert.equal(len, 4);
    for (let i = 0; i < 4; i++) {
      assert.equal(intarr[i * 2], ascresult[i]);
    }
    len = HeapsortDescInt(new Uint8Array(intarr.buffer));
    assert.equal(len, 4);
    for (let i = 0; i < 4; i++) {
      assert.equal(intarr[i * 2], descresult[i]);
    }
  });
  it("Should sort float array", function () {
    const intarr = new Float64Array([20.1, 0, 30.1, 0, 10.1, 0, 40.1, 0]);
    const ascresult = [10.1, 20.1, 30.1, 40.1];
    const descresult = [40.1, 30.1, 20.1, 10.1];
    let len = HeapsortAscFloat(new Uint8Array(intarr.buffer));
    assert.equal(len, 4);
    for (let i = 0; i < 4; i++) {
      assert.equal(intarr[i * 2], ascresult[i]);
    }
    len = HeapsortDescInt(new Uint8Array(intarr.buffer));
    assert.equal(len, 4);
    for (let i = 0; i < 4; i++) {
      assert.equal(intarr[i * 2], descresult[i]);
    }
  });
});

describe("Array", function () {
  describe("#indexOf()", function () {
    it("should return -1 when the value is not present", function () {
      assert.equal([1, 2, 3].indexOf(4), -1);
    });
  });
});
