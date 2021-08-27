package gocache

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (api *API) ListDomains() (*API_Response, error) {
	resp := API_Response{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/domain", api.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, status, err := api.doRequest(req)
	resp.HTTPStatusCode = status
	if err != nil {

		jsonErr := json.Unmarshal(body, &resp)

		if jsonErr != nil {
			return &resp, err
		}

	} else {

		resp.Response = new(DomainList)

		err = json.Unmarshal(body, &resp)
		if err != nil {
			return nil, err
		}

		resp.Response = *resp.Response.(*DomainList)

	}

	return &resp, nil
}

func (api *API) GetDomain(domain string) (*API_Response, error) {
	resp := API_Response{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/domain/%s", api.HostURL, domain), nil)
	if err != nil {
		return nil, err
	}

	body, status, err := api.doRequest(req)
	resp.HTTPStatusCode = status
	if err != nil {

		jsonErr := json.Unmarshal(body, &resp)

		if jsonErr != nil {
			return &resp, err
		}

	} else {

		err = json.Unmarshal(body, &resp)
		if err != nil {
			return nil, err
		}

		resp.Response = responseConvert(resp.Response.(map[string]interface{}), domainConvert)

	}

	return &resp, nil
}

func (api *API) CreateDomain(domain string, domainInfo map[string]interface{}) (*API_Response, error) {
	resp := API_Response{}
	reqBody := formData(domainInfo, domainConvert)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/domain/%s", api.HostURL, domain), strings.NewReader(reqBody))
	if err != nil {
		return &resp, err
	}

	body, status, err := api.doRequest(req)
	resp.HTTPStatusCode = status
	if err != nil {

		jsonErr := json.Unmarshal(body, &resp)

		if jsonErr != nil {
			return &resp, err
		}

	} else {

		resp.Response = new(DomainList)

		err = json.Unmarshal(body, &resp)
		if err != nil {
			return &resp, err
		}

		resp.Response = *resp.Response.(*DomainList)

	}

	return &resp, nil
}

func (api *API) UpdateDomain(domain string, domainInfo map[string]interface{}) (*API_Response, error) {
	resp := API_Response{}

	reqBody := formData(domainInfo, domainConvert)

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/domain/%s", api.HostURL, domain), strings.NewReader(reqBody))
	if err != nil {
		return nil, err
	}

	body, status, err := api.doRequest(req)
	resp.HTTPStatusCode = status
	if err != nil {

		jsonErr := json.Unmarshal(body, &resp)

		if jsonErr != nil {
			return &resp, err
		}

	} else {

		err = json.Unmarshal(body, &resp)
		if err != nil {
			return nil, err
		}

		resp.Response = responseConvert(resp.Response.(map[string]interface{}), domainConvert)

	}

	return &resp, nil
}

func (api *API) DeleteDomain(domain string) (*API_Response, error) {
	resp := API_Response{}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/domain/%s", api.HostURL, domain), nil)
	if err != nil {
		return nil, err
	}

	body, status, err := api.doRequest(req)
	resp.HTTPStatusCode = status

	jsonErr := json.Unmarshal(body, &resp)
	if jsonErr != nil {
		return &resp, err
	}

	return &resp, nil
}
