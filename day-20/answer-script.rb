# frozen_string_literal: true

# content = <<~DATA
# Tile 2311:
# ..##.#..#.
# ##..#.....
# #...##..#.
# ####.#...#
# ##.##.###.
# ##...#.###
# .#.#.#..##
# ..#....#..
# ###...#.#.
# ..###..###
# 
# Tile 1951:
# #.##...##.
# #.####...#
# .....#..##
# #...######
# .##.#....#
# .###.#####
# ###.##.##.
# .###....#.
# ..#.#..#.#
# #...##.#..
# 
# Tile 1171:
# ####...##.
# #..##.#..#
# ##.#..#.#.
# .###.####.
# ..###.####
# .##....##.
# .#...####.
# #.##.####.
# ####..#...
# .....##...
# 
# Tile 1427:
# ###.##.#..
# .#..#.##..
# .#.##.#..#
# #.#.#.##.#
# ....#...##
# ...##..##.
# ...#.#####
# .#.####.#.
# ..#..###.#
# ..##.#..#.
# 
# Tile 1489:
# ##.#.#....
# ..##...#..
# .##..##...
# ..#...#...
# #####...#.
# #..#.#.#.#
# ...#.#.#..
# ##.#...##.
# ..##.##.##
# ###.##.#..
# 
# Tile 2473:
# #....####.
# #..#.##...
# #.##..#...
# ######.#.#
# .#...#.#.#
# .#########
# .###.#..#.
# ########.#
# ##...##.#.
# ..###.#.#.
# 
# Tile 2971:
# ..#.#....#
# #...###...
# #.#.###...
# ##.##..#..
# .#####..##
# .#..####.#
# #..#.#..#.
# ..####.###
# ..#.#.###.
# ...#.#.#.#
# 
# Tile 2729:
# ...#.#.#.#
# ####.#....
# ..#.#.....
# ....#..#.#
# .##..##.#.
# .#.####...
# ####.#.#..
# ##.####...
# ##..#.##..
# #.##...##.
# 
# Tile 3079:
# #.#.#####.
# .#..######
# ..#.......
# ######....
# ####.#..#.
# .#...#.##.
# #.#####.##
# ..#.###...
# ..#.......
# ..#.###...
# DATA
content = File.read(File.expand_path('input.txt', __dir__))
tiles = {}
content.split("\n\n").each do |tile|
  lines = tile.split("\n")
  id = Integer(lines[0].split(' ')[1].gsub(':', ''))
  raise "saw id #{id} before" if tiles.key?(id)
  image = lines[1..-1].map { |l| l.split('') }
  tiles[id] = image
end

# 0 => right
# 1 => bottom
# 2 => left
# 3 => top
def get_edge(no, tile)
  case no
  when 0
    tile.map { |r| r[r.size - 1] }
  when 1
    tile[tile.size - 1]
  when 2
    tile.map { |r| r[0] }
  when 3
    tile[0]
  end
end

def rotate(tile, degrees)
  tile = tile
  (degrees / 90).times do
    tile = tile.map.with_index { |_, i| tile.map { |t| t[i] }.reverse }
  end
  tile
end

def flip_h(tile)
  tile.map(&:reverse)
end

def flip_v(tile)
  tile.reverse
end

find_matching_tile = -> (tile_id, edge_index) do
  edge = get_edge(edge_index, tiles[tile_id])
  tile_ids = tiles.keys - [tile_id]
  i = 0
  current_edge_index = 0
  reversed = false
  found = false
  while i < tile_ids.size
    current_tile = tiles[tile_ids[i]]
    current_edge_index = 0
    while current_edge_index <= 3
      current_edge = get_edge(current_edge_index, current_tile)
      if current_edge == edge
        found = true
        break
      elsif current_edge.reverse == edge
        reversed = true
        found = true
        break
      end
      current_edge_index += 1
    end
    break if found
    i += 1
  end
  if found
    { id: tile_ids[i], edge: current_edge_index, reversed: reversed }
  else
    nil
  end
end

