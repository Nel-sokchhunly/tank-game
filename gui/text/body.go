package text

import (
	"github.com/Nel-sokchhunly/tank-game/config"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Body(text string, posX int32, posY int32) {
	rl.DrawText(text, posX, posY, 20, config.GameColor.Text)
}
