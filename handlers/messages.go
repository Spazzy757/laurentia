package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/Spazzy757/laurentia/messages"
	"strconv"
	"strings"
	"log"
	"encoding/json"
)

type JSONString string

type DynamicMessage struct {
	Key string `json:"key"`
	ID  string `json:"id"`
	Payload interface{} `json:"payload"`
}

func GetMessagesHandler(c *gin.Context) {
	limit, err := strconv.Atoi(c.Request.URL.Query().Get("limit"))
	if err != nil {limit = 10}
	page, err := strconv.Atoi(c.Request.URL.Query().Get("page"))
	if err != nil {page = 0}
	messageList, _ := messages.GetMessageList(limit, page)
	var messageJson []DynamicMessage
	for i := 0; i < len(messageList); i++ {
		message := formatPythonDict(messageList[i].Message)
		var m DynamicMessage
		if err := json.Unmarshal([]byte(message), &m); err != nil {
			log.Fatal(err)
		}
		messageJson = append(messageJson, m)
	}
	c.JSON(http.StatusOK, gin.H{
		"messages": messageJson,
	})
}

func GetSubscriberList(c *gin.Context) {
	event := c.Request.URL.Query().Get("event")
	lookUp := "pubsub.events." + event + ".subscribers"
	subscriberList := messages.GetSMembers(lookUp)
	c.JSON(http.StatusOK, gin.H{
		"subscribers": subscriberList,
	})
}

func GetAcknowledgedSubscribers(c *gin.Context)  {
	event := c.Request.URL.Query().Get("event")
	messageID := c.Query("messageID")
	lookUp := "pubsub.events.actions." + event + "." + messageID + ".received"
	acknowledgedList := messages.GetSMembers(lookUp)
	c.JSON(http.StatusOK, gin.H{
		"event": event,
		"eventID": messageID,
		"acknowledged": acknowledgedList,
	})
}

func formatPythonDict(message string) string {
	message = strings.Replace(message, `None`, `null`, -1)
	message = strings.Replace(message, `True`, `true`, -1)
	message = strings.Replace(message, `False`, `false`, -1)
	message = strings.Replace(message, `'`, `"`, -1)
	return message
}