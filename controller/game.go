package controller

import (
	"container/list"
	"os"

	"github.com/agerber/asteroids_go/common"
	"github.com/agerber/asteroids_go/model"
	"github.com/agerber/asteroids_go/model/prime"
	"github.com/agerber/asteroids_go/view"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	gamePanel *view.GamePanel
}

func NewGame() *Game {
	gamePanel := view.NewGamePanel(common.DIM)
	return &Game{
		gamePanel: gamePanel,
	}
}

func (g *Game) Update() error {
	g.checkPressedKey()
	g.checkReleasedKey()
	g.checkCollisions()
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

func (g *Game) checkCollisions() {
	var pntFriendCenter, pntFoeCenter prime.Point
	var radFriend, radFoe int
	for e := common.GetCommandCenterInstance().GetMovFriends().Front(); e != nil; e = e.Next() {
		movFriend := e.Value.(common.Movable)
		for f := common.GetCommandCenterInstance().GetMovFoes().Front(); f != nil; f = f.Next() {
			movFoe := f.Value.(common.Movable)

			pntFriendCenter = movFriend.GetCenter()
			pntFoeCenter = movFoe.GetCenter()
			radFriend = movFriend.GetRadius()
			radFoe = movFoe.GetRadius()

			if common.DistanceBetween2Points(pntFriendCenter, pntFoeCenter) < float64(radFriend+radFoe) {
				common.GetCommandCenterInstance().GetGameOpsQueue().Enqueue(movFriend, common.REMOVE)
				common.GetCommandCenterInstance().GetGameOpsQueue().Enqueue(movFoe, common.REMOVE)
			}
		}
	}

	pntFalCenter := common.GetCommandCenterInstance().GetFalcon().GetCenter()
	radFalcon := common.GetCommandCenterInstance().GetFalcon().GetRadius()
	var pntFloaterCenter prime.Point
	var radFloater int
	for e := common.GetCommandCenterInstance().GetMovFloaters().Front(); e != nil; e = e.Next() {
		movFloater := e.Value.(common.Movable)
		pntFloaterCenter = movFloater.GetCenter()
		radFloater = movFloater.GetRadius()

		if common.DistanceBetween2Points(pntFalCenter, pntFloaterCenter) < float64(radFalcon+radFloater) {
			common.GetCommandCenterInstance().GetGameOpsQueue().Enqueue(movFloater, common.REMOVE)
		}
	}
}

func (g *Game) spawnBigAsteroids(num int) {
	for i := 0; i < num; i++ {
		common.GetCommandCenterInstance().GetGameOpsQueue().Enqueue(model.NewAsteroid(0), common.ADD)
	}
}

func (g *Game) checkPressedKey() {
	falcon := common.GetCommandCenterInstance().GetFalcon()

	switch {
	case inpututil.IsKeyJustReleased(ebiten.KeySpace):
		common.GetCommandCenterInstance().GetGameOpsQueue().Enqueue(model.NewBullet(falcon), common.ADD)
	case inpututil.IsKeyJustReleased(ebiten.KeyF):
		common.GetCommandCenterInstance().GetGameOpsQueue().Enqueue(model.NewNuke(falcon), common.ADD)
	case ebiten.IsKeyPressed(ebiten.KeyUp):
		falcon.SetThrusting(true)
		common.PlaySound("whitenoise_loop.wav")
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
		common.StopSound("whitenoise_loop.wav")
	case inpututil.IsKeyJustReleased(ebiten.KeyP):
		commandCenter.SetPaused(!commandCenter.IsPaused())
	case inpututil.IsKeyJustReleased(ebiten.KeyQ):
		common.CloseSound()
		os.Exit(0)
	case inpututil.IsKeyJustReleased(ebiten.KeyA):
		commandCenter.SetRadar(!commandCenter.IsRadar())
	case inpututil.IsKeyJustReleased(ebiten.KeyM):
		if commandCenter.IsThemeMusic() {
			common.StopSound("dr_loop.wav")
		} else {
			common.PlaySound("dr_loop.wav")
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

	commandCenter := common.GetCommandCenterInstance()
	falcon := commandCenter.GetFalcon()

	level := common.GetCommandCenterInstance().GetLevel()
	commandCenter.SetScore(commandCenter.GetScore() + int64(level)*10000)
	falcon.SetCenter(prime.Point{
		X: float64(common.DIM.Width / 2),
		Y: float64(common.DIM.Height / 2),
	})

	ordinal := level % 6
	commandCenter.SetUniverse(common.Universe(ordinal))
	commandCenter.SetRadar(ordinal > 1)

	level++
	commandCenter.SetLevel(level)
	g.spawnBigAsteroids(level)
	falcon.SetShield(model.FALCON_INITIAL_SPAWN_TIME)
	falcon.SetShowLevel(model.FALCON_INITIAL_SPAWN_TIME)
}
