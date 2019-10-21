package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
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

	url := getRemoteURLAndRemoveFromHeaders(c)

	var responseData map[string]interface{}
	restyClient := resty.New()

	resp, _ := restyClient.
		R().
		SetResult(&responseData).
		//SetHeaders(c.Request.Header).
		Get(url)

	c.JSON(
		resp.StatusCode(),
		responseData)
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

	url := getRemoteURLAndRemoveFromHeaders(c)
	fmt.Println(url)
}
