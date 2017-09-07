// Copyright 2017 Valerian Saliou. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package graphmob


import (
  "fmt"
  "net/url"
)


// SearchService service
type SearchService service


// LookupCompaniesData mapping
type LookupCompaniesData struct {
  Companies  *[]Company  `json:"companies,omitempty"`
}

// LookupEmailsData mapping
type LookupEmailsData struct {
  CompanyCard  *LookupEmailsCompanyCard  `json:"company_card,omitempty"`
  Emails       *[]LookupEmailsEmail      `json:"emails,omitempty"`
}

// SuggestCompaniesData mapping
type SuggestCompaniesData struct {
  Companies  *[]SuggestCompaniesItem  `json:"companies,omitempty"`
}

// LookupEmailsCompanyCard mapping
type LookupEmailsCompanyCard struct {
  ID            *string  `json:"id,omitempty"`
  LegalName     *string  `json:"legal_name,omitempty"`
  EmailPattern  *string  `json:"email_pattern,omitempty"`
}

// LookupEmailsEmail mapping
type LookupEmailsEmail struct {
  Email       *string                       `json:"email,omitempty"`
  PersonCard  *LookupEmailsEmailPersonCard  `json:"person_card,omitempty"`
}

// LookupEmailsEmailPersonCard mapping
type LookupEmailsEmailPersonCard struct {
  ID       *string   `json:"id,omitempty"`
  Name     *Name     `json:"name,omitempty"`
  Contact  *Contact  `json:"contact,omitempty"`
}

// SuggestCompaniesItem mapping
type SuggestCompaniesItem struct {
  ID      *string  `json:"id,omitempty"`
  Name    *string  `json:"name,omitempty"`
  Domain  *string  `json:"domain,omitempty"`
  Logo    *string  `json:"logo,omitempty"`
}


// String returns the string representation of LookupCompaniesData
func (instance LookupCompaniesData) String() string {
  return Stringify(instance)
}

// String returns the string representation of LookupEmailsData
func (instance LookupEmailsData) String() string {
  return Stringify(instance)
}

// String returns the string representation of SuggestCompaniesData
func (instance SuggestCompaniesData) String() string {
  return Stringify(instance)
}

// String returns the string representation of LookupEmailsCompanyCard
func (instance LookupEmailsCompanyCard) String() string {
  return Stringify(instance)
}

// String returns the string representation of LookupEmailsEmail
func (instance LookupEmailsEmail) String() string {
  return Stringify(instance)
}

// String returns the string representation of LookupEmailsEmailPersonCard
func (instance LookupEmailsEmailPersonCard) String() string {
  return Stringify(instance)
}


// LookupCompaniesBy lookups in companies worldwide, provided search parameters.
func (service *SearchService) LookupCompaniesBy(pageNumber uint, key string, value string) (*LookupCompaniesData, *Response, error) {
  url := fmt.Sprintf("search/lookup/companies/%d?%s=%s", pageNumber, key, url.QueryEscape(value))
  req, _ := service.client.NewRequest("GET", url, nil)

  data := new(LookupCompaniesData)
  resp, err := service.client.Do(req, data)
  if err != nil {
    return nil, resp, err
  }

  return data, resp, err
}


// LookupEmails performs an email lookup, given the company email domain or legal name.
func (service *SearchService) LookupEmails(pageNumber uint, emailDomain string, legalName string) (*LookupEmailsData, *Response, error) {
  url := fmt.Sprintf("search/lookup/emails/%d?email_domain=%s&legal_name=%s", pageNumber, url.QueryEscape(emailDomain), url.QueryEscape(legalName))
  req, _ := service.client.NewRequest("GET", url, nil)

  data := new(LookupEmailsData)
  resp, err := service.client.Do(req, data)
  if err != nil {
    return nil, resp, err
  }

  return data, resp, err
}


// SuggestCompanies suggests companies based on the provided company name.
func (service *SearchService) SuggestCompanies(pageNumber uint, companyName string) (*SuggestCompaniesData, *Response, error) {
  url := fmt.Sprintf("search/suggest/companies/%d?company_name=%s", pageNumber, url.QueryEscape(companyName))
  req, _ := service.client.NewRequest("GET", url, nil)

  data := new(SuggestCompaniesData)
  resp, err := service.client.Do(req, data)
  if err != nil {
    return nil, resp, err
  }

  return data, resp, err
}
