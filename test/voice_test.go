package test

import (
	"net/http"
	"testing"

	telesign "github.com/feelinc/go_telesign"
	"github.com/feelinc/go_telesign/voice"
	"github.com/stretchr/testify/assert"
)

func TestNewVoiceVerify(t *testing.T) {
	req := voice.NewVerify(fakeIP, fakePhone, telesign.UcidAtck, fakeLang,
		fakeCode, fakeTemplate, telesign.CallForwardActionBlock)

	api := voice.NewClient(telesign.SetConEnv(telesign.EnvEnterprise),
		telesign.SetConCustomerID(fakeCustomerID),
		telesign.SetConAPIKey(fakeAPIKey))

	resp, err := api.Execute(req)
	if err != nil {
		panic(err)
	}

	assert.IsType(t, &voice.VerifyRequest{}, req, "Should return Verify request")
	assert.IsType(t, &telesign.Voice{}, api, "Should return Voice")
	assert.Equal(t, http.StatusUnauthorized, resp.GetStatusCode(), "Should return 401")
}
