package wercker

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wercker/go-wercker-api/credentials"
)

type testResult struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func TestClientMakeRequestGET400(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		result := `{"statusCode":400,"error":"Bad Request","message":"result must be one of aborted, unknown, passed, failed","details":[{"message":"result must be one of aborted, unknown, passed, failed","path":"result","type":"any.allowOnly","context":{"valids":"aborted, unknown, passed, failed","key":"result"}}]}`
		w.WriteHeader(400)
		w.Write([]byte(result))
	}))
	defer ts.Close()

	options := &Options{Endpoint: ts.URL}
	client := NewClient(options)
	err := client.MakeRequest("GET", "/", nil)

	require.Error(t, err, "")
	assert.Equal(t, "result must be one of aborted, unknown, passed, failed", err.Error(), "")
}

func TestClientMakeRequestGET200Anon(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		result := `{"key": "some key","value":"some value"}`
		w.WriteHeader(200)
		w.Write([]byte(result))
	}))
	defer ts.Close()

	result := &testResult{}
	options := &Options{Endpoint: ts.URL}
	client := NewClient(options)
	err := client.MakeRequest("GET", "/", result)

	require.NoError(t, err, "")
	assert.Equal(t, "some key", result.Key, "")
	assert.Equal(t, "some value", result.Value, "")
}

func TestClientMakeRequestGET200Token(t *testing.T) {
	tokenSet := false
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "Bearer secret_token" {
			tokenSet = true
		}
		result := `{"key": "some key","value":"some value"}`
		w.WriteHeader(200)
		w.Write([]byte(result))
	}))
	defer ts.Close()

	result := &testResult{}
	options := &Options{Endpoint: ts.URL, Creds: credentials.Token("secret_token")}
	client := NewClient(options)
	err := client.MakeRequest("GET", "/", result)

	require.NoError(t, err, "")
	assert.True(t, tokenSet, "")
	assert.Equal(t, "some key", result.Key, "")
	assert.Equal(t, "some value", result.Value, "")
}

func TestClientMakeRequestGET200UsernamePassword(t *testing.T) {
	authSet := false
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if user, pass, ok := r.BasicAuth(); ok && user == "secret username" && pass == "secret password" {
			authSet = true
		}
		result := `{"key": "some key","value":"some value"}`
		w.WriteHeader(200)
		w.Write([]byte(result))
	}))
	defer ts.Close()

	result := &testResult{}
	options := &Options{Endpoint: ts.URL, Creds: credentials.UsernamePassword("secret username", "secret password")}
	client := NewClient(options)
	err := client.MakeRequest("GET", "/", result)

	require.NoError(t, err, "")
	assert.True(t, authSet, "")
	assert.Equal(t, "some key", result.Key, "")
	assert.Equal(t, "some value", result.Value, "")
}
