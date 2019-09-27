// Copyright (c) 2019 Sulaeman (me@sulaeman.com), All rights reserved.
// This source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package sms

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	telesign "github.com/feelinc/go_telesign"
	"github.com/stretchr/testify/assert"
)

var expectedVerifyReqBody = fmt.Sprintf("language=%s&originating_ip=%s&phone_number=%s&template=%s&ucid=%s&verify_code=%s", fakeLang, fakeIP,
	fakePhone, url.QueryEscape(fakeTemplate), telesign.UcidAtck, fakeCode)

func testNewRequest() telesign.Request {
	return NewVerify(fakeIP, fakePhone, telesign.UcidAtck, fakeLang, fakeCode,
		fakeTemplate)
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
	assert.Equal(t, verifyURI, testNewRequest().GetURI(), "Should similar to")
}

func TestPath(t *testing.T) {
	assert.Equal(t, verifyURI, testNewRequest().GetPath(), "Should similar to")
}

func TestVerifyBody(t *testing.T) {
	assert.Equal(t, expectedVerifyReqBody, testNewRequest().GetBody(),
		"Body should return equal to")
}
