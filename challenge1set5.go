package main

import (
	"fmt"
	"io"
	"os"

	"c1s5"
)

const (
	blockSize = 1000
	key       = "superseekritkey"
)

func main() {
	for {
		pt := make([]byte, blockSize)
		n, err := os.Stdin.Read(pt)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		ct := c1s5.RepeatingKeyXOR(pt[:n], []byte(key))
		if _, err := os.Stdout.Write(ct); err != nil {
			fmt.Println(err)
			return
		}
	}
}
