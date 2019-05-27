package main

import (
	"crypto/md5"
	"fmt"

	"github.com/fogleman/gg"
)

func getMD5(data []byte) [md5.Size]byte {
	return md5.Sum(data)
}

func drawPane(dc *gg.Context, x, y, r, g, b float64) {
	dc.Push()
	dc.SetRGB(r, g, b)
	dc.DrawRectangle(x, y, 200, 200)
	dc.Fill()
	dc.Pop()
}

func drawCorner(dc *gg.Context) {
	// Draw top left
	drawPane(dc, 0, 0, 0, 0, 1)
	// Draw top right
	drawPane(dc, 400, 0, 0, 0, 1)
	// Draw bottom left
	drawPane(dc, 0, 400, 0, 0, 1)
	// Draw bottom right
	drawPane(dc, 400, 400, 0, 0, 1)

}

func main() {
	// Generate MD5
	text := getMD5([]byte("teste"))
	fmt.Printf("%x\n", text)

	// Constants
	const Size = 600

	// identicon is 600x600
	dc := gg.NewContext(Size, Size)

	drawCorner(dc)

	dc.SaveJPG("teste.jpg", 100)

}
