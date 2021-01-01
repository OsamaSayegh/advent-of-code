# frozen_string_literal: true

content = File.read(File.expand_path('input.txt', __dir__))
PASSPORTS = content.split("\n\n").map do |raw|
  hash = {}
  hash[:byr] = $1 if raw =~ /byr:([^\s]+)/
  hash[:iyr] = $1 if raw =~ /iyr:([^\s]+)/
  hash[:eyr] = $1 if raw =~ /eyr:([^\s]+)/
  hash[:hgt] = $1 if raw =~ /hgt:([^\s]+)/
  hash[:hcl] = $1 if raw =~ /hcl:([^\s]+)/
  hash[:ecl] = $1 if raw =~ /ecl:([^\s]+)/
  hash[:pid] = $1 if raw =~ /pid:([^\s]+)/
  hash[:cid] = $1 if raw =~ /cid:([^\s]+)/
  hash
end

valid1 = PASSPORTS.count do |hash|
  hash.keys.size == 8 || (hash.keys.size == 7 && !hash.key?(:cid))
end
puts "PART 1: #{valid1}"

valid2 = PASSPORTS.count do |hash|
  byr = hash[:byr]
  iyr = hash[:iyr]
  eyr = hash[:eyr]
  hgt = hash[:hgt]
  hcl = hash[:hcl]
  ecl = hash[:ecl]
  pid = hash[:pid]
  (byr && (1920..2002).include?(byr.to_i)) &&
  (iyr && (2010..2020).include?(iyr.to_i)) &&
  (eyr && (2020..2030).include?(eyr.to_i)) &&
  (hgt && hgt =~ /^([\d]+)(cm|in)$/ && (($2 == 'in' && (59..76).include?($1.to_i)) || ($2 == 'cm' && (150..193).include?($1.to_i)))) &&
  (hcl && hcl.downcase =~ /^#\h{6}$/) &&
  (%w[amb blu brn gry grn hzl oth].include?(ecl)) &&
  (pid && pid =~ /^\d{9}$/)
end

puts "PART 2: #{valid2}"
