package main

import "fmt"

func getCenterPatch(text []byte) byte {
	var mask byte = 0xC0

	fmt.Printf("Mask: %b\nText: %b\n", mask, text[0])
	result := mask & text[0]
	result = result >> 6

	return result
}

func getHexColor(text []byte) string {
	color := fmt.Sprintf("#%X", text[0:3])
	return color
}

func getPatch(text []byte) (byte, byte, byte) {
	var cornerMask byte = 0xF0
	var sideMask byte = 0x0F

	cornerPatch := (cornerMask & text[3]) >> 4
	sidePatch := sideMask & text[3]

	return cornerPatch, sidePatch, byte(10)
}
