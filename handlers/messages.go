package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/Spazzy757/laurentia/messages"
	"strconv"
	"fmt"
)

func GetMessagesHandler(c *gin.Context) {
	limit, err := strconv.Atoi(c.Request.URL.Query().Get("limit"))
	if err != nil {limit = 10}
	page, err := strconv.Atoi(c.Request.URL.Query().Get("page"))
	if err != nil {page = 0}
	messageList, _ := messages.GetMessageList(limit, page)
	messageJson := make([]string, 0)
	for i := 0; i < len(messageList); i++ {
		fmt.Printf("%T", messageList[i].Message)
		messageJson = append(messageJson, messageList[i].Message)
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