package c1s6

import (
	"io/ioutil"
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
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	// For this cryptotext, the keysize appears to be 20.  So we expect that
	// keysize to be in the top 4 candidates for a variety of block counts.
	want := 20
	for _, cnt := range []int{4, 10, 30, 50} {
		k, err := keysizesByDistance(b, cnt)
		if err != nil {
			t.Fatal(err)
		}
		found := false
		for i := range k {
			found = found || k[i].size == want
		}
		if !found {
			t.Errorf("keysizesByDistance(b, %d): keysize of %d not found in top 4\n%v", cnt, want, k)
		}
	}
}
