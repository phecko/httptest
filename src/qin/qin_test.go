package qin

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func PerformRequest(method, path, body string) {
	req, _ := http.NewRequest(method, path, nil)
	rsp := httptest.NewRecorder()
	engine := NewEngine()
	engine.ServeHTTP(rsp, req)
}

func TestUrl(t *testing.T) {
	PerformRequest("GET", "/hello", "")
	PerformRequest("GET", "/gogo", "")
}
