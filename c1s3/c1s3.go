package c1s3

import (
	"sort"
	"strings"
)

var scoreMap map[byte]int

func init() {
	scoreMap = makeScoreMap()
}

func makeScoreMap() map[byte]int {
	raw := map[string]int {
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
	}
	result := map[byte]int{}
	for k, v := range raw {
		result[[]byte(k)[0]] = v
		result[[]byte(strings.ToLower(k))[0]] = v
	}
	return result
}

type Score struct {
	Value int
	Text []byte
}
type Scores []Score

// TopScorers returns the top n scores in an array of Scores.
func TopScorers(s Scores, n int) []Score {
	sort.Sort(s)
	return s[len(s)-n: len(s)]
}

func (s Scores) Less(i, j int) bool { return s[i].Value < s[j].Value }
func (s Scores) Len() int { return len(s) }
func (s Scores) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func ScorePlaintext(t []byte) int {
	result := 0
	for _, b := range t {
		result += scoreMap[b]
	}
	return result
}

// FindPlaintext finds the top n-ranked []bytes that can be found by XORING
// a byte with each byte in the cryptotext.
func FindPlaintext(c []byte, n int) [][]byte {
	var s Scores
	for i := 0; i < 256; i++ {
		b := EncryptSingleByteXOR(c, byte(i))
		s = append(s, Score{ScorePlaintext(b), b})
	}
	var result [][]byte
	ts := TopScorers(s, n)
	for _, r := range ts {
		result = append(result, r.Text)
	}
	return result
}

func EncryptSingleByteXOR(t []byte, k byte) []byte {
	b := make([]byte, len(t))
	for i := range t {
		b[i] = t[i]^k
	}
	return b	
}