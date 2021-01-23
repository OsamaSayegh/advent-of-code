#include "../lib-aoc.cpp"
#include "map"

std::map<std::string, int> cache;
int count_orbits(std::string object, std::map<std::string, std::string> orbits) {
  if (object == "COM") {
    return 0;
  }
  if (cache.find(object) != cache.end()) {
    return cache.find(object)->second;
  } else {
    int count = 1 + count_orbits(orbits.find(object)->second, orbits);
    cache.insert(std::pair<std::string, int>(object, count));
    return count;
  }
}

int main(int argc, char** argv) {
  char* content = puzzle_input(argc, argv);
  std::vector<std::string> orbits = split(content, '\n');
  std::map<std::string, std::string> map;
  for (int i = 0; i < orbits.size(); i++) {
    std::vector<std::string> objects = split(orbits[i].c_str(), ')');
    if (objects.size() != 2) {
      printf("Expected 2 objects in %s but instead got %d\n", orbits[i].c_str(), (int) objects.size());
      exit(1);
    }
    map.insert(std::pair<std::string, std::string>(objects[1], objects[0]));
  }

  int total = 0;
  for (auto it = map.begin(); it != map.end(); it++) {
    total += count_orbits(it->first, map);
  }
  std::vector<std::string> you_to_com;
  std::string src = std::string("YOU");
  while(src != "COM") {
    src = map.find(src)->second;
    you_to_com.push_back(src);
  }
  std::vector<std::string> san_to_com;
  src = std::string("SAN");
  while(src != "COM") {
    src = map.find(src)->second;
    san_to_com.push_back(src);
  }

  int min = 0x7fffffff;
  int you_steps = 0;
  for (int i = 0; i < you_to_com.size(); i++) {
    int san_steps = 0;
    for (int j = 0; j < san_to_com.size(); j++) {
      if (you_to_com[i] == san_to_com[j] && min > you_steps + san_steps) {
        min = you_steps + san_steps;
      }
      san_steps++;
    }
    you_steps++;
  }
  printf("%d\n", total);
  printf("%d\n", min);
  return 0;
}
