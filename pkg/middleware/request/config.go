package request

import (
	"github.com/gofiber/fiber/v2"
)

// Config defines the config for middleware.
type Config struct {
	// PreRequest defines a function to process pre-request
	//
	// Optional. Default: nil
	PreRequest func(c *fiber.Ctx) error
	// PostRequest defines a function to process post-request
	//
	// Optional. Default: nil
	PostRequest func(c *fiber.Ctx) error
}

// ConfigDefault is the default config
var ConfigDefault = Config{
	PreRequest:  nil,
	PostRequest: nil,
}

// Helper function to set default values
func configDefault(config ...Config) Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]
	return cfg
}
