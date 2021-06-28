#include <stdio.h>
#include <stdlib.h>
#include <vector>
#include <string>
#include <string.h>

char* puzzle_input(int argc, char** argv) {
  FILE* input = fopen(argc > 1 ? argv[1] : "input.txt", "rb");
  if (input == NULL) {
    printf("No such file %s\n", argv[1]);
    exit(1);
  }
  fseek(input, 0, SEEK_END);
  size_t size = ftell(input);
  char* content = (char*) malloc(size + 1);
  fseek(input, 0, SEEK_SET);
  fread(content, 1, size, input);
  fclose(input);
  content[size] = '\0';
  return content;
}

std::vector<std::string> split(const char* input, char del) {
  long unsigned int len = strlen(input);
  size_t start = 0;
  std::vector<std::string> res;
  for (long unsigned int i = 0; i < len; i++) {
    if (input[i] == del) {
      res.push_back(std::string(input + start, i - start));
      start = i + 1;
    }
  }
  res.push_back(std::string(input + start, len - start));
  return res;
}

std::vector<long> intcode_instructions(int argc, char **argv) {
  char *content = puzzle_input(argc, argv);
  std::vector<std::string> str_instructions = split(content, ',');
  free(content);
  std::vector<long> inst;
  for (size_t i = 0; i < str_instructions.size(); i++) {
    inst.push_back(stol(str_instructions[i]));
  }
  return inst;
}

int extract_digit(long a, int n) {
  int nth_power = 1;
  for (int i = 1; i <= n; i++) {
    nth_power *= 10;
  }
  return ((int) (a % nth_power) - (int) (a % (nth_power / 10))) / (nth_power / 10);
}

struct intcode_computer {
  std::vector<long> instructions;
  std::vector<long> outputs;
  intcode_computer *prev;
  long unsigned int inst_pointer = 0;
  long relative_base = 0;
  bool dynamic_mem = false;
  bool halted = false;

  void write(long position, long value) {
    if (position < 0 || ((size_t) position >= instructions.size() && !dynamic_mem)) {
      printf("write error position: %ld, size: %ld\n", position, instructions.size());
      exit(1);
    }
    if ((size_t) position >= instructions.size()) {
      instructions.resize(position + 1);
    }
    instructions[position] = value;
  }

  long read(long position) {
    if (position < 0 || ((size_t) position >= instructions.size() && !dynamic_mem)) {
      printf("read error position: %ld, size: %ld\n", position, instructions.size());
      exit(1);
    }
    if ((size_t) position >= instructions.size()) {
      instructions.resize(position + 1);
    }
    return instructions[position];
  }

  long read_output(size_t index) {
    if (index >= outputs.size()) {
      printf("Invalid output reads at index %ld, size: %ld\n", index, outputs.size());
      exit(1);
    }
    return outputs[index];
  }

  void set_operand(long *operand, int p_mode, int offset) {
    *operand = read(inst_pointer + offset);
    if (p_mode == 0) *operand = read(*operand);
    if (p_mode == 2) *operand = read(relative_base + *operand);
  }

  void reset(std::vector<long> new_instructions) {
    instructions = new_instructions;
    outputs.clear();
    inst_pointer = 0;
    halted = false;
    relative_base = 0;
  }

  void run(bool suspend_on_output) {
    std::vector<long> *inputs = &prev->outputs;
    long output = 0;
    while (inst_pointer < instructions.size()) {
      long inst = read(inst_pointer);
      int op = extract_digit(inst, 2) * 10 + extract_digit(inst, 1);
      int p1_mode = extract_digit(inst, 3);
      int p2_mode = extract_digit(inst, 4);
      int p3_mode = extract_digit(inst, 5);
      long result = 0, operand1 = 0, operand2 = 0;
      switch(op) {
        case 1:
        case 2:
          set_operand(&operand1, p1_mode, 1);
          set_operand(&operand2, p2_mode, 2);
          result = op == 1 ? operand1 + operand2 : operand1 * operand2;
          if (p3_mode == 0) {
            write(read(inst_pointer + 3), result);
          } else if (p3_mode == 1) {
            printf("Unexpected immediate mode for p3 in instruction %ld\n", inst);
            exit(1);
          } else if (p3_mode == 2) {
            write(relative_base + read(inst_pointer + 3), result);
          } else {
            printf("Unknown mode for p3 in instruction %ld\n", inst);
            exit(1);
          }
          inst_pointer += 4;
          break;
        case 3:
          if (inputs->size() == 0) {
            prev->run(true);
            run(suspend_on_output);
            return;
          }
          if (p1_mode == 0) {
            write(read(inst_pointer + 1), (*inputs)[0]);
            inputs->erase(inputs->begin());
          } else if (p1_mode == 1) {
            printf("Unexpected immediate mode for p1 in instruction %ld\n", inst);
            exit(1);
          } else if (p1_mode == 2) {
            write(relative_base + read(inst_pointer + 1), (*inputs)[0]);
            inputs->erase(inputs->begin());
          } else {
            printf("Unknown mode for p1 in instruction %ld\n", inst);
            exit(1);
          }
          inst_pointer += 2;
          break;
        case 4:
          if (p1_mode == 0) {
            output = read(read(inst_pointer + 1));
          } else if (p1_mode == 1) {
            output = read(inst_pointer + 1);
          } else if (p1_mode == 2) {
            output = read(relative_base + read(inst_pointer + 1));
          } else {
            printf("Unknown mode for p1 in instruction %ld\n", inst);
            exit(1);
          }
          outputs.push_back(output);
          inst_pointer += 2;
          if (suspend_on_output) {
            goto end_prog;
          }
          break;
        case 5:
        case 6:
          set_operand(&operand1, p1_mode, 1);
          set_operand(&operand2, p2_mode, 2);
          if ((op == 5 && operand1 != 0) || (op == 6 && operand1 == 0)) {
            inst_pointer = operand2;
          } else {
            inst_pointer += 3;
          }
          break;
        case 7:
        case 8:
          set_operand(&operand1, p1_mode, 1);
          set_operand(&operand2, p2_mode, 2);
          result = 0;
          if ((op == 7 && operand1 < operand2) || (op == 8 && operand1 == operand2)) {
            result = 1;
          }
          if (p3_mode == 0) {
            write(read(inst_pointer + 3), result);
          } else if (p3_mode == 1) {
            printf("Unexpected immediate mode for p3 in instruction %ld\n", inst);
            exit(1);
          } else if (p3_mode == 2) {
            write(relative_base + read(inst_pointer + 3), result);
          } else {
            printf("Unknown mode for p3 in instruction %ld\n", inst);
            exit(1);
          }
          inst_pointer += 4;
          break;
        case 9:
          set_operand(&operand1, p1_mode, 1);
          relative_base += operand1;
          inst_pointer += 2;
          break;
        case 99:
          halted = true;
          goto end_prog;
      }
    }
    end_prog: ;
  }
};
