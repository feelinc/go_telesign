package status

import (
	"net/http"
	"path"
	"testing"

	telesign "github.com/feelinc/go_telesign"
	"github.com/stretchr/testify/assert"
)

const (
	fakeCustomerID = "12345678-9ABC-DEF0-1234-56789ABCDEF0"
	fakeAPIKey     = "vjE/ZDfPvDkuGNsuqCFFO4neYIs="
	fakeID         = "0123456789ABCDEF0123456789ABCDEF"
	fakeCode       = ""
)

func testNewRequest(code string) telesign.Request {
	return New(fakeID, code)
}

func TestNewClient(t *testing.T) {
	api := NewClient(telesign.SetConEnv(telesign.EnvStandard),
		telesign.SetConCustomerID(fakeCustomerID),
		telesign.SetConAPIKey(fakeAPIKey))

	assert.IsType(t, &telesign.Status{}, api, "Should return Status")
}

func TestNew(t *testing.T) {
	assert.IsType(t, &Request{}, testNewRequest(fakeCode),
		"Should return  Request")
}

func TestMethod(t *testing.T) {
	assert.Equal(t, http.MethodGet, testNewRequest(fakeCode).GetMethod(),
		"Should similar to")
}

func TestURI(t *testing.T) {
	expected := path.Join(uri, fakeID)
	assert.Equal(t, expected, testNewRequest(fakeCode).GetURI(), "Should similar to")
}

func TestURIWithCode(t *testing.T) {
	code := "12345"
	expected := path.Join(uri, fakeID) + "?verify_code=" + code
	assert.Equal(t, expected, testNewRequest(code).GetURI(), "Should similar to")
}

func TestPath(t *testing.T) {
	expected := path.Join(uri, fakeID)
	assert.Equal(t, expected, testNewRequest(fakeCode).GetPath(), "Should similar to")
}

func TestBody(t *testing.T) {
	assert.Equal(t, "", testNewRequest(fakeCode).GetBody(),
		"Body should return equal to")
}
