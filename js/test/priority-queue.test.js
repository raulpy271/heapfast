import assert from "node:assert";

import heapfast from "../dist/main.js";

describe("Priority Queue with values", function () {
  it("Should add three items and pop highest", function () {
    const queue = new heapfast.PriorityQueue(heapfast.DESC);
    queue.add(33, "Raul");
    queue.add(20, "Max");
    queue.add(25, "Tyler");
    const [key, p] = queue.pop();
    assert.equal(key, 33);
    assert.equal(p, "Raul");
    assert.deepStrictEqual(queue.values, [undefined, "Max", "Tyler"]);
    assert.equal(queue.freeHead, 0);
    assert.equal(queue.valuesFree, 1);
  });
  it("Should add five items and pop all ascending", function () {
    const queue = new heapfast.PriorityQueue(heapfast.ASC);
    queue.add(33, "Raul");
    queue.add(20, "Max");
    queue.add(25, "Jack");
    queue.add(26, "Bob");
    queue.add(19, "Tyler");
    assert.deepStrictEqual(queue.pop(), [19, "Tyler"]);
    assert.deepStrictEqual(queue.pop(), [20, "Max"]);
    assert.deepStrictEqual(queue.pop(), [25, "Jack"]);
    assert.deepStrictEqual(queue.pop(), [26, "Bob"]);
    assert.deepStrictEqual(queue.pop(), [33, "Raul"]);
    assert.deepStrictEqual(queue.values, [3, 4, 1, 2, undefined]);
    assert.equal(queue.freeHead, 0);
    assert.equal(queue.valuesFree, 5);
  });
});
