package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv"
)

var port = 8899

func init() {
	if envPort := os.Getenv("port"); len(envPort) >= 3 {
		if portNum, err := strconv.Atoi(envPort); err == nil {
			port = portNum
		} else {
			log.Printf("Port number is not a number, instead is '%v'. Defaulted to 8899\n", envPort)
		}
	} else {
		log.Printf("Port number has to be at least 3 digits and a number, is '%v'. Defaulted to 8899\n", envPort)
	}

	log.SetFlags(log.Lshortfile)
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowMethods:    []string{"POST", "GET", "OPTIONS"},
		AllowAllOrigins: true,
		AllowHeaders: []string{
			"Accept",
			"Access-Control-Allow-Headers",
			"Access-Control-Allow-Origin",
			"Access-Control-Request-Headers",
			"Access-Control-Request-Method",
			"Authorization",
			"Content-Type",
			"Origin",
			"X-Requested-With",
			"proxy_url",
		},
		AllowCredentials:       true,
		MaxAge:                 3600,
		AllowWildcard:          true,
		AllowBrowserExtensions: true,
		AllowWebSockets:        true,
		AllowFiles:             true,
	}))

	r.GET("/proxy", handleGetRequest)
	r.POST("/proxy", handlePostRequest)
	r.POST("/proxy_xml", handlePostXmlRequest)

	log.Fatal(r.Run(":" + strconv.Itoa(port)))
}

func handleGetRequest(c *gin.Context) {
	if err, isOk := verifyContextHasRequiredValues(c); !isOk {
		log.Println(err)
		c.JSON(
			400,
			map[string]interface{}{
				"error": err,
			})
		return
	}

	url := getRemoteURLAndRemoveFromHeaders(c)

	restyClient := restyClientInit(c)
	var responseData interface{}
	resp, err := restyClient.
		R().
		SetResult(&responseData).
		Get(url)

	if err != nil {
		log.Println(err)
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

	var body interface{}
	requestBody, _ := readcloserToString(c.Request.Body)
	if err := json.Unmarshal([]byte(requestBody), &body); err != nil {
		log.Println(err)
		// fall back to using just the string, if json Unmarshaling fails
		body = requestBody
	}

	restyClient := restyClientInit(c)
	var responseData interface{}
	resp, err := restyClient.
		R().
		SetResult(&responseData).
		SetBody(body).
		Post(url)

	if err != nil {
		log.Println(err)
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

func handlePostXmlRequest(c *gin.Context) {
	if err, isOk := verifyContextHasRequiredValues(c); !isOk {
		log.Println(err)
		c.JSON(
			400,
			map[string]interface{}{
				"error": err,
			})
		return
	}

	url := getRemoteURLAndRemoveFromHeaders(c)

	var body interface{}
	requestBody, _ := readcloserToString(c.Request.Body)
	if err := json.Unmarshal([]byte(requestBody), &body); err != nil {
		// fall back to using just the string, if json Unmarshalling fails
		body = requestBody
	}

	restyClient := restyClientInit(c)
	var responseData interface{}
	resp, err := restyClient.
		R().
		SetBody(requestBody).
		SetResult(&responseData).
		Post(url)

	if err != nil {
		log.Println(err)
		c.JSON(
			400,
			map[string]interface{}{
				"error": err.Error(),
			})
		return
	}

	// response headers
	for k, v := range resp.Header() {
		c.Header(k, strings.Join(v, "; "))
	}

	// response body
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
