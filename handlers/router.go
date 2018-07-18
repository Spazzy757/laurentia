package handlers

import (
	"github.com/gin-gonic/gin"
)


func SetupRouter() *gin.Engine {
	r := gin.Default()
	//r.GET("/ws", func(c *gin.Context) {
	//	WSHandler(c.Writer, c.Request, dataChannel, mongoChannel)
	//})
	r.GET("/health", GetHealth)
	v1 := r.Group("v1")
	{
		v1.GET("/subscribers", GetSubscriberList)
		v1.GET("/messages", GetMessagesHandler)
		v1.GET("/acknowledged", GetAcknowledgedSubscribers)
	}
	return r
}


