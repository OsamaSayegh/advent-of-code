#include "../lib-aoc.cpp"

int run(int noun, int verb, std::vector<int> nums) {
  nums[1] = noun;
  nums[2] = verb;
  for (int i = 0; i < nums.size();) {
    int op = nums[i];
    if (op == 99) {
      break;
    } else if (op == 1) {
      nums[nums[i + 3]] = nums[nums[i + 1]] + nums[nums[i + 2]];
    } else if (op == 2) {
      nums[nums[i + 3]] = nums[nums[i + 1]] * nums[nums[i + 2]];
    }
    i += 4;
  }
  return nums[0];
}

int main(int argc, char** argv) {
  char* content = puzzle_input(argc, argv);
  auto s_nums = split(content, ',');
  std::vector<int> nums;
  for (int i = 0; i < s_nums.size(); i++) {
    nums.push_back(std::stoi(s_nums[i]));
  }
  printf("%d\n", run(12, 2, nums));

  for (int j = 0; j <= 99; j++) {
    for (int k = 0; k <= 99; k++) {
      if (run(j, k, nums) == 19690720) {
        printf("%d\n", 100 * j + k);
        goto end;
      }
    }
  }

  end: ;
  return 0;
}
