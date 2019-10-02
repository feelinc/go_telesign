// Copyright (c) 2019 Sulaeman (me@sulaeman.com), All rights reserved.
// This source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package voice

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	telesign "github.com/feelinc/go_telesign"
)

const uri = "/v1/verify/call"

// VerifyRequestData data
type VerifyRequestData struct {
	OriginatingIP     string `schema:"originating_ip"`
	PhoneNumber       string `schema:"phone_number"`
	Ucid              string `schema:"ucid"`
	Language          string `schema:"language"`
	VerifyCode        string `schema:"verify_code"`
	TTSMessage        string `schema:"tts_message"`
	CallForwardAction string `schema:"call_forward_action"`
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
	SubResource     string                  `json:"sub_resource"`
	Errors          []telesign.Error        `json:"errors"`
	Verify          VerifyDataResponse      `json:"verify"`
	Voice           VerifyVoiceResponse     `json:"voice"`
	PhoneType       VerifyPhoneTypeResponse `json:"phone_type"`
	Risk            VerifyRiskResponse      `json:"risk"`
	Numbering       VerifyNumberingResponse `json:"numbering"`
	ExternalID      string                  `json:"external_id"`
	SignatureString string                  `json:"signature_string"`
}

// VerifyDataResponse returned by telesign API
type VerifyDataResponse struct {
	CodeState   string `json:"code_state"`
	CodeEntered string `json:"code_entered"`
}

// VerifyVoiceResponse returned by telesign API
type VerifyVoiceResponse struct {
	CallerID string `json:"caller_id"`
}

// VerifyPhoneTypeResponse returned by telesign API
type VerifyPhoneTypeResponse struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

// VerifyRiskResponse returned by telesign API
type VerifyRiskResponse struct {
	Score          int64  `json:"score"`
	Recommendation string `json:"recommendation"`
	Level          string `json:"level"`
}

// VerifyNumberingResponse returned by telesign API
type VerifyNumberingResponse struct {
	PhoneNumber string `json:"phone_number"`
	MinLength   int    `json:"min_length"`
	MaxLength   int    `json:"max_length"`
	CountryCode string `json:"country_code"`
}

// NewVerify return new Verify request
func NewVerify(ip string, phone string, ucid string, lang string, code string,
	msg string, callForwardAction string) telesign.Request {
	// convert the spaces into "%20" first
	// later we replace all "%2520" into "%20"
	t := url.URL{Path: msg}
	msg = t.String()

	return &VerifyRequest{
		uri: uri,
		data: VerifyRequestData{
			OriginatingIP:     ip,
			PhoneNumber:       phone,
			Ucid:              ucid,
			Language:          lang,
			VerifyCode:        code,
			TTSMessage:        msg,
			CallForwardAction: callForwardAction,
		},
	}
}
