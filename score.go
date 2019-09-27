// Copyright (c) 2019 Sulaeman (me@sulaeman.com), All rights reserved.
// This source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package telesign

// Score object
type Score struct {
	conf Config
	req  Request
}

// SetEnv set the current environment
func (s *Score) SetEnv(env string) {
	s.conf.Env = env
}

// SetCustomerID set the current customer ID
func (s *Score) SetCustomerID(id string) {
	s.conf.CustomerID = id
}

// SetAPIKey set the current API key
func (s *Score) SetAPIKey(key string) {
	s.conf.APIKey = key
}

// SetHTTPTimeout set the current HTTP request timeout
func (s *Score) SetHTTPTimeout(i int) {
	s.conf.HTTPTimeout = i
}

// Execute the request
func (s Score) Execute(req Request) (Response, error) {
	return execute(s.conf, req)
}

// NewScore return a new Score connection
func NewScore(options ...OptConFunc) Connection {
	c := &Score{
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
