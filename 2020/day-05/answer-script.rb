# frozen_string_literal: true

content = File.read(File.expand_path('input.txt', __dir__))
seats = content.split("\n")

ROWS = (0..127).to_a
COLS = (0..7).to_a

def recursive(seat, rows)
  dir = seat[0]
  if dir == 'F' || dir == 'L'
    return rows[0] if rows.size == 2
    recursive(seat[1..-1], rows[0..(rows.size / 2 -1)])
  else
    return rows[1] if rows.size == 2
    recursive(seat[1..-1], rows[(rows.size / 2)..-1])
  end
end

ids = seats.map do |seat|
  recursive(seat[0...7], ROWS) * 8 + recursive(seat[7..-1], COLS)
end

# ids = %w[BFFFBBFRRR FFFBBBFRRR BBFFBBFRLL].map do |seat|
#   recursive(seat[0...7], ROWS) * 8 + recursive(seat[7..-1], COLS)
# end

puts "PART 1: #{ids.max}"

puts "PART 2: #{(ids.min..ids.max).to_a - ids}"
