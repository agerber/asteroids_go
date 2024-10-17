package commandcenter

import "github.com/agerber/asteroids_go/model"

const MAX_GAME_OPS_QUEUE_SIZE = 100000

type GameOpsQueue struct {
	queue chan *GameOp
}

func NewGameOpsQueue() *GameOpsQueue {
	return &GameOpsQueue{
		queue: make(chan *GameOp, MAX_GAME_OPS_QUEUE_SIZE),
	}
}

func (q *GameOpsQueue) Enqueue(movable model.Movable, action Action) {
	q.queue <- &GameOp{
		Movable: movable,
		Action:  action,
	}
}

func (q *GameOpsQueue) Dequeue() <-chan *GameOp {
	return q.queue
}
