#include "../lib-aoc.cpp"
#include <queue>
#include <map>

inline size_t pos_to_num(size_t y, size_t x) {
  return y * 80 + x;
}

inline void num_to_pos(uint64_t pos, uint64_t& x, uint64_t& y) {
  x = pos % 80;
  pos -= x;
  y = pos / 80;
}

uint64_t build_state(uint64_t position, uint64_t keys, uint64_t dir, uint64_t count) {
  if (position >= (1 << 13)) {
    printf("Error: position is %ld, but it can't be larger than 2**13 - 1\n", position);
    exit(1);
  }
  if (keys >= (1 << 26)) {
    printf("Error: keys is %ld, but it can't be larger than 2**26 - 1\n", keys);
    exit(1);
  }
  if (dir >= (1 << 2)) {
    printf("Error: dir is %ld, but it can't be larger than 2**2 - 1\n", dir);
    exit(1);
  }
  if (count >= (1 << 23)) {
    printf("Error: count is %ld, but it can't be larger than 2**23 - 1\n", count);
    exit(1);
  }
  uint64_t state = 0;
  state |= keys;
  state |= dir << 26;
  state |= position << 28;
  state |= count << 41;
  return state;
}

void extract_state(uint64_t state, uint64_t& position, uint64_t& keys, uint64_t& dir, uint64_t& count) {
  keys = state & ((1 << 26) - 1);
  state = state >> 26;
  dir = state & ((1 << 2) - 1);
  state = state >> 2;
  position = state & ((1 << 13) - 1);
  state = state >> 13;
  count = state;
}

uint64_t build_footprint(uint64_t position, uint64_t keys) {
  uint64_t footprint = 0;
  footprint |= keys;
  footprint |= position << 26;
  return footprint;
}

inline char get_maze_location(const std::vector<std::vector<char>>& maze, uint64_t position) {
  uint64_t x, y;
  num_to_pos(position, x, y);
  return maze.at(y).at(x);
}

inline bool is_key(char location) {
  return location >= 97 && location <= 122;
}

inline bool is_door(char location) {
  return location >= 65 && location <= 90;
}

inline void capture_key(char key, uint64_t& keys) {
  if (key < 97 || key > 122) {
    printf("Error: capture_key received %c which can't be a door key.\n", key);
    exit(1);
  }
  keys |= 1 << (((uint64_t) key) - 97);
}

inline bool has_door_key(char door, uint64_t keys) {
  if (door < 65 || door > 90) {
    printf("Error: has_door_key received %c which can't be a door.\n", door);
    exit(1);
  }
  return (keys & (1 << (((uint64_t) door) - 65))) != 0;
}

inline bool has_all_keys(uint64_t keys, uint64_t keys_count) {
  uint64_t mask = (1 << keys_count) - 1;
  return keys == mask;
}

inline uint64_t caputred_keys_count(uint64_t keys) {
  uint64_t count = 0;
  for (uint64_t i = 0; i < 26; i++) {
    if ((keys & (1 << i)) != 0) {
      count++;
    }
  }
  return count;
}

inline bool __has_to_go_back(const std::vector<std::vector<char>>& maze, uint64_t position, uint64_t keys) {
  uint64_t x, y;
  num_to_pos(position, x, y);
  int count = 0;
  char bottom = get_maze_location(maze, pos_to_num(y + 1, x));
  char top = get_maze_location(maze, pos_to_num(y - 1, x));
  char left = get_maze_location(maze, pos_to_num(y, x + 1));
  char right = get_maze_location(maze, pos_to_num(y, x - 1));

  if (bottom == '#') count++;
  if (top == '#') count++;
  if (left == '#') count++;
  if (right == '#') count++;

  if (is_door(bottom) && !has_door_key(bottom, keys)) count++;
  if (is_door(top) && !has_door_key(top, keys)) count++;
  if (is_door(left) && !has_door_key(left, keys)) count++;
  if (is_door(right) && !has_door_key(right, keys)) count++;
  return count >= 3;
}

uint64_t captured_keys_count(uint64_t keys) {
  uint64_t count = 0;
  for (uint64_t i = 0; i < 26; i++) {
    if ((keys & (1 << i)) != 0) {
      count++;
    }
  }
  return count;
}

