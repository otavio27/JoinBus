package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/otavio27/JoinBus-APP/back-end/controllers"
	"github.com/otavio27/JoinBus-APP/back-end/onibus"
	"github.com/vingarcia/krest"
)

func main() {
	ctx := context.Background()
	app := fiber.New()
	http := krest.New(30 * time.Second)
	ons := onibus.New(http, ctx)
	cto := controllers.New(ctx, http, *ons)

	app.Get("/api/geolocation", cto.GetLocation)
	app.Get("/api/linhas/:id", cto.GetLines)
	app.Get("/api/terminais", cto.GetTerminals)

	log.Fatal(app.Listen(":8000"))
}
