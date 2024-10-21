package common

import (
	"container/list"
	"sync"
)

type ICommandCenter interface {
	InitGame()
	GetFrame() int64
	IncrementFrame()
	IsGameOver() bool
	GetUniDim() Dimension
	IsFalconPositionFixed() bool
	GetMovDebris() *list.List
	GetMovFriends() *list.List
	GetMovFoes() *list.List
	GetMovFloaters() *list.List
	GetGameOpsQueue() *GameOpsQueue
	GetScore() int64
	SetScore(score int64)
	GetLevel() int
	SetLevel(level int)
	GetFalcon() IFalcon
	GetNumFalcons() int
	SetNumFalcons(numFalcons int)
}

var singletonCommandCenter ICommandCenter

func SetCommandCenterInstance(commandCenter ICommandCenter) {
	sync.OnceFunc(func() {
		singletonCommandCenter = commandCenter
	})()
}

func GetCommandCenterInstance() ICommandCenter {
	return singletonCommandCenter
}
