package test

import (
	"net/http"
	"testing"

	telesign "github.com/feelinc/go_telesign"
	"github.com/feelinc/go_telesign/score"
	"github.com/stretchr/testify/assert"
)

func TestNewScore(t *testing.T) {
	req := score.New(fakePhone, telesign.AccountLifecycleEventCreate)

	api := score.NewClient(telesign.SetConEnv(telesign.EnvStandard),
		telesign.SetConCustomerID(fakeCustomerID),
		telesign.SetConAPIKey(fakeAPIKey))

	resp, err := api.Execute(req)
	if err != nil {
		panic(err)
	}

	assert.IsType(t, &score.Request{}, req, "Should return request")
	assert.IsType(t, &telesign.Score{}, api, "Should return Score")
	assert.Equal(t, http.StatusUnauthorized, resp.GetStatusCode(), "Should return 401")
}
