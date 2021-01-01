# frozen_string_literal: true

# content = <<~DATA
# DATA
content = File.read(File.expand_path('input.txt', __dir__))
cups = content.strip.split('').map(&:to_i)
def play(cups, moves)
  size = cups.size
  current = cups[0]
  after_hash = {}
  cups.each_with_index do |c, i|
    after_hash[c] = cups[(i + 1) % size]
  end
  moves.times do
    a1 = after_hash[current]
    a2 = after_hash[a1]
    a3 = after_hash[a2]
    dest = current - 1
    dest = size if dest < 1
    while dest == a1 || dest == a2 || dest == a3
      dest -= 1
      dest = size if dest < 1
    end
    after_dest = after_hash[dest]
    after_hash[dest] = a1
    after_hash[current] = after_hash[a3]
    after_hash[a3] = after_dest
    current = after_hash[current]
  end
  a = after_hash[1]
  final = []
  final[0] = 1
  while a != 1
    final << a
    a = after_hash[a]
  end
  final
end

puts play(cups, 100)[1..-1].join('')
million_cups = Array.new(1_000_000 + 1, &:itself)
million_cups[0..cups.size] = cups
puts play(million_cups, 10_000_000).yield_self { |res| res[1] * res[2] }
