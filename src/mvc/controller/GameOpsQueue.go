package controller

import (
	"sync"

	"github.com/agerber/asteroids_go/src/mvc/model"
	"github.com/agerber/asteroids_go/src/mvc/model/prime"
)

// GameOpsQueue represents a queue of game operations.
type GameOpsQueue struct {
	mu sync.Mutex
	ll *prime.LinkedList // Pointer to a linked list object
}

// NewGameOpsQueue creates a new GameOpsQueue object.
func NewGameOpsQueue() *GameOpsQueue {
	return &GameOpsQueue{
		mu: sync.Mutex{},
		ll: prime.NewLinkedList(), // Assuming you have a NewLinkedList function
	}
}

// Enqueue adds a game operation to the queue.
func (q *GameOpsQueue) Enqueue(mov model.Movable, action GameOpAction) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.ll.Enqueue(&GameOp{Movable: mov, Action: action})
}

// Dequeue removes and returns the first game operation from the queue.
func (q *GameOpsQueue) Dequeue() *GameOp {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.ll.Dequeue().(*GameOp) // Type cast to GameOp
}

// Len returns the number of elements in the queue.
func (q *GameOpsQueue) Len() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.ll.Len()
}
