package framework

import "net/http"

type MiddlewareHandler func(rw http.ResponseWriter, req *http.Request, next func(), args ...interface{})

type Middleware struct {
	handler MiddlewareHandler
	args    []interface{}
}

func addRouteMiddlewares(handlers ...MiddlewareHandler) []Middleware {
	routeMiddlewares := []Middleware{}
	for _, handler := range handlers {
		middleware := Middleware{handler: handler}
		routeMiddlewares = append(routeMiddlewares, middleware)
	}
	return routeMiddlewares
}

func processMiddlewares(mw []Middleware, rw http.ResponseWriter, req *http.Request, fn HandlerFn, extraMw ...MiddlewareHandler) {
	middlewares := []Middleware{}
	middlewares = append(mw, addRouteMiddlewares(extraMw...)...)
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
