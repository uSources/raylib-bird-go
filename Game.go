package main

import (
	b64 "encoding/base64"
	"os"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	scoreSound = rl.Sound{}
	wingSound  = rl.Sound{}
	dieSound   = rl.Sound{}
)

type Game struct {
	screen Screen
}

func (g *Game) Init() {
	rl.InitWindow(screenWidth, screenHeight, screenTitle)

	rl.InitAudioDevice()

	rl.SetTargetFPS(60)

	g.Load()

	g.ChangeScreen(menuScreen)

	for !rl.WindowShouldClose() {
		g.Update()
		g.Draw()
	}
	g.Flush()

}

func (g *Game) Update() {
	game.screen.Update()
}

func (g *Game) Draw() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.SkyBlue)

	game.screen.Draw()

	rl.EndDrawing()
}

func (g *Game) Load() {
	dieSound = rl.LoadSound("resources/die.mp3")
	scoreSound = rl.LoadSound("resources/point.mp3")
	wingSound = rl.LoadSound("resources/wing.mp3")
}

func (g *Game) ChangeScreen(s Screen) {
	if game.screen != nil {
		game.screen.Flush()
	}
	game.screen = s
	game.screen.Load()
	game.screen.Init()
}

func (g *Game) Flush() {
	game.screen.Flush()
	rl.UnloadSound(dieSound)
	rl.UnloadSound(scoreSound)
	rl.UnloadSound(wingSound)

	rl.CloseWindow()
}

func SaveScore() {
	if score > topScore {
		f, err := os.Create(scoreFileText)
		Error(err)
		encodeText := b64.StdEncoding.EncodeToString([]byte(strconv.Itoa(score)))
		f.Write([]byte(encodeText))
	}
}

func LoadScore() {
	if _, err := os.Stat(scoreFileText); err == nil {
		dat, err := os.ReadFile(scoreFileText)
		Error(err)
		decodeText, err := b64.StdEncoding.DecodeString(string(dat))
		Error(err)

		i, err := strconv.Atoi(string(decodeText))
		topScore = i
	}

}

func Error(e error) {
	if e != nil {
		panic(e)
	}
}
