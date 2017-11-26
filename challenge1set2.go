package main

import (
    "c1s2"
	"fmt"
)

const (
	h1 = "1c0111001f010100061a024b53535009181c"
	h2 = "686974207468652062756c6c277320657965"
)

func main() {
    fmt.Println(c1s2.EncodeString(h1, h2))
}
