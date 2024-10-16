package model

import (
	"image"
	"image/color"
	"math"
	"math/rand"
	"sync"
	"time"

	"github.com/fogleman/gg"
)

// Team enumeration
type Team int

const (
	Friend Team = iota
	Foe
	Floater
	Debris
)

// Point alias for image.Point
type Point = image.Point

// PolarPoint represents a point in polar coordinates
type PolarPoint struct {
	R     float64
	Theta float64
}

// Sprite struct represents a game sprite
type Sprite struct {
	Center      Point
	DeltaX      float64
	DeltaY      float64
	Team        Team
	Radius      int
	Orientation int
	Expiry      int
	Spin        int
	Cartesians  []Point
	Color       color.Color
	RasterMap   map[interface{}]image.Image
}

// NewSprite constructs a new Sprite with a random position
func NewSprite() *Sprite {
	s := &Sprite{}
	s.Center = Point{
		X: RandInt(0, GameInstance().DIM.Width),
		Y: RandInt(0, GameInstance().DIM.Height),
	}
	return s
}

// Move updates the sprite's position and handles screen wrapping
func (s *Sprite) Move() {
	scalarX := CommandCenterInstance().UniDim.Width
	scalarY := CommandCenterInstance().UniDim.Height

	// Screen wrapping logic
	if s.Center.X > scalarX*GameInstance().DIM.Width {
		s.Center.X = 1
	} else if s.Center.X < 0 {
		s.Center.X = scalarX*GameInstance().DIM.Width - 1
	} else if s.Center.Y > scalarY*GameInstance().DIM.Height {
		s.Center.Y = 1
	} else if s.Center.Y < 0 {
		s.Center.Y = scalarY*GameInstance().DIM.Height - 1
	} else {
		newXPos := float64(s.Center.X)
		newYPos := float64(s.Center.Y)
		if CommandCenterInstance().IsFalconPositionFixed {
			newXPos -= CommandCenterInstance().Falcon.DeltaX
			newYPos -= CommandCenterInstance().Falcon.DeltaY
		}
		s.Center.X = int(math.Round(newXPos + s.DeltaX))
		s.Center.Y = int(math.Round(newYPos + s.DeltaY))
	}

	if s.Expiry > 0 {
		s.Expire()
	}

	if s.Spin != 0 {
		s.Orientation += s.Spin
	}
}

// Expire decrements the sprite's expiry and enqueues it for removal if necessary
func (s *Sprite) Expire() {
	if s.Expiry == 1 {
		CommandCenterInstance().OpsQueue.Enqueue(s, Remove)
	}
	s.Expiry--
}

// SomePosNegValue generates a random positive or negative integer
func (s *Sprite) SomePosNegValue(seed int) int {
	randomNumber := RandInt(0, seed)
	if randomNumber%2 == 0 {
		return randomNumber
	}
	return -randomNumber
}

// RenderRaster renders the sprite in raster mode using the gg library
func (s *Sprite) RenderRaster(dc *gg.Context, bufferedImage image.Image) {
	if bufferedImage == nil {
		return
	}

	centerX := float64(s.Center.X)
	centerY := float64(s.Center.Y)
	width := float64(s.Radius * 2)
	height := float64(s.Radius * 2)
	angleRadians := math.Pi * float64(s.Orientation) / 180.0

	iw := float64(bufferedImage.Bounds().Dx())
	ih := float64(bufferedImage.Bounds().Dy())

	scaleX := width / iw
	scaleY := height / ih

	dc.Push()
	defer dc.Pop()

	dc.Translate(centerX, centerY)
	dc.Scale(scaleX, scaleY)
	if angleRadians != 0 {
		dc.Rotate(angleRadians)
	}
	dc.Translate(-iw/2, -ih/2)
	dc.DrawImage(bufferedImage, 0, 0)
}

