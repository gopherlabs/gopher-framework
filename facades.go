package framework

import (
	"net/http"
)

// Router Facade
type RouteFacade struct {
	http.Handler
	provider    Routable
	middlewares []Middleware
}

func (r *RouteFacade) Register(config map[string]interface{}) interface{} {
	r = new(RouteFacade)
	r.provider = c.providers[ROUTER].(Providerable).Register(config).(Routable)
	return r
}

func (r *RouteFacade) GetKey() string {
	return r.provider.(Providerable).GetKey()
}

/*
func (r *RouteFacade) SubRouter() Routable {
	sub := new(RouteFacade)
	sub.middlewares = r.middlewares
	sub.provider = c.providers[ROUTER].(Routable).SubRouter()
	return sub
}
*/

func (r *RouteFacade) Get(path string, fn HandlerFn, mw ...MiddlewareHandler) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		processMiddlewares(append(c.middlewares, r.middlewares...), rw, req, fn, mw...)
	}
	r.provider.Get(path, nfn)
}

func (r *RouteFacade) Head(path string, fn HandlerFn, mw ...MiddlewareHandler) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		processMiddlewares(append(c.middlewares, r.middlewares...), rw, req, fn, mw...)
	}
	r.provider.Head(path, nfn)
}

func (r *RouteFacade) Post(path string, fn HandlerFn, mw ...MiddlewareHandler) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		processMiddlewares(append(c.middlewares, r.middlewares...), rw, req, fn, mw...)
	}
	r.provider.Post(path, nfn)
}

func (r *RouteFacade) Put(path string, fn HandlerFn, mw ...MiddlewareHandler) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		processMiddlewares(append(c.middlewares, r.middlewares...), rw, req, fn, mw...)
	}
	r.provider.Put(path, nfn)
}

func (r *RouteFacade) Patch(path string, fn HandlerFn, mw ...MiddlewareHandler) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		processMiddlewares(append(c.middlewares, r.middlewares...), rw, req, fn, mw...)
	}
	r.provider.Patch(path, nfn)
}

func (r *RouteFacade) Delete(path string, fn HandlerFn, mw ...MiddlewareHandler) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		processMiddlewares(append(c.middlewares, r.middlewares...), rw, req, fn, mw...)
	}
	r.provider.Delete(path, nfn)
}

func (r *RouteFacade) Options(path string, fn HandlerFn, mw ...MiddlewareHandler) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		processMiddlewares(append(c.middlewares, r.middlewares...), rw, req, fn, mw...)
	}
	r.provider.Options(path, nfn)
}

func (r *RouteFacade) Match(path string, fn HandlerFn, verbs []string, mw ...MiddlewareHandler) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		processMiddlewares(append(c.middlewares, r.middlewares...), rw, req, fn, mw...)
	}
	r.provider.Match(path, nfn, verbs)
}

func (r *RouteFacade) All(path string, fn HandlerFn, mw ...MiddlewareHandler) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		processMiddlewares(append(c.middlewares, r.middlewares...), rw, req, fn, mw...)
	}
	r.provider.All(path, nfn)
}

func (r *RouteFacade) NotFound(fn HandlerFn, mw ...MiddlewareHandler) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		processMiddlewares(append(c.middlewares, r.middlewares...), rw, req, fn, mw...)
	}
	r.provider.NotFound(nfn)
}

func (r *RouteFacade) Vars(req *http.Request) map[string]string {
	return r.provider.(Routable).Vars(req)
}

func (r *RouteFacade) Var(req *http.Request, name string) string {
	return r.provider.(Routable).Var(req, name)
}

func (r *RouteFacade) Serve() {
	c.showBanner()
	r.provider.(Servable).Serve()
}

// Middleware
func (r *RouteFacade) Use(mw MiddlewareHandler, args ...interface{}) {
	middleware := Middleware{handler: mw, args: args}
	r.middlewares = append(r.middlewares, middleware)
}
