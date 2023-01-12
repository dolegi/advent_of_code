local inspect = require 'inspect'
local readFile = require 'readFile'

local exampleInput = [[R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
]]

local realInput = readFile('input.txt')

local function parseInput(input)
  local result = {}
  for line in input:gmatch("[^\r\n]+") do
    local direction, distance = line:match("(%a) (%d+)")
    result[#result + 1] = { direction, tonumber(distance) }
  end
  return result
end

local function areNeighbours(a, b)
  if a.x == b.x and a.y == b.y then return true end
  if a.x == b.x - 1 and a.y == b.y then return true end
  if a.x == b.x + 1 and a.y == b.y then return true end
  if a.x == b.x and a.y == b.y + 1 then return true end
  if a.x == b.x and a.y == b.y - 1 then return true end
  if a.x == b.x - 1 and a.y == b.y + 1 then return true end
  if a.x == b.x + 1 and a.y == b.y + 1 then return true end
  if a.x == b.x + 1 and a.y == b.y - 1 then return true end
  if a.x == b.x - 1 and a.y == b.y - 1 then return true end

  return false
end

local function arrContains(arr, item)
  for _, v in ipairs(arr) do
    if v == item then
      return true
    end
  end
  return false
end

local parsed = parseInput(realInput)

print(inspect(parsed))

local H = { x = 1, y = 1 }
local T = { x = 1, y = 1 }

local function move(val, direction, distance)
  if direction == "R" then
    val.x = val.x + distance
  elseif direction == "L" then
    val.x = val.x - distance
  elseif direction == "U" then
    val.y = val.y + distance
  elseif direction == "D" then
    val.y = val.y - distance
  end
end

local visited = { '1|1' }
local function emptyGrid()
  return {
    { '.', '.', '.', '.', '.', '.' },
    { '.', '.', '.', '.', '.', '.' },
    { '.', '.', '.', '.', '.', '.' },
    { '.', '.', '.', '.', '.', '.' },
    { '.', '.', '.', '.', '.', '.' }
  }
end

local grid = {
  { '.', '.', '.', '.', '.', '.' },
  { '.', '.', '.', '.', '.', '.' },
  { '.', '.', '.', '.', '.', '.' },
  { '.', '.', '.', '.', '.', '.' },
  { 'H', '.', '.', '.', '.', '.' }
}

for i = 1, #grid do
  print(inspect(grid[i]))
end
print('')

for _, v in ipairs(parsed) do
  local direction, distance = v[1], v[2]

  while distance > 0 do
    move(H, direction, 1)
    if not areNeighbours(H, T) then
      if direction == 'R' then
        T.x = H.x - 1
        T.y = H.y
      elseif direction == 'L' then
        T.x = H.x + 1
        T.y = H.y
      elseif direction == 'U' then
        T.x = H.x
        T.y = H.y - 1
      elseif direction == 'D' then
        T.x = H.x
        T.y = H.y + 1
      end
      -- move(T, direction, 1) -- this is wrong, could be a diagonal

      local position = tostring(T.x) .. '|' .. tostring(T.y)
      if not arrContains(visited, position) then
        table.insert(visited, position)
      end
    end
    distance = distance - 1
    --
    -- grid = emptyGrid()
    -- grid[T.y][T.x] = 'T'
    -- grid[H.y][H.x] = 'H'
    --
    -- print(v[1], v[2])
    -- for i = #grid, 1, -1 do
    --   print(inspect(grid[i]))
    -- end
    -- print('')
  end
end

print(#visited)
