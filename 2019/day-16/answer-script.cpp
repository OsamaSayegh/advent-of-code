#include "../lib-aoc.cpp"
#include <array>
#include <algorithm>

std::array<int, 4> base = { 0, 1, 0, -1 };
int pattern_number(size_t repeat_count, size_t index) {
  size_t full_size = 4 * repeat_count;
  index %= full_size;
  return base[index / repeat_count];
}

int extract_result(const std::vector<int>& numbers) {
  int out = 0;
  int m = 10'000'000;
  for (int o = 0; o < 8; o++) {
    out += m * numbers[o];
    m /= 10;
  }
  return out;
}

int fft(std::vector<int> numbers) {
  for (int c = 1; c <= 100; c++) {
    std::vector<int> output;
    for (size_t i = 0; i < numbers.size(); i++) {
      size_t base_repeat = i + 1;
      int new_val = 0;
      for (size_t n = 0; n < numbers.size(); n++) {
        int p = pattern_number(base_repeat, n + 1);
        if (p == 0) {
          if (pattern_number(base_repeat, n) == 0) n--;
          n += base_repeat - 1;
          continue;
        }
        new_val += numbers[n] * p;
      }
      output.push_back(abs(new_val) % 10);
    }
    numbers = output;
  }
  return extract_result(numbers);
}

void repeat_vector(std::vector<int>& vec, int repeat) {
  size_t orig_size = vec.size();
  vec.reserve(orig_size * repeat);
  for (int r = 1; r <= repeat - 1; r++) {
    std::copy_n(vec.begin(), orig_size, std::back_inserter(vec));
  }
}

int lazy_fft(std::vector<int> numbers, int offset, int repeat) {
  repeat_vector(numbers, repeat);
  size_t start = offset;
  for (int p = 1; p <= 100; p++) {
    std::vector<int> output;
    output.push_back(numbers[numbers.size() - 1]);
    int output_index = 0;
    for (size_t i = numbers.size() - 1; i > start; i--) {
      size_t ii = i - 1;
      int res = (numbers[ii] + output[output_index++]) % 10;
      output.push_back(res);
    }
    numbers = output;
    std::reverse(numbers.begin(), numbers.end());
    start = 0;
  }
  return extract_result(numbers);
}

int main(int argc, char **argv) {
  char *content = puzzle_input(argc, argv);
  size_t content_len = strlen(content);
  std::vector<int> numbers;
  for (size_t i = 0; i < content_len; i++) {
    if (content[i] == '\n') break;
    numbers.push_back(content[i] - '0');
  }
  free(content);

  printf("%d\n", fft(numbers));

  int m = 1'000'000;
  int offset = 0;
  for (int k = 0; k < 7; k++) {
    offset += m * numbers[k];
    m /= 10;
  }
  printf("%d\n", lazy_fft(numbers, offset, 10'000));

  return 0;
}
