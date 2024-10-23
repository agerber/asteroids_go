package controller

import (
	"container/list"
	"os"

	"github.com/agerber/asteroids_go/common"
	"github.com/agerber/asteroids_go/model"
	"github.com/agerber/asteroids_go/view"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	gamePanel *view.GamePanel
}

func NewGame() *Game {
	gamePanel := view.NewGamePanel(common.DIM)

	// TODO: remove it
	common.GetCommandCenterInstance().SetLevel(1)

	return &Game{
		gamePanel: gamePanel,
	}
}

func (g *Game) Update() error {
	g.checkPressedKey()
	g.checkReleasedKey()
	g.checkNewLevel()
	g.checkFloaters()
	g.processGameOpsQueue()
	// keep track of the frame for development purposes
	common.GetCommandCenterInstance().IncrementFrame()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.gamePanel.Draw(screen)
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return common.DIM.Width, common.DIM.Height
}

func (g *Game) processGameOpsQueue() {
	for {
		select {
		case gameOp := <-common.GetCommandCenterInstance().GetGameOpsQueue().Dequeue():
			var list *list.List
			switch gameOp.Movable.GetTeam() {
			case common.FOE:
				list = common.GetCommandCenterInstance().GetMovFoes()
			case common.FRIEND:
				list = common.GetCommandCenterInstance().GetMovFriends()
			case common.FLOATER:
				list = common.GetCommandCenterInstance().GetMovFloaters()
			case common.DEBRIS:
				list = common.GetCommandCenterInstance().GetMovDebris()
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
	if common.GetCommandCenterInstance().GetFrame()%common.SPAWN_SHIELD_FLOATER == 0 {
		common.GetCommandCenterInstance().GetGameOpsQueue().Enqueue(model.NewShieldFloater(), common.ADD)
	}
}

func (g *Game) spawnNukeFloater() {
	if common.GetCommandCenterInstance().GetFrame()%common.SPAWN_NUKE_FLOATER == 0 {
		common.GetCommandCenterInstance().GetGameOpsQueue().Enqueue(model.NewNukeFloater(), common.ADD)
	}
}

func (g *Game) checkFloaters() {
	g.spawnShieldFloater()
	g.spawnNukeFloater()
}

func (g *Game) spawnBigAsteroids(num int) {
	for i := 0; i < num; i++ {
		common.GetCommandCenterInstance().GetGameOpsQueue().Enqueue(model.NewAsteroid(0), common.ADD)
	}
}

func (g *Game) checkPressedKey() {
	falcon := common.GetCommandCenterInstance().GetFalcon()

	switch {
	case ebiten.IsKeyPressed(ebiten.KeySpace):
		//CommandCenter.getInstance().getOpsQueue().enqueue(new Bullet(falcon), GameOp.Action.ADD);
	case ebiten.IsKeyPressed(ebiten.KeyF):
		//CommandCenter.getInstance().getOpsQueue().enqueue(new Nuke(falcon), GameOp.Action.ADD);
	case ebiten.IsKeyPressed(ebiten.KeyUp):
		falcon.SetThrusting(true)
		//SoundLoader.playSound("whitenoise_loop.wav");
	case ebiten.IsKeyPressed(ebiten.KeyLeft):
		falcon.SetTurnState(common.LEFT)
	case ebiten.IsKeyPressed(ebiten.KeyRight):
		falcon.SetTurnState(common.RIGHT)
	}
}

func (g *Game) checkReleasedKey() {
	falcon := common.GetCommandCenterInstance().GetFalcon()
	commandCenter := common.GetCommandCenterInstance()

	switch {
	case inpututil.IsKeyJustReleased(ebiten.KeyS) && commandCenter.IsGameOver():
		commandCenter.InitGame()
	case inpututil.IsKeyJustReleased(ebiten.KeyLeft):
		falcon.SetTurnState(common.IDLE)
	case inpututil.IsKeyJustReleased(ebiten.KeyRight):
		falcon.SetTurnState(common.IDLE)
	case inpututil.IsKeyJustReleased(ebiten.KeyUp):
		falcon.SetThrusting(false)
		//SoundLoader.stopSound("whitenoise_loop.wav")
	case inpututil.IsKeyJustReleased(ebiten.KeyP):
		commandCenter.SetPaused(!commandCenter.IsPaused())
	case inpututil.IsKeyJustReleased(ebiten.KeyQ):
		os.Exit(0)
	case inpututil.IsKeyJustReleased(ebiten.KeyA):
		commandCenter.SetRadar(!commandCenter.IsRadar())
	case inpututil.IsKeyJustReleased(ebiten.KeyM):
		if commandCenter.IsThemeMusic() {
			//SoundLoader.stopSound("dr_loop.wav")
		} else {
			//SoundLoader.playSound("dr_loop.wav");
		}
		commandCenter.SetThemeMusic(!commandCenter.IsThemeMusic())
	}
}

func (g *Game) isLevelClear() bool {
	// If there are no more Asteroids on the screen
	for e := common.GetCommandCenterInstance().GetMovFoes().Front(); e != nil; e = e.Next() {
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

	level := common.GetCommandCenterInstance().GetLevel()

	level++
	g.spawnBigAsteroids(level)
}
