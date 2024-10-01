package prime

type Color struct {
	R, G, B int
}

func (c Color) RGB() (int, int, int) {
	return c.R, c.G, c.B
}

func FromRGB(r, g, b int) Color {
	return Color{r, g, b}
}

// Predefined static colors
var (
	Red    = FromRGB(255, 0, 0)
	Green  = FromRGB(0, 255, 0)
	Blue   = FromRGB(0, 0, 255)
	Cyan   = FromRGB(0, 200, 200)
	Orange = FromRGB(255, 165, 0)
	White  = FromRGB(255, 255, 255)
	Black  = FromRGB(0, 0, 0)
	Gray   = FromRGB(128, 128, 128)
	Yellow = FromRGB(255, 255, 0)
)
