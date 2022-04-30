package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//Cors global CORS middleware
func Cors(headers ...string) gin.HandlerFunc {
	defaultHeaders := []string{"utoken", "lang", "Content-Type", "Accept", "Origin", "Access-Control-Allow-Origin", "Cache-Control"}
	for _, h := range headers {
		defaultHeaders = append(defaultHeaders, h)
	}
	headerStr := strings.Join(defaultHeaders, ",")
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Add("Access-Control-Allow-Headers", headerStr)
		c.Writer.Header().Add("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,DELETE,PATCH")
		c.Writer.Header().Add("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.Status(http.StatusOK)
			c.Abort()
			return
		}
		c.Next()
	}
}
