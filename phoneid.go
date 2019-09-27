// Copyright (c) 2019 Sulaeman (me@sulaeman.com), All rights reserved.
// This source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package telesign

// PhoneID object
type PhoneID struct {
	conf Config
	req  Request
}

// SetEnv set the current environment
func (s *PhoneID) SetEnv(env string) {
	s.conf.Env = env
}

// SetCustomerID set the current customer ID
func (s *PhoneID) SetCustomerID(id string) {
	s.conf.CustomerID = id
}

// SetAPIKey set the current API key
func (s *PhoneID) SetAPIKey(key string) {
	s.conf.APIKey = key
}

// SetHTTPTimeout set the current HTTP request timeout
func (s *PhoneID) SetHTTPTimeout(i int) {
	s.conf.HTTPTimeout = i
}

// Execute the request
func (s PhoneID) Execute(req Request) (Response, error) {
	return execute(s.conf, req)
}

// NewPhoneID return a new PhoneID connection
func NewPhoneID(options ...OptConFunc) Connection {
	c := &PhoneID{
		conf: Config{
			HTTPTimeout: defaultHTTPTimeout,
		},
	}

	// Run the options on it
	for _, option := range options {
		if err := option(c); err != nil {
			return nil
		}
	}

	return c
}
