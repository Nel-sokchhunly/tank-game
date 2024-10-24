package state

import (
	"github.com/Nel-sokchhunly/tank-game/config"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Paddle struct {
	Position rl.Vector2
	Size     rl.Vector2
	AutoPlay bool
}

func (paddle *Paddle) Draw(movement string) {
	rl.DrawRectangleV(paddle.Position, paddle.Size, config.GameColor.Secondary)
	paddle.checkCollisionWithBall(&BallState)

	if !paddle.AutoPlay {
		if movement == "ws" {
			paddle.initMovementWS()
		} else if movement == "arrows" {
			paddle.initMovementArrows()
		} else {
			paddle.initMovementWS()
		}
	} else {
		// move the paddle with the ball y position
		if paddle.Position.Y >= 10 || paddle.Position.Y <= float32(rl.GetScreenHeight())-paddle.Size.Y-10 {
			paddle.Position.Y = (BallState.Position.Y - paddle.Size.Y/2)
		}
	}

}

func (paddle *Paddle) checkCollisionWithBall(ball *BouncingBall) {
	if rl.CheckCollisionCircleRec(ball.Position, ball.Radius, rl.NewRectangle(paddle.Position.X, paddle.Position.Y, paddle.Size.X, paddle.Size.Y)) {
		ball.Speed.X = -ball.Speed.X * 1.1
	}

}

func (paddle *Paddle) initMovementWS() {
	if rl.IsKeyDown(rl.KeyW) && paddle.Position.Y > 10 {
		paddle.Position.Y -= 5
	}

	if rl.IsKeyDown(rl.KeyS) && paddle.Position.Y < float32(rl.GetScreenHeight())-paddle.Size.Y-10 {
		paddle.Position.Y += 5
	}

}

func (paddle *Paddle) initMovementArrows() {
	if rl.IsKeyDown(rl.KeyUp) && paddle.Position.Y > 10 {
		paddle.Position.Y -= 5
	}

	if rl.IsKeyDown(rl.KeyDown) && paddle.Position.Y < float32(rl.GetScreenHeight())-paddle.Size.Y-10 {
		paddle.Position.Y += 5
	}
}

var LeftPaddleInstance = Paddle{
	Position: rl.NewVector2(10, float32(config.Screen.Height/2-50)),
	Size:     rl.NewVector2(10, 100),
	AutoPlay: false,
}

var RightPaddleInstance = Paddle{
	Position: rl.NewVector2(float32(config.Screen.Width-10-10), float32(config.Screen.Height/2-50)),
	Size:     rl.NewVector2(10, 100),
	AutoPlay: false,
}
