#include "../lib-aoc.cpp"

int main(int argc, char **argv) {
  const char *content = puzzle_input(argc, argv);
  std::vector<int> digits;
  for (int i = 0; i < strlen(content); i++) {
    if (content[i] == '\n') continue;
    digits.push_back(std::stoi(std::string(&content[i], 1)));
  }
  std::vector<std::vector<std::vector<int>>> layers;
  int min_zeros = 0x7fffffff;
  int min_zeros_ones_by_tows = 0;
  for (int i = 0; i < digits.size(); i += 25 * 6) {
    std::vector<std::vector<int>> layer;
    int zeros = 0;
    int ones = 0;
    int twos = 0;
    for (int j = 0; j < 25 * 6; j += 25) {
      std::vector<int> row;
      for (int k = 0; k < 25; k++) {
        int digit = digits[i + j + k];
        row.push_back(digit);
        if (digit == 0) zeros++;
        if (digit == 1) ones++;
        if (digit == 2) twos++;
      }
      layer.push_back(row);
    }
    if (zeros < min_zeros) {
      min_zeros = zeros;
      min_zeros_ones_by_tows = ones * twos;
    }
    layers.push_back(layer);
  }
  printf("%d\n", min_zeros_ones_by_tows);
  for (int j = 0; j < 6; j++) {
    for (int i = 0; i < 25; i++) {
      int final_color = 2;
      for (int k = 0; k < layers.size(); k++) {
        if (layers[k][j][i] != final_color) {
          final_color = layers[k][j][i];
          break;
        }
      }
      if (final_color == 1) {
        printf("O");
      } else {
        printf(" ");
      }
    }
    printf("\n");
  }
  return 0;
}
