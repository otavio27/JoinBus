package onibus

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/otavio27/JoinBus-APP/back-end/structs"
	"github.com/vingarcia/ddd-go-template/v2-domain-adapters-and-helpers/domain"
	"github.com/vingarcia/krest"
)

type Adapter struct {
	module string
	http   krest.Provider
	ctx    context.Context
}

func New(http krest.Provider, ctx context.Context) *Adapter {
	return &Adapter{
		module: "Adapter",
		http:   http,
		ctx:    ctx,
	}
}

// GetLocation busca as linhas mais proximas da localização passada pelo app
func (a Adapter) GetGeoLocation(latitude string, longitude string) ([]byte, []string, error) {
	url := os.Getenv("Stopsnear") + latitude + "&lng=" + longitude
	resp, err := a.http.Get(a.ctx, url, krest.RequestData{
		Headers: map[string]string{
			"Referer":    "https://onibus.info/mapa/",
			"Connection": "keep-alive",
			"Host":       "onibus.info",
		},
	})
	if err != nil {

		if resp.StatusCode == 404 {
			return nil, nil, fmt.Errorf("onibus.info/api/stopsnear?lat= was not found! %s", err)
		}
		return nil, nil, fmt.Errorf("unexpected error when fetching geolocation: %s", err)
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

	var local []structs.Location
	if err := json.Unmarshal(location, &local); err != nil {
		return nil, nil, fmt.Errorf("fail to parse body as JSON: " + err.Error())
	}

	var stop_id, stop_name []string
	for _, lctn := range local {
		stop_id = append(stop_id, lctn.StopID)
		stop_name = append(stop_name, lctn.StopName)
	}

	return a.GetStopTripList(a.ctx, stop_id, stop_name)
}

// GetStopTripList busca as linhas que passam pelo ponto informado através da localização
func (a Adapter) GetStopTripList(ctx context.Context, stop []string, stopName []string) ([]byte, []string, error) {
	var stoplist []byte

	if len(stop) == 0 {
		return nil, nil, fmt.Errorf("variavel stop está vazia")
	}

	for _, std := range stop[:2] {
		url := os.Getenv("Stoptripslist") + std

		resp, err := a.http.Get(ctx, url, krest.RequestData{
			Headers: map[string]string{
				"Referer":    "https://onibus.info/mapa/",
				"Connection": "keep-alive",
				"Host":       "onibus.info",
			},
		})
		if err != nil {
			if resp.StatusCode == 404 {
				return nil, nil, fmt.Errorf("onibus.info/api/stopsnear?lat= was not found! %s", err)
			}
			return nil, nil, fmt.Errorf("unexpected error when fetching stop list: %s", err)
		}

		var reader io.ReadCloser
		switch resp.Header.Get("Content-Encoding") {
		case "gzip":
			reader, _ = gzip.NewReader(resp)
			defer reader.Close()
		default:
			reader = resp
		}

		stoplist, err = ioutil.ReadAll(reader)
		if err != nil {
			return nil, nil, fmt.Errorf("expected error reading reader: %s", err)
		}

	}

	return stoplist, stopName, nil
}

// GetjsonLines função que tem responssabilidade de busacar os horários da linha quando passada por nome ou ID
func (a Adapter) GetItineraries(ctx context.Context, id string) ([]structs.Itinerary, error) {
	url := os.Getenv("Timetable") + id

	resp, err := a.http.Get(ctx, url, krest.RequestData{
		Headers: map[string]string{
			"Referer":    "https://onibus.info/mapa/",
			"Connection": "keep-alive",
			"Host":       "onibus.info",
		},
	})

	if err != nil {
		if resp.StatusCode == 404 {
			return nil, domain.NotFoundErr("itineraries not found", map[string]any{
				"line_id": id,
			})
		}
		return nil, domain.InternalErr("unexpected error when fetching itineraries", map[string]any{
			"error":   err.Error(),
			"line_id": id,
		})
	}

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp)
		if err != nil {
			return nil, domain.InternalErr("unexpected error unzipping itineraries from external api", map[string]any{
				"error":   err.Error(),
				"line_id": id,
			})
		}
		defer reader.Close()
	default:
		reader = resp
	}

	hours, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, domain.InternalErr("unexpected error reading itineraries from external api", map[string]any{
			"error":   err.Error(),
			"line_id": id,
		})
	}

	var itineraries []structs.Itinerary
	err = json.Unmarshal(hours, &itineraries)
	if err != nil {
		return nil, domain.InternalErr("error parsing itineraries as json", map[string]any{
			"error":   err.Error(),
			"payload": string(hours),
			"line_id": id,
		})
	}

	return itineraries, nil
}

// GetjsonTerminals busca todas as linhas de cada terminal
func (a Adapter) GetjsonTerminals(ctx context.Context) ([]byte, error) {
	url := os.Getenv("Group")

	resp, err := a.http.Get(ctx, url, krest.RequestData{
		Headers: map[string]string{
			"Referer":    "https://onibus.info/mapa/",
			"Connection": "keep-alive",
			"Host":       "onibus.info",
		},
	})
	if err != nil {
		if resp.StatusCode == 404 {
			return nil, fmt.Errorf("onibus.info/api/routes/group/ was not found! %s", err)
		}
		return nil, fmt.Errorf("unexpected error when fetching terminals: %s", err)
	}

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ = gzip.NewReader(resp)
		defer reader.Close()
	default:
		reader = resp
	}

	term, _ := ioutil.ReadAll(reader)
	return term, nil
}
