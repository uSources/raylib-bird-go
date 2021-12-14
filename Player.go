package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var player = Player{}

type Player struct {
	rect         rl.Rectangle
	state        int
	acceleration rl.Vector2
	rotation     float32
}

func (p *Player) Init() {
	p.Reset()
}

func (p *Player) Reset() {
	p.rect = rl.Rectangle{
		X:      0 + 40,
		Y:      screenHeight/2 - 20,
		Width:  40,
		Height: 40,
	}
	p.acceleration = rl.Vector2Zero()
	p.rotation = 0
}

func (p *Player) Update() {
	p.Physics()
	p.Input()
	p.Draw()
}

func (p *Player) Draw() {
	rl.DrawRectangleRec(p.rect, rl.Beige)
}

func (p *Player) Physics() {

	p.acceleration.Y += gravity * rl.GetFrameTime()
	destination := rl.Vector2Add(rl.Vector2{
		X: p.rect.X,
		Y: p.rect.Y,
	}, p.acceleration)
	p.rect.X = destination.X
	p.rect.Y = destination.Y

	if p.rect.Y+p.rect.Width < 0 {
		p.Die()
	}

	if rl.CheckCollisionRecs(player.rect, floor.rect) {
		p.Die()
	}

	for _, p := range pipes {
		if rl.CheckCollisionRecs(player.rect, p.top.rect) || rl.CheckCollisionRecs(player.rect, p.bottom.rect) {
			player.Die()
		}
	}

}

func (p *Player) Die() {
	rl.PlaySound(dieSound)
	SaveScore()
	game.ChangeScreen(deadScreen)
}

func (p *Player) Input() {
	if rl.IsKeyPressed(rl.KeySpace) || rl.GetTouchPointCount() > 0 || rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		p.acceleration.Y = -300 * rl.GetFrameTime()
		rl.PlaySound(wingSound)
	}
}
