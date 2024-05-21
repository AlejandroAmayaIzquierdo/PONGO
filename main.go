package main

import (
	game "WMG/src"

	sdl "github.com/veandco/go-sdl2/sdl"
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic("Failed to initialize a window")
	}
	win, err := sdl.CreateWindow("GO SDL Window",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		1280, 720, sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE)
	if err != nil {
		panic("Failed to create the SDL window")
	}
	game := game.NewGame(win)
	game.Handle()
	win.Destroy()
}
