package server

import (
	"fmt"

	"github.com/gofiber/fiber"
)

func Serve() {
	host := fiber.New()
	host.Listen(":3000")
	fmt.Println("listening on port :3000")
}
