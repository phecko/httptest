package qin

import (
	"testing"
)

func TestRouter(t *testing.T) {

	router := NewRouter()

	passHello := false

	router.Get("/hello", func(c *Context) {
		passHello = true
	})

	passWorld := false
	router.Post("/world", func(c *Context) {
		passWorld = true
	})

	if len(router.MethodMap) != 2 {
		t.Fatal("router count wrong")
	}

}
