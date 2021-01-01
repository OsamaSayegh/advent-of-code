# frozen_string_literal: true

require 'set'
# content = <<~DATA
# Player 1:
# 9
# 2
# 6
# 3
# 1
# 
# Player 2:
# 5
# 8
# 4
# 7
# 10
# DATA
content = File.read(File.expand_path('input.txt', __dir__))
p1, p2 = content.split("\n\n").map { |lines| lines.split("\n")[1..-1].map(&:to_i) }
while !p1.empty? && !p2.empty?
  c1 = p1.shift
  c2 = p2.shift
  if c1 > c2
    p1 << c1
    p1 << c2
  else
    p2 << c2
    p2 << c1
  end
end

winner = p1.empty? ? p2 : p1
puts winner.reverse.map.with_index { |n, i| n * (i + 1) }.sum

p1, p2 = content.split("\n\n").map { |lines| lines.split("\n")[1..-1].map(&:to_i) }
def play_game(p1, p2)
  rounds = Set.new
  r = 1
  while !p1.empty? && !p2.empty?
    hash = [p1, p2].hash
    if rounds.include?(hash)
      return 'p1'
    else
      rounds << hash
    end
    c1 = p1.shift
    c2 = p2.shift
    if p1.size >= c1 && p2.size >= c2
      if play_game(p1[0...c1], p2[0...c2]) == 'p1'
        p1 << c1
        p1 << c2
      else
        p2 << c2
        p2 << c1
      end
    else
      if c1 > c2
        p1 << c1
        p1 << c2
      else
        p2 << c2
        p2 << c1
      end
    end
    r += 1
  end
  p1.empty? ? 'p2' : 'p1'
end

winner = play_game(p1, p2) == 'p1' ? p1 : p2
puts winner.reverse.map.with_index { |n, i| n * (i + 1) }.sum
