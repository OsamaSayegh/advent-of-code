#include "../lib-aoc.cpp"
#include <map>

inline size_t pos_to_num(size_t y, size_t x) {
  return y * 80 + x;
}

inline void num_to_pos(uint64_t pos, uint64_t& x, uint64_t& y) {
  x = pos % 80;
  y = pos / 80;
}

const uint64_t POSITION_LIMIT = 1 << 13;
const uint64_t KEYS_LIMIT = 1 << 26;
const uint64_t COUNT_LIMIT = 1 << 23;

uint64_t build_state(uint64_t position, uint64_t keys, uint64_t count) {
  if (position >= POSITION_LIMIT) {
    printf("Error: position is %ld, but it can't be larger than 2**13 - 1\n", position);
    exit(1);
  }
  if (keys >= KEYS_LIMIT) {
    printf("Error: keys is %ld, but it can't be larger than 2**26 - 1\n", keys);
    exit(1);
  }
  if (count >= COUNT_LIMIT) {
    printf("Error: count is %ld, but it can't be larger than 2**23 - 1\n", count);
    exit(1);
  }
  uint64_t state = 0;
  state |= keys;
  state |= position << 26;
  state |= count << 39;
  return state;
}

void extract_state(uint64_t state, uint64_t& position, uint64_t& keys, uint64_t& count) {
  keys = state & (KEYS_LIMIT - 1);
  state = state >> 26;
  position = state & (POSITION_LIMIT - 1);
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

int main(int argc, char** argv) {
  std::vector<std::vector<char>> maze;
  {
    char* content = puzzle_input(argc, argv);
    size_t content_len = strlen(content);
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
    free(content);
  }

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

  std::vector<uint64_t> queue;
  std::map<uint64_t, bool> footprints;
  queue.push_back(build_state(start_pos, 0, 0));
  while (true) {
    std::vector<uint64_t> new_queue;
    for (size_t k = 0; k < queue.size(); k++) {
      uint64_t state = queue[k];
      uint64_t position = 0, count = 0, keys = 0;
      extract_state(state, position, keys, count);
      uint64_t xx, yy;
      num_to_pos(position, xx, yy);
      for (uint64_t i = 0; i < 4; i++) {
        uint64_t x = xx, y = yy, new_keys = keys;
        if (i == 0 && y != 0) {
          y--;
        } else if (i == 1 && x < maze_width - 1) {
          x++;
        } else if (i == 2 && x != 0) {
          x--;
        } else if (i == 3 && y < maze_height - 1) {
          y++;
        } else {
          continue;
        }
        uint64_t new_position = pos_to_num(y, x);
        char location = get_maze_location(maze, new_position);
        uint64_t footprint = build_footprint(new_position, keys);
        if (location == '#') {
          continue;
        } else if (is_door(location) && !has_door_key(location, new_keys)) {
          continue;
        } else if (footprints.contains(footprint)) {
          continue;
        } else if (is_key(location)) {
          capture_key(location, new_keys);
          if (has_all_keys(new_keys, keys_count)) {
            printf("%ld\n", count + 1);
            goto p1_done;
          }
        }
        footprints[footprint] = true;
        new_queue.push_back(build_state(new_position, new_keys, count + 1));
      }
    }
    queue = new_queue;
  }

  p1_done:
  ;
  return 0;
}
