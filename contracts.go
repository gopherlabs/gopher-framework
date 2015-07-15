package framework

import (
	"net/http"
)

type HandlerFn func(rw http.ResponseWriter, req *http.Request)

type Providerable interface {
	Register(config map[string]interface{}) interface{}
	GetKey() string
}

type Servable interface {
	http.Handler
	Serve()
}

type Loggable interface {
	Info(msg string, args ...interface{})
	Debug(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Fatal(msg string, args ...interface{})
	Panic(msg string, args ...interface{})
}

type Renderable interface {
	Data(rw http.ResponseWriter, data []byte, status ...int)
	Text(rw http.ResponseWriter, data string, status ...int)
	JSON(rw http.ResponseWriter, data interface{}, status ...int)
	JSONP(rw http.ResponseWriter, callback string, data interface{}, status ...int)
	XML(rw http.ResponseWriter, data interface{}, status ...int)
	View(rw http.ResponseWriter, name string, binding interface{}, status ...int)
}

type Routable interface {
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
	Vars(r *http.Request) map[string]string
	Var(r *http.Request, param string) string
	Use(fn MiddlewareHandler, args ...interface{})
}

type Mappable interface {
	Get(key string) interface{}
	Has(key string) bool
	Set(key string, value interface{})
	Remove(key string)
}
