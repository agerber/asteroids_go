package commandcenter

import (
	"container/list"
	"math"
	"sync"

	"github.com/agerber/asteroids_go/config"
	"github.com/agerber/asteroids_go/model"
)

var (
	singletonLock          = &sync.Mutex{}
	singletonCommandCenter *CommandCenter
)

func GetCommandCenterInstance() *CommandCenter {
	if singletonCommandCenter == nil {
		singletonLock.Lock()
		defer singletonLock.Unlock()
		if singletonCommandCenter == nil {
			singletonCommandCenter = NewCommandCenter()
		}
	}

	return singletonCommandCenter
}

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
	universe   Universe
	numFalcons int
	level      int
	score      int64
	paused     bool
	themeMusic bool
	radar      bool
	frame      int64
	//private final Falcon falcon  = new Falcon();
	miniDimHash map[Universe]config.Dimension
	//private final MiniMap miniMap = new MiniMap();
	MovDebris    *list.List
	MovFriends   *list.List
	MovFoes      *list.List
	MovFloaters  *list.List
	GameOpsQueue *GameOpsQueue
}

func NewCommandCenter() *CommandCenter {
	return &CommandCenter{
		//new Falcon();
		miniDimHash: make(map[Universe]config.Dimension),
		//new MiniMap();
		MovDebris:    list.New(),
		MovFriends:   list.New(),
		MovFoes:      list.New(),
		MovFloaters:  list.New(),
		GameOpsQueue: NewGameOpsQueue(),
	}
}

func (c *CommandCenter) InitGame() {
	c.clearAll()
	c.generateStarField()
	c.setDimHash()
	c.level = 0
	c.score = 0
	c.paused = false
	c.numFalcons = 4
	//falcon.decrementFalconNumAndSpawn()
	//opsQueue.enqueue(falcon, GameOp.Action.ADD)
	//opsQueue.enqueue(miniMap, GameOp.Action.ADD)
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

func (c *CommandCenter) GetUniDim() config.Dimension {
	return c.miniDimHash[c.universe]
}

func (c *CommandCenter) IsFalconPositionFixed() bool {
	return c.universe != FREE_FLY
}

func (c *CommandCenter) setDimHash() {
	c.miniDimHash[FREE_FLY] = config.Dimension{Width: 1, Height: 1}
	c.miniDimHash[CENTER] = config.Dimension{Width: 1, Height: 1}
	c.miniDimHash[BIG] = config.Dimension{Width: 2, Height: 2}
	c.miniDimHash[HORIZONTAL] = config.Dimension{Width: 3, Height: 1}
	c.miniDimHash[VERTICAL] = config.Dimension{Width: 1, Height: 3}
	c.miniDimHash[DARK] = config.Dimension{Width: 4, Height: 4}
}

func (c *CommandCenter) generateStarField() {
	for i := 0; i < 100; i++ {
		c.GameOpsQueue.Enqueue(model.NewStar(), ADD)
	}
}

func (c *CommandCenter) clearAll() {
	c.MovDebris.Init()
	c.MovFriends.Init()
	c.MovFoes.Init()
	c.MovFloaters.Init()
}
