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
	// keysize to be in the top 5 candidates for all block counts.  (Indeed,
	// experimentation shows that it's in the top 4 candidates for all block
	// counts except 30.)
	want := 20
	for cnt := 4; cnt <= 50; cnt++ {
		k, err := keysizesByDistance(b, cnt)
		if err != nil {
			t.Fatal(err)
		}
		found := false
		for i := range k[:5] {
			found = found || k[i].size == want
		}
		if !found {
			t.Errorf("keysizesByDistance(b, %d): keysize of %d not found in top 4\n%v", cnt, want, k)
		}
	}
}

func TestFindKey(t *testing.T) {
	const keysize = 20
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	k := findKey(b, keysize)
	t.Logf(k)
	t.Fail()
}
