# frozen_string_literal: true

# content = <<~DATA
# mask = 000000000000000000000000000000X1001X
# mem[42] = 100
# mask = 00000000000000000000000000000000X0XX
# mem[26] = 1
# DATA
content = File.read(File.expand_path('input.txt', __dir__))
lines = content.split("\n")
mask = nil
memory = {}
lines.each do |line|
  if line =~ /^mask = (.+)$/
    mask = $1
  elsif line =~ /^mem\[(\d+)\] = (\d+)$/
    address = Integer($1)
    val = Integer($2)
    res = val.to_s(2).rjust(36, '0')
    mask.split('').each_with_index do |c, i|
      if c != 'X'
        res[i] = c
      end
    end
    memory[address] = res.to_i(2)
  else
    raise line
  end
end

puts memory.values.sum

memory = {}
mask = nil

lines.each do |line|
  if line =~ /^mask = (.+)$/
    mask = $1
  elsif line =~ /^mem\[(\d+)\] = (\d+)$/
    val = Integer($2)
    address = Integer($1)
    bin_address = address.to_s(2).rjust(36, '0')
    mask.split('').each_with_index do |c, i|
      if c == '1' || c == 'X'
        bin_address[i] = c
      end
    end
    addresses = [bin_address]
    while !addresses.all? { |a| !a.include?('X') }
      new_addresses = []
      addresses.each do |a|
        new_addresses << a.sub('X', '0')
        new_addresses << a.sub('X', '1')
      end
      addresses = new_addresses
    end
    addresses.each do |a|
      memory[a.to_i(2)] = val
    end
  else
    raise line
  end
end

puts memory.values.sum
