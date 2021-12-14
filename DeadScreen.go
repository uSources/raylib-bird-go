package main

import (
	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type DeadScreen struct {
}

func (m *DeadScreen) Draw() {
	rl.DrawText(gameOverText, (screenWidth/2)-(rl.MeasureText(gameOverText, defaultFontSize)/2), screenHeight/2-50, defaultFontSize, rl.RayWhite)
	if raygui.Button(rl.NewRectangle(screenWidth/2-40, screenHeight/2, 80, 40), playButtonText) {
		game.ChangeScreen(gameScreen)
	}
	if raygui.Button(rl.NewRectangle(screenWidth/2-40, screenHeight/2+50, 80, 40), menuButtonText) {
		game.ChangeScreen(menuScreen)
	}
}

func (m *DeadScreen) Init()   {}
func (m *DeadScreen) Update() {}
func (m *DeadScreen) Load()   {}
func (m *DeadScreen) Flush()  {}
