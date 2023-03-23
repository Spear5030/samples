package main

import (
	"bytes"
	"fmt"
	"time"
)

const length = 6

var pwd = []byte{'z', '1', 'y', '2', '2', '4'}
var abc = []byte{'x', 'y', 'z', '1', '2', '3', '4'}
var queue chan []byte

func main() {
	start := time.Now()
	queue = make(chan []byte, 10)
	go makeQueue(nil)
	go func() {
		for {
			select {
			case x := <-queue:
				if bytes.Equal(pwd, x) {
					fmt.Println("Bingo ", string(x))
					fmt.Println(time.Since(start))
					break
				}
			}
		}
	}()
	time.Sleep(time.Second * 10)
}

func makeQueue(pref []byte) {
	for _, b := range abc {
		if len(pref) == length {
			return
		}
		x := makeInput(pref, b)
		queue <- x
		makeQueue(x)
	}
}

func makeInput(pref []byte, b byte) []byte {
	res := make([]byte, len(pref), len(pref)+1)
	copy(res, pref)
	res = append(res, b)
	return res
}
