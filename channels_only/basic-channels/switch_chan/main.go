package main

import (
	"fmt"
	"time"
)

func main() {
	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)

	msg := Message{
		To:      []string{"frodo@underhill.shire"},
		From:    "gandalf@whitecouncil.org",
		Content: "Keep it secret, keep it safe.",
	}

	failedMessage := FailedMessage{
		ErrorMessage:    "Message intercepted by black rider",
		OriginalMessage: Message{},
	}

	msgCh <- msg
	errCh <- failedMessage

	for {
		select {
		case receivedMsg := <-msgCh:
			fmt.Println(receivedMsg)
		case receivedError := <-errCh:
			fmt.Println(receivedError)
		case <-time.After(50 * time.Millisecond):
			fmt.Println("Timeout")
			return
		}
	}
}

type Message struct {
	To      []string
	From    string
	Content string
}

type FailedMessage struct {
	ErrorMessage    string
	OriginalMessage Message
}
