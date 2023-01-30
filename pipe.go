package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
)

type Pipe struct {
	top       int
	bottom    int
	pos       rl.Vector2
	width     int32
	highlight bool
}

func newPipe() Pipe {
	spacing := 80
	centery := rand.Intn((screenHeight-spacing)-spacing) + spacing
	return Pipe{centery - spacing/2, screenHeight - (centery + spacing/2), rl.Vector2{screenWidth, 0}, 80, false}
}

func (p Pipe) show() {
	var color rl.Color
	if p.highlight {
		color = rl.Red
	} else {
		color = rl.Black
	}

	rl.DrawRectangle(int32(p.pos.X), int32(p.pos.Y), p.width, int32(p.top), color)
	rl.DrawRectangle(int32(p.pos.X), int32(screenHeight-p.bottom), p.width, int32(p.bottom), color)
}

func (p *Pipe) update(b Bird) {
	p.pos.X -= pipesSpeed
	if p.hit(b) {
		p.highlight = true
	} else {
		p.highlight = false
	}
}

func (p Pipe) offScreen() bool {
	return p.pos.X < -80
}

func (p Pipe) hit(b Bird) bool {
	if rl.CheckCollisionCircleRec(b.pos, b.radius, rl.Rectangle{
		p.pos.X,
		p.pos.Y,
		float32(p.width),
		float32(p.top),
	}) ||
		rl.CheckCollisionCircleRec(b.pos, b.radius, rl.Rectangle{
			p.pos.X,
			float32(screenHeight - p.bottom),
			float32(p.width),
			float32(p.bottom),
		}) {
		return true
	} else {
		return false
	}
}
