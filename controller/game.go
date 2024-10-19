package controller

import (
	"container/list"

	"github.com/agerber/asteroids_go/common"
	"github.com/agerber/asteroids_go/config"
	"github.com/agerber/asteroids_go/model"
	"github.com/agerber/asteroids_go/view"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	commandCenter common.ICommandCenter
	gamePanel     *view.GamePanel
}

func NewGame() *Game {
	commandCenter := NewCommandCenter()
	gamePanel := view.NewGamePanel(config.DIM, commandCenter)

	// Move to correct location
	commandCenter.InitGame()
	// TODO: remove it
	commandCenter.SetLevel(5)

	return &Game{
		commandCenter: commandCenter,
		gamePanel:     gamePanel,
	}
}

func (g *Game) Update() error {
	g.checkNewLevel()
	g.checkFloaters()
	g.processGameOpsQueue()
	// keep track of the frame for development purposes
	g.commandCenter.IncrementFrame()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.gamePanel.Draw(screen)
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return config.DIM.Width, config.DIM.Height
}

func (g *Game) processGameOpsQueue() {
	for {
		select {
		case gameOp := <-g.commandCenter.GetGameOpsQueue().Dequeue():
			var list *list.List
			switch gameOp.Movable.GetTeam() {
			case common.FOE:
				list = g.commandCenter.GetMovFoes()
			case common.FRIEND:
				list = g.commandCenter.GetMovFriends()
			case common.FLOATER:
				list = g.commandCenter.GetMovFloaters()
			case common.DEBRIS:
				list = g.commandCenter.GetMovDebris()
			default:
				return
			}

			switch gameOp.Action {
			case common.ADD:
				gameOp.Movable.AddToGame(list)
			case common.REMOVE:
				gameOp.Movable.RemoveFromGame(list)
			}
		default:
			return
		}
	}
}

func (g *Game) spawnShieldFloater() {
	if g.commandCenter.GetFrame()%config.SPAWN_SHIELD_FLOATER == 0 {
		g.commandCenter.GetGameOpsQueue().Enqueue(model.NewShieldFloater(g.commandCenter), common.ADD)
	}
}

func (g *Game) spawnNukeFloater() {
	if g.commandCenter.GetFrame()%config.SPAWN_NUKE_FLOATER == 0 {
		g.commandCenter.GetGameOpsQueue().Enqueue(model.NewNukeFloater(g.commandCenter), common.ADD)
	}
}

func (g *Game) checkFloaters() {
	g.spawnShieldFloater()
	g.spawnNukeFloater()
}

func (g *Game) spawnBigAsteroids(num int) {
	for i := 0; i < num; i++ {
		g.commandCenter.GetGameOpsQueue().Enqueue(model.NewAsteroid(0, g.commandCenter), common.ADD)
	}
}

func (g *Game) isLevelClear() bool {
	// If there are no more Asteroids on the screen
	for e := g.commandCenter.GetMovFoes().Front(); e != nil; e = e.Next() {
		if _, ok := e.Value.(*model.Asteroid); ok {
			return false
		}
	}
	return true
}

func (g *Game) checkNewLevel() {
	if !g.isLevelClear() {
		return
	}

	level := g.commandCenter.GetLevel()

	level++
	g.spawnBigAsteroids(level)
}
