# frozen_string_literal: true

content = <<~DATA
5764801
17807724
DATA
content = File.read(File.expand_path('input.txt', __dir__))

def brute_force(target)
  i = 0
  v = 1
  while v != target
    v *= 7
    v %= 20201227
    i += 1
  end
  i
end

def calc_enc_key(pri, pub)
  key = 1
  pri.times do
    key *= pub
    key %= 20201227
  end
  key
end

c_pub, d_pub = content.split("\n").map(&:to_i)
c_pri, d_pri = brute_force(c_pub), brute_force(d_pub)

raise 'encryption keys dont match' if calc_enc_key(c_pri, d_pub) != calc_enc_key(d_pri, c_pub)
puts calc_enc_key(c_pri, d_pub)
