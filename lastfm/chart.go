package lastfm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type _WeeklyTrackChart struct {
	Tracks []Track `json:"track,string"`
}

type WeeklyTrackChart struct {
	_WeeklyTrackChart `json:"weeklytrackchart,string"`
	Message           string `json:"message"`
	Error             int    `json:"error,int"`
}

func (w *WeeklyTrackChart) Fetch(user string, apikey string, from int64, to int64, max int) error {
	endpoint := fmt.Sprintf("%s?method=user.getweeklytrackchart&user=%s&api_key=%s&format=json&from=%d&to=%d&limit=%d",
		ENDPOINT_BASE, user, apikey, from, to, max)

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	invalid, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// `#` is an invalid character at json key
	valid := strings.ReplaceAll(string(invalid), "#text", "text")

	if err = json.Unmarshal([]byte(valid), &w); err != nil {
		return err
	} else if w.Error > 0 {
		return fmt.Errorf("%s (error code = %d)\n", w.Message, w.Error)
	}

	return nil
}
