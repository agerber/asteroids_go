package model

import (
	"container/list"
	"github.com/agerber/asteroids_go/common"
	"github.com/agerber/asteroids_go/model/prime"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"math"
)

const (
	TURN_STEP          = 5
	INITIAL_SPAWN_TIME = 48
	MAX_SHIELD         = 200
	MAX_NUKE           = 600
	MIN_RADIUS         = 28
)

type ImageState int

const (
	FALCON_INVISIBLE  ImageState = iota //for pre-spawning
	FALCON                              //normal ship
	FALCON_THR                          //normal ship thrusting
	FALCON_SHIELD                       //shielded ship (cyan)
	FALCON_SHIELD_THR                   //shielded ship (cyan) thrusting
)

type Falcon struct {
	*Sprite

	shield           int
	nukeMeter        int
	invisible        int
	maxSpeedAttained bool
	showLevel        int
	turnState        common.TurnState
	thrusting        bool
}

func NewFalcon() *Falcon {
	falcon := &Falcon{
		Sprite:    NewSprite(),
		turnState: common.IDLE,
	}

	falcon.team = common.FRIEND
	falcon.radius = MIN_RADIUS

	falcon.rasterMap = make(map[interface{}]*ebiten.Image)
	falcon.rasterMap[FALCON_INVISIBLE] = nil
	falcon.rasterMap[FALCON] = ebiten.NewImageFromImage(common.GetImage("assets/imgs/fal/falcon125.png"))                       //normal ship
	falcon.rasterMap[FALCON_THR] = ebiten.NewImageFromImage(common.GetImage("assets/imgs/fal/falcon125_thr.png"))               //normal ship thrusting
	falcon.rasterMap[FALCON_SHIELD] = ebiten.NewImageFromImage(common.GetImage("assets/imgs/fal/falcon125_SHIELD.png"))         //SHIELD
	falcon.rasterMap[FALCON_SHIELD_THR] = ebiten.NewImageFromImage(common.GetImage("assets/imgs/fal/falcon125_SHIELD_thr.png")) //S+THR

	return falcon
}

func (f *Falcon) Move() {
	if !common.GetCommandCenterInstance().IsFalconPositionFixed() {
		f.move(f)
		return
	}

	if f.invisible > 0 {
		f.invisible--
	}
	if f.shield > 0 {
		f.shield--
	}
	if f.nukeMeter > 0 {
		f.nukeMeter--
	}
	if f.showLevel > 0 {
		f.showLevel--
	}

	thrust := 0.85
	maxVelocity := 39

	if f.thrusting {
		vectorX := math.Cos(f.orientation*math.Pi/180) * thrust
		vectorY := math.Sin(f.orientation*math.Pi/180) * thrust

		absVelocity := int(math.Sqrt(math.Pow(f.deltaX+vectorX, 2) + math.Pow(f.deltaY+vectorY, 2)))

		if absVelocity < maxVelocity {
			f.deltaX += vectorX
			f.deltaY += vectorY
			f.radius = MIN_RADIUS + absVelocity/3
			f.maxSpeedAttained = false
		} else {
			f.maxSpeedAttained = true
		}
	}

	switch f.turnState {
	case common.LEFT:
		if f.orientation <= 0 {
			f.orientation = 360 - TURN_STEP
		} else {
			f.orientation -= TURN_STEP
		}
	case common.RIGHT:
		if f.orientation >= 360 {
			f.orientation = TURN_STEP
		} else {
			f.orientation += TURN_STEP
		}
	case common.IDLE:
	default:
	}
}

func (f *Falcon) Draw(screen *ebiten.Image) {
	// TODO: remove this
	f.nukeMeter = 1
	f.invisible = 0
	f.shield = 1

	if f.nukeMeter > 0 {
		f.drawNukeHalo(screen)
	}

	var imageState ImageState
	if f.invisible > 0 {
		imageState = FALCON_INVISIBLE
	} else if f.shield > 0 {
		if f.thrusting {
			imageState = FALCON_SHIELD_THR
		} else {
			imageState = FALCON_SHIELD
		}
		f.drawShieldHalo(screen)
	} else {
		if f.thrusting {
			imageState = FALCON_THR
		} else {
			imageState = FALCON
		}
	}

	f.renderRaster(screen, f.rasterMap[imageState])
}

func (f *Falcon) GetCenter() prime.Point {
	return f.center
}

func (f *Falcon) GetRadius() int {
	return f.radius
}

func (f *Falcon) GetTeam() common.Team {
	return f.team
}

func (f *Falcon) AddToGame(list *list.List) {
	f.addToGame(list, f)
}

func (f *Falcon) RemoveFromGame(list *list.List) {
	// The falcon is never actually removed from the game-space; instead we decrement numFalcons
	// only execute the decrementFalconNumAndSpawn() method if shield is down.
	if f.shield == 0 {
		f.DecrementFalconNumAndSpawn()
	}
}

func (f *Falcon) GetDeltaX() float64 {
	return f.deltaX
}

func (f *Falcon) GetDeltaY() float64 {
	return f.deltaY
}

func (f *Falcon) drawShieldHalo(screen *ebiten.Image) {
	vector.StrokeCircle(screen, float32(f.center.X), float32(f.center.Y), float32(f.radius), 1, CyanColor, false)
}

func (f *Falcon) drawNukeHalo(screen *ebiten.Image) {
	if f.invisible > 0 {
		return
	}

	vector.StrokeCircle(screen, float32(f.center.X), float32(f.center.Y), float32(f.radius)-10, 1, YellowColor, false)
}

func (f *Falcon) DecrementFalconNumAndSpawn() {
	common.GetCommandCenterInstance().SetNumFalcons(common.GetCommandCenterInstance().GetNumFalcons() - 1)
	if common.GetCommandCenterInstance().IsGameOver() {
		return
	}
	common.PlaySound("shipspawn.wav")
	f.shield = INITIAL_SPAWN_TIME
	f.invisible = INITIAL_SPAWN_TIME / 5
	f.orientation = common.GenerateRandomFloat64(360/TURN_STEP) * TURN_STEP
	f.deltaX = 0
	f.deltaY = 0
	f.radius = MIN_RADIUS
	f.maxSpeedAttained = false
	f.nukeMeter = 0
}

func (f *Falcon) SetThrusting(thrusting bool) {
	f.thrusting = thrusting
}

func (f *Falcon) SetTurnState(turnState common.TurnState) {
	f.turnState = turnState
}
