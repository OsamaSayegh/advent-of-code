#include "../lib-aoc.cpp"
#include <map>

int pos_to_int(int a, int b) {
  return a * 1000 + b;
}

int paint(std::vector<long>& instructions, long first_square_color) {
  intcode_computer input_computer;
  intcode_computer primary;
  primary.instructions = instructions;
  primary.prev = &input_computer;
  primary.dynamic_mem = true;
  int pos[2] = { 0, 0 };
  int steps[2] = { 0, -1 };
  std::map<int, long> grid;
  grid[0] = first_square_color;
  int count = 1;

  int max_x = 0x80000000;
  int min_x = 0x7fffffff;
  int max_y = 0x80000000;
  int min_y = 0x7fffffff;
  while (true) {
    if (pos[1] > max_y) max_y = pos[1];
    if (pos[1] < min_y) min_y = pos[1];
    if (pos[0] > max_x) max_x = pos[0];
    if (pos[0] < min_x) min_x = pos[0];

    int location = pos_to_int(pos[0], pos[1]);
    long color = grid.count(location) == 0 ? 0 : grid[location];
    input_computer.outputs.push_back(color);

    primary.run(true);
    primary.run(true);
    if (primary.halted) {
      break;
    }

    long new_color = primary.outputs[0];
    if (new_color != 0 && new_color != 1) {
      printf("Unexpected color value: %ld\n", new_color);
      exit(1);
    }
    if (grid.count(location) == 0) count++;
    grid[location] = new_color;
    primary.outputs.erase(primary.outputs.begin());

    long new_dir = primary.outputs[1];
    if (new_dir == 0) {
      int tmp = steps[0];
      steps[0] = steps[1];
      steps[1] = -tmp;
    } else if (new_dir == 1) {
      int tmp = steps[0];
      steps[0] = -steps[1];
      steps[1] = tmp;
    } else {
      printf("Unexpected new direction value: %ld\n", new_dir);
      exit(1);
    }
    primary.outputs.erase(primary.outputs.begin());
    pos[0] += steps[0];
    pos[1] += steps[1];
  }
  if (first_square_color == 1) {
    for (int y = min_y; y <= max_y; y++) {
      for (int x = min_x; x <= max_x; x++) {
        int location = pos_to_int(x, y);
        long color = grid.count(location) == 0 ? 0 : grid[location];
        if (color == 1) {
          printf("O");
        } else {
          printf(" ");
        }
      }
      printf("\n");
    }
  }
  return count;
}

int main(int argc, char **argv) {
  char *content = puzzle_input(argc, argv);
  std::vector<std::string> str_instructions = split(content, ',');
  free(content);
  std::vector<long> instructions;
  for (size_t i = 0; i < str_instructions.size(); i++) {
    instructions.push_back(stol(str_instructions[i]));
  }
  printf("%d\n", paint(instructions, 0));
  paint(instructions, 1);
  return 0;
}
