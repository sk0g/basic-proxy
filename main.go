package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	_ "github.com/joho/godotenv"
	"log"
	"strconv"
)

const port = 8899

func main() {
	r := gin.Default()

	r.GET("/proxy", handleGetRequest)
	r.POST("/proxy", handlePostRequest)

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
	headers := extractHeadersFrom(c.Request.Header)

	var responseData interface{}
	restyClient := resty.New()

	resp, _ := restyClient.
		R().
		SetResult(&responseData).
		SetHeaders(headers).
		Get(url)

	if responseData == nil {
		c.String(
			resp.StatusCode(),
			string(resp.Body()))
	} else {
		c.JSON(
			resp.StatusCode(),
			responseData)
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

	url := getRemoteURLAndRemoveFromHeaders(c)
	headers := extractHeadersFrom(c.Request.Header)

	var responseData interface{}
	restyClient := resty.New()

	var body interface{}
	if err := json.NewDecoder(c.Request.Body).Decode(&body); err != nil {
		log.Println("Error decoding JSON", err)
	}

	resp, _ := restyClient.
		R().
		SetResult(&responseData).
		SetHeaders(headers).
		SetBody(body).
		Post(url)

	if responseData == nil {
		c.String(
			resp.StatusCode(),
			string(resp.Body()))
	} else {
		c.JSON(
			resp.StatusCode(),
			responseData)
	}
}
