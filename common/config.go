package common

type Dimension struct {
	Width  int
	Height int
}

var DIM = Dimension{
	Width:  1500,
	Height: 950,
}

const (
	WINDOW_TITLE               = "Game Base"
	JAVA_ANIMATION_DELAY       = 40
	JAVA_FRAMES_PER_SECOND     = 1000 / JAVA_ANIMATION_DELAY
	GOLANG_FRAMES_PER_SECOND   = 58
	GOLANG_FRAMES_SCALE_FACTOR = float64(GOLANG_FRAMES_PER_SECOND) / float64(JAVA_FRAMES_PER_SECOND)
	LARGE_RADIUS               = 110
	MAX_RADIANS_X1000          = 6283
	PRECISION                  = 1000.0
	SPAWN_SHIELD_FLOATER       = GOLANG_FRAMES_PER_SECOND * 25
	SPAWN_NUKE_FLOATER         = GOLANG_FRAMES_PER_SECOND * 12
)
