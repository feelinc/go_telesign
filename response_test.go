package telesign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	stateCode := 200
	resp := MainResponse{
		StatusCode: stateCode,
	}

	assert.False(t, resp.Failure(), "Should only in 200 or 204")
	assert.Equal(t, stateCode, resp.GetStatusCode(), "Should having similar status code")
	assert.Empty(t, resp.Message(), "Should empty")
}
