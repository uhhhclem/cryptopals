package c1s7

import (
	"strings"
	"testing"
)

const filename = "ciphertext.txt"

func TestDecodeBase64File(t *testing.T) {
	_, err := DecodeBase64File(filename)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDecodeAES128ECB(t *testing.T) {
	key := []byte("YELLOW SUBMARINE")
	ct, err := DecodeBase64File(filename)
	if err != nil {
		t.Fatal(err)
	}
	pt, err := DecodeAES128ECB(ct, key)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(pt))
	if strings.Index(string(pt), "Play that funky music, white boy") < 0 {
		t.Fail()
	}
}
