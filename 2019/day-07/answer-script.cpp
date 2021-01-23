#include "../lib-aoc.cpp"
#include <algorithm>

int extract_digit(int a, int n) {
  int nth_power = 1;
  for (int i = 1; i <= n; i++) {
    nth_power *= 10;
  }
  return ((a % nth_power) - (a % (nth_power / 10))) / (nth_power / 10);
}

struct amplifier {
  std::vector<int> instructions;
  std::vector<int> outputs;
  amplifier *prev;
  int instructions_pointer = 0;
};

void run(amplifier *amp, bool exit_on_output) {
  int start = amp->instructions_pointer;
  std::vector<int> *instructions = &amp->instructions;
  std::vector<int> *inputs = &amp->prev->outputs;
  std::vector<int> *outputs = &amp->outputs;
  int output = 0;
  while (start < instructions->size()) {
    int inst = instructions->operator[](start);
    int op = extract_digit(inst, 2) * 10 + extract_digit(inst, 1);
    int p1_mode = extract_digit(inst, 3);
    int p2_mode = extract_digit(inst, 4);
    int p3_mode = extract_digit(inst, 5);
    int result = 0, operand1 = 0, operand2 = 0;
    switch(op) {
      case 1:
      case 2:
        operand1 = instructions->operator[](start + 1);
        if (p1_mode == 0) operand1 = instructions->operator[](operand1);
        operand2 = instructions->operator[](start + 2);
        if (p2_mode == 0) operand2 = instructions->operator[](operand2);
        result = op == 1 ? operand1 + operand2 : operand1 * operand2;
        if (p3_mode == 0) {
          instructions->operator[](instructions->operator[](start + 3)) = result;
        } else {
          printf("Unexpected immediate mode for p3 in instruction %d\n", inst);
          exit(1);
        }
        start += 4;
        break;
      case 3:
        if (p1_mode == 0) {
          if (inputs->size() == 0) {
            amp->instructions_pointer = start;
            run(amp->prev, true);
            run(amp, exit_on_output);
            return;
          } else {
            instructions->operator[](instructions->operator[](start + 1)) = inputs->operator[](0);
            inputs->erase(inputs->begin());
          }
        } else {
          printf("Unexpected immediate mode for p3 in instruction %d\n", inst);
          exit(1);
        }
        start += 2;
        break;
      case 4:
        if (p1_mode == 0) {
          output = instructions->operator[](instructions->operator[](start + 1));
        } else {
          output = instructions->operator[](start + 1);
        }
        outputs->push_back(output);
        start += 2;
        if (exit_on_output) {
          amp->instructions_pointer = start;
          return;
        }
        break;
      case 5:
      case 6:
        operand1 = instructions->operator[](start + 1);
        if (p1_mode == 0) operand1 = instructions->operator[](operand1);
        operand2 = instructions->operator[](start + 2);
        if (p2_mode == 0) operand2 = instructions->operator[](operand2);
        if ((op == 5 && operand1 != 0) || (op == 6 && operand1 == 0)) {
          start = operand2;
        } else {
          start += 3;
        }
        break;
      case 7:
      case 8:
        operand1 = instructions->operator[](start + 1);
        if (p1_mode == 0) operand1 = instructions->operator[](operand1);
        operand2 = instructions->operator[](start + 2);
        if (p2_mode == 0) operand2 = instructions->operator[](operand2);
        result = 0;
        if ((op == 7 && operand1 < operand2) || (op == 8 && operand1 == operand2)) {
          result = 1;
        }
        if (p3_mode == 0) {
          instructions->operator[](instructions->operator[](start + 3)) = result;
        } else {
          printf("Unexpected immediate mode for p3 in instruction %d\n", inst);
          exit(1);
        }
        start += 4;
        break;
      case 99:
        goto end_prog;
    }
  }
  end_prog: ;
  amp->instructions_pointer = start;
  return;
}

int main(int argc, char **argv) {
  const char *content = puzzle_input(argc, argv);
  std::vector<std::string> str_instructions = split(content, ',');
  std::vector<int> instructions;
  for (int i = 0; i < str_instructions.size(); i++) {
    instructions.push_back(stoi(str_instructions[i]));
  }
  int sequence[] = { 0, 1, 2, 3, 4 };
  int max = 0;
  do {
    int output = 0;
    for (int i = 0; i < 5; i++) {
      amplifier amp;
      amplifier prev;
      amp.instructions = instructions;
      amp.prev = &prev;
      prev.outputs.push_back(sequence[i]);
      prev.outputs.push_back(output);
      run(&amp, false);
      output = amp.outputs[0];
    }
    if (output > max) {
      max = output;
    }
  } while(std::next_permutation(sequence, sequence + 5));
  printf("%d\n", max);

  max = 0;
  int sequence2[] = { 5, 6, 7, 8, 9 };
  do {
    std::vector<amplifier> amps;
    amplifier amp0, amp1, amp2, amp3, amp4;

    amp0.instructions = instructions;
    amp1.instructions = instructions;
    amp2.instructions = instructions;
    amp3.instructions = instructions;
    amp4.instructions = instructions;

    amp4.outputs.push_back(sequence2[0]);
    amp4.outputs.push_back(0);
    amp0.outputs.push_back(sequence2[1]);
    amp1.outputs.push_back(sequence2[2]);
    amp2.outputs.push_back(sequence2[3]);
    amp3.outputs.push_back(sequence2[4]);

    amp0.prev = &amp4;
    amp1.prev = &amp0;
    amp2.prev = &amp1;
    amp3.prev = &amp2;
    amp4.prev = &amp3;

    run(&amp4, false);
    if (amp4.outputs[0] > max) {
      max = amp4.outputs[0];
    }
  } while(std::next_permutation(sequence2, sequence2 + 5));
  printf("%d\n", max);
  return 0;
}
