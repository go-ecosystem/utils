package http

import (
	"net/http"
	"testing"
)

func TestClientIP(t *testing.T) {
	want := ""
	req, _ := http.NewRequest("GET", "http://api.makeoptim.com", nil)

	// ""
	got := ClientIP(req)
	if got != want {
		t.Errorf("ClientIP() = %v, want %v", got, want)
	}

	// X-Envoy-External-Address
	want = "192.168.1.1"
	req.Header.Set("X-Envoy-External-Address", want)
	got = ClientIP(req)
	if got != want {
		t.Errorf("ClientIP() = %v, want %v", got, want)
	}

	// X-Envoy-External-Address and X-Forwarded-For
	req.Header.Set("X-Forwarded-For", "19.122.232.1")
	req.Header.Set("X-Envoy-External-Address", want)
	got = ClientIP(req)
	if got != want {
		t.Errorf("ClientIP() = %v, want %v", got, want)
	}
}
