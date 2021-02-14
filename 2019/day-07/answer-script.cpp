#include "../lib-aoc.cpp"
#include <algorithm>

int main(int argc, char **argv) {
  char *content = puzzle_input(argc, argv);
  std::vector<std::string> str_instructions = split(content, ',');
  free(content);
  std::vector<long> instructions;
  for (long unsigned int i = 0; i < str_instructions.size(); i++) {
    instructions.push_back(stoi(str_instructions[i]));
  }
  int sequence[] = { 0, 1, 2, 3, 4 };
  long max = 0;
  do {
    long output = 0;
    for (int i = 0; i < 5; i++) {
      intcode_computer computer;
      intcode_computer prev;
      computer.instructions = instructions;
      computer.prev = &prev;
      prev.outputs.push_back(sequence[i]);
      prev.outputs.push_back(output);
      computer.run(false);
      output = computer.outputs[0];
    }
    if (output > max) {
      max = output;
    }
  } while(std::next_permutation(sequence, sequence + 5));
  printf("%ld\n", max);

  max = 0;
  int sequence2[] = { 5, 6, 7, 8, 9 };
  do {
    intcode_computer computer0, computer1, computer2, computer3, computer4;

    computer0.instructions = instructions;
    computer1.instructions = instructions;
    computer2.instructions = instructions;
    computer3.instructions = instructions;
    computer4.instructions = instructions;

    computer4.outputs.push_back(sequence2[0]);
    computer4.outputs.push_back(0);
    computer0.outputs.push_back(sequence2[1]);
    computer1.outputs.push_back(sequence2[2]);
    computer2.outputs.push_back(sequence2[3]);
    computer3.outputs.push_back(sequence2[4]);

    computer0.prev = &computer4;
    computer1.prev = &computer0;
    computer2.prev = &computer1;
    computer3.prev = &computer2;
    computer4.prev = &computer3;

    computer4.run(false);
    if (computer4.outputs[0] > max) {
      max = computer4.outputs[0];
    }
  } while(std::next_permutation(sequence2, sequence2 + 5));
  printf("%ld\n", max);
  return 0;
}
