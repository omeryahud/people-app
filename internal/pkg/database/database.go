package database

import (
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber"
)

func Main() {
	if err := validateEnvironmentVariables(); err != nil {
		log.Fatalln(err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, Backend!")
	})

	strPort, _ := os.LookupEnv(httpPortKey)
	port, err := strconv.Atoi(strPort)
	if err != nil {
		log.Fatalf("could not convert port to integer %v", err)
	}

	err = app.Listen(port)
	if err != nil {
		log.Fatalf("failed to listen on port %d: %v", port, err)
	}
}
