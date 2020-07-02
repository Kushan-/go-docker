package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	// "github.com/Kushan-/go-docker/goodbye"
	"github.com/Kushan-/go-docker/pkg/router"
	// "github.com/gofiber/fiber/middleware"

)

func main() {

	// Pass Settings creating a new instance
    app := fiber.New(&fiber.Settings{
        Prefork:       true,
        CaseSensitive: true,
        ServerHeader:  "Fiber",
	})
	// app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) {
		// Set some security headers:
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-Download-Options", "noopen")
		c.Set("Strict-Transport-Security", "max-age=5184000")
		c.Set("X-Frame-Options", "SAMEORIGIN")
		c.Set("X-DNS-Prefetch-Control", "off")
	  
		// Go to next middleware:
		c.Next()
	  
		// End of the chain
		fmt.Println("Bye 👋!")
		fmt.Println(c.OriginalURL())
		fmt.Println(c)
		
	})

	
	// Custom logging format
	// app.Use(middleware.Logger("${method} - ${path}"))

	router.SetupRoutes(app)
	app.Static("/", "./static")

	app.Listen(9120)
}

//https://github.com/gofiber/recipes/tree/master/auth-jwt