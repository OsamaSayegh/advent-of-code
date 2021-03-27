#include "../lib-aoc.cpp"
#include <map>

struct chemical {
  std::string name;
  long quantity = 0;
};

struct reaction {
  chemical output;
  std::vector<chemical> inputs;
};

std::map<std::string, reaction> map;

void ore_rec(
  const std::string& name,
  long quantity,
  std::map<std::string, long>& produced,
  std::map<std::string, long>& consumed
) {
  const reaction& r = map.at(name);
  long min_quantity = r.output.quantity;
  long multiplier = 1;
  long available = 0;
  if (produced.count(name) == 1) {
    available = produced[name];
    if (consumed.count(name) == 1) available -= consumed[name];
  }
  if (quantity < available) {
    return;
  }
  if (min_quantity < (quantity - available)) {
    multiplier = (quantity - available) / min_quantity;
    if ((quantity - available) % min_quantity != 0) multiplier++;
  }
  for (size_t i = 1; i <= r.inputs.size(); i++) {
    const chemical& c = r.inputs[i - 1];
    if (c.name == "ORE") {
      if (produced.count("ORE") == 1) {
        produced["ORE"] += multiplier * c.quantity;
      } else {
        produced["ORE"] = multiplier * c.quantity;
      }
    } else {
      ore_rec(c.name, multiplier * c.quantity, produced, consumed);
      if (consumed.count(c.name) == 1) {
        consumed[c.name] += multiplier * c.quantity;
      } else {
        consumed[c.name] = multiplier * c.quantity;
      }
    }
  }
  if (produced.count(name) == 1) {
    produced[name] += multiplier * r.output.quantity;
  } else {
    produced[name] = multiplier * r.output.quantity;
  }
}

long ore_needed_for_fuel(long fuel) {
  std::map<std::string, long> produced;
  std::map<std::string, long> consumed;
  ore_rec("FUEL", fuel, produced, consumed);
  return produced["ORE"];
}

int main(int argc, char **argv) {
  char *input = puzzle_input(argc, argv);
  const std::vector<std::string> lines = split(input, '\n');
  free(input);
  for (size_t i = 0; i < lines.size(); i++) {
    const std::string& current = lines[i];
    if (current.size() == 0) continue;
    size_t prev = 0;
    size_t found = current.find_first_of(",");
    if (found == std::string::npos) {
      found = current.find_first_of("=");
    }
    std::vector<chemical> inputs;
    do {
      long quantity;
      char name[found - prev];
      std::string chunk = current.substr(prev, found - prev);
      if (sscanf(chunk.c_str(), "%ld %s", &quantity, name) == 2) {
        chemical c;
        c.name = std::string(name);
        c.quantity = quantity;
        inputs.push_back(c);
      } else {
        printf("Unexpected error when parsing input chemicals out of line: %s.\n", current.c_str());
        exit(1);
      }
      prev = found + 2;
      found = current.find_first_of(",", prev);
      if (found == std::string::npos) {
        found = current.find_first_of("=", prev);
      }
    } while (found != std::string::npos);
    long quantity = 0;
    char name[current.size() - prev];
    std::string chunk = current.substr(prev, current.size() - prev);
    if (sscanf(chunk.c_str(), "%ld %s", &quantity, name) == 2) {
      chemical output;
      output.name = std::string(name);
      output.quantity = quantity;
      reaction r;
      r.output = output;
      r.inputs = inputs;
      map.insert(std::pair<std::string, reaction>(r.output.name, r));
    } else {
      printf("Unexpected error when parsing output chemical out of line: %s.\n", current.c_str());
      exit(1);
    }
  }
  long one_fuel = ore_needed_for_fuel(1);
  printf("%ld\n", one_fuel);
  long tril = 1'000'000'000'000;
  long min = tril / one_fuel;
  long max = min * 2;
  while (ore_needed_for_fuel(max) < tril) max *= 2;
  long mid = (min + max) / 2;
  while (min < max - 1) {
    long needed = ore_needed_for_fuel(mid);
    if (needed < tril) {
      min = mid;
    } else {
      max = mid;
    }
    mid = (min + max) / 2;
  }
  printf("%ld\n", mid);
  return 0;
}
