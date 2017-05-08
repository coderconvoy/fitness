package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

const (
	geokey = "AIzaSyCtwU6GvHv9yY543PeqFDUbeR6kg09pUTw"
)

type LatLong struct {
	Lat float32 `json:"lat"`
	Lng float32 `json:"lng"`
}

type Location struct {
	S     string
	xy    LatLong
	isset bool
}

func (l *Location) Load() error {
	if l.isset {
		return nil
	}

	//todo fix to remove "&"from String l.S
	url := "https://maps.googleapis.com/maps/api/geocode/json"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return errors.Wrap(err, "Error Making Log:")
	}

	q := req.URL.Query()
	q.Add("key", geokey)
	q.Add("address", l.S)
	fmt.Println(req.URL.String())
	req.URL.RawQuery = q.Encode()

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "Error asking google for location:")
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(b))
	var geo GeoResp
	err = json.Unmarshal(b, &geo)
	fmt.Println(geo)
	if err != nil {
		return errors.Wrap(err, "Json marshal failed")
	}

	l.xy = geo.Results[0].Geometry.Location

	l.isset = true
	return nil
}

type GeoMet struct {
	Location LatLong `json:location`
}
type GeoRes struct {
	Geometry GeoMet `json:"geometry"`
}
type GeoResp struct {
	Results []GeoRes `json:"results"`
}
