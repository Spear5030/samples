package main

import (
	"context"
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
	sigint := make(chan os.Signal)
	signal.Notify(sigint, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-sigint
		fmt.Println("Shutdown with last goroutine")
		cancel()
	}()

	queue := make(chan int32, 5)
	go hardWorker(ctx, queue)

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
		queue <- x
	}

	for count.Load() != 0 {
		fmt.Println("Waiting goroutines...")
		time.Sleep(10 * time.Second)
	}
}

func hardWorker(ctx context.Context, q chan int32) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Goroutine aborted from context\n")
			return
		case x := <-q:
			count.Add(1)
			time.Sleep(time.Second * time.Duration(x*rand.Int31n(5)))
			fmt.Printf("Goroutine with %d finished\n", x)
			count.Add(-1)
		}
	}
}
