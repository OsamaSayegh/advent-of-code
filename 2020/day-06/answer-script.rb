# frozen_string_literal: true

content = File.read(File.expand_path('input.txt', __dir__))
GROUPS = content.split("\n\n")

sum = GROUPS.sum do |lines|
  all = lines.gsub("\n", '')
  all = all.split('').uniq.join('')
  all.size
end

puts sum

sum = GROUPS.sum do |lines|
  arr = lines.split("\n")
  chars = arr[0].split('').uniq
  arr.each do |line|
    chars = chars & line.split('').uniq
  end
  chars.size
end

puts sum
