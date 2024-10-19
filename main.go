package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	screenWidth  = 1280
	screenheight = 720
)

var (
	running   = true
	bkgcolour = rl.NewColor(147, 211, 196, 255)
)

func drawScene() {}

func input() {}

func render() {
	rl.BeginDrawing()

	rl.ClearBackground(bkgcolour)
	drawScene()

	rl.EndDrawing()
}

func update() {
	running = !rl.WindowShouldClose()
}

func init() {
	rl.InitWindow(screenWidth, screenheight, "first game using raylib & go")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)
}

func quit() {
	rl.CloseWindow()
}

func main() {

	for running {
		input()
		update()
		render()
	}
	quit()
}
