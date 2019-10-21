package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

const port = 8899

func main() {
	r := gin.Default()

	log.Fatal(r.Run(":" + strconv.Itoa(port)))
}
