-- global.lua

-- module initialize
local M = {}

-- constants
M.SCREEN_WIDTH = 1280
M.SCREEN_HEIGHT = 720
M.LEVEL_COLORS = {
	vmath.vector4(0.4, 1.0, 0.4, 1.0), -- 001
	vmath.vector4(0.4, 0.4, 1.0, 1.0), -- 002
	vmath.vector4(1.0, 0.4, 0.4, 1.0), -- 003
	vmath.vector4(0.8, 0.8, 0.5, 1.0), -- 004
	vmath.vector4(0.5, 0.8, 0.8, 1.0), -- 005
	vmath.vector4(0.8, 0.5, 0.8, 1.0), -- 006
}

-- variables
M.selected_level = 1
M.max_unlucked_level = 6
M.gold = 0
M.eggs = 0
M.player_position = vmath.vector3(200, 360, 0.8)
M.last_battle_result = {
	win = true,
	time = 0,
}

-- module return
return M
