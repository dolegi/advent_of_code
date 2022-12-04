local readFile = require('readFile')

local input = readFile('./input.txt')
local exampleInput = [[A Y
B X
C Z
]]

local function part1()
	local scores = {
		AX = 4;
		AY = 8;
		AZ = 3;
		BX = 1;
		BY = 5;
		BZ = 9;
		CX = 7;
		CY = 2;
		CZ = 6;
	}

	local total = 0
	local index = 1
	for token in string.gmatch(input, "[^\r\n]+") do
		local code = ""
		for play in string.gmatch(token, "%S") do
			code = code .. play
		end
		total = total + scores[code]
		index = index + 1
	end
	return total
end

print('Part 1:', part1())

local function part2()
	local scores = {
		A = { X = 3, Y = 4, Z = 8 };
		B = { X = 1, Y = 5, Z = 9 };
		C = { X = 2, Y = 6, Z = 7 };
	}
	local total = 0

	local index = 1
	for token in string.gmatch(input, "[^\r\n]+") do
		local code = {}
		local index2 = 1
		for play in string.gmatch(token, "%S") do
			code[index2] = play
			index2 = index2 + 1
		end
		local them = code[1]
		local whatToDo = code[2]

		total = total + scores[them][whatToDo]

		index = index + 1
	end
	return total
end

print('Part 2:', part2())
