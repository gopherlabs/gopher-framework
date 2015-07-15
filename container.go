package framework

const (
	LOGGER   = "LOGGER"
	ROUTER   = "ROUTER"
	RENDERER = "RENDERER"
	PARAMS   = "PARAMS"
	MAPPER   = "MAPPER"
)

var c *Container

type Config map[string]map[string]interface{}

type Container struct {
	config      Config
	providers   map[string]Providerable
	middlewares []Middleware
	Log         Loggable
	Route       Routable
	Context     Mappable
	Render      Renderable
}

func NewContainer(config ...Config) *Container {
	c = new(Container)
	c.providers = make(map[string]Providerable)
	if len(config) > 0 {
		c.config = config[0]
	}
	return c
}

func (container *Container) RegisterProvider(provider interface{}) {
	key := provider.(Providerable).GetKey()
	config := defaultConfig[key]
	if len(container.config) > 0 {
		config = container.config[key]
	}
	container.providers[key] = provider.(Providerable).Register(config).(Providerable)
	switch key {
	case LOGGER:
		container.Log = container.providers[key].(Loggable)
		showLoadingHeader()
	case MAPPER:
		container.Context = container.providers[key].(Mappable)
	case RENDERER:
		container.Render = container.providers[key].(Renderable)
	case ROUTER:
		container.Route = new(RouteFacade).Register(config).(Routable)
	}
	container.Log.Info("| * " + key + " ✓")
}

// Middleware
func (container *Container) Use(mw MiddlewareHandler, args ...interface{}) {
	middleware := Middleware{handler: mw, args: args}
	container.middlewares = append(container.middlewares, middleware)
}

func showLoadingHeader() {
	c.Log.Info(`|----------------------------------------|`)
	c.Log.Info(`| LOADING SERVICE PROVIDERS ...`)
	c.Log.Info(`|----------------------------------------|`)
}

func (container *Container) showBanner() {
	c.Log.Info(`|----------------------------------------|`)
	c.Log.Info(`|    _____                                `)
	c.Log.Info(`|   / ____|           | |                 `)
	c.Log.Info(`|  | |  __  ___  _ __ | |__   ___ _ __    `)
	c.Log.Info(`|  | | |_ |/ _ \| '_ \| '_ \ / _ \ '__|   `)
	c.Log.Info(`|  | |__| | (_) | |_) | | | |  __/ |      `)
	c.Log.Info(`|   \_____|\___/| .__/|_| |_|\___|_|      `)
	c.Log.Info(`|               | |                       `)
	c.Log.Info(`|               |_|                       `)
	c.Log.Info(`|----------------------------------------|`)
	c.Log.Info(`| GOPHER READY FOR ACTION ON PORT 3000	`)
	c.Log.Info(`|----------------------------------------|`)
}
