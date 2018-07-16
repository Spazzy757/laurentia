package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"../messages"
	"fmt"
)

func GetMessagesHandler(c *gin.Context) {
	messageList, err := messages.GetMessageList(2, 0)
	if err != nil {panic(err)}
	fmt.Print(messageList)
	c.String(http.StatusOK, "bar")
}