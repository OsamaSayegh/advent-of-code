# frozen_string_literal: true

# content = <<~DATA
# 0: 4 1 5
# 1: 2 3 | 3 2
# 2: 4 4 | 5 5
# 3: 4 5 | 5 4
# 4: "a"
# 5: "b"
# 
# ababbb
# bababa
# abbbab
# aaabbb
# aaaabbb
# DATA
# content = <<~DATA
# 42: 9 14 | 10 1
# 9: 14 27 | 1 26
# 10: 23 14 | 28 1
# 1: "a"
# 11: 42 31
# 5: 1 14 | 15 1
# 19: 14 1 | 14 14
# 12: 24 14 | 19 1
# 16: 15 1 | 14 14
# 31: 14 17 | 1 13
# 6: 14 14 | 1 14
# 2: 1 24 | 14 4
# 0: 8 11
# 13: 14 3 | 1 12
# 15: 1 | 14
# 17: 14 2 | 1 7
# 23: 25 1 | 22 14
# 28: 16 1
# 4: 1 1
# 20: 14 14 | 1 15
# 3: 5 14 | 16 1
# 27: 1 6 | 14 18
# 14: "b"
# 21: 14 1 | 1 14
# 25: 1 1 | 1 14
# 22: 14 14
# 8: 42
# 26: 14 22 | 1 20
# 18: 15 15
# 7: 14 5 | 1 21
# 24: 14 1
# 
# abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa
# bbabbbbaabaabba
# babbbbaabbbbbabbbbbbaabaaabaaa
# aaabbbbbbaaaabaababaabababbabaaabbababababaaa
# bbbbbbbaaaabbbbaaabbabaaa
# bbbababbbbaaaaaaaabbababaaababaabab
# ababaaaaaabaaab
# ababaaaaabbbaba
# baabbaaaabbaaaababbaababb
# abbbbabbbbaaaababbbbbbaaaababb
# aaaaabbaabaaaaababaa
# aaaabbaaaabbaaa
# aaaabbaabbaaaaaaabbbabbbaaabbaabaaa
# babaaabbbaaabaababbaabababaaab
# aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba
# DATA
content = File.read(File.expand_path('input.txt', __dir__))
$raw_rules, input = content.split("\n\n").map { |p| p.split("\n") }

def resolve(id, p2)
  raw = $raw_rules.find { |r| r.start_with?("#{id}:") }
  raise "no rule found: #{id}" unless raw
  raw = raw.split(": ")[1]
  raw = "42 31 | 42 11 31" if id == 11 && p2
  raw = "42 | 42 8" if id == 8 && p2
  if p2 && id == 11
    x = resolve(42, p2)
    y = resolve(31, p2)
    "(#{x}(#{x}(#{x}(#{x}(#{x}(#{x}(#{x}(#{x}(#{x}(#{x}(#{x}#{y})?#{y})?#{y})?#{y})?#{y})?#{y})?#{y})?#{y})?#{y})?#{y})?#{y})"
  elsif p2 && id == 8
    x = resolve(42, p2)
    "(#{x}+)"
  elsif raw =~ /"(\w+)"/
    $1
  else
    res = raw.split(" | ")
      .map do |ids|
        ids.split(" ")
          .map do |rid|
            rid = Integer(rid)
            resolve(rid, p2)
          end
          .join('')
      end
      .join('|')
    "(#{res})"
  end
end
regex = /^#{resolve(0, false)}$/
puts input.count { |i| regex.match?(i) }
regex = /^#{resolve(0, true)}$/
puts input.count { |i| regex.match?(i) }
