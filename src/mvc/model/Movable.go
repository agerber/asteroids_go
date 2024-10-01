package model

import "github.com/agerber/asteroids_go/src/mvc/model/prime"

// Movable interface defines methods for objects in the game
type Movable interface {
	// Move()
	// Draw(imgOff interface{}) // Interface allows for flexibility in drawing surface
	GetRadius() int
	GetTeam() Team
	GetCenter() prime.Point // Assuming Point is defined elsewhere
	// AddToGame(list interface{}) // Interface allows for flexibility in list type
	// RemoveFromGame(list interface{})
}
