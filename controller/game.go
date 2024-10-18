package controller

import (
	"container/list"

	"github.com/agerber/asteroids_go/common"
	"github.com/agerber/asteroids_go/config"
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

	return &Game{
		commandCenter: commandCenter,
		gamePanel:     gamePanel,
	}
}

func (g *Game) Update() error {
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
