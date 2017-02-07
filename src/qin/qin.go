package qin

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type Engine struct {
	Router  *Router
	CtxPool sync.Pool
}

type Context struct {
	Engine  *Engine
	Request *http.Request
}

func NewEngine() *Engine {
	engine := &Engine{}
	engine.CtxPool.New = func() interface{} {
		c := &Context{}
		c.Engine = engine
		return c
	}
	engine.Router = NewRouter()
	return engine
}

func (engine *Engine) ServeHTTP(rspWriter http.ResponseWriter, req *http.Request) {
	fmt.Println("Get Requset : " + req.URL.String())
}

func (engine *Engine) Run(addr, port string) {
	log.Fatal(http.ListenAndServe(addr+":"+port, engine))
}
