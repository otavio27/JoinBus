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

var ctx = context.Background()
var app = fiber.New()
var http = krest.New(30 * time.Second)
var ons = onibus.New(http, ctx)
var cto = controllers.New(ctx, http, *ons)

func main() {
	app.Get("/api/geolocation", cto.GetLocation)
	app.Get("/api/linhas/:id", cto.GetLines)
	app.Get("/api/terminais", cto.GetTerminals)

	log.Fatal(app.Listen(":3000"))
}
