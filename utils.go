package main

import "github.com/gin-gonic/gin"

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
