# frozen_string_literal: true

# content = <<~DATA
# mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
# trh fvjkl sbzzf mxmxvkd (contains dairy)
# sqjhc fvjkl (contains soy)
# sqjhc mxmxvkd sbzzf (contains fish)
# DATA
content = File.read(File.expand_path('input.txt', __dir__))
lines = content.split("\n")
allergens = {}
ingredients = []
food = []
lines.each do |f|
  f =~ /^(.+) \(contains (.+)\)$/
  ing = $1.split(" ")
  aller = $2.split(", ")
  food << ing
  ingredients |= ing
  aller.each do |a|
    if allergens[a]
      allergens[a] &= ing
    else
      allergens[a] = ing
    end
  end
end

while !allergens.all? { |_, ing| ing.size == 1 }
  allergens.each do |name, ing|
    if ing.size == 1
      allergens.each do |name2, ing2|
        next if name == name2
        ing2.delete(ing[0])
      end
    end
  end
end
puts (ingredients - allergens.values.flatten).sum { |ing| food.count { |ings| ings.include?(ing) } }
puts allergens.sort_by { |k,v| k }.to_h.values.flatten.join(',')
