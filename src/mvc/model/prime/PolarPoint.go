package prime

// PolarPoint represents a point in polar coordinates.
type PolarPoint struct {
	R     float64 // Radius (between 0 and 1)
	Theta float64 // Angle in radians (between 0 and 2*pi)
}

// NewPolarPoint creates a new PolarPoint instance.
func NewPolarPoint(r, theta float64) PolarPoint {
	return PolarPoint{
		R:     r,
		Theta: theta,
	}
}

// CompareTheta compares the theta values of two PolarPoints.
func (p *PolarPoint) CompareTheta(other PolarPoint) int {
	return compareFloat64(p.Theta, other.Theta)
}

// compareFloat64 compares two float64 values for sorting.
func compareFloat64(a, b float64) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}
