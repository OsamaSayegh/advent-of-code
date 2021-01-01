# frozen_string_literal: true

content = File.read(File.expand_path('input.txt', __dir__))
$numbers = content.split("\n").map { |n| Integer(n) }

YEAR = 2020

def part_1
  $numbers.each_with_index do |n1, ind|
    $numbers[ind + 1..-1].each do |n2|
      sum = n1 + n2
      if sum == YEAR
        prod = n1 * n2
        puts "My numbers are: #{n1} and #{n2}. Their sum is #{sum} and their product is #{prod}"
        return
      end
    end
  end
end

def part_2
  $numbers.each_with_index do |n1, ind|
    $numbers[ind + 1..-1].each_with_index do |n2, ind2|
      $numbers[ind2 + 1..-1].each do |n3|
        sum = n1 + n2 + n3
        if sum == YEAR
          prod = n1 * n2 * n3
          puts "My numbers are: #{n1}, #{n2} and #{n3}. Their sum is #{sum} and their product is #{prod}"
          return
        end
      end
    end
  end
end

part_1
part_2
