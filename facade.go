package framework

import (
	"net/http"
)

// Sample facade
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
	//c.providers[LOGGER].(Loggable).Info("=================== BEFORE GetName()")
	name := p.provider.GetName()
	//c.providers[LOGGER].(Loggable).Info("=================== AFTER GetName()")
	return name
}

func (p *SampleFacade) SetName(name string) {
	p.provider.SetName("facade added: " + name)
}

// Router Facade
type RouteFacade struct {
	http.Handler
	provider Routable
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

func (r *RouteFacade) Get(path string, fn func(http.ResponseWriter, *http.Request)) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
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

func (r *RouteFacade) Serve() {
	c.showBanner()
	r.provider.Serve()
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

//func (s *sampleFacade) NewSample() Samplable {
//	return container.providers[SAMPLE].(Samplable).NewSample()
//}

//func (s *sampleFacade) GetName() string {
//	return sampleFac.GetName()
//}

//func (s sampleFacade) SetName(name string) {
//	sampleFac.self.SetName(name)
//}

/*
func (s sampleFacade) Register(config map[string]interface{}) interface{} {
	return s
}

func (s sampleFacade) GetKey() string {
	return "SAMPLE"
}
*/
/*
type routerFacade struct {
	http.Handler
	name string
}

var routerFac = routerFacade{}

// Router()
func (c Container) Router() Routable {
	return routerFac
}

func (r routerFacade) Register(config map[string]interface{}) interface{} {
	return r
}

func (r routerFacade) GetKey() string {
	return "ROUTER"
}

func (r routerFacade) Get(path string, fn func(rw http.ResponseWriter, req *http.Request)) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		routeMiddleware(container, rw, req, fn)
	}
	container.providers[ROUTER].(Routable).Get(path, nfn)
}

func (r routerFacade) Head(path string, fn func(rw http.ResponseWriter, req *http.Request)) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		routeMiddleware(container, rw, req, fn)
	}
	container.providers[ROUTER].(Routable).Head(path, nfn)
}

func (r routerFacade) Post(path string, fn func(rw http.ResponseWriter, req *http.Request)) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		routeMiddleware(container, rw, req, fn)
	}
	container.providers[ROUTER].(Routable).Post(path, nfn)
}

func (r routerFacade) Put(path string, fn func(rw http.ResponseWriter, req *http.Request)) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		routeMiddleware(container, rw, req, fn)
	}
	container.providers[ROUTER].(Routable).Put(path, nfn)
}

func (r routerFacade) Patch(path string, fn func(rw http.ResponseWriter, req *http.Request)) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		routeMiddleware(container, rw, req, fn)
	}
	container.providers[ROUTER].(Routable).Patch(path, nfn)
}

func (r routerFacade) Delete(path string, fn func(rw http.ResponseWriter, req *http.Request)) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		routeMiddleware(container, rw, req, fn)
	}
	container.providers[ROUTER].(Routable).Delete(path, nfn)
}

func (r routerFacade) Options(path string, fn func(rw http.ResponseWriter, req *http.Request)) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		routeMiddleware(container, rw, req, fn)
	}
	container.providers[ROUTER].(Routable).Options(path, nfn)
}

func (r routerFacade) Match(
	path string,
	fn func(rw http.ResponseWriter, req *http.Request), verbs ...string) {

	nfn := func(rw http.ResponseWriter, req *http.Request) {
		routeMiddleware(container, rw, req, fn)
	}
	container.providers[ROUTER].(Routable).Match(path, nfn, verbs...)

}

func (r routerFacade) All(path string, fn func(rw http.ResponseWriter, req *http.Request)) {
	nfn := func(rw http.ResponseWriter, req *http.Request) {
		routeMiddleware(container, rw, req, fn)
	}
	container.providers[ROUTER].(Routable).All(path, nfn)
}

func (r routerFacade) Serve() {
	container.providers[ROUTER].(Routable).Serve()
}

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
