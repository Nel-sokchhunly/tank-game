package config

import "image/color"

type ColorPallet struct {
	Background          color.RGBA
	SecondaryBackground color.RGBA
	Primary             color.RGBA
	Secondary           color.RGBA
	Text                color.RGBA
}

var GameColor = ColorPallet{
	Background:          color.RGBA{16, 11, 42, 255},
	SecondaryBackground: color.RGBA{88, 45, 39, 255},
	Primary:             color.RGBA{151, 65, 51, 255},
	Secondary:           color.RGBA{49, 71, 67, 255},
	Text:                color.RGBA{236, 254, 202, 255},
}
