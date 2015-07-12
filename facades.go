package framework

import "net/http"

// Router Facade
type RouteFacade struct {
	http.Handler
	provider    Routable
	middlewares []Middleware
}

func (c Container) NewRouter() Routable {
	return new(RouteFacade).NewRouter()
}

func (r *RouteFacade) Register(config map[string]interface{}) interface{} {
	return r
}

func (r *RouteFacade) GetKey() string {
	return r.provider.GetKey()
}

func (r *RouteFacade) NewRouter() Routable {
	r = new(RouteFacade)
	r.middlewares = c.middlewares
	r.provider = c.providers[ROUTER].(Routable).NewRouter()
	return r
}

func (r *RouteFacade) SubRouter() Routable {
	sub := new(RouteFacade)
	sub.middlewares = r.middlewares
	sub.provider = c.providers[ROUTER].(Routable).SubRouter()
	return sub
}

func (r *RouteFacade) Get(path string, fn HandlerFn, mw ...MiddlewareHandler) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		processMiddlewares(r.middlewares, rw, req, fn, mw...)
	}
	r.provider.Get(path, nfn)
}

func (r *RouteFacade) Head(path string, fn HandlerFn, mw ...MiddlewareHandler) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		processMiddlewares(r.middlewares, rw, req, fn, mw...)
	}
	r.provider.Head(path, nfn)
}

func (r *RouteFacade) Post(path string, fn HandlerFn, mw ...MiddlewareHandler) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		processMiddlewares(r.middlewares, rw, req, fn, mw...)
	}
	r.provider.Post(path, nfn)
}

func (r *RouteFacade) Put(path string, fn HandlerFn, mw ...MiddlewareHandler) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		processMiddlewares(r.middlewares, rw, req, fn, mw...)
	}
	r.provider.Put(path, nfn)
}

func (r *RouteFacade) Patch(path string, fn HandlerFn, mw ...MiddlewareHandler) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		processMiddlewares(r.middlewares, rw, req, fn, mw...)
	}
	r.provider.Patch(path, nfn)
}

func (r *RouteFacade) Delete(path string, fn HandlerFn, mw ...MiddlewareHandler) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		processMiddlewares(r.middlewares, rw, req, fn, mw...)
	}
	r.provider.Delete(path, nfn)
}

func (r *RouteFacade) Options(path string, fn HandlerFn, mw ...MiddlewareHandler) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		processMiddlewares(r.middlewares, rw, req, fn, mw...)
	}
	r.provider.Options(path, nfn)
}

func (r *RouteFacade) Match(path string, fn HandlerFn, verbs []string, mw ...MiddlewareHandler) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		processMiddlewares(r.middlewares, rw, req, fn, mw...)
	}
	r.provider.Match(path, nfn, verbs)
}

func (r *RouteFacade) All(path string, fn HandlerFn, mw ...MiddlewareHandler) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		processMiddlewares(r.middlewares, rw, req, fn, mw...)
	}
	r.provider.All(path, nfn)
}

func (r *RouteFacade) NotFound(fn HandlerFn, mw ...MiddlewareHandler) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		processMiddlewares(r.middlewares, rw, req, fn, mw...)
	}
	r.provider.NotFound(nfn)
}

func (r *RouteFacade) Serve() {
	c.showBanner()
	r.provider.Serve()
}

// Middleware
func (r *RouteFacade) Use(mw MiddlewareHandler, args ...interface{}) {
	middleware := Middleware{handler: mw, args: args}
	r.middlewares = append(r.middlewares, middleware)
}

// Logger
func (c Container) NewLog() Loggable {
	return c.providers[LOGGER].(Loggable).NewLog()
}

// Parameters
func (c Container) PathParams(req *http.Request) map[string]string {
	return c.providers[PARAMS].(Parametable).PathParams(req)
}

// Parameters
func (c Container) PathParam(req *http.Request, param string) string {
	return c.providers[PARAMS].(Parametable).PathParam(req, param)
}

// Renderer
func (c Container) View(rw http.ResponseWriter, status int, name string, binding interface{}) {
	c.providers[RENDERER].(Renderable).View(rw, status, name, binding)
}

// Context
func (c Container) Get(key string) (value interface{}) {
	return c.providers[MAPPER].(Mappable).Get(key)
}

func (c Container) Has(key string) bool {
	return c.providers[MAPPER].(Mappable).Has(key)
}

func (c Container) Set(key string, value interface{}) {
	c.providers[MAPPER].(Mappable).Set(key, value)
}

func (c Container) Remove(key string) {
	c.providers[MAPPER].(Mappable).Remove(key)
}
