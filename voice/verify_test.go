// Copyright (c) 2019 Sulaeman (me@sulaeman.com), All rights reserved.
// This source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package voice

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	telesign "github.com/feelinc/go_telesign"
	"github.com/stretchr/testify/assert"
)

var expectedVerifyReqBody = fmt.Sprintf("call_forward_action=%s&language=%s&originating_ip=%s&phone_number=%s&tts_message=%s&ucid=%s&verify_code=%s",
	telesign.CallForwardActionBlock, fakeLang, fakeIP,
	fakePhone, url.QueryEscape(fakeTemplate), telesign.UcidAtck, fakeCode)

func testNewRequest() telesign.Request {
	return NewVerify(fakeIP, fakePhone, telesign.UcidAtck, fakeLang, fakeCode,
		fakeTemplate, telesign.CallForwardActionBlock)
}

func TestNewVerify(t *testing.T) {
	assert.IsType(t, &VerifyRequest{}, testNewRequest(),
		"Should return Verify request")
}

func TestMethod(t *testing.T) {
	assert.Equal(t, http.MethodPost, testNewRequest().GetMethod(),
		"Should similar to")
}

func TestURI(t *testing.T) {
	assert.Equal(t, uri, testNewRequest().GetURI(), "Should similar to")
}

func TestPath(t *testing.T) {
	assert.Equal(t, uri, testNewRequest().GetPath(), "Should similar to")
}

func TestVerifyBody(t *testing.T) {
	assert.Equal(t, expectedVerifyReqBody, testNewRequest().GetBody(),
		"Body should return equal to")
}
