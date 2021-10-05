package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"log"
	"time"
)

func main() {
	app := fiber.New()
	app.Get("", LimiterApi(), func(ctx *fiber.Ctx) error {
		return ctx.SendString(string(ctx.Response().Header.Peek("X-Ratelimit-Limit")) + " - " + string(ctx.Response().Header.Peek("X-Ratelimit-Remaining")))
	})
	if err := app.Listen(fmt.Sprintf("0.0.0.0:%d", 3000)); err != nil {
		log.Fatal(err)
	}
}
func LimiterApi() fiber.Handler {
	isLocal := false
	return limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return isLocal
		},
		Max:        5,
		Expiration: time.Second * 5,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendStatus(429)
		},
	})
}
