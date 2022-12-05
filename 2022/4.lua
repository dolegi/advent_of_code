local readFile = require('readFile')

local input = readFile('input.txt')

local exampleInput = [[
2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
]]

local function getRange(str)
  local start = tonumber(string.match(str, "(%d+)-"))
  local finish = tonumber(string.match(str, ".*-(%d+)"))
  local result = ''

  for i = start, finish do
    result = result .. ',' .. tostring(i)
  end
  return result .. ','
end

local function part1()
  local total = 0

  for pair in string.gmatch(input, "[^\r\n]+") do
    for a, b in string.gmatch(pair, "(.+),(.+)") do
      local aRange = getRange(a)
      local bRange = getRange(b)

      local aMatch = string.find(aRange, '.*' .. bRange .. '.*')
      local bMatch = string.find(bRange, '.*' .. aRange .. '.*')
      if aMatch ~= nil or bMatch ~= nil then
        total = total + 1
      end
    end
  end
  return total
end

print('Part 1:', part1())

local function part2()
  local total = 0

  for pair in string.gmatch(input, "[^\r\n]+") do
    for a, b in string.gmatch(pair, "(.+),(.+)") do
      local aStart = tonumber(string.match(a, "(%d+)-"))
      local aFinish = tonumber(string.match(a, ".*-(%d+)"))
      local bStart = tonumber(string.match(b, "(%d+)-"))
      local bFinish = tonumber(string.match(b, ".*-(%d+)"))

      if not (aStart > bFinish or bStart > aFinish) then
        total = total + 1
      end
    end
  end
  return total
end

print('Part 2:', part2())
