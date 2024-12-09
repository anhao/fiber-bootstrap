package router

import (
	"github.com/anhao/fiber-bootstrap/app/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
)

type HttpRouter struct{}

func (h HttpRouter) InstallRouter(app *fiber.App) {
	group := app.Group("", cors.New(), csrf.New())
	group.Get("", controllers.RenderIndex)
}

func NewHttpRouter() *HttpRouter {
	return &HttpRouter{}
}
