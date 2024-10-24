package state

import (
	"github.com/Nel-sokchhunly/tank-game/config"
)

type ScreenState struct {
	currentScreen string
}

func (s *ScreenState) GetScreen() string {
	return s.currentScreen
}

func (s *ScreenState) SetCurrentScreen(screen string) {
	s.currentScreen = screen
}

var ScreenStateInstance = ScreenState{currentScreen: config.ScreenState.HOME}
