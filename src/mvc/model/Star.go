package model

import (
	"math/rand"

	"github.com/agerber/asteroids_go/src/mvc/model/prime"
)

// const (
// 	starBrightnessMax = 226
// 	starRadius        = 1
// )

// Star represents a star object in the game.
type Star struct {
	Movable
	center prime.Point
	color  prime.Color
}

// NewStar creates a new Star object.
func NewStar() *Star {
	bright := rand.Intn(226) // Generate random int between 0 (inclusive) and 225 (exclusive)
	return &Star{
		center: prime.Point{X: rand.Intn(prime.DIM.Width) + 1, Y: rand.Intn(prime.DIM.Height) + 1},
		color:  prime.Color{R: bright, G: bright, B: bright},
	}
}

// Move updates the star's position based on the falcon's movement (if not fixed).
// func (s *Star) Move() {
// 	if !CommandCenter.GetInstance().IsFalconPositionFixed() {
// 		// Right-bounds reached
// 		if s.center.X > prime.DIM.Width {
// 			s.center.X = 1
// 		}
// 		// Left-bounds reached
// 		if s.center.X < 0 {
// 			s.center.X = prime.DIM.Width - 1
// 		}
// 		// Bottom-bounds reached
// 		if s.center.Y > prime.DIM.Height {
// 			s.center.Y = 1
// 		}
// 		// Top-bounds reached
// 		if s.center.Y < 0 {
// 			s.center.Y = prime.DIM.Height - 1
// 		} else {
// 			// In-bounds - move in opposite direction of falcon
// 			deltaX := CommandCenter.GetInstance().GetFalcon().GetDeltaX()
// 			deltaY := CommandCenter.GetInstance().GetFalcon().GetDeltaY()
// 			newX := int(math.Round(float64(s.center.X) - deltaX))
// 			newY := int(math.Round(float64(s.center.Y) - deltaY))
// 			s.center.X = newX
// 			s.center.Y = newY
// 		}
// 	}
// }

// Draw draws the star on the provided image.
// func (s *Star) Draw(img image.Image) {
// 	// Assuming the image draw functionality is implemented elsewhere
// 	// You'll need to replace this with your specific implementation
// 	draw := image.NewDraw(img)
// 	radius := s.GetRadius()
// 	centerX, centerY := s.center.X, s.center.Y
// 	draw.Ellipse(image.Rect(centerX-radius, centerY-radius, centerX+radius, centerY+radius), s.color)
// }

// GetRadius returns the radius of the star.
func (s *Star) GetRadius() int {
	return 1
}

// GetTeam returns the team of the star (debris).
func (s *Star) GetTeam() Team {
	return Debris
}

// GetCenter returns the center point of the star.
func (s *Star) GetCenter() prime.Point {
	return s.center
}

// AddToGame is not implemented in Golang as you would likely use slices or maps
// to manage game objects.

// RemoveFromGame is not implemented in Golang as you would likely use slices or maps
// to manage game objects.

// // AddToGame adds the Star to the game world.
// func (s *Star) AddToGame(list []Movable) {
// 	list = append(list, s)
// }

// // RemoveFromGame removes the Star from the game world.
// func (s *Star) RemoveFromGame(list []Movable) {
// 	list = remove(list, s)
// }

// func remove(slice []Movable, element *Star) []Movable {
// 	var i int
// 	for i = range slice {
// 		if slice[i] == element {
// 			slice = append(slice[:i], slice[i+1:]...)
// 			break
// 		}
// 	}
// 	return slice
// }
