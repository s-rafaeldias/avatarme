package util

import "crypto/md5"

func GetMD5(text string) [md5.Size]byte {
	data := []byte(text)
	return md5.Sum(data)
}
