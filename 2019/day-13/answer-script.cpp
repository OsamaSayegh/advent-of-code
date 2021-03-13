#include "../lib-aoc.cpp"
#include <map>

int part1(intcode_computer computer) {
  std::map<long, long> grid;
  int block_count = 0;

  while(true) {
    computer.run(true);
    computer.run(true);
    computer.run(true);
    if (computer.halted) break;
    long x = computer.read_output(0);
    long y = computer.read_output(1);
    long t = computer.read_output(2);
    computer.outputs.erase(computer.outputs.begin());
    computer.outputs.erase(computer.outputs.begin());
    computer.outputs.erase(computer.outputs.begin());
    long key = x * 1000 + y;
    long existing_t = grid.count(key) == 0 ? -1 : grid[key];
    grid[key] = t;
    if (existing_t == 2 && t != 2) block_count--;
    if (existing_t != 2 && t == 2) block_count++;
  }
  return block_count;
}

long part2(intcode_computer computer) {
  computer.instructions[0] = 2;
  intcode_computer input_computer;
  computer.prev = &input_computer;
  long score = 0;
  long ball = 0;
  long paddle = 0;
  while(true) {
    computer.run(true);
    computer.run(true);
    computer.run(true);
    if (computer.halted) break;
    long x = computer.read_output(0);
    long y = computer.read_output(1);
    long t = computer.read_output(2);
    computer.outputs.erase(computer.outputs.begin());
    computer.outputs.erase(computer.outputs.begin());
    computer.outputs.erase(computer.outputs.begin());
    if (x == -1 && y == 0) {
      score = t;
      continue;
    } else if (t == 4) {
      ball = x;
    } else if (t == 3) {
      paddle = x;
      continue;
    } else {
      continue;
    }
    if (paddle < ball) {
      input_computer.outputs.push_back(1);
    } else if (paddle > ball) {
      input_computer.outputs.push_back(-1);
    } else {
      input_computer.outputs.push_back(0);
    }
  }
  return score;
}

int main(int argc, char **argv) {
  std::vector<long> instructions = intcode_instructions(argc, argv);
  intcode_computer computer;
  computer.instructions = instructions;
  computer.dynamic_mem = true;
  printf("%d\n", part1(computer));
  printf("%ld\n", part2(computer));
  return 0;
}
