package main

import (
	"crypto/md5"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func getMD5(data []byte) [md5.Size]byte {
	return md5.Sum(data)
}

func main() {
	// Generate MD5
	text := getMD5([]byte("teste"))
	fmt.Printf("%x\n", text)

	// The identicon will be 60x60, so each pane is 20x20, making an identicon a 3 pane x 3 pane
	filename := "teste.png"

	// (x0, y0, x1, y1)
	img := image.NewRGBA(image.Rect(0, 0, 200, 200))
	green := color.RGBA{0, 100, 0, 255}
	red := color.RGBA{100, 0, 0, 255}

	// Coloring identicon as red
	draw.Draw(img, img.Bounds(), &image.Uniform{red}, image.ZP, draw.Src)
	for i := 0; i <= 200; i++ {
		draw.Draw(img, image.Rect(i, i, 0, 200), &image.Uniform{green}, image.ZP, draw.Src)
	}

	// Wrinting to file
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Deu ruim")
	}
	png.Encode(file, img)
}
