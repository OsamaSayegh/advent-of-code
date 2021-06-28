#include "../lib-aoc.cpp"
#include <algorithm>

inline bool is_bug(int grid, int pos) {
  return (grid & (1 << pos)) != 0;
}

const int size = 5;
const int size_square = size * size;

void check_adjacent(int pos, int distance, int current, int inner, int outer, int& bugs) {
  int adjacent = pos + distance;
  if (adjacent == size_square / 2) {
    if (distance == 1) {
      for (int i = 0; i < size_square; i += size) {
        if (is_bug(inner, i)) bugs++;
      }
    } else if (distance == -1) {
      for (int i = size - 1; i < size_square; i += size) {
        if (is_bug(inner, i)) bugs++;
      }
    } else if (distance == size) {
      for (int i = 0; i < size; i++) {
        if (is_bug(inner, i)) bugs++;
      }
    } else if (distance == -size) {
      for (int i = size_square - 1; i >= size_square - size; i--) {
        if (is_bug(inner, i)) bugs++;
      }
    } else {
      printf("Error: unexpected travel distance %d\n", distance);
      exit(1);
    }
  } else if (pos % size == size - 1 && distance == 1) {
    if (is_bug(outer, size_square / 2 + 1)) bugs++;
  } else if (pos % size == 0 && distance == -1) {
    if (is_bug(outer, size_square / 2 - 1)) bugs++;
  } else if (pos < size && distance == -size) {
    if (is_bug(outer, size_square / 2 - size)) bugs++;
  } else if (pos >= size_square - size && distance == size) {
    if (is_bug(outer, size_square / 2 + size)) bugs++;
  } else {
    if (is_bug(current, adjacent)) bugs++;
  }
}

int main(int argc, char **argv) {
  char *content = puzzle_input(argc, argv);
  int first = 0;
  int c = 0;
  for (char *copy = content; *copy != '\0'; copy++) {
    char cur = *copy;
    if (cur == '#' || cur == '.') {
      if (cur == '#') first |= 1 << c;
      c++;
    }
  }
  free(content);
  std::vector<int> stages = { first };
  int i = 0;
  while (true) {
    int prev = stages[i];
    int grid = prev;
    for (int j = size_square - 1; j >= 0; j--) {
      int bugs = 0;
      if (j + 1 <= size_square - size && is_bug(prev, j + size)) bugs++;
      if (j >= size && is_bug(prev, j - size)) bugs++;
      if (j % size != 0 && is_bug(prev, j - 1)) bugs++;
      if (j % size != (size - 1) && is_bug(prev, j + 1)) bugs++;
      if (is_bug(prev, j) && bugs != 1) {
        grid &= ~(1 << j);
      }
      if (!is_bug(prev, j) && (bugs == 1 || bugs == 2)) {
        grid |= 1 << j;
      }
    }
    if (std::count(stages.begin(), stages.end(), grid)) {
      printf("%d\n", grid);
      break;
    }
    stages.push_back(grid);
    i++;
  }

  std::vector<int> levels = { 0, first, 0 };
  for (int i = 1; i <= 200; i++) {
    std::vector<int> levels_copy = levels;
    for (size_t j = 0; j < levels.size(); j++) {
      int current = levels[j];
      int outer = j == 0 ? 0 : levels[j - 1];
      int inner = j + 1 == levels.size() ? 0 : levels[j + 1];

      int new_current = current;
      for (int p = size_square - 1; p >= 0; p--) {
        if (p == size_square / 2) continue;
        int bugs = 0;
        for (int d = 1; d <= size; d += size - 1) {
          do {
            check_adjacent(p, d, current, inner, outer, bugs);
            d *= -1;
          } while(d < 0);
        }
        if (is_bug(current, p) && bugs != 1) {
          new_current &= ~(1 << p);
        }
        if (!is_bug(current, p) && (bugs == 1 || bugs == 2)) {
          new_current |= 1 << p;
        }
      }
      levels_copy[j] = new_current;
    }
    if (levels_copy[0] != 0) levels_copy.insert(levels_copy.begin(), 0);
    if (levels_copy[levels_copy.size() - 1] != 0) levels_copy.push_back(0);
    levels = levels_copy;
  }
  int count = 0;
  for (size_t m = 0; m < levels.size(); m++) {
    for (int b = 0; b < (1 << 5); b++) {
      if (is_bug(levels[m], b)) {
        if (b >= size_square) {
          printf("Error: Unexpected set bit in position %d! value: %d.\n", b, levels[m]);
          exit(1);
        }
        count++;
      }
    }
  }
  printf("%d\n", count);
  return 0;
}
