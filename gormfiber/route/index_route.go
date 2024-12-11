package route

import (
	"gofiber/config"
	"gofiber/handler"
	"gofiber/middleware"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Static("/public", config.ProjectRootPath + "/public/asset")
	r.Get("/user", middleware.Auth ,handler.UserHandlerGetAll)
	r.Get("/user/:id", handler.UserHandlerGetById)
	r.Post("/login", handler.LoginHandler)
	r.Post("/user", handler.UserHandlerCreate)
	r.Put("/user/:id", handler.UserHandlerUpdate)
	r.Put("/user/:id/update-email", handler.UserHandlerUpdateEmail)
	r.Delete("/user/:id", handler.UserHandlerDelete)
}