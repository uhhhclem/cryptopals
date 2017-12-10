package c1s7

import (
	"crypto/aes"
	"encoding/base64"
	"io/ioutil"
)

// DecodeBase64File reads base64-encoded data from filename and returns it as
// a decoded byte slice, padded to a multiple of blocksize bytes.
func DecodeBase64File(filename string) ([]byte, error) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	b, err := base64.StdEncoding.DecodeString(string(f))
	if err != nil {
		return nil, err
	}
	return b, nil
}

func DecodeAES128ECB(ct, key []byte) ([]byte, error) {
	cb, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	var pt []byte
	cnt := len(ct) / cb.BlockSize()
	for i := 0; i < cnt; i += 1 {
		src := ct[i*cb.BlockSize() : (i+1)*cb.BlockSize()]
		dst := make([]byte, cb.BlockSize())
		cb.Decrypt(dst, src)
		pt = append(pt, dst...)
	}
	return pt, nil
}
