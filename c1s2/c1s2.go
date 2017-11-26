package c1s2

import (
	"encoding/hex"
    "errors"
)

// EncodeString produces an encoded string by XORing the bytes in two strings together.
func EncodeString(s1, s2 string) (string, error) {
	b1, err := hex.DecodeString(s1)
	if err != nil {
		return "", err
	}
	b2, err := hex.DecodeString(s2)
	if err != nil {
		return "", err
	}
	if len(b1) != len(b2) {
		return "", errors.New("Invalid input:  must be byte slices of equal length")
	}
	b3 := make([]byte, len(b1))
	for i := range b1 {
		b3[i] = b1[i]^b2[i]
	}
	return hex.EncodeToString(b3), nil
}