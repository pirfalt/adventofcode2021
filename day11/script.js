import fs from "fs";

const ex = fs.readFileSync("./example.txt");
const inp = fs.readFileSync("./input.txt");

console.log("part1");
run(ex, 100);
run(inp, 100);

console.log("part2");
run(ex, 200, true);
run(inp, 1000, true);

function run(data, steps = 10, breakOnFull = false) {
  const r = parseInput(data);
  let sum = 0;
  let base = r;
  let flashing = new Set();
  let visited = new Set();

  for (let i = 0; i < steps; i++) {
    base = base.map((v) => v + 1);
    flashing = new Set(base.flatMap((v, i) => (v >= 10 ? [i] : [])));
    visited = new Set();

    while (true) {
      for (let i of flashing) {
        if (visited.has(i)) continue;
        visited.add(i);

        for (let ai of around(i)) {
          base[ai] += 1;
          if (base[ai] >= 10) {
            flashing.add(ai);
          }
        }
      }

      if (flashing.size == visited.size) {
        break;
      }
    }

    for (let i of flashing) {
      sum += 1;
      base[i] = 0;
    }

    if (breakOnFull && flashing.size == 100) {
      console.log(`full at ${i + 1}`);
      break;
    }
  }

  console.log(printable(base), sum);
}

function two(data) {}

function parseInput(data) {
  return data
    .toString()
    .split("\n")
    .flatMap((line) => line.split("").map((c) => Number.parseInt(c, 10)));
}

function around(i) {
  const rL = 10;
  const [iul, iu, iur] = [i - rL - 1, i - rL, i - rL + 1];
  const [il, ir] = [i + 0 - 1, i + 0 + 1];
  const [idl, id, idr] = [i + rL - 1, i + rL, i + rL + 1];
  const ii = new Set([iul, iu, iur, il, ir, idl, id, idr]);

  if (i % rL == 0) [iul, il, idl].forEach((i) => ii.delete(i)); // l
  if (i % rL == 9) [iur, ir, idr].forEach((i) => ii.delete(i)); // r
  if (i <= 9) [iul, iu, iur].forEach((i) => ii.delete(i)); // u
  if (i >= 90) [idl, id, idr].forEach((i) => ii.delete(i)); // d

  return ii;
}

function printable(data) {
  return (
    "-------------------\n" +
    [
      data.slice(0, 10).join(" "),
      data.slice(10, 20).join(" "),
      data.slice(20, 30).join(" "),
      data.slice(30, 40).join(" "),
      data.slice(40, 50).join(" "),
      data.slice(50, 60).join(" "),
      data.slice(60, 70).join(" "),
      data.slice(70, 80).join(" "),
      data.slice(80, 90).join(" "),
      data.slice(90, 100).join(" "),
    ].join("\n") +
    "\n-------------------"
  );
}
