# frozen_string_literal: true

require 'set'
content = File.read(File.expand_path('input.txt', __dir__))
# content = <<~DATA
# nop +0
# acc +1
# jmp +4
# acc +3
# jmp -3
# acc -99
# acc +1
# jmp -4
# acc +6
# DATA
instructions = content.split("\n").map do |inst|
  [inst.split(' ')[0], inst.split(' ')[1].to_i]
end

def will_finish(instructions, print_acc: false)
  pos = 0
  seen = []
  acc = 0

  while pos < instructions.size
    instr = instructions[pos]
    if seen.include?(pos)
      puts acc if print_acc
      return false
    end
    seen << pos
    if instr[0] == 'nop'
      pos += 1
      next
    end
    if instr[0] == 'acc'
      pos += 1
      acc += instr[1]
      next
    end
    if instr[0] == 'jmp'
      pos += instr[1]
      next
    end
  end
  puts acc
  true
end

will_finish(instructions, print_acc: true)

g = 0
while g < instructions.size
  if instructions[g][0] == 'jmp'
    dup = instructions.map { |i| i.dup }
    dup[g][0] = 'nop'
    will_finish(dup)
  elsif instructions[g][0] == 'nop'
    dup = instructions.map { |i| i.dup }
    dup[g][0] = 'jmp'
    will_finish(dup)
  end
  g += 1
end
