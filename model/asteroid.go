package model

import (
	"container/list"
	"image/color"
	"sort"

	"github.com/agerber/asteroids_go/common"
	"github.com/agerber/asteroids_go/config"
	"github.com/agerber/asteroids_go/model/prime"
	"github.com/agerber/asteroids_go/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Asteroid struct {
	*Sprite

	commandCenter common.ICommandCenter
}

func NewAsteroid(size int, commandCenter common.ICommandCenter) *Asteroid {
	asteroid := &Asteroid{
		Sprite:        NewSprite(commandCenter),
		commandCenter: commandCenter,
	}

	if size == 0 {
		asteroid.radius = config.LARGE_RADIUS
	} else {
		asteroid.radius = config.LARGE_RADIUS / (size * 2)
	}
	asteroid.team = common.FOE
	asteroid.color = color.White
	asteroid.spin = asteroid.somePosNegValue(0.05)
	asteroid.deltaX = asteroid.somePosNegValue(5)
	asteroid.deltaY = asteroid.somePosNegValue(5)
	asteroid.cartesians = asteroid.generateVertices()

	return asteroid
}

func NewAsteroidFromExisting(astExploded *Asteroid) *Asteroid {
	newSmallerSize := astExploded.getSize() + 1

	newAsteroid := NewAsteroid(newSmallerSize, astExploded.commandCenter)
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
	a.removeFromGame(list, a)

	a.spawnSmallerAsteroidsOrDebris()
	a.commandCenter.SetScore(a.commandCenter.GetScore() + 10*(int64(a.getSize())+1))

	if a.getSize() > 1 {
		//SoundLoader.playSound("pillow.wav")
	} else {
		//SoundLoader.playSound("kapow.wav");
	}
}

func (a *Asteroid) generateVertices() []prime.Point {
	vertices := utils.GenerateRandomInt(7) + 25
	polars := make([]prime.PolarPoint, vertices)

	for i := 0; i < vertices; i++ {
		r := (800 + float64(utils.GenerateRandomInt(200))) / config.PRECISION
		theta := float64(utils.GenerateRandomInt(config.MAX_RADIANS_X1000)) / config.PRECISION
		polars[i] = prime.PolarPoint{R: r, Theta: theta}
	}

	sort.Slice(polars, func(i, j int) bool {
		return polars[i].Theta < polars[j].Theta
	})

	points := make([]prime.Point, vertices)
	for i, p := range polars {
		points[i] = utils.PolarToCartesian(p, config.PRECISION)
	}

	return points
}

func (a *Asteroid) getSize() int {
	switch a.radius {
	case config.LARGE_RADIUS:
		return 0
	case config.LARGE_RADIUS / 2:
		return 1
	case config.LARGE_RADIUS / 4:
		return 2
	default:
		return 0
	}
}

func (a *Asteroid) spawnSmallerAsteroidsOrDebris() {
	size := a.getSize()

	if size > 1 {
		// Add WhiteCloudDebris to the game
		//a.commandCenter.GetGameOpsQueue().Enqueue(NewWhiteCloudDebris(a), common.ADD)
	} else {
		// For large (0) and medium (1) sized Asteroids only, spawn 2 or 3 smaller asteroids respectively
		size += 2
		for size > 0 {
			// Add new Asteroid to the game
			a.commandCenter.GetGameOpsQueue().Enqueue(NewAsteroidFromExisting(a), common.ADD)
			size--
		}
	}
}
