package pagarmeapisdk

import (
	"os"

	"github.com/apimatic/go-core-runtime/https"
)

type ConfigurationOptions func(*Configuration)

type Configuration struct {
	serviceRefererName string
	environment        Environment
	httpConfiguration  https.HttpConfiguration
	basicAuthUserName  string
	basicAuthPassword  string
}

func newConfiguration(options ...ConfigurationOptions) Configuration {
	config := Configuration{}

	for _, option := range options {
		option(&config)
	}
	return config
}

func WithServiceRefererName(serviceRefererName string) ConfigurationOptions {
	return func(c *Configuration) {
		c.serviceRefererName = serviceRefererName
	}
}

func WithEnvironment(environment Environment) ConfigurationOptions {
	return func(c *Configuration) {
		c.environment = environment
	}
}

func WithHttpConfiguration(httpConfiguration https.HttpConfiguration) ConfigurationOptions {
	return func(c *Configuration) {
		c.httpConfiguration = httpConfiguration
	}
}

func WithBasicAuthUserName(basicAuthUserName string) ConfigurationOptions {
	return func(c *Configuration) {
		c.basicAuthUserName = basicAuthUserName
	}
}

func WithBasicAuthPassword(basicAuthPassword string) ConfigurationOptions {
	return func(c *Configuration) {
		c.basicAuthPassword = basicAuthPassword
	}
}

func (c *Configuration) ServiceRefererName() string {
	return c.serviceRefererName
}

func (c *Configuration) Environment() Environment {
	return c.environment
}

func (c *Configuration) HttpConfiguration() https.HttpConfiguration {
	return c.httpConfiguration
}

func (c *Configuration) BasicAuthUserName() string {
	return c.basicAuthUserName
}

func (c *Configuration) BasicAuthPassword() string {
	return c.basicAuthPassword
}

func CreateConfigurationFromEnvironment(options ...ConfigurationOptions) Configuration {
	config := DefaultConfiguration()

	serviceRefererName := os.Getenv("PAGARMEAPISDK_SERVICE_REFERER_NAME")
	if serviceRefererName != "" {
		config.serviceRefererName = serviceRefererName
	}
	environment := os.Getenv("PAGARMEAPISDK_ENVIRONMENT")
	if environment != "" {
		config.environment = Environment(environment)
	}
	basicAuthUserName := os.Getenv("PAGARMEAPISDK_BASIC_AUTH_USER_NAME")
	if basicAuthUserName != "" {
		config.basicAuthUserName = basicAuthUserName
	}
	basicAuthPassword := os.Getenv("PAGARMEAPISDK_BASIC_AUTH_PASSWORD")
	if basicAuthPassword != "" {
		config.basicAuthPassword = basicAuthPassword
	}
	for _, option := range options {
		option(&config)
	}
	return config
}

// Available Servers
type Server string

const (
	CORE_V5 Server = "core_v5"
	BIN_V1  Server = "bin_v1"
)

// Available Environments
type Environment string

const (
	PRODUCTION Environment = "production"
)

func CreateRetryConfiguration(options ...https.RetryConfigurationOptions) https.RetryConfiguration {
	config := DefaultRetryConfiguration()

	for _, option := range options {
		option(&config)
	}
	return config
}

func CreateHttpConfiguration(options ...https.HttpConfigurationOptions) https.HttpConfiguration {
	config := DefaultHttpConfiguration()

	for _, option := range options {
		option(&config)
	}
	return config
}

func CreateConfiguration(options ...ConfigurationOptions) Configuration {
	config := DefaultConfiguration()

	for _, option := range options {
		option(&config)
	}
	return config
}
