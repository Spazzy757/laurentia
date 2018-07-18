package handlers

import (
	"github.com/gin-gonic/gin"
)


func SetupRouter() *gin.Engine {
	r := gin.Default()
	//r.GET("/ws", func(c *gin.Context) {
	//	WSHandler(c.Writer, c.Request, dataChannel, mongoChannel)
	//})
	v1 := r.Group("v1")
	{
		v1.GET("/subscribers", GetSubScriberList)
		v1.GET("/messages", GetMessagesHandler)
	}
	return r
}


