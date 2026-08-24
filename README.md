# :spider_web: HeapSort compiled to WebAssembly

[![Unit Tests](https://github.com/raulpy271/heapfast/actions/workflows/go.yml/badge.svg)](https://github.com/raulpy271/heapfast/actions/workflows/go.yml)
[![Downloads](https://img.shields.io/npm/dt/heapfast.svg)](https://npmjs.com/package/heapfast)
[![Github stars](https://img.shields.io/github/stars/raulpy271/heapfast.svg)](https://github.com/raulpy271/heapfast)
[![npm version](https://img.shields.io/npm/v/heapfast.svg)](https://npmjs.com/package/heapfast)
[![license MIT](https://img.shields.io/npm/l/heapfast.svg)](https://github.com/raulpy271/epqueue/blob/main/LICENSE)

A HeapSort algorithm written in Go Lang and compiled to WebAssembly. The compiled code is available in a NPM package, it's possible to use in any JS environment that supports WASM.

To install and import the package, is simple as any other nodejs library:

```sh
npm install heapfast
```

## :children_crossing: Usage example

### Sorting an array of numbers

The `heapsort` function accept as the first argument an array of numbers and also a typed array.

```js
import heapfast from "heapfast";

// Sorting a array of numbers
const arr = [23, 534.123, 0.233, 88.23, 1.33];
heapfast.heapsort(arr, heapfast.DESC);
console.log(arr);
// [ 534.123, 88.23, 23, 1.33, 0.233 ]

// Sorting a big int array
const bgarr = new BigInt64Array([82n, 234n, 304000n, 234423n, 432n]);
heapfast.heapsort(bgarr, heapfast.ASC);
console.log(bgarr);
// BigInt64Array(5) [ 82n, 234n, 432n, 234423n, 304000n ]
```

### Sorting an array of objects

To sort an array of objects the user needs to provider a sorting key, which could be a string or a function. If it is a string then that string will be the object property used to sort the array. See an example:

```js
import heapfast from "heapfast";

const persons = [
  { name: "Raul", skill: "dumb", age: 33 },
  { name: "Otwell", skill: "PHP", age: 23 },
  { name: "Ken", skill: "Go", age: 13 },
  { name: "Peter", skill: "AI", age: 58 },
  { name: "Linus", skill: "system design", age: 18 },
  { name: "Bob", skill: "clean code", age: 40 },
];

// Sort inplace the array of object by age
heapfast.heapsort(persons, heapfast.ASC, "age");
```

As we said, the third parameter can be a function, so in the last example if we pass the arrow function `(person) => person.age` we would get the same result.

## :envelope: License

Released under the MIT License. See [LICENSE.md](/LICENSE) for details.
