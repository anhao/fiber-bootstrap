package controllers

import "github.com/gofiber/fiber/v2"

func RenderIndex(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	})
}
