package main

import (
	"github.com/Spazzy757/laurentia/handlers"
	"github.com/Spazzy757/laurentia/messages"
	"log"
)

func main() {
	messageChan := make(chan string)
	go messages.PubSubListener(messageChan)
	go func() {
		var msg string 
		for {
			msg = <-messageChan
			log.Println(`*************************************************`)
			log.Println("Main Message Routine")
			log.Println(msg)
			log.Println(messages.SaveMessage(msg))
			messages.VerifyMessageAndNotify(msg)
			handlers.AddMessageToChannel(msg)
			log.Println(`*************************************************`)
		}
	}()
	r := handlers.SetupRouter()
	r.Run("0.0.0.0:8000")
}


