package main

import (
	"strconv"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameState int16

const (
	MENU GameState = iota
	RUN
	DEAD
)

const (
	screenWidth       = 640
	screenHeight      = 480
	screenTitle       = "Flappy GO!"
	gravity           = 9.8 * 2
	gameOverText      = "Game over!"
	gameTitle         = "Flappy Raylib"
	defaultFontSize   = 32
	playButtonText    = "Play"
	optionsButtonText = "Options"
	menuButtonText    = "Menu"
)

var (
	startPosition = rl.Vector2{}
	gameState     = MENU
	floor         = Floor{}
	pipes         = make([]*PipeGroup, 0)
	player        = Player{}
	pointSound    = rl.Sound{}
	wingSound     = rl.Sound{}
	dieSound      = rl.Sound{}
	score         = 0
)

type Player struct {
	rect         rl.Rectangle
	state        int
	acceleration rl.Vector2
	rotation     float32
}

type Floor struct {
	rect rl.Rectangle
}

type Pipe struct {
	rect rl.Rectangle
}

type PipeGroup struct {
	top     Pipe
	bottom  Pipe
	crossed bool
}

func main() {

	rl.InitWindow(screenWidth, screenHeight, screenTitle)

	rl.InitAudioDevice()

	rl.SetTargetFPS(60)

	//Sounds

	dieSound = rl.LoadSound("resources/die.mp3")
	wingSound = rl.LoadSound("resources/wing.mp3")
	pointSound = rl.LoadSound("resources/point.mp3")

	resetPlayer()
	makeFloor()
	makePipes()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.SkyBlue)

		if gameState == MENU {
			drawMenu()
		}

		if gameState == RUN {
			updatePipes()
			updateInput()
			updateRotation()
			updatePhysics()
			checkCol()
			drawFloor()
			drawPipes()
			drawPlayer(player)
			rl.DrawText("Score: "+strconv.Itoa(score), screenWidth-150, 10, 24, rl.DarkGreen)

		}

		if gameState == DEAD {
			drawDead()
		}

		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func resetPlayer() {
	player.rect = rl.Rectangle{
		X:      startPosition.X + 40,
		Y:      startPosition.Y,
		Width:  40,
		Height: 40,
	}
	player.acceleration = rl.Vector2Zero()
	player.rotation = 0
}

func resetPipe(pipe *PipeGroup, i int) {

	startPosition := float32(screenHeight)

	if i == 0 {
		startPosition = float32(pipes[len(pipes)-1].top.rect.X)
	} else {
		startPosition = float32(pipes[i-1].top.rect.X)
	}
	makePipe(pipe, startPosition)
}

func updatePhysics() {
	player.acceleration.Y += gravity * rl.GetFrameTime()
	destination := rl.Vector2Add(rl.Vector2{
		X: player.rect.X,
		Y: player.rect.Y,
	}, player.acceleration)
	player.rect.X = destination.X
	player.rect.Y = destination.Y
}

func updateInput() {
	if rl.IsKeyPressed(rl.KeySpace) || rl.GetTouchPointCount() > 0 || rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		player.acceleration.Y = -300 * rl.GetFrameTime()
		rl.PlaySound(wingSound)
	}
}

func updateRotation() {
	player.rotation = rl.Clamp(player.acceleration.Y*5, -15, 15)
}

func checkCol() {

	if player.rect.Y+player.rect.Width < 0 {
		die()
	}

	if rl.CheckCollisionRecs(player.rect, floor.rect) {
		die()
	}

	for _, p := range pipes {
		if rl.CheckCollisionRecs(player.rect, p.top.rect) || rl.CheckCollisionRecs(player.rect, p.bottom.rect) {
			die()
		}
	}

}

func die() {
	gameState = DEAD
	rl.PlaySound(dieSound)
}

func drawPlayer(player Player) {
	rl.DrawRectangleRec(player.rect, rl.Beige)
}

func makeFloor() {
	floor = Floor{
		rect: rl.Rectangle{
			X:      0,
			Y:      screenHeight - 40,
			Width:  screenWidth,
			Height: 40,
		},
	}
}

func makePipes() {
	startPosition := float32(screenWidth)
	for i := 0; i < 5; i++ {
		p := PipeGroup{}
		pipes = append(pipes, &p)
		makePipe(&p, startPosition)
		startPosition = p.top.rect.X
	}
}

func drawFloor() {
	rl.DrawRectangleRec(floor.rect, rl.Gray)
}

func drawPipes() {
	for _, p := range pipes {
		rl.DrawRectangleRec(p.top.rect, rl.Green)
		rl.DrawRectangleRec(p.bottom.rect, rl.Green)
	}
}

func makePipe(group *PipeGroup, startPosition float32) {

	distance := rl.GetRandomValue(50, 100)
	YStartPosition := rl.GetRandomValue(-64, 64)

	group.top.rect.Width = 40
	group.bottom.rect.Width = 40
	group.top.rect.Height = 300
	group.bottom.rect.Height = 300
	group.crossed = false

	group.top.rect.X = float32(distance) + startPosition + float32(distance)
	group.bottom.rect.X = float32(distance) + startPosition + float32(distance)

	group.top.rect.Y = -(group.top.rect.Height / 2) + float32(YStartPosition)
	group.bottom.rect.Y = (-(group.top.rect.Height / 2) + screenHeight) + float32(YStartPosition)

}

func updatePipes() {
	for i, p := range pipes {
		p.top.rect.X -= 300 * rl.GetFrameTime()
		p.bottom.rect.X -= 300 * rl.GetFrameTime()

		distance := p.top.rect.X + float32(p.top.rect.Width) - player.rect.X + float32(p.top.rect.Width)/2
		if distance < 0.1 && p.crossed == false {
			score += 1
			p.crossed = true
			rl.PlaySound(pointSound)
		}

		if p.top.rect.X+float32(p.top.rect.Width) < 0 {
			resetPipe(p, i)
		}

	}
}

func debugPipes() {
	for _, p := range pipes {

		rl.DrawRectangleRec(p.top.rect, rl.White)
		rl.DrawRectangleRec(p.bottom.rect, rl.White)

	}
}

func drawMenu() {
	rl.DrawText(gameTitle, screenWidth-(rl.MeasureText(gameTitle, defaultFontSize)*2), screenHeight/2-50, defaultFontSize, rl.Black)
	if raygui.Button(rl.NewRectangle(screenWidth/2-40, screenHeight/2, 80, 40), playButtonText) {
		gameState = RUN
	}
	if raygui.Button(rl.NewRectangle(screenWidth/2-40, screenHeight/2+50, 80, 40), optionsButtonText) {
		gameState = RUN
	}
}

func resetGame() {
	resetPlayer()
	score = 0
	pipes = make([]*PipeGroup, 0)
	makePipes()
}

func drawDead() {
	resetGame()
	rl.DrawText(gameOverText, (screenWidth/2)-(rl.MeasureText(gameOverText, defaultFontSize)/2), screenHeight/2-50, defaultFontSize, rl.RayWhite)
	if raygui.Button(rl.NewRectangle(screenWidth/2-40, screenHeight/2, 80, 40), playButtonText) {
		gameState = RUN
	}
	if raygui.Button(rl.NewRectangle(screenWidth/2-40, screenHeight/2+50, 80, 40), menuButtonText) {
		gameState = MENU
	}
}
