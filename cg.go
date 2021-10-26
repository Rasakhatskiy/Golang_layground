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
	Y := math.Sqrt(16-X*X) * SCALE
	return int(Y)
}

func (g Game) Draw(screen *ebiten.Image) {

	var prevr, prevg, prevb uint32 = 0, 0, 0

	if max > 0 {
		for i := 0; i < W; i++ {
			for j := 0; j < H; j++ {
				red, green, blue, alpha := img.At(i, j).RGBA()


				img.Set(
					i,
					j,
					color.RGBA{
						R: uint8(red + prevr),
						G: uint8(green + prevg),
						B: uint8(blue + prevb),
						A: uint8(alpha),
					},
				)
				prevr, prevg, prevb = red, green, blue
			}
		}
		max--
	}



	screen.DrawImage(img, nil)

	//img := ebiten.NewImage(W, H)
	//for i := 0; i < W; i++ {
	//	img.Set(i, H/2, color.Gray{
	//		Y: 255,
	//	})
	//
	//	img.Set(H/2, i, color.Gray{
	//		Y: 255,
	//	})
	//
	//	img.Set(i, f(i-H/2)+H/2, color.RGBA{
	//		B: 255,
	//		A: 255,
	//	})
	//
	//	img.Set(i, f2(i-H/2)+H/2, color.RGBA{
	//		R: 255,
	//		A: 255,
	//	})
	//
	//	img.Set(i, f3(i-H/2)+H/2, color.RGBA{
	//		G: 255,
	//		A: 255,
	//	})
	//
	//	img.Set(i, -f3(i-H/2)+H/2, color.RGBA{
	//		G: 255,
	//		A: 255,
	//	})
	//}

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
}

func (g Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return W, H
}
