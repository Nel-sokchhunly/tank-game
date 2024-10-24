package main

import (
	"github.com/Nel-sokchhunly/tank-game/config"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func GameInit() {
	rl.InitWindow(int32(config.Screen.Width), int32(config.Screen.Height), "Tank Game")
	rl.SetTargetFPS(60)
}
