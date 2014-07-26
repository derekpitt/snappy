// Package snappy provides a simple interface with the Snappy API
package snappy

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Snappy is the main type
type Snappy struct {
	username       string
	password       string
	endpointPrefix string
}

type urlAndParams struct {
	url    string
	params url.Values
}

const (
	defaultEnpointPrefix = "https://app.besnappy.com/api/v1"
	version              = "0.0.1"
)

// WithAPIKey creates a new snappy client using your API key
func WithAPIKey(apiKey string) *Snappy {
	return &Snappy{
		username:       apiKey,
		password:       "x",
		endpointPrefix: defaultEnpointPrefix,
	}
}

// WithUsernameAndPassword creates a new snappy client using your Username and Password
func WithUsernameAndPassword(username, password string) *Snappy {
	return &Snappy{
		username:       username,
		password:       password,
		endpointPrefix: defaultEnpointPrefix,
	}
}

func (up urlAndParams) finalURL(endpointPrefix string) string {
	fullURL := fmt.Sprintf("%s%s", endpointPrefix, up.url)

	if len(up.params) > 0 {
		fullURL = fmt.Sprintf("%s?%s", fullURL, up.params.Encode())
	}

	return fullURL
}

func (s *Snappy) doRequest(requestType string, up urlAndParams, contentType string, body io.Reader) (reader io.ReadCloser, err error) {
	client := &http.Client{}
	request, err := http.NewRequest(requestType, up.finalURL(s.endpointPrefix), body)

	if err != nil {
		return
	}

	request.SetBasicAuth(s.username, s.password)

	request.Header.Set("User-Agent", "Snappy go client ("+version+")")
	if len(contentType) > 0 {
		request.Header.Set("Content-Type", contentType)
	}

	res, err := client.Do(request)

	if err != nil {
		return
	}

	// TODO: double check to see if their api returns anything other
	// than a 200 when a request is bad
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("Status NOT OK")
	}

	return res.Body, nil
}

func (s *Snappy) get(up urlAndParams) (reader io.ReadCloser, err error) {
	return s.doRequest("GET", up, "", nil)
}

func (s *Snappy) post(up urlAndParams, contentType string, body io.Reader) (reader io.ReadCloser, err error) {
	return s.doRequest("POST", up, contentType, body)
}

func (s *Snappy) postForm(up urlAndParams, values url.Values) (reader io.ReadCloser, err error) {
	bodyReader := strings.NewReader(values.Encode())

	return s.post(up, "application/x-www-form-urlencoded", bodyReader)
}

func (s *Snappy) postAsJSON(up urlAndParams, v interface{}) (reader io.ReadCloser, err error) {
	b, err := json.Marshal(v)

	if err != nil {
		return
	}

	return s.post(up, "application/json", strings.NewReader(string(b)))
}

func (s *Snappy) del(up urlAndParams) (reader io.ReadCloser, err error) {
	return s.doRequest("DELETE", up, "", nil)
}

func (s *Snappy) unmarshalJSONAtURL(up urlAndParams, v interface{}) (err error) {
	rc, err := s.get(up)

	if err != nil {
		return
	}

	defer rc.Close()
	return json.NewDecoder(rc).Decode(&v)
}
