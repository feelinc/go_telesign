package sms

import (
	"testing"

	telesign "github.com/feelinc/go_telesign"
	"github.com/stretchr/testify/assert"
)

func TestNewSMS(t *testing.T) {
	api := NewClient(telesign.SetConEnv(telesign.EnvStandard),
		telesign.SetConCustomerID(fakeCustomerID),
		telesign.SetConAPIKey(fakeAPIKey))

	assert.IsType(t, &telesign.SMS{}, api, "Should return SMS")
}
