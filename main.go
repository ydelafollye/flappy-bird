package main

import (
	"fmt"
	"github.com/gen2brain/raylib-go/raylib"
	"math/rand"
	"strconv"
	"time"
)

const (
	screenWidth  = 800
	screenHeight = 600
	gravity      = 0.6
	lift         = 15
	velocityRate = 0.9
	pipesSpeed   = 3
)

var framesCounter int

type Game struct {
	WindowsShouldClose bool
	pipe               []Pipe
	bird               Bird
}

func NewGame() (g Game) {
	g.Init()
	return
}

func (g *Game) Init() {
	g.WindowsShouldClose = false

	g.pipe = append(g.pipe, newPipe())
	g.bird = newBird()

}

func (g *Game) Update() {
	if rl.WindowShouldClose() {
		g.WindowsShouldClose = true
	}

	g.bird.update()

	if framesCounter%100 == 0 {
		g.pipe = append(g.pipe, newPipe())
	}
	for i := 0; i < len(g.pipe); i++ {
		g.pipe[i].update(g.bird)
		if g.pipe[i].offScreen() {
			g.pipe = g.pipe[1:]
			i--
		}
	}

}

func (g Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	fps := strconv.Itoa(int(rl.GetFPS()))
	rl.DrawText("FPS:"+fps, 10, 10, 20, rl.Black)

	g.bird.show()
	for _, p := range g.pipe {
		p.show()
	}
	rl.EndDrawing()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	framesCounter = 0
	game := NewGame()
	rl.InitWindow(screenWidth, screenHeight, "Smart Dot")
	rl.SetTargetFPS(60)

	for !game.WindowsShouldClose {
		if framesCounter == 600 {
			fmt.Println(framesCounter)
			framesCounter = 0
		}
		framesCounter++
		game.Draw()
		game.Update()
	}

	rl.CloseWindow()
}
