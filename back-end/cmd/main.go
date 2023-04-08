package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/otavio27/JoinBus-APP/back-end/onibus"
	"github.com/otavio27/JoinBus-APP/back-end/structs"
	"github.com/vingarcia/krest"
)

func main() {
	ctx := context.Background()

	http := krest.New(30 * time.Second)
	app := fiber.New()
	onibus := onibus.New(http)

	app.Get("/api/geolocation", func(c *fiber.Ctx) error {
		var coord structs.Coordinate
		err := json.Unmarshal(c.Body(), &coord)
		if err != nil {
			return err
		}

		body, err := onibus.GetGeoLocation(ctx, coord.Latitude, coord.Longitude)
		if err != nil {
			return err
		}

		return c.JSON(body)
	})

	log.Fatal(app.Listen(":3000"))
}
