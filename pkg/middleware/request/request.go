package request

import (
	"github.com/gofiber/fiber/v2"
)

// New creates a new middleware handler
func New(config ...Config) fiber.Handler {
	// Set default config
	cfg := configDefault(config...)

	// Return new handler
	return func(c *fiber.Ctx) error {
		// Don't execute middleware if Next returns true
		if cfg.PreRequest != nil {
			cfg.PreRequest(c)
		}
		// Continue stack
		if err := c.Next(); err != nil {
			return err
		}
		// Don't execute middleware if Next returns true
		if cfg.PostRequest != nil {
			cfg.PostRequest(c)
		}
		// Continue stack
		return nil
	}
}
