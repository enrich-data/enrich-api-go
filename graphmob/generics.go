// Copyright 2017 Valerian Saliou. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package graphmob


// Person mapping
type Person struct {
  ID           *string              `json:"id,omitempty"`
  Name         *Name                `json:"name,omitempty"`
  Avatar       *string              `json:"avatar,omitempty"`
  Gender       *string              `json:"gender,omitempty"`
  Description  *string              `json:"description,omitempty"`
  Timezone     *string              `json:"timezone,omitempty"`
  Contact      *Contact             `json:"contact,omitempty"`
  Social       *PersonSocial        `json:"social,omitempty"`
  Address      *Address             `json:"address,omitempty"`
  Employments  *[]PersonEmployment  `json:"employments,omitempty"`
  Geolocation  *Geolocation         `json:"geolocation,omitempty"`
  Locales      *[]string            `json:"locales,omitempty"`
}

// PersonSocial mapping
type PersonSocial struct {
  Facebook  *PersonSocialNetwork  `json:"facebook,omitempty"`
  Twitter   *PersonSocialNetwork  `json:"twitter,omitempty"`
  LinkedIn  *PersonSocialNetwork  `json:"linkedin,omitempty"`
  GitHub    *PersonSocialNetwork  `json:"github,omitempty"`
}

// PersonSocialNetwork mapping
type PersonSocialNetwork struct {
  Handle  *string  `json:"handle,omitempty"`
  URL     *string  `json:"url,omitempty"`
}

// PersonEmployment mapping
type PersonEmployment struct {
  ID         *string  `json:"id,omitempty"`
  Name       *string  `json:"name,omitempty"`
  Domain     *string  `json:"domain,omitempty"`
  Title      *string  `json:"title,omitempty"`
  Role       *string  `json:"role,omitempty"`
  Seniority  *string  `json:"seniority,omitempty"`
}

// Company mapping
type Company struct {
  ID           *string           `json:"id,omitempty"`
  Name         *string           `json:"name,omitempty"`
  LegalName    *string           `json:"legal_name,omitempty"`
  Logo         *string           `json:"logo,omitempty"`
  Description  *string           `json:"description,omitempty"`
  Kind         *string           `json:"kind,omitempty"`
  Founded      *uint16           `json:"founded,omitempty"`
  Timezone     *string           `json:"timezone,omitempty"`
  Contact      *Contact          `json:"contact,omitempty"`
  Category     *CompanyCategory  `json:"category,omitempty"`
  Address      *Address          `json:"address,omitempty"`
  Metrics      *CompanyMetrics   `json:"metrics,omitempty"`
}

// CompanyCategory mapping
type CompanyCategory struct {
  Industry      *string    `json:"industry,omitempty"`
  Specialities  *[]string  `json:"specialities,omitempty"`
}

// CompanyMetrics mapping
type CompanyMetrics struct {
  AnnualRevenue     *CompanyMetricsAnnualRevenue  `json:"annual_revenue,omitempty"`
  Employees         *uint32                       `json:"employees,omitempty"`
  FacebookLikes     *uint32                       `json:"facebook_likes,omitempty"`
  TwitterFollowers  *uint32                       `json:"twitter_followers,omitempty"`
}

// CompanyMetricsAnnualRevenue mapping
type CompanyMetricsAnnualRevenue struct {
  Amount    *int64   `json:"amount,omitempty"`
  Currency  *string  `json:"currency,omitempty"`
}

// Network mapping
type Network struct {
  ID           *string          `json:"id,omitempty"`
  IP           *string          `json:"ip,omitempty"`
  Kind         *string          `json:"kind,omitempty"`
  Host         *NetworkHost     `json:"host,omitempty"`
  Reverse      *NetworkReverse  `json:"reverse,omitempty"`
  Geolocation  *Geolocation     `json:"geolocation,omitempty"`
  Block        *NetworkBlock    `json:"block,omitempty"`
  Usage        *NetworkUsage    `json:"usage,omitempty"`
}

// NetworkHost mapping
type NetworkHost struct {
  Reachable  *bool  `json:"reachable,omitempty"`
}

// NetworkReverse mapping
type NetworkReverse struct {
  Hostname  *string  `json:"hostname,omitempty"`
  Matches   *bool    `json:"matches,omitempty"`
}

