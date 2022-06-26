package main

import (
	"log"

	"siakad-rest/config"
	"siakad-rest/db"
	"siakad-rest/pkg/auth"
	"siakad-rest/pkg/user"

	"github.com/gofiber/fiber/v2"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	app := fiber.New()
	db := db.Init(c.DBUrl)

	auth.RegisterRoutes(app, db)
	user.RegisterRoutes(app, db)

	// books.RegisterRoutes(app, db)

	app.Listen(c.Port)
}
