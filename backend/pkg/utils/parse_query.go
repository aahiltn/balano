package utils

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// ParseQueryParams extracts limit and offset from query parameters
func ParseQueryParams(c *fiber.Ctx) (int, int) {
	limit := 10 // default limit
	offset := 0 // default offset

	if c.Query("limit") != "" {
		if parsedLimit, err := strconv.Atoi(c.Query("limit")); err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	if c.Query("offset") != "" {
		if parsedOffset, err := strconv.Atoi(c.Query("offset")); err == nil && parsedOffset >= 0 {
			offset = parsedOffset
		}
	}

	return limit, offset
}
