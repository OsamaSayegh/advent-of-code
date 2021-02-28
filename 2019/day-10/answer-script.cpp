#include "../lib-aoc.cpp"
#include <algorithm>
#include <array>
#include <map>
#include <math.h>

void rotate_90_clock(int (&point)[2]) {
  int t = point[0];
  point[0] = -point[1];
  point[1] = t;
}

struct asteroid {
  double slope = 0.0;
  double distance = 0.0;
  int pos[2] = { -1, -1 };
  int quarter = -1;
};

int main(int argc, char **argv) {
  char *input = puzzle_input(argc, argv);
  std::vector<std::string> lines = split(input, '\n');
  free(input);
  lines.erase(lines.end());
  std::vector<std::vector<char>> asteroids;
  for (size_t i = 0; i < lines.size(); i++) {
    const char *current = lines[i].c_str();
    std::vector<char> row;
    for (size_t j = 0; j < strlen(current); j++) {
      row.push_back(current[j]);
    }
    asteroids.push_back(row);
  }
  int size = (int) asteroids.size();
  int max = 0;
  std::vector<asteroid> vaporized;

  for (int y = 0; y < size; y++) {
    std::vector<char> row = asteroids[y];
    for (int x = 0; x < size; x++) {
      if (row[x] == '#') {
        int total = 0;
        std::vector<asteroid> current_vaporized;
        for (int quarter = 0; quarter < 4; quarter++) {
          int distance = 1;
          std::vector<double> blocked_slopes;
          while (distance < size) {
            int r_pos[2] = { 0, -distance };
            int steps[2] = { 1, 0 };
            for (int r = 0; r < quarter; r++) {
              rotate_90_clock(r_pos);
              rotate_90_clock(steps);
            }
            for (int i = 0; i <= 1; i++) {
              for (int s = 1; s <= distance; s++) {
                int x1 = x + r_pos[0];
                int y1 = y + r_pos[1];
                double slope = (r_pos[1]) / (double) r_pos[0];
                if (y1 >= 0 && size > y1 && x1 >= 0 && size > x1 && asteroids[y1][x1] == '#') {
                  bool blocked = std::find(blocked_slopes.begin(), blocked_slopes.end(), slope) != blocked_slopes.end();
                  if (!blocked) {
                    total++;
                    blocked_slopes.push_back(slope);
                  }
                  double distance = sqrt(pow(r_pos[0], 2.0) + pow(r_pos[1], 2.0));
                  asteroid s;
                  s.pos[0] = x1;
                  s.pos[1] = y1;
                  s.distance = distance;
                  s.quarter = quarter;
                  s.slope = slope;
                  current_vaporized.push_back(s);
                }
                r_pos[0] += steps[0];
                r_pos[1] += steps[1];
              }
              rotate_90_clock(steps);
            }
            distance++;
          }
        }
        if (total > max) {
          max = total;
          vaporized.clear();
          for (asteroid s : current_vaporized) {
            vaporized.push_back(s);
          }
        }
      }
    }
  }
  printf("%d\n", max);
  sort(vaporized.begin(), vaporized.end(), [](const asteroid &a, const asteroid &b) -> bool {
    if (a.quarter > b.quarter) {
      return false;
    }
    if (a.quarter < b.quarter) {
      return true;
    }
    if (a.slope < b.slope) {
      return true;
    }
    if (a.slope > b.slope) {
      return false;
    }
    if (a.distance > b.distance) {
      return false;
    }
    if (a.distance < b.distance) {
      return true;
    }
    printf("Unepxected condition when sorting asteroids.\n");
    exit(1);
  });
  int no = 1;
  int index = 0;
  asteroid the200th = vaporized[index];
  while(no < 200) {
    ++index;
    index %= (int) vaporized.size();
    if (the200th.slope == vaporized[index].slope) {
      continue;
    }
    the200th = vaporized[index];
    no++;
  }
  printf("%d\n", the200th.pos[0] * 100 + the200th.pos[1]);
  return 0;
}
