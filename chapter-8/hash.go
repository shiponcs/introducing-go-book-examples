package main

import (
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func getHash(fileName string) (uint32, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	h := crc32.NewIEEE()
	if _, err := io.Copy(h, f); err != nil {
		return 0, err
	}
	return h.Sum32(), nil
}

func main() {
	h1, err := getHash("test.txt")
	if err != nil {
		return
	}
	h2, err := getHash("test1.txt")
	if err != nil {
		return
	}

	fmt.Println(h1, h2, h1 == h2)
}
