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

const verifyURI = "/v1/verify/sms"

// VerifyRequestData data
type VerifyRequestData struct {
	OriginatingIP string `schema:"originating_ip"`
	PhoneNumber   string `schema:"phone_number"`
	Ucid          string `schema:"ucid"`
	Language      string `schema:"language"`
	VerifyCode    string `schema:"verify_code"`
	Template      string `schema:"template"`
}

// VerifyRequest object
type VerifyRequest struct {
	uri  string
	data VerifyRequestData
	body string
}

// GetMethod return method request
func (r VerifyRequest) GetMethod() string {
	return http.MethodPost
}

// GetURI return uri request
func (r VerifyRequest) GetURI() string {
	return r.GetPath()
}

// GetPath return path request
func (r VerifyRequest) GetPath() string {
	return r.uri
}

// GetBody return body request
func (r *VerifyRequest) GetBody() string {
	if r.body == "" {
		r.body = telesign.StructToURLValues(r.data).Encode()
		r.body = strings.ReplaceAll(r.body, "%2520", "%20")
	}

	return r.body
}

// ParseResponse return parsed response
func (r VerifyRequest) ParseResponse(statusCode int, content []byte) (telesign.Response, error) {
	var resp VerifyResponse

	err := json.Unmarshal(content, &resp)
	if err != nil {
		return nil, err
	}

	resp.StatusCode = statusCode

	return resp, nil
}

// VerifyResponse returned by telesign API
type VerifyResponse struct {
	telesign.MainResponse
	SubResource     string             `json:"sub_resource"`
	Errors          []telesign.Error   `json:"errors"`
	Verify          VerifyDataResponse `json:"verify"`
	ExternalID      string             `json:"external_id"`
	SignatureString string             `json:"signature_string"`
}

// VerifyDataResponse returned by telesign API
type VerifyDataResponse struct {
	CodeState   string `json:"code_state"`
	CodeEntered string `json:"code_entered"`
}

// NewVerify return new Verify request
func NewVerify(ip string, phone string, ucid string, lang string, code string,
	template string) telesign.Request {
	// convert the spaces into "%20" first
	// later we replace all "%2520" into "%20"
	t := url.URL{Path: template}
	template = t.String()

	return &VerifyRequest{
		uri: verifyURI,
		data: VerifyRequestData{
			OriginatingIP: ip,
			PhoneNumber:   phone,
			Ucid:          ucid,
			Language:      lang,
			VerifyCode:    code,
			Template:      template,
		},
	}
}
