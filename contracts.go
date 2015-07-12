package framework

import (
	"net/http"
)

type HandlerFn func(rw http.ResponseWriter, req *http.Request)

type Providerable interface {
	Register(config map[string]interface{}) interface{}
	GetKey() string
}

type Loggable interface {
	Providerable
	Info(msg string, args ...interface{})
	Debug(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Fatal(msg string, args ...interface{})
	Panic(msg string, args ...interface{})
}

type Parametable interface {
	Providerable
	PathParams(r *http.Request) map[string]string
	PathParam(r *http.Request, param string) string
}

type Renderable interface {
	Providerable
	View(rw http.ResponseWriter, status int, name string, binding interface{})
}

type Routable interface {
	Providerable
	http.Handler
	NewRouter() Routable
	//SubRouter() Routable
	Get(path string, fn HandlerFn, mw ...MiddlewareHandler)
	Head(path string, fn HandlerFn, mw ...MiddlewareHandler)
	Post(path string, fn HandlerFn, mw ...MiddlewareHandler)
	Put(path string, fn HandlerFn, mw ...MiddlewareHandler)
	Patch(path string, fn HandlerFn, mw ...MiddlewareHandler)
	Delete(path string, fn HandlerFn, mw ...MiddlewareHandler)
	Options(path string, fn HandlerFn, mw ...MiddlewareHandler)
	Match(path string, fn HandlerFn, verbs []string, mw ...MiddlewareHandler)
	All(path string, fn HandlerFn, mw ...MiddlewareHandler)
	NotFound(fn HandlerFn, mw ...MiddlewareHandler)
	Serve()
	Use(fn MiddlewareHandler, args ...interface{})
}

type Mappable interface {
	Providerable
	Get(key string) interface{}
	Has(key string) bool
	Set(key string, value interface{})
	Remove(key string)
}
