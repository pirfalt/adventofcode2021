import fs from "fs";

const ex = fs.readFileSync("./example.txt");
const inp = fs.readFileSync("./input.txt");

one(ex);
// one(inp);

console.log("------------------------------");

two(ex, 40);
// two(inp);

function one(data, size = 10) {
  const [template, pairsRaw] = data.toString().trim().split("\n\n");
  const rules = pairsRaw
    .split("\n")
    .map((l) => /(\w\w) -> (\w)/.exec(l).slice(1))
    .reduce((o, [match, target]) => {
      o[match] = target;
      return o;
    }, {});

  const output = range(size).reduce((template) => {
    const tpairs = range(template.length - 1)
      .map((i) => template.substr(i, 2))
      .map((p) => p[0] + rules[p]);
    return tpairs.join("") + template.at(-1);
  }, template);

  const chars = {};
  for (const c of output) {
    chars[c] = chars[c] ?? 0;
    chars[c] += 1;
  }

  const max = Math.max(...Object.values(chars));
  const min = Math.min(...Object.values(chars));

  console.log({ template, output, rules, chars, max, min, result: max - min });
}

function two(data, size = 10) {
  const [template, pairsRaw] = data.toString().trim().split("\n\n");
  const rules = pairsRaw
    .split("\n")
    .map((l) => /(\w\w) -> (\w)/.exec(l).slice(1))
    .reduce((o, [match, target]) => {
      o[match] = target;
      return o;
    }, {});

  let pairs = range(template.length - 1)
    .map((i) => template.substr(i, 2))
    .reduce((o, p) => {
      o[p] = (o[p] ?? 0) + 1;
      return o;
    }, {});

  let chars = template.split("").reduce((o, p) => {
    o[p] = (o[p] ?? 0) + 1;
    return o;
  }, {});

  range(size).forEach(() => {
    const inp = [...Object.entries(pairs)].filter(([_, v]) => v > 0);

    inp.forEach(([p, count]) => {
      const char = rules[p];
      const n1 = p[0] + char;
      const n2 = char + p[1];

      pairs[p] = pairs[p] - count;
      pairs[n1] = (pairs[n1] ?? 0) + count;
      pairs[n2] = (pairs[n2] ?? 0) + count;

      chars[char] = (chars[char] ?? 0) + count;
    });
  });

  const max = Math.max(...Object.values(chars));
  const min = Math.min(...Object.values(chars));

  console.log({ rules, pairs, chars, max, min, result: max - min });
}

function range(len) {
  return [...Array(len)].map((_, i) => i);
}
