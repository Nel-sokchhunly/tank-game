package button

import (
	"github.com/Nel-sokchhunly/tank-game/config"
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func CenterLabelButton(text string, onClick func()) {

	x, y := config.Screen.CenterScreen()
	textWidth := rl.MeasureText(text, 20)

	positionX := x - int32(textWidth/2)

	button := gui.LabelButton(rl.NewRectangle(float32(positionX), float32(y), float32(textWidth), 20), text)
	if button {
		onClick()
	}
}
