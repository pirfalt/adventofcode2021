import fs from "fs";

const ex = fs.readFileSync("./example.txt");
const inp = fs.readFileSync("./input.txt");

console.log(one(ex));
console.log(one(inp));

console.log(printable(two(ex)));
console.log(printable(two(inp)));

function parseInput(data) {
  const [dotLines, foldLines] = data.toString().trim().split("\n\n");

  const dots = dotLines
    .split("\n")
    .map((l) => l.split(",").map((v) => Number.parseInt(v, 10)));

  const folds = foldLines.split("\n").map((l) => {
    const fold = l.split(" ")[2];
    const [ax, v] = fold.split("=");
    return { ax, v: Number.parseInt(v, 10) };
  });

  return { dots, folds };
}

function one(data) {
  const { dots, folds } = parseInput(data);

  const f = folds[0];
  console.log(f);

  let folded;
  if (f.ax == "y") {
    // fold up
    folded = dots.map(([x, y]) => {
      if (y > f.v) {
        return [x, f.v - (y - f.v)];
      }
      return [x, y];
    });
  }
  if (f.ax == "x") {
    // fold left
    folded = dots.map(([x, y]) => {
      if (x > f.v) {
        return [f.v - (x - f.v), y];
      }
      return [x, y];
    });
  }

  const d = new Set(dots.map((d) => JSON.stringify(d))).size;
  const s = new Set(folded.map((d) => JSON.stringify(d))).size;

  console.log({ d, s });

  return { dots, f };
}

function two(data) {
  const { dots, folds } = parseInput(data);

  const r = folds.reduce((dots, f) => {
    let folded;
    if (f.ax == "y") {
      // fold up
      folded = dots.map(([x, y]) => {
        if (y > f.v) {
          return [x, f.v - (y - f.v)];
        }
        return [x, y];
      });
    }

    if (f.ax == "x") {
      // fold left
      folded = dots.map(([x, y]) => {
        if (x > f.v) {
          return [f.v - (x - f.v), y];
        }
        return [x, y];
      });
    }

    const s = new Set(folded.map((d) => JSON.stringify(d)));
    const _dots = [...s.values()].map((s) => JSON.parse(s));

    return _dots;
  }, dots);

  return r;
}

function printable(dots) {
  const max = dots.reduce(
    ({ mx, my }, [x, y]) => ({
      mx: Math.max(mx, x),
      my: Math.max(my, y),
    }),
    { mx: 0, my: 0 }
  );

  const matrix = [...Array(max.my + 1)].map(() =>
    [...Array(max.mx + 1)].map(() => false)
  );

  for (const [x, y] of dots) {
    matrix[y][x] = true;
  }

  let out = "=====================\n";
  out += matrix
    .map((row) => {
      return row.map((c) => (c ? "#" : ".")).join("");
    })
    .join("\n");
  out += "\n---------------------";

  return out;
}
