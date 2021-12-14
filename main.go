package main

var (
	cryptokey  = []byte("example key 1234")
	game       = Game{}
	menuScreen = &MenuScreen{}
	gameScreen = &PlayScreen{}
	deadScreen = &DeadScreen{}
)

func main() {
	game.Init()

}
