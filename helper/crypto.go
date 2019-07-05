package helper

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"io"
)

func Hash512(input []byte, isNoise bool) (output string) {
	// noise hash
	if isNoise {
		_, _ = rand.Read(input)
	}

	// static hash
	hasher := sha512.New()
	_, _ = hasher.Write(input)
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func EncodeBase64(input string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(input))
	return encoded
}

func HashMD5(input string) string {
	h := md5.New()
	_, _ = io.WriteString(h, input)
	return string(h.Sum(nil))
}
