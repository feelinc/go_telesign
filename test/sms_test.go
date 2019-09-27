package test

import (
	"net/http"
	"testing"

	telesign "github.com/feelinc/go_telesign"
	"github.com/feelinc/go_telesign/sms"
	"github.com/stretchr/testify/assert"
)

func TestNewSMS(t *testing.T) {
	req := sms.New(fakeIP, fakePhone, fakeTemplate, telesign.MessageARN)

	api := sms.NewClient(telesign.SetConEnv(telesign.EnvEnterprise),
		telesign.SetConCustomerID(fakeCustomerID),
		telesign.SetConAPIKey(fakeAPIKey))

	resp, err := api.Execute(req)
	if err != nil {
		panic(err)
	}

	assert.IsType(t, &sms.Request{}, req, "Should return Request")
	assert.IsType(t, &telesign.SMS{}, api, "Should return SMS")
	assert.Equal(t, http.StatusUnauthorized, resp.GetStatusCode(), "Should return 401")
}

func TestNewSMSVerify(t *testing.T) {
	req := sms.NewVerify(fakeIP, fakePhone, telesign.UcidAtck, fakeLang, fakeCode,
		fakeTemplate)

	api := sms.NewClient(telesign.SetConEnv(telesign.EnvEnterprise),
		telesign.SetConCustomerID(fakeCustomerID),
		telesign.SetConAPIKey(fakeAPIKey))

	resp, err := api.Execute(req)
	if err != nil {
		panic(err)
	}

	assert.IsType(t, &sms.VerifyRequest{}, req, "Should return Verify request")
	assert.IsType(t, &telesign.SMS{}, api, "Should return SMS")
	assert.Equal(t, http.StatusUnauthorized, resp.GetStatusCode(), "Should return 401")
}