// RenderVector renders the sprite in vector mode using the gg library
func (s *Sprite) RenderVector(dc *gg.Context) {
	dc.SetColor(s.Color)

	polars := cartesiansToPolars(s.Cartesians)

	rotatedPolars := make([]PolarPoint, len(polars))
	for i, pp := range polars {
		rotatedPolars[i] = PolarPoint{
			R:     pp.R,
			Theta: pp.Theta + math.Pi*float64(s.Orientation)/180.0,
		}
	}

	dc.NewSubPath()
	for i, pp := range rotatedPolars {
		x := pp.R * float64(s.Radius) * math.Sin(pp.Theta)
		y := pp.R * float64(s.Radius) * math.Cos(pp.Theta)
		x += float64(s.Center.X)
		y = float64(s.Center.Y) - y
		if i == 0 {
			dc.MoveTo(x, y)
		} else {
			dc.LineTo(x, y)
		}
	}
	dc.ClosePath()
	dc.Stroke()
}

// AddToGame adds the sprite to the game
func (s *Sprite) AddToGame(list *[]Movable) {
	*list = append(*list, s)
}

// RemoveFromGame removes the sprite from the game
func (s *Sprite) RemoveFromGame(list *[]Movable) {
	for i, m := range *list {
		if m == s {
			*list = append((*list)[:i], (*list)[i+1:]...)
			break
		}
	}
}

// Draw is an abstract method to be implemented by subclasses
func (s *Sprite) Draw(dc *gg.Context) {
	panic("Draw method not implemented")
}

// cartesiansToPolars converts cartesian points to polar points
func cartesiansToPolars(points []Point) []PolarPoint {
	polars := make([]PolarPoint, len(points))
	for i, p := range points {
		r := math.Hypot(float64(p.X), float64(p.Y))
		theta := math.Atan2(float64(p.Y), float64(p.X))
		polars[i] = PolarPoint{R: r, Theta: theta}
	}
	return polars
}

// Random number generator setup
var randSrc = rand.New(rand.NewSource(time.Now().UnixNano()))

// RandInt generates a random integer between min and max
func RandInt(min, max int) int {
	return min + randSrc.Intn(max-min+1)
}

// Dimension represents the game's dimensions
type Dimension struct {
	Width, Height int
}

// Game struct represents the game state
type Game struct {
	DIM Dimension
}

var gameInstance *Game
var gameOnce sync.Once

// GameInstance returns a singleton instance of Game
func GameInstance() *Game {
	gameOnce.Do(func() {
		gameInstance = &Game{
			DIM: Dimension{Width: 800, Height: 600},
		}
	})
	return gameInstance
}

// Action enumeration for game operations
type Action int

const (
	Add Action = iota
	Remove
)

// GameOp represents a game operation
type GameOp struct {
	Movable Movable
	Action  Action
}

// OpsQueue manages game operations
type OpsQueue struct {
	queue []GameOp
	mutex sync.Mutex
}

// Enqueue adds a game operation to the queue
func (oq *OpsQueue) Enqueue(m Movable, action Action) {
	oq.mutex.Lock()
	defer oq.mutex.Unlock()
	oq.queue = append(oq.queue, GameOp{Movable: m, Action: action})
}

// CommandCenter manages global game state
type CommandCenter struct {
	UniDim                Dimension
	Falcon                *Sprite
	IsFalconPositionFixed bool
	OpsQueue              *OpsQueue
}

var commandCenterInstance *CommandCenter
var commandCenterOnce sync.Once

// CommandCenterInstance returns a singleton instance of CommandCenter
func CommandCenterInstance() *CommandCenter {
	commandCenterOnce.Do(func() {
		commandCenterInstance = &CommandCenter{
			UniDim:                Dimension{Width: 1, Height: 1},
			Falcon:                &Sprite{},
			IsFalconPositionFixed: false,
			OpsQueue:              &OpsQueue{},
		}
	})
	return commandCenterInstance
}

// Movable interface represents movable game objects
type Movable interface {
	Move()
	Draw(dc *gg.Context)
	AddToGame(list *[]Movable)
	RemoveFromGame(list *[]Movable)
}
