package helper

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
)

func Hash256(input []byte, isNoise bool) (output string) {
	// noise hash
	if isNoise {
		rand.Read(input)
	}

	// static hash
	hasher := sha512.New()
	hasher.Write(input)
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
