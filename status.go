// Copyright (c) 2019 Sulaeman (me@sulaeman.com), All rights reserved.
// This source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package telesign

// Status object
type Status struct {
	conf Config
	req  Request
}

// SetEnv set the current environment
func (s *Status) SetEnv(env string) {
	s.conf.Env = env
}

// SetCustomerID set the current customer ID
func (s *Status) SetCustomerID(id string) {
	s.conf.CustomerID = id
}

// SetAPIKey set the current API key
func (s *Status) SetAPIKey(key string) {
	s.conf.APIKey = key
}

// SetHTTPTimeout set the current HTTP request timeout
func (s *Status) SetHTTPTimeout(i int) {
	s.conf.HTTPTimeout = i
}

// Execute the request
func (s Status) Execute(req Request) (Response, error) {
	return execute(s.conf, req)
}

// NewStatus return a new Status connection
func NewStatus(options ...OptConFunc) Connection {
	c := &Status{
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
