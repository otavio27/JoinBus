package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/otavio27/JoinBus-APP/back-end/onibus"
	"github.com/otavio27/JoinBus-APP/back-end/structs"
	"github.com/vingarcia/krest"
)

type Controllers struct {
	module string
	http   krest.Provider
	ons    onibus.Adapter
}

func New(ctx context.Context, http krest.Provider, ons onibus.Adapter) *Controllers {
	return &Controllers{
		module: "Controllers",
		http:   http,
		ons:    ons,
	}
}

func (cto Controllers) GetLocation(c *fiber.Ctx) error {
	latitude := c.Params("lat")
	longitude := c.Params("lng")

	body, stopName, err := cto.ons.GetGeoLocation(latitude, longitude)
	if err != nil {
		return err
	}

	var stoplists []structs.MyStopList
	if err := json.Unmarshal(body, &stoplists); err != nil {
		return fmt.Errorf("fail to parse body as JSON: " + err.Error())
	}

	var linhas []map[string]any
	var Name, Id, Router string
	var hours []string

	for _, stds := range stoplists {
		Name = strings.Split(stds.TripHeadsign, "-")[0]
		Id = strings.Split(stds.ShapeID, "-")[0]
		Router = strings.Split(stds.TripHeadsign, "-")[1]
		for _, hora := range stds.Trips {
			HRS := strings.Split(hora.Eta, ":")
			hours = append(hours, HRS[0]+":"+HRS[1])
		}
		linhas = append(linhas, map[string]any{
			"NearPoint": stopName[0],
			"ID":        Id,
			"Name":      Name,
			"Direction": Router,
			"Hours":     strings.Join(hours, " "),
		})
	}

	if len(linhas) == 0 {
		return c.JSON(map[string]any{"Warning": "Não há linhas que rodam nesta localização, ou nas próximas horas."})
	}

	return c.JSON(linhas)
}

func (cto Controllers) GetLines(c *fiber.Ctx) error {
	var id = c.Params("id")

	body, err := cto.ons.GetjsonLines(c.Context(), id)
	if err != nil {
		return err
	}

	var lines []structs.LineStruct
	err = json.Unmarshal(body, &lines)
	if err != nil {
		return fmt.Errorf("Unmarshal error, not found files %s", err)
	}

	var direction string
	var operatesToday bool
	var hours []string
	var linha []map[string]any

	for _, data := range lines {
		for _, stopData := range data.StopData {
			switch data.Direction {
			case "Ida":
				direction = "Saída: " + stopData.StopName
			case "Volta":
				hours = []string{}
				direction = "Saída: " + stopData.StopName
			}

			for _, servData := range stopData.ServiceData {
				if servData.Type == cto.getServiceTypeForToday() {
					operatesToday = true
					for _, timeData := range servData.TimeData {
						for _, schedules := range timeData {
							hours = append(hours, schedules.DepartureTime)
						}
					}
				}
			}

			linha = append(linha, map[string]any{
				"Weekday":   cto.getServiceTypeForToday(),
				"ID":        id,
				"Name":      cto.getNameLines(c.Context(), id),
				"Station":   stopData.StopName,
				"Direction": direction,
				"Hours":     hours,
			})
		}
	}

	if operatesToday {
		return c.JSON(linha)
	}
	return c.JSON(map[string]any{"Warning": "Linha sem operação nesta data!"})
}

// GetTerminals é uma função que retorna informações sobre terminais em um objeto JSON.
func (cto Controllers) GetTerminals(c *fiber.Ctx) error {
	body, err := cto.ons.GetjsonTerminals(c.Context())
	if err != nil {
		return err
	}

	var Stations []structs.MyStations
	err = json.Unmarshal(body, &Stations)
	if err != nil {
		return fmt.Errorf("Unmarshal error, not found files %s", err)
	}

	var terminals []map[string]string
	for _, TRM := range Stations {
		terminals = append(terminals, map[string]string{
			"Station": TRM.StationName,
		})
	}
	return c.JSON(terminals)
}

// getNameLines retona os nomes das linhas contidas no terminal
func (cto Controllers) getNameLines(ctx context.Context, text string) string {
	var linename string

	body, err := cto.ons.GetjsonTerminals(ctx)
	if err != nil {
		return err.Error()
	}

	var Stations []structs.MyStations
	err = json.Unmarshal(body, &Stations)
	if err != nil {
		return err.Error()
	}

	for _, TRM := range Stations {
		for _, lines := range TRM.Routes {
			if lines.RouteID == text {
				linename = lines.RouteLongName
			}
		}
	}
	return linename
}

func (cto Controllers) getServiceTypeForToday() string {
	var daysweek string

	switch time.Now().Local().Weekday() {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		daysweek = "Dias Úteis"
	case time.Saturday:
		daysweek = "Sábados"
	case time.Sunday:
		daysweek = "Domingos"
	}
	return daysweek
}
