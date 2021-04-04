#include "../lib-aoc.cpp"
#include <map>

const long NORTH = 1, SOUTH = 2, WEST = 3, EAST = 4;
const long WALL = 0, MOVED = 1, OXYGEN = 2, UNKNOWN = -1;

long max_x = (long) 1 << 63;
long min_x = ~max_x;
long max_y = max_x;
long min_y = min_x;

long inline pos_to_num(long a, long b) {
  return a * 10000 + b;
}

long inline opposite_dir(long dir) {
  switch (dir) {
    case NORTH:
      return SOUTH;
    case SOUTH:
      return NORTH;
    case EAST:
      return WEST;
    case WEST:
      return EAST;
    default:
      printf("Invalid direction %ld\n", dir);
      exit(1);
  }
}

void inline add_to_maze(long x, long y, long type, std::map<long, long>& maze) {
  if (x > max_x) {
    max_x = x;
  }
  if (x < min_x) {
    min_x = x;
  }
  if (y > max_y) {
    max_y = y;
  }
  if (y < min_y) {
    min_y = y;
  }
  maze[pos_to_num(x, y)] = type;
}

long move_droid(long& x, long& y, intcode_computer& computer, long dir) {
  computer.prev->outputs.push_back(dir);
  computer.run(true);
  if (computer.outputs.size() != 1) {
    printf("Expected a single output but saw %ld\n", computer.outputs.size());
    exit(1);
  }
  long out = computer.outputs[0];
  if (out != WALL) {
    switch(dir) {
      case NORTH:
        y += 1; break;
      case EAST:
        x += 1; break;
      case SOUTH:
        y -= 1; break;
      case WEST:
        x -= 1; break;
      default:
        printf("Unknown direction %ld\n", dir);
        exit(1);
    }
  }
  computer.outputs.erase(computer.outputs.begin());
  return out;
}

bool explore(
  long& x,
  long& y,
  intcode_computer& computer,
  std::map<long, long>& maze,
  long& path,
  bool full_map
) {
  long num = pos_to_num(x, y);
  long current = maze.count(num) == 1 ? maze[num] : UNKNOWN;
  if (current == WALL || current == MOVED) {
    return false;
  }
  add_to_maze(x, y, MOVED, maze);
  bool found = false;
  for (long dir : { NORTH, EAST, SOUTH, WEST }) {
    long output = move_droid(x, y, computer, dir);
    if (output == OXYGEN) {
      add_to_maze(x, y, OXYGEN, maze);
      found = true;
      if (!full_map) {
        path += 1;
        break;
      }
      move_droid(x, y, computer, opposite_dir(dir));
    }
    if (output == WALL) {
      long wx = x, wy = y;
      switch (dir) {
        case NORTH:
          wy += 1; break;
        case SOUTH:
          wy -= 1; break;
        case EAST:
          wx += 1; break;
        case WEST:
          wx -= 1; break;
        default:
          printf("Unknown direction %ld\n", dir);
          exit(1);
      }
      add_to_maze(wx, wy, WALL, maze);
    }
    if (output == MOVED) {
      found = explore(x, y, computer, maze, path, full_map);
      if (found) {
        if (!full_map) {
          path += 1;
          break;
        }
      }
      move_droid(x, y, computer, opposite_dir(dir));
    }
  }
  return found;
}

int main(int argc, char **argv) {
  std::vector<long> instructions = intcode_instructions(argc, argv);
  intcode_computer computer;
  computer.instructions = instructions;
  computer.dynamic_mem = true;
  intcode_computer input_computer;
  computer.prev = &input_computer;
  std::map<long, long> maze;
  long x = 0, y = 0, path = 0;

  explore(x, y, computer, maze, path, false);
  printf("%ld\n", path);

  maze.clear();
  computer.inst_pointer = 0;
  computer.instructions = instructions;
  computer.relative_base = 0;
  x = 0, y = 0, path = 0;

  explore(x, y, computer, maze, path, true);

  long mins = 0;
  while (true) {
    std::vector<std::array<long, 2>> oxygens;
    for (long j = min_y; j <= max_y; j++) {
      for (long i = min_x; i <= max_x; i++) {
        long key = pos_to_num(i, j);
        long val = maze.count(key) == 1 ? maze[key] : UNKNOWN;
        if (val == OXYGEN) {
          oxygens.push_back({ i, j });
        }
      }
    }

    bool changed = false;
    for (std::array<long, 2>& oxy : oxygens) {
      long i = oxy[0], j = oxy[1];
      long mx = 0;
      long my = 1;
      for (long b = 1; b <= 4; b++) {
        long key = pos_to_num(i + mx, j + my);
        long val = maze.count(key) == 1 ? maze[key] : UNKNOWN;
        if (val == MOVED) {
          changed = true;
          maze[key] = OXYGEN;
        }
        long temp = mx;
        mx = my;
        my = -temp;
      }
    }
    if (!changed) break;
    mins += 1;
  }
  printf("%ld\n", mins);
  return 0;
}
