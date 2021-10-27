package main

import (
	"flag"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"math/rand"
	"os"
	"time"
)

type MyColor struct {
	r, g, b, a uint8
}

func fromUint32(r, g, b, a uint32) MyColor {
	var c MyColor
	c.r = uint8(r)
	c.g = uint8(g)
	c.b = uint8(b)
	c.a = uint8(a)
	return c
}

type GrayScaleMatrix [][]uint8
type RGBAMatrix [][]MyColor

var grayScale *GrayScaleMatrix
var mainMatrix *RGBAMatrix

var whatToShow = 0
var img, resultImg, grayScaleImage, edgeImage, rgbPaletteImage *ebiten.Image
var trashold = 50.0
var isTreshold = true
var toRecalculate = true
var toMakeGray = true
var invert = true
var stopper = true

const (
	ShowResult    int = 0
	ShowEdges         = 1
	ShowGrayScale     = 2
	ShowRGB           = 3
)

type Game struct {
}

func (g Game) Update() error {
	//g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	return nil
}

func findEdge(image *GrayScaleMatrix, w, h int) *GrayScaleMatrix {
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

	result := make(GrayScaleMatrix, w)
	for i := 0; i < w; i++ {
		result[i] = make([]uint8, h)
	}

	for i := 0; i < w-3; i++ {
		for j := 0; j < h-3; j++ {
			var t1 float64 = 0
			var t2 float64 = 0
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					t1 += float64(int((*image)[i+k][j+l]) * maskV[k][l])
					t2 += float64(int((*image)[i+k][j+l]) * maskH[k][l])
				}
			}
			var middle = math.Sqrt(t1*t1 + t2*t2)

			if isTreshold {
				if middle < trashold {
					middle = 0
				}
				if middle >= trashold {
					middle = 255
				}
			}

			if invert {
				middle = 255 - middle
			}
			result[i+1][j+1] = uint8(middle)
		}
	}
	return &result
}

func grayscale(matrix *RGBAMatrix, w, h int) *GrayScaleMatrix {
	var newImage = make(GrayScaleMatrix, w)
	for i := 0; i < W; i++ {
		newImage[i] = make([]uint8, h)
		for j := 0; j < H; j++ {
			c := (*matrix)[i][j]
			newImage[i][j] = uint8(float64(c.r)*0.299 + float64(c.g)*0.587 + float64(c.b)*0.114)
		}
	}
	return &newImage
}

func makeMainMatrix(image *ebiten.Image) *RGBAMatrix {
	w, h := image.Size()
	result := make(RGBAMatrix, w)
	for i := 0; i < w; i++ {
		result[i] = make([]MyColor, h)
		for j := 0; j < h; j++ {
			r, g, b, a := image.At(i, j).RGBA()
			result[i][j] = fromUint32(r, g, b, a)
		}
	}
	return &result
}

func grayScaleToImage(matrix *GrayScaleMatrix, w, h int) *ebiten.Image {
	newImage := ebiten.NewImage(w, h)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			newImage.Set(i, j, color.Gray{Y: (*matrix)[i][j]})
		}
	}
	return newImage
}

func coloredToImage(matrix *RGBAMatrix, w, h int) *ebiten.Image {
	newImage := ebiten.NewImage(w, h)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			c := (*matrix)[i][j]
			newImage.Set(i, j, color.RGBA{
				R: c.r,
				G: c.g,
				B: c.b,
				A: c.a,
			})
		}
	}
	return newImage
}

func toMatrix(image *ebiten.Image) ([][]color.Color, int, int) {
	w, h := image.Size()
	var matrix = make([][]color.Color, w)
	for i := 0; i < w; i++ {
		matrix[i] = make([]color.Color, h)
		for j := 0; j < h; j++ {
			matrix[i][j] = image.At(i, j)
		}
	}
	return matrix, w, h
}

func floyd(matrix [][]color.Color, w, h int) *ebiten.Image {
	var image = ebiten.NewImage(w, h)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {

		}
	}
	return image
}

