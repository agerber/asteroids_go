package config

type Dimension struct {
	Width  int
	Height int
}

var DIM = Dimension{
	Width:  1500,
	Height: 950,
}

const (
	WINDOW_TITLE         = "Game Base"
	ANIMATION_DELAY      = 40
	FRAMES_PER_SECOND    = 1000 / ANIMATION_DELAY
	LARGE_RADIUS         = 110
	MAX_RADIANS_X1000    = 6283
	PRECISION            = 1000.0
	SPAWN_SHIELD_FLOATER = FRAMES_PER_SECOND * 25
	SPAWN_NUKE_FLOATER   = FRAMES_PER_SECOND * 12
)
