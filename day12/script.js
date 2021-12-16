import fs from "fs";

const ex1 = fs.readFileSync("./ex1.txt");
const ex2 = fs.readFileSync("./ex2.txt");
const ex3 = fs.readFileSync("./ex3.txt");
const inp = fs.readFileSync("./input.txt");

console.log("--- Part 1 ---------------------------");
// one(ex1);
one(inp);

console.log("--- Part 2 ---------------------------");
// two(ex1);
two(inp);

function one(data) {
  const _rules = parseInput(data);
  const rules = _rules.concat(_rules.map(([from, to]) => [to, from]));
  console.log(rules);

  function step(paths, pos) {
    if (pos == "end") return [{ end: [...paths, pos] }];
    if (pos == pos.toLowerCase() && paths.includes(pos))
      return [{ stuck: [...paths, pos] }];

    const opts = rules.filter(([from, to]) => from == pos);
    // console.log({ opts });
    const moves = opts.flatMap(([_, to]) => step([...paths, pos], to));
    // console.log({ moves });
    return moves;
  }

  const paths = step([], "start");
  // console.log(paths);
  console.log(paths.filter((o) => "end" in o).length);
}

function two(data) {
  const _rules = parseInput(data);
  const rules = _rules.concat(_rules.map(([from, to]) => [to, from]));
  console.log({ rules });

  const smallSet = new Set(
    rules.map(([from]) => from).filter(([from]) => from == from.toLowerCase())
  );
  smallSet.delete("start");
  smallSet.delete("end");
  const small = [...smallSet];
  console.log({ small });

  function step(paths, pos, duplicateSmall) {
    if (pos == "end") {
      return [{ end: [...paths, pos] }];
    }
    if (
      pos == pos.toLowerCase() &&
      pos != duplicateSmall &&
      paths.includes(pos)
    ) {
      return [{ stuck1: [...paths, pos] }];
    }
    if (pos == duplicateSmall && paths.filter((p) => p == pos).length >= 2) {
      return [{ stuck2: [...paths, pos] }];
    }

    const opts = rules.filter(([from, to]) => from == pos);
    // console.log({ opts });
    const moves = opts.flatMap(([_, to]) =>
      step([...paths, pos], to, duplicateSmall)
    );
    // console.log({ moves });
    return moves;
  }

  const paths = small.flatMap((s) => {
    return step([], "start", s);
  });

  // console.log(paths);
  const validPaths = new Set(
    paths.filter((o) => "end" in o).map(({ end }) => end.join(","))
  );
  console.log({
    l: validPaths.size,
  });
}

function parseInput(data) {
  return data
    .toString()
    .trim()
    .split("\n")
    .map((l) => l.split("-"));
}

function range(len) {
  return [...Array(len)].map((_, i) => i);
}
