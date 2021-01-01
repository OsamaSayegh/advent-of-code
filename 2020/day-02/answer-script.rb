# frozen_string_literal: true

entries = File.read(File.expand_path('input.txt', __dir__)).split("\n")

entries.map! do |entry|
  data = entry.match(/^(?<min>\d+)-(?<max>\d+)\s(?<char>\w{1}):\s(?<password>\w+)$/)
  min = Integer(data[:min])
  max = Integer(data[:max])
  char = data[:char]
  password = data[:password]

  [min, max, char, password]
end

part_1_valid = entries.count do |min, max, char, password|
  occurrences = password.size - password.gsub(char, '').size
  occurrences <= max && occurrences >= min
end

part_2_valid = entries.count do |min, max, char, password|
  c1 = password[min - 1]
  c2 = password[max - 1]
  (c1 == char && c2 != char) || (c1 != char && c2 == char)
end

puts "[PART 1] Number of valid entries is #{part_1_valid}"
puts "[PART 2] Number of valid entries is #{part_2_valid}"
