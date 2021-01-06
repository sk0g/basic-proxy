package main

import (
	"bytes"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func verifyContextHasRequiredValues(c *gin.Context) (errorMessage string, isOk bool) {
	urlProxyingTo := c.GetHeader("proxy_url")

	if urlProxyingTo == "" {
		return "No URL provided", false
	}

	return "", true
}

func getRemoteURLAndRemoveFromHeaders(c *gin.Context) string {
	url := c.GetHeader("proxy_url")

	// no need to pass on the proxy_url header
	c.Header("proxy_url", "")

	return url
}

func getInsecureSkipVerifyAndRemoveFromHeaders(c *gin.Context) bool {
	skipVerifyCheck, err := strconv.ParseBool(c.GetHeader("cert_verify"))
	// fallback to verifying certs if header value can not be determined
	if err != nil {
		skipVerifyCheck = false
	}

	// no need to pass on the skipVerifyCheck header
	c.Header("cert_verify", "")

	return skipVerifyCheck
}


func extractHeadersFrom(headers http.Header) map[string]string {
	processedHeaders := make(map[string]string)

	for name, val := range headers {
		// http.Header is in the form of map[string][]string, want map[string]string
		processedHeaders[name] = strings.Join(val, ", ")
	}

	return processedHeaders
}

func readcloserToString(i io.ReadCloser) (string, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(i)
	return buf.String(), err
}
