package model

import (
	"image"
)

// Team enumeration represents the different teams in the game
type Team int

const (
	FRIEND Team = iota
	FOE
	FLOATER
	DEBRIS
)

/*
Movable interface provides a simplified interface to a complex subsystem or set of classes.
It hides the complexity by offering a more straightforward and unified API.
The goal is to make subsystems easier to use by providing a higher-level interface that clients can interact with.
This is an example of the Facade design pattern.
*/
type Movable interface {
	// Move updates the position and state of the movable object.
	Move()

	// Draw renders the movable object using the provided graphics context.
	Draw(g Graphics)

	// GetCenter returns the center point of the movable object.
	GetCenter() image.Point

	// GetRadius returns the radius of the movable object, used for collision detection.
	GetRadius() int

	// GetTeam returns the team to which the movable object belongs.
	GetTeam() Team

	// AddToGame adds the movable object to the game space.
	// The 'list' parameter will be one of the following: movFriends, movFoes, movDebris, movFloaters.
	// This is your opportunity to add sounds or perform other side effects.
	// See processGameOpsQueue() of Game class for more details.
	AddToGame(list *[]Movable)

	// RemoveFromGame removes the movable object from the game space.
	// The 'list' parameter will be one of the following: movFriends, movFoes, movDebris, movFloaters.
	RemoveFromGame(list *[]Movable)
}

// Graphics is a placeholder interface for the graphics context used in drawing.
// You should replace this with the actual graphics context/interface provided by your graphics library.
type Graphics interface {
	// Define necessary methods for drawing, e.g., DrawImage, SetColor, etc.
}
