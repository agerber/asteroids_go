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
	ORIGINAL_FRAMES_PER_SECOND = 25
	GOLANG_FRAMES_PER_SECOND   = 25
	GOLANG_FRAMES_SCALE_FACTOR = float64(GOLANG_FRAMES_PER_SECOND) / float64(ORIGINAL_FRAMES_PER_SECOND)
)
