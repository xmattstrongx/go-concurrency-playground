package main

import (
	"fmt"
	"strings"
)

func main() {
	phrase := "These are the times that try men's souls\n"

	words := strings.Split(phrase, " ")

	ch := make(chan string, len(words))

	for _, word := range words {
		ch <- word
	}

	close(ch)

	// long form of the range loop below

	// for {
	// 	if msg, ok := <-ch; ok {
	// 		fmt.Print(msg + " ")
	// 	} else {
	// 		break
	// 	}
	// }

	// range over chan will terminate of channel is drained and closed
	for _ = range ch {
		fmt.Print(<-ch + " ")
	}
}
