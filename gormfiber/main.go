package main

import (
	"gofiber/database"
	"gofiber/database/migration"
	"gofiber/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// init database
	database.DatabaseInit()

	// run migration
	migration.RunMigration()

	app := fiber.New()

	// init route
	route.RouteInit(app)

	app.Listen(":8888")
}