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
	WINDOW_TITLE      = "Game Base"
	FRAMES_PER_SECOND = 30
)
