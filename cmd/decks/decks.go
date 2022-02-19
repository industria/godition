package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/industria/godition/dredition"
)

func main() {
	log.Print("Fisk")

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.SetTrustedProxies(nil)

	r.GET("/AWS.ALB/healthcheck", healthcheck)
	r.POST("/burn", burn)
	r.Run(":8080")
}

func healthcheck(c *gin.Context) {
	c.Status(200)
}

func burn(c *gin.Context) {
	var notification dredition.Notification

	err := c.BindJSON(&notification)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	log.Printf("%v", notification)
	c.Status(200)
}
