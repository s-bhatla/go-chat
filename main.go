package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/s-bhatla/go-chat/handlers"
)

func main() {

	viewsEngine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: viewsEngine,
	})

	app.Static("/static/", "./static")

	app.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Welcome to the server")
	})

	appHandler := handlers.NewAppHandler()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return appHandler.HandleGetIndex(ctx)
	})
	server := NewWebSocket()
	app.Get("/ws", websocket.New(func(ctx *websocket.Conn) {
		server.HandleWebSocket(ctx)
	}))

	go server.HandleMessages()

	app.Listen(":3000")
}
