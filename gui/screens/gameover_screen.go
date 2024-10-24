package screens

import (
	"github.com/Nel-sokchhunly/tank-game/config"
	"github.com/Nel-sokchhunly/tank-game/gui/button"
	"github.com/Nel-sokchhunly/tank-game/gui/text"
	"github.com/Nel-sokchhunly/tank-game/state"
)

func GameOverScreen() {
	text.Heading("Game Over", 20, 20)
	button.CenterLabelButton("Restart", func() {
		// Restart the game
		state.ScreenStateInstance.SetCurrentScreen(config.ScreenState.HOME)
	})
}
