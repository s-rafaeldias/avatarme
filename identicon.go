package main

import "github.com/fogleman/gg"

type Identicon struct {
	text        []byte
	baseColor   string
	invertColor string
	size        int
	dc          *gg.Context
}

//func (i *Identicon) drawCanvas() {
//	i.dc.NewContext(i.size, i.size)
//}
