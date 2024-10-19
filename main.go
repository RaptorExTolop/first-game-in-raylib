package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	screenWidth  = 1280
	screenheight = 720
)

var (
	// uncategorised
	running     = true
	bkgcolour   = rl.NewColor(147, 211, 196, 255)
	grassSprite rl.Texture2D

	// player vars
	playerSprite rl.Texture2D
	playerSrc    rl.Rectangle
	playerDest   rl.Rectangle
	playerSpeed  float32 = 3

	// music vars
	musicPaused bool
	music       rl.Music
)

func drawScene() {
	rl.DrawTexture(grassSprite, 100, 50, rl.White)
	rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(playerDest.Width, playerDest.Height), 0, rl.White)
}

func input() {
	// movement
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		playerDest.Y -= playerSpeed
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		playerDest.Y += playerSpeed
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		playerDest.X -= playerSpeed
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		playerDest.X += playerSpeed
	}
	// debugging
	if rl.IsKeyPressed(rl.KeyQ) {
		musicPaused = !musicPaused
	}
}

func render() {
	rl.BeginDrawing()

	rl.ClearBackground(bkgcolour)
	drawScene()

	rl.EndDrawing()
}

func update() {
	// should close window
	running = !rl.WindowShouldClose()

	// bkg music
	rl.UpdateMusicStream(music) // update the music sound
	if musicPaused {
		rl.PauseMusicStream(music)
	} else {
		rl.ResumeMusicStream(music)
	} // pauses music based of off input
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
