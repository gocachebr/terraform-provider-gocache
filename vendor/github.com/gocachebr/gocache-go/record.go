package gocache

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (api *API) ListRecords(domain string) (*API_Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/dns/%s", api.HostURL, domain), nil)
	if err != nil {
		return nil, err
	}

	body, status, err := api.doRequest(req)
	if err != nil {
		return nil, err
	}

	resp := API_Response{}

	resp.Response = new(DNSResult)

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	resp.Response = *resp.Response.(*DNSResult)

	for i, r := range resp.Response.(DNSResult).Records {
		resp.Response.(DNSResult).Records[i] = responseConvert(r, recordConvert)
		proxied, ok := r["proxied"]
		if ok {
			if proxied == "1" {
				r["proxied"] = true
			} else {
				r["proxied"] = false
			}
		}
	}

	resp.HTTPStatusCode = status

	return &resp, nil
}

func (api *API) CreateRecord(domain string, recordInfo map[string]interface{}) (*API_Response, error) {
	resp := API_Response{}
	proxied, ok := recordInfo["proxied"]
	if ok {
		if proxied == true {
			recordInfo["proxied"] = "1"
		} else {
			recordInfo["proxied"] = "0"
		}
	}
	reqBody := formData(recordInfo, recordConvert)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/dns/%s", api.HostURL, domain), strings.NewReader(reqBody))
	if err != nil {
		return nil, err
	}

	body, status, err := api.doRequest(req)
	resp.HTTPStatusCode = status

	if err != nil {
		return &resp, err
	}

	resp.Response = new(DNSResult)

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return &resp, err
	}

	resp.Response = *resp.Response.(*DNSResult)

	for i, r := range resp.Response.(DNSResult).Records {
		resp.Response.(DNSResult).Records[i] = responseConvert(r, recordConvert)
		proxied, ok = r["proxied"]
		if ok {
			if proxied == "1" {
				r["proxied"] = true
			} else {
				r["proxied"] = false
			}
		}
		id, ok := r["record_id"]
		if ok {
			resp.Response.(DNSResult).Records[i]["record_id"] = fmt.Sprintf("%.0f", id)
		}
	}

	return &resp, nil
}

func (api *API) UpdateRecord(domain string, recordInfo map[string]interface{}) (*API_Response, error) {
	proxied, ok := recordInfo["proxied"]
	if ok {
		if proxied == true {
			recordInfo["proxied"] = "1"
		} else {
			recordInfo["proxied"] = "0"
		}
	}
	reqBody := formData(recordInfo, recordConvert)

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/dns/%s", api.HostURL, domain), strings.NewReader(reqBody))
	if err != nil {
		return nil, err
	}

	body, status, err := api.doRequest(req)
	if err != nil {
		return nil, err
	}

	resp := API_Response{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	resp.HTTPStatusCode = status

	return &resp, nil
}

func (api *API) DeleteRecord(domain string, recordId string) (*API_Response, error) {
	reqBody := fmt.Sprintf("record_id=%s", recordId)

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/dns/%s", api.HostURL, domain), strings.NewReader(reqBody))
	if err != nil {
		return nil, err
	}

	body, status, err := api.doRequest(req)
	if err != nil {
		return nil, err
	}

	resp := API_Response{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	resp.HTTPStatusCode = status

	return &resp, nil
}
