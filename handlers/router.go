package handlers

import (
	`github.com/gin-gonic/gin`
	"github.com/appleboy/gin-jwt"
	"time"
	"os"
)


func SetupRouter() *gin.Engine {
	authMiddleware := ConfigureAuthMiddleware()

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET(`/health`, GetHealth)
	r.POST("/login", authMiddleware.LoginHandler)

	auth := r.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	}
	v1 := r.Group(`v1`)
	v1.Use(authMiddleware.MiddlewareFunc())
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

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"userID": claims["id"],
		"text":   "Hello World.",
	})
}

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}


func ConfigureAuthMiddleware() *jwt.GinJWTMiddleware{
	superUserName := os.Getenv(`USERNAME`)
	superUserPassword := os.Getenv(`PASSWORD`)
	key := os.Getenv(`KEY`)
	return &jwt.GinJWTMiddleware{
		Realm:      "test zone",
		Key:        []byte(key),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(userId string, password string, c *gin.Context) (interface{}, bool) {
			if userId == superUserName && password == superUserPassword {
				return &User{
					UserName:  userId,
					LastName:  "admin",
					FirstName: "admin",
				}, true
			}

			return nil, false
		},
		Authorizator: func(user interface{}, c *gin.Context) bool {
			if v, ok := user.(string); ok && v == "admin" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	}
}