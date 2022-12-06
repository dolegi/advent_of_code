local inspect = require('inspect')
local readFile = require('readFile')

local exampleInput = 'mjqjpqmgbljsphdztnvjfqwrcgsmlb'
local realInput = readFile('input.txt')

local function areUnique(arr)
  for j = 1, #arr do
    for k = j + 1, #arr do
      if arr[j] == arr[k] then
        return false
      end
    end
  end
  return true
end

local function solution(input, marker)
  local detected = {}

  for i = 1, marker-1 do
    table.insert(detected, input:sub(i, i))
  end


  for i = marker, #input do
    local c = input:sub(i, i)
    table.insert(detected, c)
    if areUnique(detected) then
      return i
    end
    table.remove(detected, 1)
  end
  return 0
end

print('Part 1: ', solution(realInput, 4))
print('Part 2: ', solution(realInput, 14))
