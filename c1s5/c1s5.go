package c1s5

func RepeatingKeyXOR(t []byte, k []byte) []byte {
	result := make([]byte, len(t))
	for i := 0; i < len(t); i++ {
		result[i] = t[i] ^ k[i%len(k)]
	}
	return result
}
