package common

import (
	"container/list"

	"github.com/agerber/asteroids_go/config"
)

type ICommandCenter interface {
	InitGame()
	IncrementFrame()
	IsGameOver() bool
	GetUniDim() config.Dimension
	IsFalconPositionFixed() bool
	GetMovDebris() *list.List
	GetMovFriends() *list.List
	GetMovFoes() *list.List
	GetMovFloaters() *list.List
	GetGameOpsQueue() *GameOpsQueue
}
