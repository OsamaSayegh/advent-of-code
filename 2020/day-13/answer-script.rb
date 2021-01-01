# frozen_string_literal: true

# content = <<~DATA
# 939
# 7,13,x,x,59,x,31,19
# DATA
content = File.read(File.expand_path('input.txt', __dir__))
timestamp = content.split("\n")[0].to_i
ids_p1 = content.split("\n")[1].split(",").reject { |id| id == 'x' }.map(&:to_i)

res = ids_p1.map do |id|
  ceil = (timestamp / id.to_f).ceil
  [id, ceil * id]
end.sort_by { |_, b| b }.first

puts res[0] * (res[1] - timestamp)

ids_p2 = content.split("\n")[1].split(",").each_with_index.to_a
ids_p2.reject! { |a, _| a == 'x' }
ids_p2.map! { |a, ind| [a.to_i, ind] }

i = ids_p2[0][0]
inc = i
ids_p2.shift

while ids_p2.size > 0
  if (i + ids_p2[0][1]) % ids_p2[0][0] == 0
    inc *= ids_p2[0][0]
    ids_p2.shift
    next
  end
  i += inc
end

puts i
