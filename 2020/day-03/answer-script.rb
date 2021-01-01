# frozen_string_literal: true

# content = <<~DATA
# ..##.......
# #...#...#..
# .#....#..#.
# ..#.#...#.#
# .#...##..#.
# ..#.##.....
# .#.#.#....#
# .#........#
# #.##...#...
# #...##....#
# .#..#...#.#
# DATA
content = File.read(File.expand_path('input.txt', __dir__))
MATRIX = content.split("\n").map { |l| l.split('') }

TREE = '#'
ROWS = MATRIX.size
COLUMNS = MATRIX[0].size

def count_trees(x, y)
  i = j = 0
  trees_count = 0
  loop do
    trees_count += 1 if MATRIX[i][j] == TREE
    j += x
    i += y
    j -= COLUMNS if COLUMNS <= j
    break if ROWS <= i
  end
  trees_count
end

puts count_trees(3, 1)
puts count_trees(1, 1) * count_trees(3, 1) * count_trees(5, 1) * count_trees(7, 1) * count_trees(1, 2)
