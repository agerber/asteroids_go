package common

import (
	"container/list"
	"sync"
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

var universeStrings = map[Universe]string{
	FREE_FLY:   "FREE_FLY",
	CENTER:     "CENTER",
	BIG:        "BIG",
	HORIZONTAL: "HORIZONTAL",
	VERTICAL:   "VERTICAL",
	DARK:       "DARK",
}

func (u Universe) String() string {
	if str, ok := universeStrings[u]; ok {
		return str
	}
	return "UNKNOWN"
}

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
	IsPaused() bool
	SetPaused(paused bool)
	IsRadar() bool
	SetRadar(radar bool)
	IsThemeMusic() bool
	SetThemeMusic(themeMusic bool)
	SetUniverse(universe Universe)
	GetUniverse() Universe
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
