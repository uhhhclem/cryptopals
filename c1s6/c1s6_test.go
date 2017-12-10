package c1s6

import (
	"decrypt"
	"encoding/base64"
	"io/ioutil"
	"strings"
	"testing"
)

const (
	filename = "ciphertext.txt"
)

func TestHammingDistance(t *testing.T) {
	b1 := []byte("this is a test")
	b2 := []byte("wokka wokka!!!")
	got, err := HammingDistance(b1, b2)
	if err != nil {
		t.Fatal(err)
	}
	if want := 37; got != want {
		t.Errorf("HammingDistance(): got %d, want %d", got, want)
	}
}

func TestKeysizesByDistance(t *testing.T) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	b, err := base64.StdEncoding.DecodeString(string(f))
	if err != nil {
		t.Fatal(err)
	}

	cnts := map[int]int{}
	for cnt := 4; cnt <= 50; cnt++ {
		k, err := keysizesByDistance(b, cnt)
		if err != nil {
			t.Fatal(err)
		}
		cnts[k[0].size] += 1
	}
	if got, want := cnts[29], 41; got != want {
		t.Errorf("Expected keysize of 29 to appear %d times (got %d)", want, got)
	}
}

func TestFindKey(t *testing.T) {
	const keysize = 29
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	b, err := base64.StdEncoding.DecodeString(string(f))
	if err != nil {
		t.Fatal(err)
	}
	k, err := findKey(b, keysize)
	if err != nil {
		t.Fatal(err)
	}
	p := decrypt.RepeatingKeyXOR(b, k)
	if strings.Index(string(p), "Play that funky music, white boy") < 0 {
		t.Fail()
	}
}
