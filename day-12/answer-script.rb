# frozen_string_literal: true

# content = <<~DATA
# F10
# N3
# F7
# R90
# F11
# DATA
content = File.read(File.expand_path('input.txt', __dir__))
ACTIONS = content.split("\n").map { |l| [l[0], l[1..-1].to_i] }

MAP = {
  'E' => 0,
  'S' => 0,
  'W' => 0,
  'N' => 0
}

OPPSITES = {
  'N' => 'S',
  'S' => 'N',
  'W' => 'E',
  'E' => 'W'
}

def move(map, dir, size)
  oppsite = OPPSITES[dir]
  sum = map[oppsite] - size
  if sum > 0
    map[oppsite] = sum
  elsif sum < 0
    map[oppsite] = 0
    map[dir] += sum.abs
  else
    map[oppsite] = 0
  end
end

current = 'E'
ACTIONS.each do |action, int|
  if action == 'F'
    move(MAP, current, int)
  end
  if MAP.key?(action)
    move(MAP, action, int)
  end
  if action == 'R' || action == 'L'
    steps = int / 90
    steps *= -1 if action == 'L'
    dirs = MAP.keys
    current = dirs[(dirs.index(current) + steps) % dirs.size]
  end
end

puts MAP.values.sum

MAP.dup.each { |k, _| MAP[k] = 0 }

MAP_W = {
  'E' => 10,
  'S' => 0,
  'W' => 0,
  'N' => 1
}

def rotate(deg)
  dup = MAP_W.dup
  keys = dup.keys
  keys.each_with_index do |k, ind|
    MAP_W[keys[(ind + (deg / 90)) % 4]] = dup[k]
  end
end

ACTIONS.each do |action, int|
  if action == 'F'
    MAP_W.each do |d, v|
      move(MAP, d, v * int)
    end
  end
  if MAP_W.key?(action)
    move(MAP_W, action, int)
  end
  if action == 'R' || action == 'L'
    int *= -1 if action == 'L'
    rotate(int)
  end
end

puts MAP.values.sum
