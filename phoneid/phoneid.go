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

const uri = "/v1/phoneid"

// NewClient return new PhoneID API connection
func NewClient(options ...telesign.OptConFunc) telesign.Connection {
	return telesign.NewPhoneID(options...)
}

// RequestData data
type RequestData struct {
	OriginatingIP         string `schema:"originating_ip"`
	PhoneNumber           string `schema:"-"`
	AccountLifecycleEvent string `schema:"account_lifecycle_event"`
}

// Request object
type Request struct {
	uri  string
	data RequestData
	body string
}

// GetMethod return method request
func (r Request) GetMethod() string {
	return http.MethodPost
}

// GetURI return uri request
func (r Request) GetURI() string {
	return r.GetPath()
}

// GetPath return path request
func (r Request) GetPath() string {
	return path.Join(r.uri, r.data.PhoneNumber)
}

// GetBody return body request
func (r *Request) GetBody() string {
	if r.body == "" {
		r.body = telesign.StructToURLValues(r.data).Encode()
	}

	return r.body
}

// ParseResponse return parsed response
func (r Request) ParseResponse(statusCode int, content []byte) (telesign.Response, error) {
	var resp Response

	err := json.Unmarshal(content, &resp)
	if err != nil {
		return nil, err
	}

	resp.StatusCode = statusCode

	return resp, nil
}

// Response returned by telesign API
type Response struct {
	telesign.MainResponse
	Errors       []telesign.Error     `json:"errors"`
	Location     LocationResponse     `json:"location"`
	Numbering    NumberingResponse    `json:"numbering"`
	PhoneType    PhoneTypeResponse    `json:"phone_type"`
	Blocklisting BlocklistingResponse `json:"blocklisting"`
	Carrier      CarrierResponse      `json:"carrier"`
	Live         LiveResponse         `json:"live"`
}

// LocationResponse returned by telesign API
type LocationResponse struct {
	City        string              `json:"city"`
	State       string              `json:"state"`
	Zip         string              `json:"zip"`
	MetroCode   string              `json:"metro_code"`
	County      string              `json:"county"`
	Country     CountryResponse     `json:"country"`
	Coordinates CoordinatesResponse `json:"coordinates"`
	Timezone    TimezoneResponse    `json:"time_zone"`
}

// CountryResponse returned by telesign API
type CountryResponse struct {
	Name string `json:"name"`
	Iso2 string `json:"iso2"`
	Iso3 string `json:"iso3"`
}

// CoordinatesResponse returned by telesign API
type CoordinatesResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// TimezoneResponse returned by telesign API
type TimezoneResponse struct {
	Name         string `json:"name"`
	UTCOffsetMin string `json:"utc_offset_min"`
	UTCOffsetMax string `json:"utc_offset_max"`
}

// NumberingResponse returned by telesign API
type NumberingResponse struct {
	Original  NumberingOriginalResponse  `json:"original"`
	Cleansing NumberingCleansingResponse `json:"cleansing"`
}

// NumberingOriginalResponse returned by telesign API
type NumberingOriginalResponse struct {
	CompletePhoneNumber string `json:"complete_phone_number"`
	CountryCode         string `json:"country_code"`
	PhoneNumber         string `json:"phone_number"`
}

// NumberingCleansingResponse returned by telesign API
type NumberingCleansingResponse struct {
	Call NumberingCleansingItemResponse `json:"call"`
	SMS  NumberingCleansingItemResponse `json:"sms"`
}

// NumberingCleansingItemResponse returned by telesign API
type NumberingCleansingItemResponse struct {
	CountryCode  string `json:"country_code"`
	PhoneNumber  string `json:"phone_number"`
	CleansedCode int64  `json:"cleansed_code"`
	MinLength    int64  `json:"min_length"`
	MaxLength    int64  `json:"max_length"`
}

// PhoneTypeResponse returned by telesign API
type PhoneTypeResponse struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

// BlocklistingResponse returned by telesign API
type BlocklistingResponse struct {
	BlockCode        int64  `json:"block_code"`
	BlockCescription string `json:"block_description"`
	Blocked          bool   `json:"blocked"`
}

// CarrierResponse returned by telesign API
type CarrierResponse struct {
	Name string `json:"name"`
}

// New return new SMS request
func New(ip string, phone string, accLifecycleEvent string) telesign.Request {
	return &Request{
		uri: uri,
		data: RequestData{
			OriginatingIP:         ip,
			PhoneNumber:           phone,
			AccountLifecycleEvent: accLifecycleEvent,
		},
	}
}
