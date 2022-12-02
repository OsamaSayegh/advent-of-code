const fs = require("fs");
const input = fs.readFileSync(0).toString().trim();

const SCORES = {
  A: 1, // Rock
  B: 2, // Paper
  C: 3, // Scissors
};

const MAPPING = {
  A: 'C',
  B: 'A',
  C: 'B',
};

const rounds = input.split("\n").map(round => round.split(" "));
const scores = rounds.reduce((acc, round) => {
  const opponent = round[0];
  const me = String.fromCharCode(round[1].charCodeAt(0) - 23);
  const outcome = round[1];

  let score1 = acc[0];
  let score2 = acc[1];

  score1 += SCORES[me];
  if (MAPPING[me] === opponent) {
    score1 += 6;
  } else if (MAPPING[opponent] !== me) {
    score1 += 3;
  }

  if (outcome === "X") {
    score2 += SCORES[MAPPING[opponent]];
  } else if (outcome === "Y") {
    score2 += SCORES[opponent] + 3;
  } else {
    score2 += SCORES[MAPPING[MAPPING[opponent]]] + 6;
  }
  return [score1, score2];
}, [0, 0]);

console.log(scores[0]);
console.log(scores[1]);
