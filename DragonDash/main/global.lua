-- global.lua

-- module initialize
local M = {}

-- constants
M.SCREEN_WIDTH = 1280
M.SCREEN_HEIGHT = 720
M.TOTAL_LEVELS = 2
M.LEVEL_COLORS = {
	vmath.vector4(0.7, 0.7, 0.7, 1.0), -- gray
	vmath.vector4(0.6, 0.6, 1.0, 1.0), -- blue
	vmath.vector4(0.6, 1.0, 0.6, 1.0), -- green
	vmath.vector4(1.0, 0.6, 0.6, 1.0), -- red
	vmath.vector4(0.9, 0.9, 0.5, 1.0), -- 
	vmath.vector4(0.6, 0.9, 0.9, 1.0), -- 
	vmath.vector4(0.9, 0.6, 0.9, 1.0), -- 
}
M.BOSS_HEALTH = {
	400, 500, 1500, 2000, 2500, 3000, 3500, 4000, 4500, 5000,
	500, 1000, 1500, 2000, 2500, 3000, 3500, 4000, 4500, 5000,
}

-- variables
M.selected_level = 1
M.max_unlucked_level = 2
M.player_hearts = 3
M.gold = 0
M.eggs = 0
M.player_position = vmath.vector3(200, 360, 0)
M.last_battle_result = {
	win = true,
	time = 1000,
	hearts = 3,
}
M.config_music = true
M.config_sound = true

-- module return
return M
