// Copyright (c) 2019 Sulaeman (me@sulaeman.com), All rights reserved.
// This source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package phoneid

import (
	"encoding/json"
	"net/http"
	"path"

	telesign "github.com/feelinc/go_telesign"
)

const liveURI = "/v1/phoneid/live"

// LiveRequest object
type LiveRequest struct {
	uri         string
	phoneNumber string
	ucid        string
}

// GetMethod return method request
func (r LiveRequest) GetMethod() string {
	return http.MethodGet
}

// GetURI return uri request
func (r LiveRequest) GetURI() string {
	return r.GetPath() + "?ucid=" + r.ucid
}

// GetPath return path request
func (r LiveRequest) GetPath() string {
	return path.Join(r.uri, r.phoneNumber)
}

// GetBody return body request
func (r *LiveRequest) GetBody() string {
	return ""
}

// ParseResponse return parsed response
func (r LiveRequest) ParseResponse(statusCode int, content []byte) (telesign.Response, error) {
	var resp Response

	err := json.Unmarshal(content, &resp)
	if err != nil {
		return nil, err
	}

	resp.StatusCode = statusCode

	return resp, nil
}

// LiveResponse returned by telesign API
type LiveResponse struct {
	SubscriberStatus   string `json:"subscriber_status"`
	DeviceStatus       string `json:"device_status"`
	Roaming            string `json:"roaming"`
	RoamingCountry     string `json:"roaming_country"`
	RoamingCountryIso2 string `json:"roaming_country_iso2"`
}

// NewLive return new Live request
func NewLive(phone string, ucid string) telesign.Request {
	return &LiveRequest{
		uri:         liveURI,
		phoneNumber: phone,
		ucid:        ucid,
	}
}
