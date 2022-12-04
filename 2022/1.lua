local readFile = require('readFile')

local input = readFile('./input.txt')

local calories = { 0 }
local i = 1

for token in string.gmatch(input, "[^\r\n]*") do
  if token == '' then
    table.insert(calories, 0)
    i = i + 1
  else
    local number = tonumber(token)
    if number ~= nil then
      calories[i] = calories[i] + number
    end
  end
end

table.sort(calories)
print('Part 1: ', calories[#calories])
print('Part 2: ', calories[#calories] + calories[#calories - 1] + calories[#calories - 2])
