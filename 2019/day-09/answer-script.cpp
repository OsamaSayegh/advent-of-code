#include "../lib-aoc.cpp"

int main(int argc, char **argv) {
  char *content = puzzle_input(argc, argv);
  std::vector<std::string> str_instructions = split(content, ',');
  free(content);
  std::vector<long> instructions;
  for (long unsigned int i = 0; i < str_instructions.size(); i++) {
    instructions.push_back(stoi(str_instructions[i]));
  }
  for (int value = 1; value <= 2; value++) {
    std::vector<long> inputs;
    inputs.push_back(value);
    intcode_computer input_computer;
    input_computer.outputs = inputs;

    intcode_computer computer;
    computer.dynamic_mem = true;
    computer.instructions = instructions;
    computer.prev = &input_computer;
    computer.run(false);
    printf("%ld\n", computer.outputs[0]);
  }
  return 0;
}
