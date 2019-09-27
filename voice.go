// Copyright (c) 2019 Sulaeman (me@sulaeman.com), All rights reserved.
// This source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package telesign

// Voice object
type Voice struct {
	conf Config
	req  Request
}

// SetEnv set the current environment
func (s *Voice) SetEnv(env string) {
	s.conf.Env = env
}

// SetCustomerID set the current customer ID
func (s *Voice) SetCustomerID(id string) {
	s.conf.CustomerID = id
}

// SetAPIKey set the current API key
func (s *Voice) SetAPIKey(key string) {
	s.conf.APIKey = key
}

// SetHTTPTimeout set the current HTTP request timeout
func (s *Voice) SetHTTPTimeout(i int) {
	s.conf.HTTPTimeout = i
}

// Execute the request
func (s Voice) Execute(req Request) (Response, error) {
	return execute(s.conf, req)
}

// NewVoice return a new Voice connection
func NewVoice(options ...OptConFunc) Connection {
	c := &Voice{
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
