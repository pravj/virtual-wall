package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"math/rand"
	"os"
	"time"
)

const (
	ImageWidth  = 1010
	ImageHeight = 390
	MaxColumns  = 10
	MaxRows     = 11
	TileWidth   = 90
	TileHeight  = 30
	SepX        = 10
	SepY        = 5
)

func main() {

	img := image.NewRGBA(image.Rect(0, 0, ImageWidth, ImageHeight))
	fillBackground(img)

	for cols := 0; cols < MaxColumns; cols++ {
		for rows := 0; rows < MaxRows; rows++ {
			drawTile(rows, cols, img, false)
		}
	}

	i := 0
	for i <= 200 {
		drawTile(0, 0, img, true)
		i = i + 1
	}

	saveImage(img)
}

func colors() (uint8, uint8, uint8) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return uint8(r.Intn(255)), uint8(r.Intn(255)), uint8(r.Intn(255))
}

func drawTile(i, j int, img *image.RGBA, random bool) {

	startX := 0
	startY := 0

	if random {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		startX = r.Intn(970) + 10
		startY = r.Intn(350) + 5
	} else {
		startX = SepX + j*(TileWidth+SepX)
		startY = SepY + i*(TileHeight+SepY)
	}

	R, G, B := colors()

	index := 0
	for y := startY; y < (startY + TileHeight); y++ {
		d := (index % (TileHeight / 2))

		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		gapX := r.Intn(SepX)

		for x := (startX - gapX - d); x < (startX + TileWidth + gapX + d); x++ {
			img.Set(x, y, color.RGBA{R, G, B, 0xff})
		}

		index = index + 1
	}

	for x := startX; x < (startX + TileWidth); x++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		gapY := r.Intn(SepY)

		for y := (startY - gapY); y < (startY + TileHeight + gapY); y++ {
			img.Set(x, y, color.RGBA{R, G, B, 0xff})
		}
	}
}

func saveImage(img *image.RGBA) {
	file, err := os.Create("virtual-wall.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	jpeg.Encode(file, img, &jpeg.Options{100})
}

func fillBackground(img *image.RGBA) {
	for y := img.Rect.Min.Y; y < img.Rect.Max.Y; y++ {
		for x := img.Rect.Min.X; x < img.Rect.Max.X; x++ {
			img.Set(x, y, color.RGBA{0xfd, 0xfd, 0xfd, 0})
		}
	}
}
