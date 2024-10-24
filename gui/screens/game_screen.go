package screens

import (
	"strconv"

	"github.com/Nel-sokchhunly/tank-game/config"
	"github.com/Nel-sokchhunly/tank-game/state"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func endGame() {
	// End the game
	state.ScreenStateInstance.SetCurrentScreen(config.ScreenState.GAMEOVER)
}

func GameScreen() {
	// state

	// Background
	// draw left side of the screen with primary color
	rl.DrawRectangle(0, 0, int32(config.Screen.Width/2), int32(config.Screen.Height), config.GameColor.Background)
	// draw right side of the screen with secondary color
	rl.DrawRectangle(int32(config.Screen.Width/2), 0, int32(config.Screen.Width/2), int32(config.Screen.Height), config.GameColor.Primary)

	// ENTITIES
	state.LeftPaddleInstance.Draw("ws")
	state.RightPaddleInstance.Draw("arrows")
	BouncingBall()

	if state.GameState.IsGameOver() {
		endGame()
	}

	// Score
	rl.DrawText(strconv.Itoa(state.GameState.LeftScore), int32(config.Screen.Width/2)-24, 10, 20, rl.White)
	rl.DrawText(strconv.Itoa(state.GameState.RightScore), int32(config.Screen.Width/2)+10, 10, 20, rl.White)

}

func BouncingBall() {
	// rl.DrawCircleV(state.BallState.Position, state.BallState.Radius, config.GameColor.Primary)
	// if at the left side of the screen use primary color
	if state.BallState.Position.X < float32(config.Screen.Width/2) {
		rl.DrawCircleV(state.BallState.Position, state.BallState.Radius, config.GameColor.Primary)
	} else {
		// if at the right side of the screen use background color
		rl.DrawCircleV(state.BallState.Position, state.BallState.Radius, config.GameColor.Background)
	}
	state.BallState.Start()
}
