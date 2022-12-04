const fs = require("fs");

const input = fs.readFileSync(0).toString().trim();
const answers = input.split("\n").reduce((acc, line) => {
  const pair = line.split(",").map(range => range.split("-").map(num => parseInt(num, 10)));

  if ((pair[0][0] <= pair[1][0] && pair[0][1] >= pair[1][1]) || (pair[1][0] <= pair[0][0] && pair[1][1] >= pair[0][1])) {
    acc[0] += 1;
    acc[1] += 1;
  } else if ((pair[0][1] >= pair[1][0] && pair[0][0] <= pair[1][0]) || (pair[1][1] >= pair[0][0] && pair[1][0] <= pair[0][0])) {
    acc[1] += 1;
  }

  return acc;
}, [0, 0]);

console.log(answers[0]);
console.log(answers[1]);
