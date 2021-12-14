package main

import rl "github.com/gen2brain/raylib-go/raylib"

var floor = Floor{}

type Floor struct {
	rect rl.Rectangle
}

func (f *Floor) Init() {
	f.rect = rl.Rectangle{
		X:      0,
		Y:      screenHeight - 40,
		Width:  screenWidth,
		Height: 40,
	}
}

func (f *Floor) Draw() {
	rl.DrawRectangleRec(f.rect, rl.Gray)
}
