package helpers

import (
	cypto "crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
)

var letterRunes = []rune("AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz")

func GenerateHash() (string, error) {
	b := make([]byte, 32)
	_, err := cypto.Read(b)

	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(b)

	return hex.EncodeToString(hash[:]), nil
}

func GenerateStringRune(n int) string {
	res := make([]rune, n)

	for i := range res {
		res[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(res)
}
