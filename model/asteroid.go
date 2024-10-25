package model

import (
	"container/list"
	"image/color"
	"sort"

	"github.com/agerber/asteroids_go/common"
	"github.com/agerber/asteroids_go/model/prime"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ASTEROID_LARGE_RADIUS      = 110
	ASTEROID_MAX_RADIANS_X1000 = 6283
	ASTEROID_PRECISION         = 1000.0
)

type Asteroid struct {
	*Sprite
}

func NewAsteroid(size int) *Asteroid {
	asteroid := &Asteroid{
		Sprite: NewSprite(),
	}

	if size == 0 {
		asteroid.radius = ASTEROID_LARGE_RADIUS
	} else {
		asteroid.radius = ASTEROID_LARGE_RADIUS / (size * 2)
	}
	asteroid.team = common.FOE
	asteroid.color = color.White
	asteroid.spin = asteroid.somePosNegValue(10)
	asteroid.deltaX = asteroid.somePosNegValue(10)
	asteroid.deltaY = asteroid.somePosNegValue(10)
	asteroid.cartesians = asteroid.generateVertices()

	return asteroid
}

func NewAsteroidFromExisting(astExploded *Asteroid) *Asteroid {
	newSmallerSize := astExploded.GetSize() + 1

	newAsteroid := NewAsteroid(newSmallerSize)
	newAsteroid.center = astExploded.center
	newAsteroid.deltaX = astExploded.deltaX/1.5 + newAsteroid.somePosNegValue(float64(5+newSmallerSize*2))
	newAsteroid.deltaY = astExploded.deltaY/1.5 + newAsteroid.somePosNegValue(float64(5+newSmallerSize*2))

	return newAsteroid
}

func (a *Asteroid) Move() {
	a.move(a)
}

func (a *Asteroid) Draw(screen *ebiten.Image) {
	a.renderVector(screen)
}

func (a *Asteroid) GetCenter() prime.Point {
	return a.center
}

func (a *Asteroid) GetRadius() int {
	return a.radius
}

func (a *Asteroid) GetTeam() common.Team {
	return a.team
}

func (a *Asteroid) AddToGame(list *list.List) {
	a.addToGame(list, a)
}

func (a *Asteroid) RemoveFromGame(list *list.List) {
	removed := a.removeFromGame(list, a)
	if !removed {
		return
	}

	a.spawnSmallerAsteroidsOrDebris()
	common.GetCommandCenterInstance().SetScore(common.GetCommandCenterInstance().GetScore() + 10*(int64(a.GetSize())+1))

	if a.GetSize() > 1 {
		common.PlaySound("pillow.wav")
	} else {
		common.PlaySound("kapow.wav")
	}
}

func (a *Asteroid) generateVertices() []prime.Point {
	vertices := common.GenerateRandomInt(7) + 25
	polars := make([]prime.PolarPoint, vertices)

	for i := 0; i < vertices; i++ {
		r := (800 + float64(common.GenerateRandomInt(200))) / ASTEROID_PRECISION
		theta := float64(common.GenerateRandomInt(ASTEROID_MAX_RADIANS_X1000)) / ASTEROID_PRECISION
		polars[i] = prime.PolarPoint{R: r, Theta: theta}
	}

	sort.Slice(polars, func(i, j int) bool {
		return polars[i].Theta < polars[j].Theta
	})

	points := make([]prime.Point, vertices)
	for i, p := range polars {
		points[i] = common.PolarToCartesian(p, ASTEROID_PRECISION)
	}

	return points
}

func (a *Asteroid) GetSize() int {
	switch a.radius {
	case ASTEROID_LARGE_RADIUS:
		return 0
	case ASTEROID_LARGE_RADIUS / 2:
		return 1
	case ASTEROID_LARGE_RADIUS / 4:
		return 2
	default:
		return 0
	}
}

func (a *Asteroid) spawnSmallerAsteroidsOrDebris() {
	size := a.GetSize()

	if size > 1 {
		common.GetCommandCenterInstance().GetGameOpsQueue().Enqueue(NewWhiteCloudDebris(a.Sprite), common.ADD)
	} else {
		// For large (0) and medium (1) sized Asteroids only, spawn 2 or 3 smaller asteroids respectively
		size += 2
		for size > 0 {
			// Add new Asteroid to the game
			common.GetCommandCenterInstance().GetGameOpsQueue().Enqueue(NewAsteroidFromExisting(a), common.ADD)
			size--
		}
	}
}
