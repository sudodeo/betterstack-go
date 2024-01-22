package sdk

// Config contains parameters for configuring the SDK.
type Config struct {
	// APIToken is the BetterStack API token
	APIToken string
	// APIType is the BetterStack API type. "uptime" or "logs"
	APIType string
}

// NewContext creates a new Context.
func NewContext(config Config) *Context {
	return &Context{
		Config: config,
	}
}

// Context contains the SDK configuration.
type Context struct {
	Config Config
}
