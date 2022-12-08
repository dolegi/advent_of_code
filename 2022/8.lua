local inspect = require('inspect')
local readFile = require('readFile')

local exampleInput = [[30373
25512
65332
33549
35390
]]

local realInput = readFile('input.txt')

local function convertToMatrix(input)
  local matrix = {}

  local y = 1
  for line in string.gmatch(input, "[^\r\n]+") do
    matrix[y] = {}
    for tree in string.gmatch(line, "%d") do
      table.insert(matrix[y], tonumber(tree))
    end
    y = y + 1
  end
  return matrix
end

local function transpose(m)
  local res = {}

  for i = 1, #m[1] do
    res[i] = {}
    for j = 1, #m do
      res[i][j] = m[j][i]
    end
  end

  return res
end

local function reverse(arr)
  local rev = {}
  for i = #arr, 1, -1 do
    rev[#rev + 1] = arr[i]
  end
  return rev
end

local function visibleFrom(value, arr)
  for i = 1, #arr do
    if value <= arr[i] then
      return false
    end
  end
  return true
end

local function countVisibleTrees(matrix)
  local transposedMatrix = transpose(matrix)

  local total = 0
  for y = 1, #matrix do
    for x = 1, #matrix[y] do
      if y == 1 or y == #matrix or x == 1 or x == #matrix[y] then
        total = total + 1
        goto continue
      end

      local currentValue = matrix[y][x]

      local leftSide = { table.unpack(matrix[y], 1, x - 1) }
      if visibleFrom(currentValue, leftSide) then
        total = total + 1
        goto continue
      end

      local rightSide = { table.unpack(matrix[y], x + 1, #matrix[y]) }
      if visibleFrom(currentValue, rightSide) then
        total = total + 1
        goto continue
      end

      local topSide = { table.unpack(transposedMatrix[x], 1, y - 1) }
      if visibleFrom(currentValue, topSide) then
        total = total + 1
        goto continue
      end

      local bottomSide = { table.unpack(transposedMatrix[x], y + 1, #transposedMatrix[x]) }
      if visibleFrom(currentValue, bottomSide) then
        total = total + 1
        goto continue
      end

      ::continue::
    end
  end
  return total
end

local function part1(input)
  local matrix = convertToMatrix(input)

  return countVisibleTrees(matrix)
end

local function countVisibility(value, arr)
  local visible = 0
  for i = 1, #arr do
    if arr[i] < value then
      visible = visible + 1
    else
      return visible + 1
    end
  end
  return visible
end

local function part2(input)
  local matrix = convertToMatrix(input)
  local transposedMatrix = transpose(matrix)


  local bestVisibility = 0
  for y = 1, #matrix do
    for x = 1, #matrix[y] do
      if y == 1 or y == #matrix or x == 1 or x == #matrix[y] then
        goto continue
      end
      local currentValue = matrix[y][x]

      local leftSide = { table.unpack(matrix[y], 1, x - 1) }
      local rightSide = { table.unpack(matrix[y], x + 1, #matrix[y]) }
      local topSide = { table.unpack(transposedMatrix[x], 1, y - 1) }
      local bottomSide = { table.unpack(transposedMatrix[x], y + 1, #transposedMatrix[x]) }

      local currentVisibility =
      countVisibility(currentValue, reverse(leftSide)) *
          countVisibility(currentValue, rightSide) *
          countVisibility(currentValue, reverse(topSide)) *
          countVisibility(currentValue, bottomSide)

      if currentVisibility > bestVisibility then
        bestVisibility = currentVisibility
      end
      ::continue::
    end
  end
  return bestVisibility
end

print('Part 1:', part1(realInput))
print('Part 2:', part2(realInput))
