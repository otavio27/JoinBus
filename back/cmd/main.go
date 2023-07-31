package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/otavio27/JoinBus-APP/back-end/adapters/jsonlog"
	"github.com/otavio27/JoinBus-APP/back-end/adapters/onibus"
	"github.com/otavio27/JoinBus-APP/back-end/cmd/controllers"
	"github.com/otavio27/JoinBus-APP/back-end/cmd/middlewares"
	"github.com/vingarcia/krest"
)

func main() {
	stopsnear := os.Getenv("Stopsnear")
	stoptripslist := os.Getenv("Stoptripslist")
	timetable := os.Getenv("Timetable")
	group := os.Getenv("Group")
	referer := os.Getenv("Referer")
	host := os.Getenv("Host")

	ctx := context.Background()
	app := fiber.New()
	http := krest.New(30 * time.Second)
	ons := onibus.New(http, ctx, stopsnear, stoptripslist, timetable, group, referer, host)
	cto := controllers.New(ctx, http, *ons)
	logger := jsonlog.New("info")
	errHandler := middlewares.NewErrorHandler(logger)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://owtechsystems.com/, http://*.owtechsystems.com/",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS, HEAD",
	}))

	app.Use(errHandler.Middleware)

	app.Get("/api/geolocation/:lat/:lng", cto.GetLocation)
	app.Get("/api/linhas/:id", cto.GetItinerary)
	app.Get("/api/terminais", cto.GetTerminals)
	app.Get("/api/routes/:route", cto.GetRoutes)
	app.Get("/api/search/:text", cto.GetlinesRegexp)

	log.Fatal(app.Listen(":8000"))
}
