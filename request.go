// Copyright (c) 2019 Sulaeman (me@sulaeman.com), All rights reserved.
// This source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package telesign

// Request interface
type Request interface {
	GetMethod() string
	GetURI() string
	GetPath() string
	GetBody() string
	ParseResponse(statusCode int, content []byte) (Response, error)
}
