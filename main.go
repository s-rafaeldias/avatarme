package main

import (
	"crypto/md5"
	"fmt"

	"github.com/fogleman/gg"
)

func getMD5(data []byte) [md5.Size]byte {
	return md5.Sum(data)
}

func main() {
	// Generate MD5
	text := getMD5([]byte("teste"))
	fmt.Printf("%x\n", text)

	filename := "teste.png"
	const Size = 500
	dc := gg.NewContext(Size, Size)
	dc.DrawCircle(250., 250., 50.)
	dc.SetRGB(1, 1, 1)
	dc.Fill()
	dc.SavePNG(filename)
}
