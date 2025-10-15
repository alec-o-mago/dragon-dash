-- global.lua

-- module initialize
local M = {}

-- constants
M.SCREEN_WIDTH = 1280
M.SCREEN_HEIGHT = 720

-- variables
M.selected_level = 1
M.max_unlucked_level = 3
M.gold = 0
M.eggs = 0
M.last_battle_result = {
	win = true,
	time = 0,
}

-- module end
return M
