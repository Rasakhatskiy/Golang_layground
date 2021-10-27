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

func findEdge(image [][]uint8, w, h int)  *ebiten.Image{
	var maskV = [3][3]int{
		{-1, 0, 1},
		{-2, 0, 2},
		{-1, 0, 1},
	}

	var maskH = [3][3]int{
		{-1, -2, -1},
		{0, 0, 0},
		{1, 2, 1},
	}

	var newImage = ebiten.NewImage(w, h)

	for i := 0; i < w-3; i++ {
		for j := 0; j < h-3; j++ {
			var t1 float64 = 0
			var t2 float64 = 0
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					t1 += float64(int(image[i+k][j+l]) * maskV[k][l])
					t2 += float64(int(image[i+k][j+l]) * maskH[k][l])
				}
			}
			var middle = math.Sqrt(float64(t1*t1 + t2*t2))
			newImage.Set(i+1, j+1, color.Gray{Y: uint8(middle)})
		}
	}
	return newImage
}

func grayscale(image *ebiten.Image, w, h int) [][]uint8 {
	var newImage = make([][]uint8, w)

	for i := 0; i < W; i++ {
		newImage[i] = make([]uint8, h)
		for j := 0; j < H; j++ {
			ured, ugreen, ublue, _ := image.At(i, j).RGBA()
			red, green, blue := uint8(ured), uint8(ugreen), uint8(ublue)
			y := uint8(float64(red)*0.299 + float64(green)*0.587 + float64(blue)*0.114)
			newImage[i][j] = y
		}
	}
	return newImage
}

func (g Game) Draw(screen *ebiten.Image) {

	//var prevr, prevg, prevb uint32 = 0, 0, 0

	if max > 0 {
		var gray = grayscale(img, W, H)
		img = findEdge(gray, W, H)
		max--
	}
	screen.DrawImage(img, nil)
}

func (g Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return W, H
}
