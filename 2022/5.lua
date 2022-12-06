local inspect = require('inspect')
local readFile = require('readFile')

local exampleStacks = {
  { 'Z', 'N' },
  { 'M', 'C', 'D' },
  { 'P' }
}

local exampleInstructions = [[move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
]]

local realStacks = {
  -- [J]             [F] [M]
  -- [Z] [F]     [G] [Q] [F]
  -- [G] [P]     [H] [Z] [S] [Q]
  -- [V] [W] [Z] [P] [D] [G] [P]
  -- [T] [D] [S] [Z] [N] [W] [B] [N]
  -- [D] [M] [R] [J] [J] [P] [V] [P] [J]
  -- [B] [R] [C] [T] [C] [V] [C] [B] [P]
  -- [N] [S] [V] [R] [T] [N] [G] [Z] [W]
  --  1   2   3   4   5   6   7   8   9
  { 'N', 'B', 'D', 'T', 'V', 'G', 'Z', 'J' },
  { 'S', 'R', 'M', 'D', 'W', 'P', 'F' },
  { 'V', 'C', 'R', 'S', 'Z' },
  { 'R', 'T', 'J', 'Z', 'P', 'H', 'G' },
  { 'T', 'C', 'J', 'N', 'D', 'Z', 'Q', 'F' },
  { 'N', 'V', 'P', 'W', 'G', 'S', 'F', 'M' },
  { 'G', 'C', 'V', 'B', 'P', 'Q' },
  { 'Z', 'B', 'P', 'N' },
  { 'W', 'P', 'J' }
}

local realInstructions = readFile('input.txt')

local function part1(instructions, stacks)
  for line in string.gmatch(instructions, "[^\r\n]+") do
    for move, from, to in string.gmatch(line, "(%d+).-(%d+).-(%d+).-") do

      local fromStack = stacks[tonumber(from)]
      local toStack = stacks[tonumber(to)]

      for i = 1, move do
        local top = table.remove(fromStack)
        table.insert(toStack, top)
      end

    end
  end
  local result = ''
  for i = 1, #stacks do
    local stack = stacks[i]
    result = result .. stack[#stack]
  end
  return result
end

-- print('Part 1:', part1(realInstructions, realStacks))

local function part2(instructions, stacks)
  for line in string.gmatch(instructions, "[^\r\n]+") do
    for move, from, to in string.gmatch(line, "(%d+).-(%d+).-(%d+).-") do

      local fromStack = stacks[tonumber(from)]
      local toStack = stacks[tonumber(to)]

      local top = {}
      for i = 1, move do
        table.insert(top, table.remove(fromStack))
      end
      local rev = {}
      for i = #top, 1, -1 do
        rev[#rev + 1] = top[i]
      end
      top = rev

      for i = 1, #top do
        table.insert(toStack, top[i])
      end
    end
  end
  local result = ''
  for i = 1, #stacks do
    local stack = stacks[i]
    result = result .. stack[#stack]
  end
  return result
end

print('Part 2:', part2(realInstructions, realStacks))
