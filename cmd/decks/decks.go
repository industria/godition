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

/*

curl -v -XPOST "http://localhost:8080/burn" -H "Content-Type: application/json" -d '{
    "data": {
        "product": {
            "id": "5d5a8cc11574df09cd20ecac",
            "name": "ekstrabladet-dk",
            "productType": "frontpage"
        },
        "edition": {
            "id": "5d5a8cf857cd2009c74b6378",
            "name": "manuel-top"
        },
        "burned": {
            "htmlHash": "fab7859c399e7721",
            "htmlUpdatedAt": "2022-02-20T06:10:07.076Z",
            "htmlUrl": "https://smooth-storage.aptoma.no/users/drf-eb/files/sphynx/2022/2/20/fab7859c399e7721.html",
            "cssUpdatedAt": "2022-02-14T09:03:41.489Z",
            "clientId": "eb",
            "editionId": "5d5a8cf857cd2009c74b6378",
            "previewUrl": "https://smooth-storage.aptoma.no/users/drf-eb/files/sphynx/2022/2/20/fab7859c399e7721-preview.html",
            "cssUrl": "https://smooth-storage.aptoma.no/users/drf-eb/files/sphynx/2022/2/14/236a375ec45e1904.css",
            "cssHash": "236a375ec45e1904"
        }
    },
    "event": "SphynxPostPublishEvent"
}'
*/