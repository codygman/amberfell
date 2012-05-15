/*
  To the extent possible under law, Ian Davis has waived all copyright
  and related or neighboring rights to this Amberfell Source Code file.
  This work is published from the United Kingdom. 
*/
package main

const (
	// GroundLevel = 807

	CHUNK_WIDTH  = 8
	CHUNK_HEIGHT = 128

	MAP_DIAM       = 64000
	PLAYER_START_X = 32122 //32035 //31972 //31768 //31559 //31445 //32011 // 31767 //MAP_DIAM / 2
	PLAYER_START_Z = 31029 //31426 //31450 //31621 //31488 //32137 //31058 // 32009 //MAP_DIAM / 2
	NOISE_SCALE    = 16

	TREE_PRECIPITATION_MIN = 0.2 //0.44
	TREE_DENSITY_PCT       = 5

	BUSH_PRECIPITATION_MIN = 0.1 //0.44
	BUSH_DENSITY_PCT       = 1

	VERTEX_BUFFER_CAPACITY = 20000

	MAX_ITEMS_IN_INVENTORY = 999

	XAXIS = 0
	YAXIS = 1
	ZAXIS = 2

	TILES_HORZ = 16
	TILES_VERT = 8

	TILE_WIDTH   = 48                 // Height and width of a block texture in pixels
	SCREEN_SCALE = 1.0 * TILE_WIDTH   // Width of one world coordinate unit in pixels
	PIXEL_SCALE  = 1.0 / SCREEN_SCALE // Width of one pixel in world coordinate units

	KEY_DEBOUNCE_DELAY = 3e8 // nanoseconds

	CAMPFIRE_INTENSITY = 6
	MAX_LIGHT_LEVEL    = 12

	CAMPFIRE_DURATION = 10

	AMBERFELL_UNITS_PER_SECOND_UNPOWERED = 0.01
	AMBERFELL_UNITS_PER_SECOND_POWERED   = 0.2
	AMBERFELL_PUMP_CAPACITY              = 5
)

const (
	EAST_FACE  = iota // +ve x
	WEST_FACE         // -ve x
	NORTH_FACE        // -ve z
	SOUTH_FACE        // +ve z
	UP_FACE           // +ve y
	DOWN_FACE         // -ve y

	DIR_NE
	DIR_SE
	DIR_SW
	DIR_NW
	DIR_UN
	DIR_UE
	DIR_US
	DIR_UW
	DIR_DN
	DIR_DE
	DIR_DS
	DIR_DW
	FACE_NONE
)

const (
	ORIENT_EAST = iota
	ORIENT_NORTH
	ORIENT_SOUTH
	ORIENT_WEST
)

const (
	BLOCK_AIR = iota
	BLOCK_STONE
	BLOCK_DIRT
	BLOCK_TRUNK
	BLOCK_LEAVES
	BLOCK_LOG_WALL
	BLOCK_LOG_SLAB

	BLOCK_AMBERFELL_SOURCE
	BLOCK_COAL_SEAM
	BLOCK_COPPER_SEAM
	BLOCK_IRON_SEAM
	BLOCK_CARVED_STONE
	BLOCK_CAMPFIRE
	BLOCK_BURNT_GRASS
	BLOCK_BUSH

	BLOCK_AMBERFELL_PUMP
	BLOCK_STEAM_GENERATOR
	BLOCK_CARPENTERS_BENCH
	BLOCK_MAGNETITE_SEAM

	BLOCK_PLANK_WALL
	BLOCK_PLANK_SLAB
	BLOCK_AMBERFELL_CONDENSER
)

const (
	ITEM_FIREWOOD = iota + 512
	ITEM_PLANK
	ITEM_COAL
	ITEM_IRON_ORE
	ITEM_COPPER_ORE
	ITEM_IRON_INGOT
	ITEM_COPPER_INGOT
	ITEM_AMBERFELL
	ITEM_MAGNETITE_ORE
	ITEM_LODESTONE
	ITEM_AMBERFELL_CRYSTAL
)

const (
	MAX_ITEMS = 4096
	ITEM_NONE = MAX_ITEMS - 1
)

const (
	BEHAVIOUR_WANDER = iota
	BEHAVIOUR_SEPARATE
	BEHAVIOUR_GATHER
	BEHAVIOUR_ALIGN
	BEHAVIOUR_PURSUE
	BEHAVIOUR_EVADE
)

const (
	TARGET_ANY = iota
	TARGET_PLAYER
	TARGET_WOLF
	TARGET_CAMPFIRE
	TARGET_AMBERFELL_SOURCE
)

const (
	ACTION_HAND = iota
	ACTION_BREAK
	ACTION_WEAPON
	ACTION_ITEM0
	ACTION_ITEM1
	ACTION_ITEM2
	ACTION_ITEM3
	ACTION_ITEM4
)

