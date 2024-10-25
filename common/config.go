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
	GOLANG_FRAMES_PER_SECOND   = 25
	GOLANG_FRAMES_SCALE_FACTOR = float64(GOLANG_FRAMES_PER_SECOND) / float64(JAVA_FRAMES_PER_SECOND)
)
