package middleware

import (
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// Recovery Recovery
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				//write log
				log.Panicf("gin err: %v", err)

				// If the connection is dead, we can't write a status to it.
				if brokenPipe {
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
				} else {
					c.AbortWithStatus(http.StatusInternalServerError)
				}
			}
		}()
		c.Next()
	}
}
