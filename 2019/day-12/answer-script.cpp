#include "../lib-aoc.cpp"
#include <numeric>

struct moon {
  int x, y, z;
  int vx = 0, vy = 0, vz = 0;
};

// thanks to https://stackoverflow.com/a/4229930
long gcd(long a, long b) {
  for (;;) {
    if (a == 0) return b;
    b %= a;
    if (b == 0) return a;
    a %= b;
  }
}

long lcm(long a, long b) {
  long temp = gcd(a, b);
  return temp ? (a / temp * b) : 0;
}

moon parse(std::string line) {
  moon m;
  size_t s = line.find_first_of("=") + 1;
  size_t e = line.find_first_of(",");
  m.x = stoi(line.substr(s, e - s));
  s = line.find_first_of("=", s) + 1;
  e = line.find_first_of(",", e + 1);
  m.y = stoi(line.substr(s, e - s));
  s = line.find_first_of("=", s) + 1;
  e = line.find_first_of(">", e + 1);
  m.z = stoi(line.substr(s, e - s));
  return m;
}

int main(int argc, char **argv) {
  char *content = puzzle_input(argc, argv);
  std::vector<std::string> lines = split(content, '\n');
  std::vector<moon> moons;
  for (size_t i = 0; i < lines.size(); i++) {
    if (lines[i].size() == 0) continue;
    moons.push_back(parse(lines[i]));
  }
  std::vector<moon> init_state = moons;
  std::vector<moon> moons2 = moons;
  free(content);

  for (int s = 1; s <= 1000; s++) {
    for (size_t i = 0; i < moons.size() - 1; i++) {
      moon& m1 = moons[i];
      for (size_t j = i + 1; j < moons.size(); j++) {
        moon& m2 = moons[j];
        if (m1.x < m2.x) {
          m1.vx++;
          m2.vx--;
        } else if (m1.x > m2.x) {
          m1.vx--;
          m2.vx++;
        }
        if (m1.y < m2.y) {
          m1.vy++;
          m2.vy--;
        } else if (m1.y > m2.y) {
          m1.vy--;
          m2.vy++;
        }
        if (m1.z < m2.z) {
          m1.vz++;
          m2.vz--;
        } else if (m1.z > m2.z) {
          m1.vz--;
          m2.vz++;
        }
      }
    }
    for (size_t g = 0; g < moons.size(); g++) {
      moon& m = moons[g];
      m.x += m.vx;
      m.y += m.vy;
      m.z += m.vz;
    }
  }
  int total = 0;
  for (size_t i = 0; i < moons.size(); i++) {
    moon m = moons[i];
    int pe = abs(m.x) + abs(m.y) + abs(m.z);
    int ke = abs(m.vx) + abs(m.vy) + abs(m.vz);
    total += pe * ke;
  }
  printf("%d\n", total);

  long x_steps = 0;
  while (true) {
    x_steps++;
    for (size_t i = 0; i < moons2.size(); i++) {
      moon& m1 = moons2[i];
      for (size_t j = i + 1; j < moons2.size(); j++) {
        moon& m2 = moons2[j];
        if (m1.x < m2.x) {
          m1.vx++;
          m2.vx--;
        } else if (m1.x > m2.x) {
          m1.vx--;
          m2.vx++;
        }
      }
    }
    bool idential = true;
    for (size_t g = 0; g < moons2.size(); g++) {
      moon& m1 = moons2[g];
      m1.x += m1.vx;
      moon& m2 = init_state[g];
      if (m1.x != m2.x || m1.vx != m2.vx) idential = false;
    }
    if (idential) break;
  }

  long y_steps = 0;
  while (true) {
    y_steps++;
    for (size_t i = 0; i < moons2.size(); i++) {
      moon& m1 = moons2[i];
      for (size_t j = i + 1; j < moons2.size(); j++) {
        moon& m2 = moons2[j];
        if (m1.y < m2.y) {
          m1.vy++;
          m2.vy--;
        } else if (m1.y > m2.y) {
          m1.vy--;
          m2.vy++;
        }
      }
    }
    bool idential = true;
    for (size_t g = 0; g < moons2.size(); g++) {
      moon& m1 = moons2[g];
      m1.y += m1.vy;
      moon& m2 = init_state[g];
      if (m1.y != m2.y || m1.vy != m2.vy) idential = false;
    }
    if (idential) break;
  }

  long z_steps = 0;
  while (true) {
    z_steps++;
    for (size_t i = 0; i < moons2.size(); i++) {
      moon& m1 = moons2[i];
      for (size_t j = i + 1; j < moons2.size(); j++) {
        moon& m2 = moons2[j];
        if (m1.z < m2.z) {
          m1.vz++;
          m2.vz--;
        } else if (m1.z > m2.z) {
          m1.vz--;
          m2.vz++;
        }
      }
    }
    bool idential = true;
    for (size_t g = 0; g < moons2.size(); g++) {
      moon& m1 = moons2[g];
      m1.z += m1.vz;
      moon& m2 = init_state[g];
      if (m1.z != m2.z || m1.vz != m2.vz) idential = false;
    }
    if (idential) break;
  }
  printf("%ld\n", lcm(lcm(x_steps, y_steps), z_steps));
  return 0;
}
