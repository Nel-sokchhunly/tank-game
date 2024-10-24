package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	GameInit()
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		Render()
		rl.EndDrawing()
	}
}
