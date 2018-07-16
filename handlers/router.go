package handlers

import "github.com/gin-gonic/gin"

func GetMainEngine() *gin.Engine {
	r := gin.Default()
	r.GET("/messages", GetMessagesHandler)
	//r.GET("/ws", func(c *gin.Context) {
	//	WSHandler(c.Writer, c.Request, dataChannel, mongoChannel)
	//})
	return r
}