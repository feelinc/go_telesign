package test

import (
	"net/http"
	"testing"

	telesign "github.com/feelinc/go_telesign"
	"github.com/feelinc/go_telesign/phoneid"
	"github.com/stretchr/testify/assert"
)

func TestNewPhoneID(t *testing.T) {
	req := phoneid.New(fakeIP, fakePhone, telesign.AccountLifecycleEventCreate)

	api := phoneid.NewClient(telesign.SetConEnv(telesign.EnvStandard),
		telesign.SetConCustomerID(fakeCustomerID),
		telesign.SetConAPIKey(fakeAPIKey))

	resp, err := api.Execute(req)
	if err != nil {
		panic(err)
	}

	assert.IsType(t, &phoneid.Request{}, req, "Should return request")
	assert.IsType(t, &telesign.PhoneID{}, api, "Should return PhoneID")
	assert.Equal(t, http.StatusUnauthorized, resp.GetStatusCode(), "Should return 401")
}

func TestNewPhoneIDLive(t *testing.T) {
	req := phoneid.NewLive(fakePhone, telesign.UcidBacs)

	api := phoneid.NewClient(telesign.SetConEnv(telesign.EnvEnterprise),
		telesign.SetConCustomerID(fakeCustomerID),
		telesign.SetConAPIKey(fakeAPIKey))

	resp, err := api.Execute(req)
	if err != nil {
		panic(err)
	}

	assert.IsType(t, &phoneid.LiveRequest{}, req, "Should return request")
	assert.IsType(t, &telesign.PhoneID{}, api, "Should return PhoneID")
	assert.Equal(t, http.StatusUnauthorized, resp.GetStatusCode(), "Should return 401")
}
