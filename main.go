package main

import (
	"fmt"

	"github.com/fogleman/gg"
)

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
	// Constants
	const Size = 600

	// Generate MD5
	text := getMD5([]byte("teste"))
	fmt.Printf("%b\n", text)

	// Get base color
	baseColor := getHexColor(text[:])
	invertColor := getColorInvert(text[:])
	cornerPatch, sidePatch, centerPatch := getPatch(text[:])

	fmt.Printf("Corner patch: %d\nSide patch: %d\nCenter patch: %d\n",
		cornerPatch, sidePatch, centerPatch)
	fmt.Printf("Base color: %s\nInvert color? %d\n", baseColor, invertColor)

	// identicon is 600x600
	dc := gg.NewContext(Size, Size)
	dc.SetHexColor("#FFFFFF")
	dc.Clear()

	drawPane(dc, 1., 2., baseColor)

	dc.SavePNG("teste.png")

}
