import assert from "node:assert";

import heapfast from "../heapfast.js";

describe("sort arrays without values", function () {
  it("Should sort BigInt64Array", function () {
    const intarr = new BigInt64Array([20n, 30n, 10n, 40n]);
    const ascresult = [10n, 20n, 30n, 40n];
    const descresult = [40n, 30n, 20n, 10n];
    let len = heapfast.heapsort(intarr, heapfast.ASC);
    assert.equal(len, 4);
    for (let i = 0; i < 4; i++) {
      assert.equal(intarr[i], ascresult[i]);
    }
    len = heapfast.heapsort(intarr, heapfast.DESC);
    assert.equal(len, 4);
    for (let i = 0; i < 4; i++) {
      assert.equal(intarr[i], descresult[i]);
    }
  });
  it("Should sort Int32Array", function () {
    const intarr = new Int32Array([20, 30, 10, 40]);
    const ascresult = [10, 20, 30, 40];
    const descresult = [40, 30, 20, 10];
    let len = heapfast.heapsort(intarr, heapfast.ASC);
    assert.equal(len, 4);
    for (let i = 0; i < 4; i++) {
      assert.equal(intarr[i], ascresult[i]);
    }
    len = heapfast.heapsort(intarr, heapfast.DESC);
    assert.equal(len, 4);
    for (let i = 0; i < 4; i++) {
      assert.equal(intarr[i], descresult[i]);
    }
  });
  it("Should sort Float32Array", function () {
    const floatarr = new Float32Array([20.5, 30.5, 10.5, 40.5]);
    const ascresult = [10.5, 20.5, 30.5, 40.5];
    const descresult = [40.5, 30.5, 20.5, 10.5];
    let len = heapfast.heapsort(floatarr, heapfast.ASC);
    assert.equal(len, 4);
    for (let i = 0; i < 4; i++) {
      assert.equal(floatarr[i], ascresult[i]);
    }
    len = heapfast.heapsort(floatarr, heapfast.DESC);
    assert.equal(len, 4);
    for (let i = 0; i < 4; i++) {
      assert.equal(floatarr[i], descresult[i]);
    }
  });
  it("Should sort Uint32Array", function () {
    const uintarr = new Uint32Array([20, 30, 10, 40]);
    const ascresult = [10, 20, 30, 40];
    const descresult = [40, 30, 20, 10];
    let len = heapfast.heapsort(uintarr, heapfast.ASC);
    assert.equal(len, 4);
    for (let i = 0; i < 4; i++) {
      assert.equal(uintarr[i], ascresult[i]);
    }
    len = heapfast.heapsort(uintarr, heapfast.DESC);
    assert.equal(len, 4);
    for (let i = 0; i < 4; i++) {
      assert.equal(uintarr[i], descresult[i]);
    }
  });
  it("Should sort untyped Array of numbers", function () {
    const arr = [20.5, 30, 10.5, 40];
    const ascresult = [10.5, 20.5, 30, 40];
    const descresult = [40, 30, 20.5, 10.5];
    let len = heapfast.heapsort(arr, heapfast.ASC);
    assert.equal(len, 4);
    for (let i = 0; i < 4; i++) {
      assert.equal(arr[i], ascresult[i]);
    }
    len = heapfast.heapsort(arr, heapfast.DESC);
    assert.equal(len, 4);
    for (let i = 0; i < 4; i++) {
      assert.equal(arr[i], descresult[i]);
    }
  });
  it("Should throw when the untyped Array does not hold numbers", function () {
    const arr = ["b", "a", "c"];
    assert.throws(() => {
      heapfast.heapsort(arr, heapfast.ASC);
    }, TypeError);
  });
});

describe("sort arrays of objects", function () {
  it("Should sort using string key", function () {
    const persons = [
      { name: "Raul", skill: "dumb", age: 33 },
      { name: "Otwell", skill: "PHP", age: 23 },
      { name: "Ken", skill: "Go", age: 13 },
      { name: "Peter", skill: "AI", age: 58 },
      { name: "Linus", skill: "system design", age: 18 },
      { name: "Bob", skill: "clean code", age: 40 },
    ];
    const ascresult = [2, 4, 1, 0, 5, 3].map((i) => persons[i]);
    const descresult = [3, 5, 0, 1, 4, 2].map((i) => persons[i]);
    let len = heapfast.heapsort(persons, heapfast.ASC, "age");
    assert.equal(len, persons.length);
    for (let i = 0; i < len; i++) {
      assert.equal(persons[i], ascresult[i]);
    }
    len = heapfast.heapsort(persons, heapfast.DESC, "age");
    assert.equal(len, persons.length);
    for (let i = 0; i < len; i++) {
      assert.equal(persons[i], descresult[i]);
    }
  });
  it("Should sort using function key", function () {
    const persons = [
      { name: "Raul", skill: "dumb", age: 33 },
      { name: "Otwell", skill: "PHP", age: 23 },
      { name: "Ken", skill: "Go", age: 13 },
      { name: "Peter", skill: "AI", age: 58 },
      { name: "Linus", skill: "system design", age: 18 },
      { name: "Bob", skill: "clean code", age: 40 },
    ];
    const ascresult = [2, 4, 1, 0, 5, 3].map((i) => persons[i]);
    const descresult = [3, 5, 0, 1, 4, 2].map((i) => persons[i]);
    let len = heapfast.heapsort(persons, heapfast.ASC, (p) => p.age);
    assert.equal(len, persons.length);
    for (let i = 0; i < len; i++) {
      assert.equal(persons[i], ascresult[i]);
    }
    len = heapfast.heapsort(persons, heapfast.DESC, (p) => p.age);
    assert.equal(len, persons.length);
    for (let i = 0; i < len; i++) {
      assert.equal(persons[i], descresult[i]);
    }
  });
});
