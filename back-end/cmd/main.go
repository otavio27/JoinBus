package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	app.Use(cors.New())

	app.Get("/api/geolocation/:lat/:lng", cto.GetLocation)
	app.Get("/api/linhas/:id", cto.GetItinerary)
	app.Get("/api/terminais", cto.GetTerminals)
	app.Get("/api/routes/:route", cto.GetRoutes)
	app.Get("/api/search/:text", cto.GetlinesRegexp)

	log.Fatal(app.Listen(":8000"))
}
