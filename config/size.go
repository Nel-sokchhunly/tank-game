package config

type GameScreen struct {
	Width  int16
	Height int16
}

func (s *GameScreen) CenterScreen() (int32, int32) {
	return int32(s.Width / 2), int32(s.Height / 2)
}

var Screen = GameScreen{
	Width:  800,
	Height: 600,
}
