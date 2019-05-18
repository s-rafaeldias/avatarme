package main

import (
	"crypto/md5"
	"fmt"
)

func getMD5(data []byte) [md5.Size]byte {
	return md5.Sum(data)
}

func main() {
	text := getMD5([]byte("teste"))
	fmt.Printf("%x\n", text)
}
