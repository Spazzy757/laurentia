package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/Spazzy757/laurentia/messages"
	"encoding/json"
	"strconv"
)

func GetMessagesHandler(c *gin.Context) {
	limit, err := strconv.Atoi(c.Request.URL.Query().Get("limit"))
	if err != nil {limit = 10}
	page, err := strconv.Atoi(c.Request.URL.Query().Get("page"))
	if err != nil {page = 10}
	messageList, _ := messages.GetMessageList(limit, page)
	var messageJson []interface{}
	for i := 0; i < len(messageList); i++ {
		b := []byte(messageList[i].Message)
		var f interface{}
		err := json.Unmarshal(b, &f)
		if err != nil {panic(err)}
		messageJson = append(messageJson, f)
	}
	c.SecureJSON(http.StatusOK, messageJson)
}

func GetSubScriberList(c *gin.Context) {

}