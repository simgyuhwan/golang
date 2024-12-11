package route

import (
	"fmt"
	"gofiber/config"
	"gofiber/handler"

	"github.com/gofiber/fiber/v2"
)

func middleware(ctx *fiber.Ctx) error {
	token := ctx.Get("x-token")
	if token != "secret"{
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message" : "unauthenticated",
		})
	}
	return ctx.Next()
}

func RouteInit(r *fiber.App) {
	r.Static("/public", config.ProjectRootPath + "/public/asset")
	r.Get("/user", middleware,handler.UserHandlerGetAll)
	r.Get("/user/:id", handler.UserHandlerGetById)
	r.Post("/user", handler.UserHandlerCreate)
	r.Put("/user/:id", handler.UserHandlerUpdate)
	r.Put("/user/:id/update-email", handler.UserHandlerUpdateEmail)
	r.Delete("/user/:id", handler.UserHandlerDelete)
	fmt.Printf("-----------------")
	fmt.Printf("-----------------")
	fmt.Printf("-----------------")
	fmt.Printf("-----------------")
	fmt.Printf("-----------------")
	fmt.Printf("-----------------")
	fmt.Printf("-----------------")
	fmt.Printf("-----------------")
}