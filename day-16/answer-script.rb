# frozen_string_literal: true

# content = <<~DATA
# class: 1-3 or 5-7
# row: 6-11 or 33-44
# seat: 13-40 or 45-50
# 
# your ticket:
# 7,1,14
# 
# nearby tickets:
# 7,3,47
# 40,4,50
# 55,2,20
# 38,6,12
# DATA
content = File.read(File.expand_path('input.txt', __dir__))
fields = content.split("\n\n")[0].split("\n").map do |field|
  name = field.split(": ")[0]
  ranges = field.split(": ")[1].split(" or ").map do |range|
    (range.split("-")[0].to_i..range.split("-")[1].to_i)
  end
  { name: name, ranges: ranges }
end

my_ticket = content.split("\n\n")[1].split("\n")[1].split(",").map(&:to_i)
nearby = content.split("\n\n")[2].split("\n")[1..-1].map do |ticket|
  ticket.split(",").map(&:to_i)
end

invalid = nearby.flatten.select { |n| fields.all? { |f| f[:ranges].all? { |r| !r.include?(n) } } }
puts invalid.sum

nearby.select! { |t| t.none? { |n| invalid.include?(n) } }
field_names = fields.map { |f| f[:name] }

i = 0
size = nearby[0].size
ordered_field_names = []
size.times { ordered_field_names << nil }
while !field_names.empty?
  column = nearby.map { |t| t[i] }
  matches = fields
    .select { |f| field_names.include?(f[:name]) }
    .select { |f| column.all? { |v| f[:ranges].any? { |r| r.include?(v) } } }
  if matches.size == 1
    ordered_field_names[i] = matches[0][:name]
    field_names -= [matches[0][:name]]
  end
  i += 1
  i %= size
end

fields.sort_by! { |f| ordered_field_names.index(f[:name]) }
prod = 1
fields.each_with_index do |f, i|
  prod *= my_ticket[i] if f[:name].start_with?('departure')
end
puts prod
