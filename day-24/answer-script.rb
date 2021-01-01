# frozen_string_literal: true

# content = <<~DATA
# sesenwnenenewseeswwswswwnenewsewsw
# neeenesenwnwwswnenewnwwsewnenwseswesw
# seswneswswsenwwnwse
# nwnwneseeswswnenewneswwnewseswneseene
# swweswneswnenwsewnwneneseenw
# eesenwseswswnenwswnwnwsewwnwsene
# sewnenenenesenwsewnenwwwse
# wenwwweseeeweswwwnwwe
# wsweesenenewnwwnwsenewsenwwsesesenwne
# neeswseenwwswnwswswnw
# nenwswwsewswnenenewsenwsenwnesesenew
# enewnwewneswsewnwswenweswnenwsenwsw
# sweneswneswneneenwnewenewwneswswnese
# swwesenesewenwneswnwwneseswwne
# enesenwswwswneneswsenwnewswseenwsese
# wnwnesenesenenwwnenwsewesewsesesew
# nenewswnwewswnenesenwnesewesw
# eneswnwswnwsenenwnwnwwseeswneewsenese
# neswnwewnwnwseenwseesewsenwsweewe
# wseweeenwnesenwwwswnew
# DATA
content = File.read(File.expand_path('input.txt', __dir__))
lines = content.split("\n")
tiles = {}
lines.each do |tile|
  x = 0
  y = 0
  i = 0
  while i < tile.size
    case tile[i]
    when 'e'
      x += 1
      i += 1
    when 'w'
      x -= 1
      i += 1
    when 's'
      y -= 1
      if tile[i + 1] == 'e'
        x += 1
      end
      i += 2
    when 'n'
      y += 1
      if tile[i + 1] == 'w'
        x -= 1
      end
      i += 2
    end
  end
  key = [x, y]
  if tiles.key?(key)
    tiles[key] = !tiles[key]
  else
    tiles[key] = false
  end
end
puts tiles.count { |_, white| !white }

neighbors = [
  [1, 0],
  [-1, 0],
  [0, 1],
  [0, -1],
  [1, -1],
  [-1, 1]
]
100.times do |n|
  _tiles = tiles.dup
  tiles.each do |(x, y), white|
    next if white
    neighbors.each do |i, j|
      if !_tiles.key?([x + i, y + j])
        _tiles[[x + i, y + j]] = true
      end
    end
  end
  tiles = _tiles
  _tiles = tiles.dup
  tiles.each do |(x, y), white|
    black_count = neighbors.count do |i, j|
      k = [x + i, y + j]
      tiles.key?(k) && !tiles[k]
    end
    if !white
      _tiles[[x, y]] = true if black_count == 0 || black_count > 2
    else
      _tiles[[x, y]] = false if black_count == 2
    end
  end
  tiles = _tiles
end
puts tiles.count { |_, white| !white }
