package http

import "net/http"

var clientIPHeaders = []string{"X-Envoy-External-Address", "X-Forwarded-For", "X-Real-Ip", "X-Appengine-Remote-Addr"}

// ClientIP get client IP with possible headers.
func ClientIP(request *http.Request) string {
	for _, v := range clientIPHeaders {
		ip := request.Header.Get(v)
		if ip != "" {
			return ip
		}
	}
	return ""
}
