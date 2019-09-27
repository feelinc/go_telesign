// Copyright (c) 2019 Sulaeman (me@sulaeman.com), All rights reserved.
// This source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package telesign

// SMS object
type SMS struct {
	conf Config
	req  Request
}

// SetEnv set the current environment
func (s *SMS) SetEnv(env string) {
	s.conf.Env = env
}

// SetCustomerID set the current customer ID
func (s *SMS) SetCustomerID(id string) {
	s.conf.CustomerID = id
}

// SetAPIKey set the current API key
func (s *SMS) SetAPIKey(key string) {
	s.conf.APIKey = key
}

// SetHTTPTimeout set the current HTTP request timeout
func (s *SMS) SetHTTPTimeout(i int) {
	s.conf.HTTPTimeout = i
}

// Execute the request
func (s SMS) Execute(req Request) (Response, error) {
	return execute(s.conf, req)
}

// NewSMS return a new SMS connection
func NewSMS(options ...OptConFunc) Connection {
	c := &SMS{
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
