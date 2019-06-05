package main

import "fmt"

func getHexColor(text []byte) string {
	color := fmt.Sprintf("#%X", text[0:3])
	return color
}

func getPatch(text []byte) (byte, byte, byte) {
	var cornerMask byte = 0xF0
	var sideMask byte = 0x0F
	var centerMask byte = 0xC0

	cornerPatch := (cornerMask & text[3]) >> 4
	sidePatch := sideMask & text[3]
	centerPatch := (centerMask & text[4]) >> 6

	return cornerPatch, sidePatch, centerPatch
}
