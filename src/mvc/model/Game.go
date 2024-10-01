package model

/* type Falcon struct {
	Shield             int  // The Falcon's shield level (0-100)
	NukeMeter          int  // Nuke charge level (0-100)
	IsMaxSpeedAttained bool // Whether the Falcon is at max speed
	ShowLevel          int  // Shows the level progress
} */

type PolarPoint struct {
	R     float64 // The radial distance
	Theta float64 // The angle in radians
}
type CommandCenter struct {
	Level      int     // Current game level
	Score      int     // Player's score
	NumFalcons int     // Number of falcons (lives) remaining
	IsGameOver bool    // Indicates if the game is over
	IsPaused   bool    // Indicates if the game is paused
	Falcon     *Falcon // Player's ship
}

var commandCenter *CommandCenter

// GetCommandCenter returns the game state singleton
func GetCommandCenter() *CommandCenter {
	if commandCenter == nil {
		falcon := NewFalcon() // Make sure this returns a valid Falcon instance
		commandCenter = &CommandCenter{
			Level:      1,
			Score:      0,
			NumFalcons: 3,
			IsGameOver: true,
			IsPaused:   false,
			Falcon:     falcon,
		}
	}
	return commandCenter
}
