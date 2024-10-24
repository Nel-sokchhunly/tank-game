package button

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Text(text string, posX float32, posY float32, onClick func()) {
	textWidth := rl.MeasureText(text, 20)

	button := gui.LabelButton(rl.NewRectangle(posX, posY, float32(textWidth), 20), text)
	if button {
		onClick()
	}

}
