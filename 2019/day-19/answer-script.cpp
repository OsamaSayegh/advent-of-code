#include "../lib-aoc.cpp"
#include <math.h>

bool test_coordinates(
  intcode_computer computer,
  long x,
  long y
) {
  computer.prev->outputs.push_back(x);
  computer.prev->outputs.push_back(y);
  computer.run(true);
  if (computer.outputs.size() == 1) {
    long output = computer.outputs[0];
    if (output == 1) {
      return true;
    } else if (output == 0) {
      return false;
    } else {
      printf("Unexpected output %ld\n", output);
      exit(1);
    }
  } else {
    printf("Unexpected output size %ld\n", computer.outputs.size());
    exit(1);
  }
}

long find_beam_right_edge(intcode_computer computer, long x, long y) {
  long direction = -1;
  if (test_coordinates(computer, x, y)) {
    direction = 1;
  }
  x += direction;
  while(!test_coordinates(computer, x, y)) x += direction;
  return x;
}

int main(int argc, char **argv) {
  std::vector<long> instructions = intcode_instructions(argc, argv);
  intcode_computer computer;
  computer.dynamic_mem = true;
  computer.instructions = instructions;
  intcode_computer input_computer;
  computer.prev = &input_computer;
  int total = 0;
  long base_begin = -1;
  long base_end = -1;
  double beam_end_ratio_sum = 0.0, beam_size_ratio_sum = 0.0;
  long ratio_total = 0;
  for (long y = 0; y < 50; y++) {
    bool in_beam = false;
    for (long x = 0; x < 50; x++) {
      bool was_in_beam = in_beam;
      in_beam = test_coordinates(computer, x, y);
      if (in_beam) total++;
      if (was_in_beam != in_beam) {
        if (was_in_beam) {
          base_end = x;
        } else {
          base_begin = x;
        }
      }
    }
    if (base_end - base_begin >= 3) {
      beam_end_ratio_sum += (double) base_end / (double) y;
      beam_size_ratio_sum += (double) (base_end - base_begin) / (double) y;
      ratio_total++;
    }
  }

  printf("%d\n", total);

  long square_size = 100;
  double beam_end_ratio_avg = beam_end_ratio_sum / (double) ratio_total;
  double beam_size_ratio_avg = beam_size_ratio_sum / (double) ratio_total;
  long y = (long) ((double) square_size / beam_size_ratio_avg);
  square_size--; // to prevent off by one errors in the following while loop
  long x;
  while (true) {
    long est_x = (long) (beam_end_ratio_avg * (double) y);
    x = find_beam_right_edge(computer, est_x, y);
    if (
      test_coordinates(computer, x - square_size, y + square_size) &&
      !test_coordinates(computer, x - (square_size + 1), y + square_size)
    ) break;
    y++;
  }
  x -= square_size;
  printf("%ld\n", x * 10000 + y);
  return 0;
}
