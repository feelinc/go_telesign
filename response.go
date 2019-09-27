// Copyright (c) 2019 Sulaeman (me@sulaeman.com), All rights reserved.
// This source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package telesign

import "net/http"

// Response interface
type Response interface {
	GetStatusCode() int
	Failure() bool
	Message() string
}

// MainResponse returned by telesign API
type MainResponse struct {
	StatusCode  int
	ResourceURI string         `json:"resource_uri"`
	ReferenceID string         `json:"reference_id"`
	Status      StatusResponse `json:"status"`
}

// StatusResponse returned by telesign API
type StatusResponse struct {
	Code        int    `json:"code"`
	UpdatedOn   string `json:"updated_on"`
	Description string `json:"description"`
}

// Failure return true if failure, otherwise false
func (r MainResponse) Failure() bool {
	return r.StatusCode != http.StatusOK && r.StatusCode != http.StatusNoContent
}

// GetStatusCode return response HTTP status code
func (r MainResponse) GetStatusCode() int {
	return r.StatusCode
}

// Message return status description
func (r MainResponse) Message() string {
	return r.Status.Description
}
