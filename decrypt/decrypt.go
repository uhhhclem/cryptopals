// Package decrypt contains decryption functions developed during the cryptopals
// exercises.
package decrypt

import (
	"errors"
	"sort"
	"strings"
)

// ScoreMap is used for scoring texts; it contains a score that each character
// in the text is worth.
type ScoreMap map[byte]int

func DefaultScoreMap() ScoreMap {
	raw := map[string]int{
		"E": 1,
		"T": 1,
		"A": 1,
		"O": 1,
		"N": 1,
		"S": 1,
		"H": 1,
		"R": 1,
		"D": 1,
		"L": 1,
		"U": 1,
		" ": 1,
	}
	result := map[byte]int{}
	for k, v := range raw {
		result[[]byte(k)[0]] = v
		result[[]byte(strings.ToLower(k))[0]] = v
	}
	return result
}

// Score represents a block of text that has been scored via some method.
type Score struct {
	Value int
	Text  []byte
}

// Scores implements sort.Interface for sorting in reverse order by Value.
type Scores []Score

func (s Scores) Less(i, j int) bool { return s[i].Value < s[j].Value }
func (s Scores) Len() int           { return len(s) }
func (s Scores) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// TopScorers returns the top n scores in an array of Scores.
func TopScorers(s Scores, n int) []Score {
	sort.Sort(s)
	return s[len(s)-n : len(s)]
}

func ScorePlaintext(t []byte, m ScoreMap) int {
	result := 0
	for _, b := range t {
		result += m[b]
	}
	return result
}

// FindPlaintext finds the top n-ranked []bytes that can be found by XORING
// a byte with each byte in the cryptotext.
func FindPlaintext(c []byte, n int, m ScoreMap) [][]byte {
	var s Scores
	for i := 0; i < 256; i++ {
		b := EncryptSingleByteXOR(c, byte(i))
		s = append(s, Score{ScorePlaintext(b, m), b})
	}
	var result [][]byte
	ts := TopScorers(s, n)
	for _, r := range ts {
		result = append(result, r.Text)
	}
	return result
}

// FindSingleByteKey finds the byte that produces the highest-scoring plaintext
// when XORed with every byte of a cryptotext.
func FindSingleByteKey(c []byte, m ScoreMap) (byte, error) {
	var (
		key       byte
		highScore int
		found     bool
	)
	for i := 0; i < 256; i++ {
		b := EncryptSingleByteXOR(c, byte(i))
		s := ScorePlaintext(b, m)
		if s >= highScore {
			highScore = s
			key = byte(i)
			found = true
		}
	}
	if found {
		return key, nil
	}
	return 0, errors.New("no key found")
}

// EncryptSingleByteXOR encrypts t by XORing each byte in it with k.
func EncryptSingleByteXOR(t []byte, k byte) []byte {
	b := make([]byte, len(t))
	for i := range t {
		b[i] = t[i] ^ k
	}
	return b
}

// RepeatingKeyXOR transforms t by cyclically XORing each byte in it with
// a byte from k.
func RepeatingKeyXOR(t []byte, k []byte) []byte {
	result := make([]byte, len(t))
	for i := 0; i < len(t); i++ {
		result[i] = t[i] ^ k[i%len(k)]
	}
	return result
}
