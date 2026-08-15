import assert from "node:assert";

import heapfast from "../heapfast.js";

describe("HeapsortExported", function () {
  it("Should sort integer array", function () {
    const intarr = new BigInt64Array([20n, 30n, 10n, 40n]);
    const ascresult = [10n, 20n, 30n, 40n];
    const descresult = [40n, 30n, 20n, 10n];
    let len = HeapsortAscIntK(new Uint8Array(intarr.buffer));
    assert.equal(len, 4);
    for (let i = 0; i < 4; i++) {
      assert.equal(intarr[i], ascresult[i]);
    }
    len = HeapsortDescIntK(new Uint8Array(intarr.buffer));
    assert.equal(len, 4);
    for (let i = 0; i < 4; i++) {
      assert.equal(intarr[i], descresult[i]);
    }
  });
  it("Should sort float array", function () {
    const intarr = new Float64Array([20.1, 30.1, 10.1, 40.1]);
    const ascresult = [10.1, 20.1, 30.1, 40.1];
    const descresult = [40.1, 30.1, 20.1, 10.1];
    let len = HeapsortAscFloatK(new Uint8Array(intarr.buffer));
    assert.equal(len, 4);
    for (let i = 0; i < 4; i++) {
      assert.equal(intarr[i], ascresult[i]);
    }
    len = HeapsortDescFloatK(new Uint8Array(intarr.buffer));
    assert.equal(len, 4);
    for (let i = 0; i < 4; i++) {
      assert.equal(intarr[i], descresult[i]);
    }
  });
  it("Should sort integer map", function () {
    const intarr = new BigInt64Array([2n, 20n, 3n, 30n, 1n, 10n, 4n, 40n]);
    const ascresult = [
      { k: 10n, v: 1n },
      { k: 20n, v: 2n },
      { k: 30n, v: 3n },
      { k: 40n, v: 4n },
    ];
    const descresult = [
      { k: 40n, v: 4n },
      { k: 30n, v: 3n },
      { k: 20n, v: 2n },
      { k: 10n, v: 1n },
    ];
    let len = HeapsortAscIntKV(new Uint8Array(intarr.buffer));
    assert.equal(len, 4);
    for (let i = 0; i < 4; i++) {
      assert.equal(intarr[i * 2 + 1], ascresult[i].k);
      assert.equal(intarr[i * 2], ascresult[i].v);
    }
    len = HeapsortDescIntKV(new Uint8Array(intarr.buffer));
    assert.equal(len, 4);
    for (let i = 0; i < 4; i++) {
      assert.equal(intarr[i * 2 + 1], descresult[i].k);
      assert.equal(intarr[i * 2], descresult[i].v);
    }
  });
  it("Should sort float map", function () {
    const intarr = new Float64Array([
      2.0, 20.1, 3.0, 30.1, 1.0, 10.1, 4.0, 40.1,
    ]);
    const ascresult = [
      { k: 10.1, v: 1.0 },
      { k: 20.1, v: 2.0 },
      { k: 30.1, v: 3.0 },
      { k: 40.1, v: 4.0 },
    ];
    const descresult = [
      { k: 40.1, v: 4.0 },
      { k: 30.1, v: 3.0 },
      { k: 20.1, v: 2.0 },
      { k: 10.1, v: 1.0 },
    ];
    let len = HeapsortAscIntKV(new Uint8Array(intarr.buffer));
    assert.equal(len, 4);
    for (let i = 0; i < 4; i++) {
      assert.equal(intarr[i * 2 + 1], ascresult[i].k);
      assert.equal(intarr[i * 2], ascresult[i].v);
    }
    len = HeapsortDescIntKV(new Uint8Array(intarr.buffer));
    assert.equal(len, 4);
    for (let i = 0; i < 4; i++) {
      assert.equal(intarr[i * 2 + 1], descresult[i].k);
      assert.equal(intarr[i * 2], descresult[i].v);
    }
  });
});