bool _compare_items(uint64_t a, uint64_t b) {
  uint64_t a_keys, b_keys, a_count, b_count, _null;
  extract_state(a, _null, a_keys, _null, a_count);
  extract_state(b, _null, b_keys, _null, b_count);
  return captured_keys_count(a_keys) < captured_keys_count(b_keys);
  // return a_count > b_count;
  // return captured_keys_count(a_keys) < captured_keys_count(b_keys);// && a_count > b_count;
  // if (b_count < a_count) {
  // } else {
  //   return true;
  // }
}

struct priority_queue {
  std::vector<uint64_t> _vector;

  void add(uint64_t item) {
    _vector.push_back(item);
    sort(_vector.begin(), _vector.end(), _compare_items);
  }

  uint64_t pop() {
    uint64_t last = _vector.at(_vector.size() - 1);
    _vector.pop_back();
    return last;
  }

  size_t size() {
    return _vector.size();
  }
};

int main(int argc, char **argv) {
  char *content = puzzle_input(argc, argv);
  size_t content_len = strlen(content);
  std::vector<std::vector<char>> maze;
  {
    std::vector<char> row;
    for (size_t t = 0; t < content_len; t++) {
      char current = content[t];
      if (current == '\n') {
        maze.push_back(row);
        row = std::vector<char> { };
      } else {
        row.push_back(current);
      }
    }
  }

  // priority_queue q;
  // q.add(build_state(0, 7, 0, 0));
  // q.add(build_state(0, 65, 0, 0));
  // q.add(build_state(0, 32, 0, 0));
  // printf("size=%ld\n", q.size());
  // while (q.size() > 0) {
  //   uint64_t keys, _null;
  //   extract_state(q.pop(), _null, keys, _null, _null);
  //   printf("item: %ld\n", captured_keys_count(keys));
  // }
  // return 0;

  size_t maze_width = maze.at(0).size();
  size_t maze_height = maze.size();
  size_t start_pos = 0;
  size_t keys_count = 0;
  for (size_t i = 0; i < maze_height; i++) {
    for (size_t j = 0; j < maze_width; j++) {
      if (maze[i][j] == '@') {
        start_pos = pos_to_num(i, j);
      } else if (is_key(maze[i][j])) {
        keys_count++;
      }
    }
  }

  std::queue<uint64_t> queue;
  // priority_queue queue;
  std::map<uint64_t, bool> footprints;
  queue.push(build_state(start_pos, 0, 0, 0));
  // queue.add(build_state(start_pos, 0, 0, 0));
  while (true) {
    std::queue<uint64_t> new_queue;
    while (queue.size() > 0) {
      uint64_t state = queue.front();
      queue.pop();
      // printf("queue=%ld\n", queue.size());
      // uint64_t state = queue.pop();
      uint64_t position = 0, count = 0, dir = 0, keys = 0;
      extract_state(state, position, keys, dir, count);
      char location = get_maze_location(maze, position);
      if (location == '#') {
        continue;
      } else if (is_door(location) && !has_door_key(location, keys)) {
        continue;
      } else if (is_key(location)) {
        capture_key(location, keys);
        if (has_all_keys(keys, keys_count)) {
          printf("%ld\n", count);
          goto done;
          // break;
        }
      }
      footprints[build_footprint(position, keys)] = true;
      count++;
      for (uint64_t i = 0; i < 4; i++) {
        uint64_t x, y;
        num_to_pos(position, x, y);
        if (i == 0 && y != 0) {
          // if (y == 0) continue;
          y--;
        } else if (i == 1 && x < maze_width - 1) {
          // if (x >= maze_width - 1) continue;
          x++;
        } else if (i == 2 && x != 0) {
          // if (x == 0) continue;
          x--;
        } else if (i == 3 && y < maze_height - 1) {
          // if (y >= maze_height - 1) continue;
          y++;
        } else {
          continue;
        }
        uint64_t new_position = pos_to_num(y, x);
        if (!footprints.contains(build_footprint(new_position, keys)) && get_maze_location(maze, new_position) != '#') {
          // queue.push(build_state(new_position, keys, dir, count));
          new_queue.push(build_state(new_position, keys, dir, count));
          // queue.add(build_state(new_position, keys, dir, count));
        }
      }
    }
    queue = new_queue;
  }

  done:;
  return 0;
}
