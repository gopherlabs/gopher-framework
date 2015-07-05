package framework

import "net/http"

type handlerFn func(rw http.ResponseWriter, req *http.Request)

func routeMiddleware(c Container, rw http.ResponseWriter, req *http.Request, fn handlerFn) {
	c.providers[LOGGER].(Loggable).Info("[%s] %s", req.Method, req.URL.Path)
	fn(rw, req)
}
