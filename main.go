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
