package phoneid

import (
	"net/http"
	"path"
	"testing"

	telesign "github.com/feelinc/go_telesign"
	"github.com/stretchr/testify/assert"
)

func testNewLiveRequest() telesign.Request {
	return NewLive(fakePhone, telesign.UcidBacs)
}

func TestNewLive(t *testing.T) {
	assert.IsType(t, &LiveRequest{}, testNewLiveRequest(),
		"Should return Live Request")
}

func TestLiveMethod(t *testing.T) {
	assert.Equal(t, http.MethodGet, testNewLiveRequest().GetMethod(),
		"Should similar to")
}

func TestLiveURI(t *testing.T) {
	expected := path.Join(liveURI, fakePhone) + "?ucid=" + telesign.UcidBacs
	assert.Equal(t, expected, testNewLiveRequest().GetURI(), "Should similar to")
}

func TestLivePath(t *testing.T) {
	expected := path.Join(liveURI, fakePhone)
	assert.Equal(t, expected, testNewLiveRequest().GetPath(), "Should similar to")
}

func TestLiveBody(t *testing.T) {
	assert.Equal(t, "", testNewLiveRequest().GetBody(),
		"Body should return equal to")
}
