package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
)

var W, H = 980, 1200
var SCALE = 10.0
var img *ebiten.Image
var max = 1

func main() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("F:\\Desktop\\photo_2021-10-27_11-44-39.jpg")
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
