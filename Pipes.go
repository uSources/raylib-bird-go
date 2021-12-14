package main

import rl "github.com/gen2brain/raylib-go/raylib"

var pipes = make([]*PipeGroup, 0)

type Pipe struct {
	rect rl.Rectangle
}

type PipeGroup struct {
	top     Pipe
	bottom  Pipe
	crossed bool
}

func (pg *PipeGroup) Make(startPosition float32) {

	distance := rl.GetRandomValue(50, 100)
	YStartPosition := rl.GetRandomValue(-64, 64)

	pg.top.rect.Width = 40
	pg.bottom.rect.Width = 40
	pg.top.rect.Height = 300
	pg.bottom.rect.Height = 300
	pg.crossed = false

	pg.top.rect.X = float32(distance) + startPosition + float32(distance)
	pg.bottom.rect.X = float32(distance) + startPosition + float32(distance)

	pg.top.rect.Y = -(pg.top.rect.Height / 2) + float32(YStartPosition)
	pg.bottom.rect.Y = (-(pg.top.rect.Height / 2) + screenHeight) + float32(YStartPosition)

}

func FillPipesSlide() {
	startPosition := float32(screenWidth)
	for i := 0; i < 5; i++ {
		g := &PipeGroup{}
		pipes = append(pipes, g)
		g.Make(startPosition)
		startPosition = g.top.rect.X
	}
}

func UpdatePipes() {
	for i, p := range pipes {
		p.top.rect.X -= 300 * rl.GetFrameTime()
		p.bottom.rect.X -= 300 * rl.GetFrameTime()

		distance := p.top.rect.X + float32(p.top.rect.Width) - player.rect.X + float32(p.top.rect.Width)/2
		if distance < 0.1 && p.crossed == false {
			score += 1
			p.crossed = true
			rl.PlaySound(scoreSound)
		}

		if p.top.rect.X+float32(p.top.rect.Width) < 0 {
			p.Reset(i)
		}

	}
}

func (pg *PipeGroup) Reset(i int) {

	startPosition := float32(screenHeight)

	if i == 0 {
		startPosition = float32(pipes[len(pipes)-1].top.rect.X)
	} else {
		startPosition = float32(pipes[i-1].top.rect.X)
	}
	pg.Make(startPosition)
}

func (p *Pipe) Draw() {
	rl.DrawRectangleRec(p.rect, rl.Green)
}

func (pg *PipeGroup) Draw() {
	pg.bottom.Draw()
	pg.top.Draw()
}

func DrawPipes() {
	for _, pg := range pipes {
		pg.Draw()
	}
}

func FlushPipes() {
	pipes = nil
}
