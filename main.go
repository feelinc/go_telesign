// Copyright (c) 2018 Sulaeman (me@sulaeman.com), All rights reserved.
// This source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package telesign

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"runtime"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/schema"
)

var userAgent = fmt.Sprintf("TeleSignSDK/go-%s Go/%s", version, runtime.Version())

// OptFunc is a function that configures a Telesign
type OptFunc func(*Telesign) error

// SetEnv set the environment
func SetEnv(env string) OptFunc {
	return func(t *Telesign) error {
		t.env = env
		return nil
	}
}

// SetCustomerID set the customer ID
func SetCustomerID(id string) OptFunc {
	return func(t *Telesign) error {
		t.customerID = id
		return nil
	}
}

// SetAPIKey set the API Key
func SetAPIKey(key string) OptFunc {
	return func(t *Telesign) error {
		t.apiKey = key
		return nil
	}
}

// Telesign object
type Telesign struct {
	env        string
	customerID string
	apiKey     string
	phoneID    Connection
	score      Connection
	sms        Connection
	voice      Connection
}

// PhoneID return PhoneID service
func (t *Telesign) PhoneID(options ...OptConFunc) Connection {
	if t.phoneID == nil {
		if len(options) > 0 {
			t.phoneID = NewPhoneID(options...)
		} else {
			t.phoneID = NewPhoneID(SetConEnv(t.env), SetConCustomerID(t.customerID),
				SetConAPIKey(t.apiKey))
		}
	}
	return t.phoneID
}

// Score return Score service
func (t *Telesign) Score(options ...OptConFunc) Connection {
	if t.score == nil {
		if len(options) > 0 {
			t.score = NewScore(options...)
		} else {
			t.score = NewScore(SetConEnv(t.env), SetConCustomerID(t.customerID),
				SetConAPIKey(t.apiKey))
		}
	}
	return t.score
}

// SMS return SMS service
func (t *Telesign) SMS(options ...OptConFunc) Connection {
	if t.sms == nil {
		if len(options) > 0 {
			t.sms = NewSMS(options...)
		} else {
			t.sms = NewSMS(SetConEnv(t.env), SetConCustomerID(t.customerID),
				SetConAPIKey(t.apiKey))
		}
	}
	return t.sms
}

// Voice return Voice service
func (t *Telesign) Voice(options ...OptConFunc) Connection {
	if t.voice == nil {
		if len(options) > 0 {
			t.voice = NewVoice(options...)
		} else {
			t.voice = NewVoice(SetConEnv(t.env), SetConCustomerID(t.customerID),
				SetConAPIKey(t.apiKey))
		}
	}
	return t.voice
}

// New return a new Telesign
func New(options ...OptFunc) *Telesign {
	t := &Telesign{}

	// Run the options on it
	for _, option := range options {
		if err := option(t); err != nil {
			return nil
		}
	}

	return t
}

type responseParser func(statusCode int, content []byte) (Response, error)

// Execute execute the request return a response
func Execute(customerID string, apiKey string, httpTimeout int,
	req *http.Request, resource string, body string, resParser responseParser) (Response, error) {
	nonce, _ := uuid.NewRandom()
	sigData := buildSignature(time.Now(), nonce.String(), req.Method, resource,
		req.Header.Get("Content-Type"), body)
	sig := fmt.Sprintf("TSA %s:%s", customerID, createSignature(apiKey, sigData))

	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("X-TS-Auth-Method", authMethod)
	req.Header.Add("Authorization", sig)
	req.Header.Add("X-TS-Nonce", sigData.Nonce)
	req.Header.Add("Date", sigData.Date)

	client := getClient(httpTimeout)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	respBody, _ := ioutil.ReadAll(resp.Body)

	// fmt.Println(string(respBody))

	return resParser(resp.StatusCode, respBody)
}

// StructToURLValues convert struct data to URL Values
func StructToURLValues(i interface{}) url.Values {
	val := url.Values{}
	err := schema.NewEncoder().Encode(i, val)
	if err != nil {
		panic(err)
	}
	return val
}

func execute(conf Config, req Request) (Response, error) {
	uri := buildRequestURI(getDomain(conf.Env), req.GetURI())

	httpReq, err := buildRequest(req.GetMethod(), uri,
		bytes.NewBuffer([]byte(req.GetBody())))
	if err != nil {
		return nil, err
	}

	return Execute(conf.CustomerID, conf.APIKey, conf.HTTPTimeout, httpReq,
		req.GetPath(), req.GetBody(), func() responseParser {
			return func(statusCode int, content []byte) (Response, error) {
				return req.ParseResponse(statusCode, content)
			}
		}())
}

func buildRequestURI(domain string, uri string) string {
	return fmt.Sprintf("%s%s", domain, uri)
}

func buildRequest(method string, uri string, body *bytes.Buffer) (*http.Request, error) {
	httpReq, err := http.NewRequest(method, uri, body)
	if err != nil {
		return nil, err
	}

	if method == http.MethodPost {
		httpReq.Header.Add("Content-Type",
			"application/x-www-form-urlencoded; charset=utf-8")
	}

	return httpReq, nil
}

func getDomain(env string) string {
	if env == EnvEnterprise {
		return domainEnterprise
	}
	return domain
}

type signatureData struct {
	HTTPMethod  string
	Resource    string
	ContentType string
	Date        string
	Nonce       string
	Body        string
}

func buildSignature(t time.Time, nonce string, method string, resource string,
	contentType string, body string) signatureData {
	return signatureData{
		HTTPMethod:  method,
		Resource:    resource,
		ContentType: contentType,
		Date:        t.Format(timeFormat),
		Nonce:       nonce,
		Body:        body,
	}
}

func createSignature(apiKey string, data signatureData) string {
	var str string
	switch data.HTTPMethod {
	case http.MethodGet, http.MethodDelete:
		str = fmt.Sprintf("%s\n%s\n%s\nx-ts-auth-method:%s\nx-ts-nonce:%s\n%s",
			data.HTTPMethod, data.ContentType, data.Date, authMethod, data.Nonce,
			data.Resource)
		break
	default:
		str = fmt.Sprintf("%s\n%s\n%s\nx-ts-auth-method:%s\nx-ts-nonce:%s\n%s\n%s",
			data.HTTPMethod, data.ContentType, data.Date, authMethod, data.Nonce,
			data.Body, data.Resource)
	}

	key, _ := base64.StdEncoding.DecodeString(apiKey)
	hmac := hmacSHA256(key, str)
	signature := base64.StdEncoding.EncodeToString(hmac)

	return signature
}

func hmacSHA256(key []byte, content string) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(content))
	return mac.Sum(nil)
}

func getClient(timeout int) http.Client {
	return http.Client{
		Timeout: time.Duration(time.Duration(timeout) * time.Second),
	}
}
