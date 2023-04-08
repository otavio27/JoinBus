package onibus

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/otavio27/JoinBus-APP/back-end/structs"
	"github.com/vingarcia/krest"
)

type Adapter struct {
	module string
	http   krest.Provider
}

func New(http krest.Provider) *Adapter {
	return &Adapter{
		module: "Adapter",
		http:   http,
	}
}

// GetLocation busca as linhas mais proximas da localização passada pelo app
func (a Adapter) GetGeoLocation(ctx context.Context, latitude float64, longitude float64) ([]map[string]string, error) {
	latd := fmt.Sprintf("%f", latitude)
	logd := fmt.Sprintf("%f", longitude)
	url := "https://onibus.info/api/stopsnear?lat=" + latd + "&lng=" + logd
	resp, err := a.http.Get(ctx, url, krest.RequestData{
		Headers: map[string]string{
			"Referer":    "https://onibus.info/mapa/",
			"Connection": "keep-alive",
			"Host":       "onibus.info",
		},
	})
	if err != nil {
		if resp.StatusCode == 404 {
			return nil, fmt.Errorf("onibus.info/api/stopsnear?lat= was not found! %s", err)
		}
		return nil, fmt.Errorf("unexpected error when fetching example.com: %s", err)
	}

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ = gzip.NewReader(resp)
		defer reader.Close()
	default:
		reader = resp
	}

	location, _ := ioutil.ReadAll(reader)

	var local []structs.MyLocation
	if err := json.Unmarshal(location, &local); err != nil {
		return nil, fmt.Errorf("fail to parse body as JSON: " + err.Error())
	}

	var stop_id, stop_name []string
	for _, lctn := range local {
		stop_id = append(stop_id, lctn.StopID)
		stop_name = append(stop_name, lctn.StopName)
	}
	return a.GetStopTripList(ctx, stop_id, stop_name)

}

// GetStopTripList busca as linhas que passam pelo ponto informado através da localização
func (a Adapter) GetStopTripList(ctx context.Context, stop []string, stopName []string) ([]map[string]string, error) {
	var linhas []map[string]string
	for _, std := range stop[:2] {
		url := "https://onibus.info/api/stoptrips/" + std

		resp, err := a.http.Get(ctx, url, krest.RequestData{
			Headers: map[string]string{
				"Referer":    "https://onibus.info/mapa/",
				"Connection": "keep-alive",
				"Host":       "onibus.info",
			},
		})
		if err != nil {
			if resp.StatusCode == 404 {
				return nil, fmt.Errorf("onibus.info/api/stopsnear?lat= was not found! %s", err)
			}
			return nil, fmt.Errorf("unexpected error when fetching example.com: %s", err)
		}

		var reader io.ReadCloser
		switch resp.Header.Get("Content-Encoding") {
		case "gzip":
			reader, _ = gzip.NewReader(resp)
			defer reader.Close()
		default:
			reader = resp
		}

		stoplist, _ := ioutil.ReadAll(reader)

		var stoplists []structs.MyStopList
		if err := json.Unmarshal(stoplist, &stoplists); err != nil {
			return nil, fmt.Errorf("fail to parse body as JSON: " + err.Error())
		}

		keys := make(map[string]bool)
		var Name, Id, Router, Hours string
		for _, stds := range stoplists {
			Hours = ""
			Name = strings.Split(stds.TripHeadsign, "-")[0]
			Id = strings.Split(stds.ShapeID, "-")[0]
			Router = strings.Split(stds.TripHeadsign, "-")[1]
			for _, hora := range stds.Trips {
				HRS := strings.Split(hora.Eta, ":")
				Hours += fmt.Sprintf("%s:%s", HRS[0], HRS[1]) + " "
			}
			if _, value := keys[Id]; !value {
				keys[Id] = true
				linhas = append(linhas, map[string]string{
					"Ponto mais próximo": stopName[0],
					"ID":                 Id,
					"Linha":              Name,
					"Direcao":            Router,
					"Horarios":           Hours,
				})
			}

		}
	}

	if len(linhas) == 0 {
		return nil, fmt.Errorf("AVISO Não há linhas que rodam nesta localização, ou nas próximas horas.")
	}

	return linhas, nil
}
