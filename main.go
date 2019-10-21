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

	r.Handle("GET", "/proxyGET", func(c *gin.Context) {
		fmt.Println("GET request detected")
	})

	log.Fatal(r.Run(":" + strconv.Itoa(port)))
}
