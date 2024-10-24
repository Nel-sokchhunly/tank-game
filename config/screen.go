package config

type TScreenState struct {
	HOME     string
	GAME     string
	GAMEOVER string
}

var ScreenState = TScreenState{
	HOME:     "HOME",
	GAME:     "GAME",
	GAMEOVER: "GAMEOVER",
}
