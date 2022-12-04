local readFile = require('readFile')
local exampleInput = [[vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
]]

local input = readFile('input.txt')

local function getCode(str)
  local code = 0
  local byte = string.byte(str)
  if byte > 90 then -- lowercase
    code = byte - 96
  else -- uppercase
    code = byte - 38
  end
  return code
end

local function part1()
  local total = 0

  for token in string.gmatch(input, "[^\r\n]+") do
    local compartment1 = string.sub(token, 0, (#token / 2))
    local compartment2 = string.sub(token, (#token / 2) + 1, #token)

    local match = string.match(compartment1, "[" .. compartment2 .. "]+")

    if #match > 1 then
      match = string.sub(match, 0, 1)
    end

    total = total + getCode(match)
  end

  return total
end

print('Part 1:', part1())

local function part2()
  local total = 0
  local index = 1
  local group = ''
  for token in string.gmatch(input, "[^\r\n]+") do
    group = group .. ' ' .. token
    if index % 3 == 0 then
      for a, b, c in string.gmatch(group, "(%w+) (%w+) (%w+)") do
        local match1 = ''
        for m in string.gmatch(a, "[" .. b .. "]+") do
          match1 = match1 .. m
        end
        local match2 = string.match(c, "[" .. match1 .. "]+")

        total = total + getCode(match2)
      end
      group = ''
    end
    index = index + 1
  end
  return total
end

print('Part 2:', part2())
