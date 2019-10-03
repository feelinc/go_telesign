package telesign

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const (
	fakeCustomerID  = "12345678-9ABC-DEF0-1234-56789ABCDEF0"
	fakeAPIKey      = "vjE/ZDfPvDkuGNsuqCFFO4neYIs="
	fakeMethod      = "POST"
	fakeResource    = "/v1/uri"
	fakeContentType = "application/x-www-form-urlencoded"
	fakeBody        = "the body"
)

func testBuildSignature(t time.Time, n string) signatureData {
	return buildSignature(t, n, fakeMethod, fakeResource, fakeContentType,
		fakeBody)
}

func TestNew(t *testing.T) {
	tel := New(SetEnv(EnvStandard), SetCustomerID(fakeCustomerID),
		SetAPIKey(fakeAPIKey))

	assert.IsType(t, &Telesign{}, tel, "Should return Telesign")
}

func TestPhoneID(t *testing.T) {
	tel := New(SetEnv(EnvStandard), SetCustomerID(fakeCustomerID),
		SetAPIKey(fakeAPIKey))

	assert.IsType(t, &PhoneID{}, tel.PhoneID(), "Should return PhoneID")
}

func TestPhoneIDConf(t *testing.T) {
	phoneID := New().PhoneID(SetConEnv(EnvStandard), SetConCustomerID(fakeCustomerID),
		SetConAPIKey(fakeAPIKey))

	assert.IsType(t, &PhoneID{}, phoneID, "Should return PhoneID")
}

func TestScore(t *testing.T) {
	tel := New(SetEnv(EnvStandard), SetCustomerID(fakeCustomerID),
		SetAPIKey(fakeAPIKey))

	assert.IsType(t, &Score{}, tel.Score(), "Should return Score")
}

func TestScoreConf(t *testing.T) {
	score := New().Score(SetConEnv(EnvStandard), SetConCustomerID(fakeCustomerID),
		SetConAPIKey(fakeAPIKey))

	assert.IsType(t, &Score{}, score, "Should return Score")
}

func TestSMS(t *testing.T) {
	tel := New(SetEnv(EnvStandard), SetCustomerID(fakeCustomerID),
		SetAPIKey(fakeAPIKey))

	assert.IsType(t, &SMS{}, tel.SMS(), "Should return SMS")
}

func TestSMSConf(t *testing.T) {
	sms := New().SMS(SetConEnv(EnvStandard), SetConCustomerID(fakeCustomerID),
		SetConAPIKey(fakeAPIKey))

	assert.IsType(t, &SMS{}, sms, "Should return SMS")
}

func TestVoice(t *testing.T) {
	tel := New(SetEnv(EnvStandard), SetCustomerID(fakeCustomerID),
		SetAPIKey(fakeAPIKey))

	assert.IsType(t, &Voice{}, tel.Voice(), "Should return Voice")
}

func TestVoiceConf(t *testing.T) {
	voice := New().Voice(SetConEnv(EnvStandard), SetConCustomerID(fakeCustomerID),
		SetConAPIKey(fakeAPIKey))

	assert.IsType(t, &Voice{}, voice, "Should return Voice")
}

func TestStatus(t *testing.T) {
	tel := New(SetEnv(EnvStandard), SetCustomerID(fakeCustomerID),
		SetAPIKey(fakeAPIKey))

	assert.IsType(t, &Status{}, tel.Status(), "Should return Status")
}

func TestStatusConf(t *testing.T) {
	status := New().Status(SetConEnv(EnvStandard), SetConCustomerID(fakeCustomerID),
		SetConAPIKey(fakeAPIKey))

	assert.IsType(t, &Status{}, status, "Should return Status")
}

func TestStructToURLValues(t *testing.T) {
	data := struct {
		name string
		add  string
	}{
		name: "name val",
		add:  "add val",
	}
	expected := "add=add+val&name=name+val"
	val := StructToURLValues(data)

	assert.IsType(t, url.Values{}, val, "Should return a URL Values")
	assert.Equal(t, expected, val.Encode(), "Should similar to")
}

func TestBuildRequestURI(t *testing.T) {
	expected := fmt.Sprintf("%s%s", getDomain(EnvStandard), fakeResource)
	uri := buildRequestURI(getDomain(EnvStandard), fakeResource)
	assert.Equal(t, expected, uri, "Should similar to")

	expected = fmt.Sprintf("%s%s", getDomain(EnvEnterprise), fakeResource)
	uri = buildRequestURI(getDomain(EnvEnterprise), fakeResource)
	assert.Equal(t, expected, uri, "Should similar to")
}

func TestBuildRequest(t *testing.T) {
	uri := buildRequestURI(getDomain(EnvStandard), fakeResource)
	req, err := buildRequest(http.MethodPost, uri, bytes.NewBuffer([]byte("")))

	assert.IsType(t, &http.Request{}, req, "Should return HTTP Request")
	assert.Nil(t, err, "Should not having error")
}

func TestGetDomain(t *testing.T) {
	assert.Equal(t, domain, getDomain(EnvStandard), "Should return domain standard")
	assert.Equal(t, domainEnterprise, getDomain(EnvEnterprise), "Should return domain enterprise")
}

func TestBuildSignature(t *testing.T) {
	tm := time.Now()
	nonce, _ := uuid.NewRandom()
	data := testBuildSignature(tm, nonce.String())

	assert.IsType(t, signatureData{}, data, "Should return signatureData struct")
	assert.NotEmpty(t, data.HTTPMethod, "Should not empty")
	assert.NotEmpty(t, data.Resource, "Should not empty")
	assert.NotEmpty(t, data.ContentType, "Should not empty")
	assert.NotEmpty(t, data.Date, "Should not empty")
	assert.NotEmpty(t, data.Nonce, "Should not empty")
	assert.NotEmpty(t, data.Body, "Should not empty")
	assert.Equal(t, fakeMethod, data.HTTPMethod, "Should similar to")
	assert.Equal(t, fakeResource, data.Resource, "Should similar to")
	assert.Equal(t, fakeContentType, data.ContentType, "Should similar to")
	assert.Equal(t, tm.Format(timeFormat), data.Date, "Should similar to")
	assert.Equal(t, fakeBody, data.Body, "Should similar to")
}

func TestCreateSignature(t *testing.T) {
	tm, _ := time.Parse(timeFormat, "Thu, 26 Sep 2019 10:00:45 +0700")
	nonce := "f2e5d1f6-ba17-48a7-8a74-b94a187db798"

	expected := "PszieAmepc085FrIe/DXGXUQQDd2b8b6AEjYvaaqJVU="
	sig := createSignature(fakeAPIKey, testBuildSignature(tm, nonce))

	assert.Equal(t, expected, sig, "Should similar to")
}

func TestGetClient(t *testing.T) {
	c := getClient(defaultHTTPTimeout)

	assert.IsType(t, http.Client{}, c, "Should return a HTTP Client")
}
