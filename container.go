package framework

const (
	LOGGER   = "LOGGER"
	ROUTER   = "ROUTER"
	RENDERER = "RENDERER"
	PARAMS   = "PARAMS"
	SAMPLE   = "SAMPLE"
)

var c *Container

type Config map[string]map[string]interface{}

//TODO Replace Context with concurrent-map: https://github.com/streamrail/concurrent-map
type Context map[string]interface{}

type Container struct {
	config      Config
	providers   map[string]Providerable
	middlewares []Middleware
	context     Context
}

func NewContainer(config ...Config) *Container {
	c = new(Container)
	c.providers = make(map[string]Providerable)
	c.context = make(Context)
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
	if key == LOGGER {
		showLoadingHeader()
	}
	container.providers[LOGGER].(Loggable).Info("| * " + key + " ✓")
}

// Middleware
func (container *Container) Use(mw MiddlewareHandler, args ...interface{}) {
	middleware := Middleware{handler: mw, args: args}
	container.middlewares = append(container.middlewares, middleware)
}

// App Container Context
func (container Container) Context() Context {
	return c.context
}

func (context Context) Read(key string) interface{} {
	return c.context[key]
}

func (context Context) Has(key string) bool {
	_, exists := c.context[key]
	return exists
}

func (context Context) Write(key string, data interface{}) {
	c.context[key] = data
}

func (context Context) Remove(key string) {
	delete(c.context, key)
}

func showLoadingHeader() {
	c.providers[LOGGER].(Loggable).Info(`|----------------------------------------|`)
	c.providers[LOGGER].(Loggable).Info(`| LOADING SERVICE PROVIDERS ...`)
	c.providers[LOGGER].(Loggable).Info(`|----------------------------------------|`)
}

func (container *Container) showBanner() {
	log := container.providers[LOGGER].(Loggable).NewLog()
	log.Info(`|----------------------------------------|`)
	log.Info(`|    _____                                `)
	log.Info(`|   / ____|           | |                 `)
	log.Info(`|  | |  __  ___  _ __ | |__   ___ _ __    `)
	log.Info(`|  | | |_ |/ _ \| '_ \| '_ \ / _ \ '__|   `)
	log.Info(`|  | |__| | (_) | |_) | | | |  __/ |      `)
	log.Info(`|   \_____|\___/| .__/|_| |_|\___|_|      `)
	log.Info(`|               | |                       `)
	log.Info(`|               |_|                       `)
	log.Info(`|----------------------------------------|`)
	log.Info(`| GOPHER READY FOR ACTION ON PORT 3000	`)
	log.Info(`|----------------------------------------|`)
}
