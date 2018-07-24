package handlers

import (
	`github.com/gin-gonic/gin`
)


func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET(`/health`, GetHealth)
	v1 := r.Group(`v1`)
	{
		v1.GET(`/subscribers`, GetSubscriberList)
		v1.GET(`/messages`, GetMessagesHandler)
		v1.GET(`/acknowledged`, GetAcknowledgedSubscribers)
	}
	ws := r.Group(`ws`)
	{
		ws.GET(`/messages`, func(c *gin.Context) {
			MessageWSHandler(c.Writer, c.Request)
		})
	}
	return r
}


