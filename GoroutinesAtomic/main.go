package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/signal"
	"sync/atomic"
	"time"
)

var count atomic.Int32

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
		if count.Load() < 5 && !gfStop {
			go hardWork(x)
		} else {
			fmt.Println("Sorry. We are busy. Try later")
		}
	}

	for count.Load() != 0 {
		fmt.Println("Waiting goroutines...")
		time.Sleep(10 * time.Second)
	}
}

func hardWork(x int32) {
	count.Add(1)
	time.Sleep(time.Second * time.Duration(x*rand.Int31n(10)))
	fmt.Printf("Goroutine with %d finished\n", x)
	count.Add(-1)
}
