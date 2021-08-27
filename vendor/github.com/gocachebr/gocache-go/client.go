package gocache

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type API struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
}

func NewClient(token string) (*API, error) {
	api := API{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    "https://api.gocache.com.br/v1",
		Token:      token,
	}

	return &api, nil
}

func (api *API) doRequest(req *http.Request) ([]byte, int, error) {
	req.Header.Set("GoCache-Token", api.Token)

	res, err := api.HTTPClient.Do(req)
	if err != nil {
		return nil, -1, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, res.StatusCode, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, res.StatusCode, err
}
