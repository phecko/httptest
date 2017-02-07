package qin

import (
	"errors"
)

type HandleFunc func(*Context)

var (
	METHOD_GET  string = "GET"
	METHOD_POST string = "GET"
)

type IRouter interface {
	Get(string, HandleFunc)
	Post(string, HandleFunc)
}

type Router struct {
	MethodMap map[string]map[string]*HandleFunc
}

func NewRouter() *Router {
	router := &Router{}
	router.MethodMap = make(map[string]map[string]*HandleFunc, 0)
	return router
}

func (router *Router) Get(path string, handle HandleFunc) error {
	return router.add(METHOD_GET, path, handle)
}

func (router *Router) Post(path string, handle HandleFunc) error {
	return router.add(METHOD_POST, path, handle)
}

func (router *Router) Match(method, path string) (bool, *HandleFunc) {
	return router.add(METHOD_POST, path, handle)
}

func (router *Router) add(method, path string, handle HandleFunc) error {
	if _, ok := router.MethodMap[method]; !ok {
		router.MethodMap[method] = make(map[string]*HandleFunc, 0)
	}

	if _, ok := router.MethodMap[method][path]; ok {
		return errors.New("Route has add :" + path)
	}

	router.MethodMap[method][path] = &handle
	return nil
}
