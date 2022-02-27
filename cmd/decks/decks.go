package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/industria/godition/dredition"
	"github.com/industria/godition/splitter"
)

var drEditionClient *dredition.Client

func main() {
	log.Print("Fisk")

	drEditionClient = dredition.NewClient()

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
	log.Printf("Notification: %v", notification)

	metadata, err := drEditionClient.Metadata(notification)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	log.Printf("Burn Metadata: %v", metadata)

	css, err := drEditionClient.CSS(metadata)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	log.Printf("CSS: %s", *css)

	// TODO: CSS source mapping if it is actually published...

	html, err := drEditionClient.HTML(metadata)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	log.Printf("HTML: %s", *html)

	htmlR := strings.NewReader(*html)
	decks, err := splitter.Split(htmlR, notification)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	for _, deck := range *decks {
		log.Printf("Deck: %v", deck)
	}

	c.Status(200)
}
