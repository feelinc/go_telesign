package phoneid

import (
	"fmt"
	"net/http"
	"path"
	"testing"

	telesign "github.com/feelinc/go_telesign"
	"github.com/stretchr/testify/assert"
)

var expectedReqBody = fmt.Sprintf("account_lifecycle_event=%s&originating_ip=%s",
	telesign.AccountLifecycleEventCreate, fakeIP)

func testNewRequest() telesign.Request {
	return New(fakeIP, fakePhone, telesign.AccountLifecycleEventCreate)
}

func TestNewClient(t *testing.T) {
	api := NewClient(telesign.SetConEnv(telesign.EnvStandard),
		telesign.SetConCustomerID(fakeCustomerID),
		telesign.SetConAPIKey(fakeAPIKey))

	assert.IsType(t, &telesign.PhoneID{}, api, "Should return PhoneID")
}

func TestNew(t *testing.T) {
	assert.IsType(t, &Request{}, testNewRequest(), "Should return Request")
}

func TestMethod(t *testing.T) {
	assert.Equal(t, http.MethodPost, testNewRequest().GetMethod(),
		"Should similar to")
}

func TestURI(t *testing.T) {
	expected := path.Join(uri, fakePhone)
	assert.Equal(t, expected, testNewRequest().GetURI(), "Should similar to")
}

func TestPath(t *testing.T) {
	expected := path.Join(uri, fakePhone)
	assert.Equal(t, expected, testNewRequest().GetPath(), "Should similar to")
}

func TestBody(t *testing.T) {
	assert.Equal(t, expectedReqBody, testNewRequest().GetBody(),
		"Body should return equal to")
}
