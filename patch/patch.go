package patch

import "fmt"

func GetHexColor(text []byte) string {
	color := fmt.Sprintf("#%X", text[0:3])
	return color
}

func GetPatch(text []byte) (byte, byte, byte) {
	var cornerMask byte = 0xF0
	var sideMask byte = 0x0F
	var centerMask byte = 0xC0

	cornerPatch := (cornerMask & text[3]) >> 4
	sidePatch := sideMask & text[3]
	centerPatch := (centerMask & text[4]) >> 6

	return cornerPatch, sidePatch, centerPatch
}

func GetColorInvert(text []byte) byte {
	var mask byte = 0x02

	invert := (text[4] & mask) >> 1
	return invert
}
