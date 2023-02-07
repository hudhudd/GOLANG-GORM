package main

import (
	"go-fiber-gorm/database"
	"go-fiber-gorm/database/migration"
	"go-fiber-gorm/route"

	"github.com/gofiber/fiber/v2"
)

// import "fmt"

func main() {
	//INITIAL DATABASE
	database.DatabaseInit()
	migration.RunMigraiton()

	app := fiber.New()

	// INITIAL ROUTE
	route.RouteInit(app)

	app.Listen(":8080")
}
