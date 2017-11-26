package c1s3

import (
    "bytes"
    "encoding/hex"
    "testing"
)

func TestFindPlaintext(t *testing.T) {
    const (
        h = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
        want = "Cooking MC's like a pound of bacon"
    )
    
    b, err := hex.DecodeString(h)
    if err != nil {
        t.Fatalf("DecodeString returned error: %s", err)
    }
    
    top := FindPlaintext(b, 5)
    
    for _, text := range top {
        if bytes.Compare(text, []byte(want)) == 0 {
            return
        }
    }
    for _, b := range top  {
        t.Logf("%s\n", string(b))
    }
    t.Error("Could not find plaintext.")
}

func TestScorePlaintext(t *testing.T) {
	tests := []struct {
		text  []byte
		want int
	}{
		{[]byte("n290g1-39xdbq-580109dja'9371-s01"), 5},
		{[]byte("Every good boy deserves favour"), 18},
	}

	for _, tt := range tests {
		if got, want := ScorePlaintext(tt.text), tt.want; got != want {
			t.Errorf("ScorePlaintext(%q): got %d, want %d", string(tt.text), got, tt.want)
		}
	}
}
