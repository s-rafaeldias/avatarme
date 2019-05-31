package main

import "fmt"

func getCenterPatch(text []byte) byte {
	var mask byte = 0xC0

	fmt.Printf("Mask: %b\nText: %b\n", mask, text[0])
	result := mask & text[0]
	result = result >> 6

	return result
}
