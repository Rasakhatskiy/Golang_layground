package main

import "github.com/hajimehoshi/ebiten/v2"

var W, H = 512, 512
var SCALE = 10.0

func main() {
	ebiten.SetWindowSize(W, H)
	ebiten.SetWindowTitle("bruh momento numero cuarenta y siete")
	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}