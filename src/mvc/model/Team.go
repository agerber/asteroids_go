package model

import "fmt"

// Team represents the available teams for movables
type Team int

const (
	Friend Team = iota
	Foe
	Floater
	Debris
)

func (t Team) String() string {
	switch t {
	case Friend:
		return "Friend"
	case Foe:
		return "Foe"
	case Floater:
		return "Floater"
	case Debris:
		return "Debris"
	default:
		return fmt.Sprintf("Unknown Team (%d)", t)
	}
}
