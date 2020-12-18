# frozen_string_literal: true

content = <<~DATA
1 + 2 * 3 + 4 * 5 + 6
DATA
content = File.read(File.expand_path('input.txt', __dir__))
expressions = content.split("\n")

def evaluate(expr, p2)
  i = 0
  operations = []
  term = nil
  operator = nil
  operations = []
  while i < expr.size
    if expr[i] == '('
      j = i + 1
      nested = 0
      while true
        nested += 1 if expr[j] == '('
        if expr[j] == ')'
          if nested == 0
            term = evaluate(expr[(i + 1)...j], p2)
            i = j
            break
          else
            nested -= 1
          end
        end
        j += 1
      end
    elsif expr[i] =~ /^\d$/
      term = Integer(expr[i])
    elsif %w[* +].include?(expr[i])
      operator = expr[i]
    else
      raise 'unexpected value' if expr[i] != ' '
    end
    if term
      operations << term
      term = nil
    end
    if operator
      operations << operator
      operator = nil
    end
    i += 1
  end
  if p2
    h = 1
    while h < operations.size - 1
      if operations[h] == '+'
        subres = operations[h - 1] + operations[h + 1]
        operations[h - 1] = subres
        operations.delete_at(h + 1)
        operations.delete_at(h)
      else
        h += 2
      end
    end
    res = eval(operations.join('')) # lol
  else
    res = operations[0]
    h = 1
    while h < operations.size - 1
      res = res.send(operations[h], operations[h + 1])
      h += 2
    end
  end
  res
end
puts expressions.map { |expr| evaluate(expr, false) }.sum
puts expressions.map { |expr| evaluate(expr, true) }.sum
