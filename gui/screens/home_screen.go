package screens

import (
	"github.com/Nel-sokchhunly/tank-game/config"
	"github.com/Nel-sokchhunly/tank-game/gui/text"
	"github.com/Nel-sokchhunly/tank-game/state"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func startGame() {
	state.ScreenStateInstance.SetCurrentScreen(config.ScreenState.GAME)
}

func HomeScreen() {
	// inputs
	HomeScreenInput()

	// UI
	text.Heading("Ping Pong", 20, 20)
	text.Body("Press spacebar to start the game", 20, int32(config.Screen.Height-40))
}

func HomeScreenInput() {
	// spacebar pressed

	if rl.IsKeyPressed(rl.KeySpace) {
		startGame()
	}
}
