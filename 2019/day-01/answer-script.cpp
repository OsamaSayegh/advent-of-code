#include "../lib-aoc.cpp"

int fuel(int a) {
  int r = a / 3 - 2;
  if (r > 0) {
    return r + fuel(r);
  } else {
    return 0;
  }
}

int main(int argc, char** argv) {
  char* content = puzzle_input(argc, argv);
  int sump1 = 0;
  int sump2 = 0;
  auto lines = split(content, '\n');
  for (int i = 0; i < lines.size(); i++) {
    if (lines[i] == "") continue;
    int number = std::stoi(lines[i]);
    sump1 += number / 3 - 2;
    sump2 += fuel(number);
  }
  printf("%d\n", sump1);
  printf("%d\n", sump2);
  return 0;
}
