package messages

import (
	"os"
	"time"
	"github.com/go-redis/redis"
	"fmt"
	"log"
)


var redisHost = os.Getenv("REDIS_HOST") + ":6379"
var redisPassword = os.Getenv("REDIS_PASSWORD")
var channel = os.Getenv("REDIS_CHANNEL")


func GetClient() *redis.Client{
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       1,
	})
	return client
}


func PubSubListener(pubSubChannel chan string) {
	client := GetClient()
	pubSub := client.Subscribe(channel)
	defer pubSub.Close()
	for {
		msg, err := pubSub.ReceiveMessage()
		if err != nil {log.Println("Pub Sub Failed to Receive Message")}
		if msg != nil {fmt.Println(msg)}
		pubSubChannel <- msg.Payload
		time.Sleep(1) // Be Kind To Redis
	}
}


func PubSubSendMessage(message string) error{
	client := GetClient()
	err := client.Publish(channel, message).Err()
	if err != nil {return err}
	return nil
}


func GetSMembers(descriptor string) []string{
	client := GetClient()
	sMem := client.SMembers(descriptor)
	return sMem.Val()
}
