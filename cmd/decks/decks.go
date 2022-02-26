package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/industria/godition/dredition"
)

var burnProcessor *dredition.BurnProcessor

func main() {
	log.Print("Fisk")

	burnProcessor = dredition.NewBurnProcessor()

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

	err = burnProcessor.Process(notification)
	if err != nil {
		log.Printf("unable to process burn notification %v : %v", notification, err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(200)
}
