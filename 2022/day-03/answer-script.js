const fs = require("fs");
const input = fs.readFileSync(0).toString().trim();

function findPriority(item) {
  let priority = item.charCodeAt(0);
  if (priority >= 97) {
    priority -= 96;
  } else {
    priority -= 38;
  }
  return priority;
}

const sums = input.split("\n").reduce((acc, line, index, array) => {
  {
    const compartment1 = new Set(line.substr(0, line.length / 2));
    const compartment2 = new Set(line.substr(line.length / 2));

    for (const item of compartment1) {
      if (compartment2.has(item)) {
        acc[0] += findPriority(item);
        break;
      }
    }
  }

  if (index % 3 === 0) {
    const compartment1 = new Set(line);
    const compartment2 = new Set(array[index + 1]);
    const compartment3 = new Set(array[index + 2]);

    for (const item of compartment1) {
      if (compartment2.has(item) && compartment3.has(item)) {
        acc[1] += findPriority(item);
        break;
      }
    }
  }
  return acc;
}, [0, 0]);

console.log(sums[0]);
console.log(sums[1]);
