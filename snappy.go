package snappy

import (
  "encoding/json"
  "errors"
  "fmt"
  "io"
  "io/ioutil"
  "net/http"
  "net/url"
)

type snappy struct {
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
func WithAPIKey(apiKey string) *snappy {
  return &snappy{
    username:       apiKey,
    password:       "x",
    endpointPrefix: defaultEnpointPrefix,
  }
}

// WithUsernameAndPassword creates a new snappy client using your Username and Password
func WithUsernameAndPassword(username, password string) *snappy {
  return &snappy{
    username:       username,
    password:       password,
    endpointPrefix: defaultEnpointPrefix,
  }
}

func (s *snappy) get(up urlAndParams) (reader io.ReadCloser, err error) {
  client := &http.Client{}
  fullURL := fmt.Sprintf("%s%s", s.endpointPrefix, up.url)

  if len(up.params) > 0 {
    fullURL = fmt.Sprintf("%s?%s", fullURL, up.params.Encode())
  }

  request, err := http.NewRequest("GET", fullURL, nil)

  if err != nil {
    return
  }

  request.Header.Set("User-Agent", "Snappy go client ("+version+")")
  request.SetBasicAuth(s.username, s.password)

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

func (s *snappy) getReadAll(up urlAndParams) (b []byte, err error) {
  rc, err := s.get(up)

  if err != nil {
    return
  }

  defer rc.Close()
  return ioutil.ReadAll(rc)
}

func (s *snappy) unmarshalJSONAtURL(up urlAndParams, v interface{}) (err error) {
  b, err := s.getReadAll(up)

  if err != nil {
    return
  }

  return json.Unmarshal(b, v)
}