corners = tiles.select do |id, tile|
  edges = []
  matching_edges = (0..3).count do |n|
    edge = get_edge(n, tile)
    tiles.all? do |sub_id, sub_tile|
      next true if id == sub_id
      (0..3).none? do |m|
        sub_edge = get_edge(m, sub_tile)
        if sub_edge == edge
          edges << n
          true
        elsif sub_edge.reverse == edge
          edges << n
          true
        else
          false
        end
      end
    end
  end
  matching_edges == 2
end
puts corners.keys.inject(:*)

my_corner_id = corners.keys[0]
edges = (0..3).reject { |i| find_matching_tile.call(my_corner_id, i) }.sort
if edges == [0, 1]
  tiles[my_corner_id] = flip_h(flip_v(tiles[my_corner_id]))
elsif edges == [1, 2]
  tiles[my_corner_id] = flip_v(tiles[my_corner_id])
elsif edges == [0, 3]
  tiles[my_corner_id] = flip_h(tiles[my_corner_id])
else
  raise ':(' unless edges == [2, 3]
end

square_size = Integer(Math.sqrt(tiles.size))

arranged = Array.new(square_size)
arranged.map! { Array.new(square_size) }
arranged[0][0] = tiles[my_corner_id]

arranged_ids = Array.new(square_size)
arranged_ids.map! { Array.new(square_size) }
arranged_ids[0][0] = my_corner_id

i = 0
while i < square_size
  j = 0
  while j < square_size
    if arranged_ids[i][j]
      j += 1
      next
    end
    top = arranged_ids[i - 1][j]
    left = arranged_ids[i][j - 1]
    if top
      output = find_matching_tile.call(top, 1)
      raise 'no top match' unless output
      rotate_degress = (360 + (3 - output[:edge]) * 90) % 360
      reversed = [90, 180].include?(rotate_degress) ? !output[:reversed] : output[:reversed]
      new_tile = rotate(tiles[output[:id]], rotate_degress)
      new_tile = flip_h(new_tile) if reversed
      tiles[output[:id]] = new_tile
      arranged[i][j] = new_tile
      arranged_ids[i][j] = output[:id]
    elsif left
      output = find_matching_tile.call(left, 0)
      raise 'no left match' unless output
      rotate_degress = (360 + (2 - output[:edge]) * 90) % 360
      reversed = rotate_degress < 180 ? output[:reversed] : !output[:reversed]
      new_tile = rotate(tiles[output[:id]], rotate_degress)
      new_tile = flip_v(new_tile) if reversed
      tiles[output[:id]] = new_tile
      arranged[i][j] = new_tile
      arranged_ids[i][j] = output[:id]
    else
      raise 'dont know what to do'
    end
    j += 1
  end
  i += 1
end

borderless = []
arranged.each_with_index do |tiles_row, index|
  (1...tiles_row[0].size - 1).each do |n|
    borderless << []
    tiles_row.each do |t|
      borderless[(index * (tiles_row[0].size - 2)) + (n - 1)].concat(t[n][1...-1])
    end
  end
end
                  #
#    ##    ##    ###
 #  #  #  #  #  #
relative_sea_monster_coordinates = [
  [0, 18],
  [1, 0],
  [1, 5],
  [1, 6],
  [1, 11],
  [1, 12],
  [1, 17],
  [1, 18],
  [1, 19],
  [2, 1],
  [2, 4],
  [2, 7],
  [2, 10],
  [2, 13],
  [2, 16]
]
check_for_sea_monsters = -> (image) do
  found_monsters = false
  i = 0
  while i < image.size - 2
    j = 0
    while j < image[0].size - 19
      if relative_sea_monster_coordinates.all? { |y, x| image[i + y][j + x] == '#' }
        relative_sea_monster_coordinates.each { |y, x| image[i + y][j + x] = 'O' }
        found_monsters = true
      end
      j += 1
    end
    i += 1
  end
  found_monsters
end

4.times do
  borderless = rotate(borderless, 90)
  break if check_for_sea_monsters.call(borderless)

  orig_borderless = borderless
  done = false
  %i[flip_h flip_v].each do |m|
     borderless = method(m).call(borderless)
     if check_for_sea_monsters.call(borderless)
       done = true
       break
     end
  end
  break if done
  borderless = orig_borderless
end
puts borderless.map { |line| line.join('') }.join("\n").count('#')
# puts borderless.map { |line| line.join('') }.join("\n")
