package framework

import (
	"os"
	"strconv"
)

const (
	LOGGER   = "LOGGER"
	ROUTER   = "ROUTER"
	RENDERER = "RENDERER"
	CONTEXT  = "CONTEXT"
)

var c *Container
var Initialized bool

type Container struct {
	config      Config
	providers   map[string]Providerable
	middlewares []Middleware
	Log         Loggable
	Route       Routable
	RouteGroup  Routegroupable
	Context     Contextable
	Render      Renderable
}

func NewContainer(config ...map[string]interface{}) *Container {
	c = new(Container)
	c.providers = make(map[string]Providerable)
	c.config = defaultConfig
	if len(config) > 0 {
		c.applyConfig(config[0])
	}
	c.setConfigFromEnv("PORT")
	c.setConfigFromEnv("HOST")
	return c
}

func (container *Container) setConfigFromEnv(env string) {
	if os.Getenv(env) != "" {
		router := c.config[ROUTER].(ConfigRouter)
		switch env {
		case "PORT":
			router.Port, _ = strconv.Atoi(os.Getenv(env))
		case "HOST":
			router.Host = os.Getenv(env)
		}
		c.config[ROUTER] = router
	}
}

func (container *Container) applyConfig(in Config) {
	if in[LOGGER] != nil {
		c.config[LOGGER] = applyConfigLogger(in)
	}
	if in[ROUTER] != nil {
		c.config[ROUTER] = applyConfigRouter(in)
	}
	if in[RENDERER] != nil {
		c.config[RENDERER] = ConfigRenderer(in[RENDERER].(ConfigRenderer))
	}
}

func applyConfigLogger(in Config) ConfigLogger {
	logger := c.config[LOGGER].(ConfigLogger)
	newLogger := ConfigLogger(in[LOGGER].(ConfigLogger))
	if newLogger.TimestampFormat != "" {
		logger.TimestampFormat = newLogger.TimestampFormat
	}
	if newLogger.LogLevel != 0 {
		logger.LogLevel = newLogger.LogLevel
	}
	return logger
}

func applyConfigRouter(in Config) ConfigRouter {
	router := c.config[ROUTER].(ConfigRouter)
	newRouter := ConfigRouter(in[ROUTER].(ConfigRouter))
	if newRouter.Host != "" {
		router.Host = newRouter.Host
	}
	if newRouter.Port != 0 {
		router.Port = newRouter.Port
	}
	if newRouter.StaticDirs != nil {
		router.StaticDirs = newRouter.StaticDirs
	}
	return router
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
	case CONTEXT:
		container.Context = provider.(Providerable).Register(c, config).(Contextable)
	case RENDERER:
		container.Render = provider.(Providerable).Register(c, config).(Renderable)
	case ROUTER:
		container.providers[key] = provider.(Providerable)
		container.Route = new(RouteFacade).Register(c, config).(Routable)
		container.RouteGroup = new(RouteGroupFacade)
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
	portInfo := ""
	if port > 0 {
		portInfo = `ON PORT ` + strconv.Itoa(port)
	}
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
	c.Log.Info(`| GOPHER READY FOR ACTION ` + portInfo)
	c.Log.Info(`|----------------------------------------|`)
}
