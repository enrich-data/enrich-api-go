// Copyright 2017 Valerian Saliou. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package enrich


import (
  "bytes"
  "encoding/json"
  "fmt"
  "time"
  "io"
  "io/ioutil"
  "net/http"
  "net/url"
)


const (
  libraryVersion = "1.2.0"
  defaultRestEndpointURL = "https://api.enrichdata.com/v1/"
  userAgent = "enrich-api-go/" + libraryVersion
  acceptContentType = "application/json"
  clientTimeout = 40
)

// ClientConfig mapping
type ClientConfig struct {
  HTTPClient *http.Client
  RestEndpointURL string
}

type auth struct {
  Available bool
  Username string
  Password string
}

// Client maps an API client
type Client struct {
  config *ClientConfig
  client *http.Client
  auth *auth

  BaseURL *url.URL
  UserAgent string

  common service

  Verify *VerifyService
  Enrich *EnrichService
}

type service struct {
  client *Client
}

// Response maps an API HTTP response
type Response struct {
  *http.Response
}

type errorResponse struct {
  Error  errorResponseError  `json:"error,omitempty"`
}

type errorResponseError struct {
  Reason   string  `json:"reason,omitempty"`
  Message  string  `json:"message,omitempty"`
}


// Error prints an error response
func (response *errorResponseError) Error() string {
  return fmt.Sprintf("%v %v", response.Reason, response.Message)
}


// NewWithConfig returns a new API client
func NewWithConfig(config ClientConfig) *Client {
  // Defaults
  if config.HTTPClient == nil {
    config.HTTPClient = http.DefaultClient
    config.HTTPClient.Timeout = time.Duration(clientTimeout * time.Second)
  }
  if config.RestEndpointURL == "" {
    config.RestEndpointURL = defaultRestEndpointURL
  }

  // Create client
  baseURL, _ := url.Parse(config.RestEndpointURL)

  client := &Client{config: &config, client: config.HTTPClient, auth: &auth{}, BaseURL: baseURL, UserAgent: userAgent}
  client.common.client = client

  // Map services
  client.Verify = (*VerifyService)(&client.common)
  client.Enrich = (*EnrichService)(&client.common)

  return client
}


// New returns a new API client
func New() *Client {
  return NewWithConfig(ClientConfig{})
}


// Authenticate saves authentication parameters
func (client *Client) Authenticate(username string, password string) {
  client.auth.Username = username
  client.auth.Password = password
  client.auth.Available = true
}


// NewRequest creates an API request
func (client *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
  rel, err := url.Parse(urlStr)
  if err != nil {
    return nil, err
  }

  url := client.BaseURL.ResolveReference(rel)

  var buf io.ReadWriter
  if body != nil {
    buf = new(bytes.Buffer)
    err := json.NewEncoder(buf).Encode(body)
    if err != nil {
      return nil, err
    }
  }

  req, err := http.NewRequest(method, url.String(), buf)
  if err != nil {
    return nil, err
  }

  if client.auth.Available == true {
    req.SetBasicAuth(client.auth.Username, client.auth.Password)
  }

  req.Header.Add("Accept", acceptContentType)
  req.Header.Add("Content-Type", acceptContentType)

  if client.UserAgent != "" {
    req.Header.Add("User-Agent", client.UserAgent)
  }

  return req, nil
}


// Do sends an API request
func (client *Client) Do(req *http.Request, v interface{}) (*Response, error) {
  return client.DoInner(req, v)
}


// DoInner sends an API request (inner)
func (client *Client) DoInner(req *http.Request, v interface{}) (*Response, error) {
  resp, err := client.client.Do(req)
  if err != nil {
    return nil, err
  }

  defer func() {
    io.CopyN(ioutil.Discard, resp.Body, 512)
    resp.Body.Close()
  }()

  response := newResponse(resp)

  err = checkResponse(resp)
  if err != nil {
    return response, err
  }

  if decodeResponse(resp, v) == true {
    err = nil
  }

  return response, err
}


// newResponse creates an HTTP response
func newResponse(httpResponse *http.Response) *Response {
  response := &Response{Response: httpResponse}

  return response
}


// checkResponse checks response for errors
func checkResponse(response *http.Response) error {
  if code := response.StatusCode; 200 <= code && code <= 299 {
    return nil
  }
  errorResponse := &errorResponse{}

  if json.NewDecoder(response.Body).Decode(errorResponse) != nil {
    errorResponse.Error = errorResponseError{Reason: "error", Message: "Request could not be submitted."}
  }

  return &errorResponse.Error
}


// decodeResponse decodes response body
func decodeResponse(resp *http.Response, v interface{}) bool {
  if v != nil {
    if w, ok := v.(io.Writer); ok {
      io.Copy(w, resp.Body)
    } else if json.NewDecoder(resp.Body).Decode(v) == io.EOF {
      return true
    }
  }
  return false
}
