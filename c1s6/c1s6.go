package c1s6

import (
	"errors"
	"sort"
)

var errInvalidArgs = errors.New("invalid arguments")

// HammingDistance calculates the Hamming distance between two byte slices
// of the same length.
func HammingDistance(b1, b2 []byte) (int, error) {
	if b1 == nil || b2 == nil || len(b1) != len(b2) {
		return 0, errInvalidArgs
	}
	d := 0
	for i := range b1 {
		d += bitwiseHammingDistance(b1[i], b2[i])
	}
	return d, nil
}

// This is cribbed from Wikipedia.  The real trick here is that if k is
// nonzero, k &= k-1 clears the least significant 1 bit of k.
func bitwiseHammingDistance(b1, b2 byte) int {
	d := 0
	for val := b1 ^ b2; val != 0; val &= val - 1 {
		d++
	}
	return d
}

type keysize struct {
	size  int
	ndist float64
}

type keysizes []keysize

func (k keysizes) Less(i, j int) bool { return k[i].ndist < k[j].ndist }
func (k keysizes) Len() int           { return len(k) }
func (k keysizes) Swap(i, j int)      { k[i], k[j] = k[j], k[i] }

// keysizesByDistance returns a slice of keysizes sorted by the average normalized
// edit distance between adjacent pairs of the first cnt blocks of b.
func keysizesByDistance(b []byte, cnt int) (keysizes, error) {
	const maxKeysize = 40
	if b == nil || cnt < 2 || len(b) < maxKeysize*cnt {
		return nil, errInvalidArgs
	}
	var result keysizes
	for size := 2; size <= maxKeysize; size++ {
		// get cnt adjacent blocks of size bytes
		blocks := make([][]byte, cnt)
		for i := 0; i < cnt; i++ {
			blocks[i] = b[i*size : (i+1)*size]
		}
		// sum up the distance between each of the three adjacent block
		// pairs
		dist := 0
		for i := 0; i < cnt-1; i++ {
			d, err := HammingDistance(blocks[i], blocks[i+1])
			if err != nil {
				return nil, err
			}
			dist += d
		}
		// normalized distance is the average of the cnt-1 distances divided
		// by the key size.
		result = append(result,
			keysize{
				size:  size,
				ndist: (float64(dist) / float64(cnt-1)) / float64(size),
			})
	}
	sort.Sort(result)
	return result, nil
}
