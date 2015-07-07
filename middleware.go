package framework

import "net/http"

type handlerFn func(rw http.ResponseWriter, req *http.Request)

type MiddlewareHandler func(rw http.ResponseWriter, req *http.Request, next func(), args ...interface{})

type Middleware struct {
	handler MiddlewareHandler
	args    []interface{}
}

func routeMiddleware(c *Container, rw http.ResponseWriter, req *http.Request, fn handlerFn) {
	fn(rw, req)
}

func routeLoggerMiddleware(rw http.ResponseWriter, req *http.Request, next func(), args ...interface{}) {
	c.providers[LOGGER].(Loggable).Info("[%s] %s", req.Method, req.URL.Path)
	next()
}
