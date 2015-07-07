package framework

import (
	"fmt"
	"net/http"
)

// Sample facade
/*
type SampleFacade struct {
	name     string
	provider Samplable
}

func (c Container) NewSample() Samplable {
	return new(SampleFacade).NewSample()
}

func (p *SampleFacade) Register(config map[string]interface{}) interface{} {
	return p
}

func (p *SampleFacade) GetKey() string {
	return p.provider.GetKey()
}

func (p *SampleFacade) NewSample() Samplable {
	p = new(SampleFacade)
	p.provider = c.providers[SAMPLE].(Samplable).NewSample()
	return p
}

func (p *SampleFacade) GetName() string {
	name := p.provider.GetName()
	return name
}

func (p *SampleFacade) SetName(name string) {
	p.provider.SetName("facade added: " + name)
}
*/

// Router Facade
type RouteFacade struct {
	http.Handler
	provider    Routable
	middlewares []func(rw http.ResponseWriter, req *http.Request, next func())
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
	r.provider = c.providers[ROUTER].(Routable).NewRouter()
	return r
}

func (r *RouteFacade) SubRouter() Routable {
	sub := new(RouteFacade)
	sub.provider = c.providers[ROUTER].(Routable).SubRouter()
	return sub
}

func (r *RouteFacade) Get(path string, fn func(http.ResponseWriter, *http.Request)) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		for _, middleware := range r.middlewares {
			next := false
			middleware(rw, req, func() { next = true })
			fmt.Printf("The value of next is %v", next)
			if next {
				continue
			} else {
				break
			}
		}
		routeMiddleware(c, rw, req, fn)
	}
	r.provider.Get(path, nfn)
}

func (r *RouteFacade) Head(path string, fn func(http.ResponseWriter, *http.Request)) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		routeMiddleware(c, rw, req, fn)
	}
	r.provider.Head(path, nfn)
}

func (r *RouteFacade) Post(path string, fn func(http.ResponseWriter, *http.Request)) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		routeMiddleware(c, rw, req, fn)
	}
	r.provider.Post(path, nfn)
}

func (r *RouteFacade) Put(path string, fn func(http.ResponseWriter, *http.Request)) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		routeMiddleware(c, rw, req, fn)
	}
	r.provider.Put(path, nfn)
}

func (r *RouteFacade) Patch(path string, fn func(http.ResponseWriter, *http.Request)) {
	r.provider.Patch(path, fn)
}

func (r *RouteFacade) Delete(path string, fn func(http.ResponseWriter, *http.Request)) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		routeMiddleware(c, rw, req, fn)
	}
	r.provider.Delete(path, nfn)
}

func (r *RouteFacade) Options(path string, fn func(http.ResponseWriter, *http.Request)) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		routeMiddleware(c, rw, req, fn)
	}
	r.provider.Options(path, nfn)
}

func (r *RouteFacade) Match(path string, fn func(http.ResponseWriter, *http.Request), verbs ...string) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		routeMiddleware(c, rw, req, fn)
	}
	r.provider.Match(path, nfn, verbs...)
}

func (r *RouteFacade) All(path string, fn func(http.ResponseWriter, *http.Request)) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		routeMiddleware(c, rw, req, fn)
	}
	r.provider.All(path, nfn)
}

func (r *RouteFacade) NotFound(fn func(http.ResponseWriter, *http.Request)) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		routeMiddleware(c, rw, req, fn)
	}
	r.provider.NotFound(nfn)
}

func (r *RouteFacade) Serve() {
	c.showBanner()
	r.provider.Serve()
}

// Middleware
func (r *RouteFacade) Use(mw func(rw http.ResponseWriter, req *http.Request, next func())) {
	r.middlewares = append(r.middlewares, mw)
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

/*

// Parameters
func (c appContainer) PathParams(req *http.Request) map[string]string {
	return c.providers[PARAMS].(Parametable).PathParams(req)
}

// Parameters
func (c appContainer) PathParam(req *http.Request, param string) string {
	return c.providers[PARAMS].(Parametable).PathParam(req, param)
}

// Logger
func (c appContainer) Log() Loggable {
	return c.providers[LOGGER].(Loggable)
}

// Renderer
func (c appContainer) View(rw http.ResponseWriter, status int, name string, binding interface{}) {
	c.providers[RENDERER].(Renderable).View(rw, status, name, binding)
}
*/
