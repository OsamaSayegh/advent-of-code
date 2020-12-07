# frozen_string_literal: true
require 'set'

content = File.read(File.expand_path('input.txt', __dir__))
# content = <<~DATA
# light red bags contain 1 bright white bag, 2 muted yellow bags.
# dark orange bags contain 3 bright white bags, 4 muted yellow bags.
# bright white bags contain 1 shiny gold bag.
# muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
# shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
# dark olive bags contain 3 faded blue bags, 4 dotted black bags.
# vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
# faded blue bags contain no other bags.
# dotted black bags contain no other bags.
# DATA
#
# content = <<~DATA
# shiny gold bags contain 2 dark red bags.
# dark red bags contain 2 dark orange bags.
# dark orange bags contain 2 dark yellow bags.
# dark yellow bags contain 2 dark green bags.
# dark green bags contain 2 dark blue bags.
# dark blue bags contain 2 dark violet bags.
# dark violet bags contain no other bags.
# DATA
RULES = content.split("\n").map do |rule|
  if rule =~ /(\w+\s\w+)\sbags\scontain no/
    { container: $1, containees: [] }
  elsif rule =~ /(\w+\s\w+)\sbags\scontain\s(\d+)\s(\w+\s\w+)bags?\.$/
    { container: $1, containees: [{ type: $3, count: $2.to_i }] }
  elsif rule =~ /(\w+\s\w+)\sbags\scontain\s(.+)/
    hash = { container: $1, containees: [] }
    $2.split(', ').each do |type|
      type =~ /(\d+)\s(\w+\s\w+)\s/
      hash[:containees] << { count: $1.to_i, type: $2 }
    end
    hash
  end
end


def find_roots(bag)
  bags = Set.new
  RULES.each do |rule|
    if rule[:containees].any? { |c| c[:type] == bag }
      bags << rule[:container]
    end
  end
  bags.dup.each do |bag|
    find_roots(bag).each { |b| bags << b }
  end
  bags
end

def count_children(bag)
  count = 0
  RULES.each do |rule|
    if rule[:container] == bag
      rule[:containees].each do |child|
        count += child[:count]
        count += count_children(child[:type]) * child[:count]
      end
    end
  end
  count
end

puts find_roots('shiny gold').size
puts count_children('shiny gold')
