package main

import (
	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type MenuScreen struct {
}

func (m *MenuScreen) Draw() {
	rl.DrawText(gameTitle, (screenWidth/2)-(rl.MeasureText(gameTitle, defaultFontSize)/2), screenHeight/2-50, defaultFontSize, rl.Black)
	if raygui.Button(rl.NewRectangle(screenWidth/2-40, screenHeight/2, 80, 40), playButtonText) {
		game.ChangeScreen(gameScreen)
	}
	if raygui.Button(rl.NewRectangle(screenWidth/2-40, screenHeight/2+50, 80, 40), exitButtonText) {
		rl.CloseWindow()
	}
}

func (m *MenuScreen) Init()   {}
func (m *MenuScreen) Update() {}
func (m *MenuScreen) Load()   {}
func (m *MenuScreen) Flush()  {}
