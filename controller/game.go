package controller

import (
	"container/list"

	"github.com/agerber/asteroids_go/commandcenter"
	"github.com/agerber/asteroids_go/config"
	"github.com/agerber/asteroids_go/model"
	"github.com/agerber/asteroids_go/view"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	gamePanel *view.GamePanel
}

func NewGame() *Game {
	// Move to correct location
	commandcenter.GetCommandCenterInstance().InitGame()
	gamePanel := view.NewGamePanel(config.DIM)
	return &Game{gamePanel: gamePanel}
}

func (g *Game) Update() error {
	processGameOpsQueue()
	// keep track of the frame for development purposes
	commandcenter.GetCommandCenterInstance().IncrementFrame()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.gamePanel.Draw(screen)
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return config.DIM.Width, config.DIM.Height
}

func processGameOpsQueue() {
	for {
		select {
		case gameOp := <-commandcenter.GetCommandCenterInstance().GameOpsQueue.Dequeue():
			var list *list.List
			switch gameOp.Movable.GetTeam() {
			case model.FOE:
				list = commandcenter.GetCommandCenterInstance().MovFoes
			case model.FRIEND:
				list = commandcenter.GetCommandCenterInstance().MovFriends
			case model.FLOATER:
				list = commandcenter.GetCommandCenterInstance().MovFloaters
			case model.DEBRIS:
				list = commandcenter.GetCommandCenterInstance().MovDebris
			default:
				return
			}

			switch gameOp.Action {
			case commandcenter.ADD:
				gameOp.Movable.AddToGame(list)
			case commandcenter.REMOVE:
				gameOp.Movable.RemoveFromGame(list)
			}
		default:
			return
		}
	}
}
