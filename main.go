package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1280
	screenheight = 720
)

var (
	// uncategorised
	running     = true
	bkgcolour   = rl.NewColor(147, 211, 196, 255)
	grassSprite rl.Texture2D
	frameCount  int

	// player vars
	playerSprite                                  rl.Texture2D
	playerSrc                                     rl.Rectangle
	playerDest                                    rl.Rectangle
	playerSpeed                                   float32 = 3
	playerMoving                                  bool
	playerDir, playerFrame                        int
	playerUp, playerDown, playerRight, playerLeft bool

	// music vars
	musicPaused bool
	music       rl.Music

	// camera
	cam rl.Camera2D
)

func drawScene() {
	rl.DrawTexture(grassSprite, 100, 50, rl.White)
	rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(playerDest.Width, playerDest.Height), 0, rl.White)
}

func input() {
	// movement input
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		playerMoving = true
		playerUp = true
		playerDir = 1 // for animation framesy value
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		playerMoving = true
		playerDown = true
		playerDir = 0 // for animation frames y value
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		playerMoving = true
		playerLeft = true
		playerDir = 2 // for animation frames y value
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		playerMoving = true
		playerRight = true
		playerDir = 3 // for animation frames y value
	}
	// debugging
	if rl.IsKeyPressed(rl.KeyQ) {
		musicPaused = !musicPaused
	}
}

func update() {
	// should close window
	running = !rl.WindowShouldClose()

	// player movement & animation
	playerSrc.X = 0 // for animation, if the player isn't moving stay at default frame in animation
	if playerMoving {
		if playerUp {
			playerDest.Y -= playerSpeed
		}
		if playerDown {
			playerDest.Y += playerSpeed
		}
		if playerLeft {
			playerDest.X -= playerSpeed
		}
		if playerRight {
			playerDest.X += playerSpeed
		}
		if frameCount%8 == 1 {
			playerFrame++
		}
		playerSrc.X = playerSrc.Width * float32(playerFrame) // animation stuff teehee
	}

	frameCount++

	if playerFrame > 3 {
		playerFrame = 0
	}

	playerSrc.Y = playerSrc.Height * float32(playerDir)

	// bkg music
	rl.UpdateMusicStream(music) // update the music sound
	if musicPaused {
		rl.PauseMusicStream(music)
	} else {
		rl.ResumeMusicStream(music)
	} // pauses music based of off input

	fmt.Println(frameCount)

	// camera
	cam.Target = rl.NewVector2(float32(playerDest.X-(playerDest.Width/2)), float32(playerDest.Y-(playerDest.Height/2)))
	playerMoving = false
	playerUp, playerDown, playerLeft, playerRight = false, false, false, false
}

func render() {
	rl.BeginDrawing()

	rl.ClearBackground(bkgcolour)
	rl.BeginMode2D(cam)
	drawScene()

	rl.EndMode2D()
	rl.EndDrawing()
}

func init() {
	rl.InitWindow(screenWidth, screenheight, "first game using raylib & go")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)

	// texture/sprite loading
	grassSprite = rl.LoadTexture("sproutLandsPack/Tilesets/Grass.png")
	playerSprite = rl.LoadTexture("sproutLandsPack/Characters/BasicCharakterSpritesheet.png")

	// player Src and Dest
	playerSrc = rl.NewRectangle(0, 0, 48, 48)
	playerDest = rl.NewRectangle(200, 200, 100, 100)

	// music
	rl.InitAudioDevice()
	music = rl.LoadMusicStream("sproutLandsPack/Music/AnimalCrossingCampMusicMorning.mp3") // sets musics file path
	musicPaused = false                                                                    // makes sure musicPaused if false so music starts playing
	rl.PlayMusicStream(music)
	rl.SetMusicVolume(music, 0.1) // sets volume for music

	// camera init
	cam = rl.NewCamera2D(rl.NewVector2(float32(screenWidth/2), float32(screenheight/2)), rl.NewVector2(float32(playerDest.X-(playerDest.Width/2)), float32(playerDest.Y-(playerDest.Height/2))), 0, 1.5)
}

func quit() {
	// unload textures
	rl.UnloadTexture(grassSprite)
	rl.UnloadTexture(playerSprite)
	rl.UnloadMusicStream(music)
	rl.CloseAudioDevice()

	// close window
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
