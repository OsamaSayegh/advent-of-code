#include "../lib-aoc.cpp"

int main(int argc, char **argv) {
  std::vector<long> instructions = intcode_instructions(argc, argv);
  for (size_t i = 0; i < instructions.size(); i++) {
    break;
    printf("%ld", instructions[i]);
    if (i + 1 < instructions.size()) printf(", ");
  }
  intcode_computer computer;
  computer.instructions = instructions;
  computer.dynamic_mem = true;
  {
    intcode_computer input_computer;
    computer.prev = &input_computer;
    std::vector<std::string> commands = {
      "NOT A J",
      "NOT B T",
      "OR T J",
      "NOT C T",
      "OR T J",
      "AND D J",
      "WALK"
    };
    for (size_t i = 0; i < commands.size(); i++) {
      std::string& command = commands[i];
      for (size_t j = 0; j < command.size(); j++) {
        input_computer.outputs.push_back((long) command[j]);
      }
      input_computer.outputs.push_back((long) '\n');
    }
    computer.run(false);

    for (size_t i = 0; i < computer.outputs.size(); i++) {
      if (computer.outputs[i] >= 1 << 7) {
        printf("%ld\n", computer.outputs[i]);
        goto part_two;
      }
    }

    for (size_t i = 0; i < computer.outputs.size(); i++) {
      printf("%c", (char) computer.outputs[i]);
    }
    printf("Failed to find the answer for part 1. Debugging information is above.\n");
    exit(1);
  }
  part_two: ;

  computer.reset(instructions);
  intcode_computer input_computer;
  computer.prev = &input_computer;
  std::vector<std::string> commands = {
    "NOT A T",
    "NOT B J",
    "OR T J",
    "NOT C T",
    "OR T J",
    "AND H J",
    "NOT A T",
    "OR T J",
    "AND D J",
    "RUN"
  };
  for (size_t i = 0; i < commands.size(); i++) {
    std::string& command = commands[i];
    for (size_t j = 0; j < command.size(); j++) {
      input_computer.outputs.push_back((long) command[j]);
    }
    input_computer.outputs.push_back((long) '\n');
  }
  computer.run(false);

  for (size_t i = 0; i < computer.outputs.size(); i++) {
    if (computer.outputs[i] >= 1 << 7) {
      printf("%ld\n", computer.outputs[i]);
      return 0;
    }
  }

  for (size_t i = 0; i < computer.outputs.size(); i++) {
    printf("%c", (char) computer.outputs[i]);
  }
  printf("Failed to find the answer for part 2. Debugging information is above.\n");
  return 1;
}
