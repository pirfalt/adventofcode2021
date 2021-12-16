import fs from "fs";

const ex = fs.readFileSync("./example.txt");
const inp = fs.readFileSync("./input.txt");

console.log(one(ex));
console.log(one(inp));

console.log(two(ex));
console.log(two(inp));

function one(data) {
  return data
    .toString()
    .split("\n")
    .map((line) => line.split(""))
    .map((l) => {
      const matches = {
        "(": ")",
        "[": "]",
        "{": "}",
        "<": ">",
      };
      const points = { ")": 3, "]": 57, "}": 1197, ">": 25137 };

      let stack = [];
      for (const c of l) {
        if (Object.keys(matches).includes(c)) {
          stack.push(c);
        }
        if (Object.values(matches).includes(c)) {
          const open = stack.pop();
          const expected = matches[open];
          const actual = c;
          if (expected != actual) {
            return points[actual];
          }
        }
      }

      return 0;
    })
    .reduce((a, b) => a + b, 0);
}

function two(data) {
  const input = data
    .toString()
    .split("\n")
    .map((line) => line.split(""));

  const scores = input.flatMap((l) => {
    const matches = {
      "(": ")",
      "[": "]",
      "{": "}",
      "<": ">",
    };
    const points = {
      ")": 1,
      "]": 2,
      "}": 3,
      ">": 4,
    };

    let stack = [];
    for (const c of l) {
      if (Object.keys(matches).includes(c)) {
        stack.push(c);
      }
      if (Object.values(matches).includes(c)) {
        const open = stack.pop();
        const expected = matches[open];
        const actual = c;
        if (expected != actual) {
          return [];
        }
      }
    }

    const close = stack.map((open) => points[matches[open]]).reverse();
    const score = close.reduce((acc, score) => acc * 5 + score, 0);

    return [score];
  });
  const mid = Math.floor(scores.length / 2);
  return scores.sort((a, b) => a - b)[mid];
}
