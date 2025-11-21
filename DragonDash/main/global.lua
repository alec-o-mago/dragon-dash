-- global.lua

-- Module initialize
local M = {}

-- Constants
M.SCREEN_WIDTH = 1280
M.SCREEN_HEIGHT = 720
M.MAX_EGGS = 60
M.TOTAL_BOSSES = 7
M.STARTING_SKIN_COST = 100
M.SKIN_COST_INCREMENT = 20
M.STARTING_BOSS_HEALTH = 100
M.BOSS_HEALTH_INCREMENT = 12
M.STARTING_GOLD = 100
M.STARTING_EGGS = 0
M.STARTING_PLAYER_1_SKIN = 1
M.STARTING_PLAYER_2_SKIN = 2
M.STARTING_BOSS_POSITION = vmath.vector3(1060, 360, 0)
M.BOSS_Y_POSITIONS = {128, 360, 592}
M.LEVEL_COLORS = {
	vmath.vector4(0.6, 0.4, 0.4, 1.0), -- red
	vmath.vector4(0.4, 0.6, 0.4, 1.0), -- green
	vmath.vector4(0.4, 0.4, 0.6, 1.0), -- blue
	vmath.vector4(0.6, 0.6, 0.3, 1.0), -- yellow
	vmath.vector4(0.6, 0.3, 0.6, 1.0), -- 
	vmath.vector4(0.3, 0.6, 0.6, 1.0), -- 
	vmath.vector4(0.5, 0.5, 0.5, 1.0), -- gray
}

-- Variables
M.player_hearts = 5
M.gold = 2000
M.eggs = 3
M.player_position = vmath.vector3(200, 360, 0)
M.player_1_selected_skin = 3
M.player_2_selected_skin = 2
M.high_score_bosses_defeated = 0
M.high_score_time = 0
M.last_battle_bosses_defeated = 1
M.last_battle_time = 120 -- Time in seconds
M.config_music = true
M.config_sound = true
M.unlocked_skins = {
	true, true, true, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false, false, false,
}

-- Module return
return M
