const fs = require("fs");
const input = fs.readFileSync(0).toString().trim();

const elves = input.split("\n\n").map(elf => {
  return elf.split("\n").map(cal => parseInt(cal, 10));
});

const sorted = elves.map(elf => elf.reduce((acc, cur) => acc + cur, 0)).sort().reverse();
console.log(sorted[0]);
console.log(sorted[0] + sorted[1] + sorted[2]);
