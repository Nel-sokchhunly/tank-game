package state

import (
	"github.com/Nel-sokchhunly/tank-game/config"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type BouncingBall struct {
	Position rl.Vector2
	Radius   float32
	Speed    rl.Vector2
}

func (ball *BouncingBall) Start() {
	ball.Position.X += ball.Speed.X
	ball.Position.Y += ball.Speed.Y

	if ball.Position.X >= float32(config.Screen.Width-int16(ball.Radius)) || ball.Position.X <= ball.Radius {
		ball.Speed.X *= -1
	}

	if ball.Position.Y > float32(config.Screen.Height-int16(ball.Radius)) || ball.Position.Y <= ball.Radius {
		ball.Speed.Y *= -1
	}

	// detect collision on both left and right side
	if ball.Position.X <= ball.Radius {
		ball.Position.X = ball.Radius
	}

	// right side collision
	if ball.Position.X >= float32(config.Screen.Width-int16(ball.Radius)) {
		GameState.UpdateRightScore()
		ball.ResetPosition("right")
	}

	// left side collision
	if ball.Position.X <= ball.Radius {
		GameState.UpdateLeftScore()
		ball.ResetPosition("left")
	}
}

func (ball *BouncingBall) ResetPosition(collisionSide string) {
	ball.Position = rl.NewVector2(float32(config.Screen.Width/2), float32(config.Screen.Height/2))
	if collisionSide == "left" {
		ball.Speed = rl.NewVector2(3, 2)
	} else if collisionSide == "right" {
		ball.Speed = rl.NewVector2(-3, 2)
	} else {
		ball.Speed = rl.NewVector2(3, 2)
	}
}

var BallState = BouncingBall{
	Position: rl.NewVector2(float32(config.Screen.Width/2), float32(config.Screen.Height/2)),
	Radius:   float32(10),
	Speed:    rl.NewVector2(3, 2),
}
