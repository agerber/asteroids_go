package controller

import (
	"container/list"
	"math"

	"github.com/agerber/asteroids_go/common"
	"github.com/agerber/asteroids_go/model"
)

type Universe int

const (
	FREE_FLY Universe = iota
	CENTER
	BIG
	HORIZONTAL
	VERTICAL
	DARK
)

type CommandCenter struct {
	universe    Universe
	numFalcons  int
	level       int
	score       int64
	paused      bool
	themeMusic  bool
	radar       bool
	frame       int64
	falcon      common.IFalcon
	miniDimHash map[Universe]common.Dimension
	//private final MiniMap miniMap = new MiniMap();
	movDebris    *list.List
	movFriends   *list.List
	movFoes      *list.List
	movFloaters  *list.List
	gameOpsQueue *common.GameOpsQueue
}

func NewCommandCenter() *CommandCenter {
	return &CommandCenter{
		falcon:      model.NewFalcon(),
		miniDimHash: make(map[Universe]common.Dimension),
		//new MiniMap();
		movDebris:    list.New(),
		movFriends:   list.New(),
		movFoes:      list.New(),
		movFloaters:  list.New(),
		gameOpsQueue: common.NewGameOpsQueue(),
	}
}

func (c *CommandCenter) InitGame() {
	c.clearAll()
	c.generateStarField()
	c.setDimHash()
	c.level = 0
	c.score = 0
	c.paused = false
	//set to one greater than number of falcons lives in your game as decrementFalconNumAndSpawn() also decrements
	c.numFalcons = 4
	c.falcon.DecrementFalconNumAndSpawn()
	c.gameOpsQueue.Enqueue(c.falcon, common.ADD)
	//opsQueue.enqueue(miniMap, GameOp.Action.ADD)
}

func (c *CommandCenter) GetFrame() int64 {
	return c.frame
}

func (c *CommandCenter) IncrementFrame() {
	if c.frame >= math.MaxInt64 {
		c.frame = 0
		return
	}
	c.frame++
}

func (c *CommandCenter) IsGameOver() bool {
	return c.numFalcons < 1
}

func (c *CommandCenter) GetUniDim() common.Dimension {
	return c.miniDimHash[c.universe]
}

func (c *CommandCenter) IsFalconPositionFixed() bool {
	return c.universe != FREE_FLY
}

func (c *CommandCenter) GetMovDebris() *list.List {
	return c.movDebris
}

func (c *CommandCenter) GetMovFriends() *list.List {
	return c.movFriends
}

func (c *CommandCenter) GetMovFoes() *list.List {
	return c.movFoes
}

func (c *CommandCenter) GetMovFloaters() *list.List {
	return c.movFloaters
}

func (c *CommandCenter) GetGameOpsQueue() *common.GameOpsQueue {
	return c.gameOpsQueue
}

func (c *CommandCenter) GetScore() int64 {
	return c.score
}

func (c *CommandCenter) SetScore(score int64) {
	c.score = score
}

func (c *CommandCenter) GetLevel() int {
	return c.level
}

func (c *CommandCenter) SetLevel(level int) {
	c.level = level
}

func (c *CommandCenter) GetFalcon() common.IFalcon {
	return c.falcon
}

func (c *CommandCenter) GetNumFalcons() int {
	return c.numFalcons
}

func (c *CommandCenter) SetNumFalcons(numFalcons int) {
	c.numFalcons = numFalcons
}

func (c *CommandCenter) setDimHash() {
	c.miniDimHash[FREE_FLY] = common.Dimension{Width: 1, Height: 1}
	c.miniDimHash[CENTER] = common.Dimension{Width: 1, Height: 1}
	c.miniDimHash[BIG] = common.Dimension{Width: 2, Height: 2}
	c.miniDimHash[HORIZONTAL] = common.Dimension{Width: 3, Height: 1}
	c.miniDimHash[VERTICAL] = common.Dimension{Width: 1, Height: 3}
	c.miniDimHash[DARK] = common.Dimension{Width: 4, Height: 4}
}

func (c *CommandCenter) generateStarField() {
	for i := 0; i < 100; i++ {
		c.gameOpsQueue.Enqueue(model.NewStar(), common.ADD)
	}
}

func (c *CommandCenter) clearAll() {
	c.movDebris.Init()
	c.movFriends.Init()
	c.movFoes.Init()
	c.movFloaters.Init()
}
