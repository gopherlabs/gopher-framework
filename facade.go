package framework

import "net/http"

var sampleFac = new(SampleFacade)

func (c Container) NewSample() Samplable {
	return sampleFac.NewSample()
}

type SampleFacade struct {
	name string
}

func (p *SampleFacade) Register(config map[string]interface{}) interface{} {
	return p
}

func (p *SampleFacade) GetKey() string {
	return "SAMPLE"
}

func (p *SampleFacade) NewSample() Samplable {
	return new(SampleFacade)
}

func (p *SampleFacade) GetName() string {
	//p.log.Info()

	//	if p.container.NewLog() != nil {
	//		p.container.NewLog().Info("=================== BEFORE GetName()")
	//	}

	return ">" + p.name
}

func (p *SampleFacade) SetName(name string) {
	p.name = name
}

// Logger
func (c Container) NewLog() Loggable {
	return c.providers[LOGGER].(Loggable).NewLog()
}

// Router
func (c Container) NewRouter() Routable {
	return c.providers[ROUTER].(Routable).NewRouter()
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
