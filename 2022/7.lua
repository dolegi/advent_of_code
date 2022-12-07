local inspect = require('inspect')
local readFile = require('readFile')

local exampleInput = [[$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
]]

local realInput = readFile('input.txt')

local function isDir(str)
  return string.match(str, '^dir.*') ~= nil
end

local function isCmd(str)
  return string.match(str, '^$.*') ~= nil
end

local function isCD(str)
  return string.match(str, '^$ cd.*') ~= nil
end

local function isLS(str)
  return string.match(str, '^$ ls.*') ~= nil
end

local function getCDdir(str)
  return string.match(str, '^$ cd (.*)')
end

local function getSpace(str)
  return tonumber(string.match(str, '.-(%d+).-'))
end

local function buildTree(input)
  local dirTree = { root = {} }
  local parents = {}
  local curDir = dirTree['root']
  local insideLS = false

  for line in string.gmatch(input, "[^\r\n]+") do
    if isDir(line) then
      goto continue
    end
    if isCmd(line) then
      insideLS = false
      if isCD(line) then
        local nextDir = getCDdir(line)
        if nextDir == '..' then
          curDir = table.remove(parents)
        else
          table.insert(parents, curDir)
          curDir[nextDir] = {}
          curDir = curDir[nextDir]
        end
      end
      if isLS(line) then
        insideLS = true
      end
    else
      if insideLS then
        if curDir['space'] == nil then
          curDir['space'] = 0
        end
        local space = getSpace(line)
        curDir['space'] = curDir['space'] + space
        for i = 1, #parents do
          local parent = parents[i]

          if parent['space'] == nil then
            parent['space'] = 0
          end

          parent['space'] = parent['space'] + space
        end
      end
    end
    ::continue::
  end
  return dirTree
end

local function findSpaceUnder100K(node)
  local space = 0

  if type(node) == 'number' then
    if node < 100000 then
      space = space + node
    end
  else
    for _, v in pairs(node) do
      space = space + findSpaceUnder100K(v)
    end
  end
  return space
end

local function findSmallestDirUnder(node, limit, currentResult)
  if type(node) == 'number' then
    if node >= limit and node <= currentResult then
      currentResult = node
    end
  else
    for _, v in pairs(node) do
      currentResult = findSmallestDirUnder(v, limit, currentResult)
    end
  end
  return currentResult
end

local function part1(input)
  local tree = buildTree(input)
  return findSpaceUnder100K(tree)
end

local function part2(input)
  local totalSpace = 70000000
  local requiredUnused = 30000000

  local tree = buildTree(input)
  local usedSpace = tree.root.space
  local unusedSpace = totalSpace - usedSpace
  local minimumToDelete = requiredUnused - unusedSpace

  local spaceToRemove = findSmallestDirUnder(tree, minimumToDelete, totalSpace)
  return spaceToRemove
end

print('Part 1:', part1(realInput))
print('Part 2:', part2(realInput))
