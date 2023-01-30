package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Bird struct {
	pos    rl.Vector2
	radius float32
	vel    float32
}

func newBird() Bird {
	return Bird{rl.Vector2{80, screenHeight / 2}, 0, 0}
}

func (b Bird) show() {
	rl.DrawCircle(int32(b.pos.X), int32(b.pos.Y), 17, rl.Black)
}

func (b *Bird) update() {
	b.vel += gravity
	b.vel *= velocityRate
	b.pos.Y += b.vel

	if b.pos.Y > screenHeight {
		b.pos.Y = screenHeight
		b.vel = 0
	}
	if b.pos.Y < 0 {
		b.pos.Y = 0
		b.vel = 0
	}
	if rl.IsKeyPressed(32) {
		b.goUp()
	}
}

func (b *Bird) goUp() {
	b.vel -= lift
}
