#include "../lib-aoc.cpp"

int main(int argc, char** argv) {
  auto content = puzzle_input(argc, argv);
  auto range = split(content, '-');
  int start = stoi(range[0]);
  int end = stoi(range[1]);

  int count_p1 = 0;
  int count_p2 = 0;
  for (int pass = start; pass <= end; pass++) {
    bool has_adjacents = false;
    bool decreases = false;
    bool exactly_two = false;
    int i = pass;
    int current_digit = i % 10;
    i /= 10;
    int prev_digit = i % 10;
    while (i > 0) {
      int counter = prev_digit == current_digit ? 1 : 0;
      do {
        if (current_digit < prev_digit) goto end_for;
        counter++;
        current_digit = prev_digit;
        i /= 10;
        prev_digit = i % 10;
      } while(prev_digit == current_digit && prev_digit != 0);
      if (counter == 2) exactly_two = true;
      if (counter >= 2) has_adjacents = true;
    }
    if (!decreases && has_adjacents) count_p1 += 1;
    if (!decreases && exactly_two) count_p2 += 1;
    end_for: ;
  }
  printf("%d\n", count_p1);
  printf("%d\n", count_p2);
  return 0;
}
