package main

import (
	"bytes"
	"fmt"
	"time"
)

const length = 6

var pwd = []byte{'z', '1', 'y', '2', '2', '4'}
var abc = []byte{'x', 'y', 'z', '1', '2', '3', '4'}

var z [][]byte

func main() {
	start := time.Now()
	makeQueue(nil)
	for _, bs := range z {
		if bytes.Equal(pwd, bs) {
			fmt.Println("Bingo ", string(bs))
			break
		}
	}
	fmt.Println(time.Since(start))
}

func makeQueue(pref []byte) {
	for _, b := range abc {
		if len(pref) == length {
			break
		}
		x := makeInput(pref, b)
		z = append(z, x)
		makeQueue(x)

	}
}

func makeInput(pref []byte, b byte) []byte {
	res := make([]byte, len(pref), len(pref)+1)
	copy(res, pref)
	res = append(res, b)
	return res
}
