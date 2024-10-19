package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	screenWidth  = 1280
	screenheight = 720
)

var ()

func drawScene() {}

func input() {}

func render() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)

	drawScene()

	rl.EndDrawing()
}

func update() {}

func main() {
	rl.InitWindow(screenWidth, screenheight, "raylib [core] example - basic window")

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		input()
		update()
		render()
	}
	rl.CloseWindow()
}
