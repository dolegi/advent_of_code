local inspect = require('inspect')

local function readAll(file)
  local f = assert(io.open(file, "rb"))
  local content = f:read("*all")
  f:close()
  return content
end

local input = readAll('./input.txt')

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
print('Part 2: ', calories[#calories] + calories[#calories-1] + calories[#calories-2])
