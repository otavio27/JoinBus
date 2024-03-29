package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/otavio27/JoinBus-APP/back-end/adapters/jsonlog"
	"github.com/otavio27/JoinBus-APP/back-end/adapters/onibus"
	"github.com/otavio27/JoinBus-APP/back-end/cmd/controllers"
	"github.com/otavio27/JoinBus-APP/back-end/cmd/middlewares"
	"github.com/otavio27/JoinBus-APP/back-end/helpers/env"
	"github.com/vingarcia/krest"
)

func main() {
	stopsnear := env.MustGetString("Stopsnear")
	stoptripslist := env.MustGetString("Stoptripslist")
	timetable := env.MustGetString("Timetable")
	group := env.MustGetString("Group")
	referer := env.MustGetString("Referer")
	host := env.MustGetString("Host")
	port := env.GetString("port", "80")

	ctx := context.Background()
	http := krest.New(30 * time.Second)
	ons := onibus.New(http, ctx, stopsnear, stoptripslist, timetable, group, referer, host)
	cto := controllers.New(ctx, http, *ons)
	logger := jsonlog.New("info")
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		startTime := time.Now()
		err := c.Next()
		logger.Info(ctx, "request received", map[string]any{
			"method":   c.Method(),
			"path":     c.Path(),
			"duration": time.Since(startTime),
		})
		return err
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]any{
			"service": "JoinBus",
		})
	})

	errHandler := middlewares.NewErrorHandler(logger)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	app.Get("/joinbus/geolocation/:lat/:lng", errHandler.Middleware, cto.GetLocation)
	app.Get("/joinbus/linhas/:id", errHandler.Middleware, cto.GetItinerary)
	app.Get("/joinbus/terminais", errHandler.Middleware, cto.GetTerminals)
	app.Get("/joinbus/routes/:route", errHandler.Middleware, cto.GetRoutes)
	app.Get("/joinbus/search/:text", errHandler.Middleware, cto.GetlinesRegexp)

	err := app.Listen(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
