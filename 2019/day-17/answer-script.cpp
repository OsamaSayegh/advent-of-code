#include "../lib-aoc.cpp"
#include <map>
#include <algorithm>

const long
  SCAFFOLD = '#',
  NEWLINE = '\n',
  UP = '^',
  RIGHT = '>',
  DOWN = 'v',
  LEFT = '<',
  COMMA = ',',
  R = 'R',
  L = 'L',
  A = 'A',
  B = 'B',
  C = 'C',
  NONE = -1;

struct match {
  long size = -1;
  std::vector<size_t> occurrences;
};

bool operator==(const match& a, const match& b) {
  if (a.size != b.size) return false;
  if (a.occurrences.size() != b.occurrences.size()) return false;
  for (size_t i = 0; i < a.occurrences.size(); i++) {
    if (a.occurrences[i] != b.occurrences[i]) return false;
  }
  return true;
}

long pos_to_num(long x, long y) {
  return x * 1000 + y;
}

long location_type(const std::map<long, long>& map, long x, long y) {
  long key = pos_to_num(x, y);
  if (map.count(key) == 1) {
    return map.at(key);
  } else {
    return NONE;
  }
}

bool is_robot(long code) {
  return code == UP || code == RIGHT || code == DOWN || code == LEFT;
}

int main(int argc, char **argv) {
  std::vector<long> instructions = intcode_instructions(argc, argv);
  intcode_computer computer;
  computer.instructions = instructions;
  computer.dynamic_mem = true;
  intcode_computer input_computer;
  computer.prev = &input_computer;
  computer.run(false);

  std::map<long, long> map;
  long x = 0, y = 0;
  long rx = -1, ry = -1;
  std::array<long, 2> rd = { -1, -1 }; // robot direction
  long width = 0, height = 0;
  while (computer.outputs.at(computer.outputs.size() - 1) == (long) NEWLINE) {
    computer.outputs.pop_back();
  }
  for (size_t i = 0; i < computer.outputs.size(); i++) {
    long current = computer.outputs[i];
    if (current == 10) {
      x = 0;
      y += 1;
      continue;
    }
    if (is_robot(current)) {
      rx = x;
      ry = y;
      switch (current) {
        case UP:
          rd = { 0, -1 };
          break;
        case RIGHT:
          rd = { 1, 0 };
          break;
        case DOWN:
          rd = { 0, 1 };
          break;
        case LEFT:
          rd = { -1, 0 };
          break;
      }
      map[pos_to_num(x, y)] = SCAFFOLD;
    } else {
      map[pos_to_num(x, y)] = current;
    }
    if (i + 1 < computer.outputs.size()) x++;
  }
  width = x + 1, height = y + 1;

  long sum = 0;
  for (long j = 0; j < height; j++) {
    for (long i = 0; i < width; i++) {
      if (
        location_type(map, i, j) == SCAFFOLD &&
        location_type(map, i + 1, j) == SCAFFOLD &&
        location_type(map, i - 1, j) == SCAFFOLD &&
        location_type(map, i, j + 1) == SCAFFOLD &&
        location_type(map, i, j - 1) == SCAFFOLD
      ) {
        sum += i * j;
      }
    }
  }
  printf("%ld\n", sum);
  std::vector<char> turns;
  std::vector<long> moves;

  if (location_type(map, rx + rd[0], ry + rd[1]) != SCAFFOLD) {
    std::array<long, 2> new_rd = { -rd[1], rd[0] };
    if (location_type(map, rx + new_rd[0], ry + new_rd[1]) == SCAFFOLD) {
      turns.push_back(R);
      moves.push_back(0);
    } else {
      new_rd = { rd[1], -rd[0] };
      if (location_type(map, rx + new_rd[0], ry + new_rd[1]) == SCAFFOLD) {
        turns.push_back(L);
        moves.push_back(0);
      } else {
        turns.push_back(L);
        turns.push_back(L);
        moves.push_back(0);
        moves.push_back(0);
        new_rd = { new_rd[1], -new_rd[0] };
      }
    }
    rd[0] = new_rd[0];
    rd[1] = new_rd[1];
  }
  while (true) {
    long next = location_type(map, rx + rd[0], ry + rd[1]);
    if (next == SCAFFOLD) {
      rx += rd[0];
      ry += rd[1];
      if (moves.size() == 0) {
        moves.push_back(1);
      } else {
        moves[moves.size() - 1] += 1;
      }
    } else {
      std::array<long, 2> new_rd = { -rd[1], rd[0] };
      if (location_type(map, rx + new_rd[0], ry + new_rd[1]) == SCAFFOLD) {
        turns.push_back(R);
        moves.push_back(0);
        rd[0] = new_rd[0];
        rd[1] = new_rd[1];
        continue;
      }
      new_rd = { rd[1], -rd[0] };
      if (location_type(map, rx + new_rd[0], ry + new_rd[1]) == SCAFFOLD) {
        turns.push_back(L);
        moves.push_back(0);
        rd[0] = new_rd[0];
        rd[1] = new_rd[1];
        continue;
      }
      break;
    }
  }
  if (moves.size() != turns.size()) {
    printf("Error: expected moves count (%ld) to equal turns count (%ld)\n", moves.size(), turns.size());
    exit(1);
  }

  std::vector<match> matches;
  for (size_t l = 4; l >= 2; l--) {
    for (size_t i = 0; i < moves.size() - l; i++) {
      size_t a = i;
      size_t b = i + l;
      match m;
      m.size = l;
      m.occurrences.push_back(i);
      for (size_t j = b; j <= moves.size() - l; j++) {
        bool has_match = true;
        for (size_t g = 0; g < l; g++) {
          if ((moves[a + g] != moves[j + g]) || (turns[a + g] != turns[j + g])) {
            has_match = false;
            break;
          }
        }
        if (has_match) m.occurrences.push_back(j);
      }
      if (m.occurrences.size() > 1) matches.push_back(m);
    }
  }
  std::vector<char> main_routine;
  match a, b, c;
  for (size_t i = 0; i < matches.size(); i++) {
    a = matches[i];
    for (size_t j = 0; j < matches.size(); j++) {
      if (j == i) continue;
      b = matches[j];
      for (size_t k = 0; k < matches.size(); k++) {
        if (k == j || k == i) continue;
        c = matches[k];
        size_t list_size = 0;
        while (list_size < moves.size()) {
          bool found = false;
          for (match m : { a, b, c }) {
            for (size_t r : m.occurrences) {
              if (r == list_size) {
                list_size += m.size;
                found = true;
                if (m == a) {
                  main_routine.push_back(A);
                } else if (m == b) {
                  main_routine.push_back(B);
                } else {
                  main_routine.push_back(C);
                }
                break;
              }
            }
            if (found) break;
          }
          if (!found) break;
        }
        if (list_size == moves.size() && main_routine.size() * 2 - 1 <= 20) {
          goto pattern_found;
        } else {
          main_routine.clear();
        }
      }
    }
  }

  printf("Error: no pattern found\n");
  exit(1);
  pattern_found: ;

  computer.instructions = instructions;
  computer.instructions[0] = 2;
  computer.outputs.clear();
  computer.relative_base = 0;
  computer.halted = false;
  computer.inst_pointer = 0;
  input_computer.outputs.clear();

  for (size_t r = 0; r < main_routine.size(); r++) {
    input_computer.outputs.push_back(main_routine[r]);
    if (r + 1 < main_routine.size()) input_computer.outputs.push_back(COMMA);
  }
  input_computer.outputs.push_back(NEWLINE);

  for (match m : { a, b, c }) {
    long first = m.occurrences[0];
    for (long x = 0; x < m.size; x++) {
      input_computer.outputs.push_back(turns[first + x]);
      input_computer.outputs.push_back(COMMA);
      long num = moves[first + x];
      std::vector<char> nums;
      while (num > 0) {
        nums.push_back((char) (num % 10 + 48));
        num /= 10;
      }
      std::reverse(nums.begin(), nums.end());
      for (char n : nums) input_computer.outputs.push_back(n);
      if (x + 1 < m.size) input_computer.outputs.push_back(COMMA);
    }
    input_computer.outputs.push_back(NEWLINE);
  }
  input_computer.outputs.push_back('n');
  input_computer.outputs.push_back(NEWLINE);

  computer.run(false);
  printf("%ld\n", computer.outputs[computer.outputs.size() - 1]);
  return 0;
}
