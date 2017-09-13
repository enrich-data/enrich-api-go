# graphmob-api-go

The Graphmob API Golang wrapper. Enrich, Search and Verify data from your Golang services.

Copyright 2017 Graphmob. See LICENSE for copying information.

* **📝 Implements**: [Graphmob REST API ~ v1](https://docs.graphmob.com/api/v1/) at reference revision: 09/12/2017
* **😘 Maintainer**: [@valeriansaliou](https://github.com/valeriansaliou)

## Usage

Import the library:

```go
import "gitlab.com/graphmob-com/graphmob-api-go/graphmob"
```

Construct a new authenticated Graphmob client with your `user_id` and `secret_key` tokens (you can generate those from your Graphmob Dashboard, [see the docs](https://docs.graphmob.com/api/v1/)).

```go
client := graphmob.New()

client.Authenticate("ui_a311da78-6b89-459c-8028-b331efab20d5", "sk_f293d44f-675d-4cb1-9c78-52b8a9af0df2")
```

Then, consume the client eg. to enrich an email address:

```go
import "fmt"

data, _, err := client.Enrich.EnrichPersonBy("email", "valerian@crisp.chat")

if err != nil {
  fmt.Printf("Error: %s", err)
} else {
  fmt.Printf("Enrich Lookup Emails: %s\n", data)
}
```

## Authentication

To authenticate against the API, generate your tokens (`user_id` and `secret_key`) **once** from your [Graphmob Dashboard](https://dashboard.graphmob.com/).

Then, pass those tokens **once** when you instanciate the Graphmob client as following:

```go
// Make sure to replace 'user_id' and 'secret_key' with your tokens
client.Authenticate("user_id", "secret_key")
```

## Data Discovery

**When Graphmob doesn't know about a given data point, eg. an email that was never enriched before, it launches a discovery. Discoveries can take a few seconds, and sometimes more than 10 seconds.**

This library implements a retry logic with a timeout if the discovery takes too long, or if the item wasn't found.

Thus, you can expect some requests, especially the Enrich requests, to take more time than expected. This is normal, and is not a performance issue on your side, or on our side. Under the hood, when you request a data point (eg. enrich a person given an email) that doesn't yet exist in our databases, the Graphmob API returns the HTTP response `201 Created`. Then, this library will poll the enrich resource for results, with intervals of a few seconds. The API will return `404 Not Found` as the discovery is still processing and no result is yet known at this point. Once a result is found, the API will reply with `200 OK` and return discovered data. If the discovery fails and no data can be aggregated for this email, the library aborts the retry after some time (less than 20 seconds), and returns a `not_found` error.

If a requested data point is already known by the Graphmob API, it will be immediately returned, which won't induce any delay.

## Resource Methods

This library implements all methods the Graphmob API provides. See the [API docs](https://docs.graphmob.com/api/v1/) for a reference of available methods, as well as how returned data is formatted.

### Search API

#### Lookup People

* **Method:** `client.Search.LookupPeopleBy(pageNumber, queryKey, queryValue)`
* **Docs:** [https://docs.graphmob.com/api/v1/#lookup-people](https://docs.graphmob.com/api/v1/#lookup-people)

```go
data, _, err := client.Search.LookupPeopleBy(1, "company_name", "Crisp")
```

#### Lookup Companies

* **Method:** `client.Search.LookupCompaniesBy(pageNumber, queryKey, queryValue)`
* **Docs:** [https://docs.graphmob.com/api/v1/#lookup-companies](https://docs.graphmob.com/api/v1/#lookup-companies)

```go
data, _, err := client.Search.LookupCompaniesBy(1, "name", "Crisp")
```

#### Lookup Emails

* **Method:** `client.Search.LookupEmails(pageNumber, emailDomain, legalName)`
* **Docs:** [https://docs.graphmob.com/api/v1/#lookup-emails](https://docs.graphmob.com/api/v1/#lookup-emails)

```go
data, _, err := client.Search.LookupEmails(1, "crisp.chat", "Crisp IM, Inc.")
```

#### Suggest Companies

* **Method:** `client.Search.SuggestCompanies(pageNumber, companyName)`
* **Docs:** [https://docs.graphmob.com/api/v1/#suggest-companies](https://docs.graphmob.com/api/v1/#suggest-companies)

```go
data, _, err := client.Search.SuggestCompanies(1, "Crisp")
```

### Verify API

#### Validate an Email

* **Method:** `client.Verify.ValidateEmail(email)`
* **Docs:** [https://docs.graphmob.com/api/v1/#validate-an-email](https://docs.graphmob.com/api/v1/#validate-an-email)

```go
data, _, err := client.Verify.ValidateEmail("valerian@crisp.chat")
```

#### Format an Email

* **Method:** `client.Verify.FormatEmail(emailDomain, firstName, lastName)`
* **Docs:** [https://docs.graphmob.com/api/v1/#format-an-email](https://docs.graphmob.com/api/v1/#format-an-email)

```go
data, _, err := client.Verify.FormatEmail("crisp.chat", "Valerian", "Saliou")
```

### Enrich API

#### Enrich a Person

* **Method:** `client.Enrich.EnrichPersonBy(key, value)`
* **Docs:** [https://docs.graphmob.com/api/v1/#enrich-a-person](https://docs.graphmob.com/api/v1/#enrich-a-person)

```go
data, _, err := client.Enrich.EnrichPersonBy("email", "valerian@crisp.chat")
```

#### Enrich a Company

* **Method:** `client.Enrich.EnrichCompanyBy(key, value)`
* **Docs:** [https://docs.graphmob.com/api/v1/#enrich-a-company](https://docs.graphmob.com/api/v1/#enrich-a-company)

```go
data, _, err := client.Enrich.EnrichCompanyBy("legal_name", "Crisp IM, Inc.")
```

#### Enrich a Network

* **Method:** `client.Enrich.EnrichNetworkBy(key, value)`
* **Docs:** [https://docs.graphmob.com/api/v1/#enrich-a-network](https://docs.graphmob.com/api/v1/#enrich-a-network)

```go
data, _, err := client.Enrich.EnrichNetworkBy("ip", "178.62.89.169")
```