// NetworkUsage mapping
type NetworkUsage struct {
  Home    *bool  `json:"home,omitempty"`
  Office  *bool  `json:"office,omitempty"`
  Mobile  *bool  `json:"mobile,omitempty"`
  Server  *bool  `json:"server,omitempty"`
  TOR     *bool  `json:"tor,omitempty"`
  VPN     *bool  `json:"vpn,omitempty"`
}

// NetworkBlock mapping
type NetworkBlock struct {
  Name   *string             `json:"name,omitempty"`
  Range  *string             `json:"range,omitempty"`
  Owner  *NetworkBlockOwner  `json:"owner,omitempty"`
}

// NetworkBlockOwner mapping
type NetworkBlockOwner struct {
  Organization  *string   `json:"organization,omitempty"`
  Person        *string   `json:"person,omitempty"`
  Contact       *Contact  `json:"contact,omitempty"`
  Address       *Address  `json:"address,omitempty"`
}

// Contact mapping
type Contact struct {
  Domain    *string    `json:"domain,omitempty"`
  Website   *string    `json:"website,omitempty"`
  Facebook  *string    `json:"facebook,omitempty"`
  Twitter   *string    `json:"twitter,omitempty"`
  LinkedIn  *uint      `json:"linkedin,omitempty"`
  Emails    *[]string  `json:"emails,omitempty"`
  Phones    *[]string  `json:"phones,omitempty"`
}

// Address mapping
type Address struct {
  Street       *string       `json:"street,omitempty"`
  Postcode     *string       `json:"postcode,omitempty"`
  City         *string       `json:"city,omitempty"`
  Region       *string       `json:"region,omitempty"`
  Country      *string       `json:"country,omitempty"`
  Coordinates  *Coordinates  `json:"coordinates,omitempty"`
}

// Geolocation mapping
type Geolocation struct {
  Country      *string       `json:"country,omitempty"`
  Region       *string       `json:"region,omitempty"`
  City         *string       `json:"city,omitempty"`
  Coordinates  *Coordinates  `json:"coordinates,omitempty"`
}

// Name mapping
type Name struct {
  Full   *string  `json:"full,omitempty"`
  First  *string  `json:"first,omitempty"`
  Last   *string  `json:"last,omitempty"`
}

// Coordinates mapping
type Coordinates struct {
  Latitude   *float32  `json:"latitude,omitempty"`
  Longitude  *float32  `json:"longitude,omitempty"`
}


// String returns the string representation of Person
func (instance Person) String() string {
  return Stringify(instance)
}

// String returns the string representation of PersonSocial
func (instance PersonSocial) String() string {
  return Stringify(instance)
}

// String returns the string representation of PersonSocialNetwork
func (instance PersonSocialNetwork) String() string {
  return Stringify(instance)
}

// String returns the string representation of PersonEmployment
func (instance PersonEmployment) String() string {
  return Stringify(instance)
}

// String returns the string representation of Company
func (instance Company) String() string {
  return Stringify(instance)
}

// String returns the string representation of CompanyCategory
func (instance CompanyCategory) String() string {
  return Stringify(instance)
}

// String returns the string representation of CompanyMetrics
func (instance CompanyMetrics) String() string {
  return Stringify(instance)
}

// String returns the string representation of CompanyMetricsAnnualRevenue
func (instance CompanyMetricsAnnualRevenue) String() string {
  return Stringify(instance)
}

// String returns the string representation of Network
func (instance Network) String() string {
  return Stringify(instance)
}

// String returns the string representation of NetworkHost
func (instance NetworkHost) String() string {
  return Stringify(instance)
}

// String returns the string representation of NetworkReverse
func (instance NetworkReverse) String() string {
  return Stringify(instance)
}

// String returns the string representation of NetworkUsage
func (instance NetworkUsage) String() string {
  return Stringify(instance)
}

// String returns the string representation of NetworkBlock
func (instance NetworkBlock) String() string {
  return Stringify(instance)
}

// String returns the string representation of NetworkBlockOwner
func (instance NetworkBlockOwner) String() string {
  return Stringify(instance)
}

// String returns the string representation of Contact
func (instance Contact) String() string {
  return Stringify(instance)
}

// String returns the string representation of Address
func (instance Address) String() string {
  return Stringify(instance)
}

// String returns the string representation of Geolocation
func (instance Geolocation) String() string {
  return Stringify(instance)
}

// String returns the string representation of Name
func (instance Name) String() string {
  return Stringify(instance)
}

// String returns the string representation of Coordinates
func (instance Coordinates) String() string {
  return Stringify(instance)
}
