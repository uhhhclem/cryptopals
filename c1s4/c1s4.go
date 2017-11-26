package c1s4

import (
	"c1s3"
	"encoding/hex"
	"sort"
	"strings"
)

func DecodeFile(f string) ([][]byte, error) {
	var result [][]byte
	for _, s := range strings.Split(f, "\n") {
		if strings.TrimSpace(s) == "" {
			continue
		}
		b, err := hex.DecodeString(s)
		if err != nil {
			return nil, err
		}
		result = append(result, b)
	}
	return result, nil
}

func FindPlaintext(f [][]byte, n int) c1s3.Scores {
	var s c1s3.Scores
	for _, b := range f {
		for i := 0; i < 256; i++ {
			c := c1s3.EncryptSingleByteXOR(b, byte(i))
			s = append(s, c1s3.Score{c1s3.ScorePlaintext(c), c})
		}
	}
	sort.Sort(s)
	var result c1s3.Scores
	for i := n; i > 0; i-- {
		result = append(result, s[len(s)-i])
	}
	return result
}
