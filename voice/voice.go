// Copyright (c) 2019 Sulaeman (me@sulaeman.com), All rights reserved.
// This source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package voice

import telesign "github.com/feelinc/go_telesign"

const voiceURI = "/v2/voice"

// NewClient return new Voice API connection
func NewClient(options ...telesign.OptConFunc) telesign.Connection {
	return telesign.NewVoice(options...)
}
