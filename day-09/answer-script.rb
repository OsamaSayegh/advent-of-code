# frozen_string_literal: true

content = <<~DATA
35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576
DATA
content = File.read(File.expand_path('input.txt', __dir__))
PREAMPLE = 25
NUMBERS = content.split("\n").map { |l| l.to_i }

res = 0
NUMBERS.each_with_index do |n, ind|
  next if ind < PREAMPLE
  if NUMBERS[(ind - PREAMPLE)...ind].combination(2).none? { |a, b| a + b == n }
    puts n
    res = n
    break
  end
end

g = 2
while true
  NUMBERS.each_with_index do |n, ind|
    sub = NUMBERS[ind...ind+g]
    if sub.sum == res
      sub.sort!
      puts sub[0] + sub[-1]
      exit
    end
  end
  g += 1
end
