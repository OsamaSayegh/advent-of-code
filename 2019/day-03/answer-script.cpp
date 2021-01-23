#include "../lib-aoc.cpp"
#include <cstdlib>

struct vec {
  int x1, y1, x2, y2;
};

vec create_vec(int amount, char dir, int x, int y) {
  vec v;
  v.x1 = x;
  v.y1 = y;
  if (dir == 'R') {
    x += amount;
  } else if (dir == 'L') {
    x -= amount;
  } else if (dir == 'U') {
    y += amount;
  } else if (dir == 'D') {
    y -= amount;
  }
  v.x2 = x;
  v.y2 = y;
  return v;
}

int main(int argc, char** argv) {
  auto content = puzzle_input(argc, argv);
  std::vector<vec> wire1_vectors;
  int x = 0, y = 0;
  auto lines = split(content, '\n');
  auto wire1 = split(lines[0].c_str(), ',');
  auto wire2 = split(lines[1].c_str(), ',');
  for (int i = 0; i < wire1.size(); i++) {
    int amount = std::stoi(wire1[i].substr(1, wire1[i].size()));
    char f = wire1[i][0];
    vec v = create_vec(amount, f, x, y);
    x = v.x2;
    y = v.y2;
    wire1_vectors.push_back(v);
  }
  x = 0;
  y = 0;
  int minp1 = 0xfffffff;
  int minp2 = 0xfffffff;

  int wire2_steps = 0;
  for (int i = 0; i < wire2.size(); i++) {
    int amount = std::stoi(wire2[i].substr(1, wire2[i].size()));
    char f = wire2[i][0];
    vec v2 = create_vec(amount, f, x, y);
    x = v2.x2;
    y = v2.y2;

    int v2x1 = v2.x1 < v2.x2 ? v2.x1 : v2.x2;
    int v2x2 = v2.x1 < v2.x2 ? v2.x2 : v2.x1;
    int v2y1 = v2.y1 < v2.y2 ? v2.y1 : v2.y2;
    int v2y2 = v2.y1 < v2.y2 ? v2.y2 : v2.y1;

    int wire1_steps = 0;
    for (int j = 0; j < wire1_vectors.size(); j++) {
      vec v1 = wire1_vectors[j];
      int v1x1 = v1.x1 < v1.x2 ? v1.x1 : v1.x2;
      int v1x2 = v1.x1 < v1.x2 ? v1.x2 : v1.x1;
      int v1y1 = v1.y1 < v1.y2 ? v1.y1 : v1.y2;
      int v1y2 = v1.y1 < v1.y2 ? v1.y2 : v1.y1;

      if (v2x1 <= v1x2 && v2x2 >= v1x1 && v2y1 <= v1y2 && v2y2 >= v1y1) {
        int ix = v2x1 == v2x2 ? v2x1 : v1x1;
        int iy = v2y1 == v2y2 ? v2y1 : v1y1;
        int dis = abs(ix) + abs(iy);
        if (dis == 0) continue;
        if (dis < minp1) {
          minp1 = dis;
        }
        int combined_steps = wire1_steps + abs(v1.x1 == v1.x2 ? iy - v1.y1 : ix - v1.x1);
        combined_steps += wire2_steps + abs(v2.x1 == v2.x2 ? iy - v2.y1 : ix - v2.x1);
        if (combined_steps < minp2) {
          minp2 = combined_steps;
        }
      }
      wire1_steps += v1x1 == v1x2 ? v1y2 - v1y1 : v1x2 - v1x1;
    }
    wire2_steps += v2x1 == v2x2 ? v2y2 - v2y1 : v2x2 - v2x1;
  }
  printf("%d\n", minp1);
  printf("%d\n", minp2);
  return 0;
}
