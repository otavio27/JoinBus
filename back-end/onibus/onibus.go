package onibus

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/otavio27/JoinBus-APP/back-end/structs"
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
func (a Adapter) GetGeoLocation(latitude float64, longitude float64) ([]byte, []string, error) {
	latd := fmt.Sprintf("%f", latitude)
	logd := fmt.Sprintf("%f", longitude)
	url := "https://onibus.info/api/stopsnear?lat=" + latd + "&lng=" + logd
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
		return nil, nil, fmt.Errorf("unexpected error when fetching example.com: %s", err)
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
				return nil, nil, fmt.Errorf("onibus.info/api/stopsnear?lat= was not found! %s", err)
			}
			return nil, nil, fmt.Errorf("unexpected error when fetching example.com: %s", err)
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
func (a Adapter) GetjsonLines(ctx context.Context, text string) ([]byte, error) {
	url := "https://onibus.info/api/timetable/" + text

	resp, err := a.http.Get(ctx, url, krest.RequestData{
		Headers: map[string]string{
			"Referer":    "https://onibus.info/mapa/",
			"Connection": "keep-alive",
			"Host":       "onibus.info",
		},
	})
	if err != nil {
		if resp.StatusCode == 404 {
			return nil, fmt.Errorf("onibus.info/api/timetable/ was not found! %s", err)
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

	hours, _ := ioutil.ReadAll(reader)
	return hours, nil
}

// GetjsonTerminals busca todas as linhas de cada terminal
func (a Adapter) GetjsonTerminals(ctx context.Context) ([]byte, error) {
	url := "https://onibus.info/api/routes/group"

	resp, err := a.http.Get(ctx, url, krest.RequestData{
		Headers: map[string]string{
			"Referer":    "https://onibus.info/mapa/",
			"Connection": "keep-alive",
			"Host":       "onibus.info",
		},
	})
	if err != nil {
		if resp.StatusCode == 404 {
			return nil, fmt.Errorf("onibus.info/api/timetable/ was not found! %s", err)
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

	term, _ := ioutil.ReadAll(reader)
	return term, nil
}
