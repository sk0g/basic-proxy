package main

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	_ "github.com/joho/godotenv"
)

const port = 8899

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowMethods:    []string{"POST", "GET", "OPTIONS"},
		AllowAllOrigins: true,
		AllowHeaders: []string{
			"proxy_url",
			"Access-Control-Allow-Headers",
			"Access-Control-Allow-Origin",
			"Origin",
			"Accept",
			"X-Requested-With",
			"Content-Type",
			"Access-Control-Request-Method",
			"Access-Control-Request-Headers"},
		AllowCredentials:       true,
		MaxAge:                 3600,
		AllowWildcard:          true,
		AllowBrowserExtensions: true,
		AllowWebSockets:        true,
		AllowFiles:             true,
	}))

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

	resp, err := restyClient.
		R().
		SetResult(&responseData).
		SetHeaders(headers).
		Get(url)

	if err != nil {
		log.Println("Error while proxying request: \n", err)
	}

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

	var body interface{}
	requestBody, _ := readcloserToString(c.Request.Body)
	if err := json.Unmarshal([]byte(requestBody), &body); err != nil {
		// fall back to using just the string, if json Unmarshalling fails
		body = requestBody
	}

	var responseData map[string]interface{}

	resp, err := resty.
		New().
		R().
		SetBody(requestBody).
		SetHeaders(headers).
		SetResult(&responseData).
		Post(url)

	if err != nil {
		c.JSON(
			400,
			map[string]interface{}{
				"error": err.Error(),
			})
		return
	}

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
