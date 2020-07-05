# Golang Miele 3rd Party API

Go client for the Miele 3rd Party API

## Usage

```go
package main

import (
    "fmt"
    "os"

    "github.com/ingmarstein/miele-go/miele"
)

func main() {
    token := os.Getenv("MIELE_AUTH_TOKEN")
    // You may optionally include your own http client
    client := miele.NewClient(nil, token)
    site, err := client.Site.List(&solaredge.ListOptions{Page: 2, PerPage: 1})
    if err != nil {
    	panic(err)
    }
    fmt.Println(site)
}
```