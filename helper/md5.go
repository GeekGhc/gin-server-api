package helper

import (
	"crypto/md5"
	"encoding/hex"
)

// md5 encrypting
func EncodeMD5(value string) string{
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}