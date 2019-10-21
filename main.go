package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

const port = 8899

func main() {
	r := gin.Default()

	r.GET("/proxyGET", handleGetRequest)
	r.POST("/proxyPOST", handlePostRequest)

	log.Fatal(r.Run(":" + strconv.Itoa(port)))
}

func handleGetRequest(c *gin.Context) {
	if err, isOk := verifyContextHasRequiredValues(c); !isOk {
		c.JSON(
			400,
			map[string]interface{}{
				"error": err,
			})
		return
	}
}

func handlePostRequest(c *gin.Context) {
	if err, isOk := verifyContextHasRequiredValues(c); !isOk {
		c.JSON(
			400,
			map[string]interface{}{
				"error": err,
			})
		return
	}
}
