package voice

import (
	"testing"

	telesign "github.com/feelinc/go_telesign"
	"github.com/stretchr/testify/assert"
)

func TestNewVoice(t *testing.T) {
	api := NewClient(telesign.SetConEnv(telesign.EnvStandard),
		telesign.SetConCustomerID(fakeCustomerID),
		telesign.SetConAPIKey(fakeAPIKey))

	assert.IsType(t, &telesign.Voice{}, api, "Should return Voice")
}
