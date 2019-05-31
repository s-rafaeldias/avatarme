package main

import (
	"crypto/md5"
	"fmt"

	"github.com/fogleman/gg"
)

func getMD5(data []byte) [md5.Size]byte {
	return md5.Sum(data)
}

func drawPane(dc *gg.Context, x, y float64, color string) {
	dc.Push()

	dc.LineTo(0, 0)
	dc.LineTo(0, 600)
	dc.LineTo(600, 0)

	dc.SetLineWidth(10)
	dc.SetHexColor("#000000")
	dc.StrokePreserve()
	dc.SetHexColor("#000000")
	dc.Fill()

	dc.Pop()
}

func main() {
	// Generate MD5
	text := getMD5([]byte("raf197795@gmail.c"))
	fmt.Printf("%b\n", text)

	centerPatch := getCenterPatch(text[:])
	fmt.Println("Center patch:", centerPatch)

	// Constants
	const Size = 600

	// identicon is 600x600
	dc := gg.NewContext(Size, Size)
	dc.SetHexColor("#FFFFFF")
	dc.Clear()

	drawPane(dc, 1., 2., "#FFCC00")
	//drawCorner(dc)
	//drawSide(dc)
	//drawCenter(dc)

	dc.SavePNG("teste.png")

}
