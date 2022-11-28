package util

import (
	"crypto/md5"
	"encoding/hex"
)

func EncodeMD5(str string) (encode string) {
	md5 := md5.New()
	md5.Write([]byte(str))
	return hex.EncodeToString(md5.Sum(nil))
}
