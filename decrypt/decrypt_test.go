package decrypt

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestRepeatingKeyXOR(t *testing.T) {
	pt := []byte(`Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`)
	key := []byte("ICE")
	ct := RepeatingKeyXOR(pt, key)
	want :=
		`0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272` +
			`a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f`
	got := hex.EncodeToString(ct)
	if got != want {
		t.Errorf("RepeatingKeyXOR(%q) failed:\n  got: %s\n  want: %s", string(pt), got, want)
	}
}

func TestScorePlaintext(t *testing.T) {
	m := DefaultScoreMap()
	tests := []struct {
		text []byte
		want int
	}{
		{[]byte("n290g1-39xdbq-580109dja'9371-s01"), 5},
		{[]byte("Every good boy deserves favour"), 22},
	}

	for _, tt := range tests {
		if got, want := ScorePlaintext(tt.text, m), tt.want; got != want {
			t.Errorf("ScorePlaintext(%q): got %d, want %d", string(tt.text), got, tt.want)
		}
	}
}

func TestFindPlaintext(t *testing.T) {
	const (
		h    = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
		want = "Cooking MC's like a pound of bacon"
	)
	m := DefaultScoreMap()
	b, err := hex.DecodeString(h)
	if err != nil {
		t.Fatalf("DecodeString returned error: %s", err)
	}

	top := FindPlaintext(b, 5, m)

	for _, text := range top {
		if bytes.Compare(text, []byte(want)) == 0 {
			return
		}
	}
	for _, b := range top {
		t.Logf("%s\n", string(b))
	}
	t.Error("Could not find plaintext.")
}

func TestFindSingleByteKey(t *testing.T) {
	tests := []struct {
		key  int
		text string
	}{
		{71, "here we are now, entertain us, i feel stupid and contagious"},
		{2, "now is the winter of our discontent made glorious summer by this son of york"},
		{42, "i haven't seen any citizen say, 'hey, wait just a minute'"},
		{111, "sumer is icumen in, lhude sing cuccu!"},
		{41, "this space intentionally left blank"},
		{98, "arnold layne had a strange hobby"},
		{102, "evenatextwithnospacesshouldbedecryptable"},
	}
	m := DefaultScoreMap()
	for _, tt := range tests {
		c := EncryptSingleByteXOR([]byte(tt.text), byte(tt.key))
		got, err := FindSingleByteKey(c, m)
		if err != nil {
			t.Fatal("FindSingleByteKey() returned %s", err)
		}
		if got, want := got, byte(tt.key); got != want {
			t.Errorf("FindSingleByteKey(): got key of %d, want %d", got, want)
		}
	}
}
