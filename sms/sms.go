// Copyright (c) 2019 Sulaeman (me@sulaeman.com), All rights reserved.
// This source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package sms

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	telesign "github.com/feelinc/go_telesign"
)

const smsURI = "/v1/messaging"

// NewClient return new SMS API connection
func NewClient(options ...telesign.OptConFunc) telesign.Connection {
	return telesign.NewSMS(options...)
}

// RequestData data
type RequestData struct {
	OriginatingIP string `schema:"originating_ip"`
	PhoneNumber   string `schema:"phone_number"`
	Message       string `schema:"message"`
	MessageType   string `schema:"message_type"`
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
	return r.uri
}

// GetBody return body request
func (r *Request) GetBody() string {
	if r.body == "" {
		r.body = telesign.StructToURLValues(r.data).Encode()
		r.body = strings.ReplaceAll(r.body, "%2520", "%20")
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
	AdditionalInfo AdditionalInfo `json:"additional_info"`
}

// AdditionalInfo returned by telesign API
type AdditionalInfo struct {
	CodeEntered       string `json:"code_entered"`
	MessagePartsCount int    `json:"message_parts_count"`
}

// New return new SMS request
func New(ip string, phone string, msg string, typ string) telesign.Request {
	// convert the spaces into "%20" first
	// later we replace all "%2520" into "%20"
	t := url.URL{Path: msg}
	msg = t.String()

	return &Request{
		uri: smsURI,
		data: RequestData{
			OriginatingIP: ip,
			PhoneNumber:   phone,
			Message:       msg,
			MessageType:   typ,
		},
	}
}
