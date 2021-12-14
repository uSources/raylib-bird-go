package main

type Screen interface {
	Init()
	Draw()
	Update()
	Load()
	Flush()
}
