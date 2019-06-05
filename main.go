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

	dc.SetLineWidth(1)
	dc.SetHexColor(color)
	dc.StrokePreserve()
	dc.SetHexColor(color)
	dc.Fill()

	dc.Pop()
}

func main() {
	// Generate MD5
	text := getMD5([]byte("aqueleteste"))
	fmt.Printf("%b\n", text)

	// Get base color
	color := getHexColor(text[:])
	cornerPatch, sidePatch, centerPatch := getPatch(text[:])
	fmt.Printf("Corner patch: %d\nSide patch: %d\nCenter patch: %d\n",
		cornerPatch, sidePatch, centerPatch)

	// Constants
	const Size = 600

	// identicon is 600x600
	dc := gg.NewContext(Size, Size)
	dc.SetHexColor("#FFFFFF")
	dc.Clear()

	drawPane(dc, 1., 2., color)
	//drawCorner(dc)
	//drawSide(dc)
	//drawCenter(dc)

	dc.SavePNG("teste.png")

}
