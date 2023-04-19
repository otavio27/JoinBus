package onibus

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/otavio27/JoinBus-APP/back-end/structs"
	"github.com/vingarcia/ddd-go-template/v2-domain-adapters-and-helpers/domain"
	"github.com/vingarcia/krest"
)

type Adapter struct {
	module        string
	http          krest.Provider
	ctx           context.Context
	stopsnear     string
	stoptripslist string
	timetable     string
	group         string
	referer       string
	host          string
}

func New(http krest.Provider, ctx context.Context,
	stopsnear string, stoptripslist string, timetable string,
	group string, referer string, host string) *Adapter {
	return &Adapter{
		module:        "Adapter",
		http:          http,
		ctx:           ctx,
		stopsnear:     stopsnear,
		stoptripslist: stoptripslist,
		timetable:     timetable,
		group:         group,
		referer:       referer,
		host:          host,
	}
}

func (a Adapter) GetGeoLocation(latitude string, longitude string) ([]structs.StopList, []string, error) {
	resp, err := a.http.Get(a.ctx, a.stopsnear+latitude+"&lng="+longitude, krest.RequestData{
		Headers: map[string]string{
			"Referer":    a.referer,
			"Connection": "keep-alive",
			"Host":       a.host,
		},
	})
	if err != nil {
		if resp.StatusCode == 404 {
			return nil, nil, domain.NotFoundErr("geolocation not found", map[string]any{
				"latitude":  latitude,
				"longitude": longitude,
			})
		}
		return nil, nil, domain.InternalErr("unexpected error when fetching geolocation", map[string]any{
			"error": err.Error(),
		})
	}

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err := gzip.NewReader(resp)
		if err != nil {
			return nil, nil, domain.InternalErr("unexpected error unzipping geolocation from external api", map[string]any{
				"error": err.Error(),
			})
		}
		defer reader.Close()
	default:
		reader = resp
	}

	location, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, nil, domain.InternalErr("unexpected error unzipping geolocation from external api", map[string]any{
			"error": err.Error(),
		})
	}

	var local []structs.Location
	err = json.Unmarshal(location, &local)
	if err != nil {
		return nil, nil, domain.InternalErr("error parsing geolocation as json", map[string]any{
			"error":   err.Error(),
			"payload": string(location),
		})
	}

	var stop_id, stop_name []string
	for _, lctn := range local {
		stop_id = append(stop_id, lctn.StopID)
		stop_name = append(stop_name, lctn.StopName)
	}

	return a.GetStopTripList(a.ctx, stop_id, stop_name)
}

func (a Adapter) GetStopTripList(ctx context.Context, stop []string, stopName []string) ([]structs.StopList, []string, error) {
	var stoplist []byte

	if len(stop) == 0 {
		return nil, nil, domain.NotFoundErr("stop variable is empty", map[string]any{
			"stop": stop,
		})
	}

	for _, std := range stop[:2] {
		resp, err := a.http.Get(ctx, a.stoptripslist+std, krest.RequestData{
			Headers: map[string]string{
				"Referer":    a.referer,
				"Connection": "keep-alive",
				"Host":       a.host,
			},
		})
		if err != nil {
			if resp.StatusCode == 404 {
				return nil, nil, domain.NotFoundErr("stoplists not found", map[string]any{
					"stop_id": std,
				})
			}
			return nil, nil, domain.InternalErr("unexpected error when fetching stoplists", map[string]any{
				"error": err.Error(),
			})
		}

		var reader io.ReadCloser
		switch resp.Header.Get("Content-Encoding") {
		case "gzip":
			reader, err := gzip.NewReader(resp)
			if err != nil {
				return nil, nil, domain.InternalErr("unexpected error unzipping stoplists from external api", map[string]any{
					"error": err.Error(),
				})
			}
			defer reader.Close()
		default:
			reader = resp
		}

		stoplist, err = ioutil.ReadAll(reader)
		if err != nil {
			return nil, nil, domain.InternalErr("unexpected error reading stoplist from external api", map[string]any{
				"error": err.Error(),
			})
		}

	}

	var stoplists []structs.StopList
	err := json.Unmarshal(stoplist, &stoplists)
	if err != nil {
		return nil, nil, domain.InternalErr("error parsing stoplists as json", map[string]any{
			"error":   err.Error(),
			"payload": string(stoplist),
		})
	}

	return stoplists, stopName, nil
}

func (a Adapter) GetItineraries(ctx context.Context, id string) ([]structs.Itinerary, error) {
	resp, err := a.http.Get(ctx, a.timetable+id, krest.RequestData{
		Headers: map[string]string{
			"Referer":    a.referer,
			"Connection": "keep-alive",
			"Host":       a.host,
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

func (a Adapter) GetjsonTerminals(ctx context.Context) ([]structs.Stations, error) {
	resp, err := a.http.Get(ctx, a.group, krest.RequestData{
		Headers: map[string]string{
			"Referer":    a.referer,
			"Connection": "keep-alive",
			"Host":       a.host,
		},
	})

	if err != nil {
		if resp.StatusCode == 404 {
			return nil, domain.NotFoundErr("terminals not found", map[string]any{
				"error": err.Error(),
			})
		}
		return nil, domain.InternalErr("unexpected error when fetching terminals", map[string]any{
			"error": err.Error(),
		})
	}

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err := gzip.NewReader(resp)
		if err != nil {
			return nil, domain.InternalErr("unexpected error reading terminals from external api", map[string]any{
				"error": err.Error(),
			})
		}
		defer reader.Close()
	default:
		reader = resp
	}

	term, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, domain.InternalErr("unexpected error reading terminals from external api", map[string]any{
			"error": err.Error(),
		})
	}

	var Stations []structs.Stations
	err = json.Unmarshal(term, &Stations)
	if err != nil {
		return nil, domain.InternalErr("error parsing terminals as json", map[string]any{
			"error":   err.Error(),
			"payload": string(term),
		})
	}

	return Stations, nil
}
