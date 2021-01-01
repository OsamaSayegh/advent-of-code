# frozen_string_literal: true

# content = <<~DATA
# L.LL.LL.LL
# LLLLLLL.LL
# L.L.L..L..
# LLLL.LL.LL
# L.LL.LL.LL
# L.LLLLL.LL
# ..L.L.....
# LLLLLLLLLL
# L.LLLLLL.L
# L.LLLLL.LL
# DATA
content = File.read(File.expand_path('input.txt', __dir__))
SEATS = content.split("\n").map { |l| l.split('') }

def iteratep1(seats)
  new_seats = []
  seats.each_with_index do |row, i|
    new_seats << []
    row.each_with_index do |seat, j|
      occupied = 0
      [
        [0, 1],
        [0, -1],
        [1, 0],
        [1, 1],
        [1, -1],
        [-1, 0],
        [-1, 1],
        [-1, -1]
      ].each do |x, y|
        i1 = i + x
        j1 = j + y
        next if i1 < 0 || j1 < 0 || i1 >= seats.size || j1 >= seats[0].size
        occupied += 1 if seats[i1][j1] == '#'
      end
      if seat == 'L' && occupied == 0
        new_seats[i][j] = '#'
      elsif seat == '#' && occupied >= 4
        new_seats[i][j] = 'L'
      else
        new_seats[i][j] = seats[i][j]
      end
    end
  end
  new_seats
end

prev = SEATS
seats = iteratep1(prev)
while prev != seats
  prev = seats
  seats = iteratep1(seats)
end

puts seats.map { |row| row.count { |s| s == '#' } }.sum

def iteratep2(seats)
  new_seats = []
  seats.each_with_index do |row, i|
    new_seats << []
    row.each_with_index do |seat, j|
      occupied = 0
      [
        [0, 1],
        [0, -1],
        [1, 0],
        [1, 1],
        [1, -1],
        [-1, 0],
        [-1, 1],
        [-1, -1]
      ].each do |x, y|
        i1 = i + x
        j1 = j + y
        while !(i1 < 0 || j1 < 0 || i1 >= seats.size || j1 >= seats[0].size)
          if seats[i1][j1] == '#'
            occupied += 1
            break
          elsif seats[i1][j1] == 'L'
            break
          end
          i1 += x
          j1 += y
        end
      end
      if seat == 'L' && occupied == 0
        new_seats[i][j] = '#'
      elsif seat == '#' && occupied >= 5
        new_seats[i][j] = 'L'
      else
        new_seats[i][j] = seats[i][j]
      end
    end
  end
  new_seats
end

prev = SEATS
seats = iteratep2(prev)
while prev != seats
  prev = seats
  seats = iteratep2(seats)
end

puts seats.map { |row| row.count { |s| s == '#' } }.sum
