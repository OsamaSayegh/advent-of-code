# frozen_string_literal: true

# content = <<~DATA
# .#.
# ..#
# ###
# DATA
content = File.read(File.expand_path('input.txt', __dir__))
$hash = {}
def lookup(x, y, z)
  v = $hash["#{x}#{y}#{z}"]
  v || '.'
end

def set(hash, x, y, z, v)
  hash["#{x}#{y}#{z}"] = v
end

neighbors = []
arr = [-1, 0, 1]
arr.each do |z|
  arr.each do |y|
    arr.each do |x|
      next if z == 0 && y == 0 && x == 0
      neighbors << [z, y, x]
    end
  end
end

min_y = 0
min_x = 0
max_y = 0
max_x = 0
content.split("\n").each_with_index do |line, y|
  max_y = [max_y, y + 1].max
  line.split('').each_with_index do |e, x|
    max_x = [max_x, x + 1].max
    set($hash, x, y, 0, e)
  end
end

(1..6).each do |n|
  hash = $hash.dup
  (-n..n).each do |z|
    ((min_y - 1)..(max_y + 1)).each do |y|
      ((min_x - 1)..(max_x + 1)).each do |x|
        active_neighbors = 0
        neighbors.each do |k, j, i|
          nz = z + k
          ny = y + j
          nx = x + i
          active_neighbors += 1 if lookup(nx, ny, nz) == '#'
        end
        element = lookup(x, y, z)
        if element == '.'
          if active_neighbors == 3
            set(hash, x, y, z, '#')
          else
            set(hash, x, y, z, '.')
          end
        else
          if active_neighbors == 3 || active_neighbors == 2
            set(hash, x, y, z, '#')
          else
            set(hash, x, y, z, '.')
          end
        end
      end
    end
  end
  min_y -= 1
  min_x -= 1
  max_y += 1
  max_x += 1
  $hash = hash
end

puts $hash.values.select { |v| v == '#' }.size


$hash = {}
def lookup(x, y, z, w)
  v = $hash["#{x}|#{y}|#{z}|#{w}"]
  v || '.'
end

def set(hash, x, y, z, w, v)
  hash["#{x}|#{y}|#{z}|#{w}"] = v
end

neighbors = []
arr = [-1, 0, 1]
arr.each do |w|
  arr.each do |z|
    arr.each do |y|
      arr.each do |x|
        next if w == 0 && z == 0 && y == 0 && x == 0
        neighbors << [w, z, y, x]
      end
    end
  end
end

min_y = 0
min_x = 0

max_y = 0
max_x = 0
content.split("\n").each_with_index do |line, y|
  max_y = [max_y, y + 1].max
  line.split('').each_with_index do |e, x|
    max_x = [max_x, x + 1].max
    set($hash, x, y, 0, 0, e)
  end
end

(1..6).each do |n|
  hash = $hash.dup
  (-n..n).each do |w|
    (-n..n).each do |z|
      ((min_y - 1)..(max_y + 1)).each do |y|
        ((min_x - 1)..(max_x + 1)).each do |x|
          active_neighbors = 0
          neighbors.each do |h, k, j, i|
            nw = w + h
            nz = z + k
            ny = y + j
            nx = x + i
            active_neighbors += 1 if lookup(nx, ny, nz, nw) == '#'
          end
          element = lookup(x, y, z, w)
          if element == '.'
            if active_neighbors == 3
              set(hash, x, y, z, w, '#')
            else
              set(hash, x, y, z, w, '.')
            end
          else
            if active_neighbors == 3 || active_neighbors == 2
              set(hash, x, y, z, w, '#')
            else
              set(hash, x, y, z, w, '.')
            end
          end
        end
      end
    end
  end
  min_y -= 1
  min_x -= 1
  max_y += 1
  max_x += 1
  $hash = hash
end

puts $hash.values.select { |v| v == '#' }.size
exit
