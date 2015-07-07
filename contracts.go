package framework

import (
	"net/http"
)

type Providerable interface {
	Register(config map[string]interface{}) interface{}
	GetKey() string
}

type Loggable interface {
	Providerable
	NewLog() Loggable
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
	SubRouter() Routable
	Get(path string, fn func(http.ResponseWriter, *http.Request), mw ...MiddlewareHandler)
	Head(path string, fn func(http.ResponseWriter, *http.Request), mw ...MiddlewareHandler)
	Post(path string, fn func(http.ResponseWriter, *http.Request), mw ...MiddlewareHandler)
	Put(path string, fn func(http.ResponseWriter, *http.Request), mw ...MiddlewareHandler)
	Patch(path string, fn func(http.ResponseWriter, *http.Request), mw ...MiddlewareHandler)
	Delete(path string, fn func(http.ResponseWriter, *http.Request), mw ...MiddlewareHandler)
	Options(path string, fn func(http.ResponseWriter, *http.Request), mw ...MiddlewareHandler)
	Match(path string, fn func(http.ResponseWriter, *http.Request), verbs []string, mw ...MiddlewareHandler)
	All(path string, fn func(http.ResponseWriter, *http.Request), mw ...MiddlewareHandler)
	NotFound(fn func(http.ResponseWriter, *http.Request), mw ...MiddlewareHandler)
	Serve()
	Use(fn MiddlewareHandler, args ...interface{})
}

/*
type Samplable interface {
	Providerable
	NewSample() Samplable
	GetName() string
	SetName(name string)
}
*/
