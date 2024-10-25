package prime

type AspectRatio struct {
	Width  float64
	Height float64
}

func (a *AspectRatio) Scale(scale float64) {
	a.Width *= scale
	a.Height *= scale
}
