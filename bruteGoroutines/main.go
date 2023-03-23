package main

import (
	"bytes"
	"fmt"
	"sync"
	"time"
)

const length = 8

var pwd = []byte{'z', '1', 'y', '2', '2', '4', '5'}
var abc = []byte{'x', 'y', 'z', '1', '2', '3', '4', '5'}
var queue chan []byte
var done chan struct{}
var wg sync.WaitGroup

func main() {
	start := time.Now()
	queue = make(chan []byte, 10)
	done = make(chan struct{})
	go makeQueue(nil)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(&wg, done)
	}
	wg.Wait()
	fmt.Println(time.Since(start))
}

func worker(wg *sync.WaitGroup, done chan struct{}) {
	for {
		select {
		case x := <-queue:
			if bytes.Equal(pwd, x) {
				fmt.Println("Bingo ", string(x))
				close(done)
			}
		case <-done:
			wg.Done()
			return
		}
	}
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
