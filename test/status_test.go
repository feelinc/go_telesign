package test

import (
	"net/http"
	"testing"

	telesign "github.com/feelinc/go_telesign"
	"github.com/feelinc/go_telesign/status"
	"github.com/stretchr/testify/assert"
)

func TestNewStatus(t *testing.T) {
	req := status.New(fakeReferenceID, fakeCode)

	api := status.NewClient(telesign.SetConEnv(telesign.EnvEnterprise),
		telesign.SetConCustomerID(fakeCustomerID),
		telesign.SetConAPIKey(fakeAPIKey))

	resp, err := api.Execute(req)
	if err != nil {
		panic(err)
	}

	assert.IsType(t, &status.Request{}, req, "Should return request")
	assert.IsType(t, &telesign.Status{}, api, "Should return Status")
	assert.Equal(t, http.StatusUnauthorized, resp.GetStatusCode(), "Should return 401")
}
