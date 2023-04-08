package controllers

import (
	"github.com/vingarcia/krest"
)

type Controllers struct {
	module string
	http   krest.Provider
}

func New(http krest.Provider) *Controllers {
	return &Controllers{
		module: "Controllers",
		http:   http,
	}
}

/*
	func (cto Controllers) GetLocation(c *fiber.Ctx) error {
	var coord structs.Coordinate
	err := json.Unmarshal(c.Body(), &coord)
	if err != nil {
		return err
	}

	body, err := onibus.GetGeoLocation(c, coord.Latitude, coord.Longitude)
	if err != nil {
		return err
	}

	return c.JSON(body)
}
*/
