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
	stopsnear := os.Getenv("Stopsnear")
	stoptripslist := os.Getenv("Stoptripslist")
	timetable := os.Getenv("Timetable")
	group := os.Getenv("Group")
	referer := os.Getenv("Referer")
	host := os.Getenv("Host")

	fmt.Println("Valor de Stopsnear: ", stopsnear)
	fmt.Println("Valor de Stoptripslist: ", stoptripslist)
	fmt.Println("Valor de Timetable: ", timetable)
	fmt.Println("Valor de Group: ", group)
	fmt.Println("Valor de Referer: ", referer)
	fmt.Println("Valor de Host: ", host)

	ctx := context.Background()
	app := fiber.New()
	http := krest.New(30 * time.Second)
	ons := onibus.New(http, ctx, stopsnear, stoptripslist, timetable, group, referer, host)
	cto := controllers.New(ctx, http, *ons)
	logger := jsonlog.New("info")
	errHandler := middlewares.NewErrorHandler(logger)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	app.Get("/api/geolocation/:lat/:lng", errHandler.Middleware, cto.GetLocation)
	app.Get("/api/linhas/:id", errHandler.Middleware, cto.GetItinerary)
	app.Get("/api/terminais", errHandler.Middleware, cto.GetTerminals)
	app.Get("/api/routes/:route", errHandler.Middleware, cto.GetRoutes)
	app.Get("/api/search/:text", errHandler.Middleware, cto.GetlinesRegexp)

	log.Fatal(app.Listen(":9750"))
}
