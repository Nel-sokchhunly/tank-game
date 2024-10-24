package state

type Game struct {
	LeftScore  int
	RightScore int
}

func (game *Game) UpdateLeftScore() {
	game.LeftScore++
}

func (game *Game) UpdateRightScore() {
	game.RightScore++
}

func (game *Game) IsGameOver() bool {
	return game.LeftScore >= 5 || game.RightScore >= 5
}

var GameState = Game{
	LeftScore:  0,
	RightScore: 0,
}
