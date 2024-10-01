package prime

type AspectRatio struct {
	Width  float64
	Height float64
}

func NewAspectRatio(width, height float64) *AspectRatio {
	return &AspectRatio{
		Width:  width,
		Height: height,
	}
}

func (ar *AspectRatio) Scale(scale float64) *AspectRatio {
	ar.Width *= scale
	ar.Height *= scale
	return ar
}
