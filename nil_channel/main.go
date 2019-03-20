package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}

func merge(a, b <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)

		for a != nil || b != nil {
			select {
			case v, ok := <-a:
				if !ok {
					a = nil
					log.Println("A is done")
					continue
				}
				ch <- v
			case v, ok := <-b:
				if !ok {
					b = nil
					log.Println("B is done")
					continue
				}
				ch <- v
			}
		}
	}()
	return ch
}

func asChan(vs ...int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, v := range vs {
			ch <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(ch)
	}()
	return ch
}
