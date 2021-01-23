#include <stdio.h>
#include <stdlib.h>
#include <vector>
#include <string>
#include <string.h>

char* puzzle_input(int argc, char** argv) {
  FILE* input = fopen(argc > 1 ? argv[1] : "input.txt", "rb");
  if (input == NULL) {
    printf("No such file %s\n", argv[1]);
    exit(1);
  }
  fseek(input, 0, SEEK_END);
  size_t size = ftell(input);
  char* content = (char*) malloc(size + 1);
  fseek(input, 0, SEEK_SET);
  fread(content, 1, size, input);
  content[size] = '\0';
  return content;
}

std::vector<std::string> split(const char* input, char del) {
  int len = strlen(input);
  size_t start = 0;
  std::vector<std::string> res;
  for (int i = 0; i < len; i++) {
    if (input[i] == del) {
      res.push_back(std::string(input + start, i - start));
      start = i + 1;
    }
  }
  res.push_back(std::string(input + start, len - start));
  return res;
}
