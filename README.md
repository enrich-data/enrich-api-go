# enrich-api-go

The Enrich API Golang wrapper. Enrich, Search and Verify data from your Golang services.

Copyright 2017 Enrich. See LICENSE for copying information.

* **📝 Implements**: [Enrich REST API ~ v1](https://docs.enrichdata.com/api/v1/) at reference revision: 10/05/2017
* **😘 Maintainer**: [@valeriansaliou](https://github.com/valeriansaliou)

## Usage

Import the library:

```go
import "github.com/enrich-data/enrich-api-go/enrich"
```

Construct a new authenticated Enrich client with your `user_id` and `secret_key` tokens (you can generate those from your Enrich Dashboard, [see the docs](https://docs.enrichdata.com/api/v1/)).

```go
client := enrich.New()

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

To authenticate against the API, generate your tokens (`user_id` and `secret_key`) **once** from your [Enrich Dashboard](https://dashboard.enrichdata.com/).

Then, pass those tokens **once** when you instanciate the Enrich client as following:

```go
// Make sure to replace 'user_id' and 'secret_key' with your tokens
client.Authenticate("user_id", "secret_key")
```

## Data Discovery

**When Enrich doesn't know about a given data point, eg. an email that was never enriched before, it launches a discovery. Discoveries can take a few seconds, and sometimes more than 10 seconds.**

This library implements a retry logic with a timeout if the discovery takes too long, or if the item wasn't found.

Thus, you can expect some requests, especially the Enrich requests, to take more time than expected. This is normal, and is not a performance issue on your side, or on our side. Under the hood, when you request a data point (eg. enrich a person given an email) that doesn't yet exist in our databases, the Enrich API returns the HTTP response `201 Created`. Then, this library will poll the enrich resource for results, with intervals of a few seconds. The API will return `404 Not Found` as the discovery is still processing and no result is yet known at this point. Once a result is found, the API will reply with `200 OK` and return discovered data. If the discovery fails and no data can be aggregated for this email, the library aborts the retry after some time (less than 20 seconds), and returns a `not_found` error.

If a requested data point is already known by the Enrich API, it will be immediately returned, which won't induce any delay.

## Resource Methods

This library implements all methods the Enrich API provides. See the [API docs](https://docs.enrichdata.com/api/v1/) for a reference of available methods, as well as how returned data is formatted.

### Verify API

#### Validate an Email

* **Method:** `client.Verify.ValidateEmail(email)`
* **Docs:** [https://docs.enrichdata.com/api/v1/#validate-an-email](https://docs.enrichdata.com/api/v1/#validate-an-email)

```go
data, _, err := client.Verify.ValidateEmail("valerian@crisp.chat")
```

### Enrich API

#### Enrich a Person

* **Method:** `client.Enrich.EnrichPersonBy(key, value)`
* **Docs:** [https://docs.enrichdata.com/api/v1/#enrich-a-person](https://docs.enrichdata.com/api/v1/#enrich-a-person)

```go
data, _, err := client.Enrich.EnrichPersonBy("email", "valerian@crisp.chat")
```

#### Enrich a Company

* **Method:** `client.Enrich.EnrichCompanyBy(key, value)`
* **Docs:** [https://docs.enrichdata.com/api/v1/#enrich-a-company](https://docs.enrichdata.com/api/v1/#enrich-a-company)

```go
data, _, err := client.Enrich.EnrichCompanyBy("domain", "crisp.chat")
```

#### Enrich a Network

* **Method:** `client.Enrich.EnrichNetworkBy(key, value)`
* **Docs:** [https://docs.enrichdata.com/api/v1/#enrich-a-network](https://docs.enrichdata.com/api/v1/#enrich-a-network)

```go
data, _, err := client.Enrich.EnrichNetworkBy("ip", "178.62.89.169")
```
