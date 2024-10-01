package model

// GameState enum to handle different views
type GameState int

const (
	MainMenu GameState = iota
	Gameplay
	Settings
)

type MainModel struct {
	State GameState // To track which state/view we are in
}

func NewMainModel() *MainModel {
	return &MainModel{
		State: MainMenu, // Start with main menu
	}
}
