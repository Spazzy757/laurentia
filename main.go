package main

import (
	"github.com/Spazzy757/laurentia/handlers"
	"github.com/Spazzy757/laurentia/messages"
)

func main() {
	messageChan := make(chan string, 1)
	go messages.PubSubListener(messageChan)
	go func() {
		for {
			msg := <-messageChan
			messages.SaveMessage(msg)
			messages.VerifyMessageAndNotify(msg)
		}
	}()
	r := handlers.SetupRouter()
	r.Run("0.0.0.0:8000")
}
