package main

import (
	"context"
	"fmt"
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
	fmt.Println("main step: 1")
	stopsnear := os.Getenv("Stopsnear")
	stoptripslist := os.Getenv("Stoptripslist")
	timetable := os.Getenv("Timetable")
	group := os.Getenv("Group")
	referer := os.Getenv("Referer")
	host := os.Getenv("Host")
	fmt.Println("main step: 2")
	ctx := context.Background()
	app := fiber.New()
	http := krest.New(30 * time.Second)
	ons := onibus.New(http, ctx, stopsnear, stoptripslist, timetable, group, referer, host)
	cto := controllers.New(ctx, http, *ons)
	logger := jsonlog.New("info")
	errHandler := middlewares.NewErrorHandler(logger)
	fmt.Println("main step: 3")
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	app.Get("/api/geolocation/:lat/:lng", errHandler.Middleware, cto.GetLocation)
	app.Get("/api/linhas/:id", errHandler.Middleware, cto.GetItinerary)
	app.Get("/api/terminais", errHandler.Middleware, cto.GetTerminals)
	app.Get("/api/routes/:route", errHandler.Middleware, cto.GetRoutes)
	app.Get("/api/search/:text", errHandler.Middleware, cto.GetlinesRegexp)
	fmt.Println("main step: 4")
	log.Fatal(app.Listen(":8750"))
}
