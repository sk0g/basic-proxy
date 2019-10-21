package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

const port = 8899

func main() {
	r := gin.Default()

	r.GET("/proxyGET", func(c *gin.Context) {
		fmt.Println("GET request detected")
	})

	r.POST("/proxyPOST", func(c *gin.Context) {
		fmt.Println("POST request detected")
	})

	log.Fatal(r.Run(":" + strconv.Itoa(port)))
}

func verifyContextHasRequiredValues(c *gin.Context) (errorMessage string, isOk bool) {
	urlProxyingTo := c.GetHeader("proxy_url")

	if urlProxyingTo == "" {
		return "No URL provided", false
	}

	return "", true
}
