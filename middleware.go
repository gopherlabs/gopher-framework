package framework

import "net/http"

type handlerFn func(rw http.ResponseWriter, req *http.Request)

type MiddlewareHandler func(rw http.ResponseWriter, req *http.Request, next func(), args ...interface{})

type Middleware struct {
	handler MiddlewareHandler
	args    []interface{}
}

func processMiddlewares(middlewares []Middleware, rw http.ResponseWriter, req *http.Request, fn handlerFn) {
	for _, middleware := range middlewares {
		next := false
		middleware.handler(rw, req, func() { next = true }, middleware.args...)
		if next {
			continue
		} else {
			return
		}
	}
	fn(rw, req)
}

func LoggerMiddleware(rw http.ResponseWriter, req *http.Request, next func(), args ...interface{}) {
	c.providers[LOGGER].(Loggable).Info("[%s] %s", req.Method, req.URL.Path)
	next()
}
