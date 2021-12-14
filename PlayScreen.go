package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	score    = 0
	topScore = 0
)

const (
	scoreText    = "Score:"
	topscoreText = "Top Score:"
)

type PlayScreen struct {
}

func (m *PlayScreen) Init() {
	score = 0
	player.Init()
	floor.Init()
	FillPipesSlide()
}
func (m *PlayScreen) Update() {
	player.Update()
	UpdatePipes()
}
func (m *PlayScreen) Load() {
	LoadScore()
}
func (m *PlayScreen) Draw() {
	player.Draw()
	floor.Draw()
	DrawPipes()
	rl.DrawText(scoreText+strconv.Itoa(score), 0, 0, defaultFontSize, rl.Black)
	rl.DrawText(topscoreText+strconv.Itoa(topScore), 0, defaultFontSize, defaultFontSize, rl.Black)
	rl.DrawFPS(0, defaultFontSize*2)
}
func (m *PlayScreen) Flush() {
	FlushPipes()
}