func toRGBPalette(matrix *RGBAMatrix, w, h int) *RGBAMatrix {
	result := make(RGBAMatrix, w)
	for i := 0; i < w; i++ {
		result[i] = make([]MyColor, h)
		for j := 0; j < h; j++ {
			c := (*matrix)[i][j]
			if c.r >= c.g && c.r >= c.b {
				result[i][j] = fromUint32(255, 0, 0, 255)
			}
			if c.g >= c.r && c.g >= c.b {
				result[i][j] = fromUint32(0, 255, 0, 255)
			}
			if c.b >= c.r && c.b >= c.g {
				result[i][j] = fromUint32(0, 0, 255, 255)
			}
		}
	}
	return &result
}

func combineLayers(matrix *RGBAMatrix, layerGray *GrayScaleMatrix, w, h int) *ebiten.Image {
	resultImage := ebiten.NewImage(w, h)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			c := (*matrix)[i][j]
			y := (*layerGray)[i][j]

			rf, gf, bf := float64(c.r)/255.0, float64(c.g)/255.0, float64(c.b)/255.0
			yf := float64(y) / 255.0

			rf *= yf
			gf *= yf
			bf *= yf
			rf *= 255
			gf *= 255
			bf *= 255

			resultImage.Set(i, j, color.RGBA{
				R: uint8(rf),
				G: uint8(gf),
				B: uint8(bf),
				A: 255,
			})
		}
	}
	return resultImage
}

func updateControls() {
	if ebiten.IsKeyPressed(ebiten.KeyDigit1) {
		trashold += 1
		toRecalculate = true
		if trashold > 255 {
			trashold = 0
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyDigit2) {
		trashold -= 1
		toRecalculate = true
		if trashold < 0 {
			trashold = 255
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyI) {
		invert = true
		toRecalculate = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyJ) {
		invert = false
		toRecalculate = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyT) {
		isTreshold = true
		toRecalculate = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyY) {
		isTreshold = false
		toRecalculate = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		whatToShow = ShowResult
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		whatToShow = ShowEdges
	}
	if ebiten.IsKeyPressed(ebiten.KeyG) {
		whatToShow = ShowGrayScale
	}
	if ebiten.IsKeyPressed(ebiten.KeyB) {
		whatToShow = ShowRGB
	}

	if ebiten.IsKeyPressed(ebiten.KeyP) && stopper {
		saveToPNG(*resultImg)
		stopper = false
	}

	if ebiten.IsKeyPressed(ebiten.KeyO) {
		stopper = true
	}

}

func (g Game) Draw(screen *ebiten.Image) {
	if toMakeGray {
		mainMatrix = makeMainMatrix(img)
		rgbPaletteImage = coloredToImage(toRGBPalette(mainMatrix, W, H), W, H)
		grayScale = grayscale(mainMatrix, W, H)
		grayScaleImage = grayScaleToImage(grayScale, W, H)
		edgeMatrix := findEdge(grayScale, W, H)
		edgeImage = grayScaleToImage(edgeMatrix, W, H)
		resultImg = combineLayers(mainMatrix, edgeMatrix, W, H)
		toRecalculate = false
		toMakeGray = false
	}

	updateControls()

	if toRecalculate {
		edgeMatrix := findEdge(grayScale, W, H)
		edgeImage = grayScaleToImage(edgeMatrix, W, H)
		resultImg = combineLayers(mainMatrix, edgeMatrix, W, H)
		toRecalculate = false
	}

	switch whatToShow {
	case ShowResult:
		screen.DrawImage(resultImg, nil)
		break
	case ShowEdges:
		screen.DrawImage(edgeImage, nil)
		break
	case ShowGrayScale:
		screen.DrawImage(grayScaleImage, nil)
		break
	case ShowRGB:
		screen.DrawImage(rgbPaletteImage, nil)
	}
}

func (g Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return W, H
}

func saveToPNG(p ebiten.Image) {
	w, h := p.Size()
	flag.Parse()
	rand.Seed(time.Now().UTC().UnixNano())

	out, err := os.Create("./output.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	render := image.NewRGBA(p.Bounds())
	draw.Draw(render, render.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			r, g, b, a := p.At(i, j).RGBA()

			fill := &image.Uniform{C: color.RGBA{
				R: uint8(r),
				G: uint8(g),
				B: uint8(b),
				A: uint8(a),
			}}

			draw.Draw(render, image.Rect(i, j, i+10, j+10), fill, image.ZP, draw.Src)
		}
	}

	err = png.Encode(out, render)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Generated image to output.png \n")
}
