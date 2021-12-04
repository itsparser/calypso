package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/workfoxes/calypso/pkg/log"
	"strconv"
)

func GetPageAndLimit(c *fiber.Ctx) (int64, int64) {
	page := GetQueryInt(c, "page", "1")
	limit := GetQueryInt(c, "limit", "10")
	return page, limit
}

func GetQueryInt(c *fiber.Ctx, key string, defaultValue ...string) int64 {
	data, err := strconv.ParseInt(GetQuery(c, key, defaultValue...), 10, 64)
	if err != nil {
		log.Error(err)
	}
	return data
}

func GetQuery(c *fiber.Ctx, key string, defaultValue ...string) string {
	return c.Query(key, defaultValue...)
}
