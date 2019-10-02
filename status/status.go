// Copyright (c) 2019 Sulaeman (me@sulaeman.com), All rights reserved.
// This source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package status

import (
	"encoding/json"
	"net/http"
	"path"

	telesign "github.com/feelinc/go_telesign"
)

const uri = "/v1/verify"

// NewClient return new Status API connection
func NewClient(options ...telesign.OptConFunc) telesign.Connection {
	return telesign.NewStatus(options...)
}

// Request object
type Request struct {
	uri  string
	id   string
	code string
}

// GetMethod return method request
func (r Request) GetMethod() string {
	return http.MethodGet
}

// GetURI return uri request
func (r Request) GetURI() string {
	if r.code != "" {
		return r.GetPath() + "?verify_code=" + r.code
	}
	return r.GetPath()
}

// GetPath return path request
func (r Request) GetPath() string {
	return path.Join(r.uri, r.id)
}

// GetBody return body request
func (r *Request) GetBody() string {
	return ""
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
	Errors                   []telesign.Error                 `json:"errors"`
	Verify                   VerifyResponse                   `json:"verify"`
	NumberDeactivationStatus NumberDeactivationStatusResponse `json:"number_deactivation_status"`
	AdditionalInfo           telesign.AdditionalInfo          `json:"additional_info"`
}

// VerifyResponse returned by telesign API
type VerifyResponse struct {
	CodeState   string `json:"code_state"`
	CodeEntered string `json:"code_entered"`
}

// NumberDeactivationStatusResponse returned by telesign API
type NumberDeactivationStatusResponse struct {
	ErrorCode       string `json:"error_code"`
	Description     string `json:"description"`
	LastDeactivated string `json:"last_deactivated"`
}

// New return new request
func New(id string, code string) telesign.Request {
	return &Request{
		uri:  uri,
		id:   id,
		code: code,
	}
}
