// Copyright (c) 2019 Sulaeman (me@sulaeman.com), All rights reserved.
// This source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package telesign

// OptConFunc is a function that configures a Connection
type OptConFunc func(Connection) error

// Connection interface
type Connection interface {
	SetEnv(env string)
	SetCustomerID(id string)
	SetAPIKey(key string)
	SetHTTPTimeout(i int)
	Execute(req Request) (Response, error)
}

// SetConEnv set the environment
func SetConEnv(env string) OptConFunc {
	return func(c Connection) error {
		c.SetEnv(env)
		return nil
	}
}

// SetConCustomerID set the customer ID
func SetConCustomerID(id string) OptConFunc {
	return func(c Connection) error {
		c.SetCustomerID(id)
		return nil
	}
}

// SetConAPIKey set the API Key
func SetConAPIKey(key string) OptConFunc {
	return func(c Connection) error {
		c.SetAPIKey(key)
		return nil
	}
}

// Config data
type Config struct {
	Env         string
	CustomerID  string
	APIKey      string
	HTTPTimeout int
}
