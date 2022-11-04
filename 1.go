package main

import (
	"fmt"
	"sync"
)

var x int64
var y int64
var wg sync.WaitGroup
var ch = make(chan int64, 1)

func add() {
	y = <-ch
	for i := 0; i < 50000; i++ {
		y = y + 1
	}
	ch <- y
	wg.Done()
}
func main() {
	ch <- x
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	x = <-ch
	fmt.Println(x)
}
