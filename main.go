package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1280
	screenheight = 720
)

var (
	// uncategorised
	running    = true
	frameCount int

	tileDest   rl.Rectangle
	tileSrc    rl.Rectangle
	tilemap    []int
	srcMap     []string
	mapW, mapH int

	bkgcolour   = rl.NewColor(147, 211, 196, 255)
	grassSprite rl.Texture2D

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
	///rl.DrawTexture(grassSprite, 100, 50, rl.White)

	for i := 0; i < len(tilemap); i++ {
		if tilemap[i] != 0 {
			tileDest.X = tileDest.Width * float32(i%mapW)
			tileDest.Y = tileDest.Height * float32(i/mapW)
			tileSrc.X = tileSrc.Width * float32((tilemap[i]-1)%int(grassSprite.Width/int32(tileSrc.Width)))
		}
		tileSrc.Y = tileSrc.Height * float32((tilemap[i]-1)/int(grassSprite.Width))
		rl.DrawTexturePro(grassSprite, tileSrc, tileDest, rl.NewVector2(tileDest.Width, tileDest.Height), 0, rl.White)
	}
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
	playerSrc.X = playerSrc.Width * float32(playerFrame) // animation stuff teehee
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
	} else if frameCount%45 == 1 {
		playerFrame++
	}

	frameCount++

	if playerFrame > 3 {
		playerFrame = 0
	}

	if !playerMoving && playerFrame > 1 {
		playerFrame = 0
	}

	playerSrc.X = playerSrc.Width * float32(playerFrame)

	playerSrc.Y = playerSrc.Height * float32(playerDir)

	// bkg music
	rl.UpdateMusicStream(music) // update the music sound
	if musicPaused {
		rl.PauseMusicStream(music)
	} else {
		rl.ResumeMusicStream(music)
	} // pauses music based of off input

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

func loadmap(mapfile string) {
	file, err := os.ReadFile(mapfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	removeNewLines := strings.Replace(string(file), "\n", " ", -1)
	sliced := strings.Split(removeNewLines, " ")
	mapW = -1
	mapH = -1
	for i := 0; i < len(sliced); i++ {
		s, _ := strconv.ParseInt(sliced[i], 10, 64)
		m := int(s)
		if mapW == -1 {
			mapW = m
		} else if mapH == -1 {
			mapH = m
		} else {
			tilemap = append(tilemap, m)
		}
	}
	if len(tilemap) > mapW*mapH {
		tilemap = tilemap[:len(tilemap)-1]
	}

	/*mapW, mapH = 5, 5
	for i := 0; i < (mapW * mapH); i++ {
		tilemap = append(tilemap, 1)
	}*/
}

func init() {
	rl.InitWindow(screenWidth, screenheight, "first game using raylib & go")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)

	tileDest = rl.NewRectangle(0, 0, 16, 16)
	tileSrc = rl.NewRectangle(0, 0, 16, 16)

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
	loadmap("one.map")
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
