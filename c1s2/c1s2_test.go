package c1s2

import "testing"

func TestEncodeStringReturnsCorrectValue(t *testing.T) {
	const (
		h1 = "1c0111001f010100061a024b53535009181c"
		h2 = "686974207468652062756c6c277320657965"
	)

    got, err := EncodeString(h1, h2)
    if err != nil {
        t.Error(err)
        return
    }

	if want := "746865206b696420646f6e277420706c6179"; got != want {
		t.Errorf("encodeString(h1, h2): got %q, want %q", got, want)
	}
}
