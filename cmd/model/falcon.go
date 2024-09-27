package model

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"math"
	"path/filepath"
	"sync"
)

// Falcon struct represents the player's ship
type Falcon struct {
	X, Y           float64 // Position of the Falcon
	DeltaX, DeltaY float64 // Velocity in X and Y
	Orientation    float64 // Rotation in degrees
	Shield         int     // Shield level (0-100)
	NukeMeter      int     // Nuke charge level (0-100)
	Thrusting      bool    // If the Falcon is thrusting
	TurnState      TurnState
	Sprite         *ebiten.Image
	SpriteThrust   *ebiten.Image
}

type TurnState int

const (
	IDLE TurnState = iota
	LEFT
	RIGHT
)

const (
	TurnStep    = 11
	ThrustPower = 0.85
	MaxSpeed    = 8.0
)

type ImageMap map[string]*ebiten.Image

var (
	images ImageMap
	once   sync.Once
)

// LoadImages function
func LoadImages() {
	var err error
	images = make(ImageMap)

	imagePaths := map[string]string{
		"falcon":     filepath.Join("resources", "imgs", "fal", "falcon125.png"),
		"falcon_thr": filepath.Join("resources", "imgs", "fal", "falcon125_thr.png"),
	}

	for key, path := range imagePaths {
		images[key], _, err = ebitenutil.NewImageFromFile(path)
		if err != nil {
			log.Fatalf("Error loading image %s: %v", path, err)
		}
	}
}

// NewFalcon creates and initializes a new Falcon instance
func NewFalcon() *Falcon {
	once.Do(LoadImages)

	return &Falcon{
		X:            320, // Starting position (center of screen)
		Y:            240,
		Shield:       100,
		NukeMeter:    0,
		TurnState:    IDLE,
		Sprite:       images["falcon"],
		SpriteThrust: images["falcon_thr"],
	}
}

// Move updates the Falcon's position and orientation
func (f *Falcon) Move() {
	const Friction float64 = 0.95
	if f.Thrusting {
		// Calculate thrust direction based on current orientation
		vectorX := math.Cos(f.Orientation * math.Pi / 180)
		vectorY := math.Sin(f.Orientation * math.Pi / 180)

		// Apply thrust to velocity
		f.DeltaX += vectorX * ThrustPower
		f.DeltaY += vectorY * ThrustPower

		// Limit velocity to max speed
		speed := math.Hypot(f.DeltaX, f.DeltaY)
		if speed > MaxSpeed {
			f.DeltaX *= MaxSpeed / speed
			f.DeltaY *= MaxSpeed / speed
		}
	} else {
		// Apply friction when not thrusting
		f.DeltaX *= Friction
		f.DeltaY *= Friction
	}

	// Update position
	f.X += f.DeltaX
	f.Y += f.DeltaY

	// Wrap around screen edges (assumes screen size is 640x480)
	f.X = math.Mod(f.X+640, 640)
	f.Y = math.Mod(f.Y+480, 480)

	// Turn Falcon based on input
	switch f.TurnState {
	case LEFT:
		f.Orientation -= TurnStep
	case RIGHT:
		f.Orientation += TurnStep
	}
}

// Draw renders the Falcon onto the screen
func (f *Falcon) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(f.Sprite.Bounds().Dx())/2, -float64(f.Sprite.Bounds().Dy())/2)
	op.GeoM.Rotate(f.Orientation * math.Pi / 180)
	op.GeoM.Translate(f.X, f.Y)

	// Choose which sprite to draw based on thrust state
	if f.Thrusting {
		screen.DrawImage(f.SpriteThrust, op)
	} else {
		screen.DrawImage(f.Sprite, op)
	}
}
