package lastfm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Artist struct {
	ArtistName string `json:"text"`
}

type Track struct {
	Artist    `json:"artist,string"`
	Name      string `json:"name"`
	Playcount int    `json:"playcount,string"`
}

type _WeeklyTrackChart struct {
	Tracks []Track `json:"track,string"`
}

type WeeklyTrackChart struct {
	_WeeklyTrackChart `json:"weeklytrackchart,string"`
}

const ENDPOINT_BASE = "http://ws.audioscrobbler.com/2.0/?method=user.getweeklytrackchart"

func GetTracks(user string, apikey string, from int64, to int64, max int) ([]Track, error) {
	if apikey == "" {
		return nil, fmt.Errorf("Last.FM API Key is required")
	}
	endpoint := fmt.Sprintf("%s&user=%s&api_key=%s&format=json&from=%d&to=%d&limit=%d",
		ENDPOINT_BASE, user, apikey, from, to, max)

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	invalid, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	valid := strings.ReplaceAll(string(invalid), "#text", "text")

	weekly := WeeklyTrackChart{}
	err = json.Unmarshal([]byte(valid), &weekly)
	if err != nil {
		return nil, err
	}

	return weekly.Tracks, nil
}
