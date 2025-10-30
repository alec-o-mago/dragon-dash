-- global.lua

-- module initialize
local M = {}

-- constants
M.SCREEN_WIDTH = 1280
M.SCREEN_HEIGHT = 720
M.TOTAL_BOSSES = 7
M.STARTING_BOSS_HEALTH = 100
M.STARTING_BOSS_POSITION = vmath.vector3(1060, 360, 0)
M.BOSS_Y_POSITIONS = {128, 360, 592}
M.LEVEL_COLORS = {
	vmath.vector4(1.0, 0.6, 0.6, 1.0), -- red
	vmath.vector4(0.6, 1.0, 0.6, 1.0), -- green
	vmath.vector4(0.6, 0.6, 1.0, 1.0), -- blue
	vmath.vector4(0.9, 0.9, 0.5, 1.0), -- 
	vmath.vector4(0.9, 0.5, 0.9, 1.0), -- 
	vmath.vector4(0.5, 0.9, 0.9, 1.0), -- 
	vmath.vector4(0.7, 0.7, 0.8, 1.0), -- gray
}

-- variables
M.selected_level = 1
M.max_unlucked_level = 2
M.player_hearts = 5
M.gold = 0
M.eggs = 3
M.player_position = vmath.vector3(200, 360, 0)
M.upgrade_levels = {
	main_power = 1,
	pet_power = 1,
	hearts = 1,
}
M.unlocked_skins = {
	false, false, false, false, false, false, false,
	false, false, false, false, false, false, false,
	false, false, false, false, false, false, false,
}
M.last_battle_result = { -- Example victory data
	bosses_defeated = 1,
	time = 120, -- time in seconds
}
M.config_music = true
M.config_sound = true

-- module return
return M
