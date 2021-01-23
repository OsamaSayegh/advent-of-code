#include "../lib-aoc.cpp"

int extract_digit(int a, int n) {
  int nth_power = 1;
  for (int i = 1; i <= n; i++) {
    nth_power *= 10;
  }
  return ((a % nth_power) - (a % (nth_power / 10))) / (nth_power / 10);
}

int run(std::vector<int> instructions, int input, bool p2) {
  int output = 0;
  for (int i = 0; i < instructions.size();) {
    int inst = instructions[i];
    int op = extract_digit(inst, 2) * 10 + extract_digit(inst, 1);
    int p1_mode = extract_digit(inst, 3);
    int p2_mode = extract_digit(inst, 4);
    int p3_mode = extract_digit(inst, 5);
    int result = 0, operand1 = 0, operand2 = 0;
    switch(op) {
      case 1:
      case 2:
        operand1 = instructions[i + 1];
        if (p1_mode == 0) operand1 = instructions[operand1];
        operand2 = instructions[i + 2];
        if (p2_mode == 0) operand2 = instructions[operand2];
        result = op == 1 ? operand1 + operand2 : operand1 * operand2;
        if (p3_mode == 0) {
          instructions[instructions[i + 3]] = result;
        } else {
          printf("Unexpected immediate mode for p3 in instruction %d\n", inst);
          exit(1);
        }
        i += 4;
        break;
      case 3:
        if (p1_mode == 0) {
          instructions[instructions[i + 1]] = input;
        } else {
          printf("Unexpected immediate mode for p3 in instruction %d\n", inst);
          exit(1);
        }
        i += 2;
        break;
      case 4:
        if (p1_mode == 0) {
          output = instructions[instructions[i + 1]];
        } else {
          output = instructions[i + 1];
        }
        i += 2;
        break;
      case 5:
      case 6:
        if (!p2) break;
        operand1 = instructions[i + 1];
        if (p1_mode == 0) operand1 = instructions[operand1];
        operand2 = instructions[i + 2];
        if (p2_mode == 0) operand2 = instructions[operand2];
        if ((op == 5 && operand1 != 0) || (op == 6 && operand1 == 0)) {
          i = operand2;
        } else {
          i += 3;
        }
        break;
      case 7:
      case 8:
        if (!p2) break;
        operand1 = instructions[i + 1];
        if (p1_mode == 0) operand1 = instructions[operand1];
        operand2 = instructions[i + 2];
        if (p2_mode == 0) operand2 = instructions[operand2];
        result = 0;
        if ((op == 7 && operand1 < operand2) || (op == 8 && operand1 == operand2)) {
          result = 1;
        }
        if (p3_mode == 0) {
          instructions[instructions[i + 3]] = result;
        } else {
          printf("Unexpected immediate mode for p3 in instruction %d\n", inst);
          exit(1);
        }
        i += 4;
        break;
      case 99:
        goto end_prog;
    }
  }
  end_prog: ;
  return output;
}

int main(int argc, char** argv) {
  char* content = puzzle_input(argc, argv);
  std::vector<std::string> str_instructions = split(content, ',');
  std::vector<int> instructions;

  for (int i = 0; i < str_instructions.size(); i++) {
    instructions.push_back(stoi(str_instructions[i]));
  }

  printf("%d\n", run(instructions, 1, false));
  printf("%d\n", run(instructions, 5, true));
  return 0;
}
