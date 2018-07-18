package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/Spazzy757/laurentia/messages"
	"strconv"
	"fmt"
	"log"
)

func GetMessagesHandler(c *gin.Context) {
	limit, err := strconv.Atoi(c.Request.URL.Query().Get("limit"))
	if err != nil {limit = 10}
	page, err := strconv.Atoi(c.Request.URL.Query().Get("page"))
	if err != nil {page = 10}
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

func GetSubScriberList(c *gin.Context) {
	event := c.Request.URL.Query().Get("event")
	log.Println(event)
	lookUp := "pubsub.events." + event + ".subscribers"
	log.Println(lookUp)
	subscriberList := messages.GetSMembers(lookUp)
	log.Println(subscriberList)
	c.JSON(http.StatusOK, gin.H{
		"subscribers": subscriberList,
	})
}