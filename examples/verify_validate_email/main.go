// Copyright 2017 Valerian Saliou. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
  "github.com/enrich-data/enrich-api-go/enrich"
  "fmt"
)

func main() {
  client := enrich.New()
  client.Authenticate("ui_a311da78-6b89-459c-8028-b331efab20d5", "sk_f293d44f-675d-4cb1-9c78-52b8a9af0df2")

  data, _, err := client.Verify.ValidateEmail("valerian@crisp.chat")

  if err != nil {
    fmt.Printf("Error: %s", err)
  } else {
    fmt.Printf("Verify Validate Email (raw): %s\n", data)
  }
}
