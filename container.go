package framework

import "strconv"

const (
	LOGGER   = "LOGGER"
	ROUTER   = "ROUTER"
	RENDERER = "RENDERER"
	MAPPER   = "MAPPER"
)

var c *Container
var Initialized bool

type Container struct {
	config      Config
	providers   map[string]Providerable
	middlewares []Middleware
	Log         Loggable
	Route       Routable
	Context     Mappable
	Render      Renderable
}

func NewContainer(config ...map[string]interface{}) *Container {
	c = new(Container)
	c.providers = make(map[string]Providerable)
	c.config = defaultConfig
	if len(config) > 0 {
		c.applyConfig(config[0])
	}
	return c
}

func (container *Container) applyConfig(in Config) {
	if in[LOGGER] != nil {
		c.config[LOGGER] = ConfigLogger(in[LOGGER].(ConfigLogger))
	}
	if in[ROUTER] != nil {
		c.config[ROUTER] = ConfigRouter(in[ROUTER].(ConfigRouter))
	}
	if in[RENDERER] != nil {
		c.config[RENDERER] = ConfigRenderer(in[RENDERER].(ConfigRenderer))
	}
}

func (container *Container) RegisterProvider(provider interface{}) {
	key := provider.(Providerable).GetKey()
	config := defaultConfig[key]
	if len(container.config) > 0 {
		config = container.config[key]
	}
	switch key {
	case LOGGER:
		container.Log = provider.(Providerable).Register(c, config).(Loggable)
		if Initialized == false {
			showLoadingHeader()
		}
	case MAPPER:
		container.Context = provider.(Providerable).Register(c, config).(Mappable)
	case RENDERER:
		container.Render = provider.(Providerable).Register(c, config).(Renderable)
	case ROUTER:
		container.providers[key] = provider.(Providerable)
		container.Route = new(RouteFacade).Register(c, config).(Routable)
	}
	if Initialized == false {
		container.Log.Info("| * " + key + " ✓")
	}
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

func (container *Container) showBanner(port int) {
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
	c.Log.Info(`| GOPHER READY FOR ACTION ON PORT ` + strconv.Itoa(port))
	c.Log.Info(`|----------------------------------------|`)
}
