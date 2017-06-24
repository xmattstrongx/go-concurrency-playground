package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

const logFile = "./log.txt"

func main() {
	runtime.GOMAXPROCS(4)

	f, _ := os.Create(logFile)
	f.Close()

	logCh := make(chan string, 50)

	go func() {
		for {
			msg, ok := <-logCh
			if ok {
				f, _ := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

				logTime := time.Now().Format(time.RFC3339)
				if _, err := f.WriteString(logTime + " - " + msg); err != nil {
					fmt.Println(err)
				}
				f.Close()
			} else {
				break
			}
		}
	}()

	// mutex := new(sync.Mutex)
	mutex := make(chan bool, 1)

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			// mutex.Lock()
			mutex <- true
			go func() {
				msg := fmt.Sprintf("%d + %d = %d\n", i, j, i+j)
				logCh <- msg
				fmt.Print(msg)
				// mutex.Unlock()
				<-mutex
			}()
		}
	}
}