// Textures in tile.png
const (

	// Row 1 (0-15)
	TEXTURE_NONE = iota
	TEXTURE_STONE_TOP
	TEXTURE_DIRT_TOP
	TEXTURE_TRUNK_TOP
	TEXTURE_FIRE
	TEXTURE_LOG_WALL_TOP
	TEXTURE_AMBERFELL_SOURCE_TOP
	TEXTURE_COAL
	TEXTURE_COPPER
	TEXTURE_IRON
	TEXTURE_BURNT_GRASS
	_ // 11
	_ // 12
	_ // 13
	_ // 14
	_ // 15

	// Row 2  (16-31)
	TEXTURE_CARVED_STONE
	TEXTURE_STONE
	TEXTURE_DIRT
	TEXTURE_TRUNK
	TEXTURE_LEAVES
	TEXTURE_LOG_WALL
	TEXTURE_AMBERFELL_SOURCE
	_ // 23
	_ // 24
	_ // 25
	_ // 26
	_ // 27
	_ // 28
	_ // 29
	_ // 30
	_ // 31

	// Row 3 (32-47)
	TEXTURE_COPPER_MACH_SIDE
	TEXTURE_COPPER_MACH_TOP
	_ // 34
	_ // 35
	_ // 36
	_ // 36
	_ // 37
	_ // 38
	_ // 39
	_ // 40
	_ // 41
	_ // 42
	_ // 43
	_ // 44
	_ // 45
	_ // 46
	_ // 47

	// Row 4  (48-63)
	TEXTURE_IRON_MACH_SIDE
	TEXTURE_IRON_MACH_TOP
	_ // 50
	_ // 51
	_ // 52
	_ // 53
	_ // 54
	_ // 55
	_ // 56
	_ // 57
	_ // 58
	_ // 59
	_ // 60
	_ // 61
	_ // 62
	_ // 63

	// Row 5  (64-79)
	TEXTURE_PLANK_WALL
	TEXTURE_CARPENTERS_BENCH_TOP
	_ // 66
	_ // 67
	_ // 68
	_ // 69
	_ // 70
	_ // 71
	_ // 72
	_ // 73
	_ // 74
	_ // 75
	_ // 76
	_ // 77
	_ // 78
	_ // 79

	// Row 6  (80-95)
	TEXTURE_ACTION_HAND
	TEXTURE_ACTION_BREAK
	TEXTURE_ACTION_WEAPON
	TEXTURE_ITEM_FIREWOOD
	TEXTURE_ITEM_PLANK
	TEXTURE_ITEM_COAL
	TEXTURE_ITEM_IRON_ORE
	TEXTURE_ITEM_COPPER_ORE
	TEXTURE_ITEM_IRON_INGOT
	TEXTURE_ITEM_COPPER_INGOT
	TEXTURE_ITEM_AMBERFELL
	TEXTURE_ITEM_LODESTONE
	TEXTURE_ITEM_AMBERFELL_CRYSTAL
	_
	_
	_
	_
	_

	// Row 7  (96-111)

)

const (
	// Player textures

	TEXTURE_HAT_FRONT = 4096
	TEXTURE_HAT_LEFT  = 4097
	TEXTURE_HAT_BACK  = 4098
	TEXTURE_HAT_RIGHT = 4099
	TEXTURE_HAT_TOP   = 4100

	TEXTURE_HEAD_FRONT  = 4101
	TEXTURE_HEAD_LEFT   = 4102
	TEXTURE_HEAD_BACK   = 4103
	TEXTURE_HEAD_RIGHT  = 4104
	TEXTURE_HEAD_BOTTOM = 4105

	TEXTURE_TORSO_FRONT = 4106
	TEXTURE_TORSO_LEFT  = 4107
	TEXTURE_TORSO_BACK  = 4108
	TEXTURE_TORSO_RIGHT = 4109
	TEXTURE_TORSO_TOP   = 4110

	TEXTURE_LEG      = 4111
	TEXTURE_LEG_SIDE = 4112
	TEXTURE_ARM      = 4113
	TEXTURE_ARM_TOP  = 4114
	TEXTURE_HAND     = 4116
	TEXTURE_BRIM     = 4117

	// Mob textures

	// HUD textures
	TEXTURE_PICKER = 5000

	// Strength of materials
	STRENGTH_STONE       = 20
	STRENGTH_DIRT        = 3
	STRENGTH_WOOD        = 5
	STRENGTH_LEAVES      = 1
	STRENGTH_IRON        = 50
	STRENGTH_UNBREAKABLE = 255
)

var (
	SUNLIGHT_LEVELS = [8]float64{0.02, 0.05, 0.08, 0.11, 0.15, 0.20, 0.30, 0.45}

	NORMALS = [6]([3]float32){[3]float32{1.0, 0.0, 0.0},
		[3]float32{-1.0, 0.0, 0.0},
		[3]float32{0.0, 0.0, -1.0},
		[3]float32{0.0, 0.0, 1.0},
		[3]float32{0.0, 1.0, 0.0},
		[3]float32{0.0, -1.0, 0.0},
	}

	LIGHT_LEVELS = [13]([4]float32){[4]float32{0, 0, 0, 1},
		[4]float32{0.04, 0.04, 0.04, 1.0},
		[4]float32{0.12, 0.12, 0.12, 1.0},
		[4]float32{0.20, 0.20, 0.20, 1.0},
		[4]float32{0.28, 0.28, 0.28, 1.0},
		[4]float32{0.36, 0.36, 0.36, 1.0},
		[4]float32{0.44, 0.44, 0.44, 1.0},
		[4]float32{0.52, 0.52, 0.52, 1.0},
		[4]float32{0.60, 0.60, 0.60, 1.0},
		[4]float32{0.68, 0.68, 0.68, 1.0},
		[4]float32{0.76, 0.76, 0.76, 1.0},
		[4]float32{0.84, 0.84, 0.84, 1.0},
		[4]float32{0.92, 0.92, 0.92, 1.0},
	}
	COLOUR_WHITE = [4]float32{1.0, 1.0, 1.0, 1.0}
	COLOURS      = [5][4]float32{
		{1.0, 1.0, 1.0, 1.0},
		{0.7, 0.7, 0.7, 1.0},
		{0.55, 0.55, 0.55, 1.0},
		{0.4, 0.4, 0.4, 1.0},
		{0.2, 0.2, 0.2, 1.0},
	}

	COLOUR_HIGH = [4]float32{96.0 / 255, 208.0 / 255, 96.0 / 255, 1.0}
)
