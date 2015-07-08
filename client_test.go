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
	result := []byte(`{"statusCode":400,"error":"Bad Request","message":"result must be one of aborted, unknown, passed, failed","details":[{"message":"result must be one of aborted, unknown, passed, failed","path":"result","type":"any.allowOnly","context":{"valids":"aborted, unknown, passed, failed","key":"result"}}]}`)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		w.Write(result)
	}))
	defer ts.Close()

	options := &Options{Endpoint: ts.URL}
	client := NewClient(options)
	_, err := client.MakeRequest("GET", "/")

	require.Error(t, err, "")
	assert.Equal(t, "result must be one of aborted, unknown, passed, failed", err.Error(), "")
}

func TestClientMakeRequestGET200Anon(t *testing.T) {
	result := []byte(`{"key": "some key","value":"some value"}`)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write(result)
	}))
	defer ts.Close()

	options := &Options{Endpoint: ts.URL}
	client := NewClient(options)
	body, err := client.MakeRequest("GET", "/")

	require.NoError(t, err, "")
	assert.Equal(t, result, body, "")
}

func TestClientMakeRequestGET200Token(t *testing.T) {
	result := []byte(`{"key": "some key","value":"some value"}`)
	tokenSet := false
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "Bearer secret_token" {
			tokenSet = true
		}
		w.WriteHeader(200)
		w.Write(result)
	}))
	defer ts.Close()

	options := &Options{Endpoint: ts.URL, Creds: credentials.Token("secret_token")}
	client := NewClient(options)
	body, err := client.MakeRequest("GET", "/")

	require.NoError(t, err, "")
	assert.True(t, tokenSet, "")
	assert.Equal(t, result, body, "")
}

func TestClientMakeRequestGET200UsernamePassword(t *testing.T) {
	authSet := false
	result := []byte(`{"key": "some key","value":"some value"}`)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if user, pass, ok := r.BasicAuth(); ok && user == "secret username" && pass == "secret password" {
			authSet = true
		}
		w.WriteHeader(200)
		w.Write(result)
	}))
	defer ts.Close()

	options := &Options{Endpoint: ts.URL, Creds: credentials.UsernamePassword("secret username", "secret password")}
	client := NewClient(options)
	body, err := client.MakeRequest("GET", "/")

	require.NoError(t, err, "")
	assert.True(t, authSet, "")
	assert.Equal(t, result, body, "")
}
