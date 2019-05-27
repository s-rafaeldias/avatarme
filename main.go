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

func drawSide(dc *gg.Context) {
	// Draw top
	drawPane(dc, 200, 0, 1, 0, 0)
	// Draw right
	drawPane(dc, 400, 200, 1, 0, 0)
	// Draw bottom
	drawPane(dc, 200, 400, 1, 0, 0)
	// Draw left
	drawPane(dc, 0, 200, 1, 0, 0)
}

func drawCenter(dc *gg.Context) {
	drawPane(dc, 200, 200, 0, 1, 0)
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
	drawSide(dc)
	drawCenter(dc)

	dc.SavePNG("teste.png")

}
