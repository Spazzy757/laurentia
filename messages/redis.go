package messages

import (
	"os"
	"time"
	"github.com/go-redis/redis"
	"fmt"
	"log"
	"strings"
	"encoding/json"
)


var redisHost = os.Getenv("REDIS_HOST") + ":6379"
var redisPassword = os.Getenv("REDIS_PASSWORD")
var channel = os.Getenv("REDIS_CHANNEL")


type DynamicMessage struct {
	Key string `json:"key"`
	ID  string `json:"id"`
	Payload  interface{} `json:"payload"`
}

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

func CheckAcknowledgment(message string) error{
	client := GetClient()
	message = strings.Replace(message, `'`, `"`, -1)
	var m DynamicMessage
	if err := json.Unmarshal([]byte(message), &m); err != nil {
		log.Fatal(err)
	}
	return nil
}


func GetSMembers(descriptor string) []string{
	client := GetClient()
	sMem := client.SMembers(descriptor)
	log.Println(sMem)
	return sMem.Val()
}
