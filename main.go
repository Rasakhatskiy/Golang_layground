package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
)

var W, H = 720, 694
var SCALE = 10.0
var img *ebiten.Image
var max = 3

func main() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("F:\\Desktop\\1.png")
	if err != nil {
		log.Fatal(err)
	}



	ebiten.SetWindowSize(W, H)
	ebiten.SetWindowTitle("bruh momento numero cuarenta y siete")
	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
