package main

import "github.com/fogleman/gg"

type Identicon struct {
	text      []byte
	baseColor string
	size      int
	dc        *gg.Context
}
