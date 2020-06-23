package main

import (
	
	"github.com/gofiber/fiber"

)

func main() {

	// Pass Settings creating a new instance
    app := fiber.New(&fiber.Settings{
        Prefork:       true,
        CaseSensitive: true,
        ServerHeader:  "Fiber",
	})
	
	app.Static("/", "./static")

	app.Listen(9120)
}

//https://github.com/gofiber/recipes/tree/master/auth-jwt