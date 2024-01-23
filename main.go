package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thematrix97/gografanaspeaker/controllers"
	"github.com/thematrix97/gografanaspeaker/services"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/event", func(ctx *gin.Context) {
		res, errors := controllers.ProcessGrafanaEvent(ctx)
		if errors != nil {
			ctx.String(http.StatusInternalServerError, "Server Error")
		} else {
			ctx.JSON(http.StatusOK, res)
		}
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

// https://stackoverflow.com/questions/69948784/how-to-handle-errors-in-gin-middleware
func ErrorHandler(c *gin.Context) {
	c.Next()

	for _, err := range c.Errors {
		log.Println(err.Error())
	}

	c.JSON(http.StatusInternalServerError, "Internal Error, Whoops") //TODO implement a proper error handling
}

func main() {
	services.LoadEnvConfig()
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
