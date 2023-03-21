package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"time"
)

type wgCount struct {
	sync.WaitGroup
	count atomic.Int32
}

func (wgCount *wgCount) Add(n int32) {
	wgCount.WaitGroup.Add(int(n))
	wgCount.count.Add(int32(n))
}

func (wgCount *wgCount) Done(n int32) {
	wgCount.WaitGroup.Done()
	wgCount.count.Add(-1)
}

var wg wgCount

func main() {
	var x int32
	var gfStop bool
	sigint := make(chan os.Signal)
	signal.Notify(sigint, os.Interrupt)

	go func() {
		<-sigint
		fmt.Println("Gracefully shutdown")
		gfStop = true
	}()

	for {
		_, err := fmt.Scan(&x)
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println("EOF detected")
				break
			}
			fmt.Println(err)
			continue
		}
		if wg.count.Load() < 5 && !gfStop {
			go hardWork(x)
		} else {
			fmt.Println("Sorry. We are busy. Try later")
		}
	}
	wg.Wait()
}

func hardWork(x int32) {
	wg.Add(1)
	time.Sleep(time.Second * time.Duration(x*rand.Int31n(10)))
	fmt.Printf("Goroutine with %d finished\n", x)
	wg.Add(-1)
}
