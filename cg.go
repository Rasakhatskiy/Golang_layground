package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"math"
)

type Game struct {
}

func (g Game) Update() error {
	return nil
}

func f(x int) int {
	X := float64(x) / SCALE
	Y := math.Sin(X) * SCALE
	return int(Y)
}

func f2(x int) int {
	X := float64(x) / SCALE
	Y := math.Cos(X) * SCALE
	return int(Y)
}

func f3(x int) int {
	X := float64(x) / SCALE
	Y :=math.Sqrt(16-X*X) * SCALE
	return int(Y)
}

func (g Game) Draw(screen *ebiten.Image) {
	img := ebiten.NewImage(W, H)
	for i := 0; i < W; i++ {
		img.Set(i, H/2, color.Gray{
			Y: 255,
		})

		img.Set(H/2, i, color.Gray{
			Y: 255,
		})

		img.Set(i, f(i-H/2)+H/2, color.RGBA{
			B: 255,
			A: 255,
		})

		img.Set(i, f2(i-H/2)+H/2, color.RGBA{
			R: 255,
			A: 255,
		})

		img.Set(i, f3(i-H/2)+H/2, color.RGBA{
			G: 255,
			A: 255,
		})

		img.Set(i, -f3(i-H/2)+H/2, color.RGBA{
			G: 255,
			A: 255,
		})
	}

	//for i := 0; i < W; i++ {
	//
	//	for j := 0; j < H; j++ {
	//		//img.Set(
	//		//	i,
	//		//	j,
	//		//	color.Gray{
	//		//		Y: 255,
	//		//	},
	//		//)
	//	}
	//}
	screen.DrawImage(img, nil)
}

func (g Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return W, H
}
