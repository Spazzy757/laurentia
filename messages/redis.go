package messages

import (
	"os"
	"time"
	"github.com/go-redis/redis"
	"fmt"
	"log"
	"strings"
	"encoding/json"
	"net/http"
	"bytes"
)


var redisHost = os.Getenv("REDIS_HOST") + ":6379"
var redisPassword = os.Getenv("REDIS_PASSWORD")
var channel = os.Getenv("REDIS_CHANNEL")


type DynamicMessage struct {
	Key string `json:"key"`
	ID  string `json:"id"`
	Payload  interface{} `json:"payload"`
}

type Acknowledgement struct {
	Event string
	ID string
	Successful []string
	Failure []string
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

func CheckAcknowledgment(message string) (Acknowledgement, error){
	message = strings.Replace(message, `'`, `"`, -1)
	var m DynamicMessage
	if err := json.Unmarshal([]byte(formatPythonDict(message)), &m); err != nil {
		return Acknowledgement{}, err
	}
	subscriberLookUp := "pubsub.events.order." + m.Key + ".subscribers"
	subscribersList := GetSMembers(subscriberLookUp)
	ackLookUp := "pubsub.events.actions." + m.Key + "." + m.ID + ".received"
	ackList := GetSMembers(ackLookUp)
	successList := make([]string, 0)
	failureList := make([]string, 0)
	for _, v := range subscribersList {
		check := checkStringInSlice(v, ackList)
		if check {
			successList = append(successList, v)
		} else {
			failureList = append(failureList, v)
		}
	}
	return Acknowledgement{m.Key, m.ID, successList, failureList}, nil
}


func GetSMembers(descriptor string) []string{
	client := GetClient()
	sMem := client.SMembers(descriptor)
	return sMem.Val()
}

func checkStringInSlice(s string, sList []string) bool{
	for _, v := range sList {
		if s == v {return true}
	}
	return false
}

func VerifyMessageAndNotify(message string) {
	ack, err := CheckAcknowledgment(message)
	if err != nil {
		log.Println("Error acknowleging event")
		return
	}
	if len(ack.Failure) > 0 {
		notifySlack(ack)
	}
}

func notifySlack(failureAck Acknowledgement) {
	url := os.Getenv("SLACK_WEBHOOK")
	messageString := fmt.Sprintf(`{
 		"channel": "#laurentia",
		"username": "Laurentia-Bot",
		"icon_emoji": ":ghost:",	
"attachments": [
		{
		"fallback": "There may be an error with you subscriber",
		"color": "#ff0000",
		"title": "PubSub Failure Has Been Detected",
		"fields": [
		{
			"title": "Event",
			"value": ">%v" ,
			"short": true
		},
		{
			"title": "ID:",
			"value": "> %v" ,
			"short": true
		},
		{
			"title": "Failures:",
			"value": "> %v" ,
			"short": false
		},
		{
			"title": "Successful:",
			"value": "> %v" ,
			"short": false
		}
		],
		"footer": "laurentia",
		}
	]
	}`, failureAck.Event, failureAck.ID, failureAck.Failure, failureAck.Successful)
	var jsonStr = []byte(messageString)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func formatPythonDict(message string) string {
	message = strings.Replace(message, `None`, `null`, -1)
	message = strings.Replace(message, `True`, `true`, -1)
	message = strings.Replace(message, `False`, `false`, -1)
	message = strings.Replace(message, `'`, `"`, -1)
	return message
}