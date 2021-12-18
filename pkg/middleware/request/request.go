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
			err := cfg.PreRequest(c)
			if err != nil {
				return err
			}
		}
		// Continue stack
		if err := c.Next(); err != nil {
			return err
		}
		// Don't execute middleware if Next returns true
		if cfg.PostRequest != nil {
			err := cfg.PostRequest(c)
			if err != nil {
				return err
			}
		}
		// Continue stack
		return nil
	}
}
