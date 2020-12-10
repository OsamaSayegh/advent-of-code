# frozen_string_literal: true

# content = <<~DATA
# 28
# 33
# 18
# 42
# 31
# 14
# 46
# 20
# 48
# 47
# 24
# 23
# 49
# 45
# 19
# 38
# 39
# 11
# 1
# 32
# 25
# 35
# 8
# 17
# 7
# 9
# 4
# 2
# 34
# 10
# 3
# DATA
content = File.read(File.expand_path('input.txt', __dir__))
ADAPTERS = content.split("\n").map(&:to_i)
ADAPTERS << 0
ADAPTERS << ADAPTERS.max + 3
ADAPTERS.sort!

diff_1 = 0
diff_3 = 0

ADAPTERS.each_with_index do |i, ind|
  next if ind == ADAPTERS.size - 1
  if ADAPTERS[ind + 1] - i == 1
    diff_1 += 1
  end
  if ADAPTERS[ind + 1] - i == 3
    diff_3 += 1
  end
end

puts diff_1 * diff_3

PATHS = {}
CACHE = {}

ADAPTERS.each_with_index do |i, ind|
  sub = ADAPTERS[ind + 1..ind + 3]
  next unless sub
  PATHS[i] = sub.select { |n| n - i <= 3 }
end

def paths(n)
  return CACHE[n] if CACHE[n]
  arr = PATHS[n]
  res = arr.map { |s| paths(s) }.sum
  res = 1 if res == 0
  CACHE[n] = res
  res
end

puts paths(0)
