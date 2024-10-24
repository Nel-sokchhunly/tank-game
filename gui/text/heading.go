package text

import rl "github.com/gen2brain/raylib-go/raylib"

func Heading(text string, posX int32, posY int32) {
	rl.DrawText(text, posX, posY, 32, rl.White)
}
