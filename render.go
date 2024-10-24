package main

import (
	cfg "github.com/Nel-sokchhunly/tank-game/config"
	"github.com/Nel-sokchhunly/tank-game/gui/screens"
	"github.com/Nel-sokchhunly/tank-game/state"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Render() {
	rl.ClearBackground(cfg.GameColor.Background)

	switch state.ScreenStateInstance.GetScreen() {
	case cfg.ScreenState.HOME:
		screens.HomeScreen()
	case cfg.ScreenState.GAME:
		screens.GameScreen()
	case cfg.ScreenState.GAMEOVER:
		screens.GameOverScreen()
	default:
		screens.HomeScreen()
	}

}
