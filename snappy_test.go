package snappy

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	client *Snappy
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = WithAPIKey("apikey")
	url, _ := url.Parse(server.URL)
	client.endpointPrefix = url.String()
}

func teardown() {
	server.Close()
}

func TestWithAPIKey(t *testing.T) {
	testClient := WithAPIKey("123")

	if testClient.username != "123" {
		t.Error("expected username == '123'")
	}

	if testClient.password != "x" {
		t.Error("expected username == 'x'")
	}
}

func TestWithUsernameAndPassword(t *testing.T) {
	testClient := WithUsernameAndPassword("username", "password")

	if testClient.username != "username" {
		t.Error("expected username == 'username'")
	}

	if testClient.password != "password" {
		t.Error("expected username == 'password'")
	}
}

func TestStatusNotOk(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `Not Found`)
	})

	_, err := client.Accounts()

	if err == nil {
		t.Error("expected err != nil")
	}
}
