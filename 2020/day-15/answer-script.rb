# frozen_string_literal: true

# content = <<~DATA
# DATA
content = File.read(File.expand_path('input.txt', __dir__))
numbers = content.split(",").map(&:to_i)

def find(numbers, position)
  hash1 = {}
  hash2 = {}
  numbers = numbers.dup
  numbers.each_with_index { |n, i| hash1[n] = hash2[n] = i }
  while numbers.size < position
    last = numbers[-1]
    f = hash1[last]
    l = hash2[last]
    index = numbers.size
    if !f || !l
      numbers << 0
      hash2[0] = hash1[0]
      hash1[0] = index
      next
    end
    new = f - l
    numbers << new
    hash2[new] = hash1[new]
    hash1[new] = index
  end
  numbers[-1]
end

puts find(numbers, 2020)
puts find(numbers, 30000000)
