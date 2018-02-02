// Copyright 2017 Valerian Saliou. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package enrich


import (
  "fmt"
  "net/url"
)


// EnrichService service
type EnrichService service


// EnrichPersonData mapping
type EnrichPersonData struct {
  Person     *Person     `json:"person,omitempty"`
  Companies  *[]Company  `json:"companies,omitempty"`
}

// EnrichNetworkData mapping
type EnrichNetworkData struct {
  Network  *Network  `json:"network,omitempty"`
  Company  *Company  `json:"company,omitempty"`
}


// String returns the string representation of EnrichPersonData
func (instance EnrichPersonData) String() string {
  return Stringify(instance)
}

// String returns the string representation of EnrichNetworkData
func (instance EnrichNetworkData) String() string {
  return Stringify(instance)
}


// EnrichPersonBy enriches data on a person with personal and company information on a person.
func (service *EnrichService) EnrichPersonBy(key string, value string) (*EnrichPersonData, *Response, error) {
  url := fmt.Sprintf("enrich/person?%s=%s", key, url.QueryEscape(value))
  req, _ := service.client.NewRequest("GET", url, nil)

  data := new(EnrichPersonData)
  resp, err := service.client.Do(req, data)
  if err != nil {
    return nil, resp, err
  }

  return data, resp, err
}

// EnrichNetworkBy enriches a network with network and company information.
func (service *EnrichService) EnrichNetworkBy(key string, value string) (*EnrichNetworkData, *Response, error) {
  url := fmt.Sprintf("enrich/network?%s=%s", key, url.QueryEscape(value))
  req, _ := service.client.NewRequest("GET", url, nil)

  data := new(EnrichNetworkData)
  resp, err := service.client.Do(req, data)
  if err != nil {
    return nil, resp, err
  }

  return data, resp, err
}
