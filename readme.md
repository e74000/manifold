# Manifold Markets API Client Library

This is a Go client library for interacting with the [Manifold Markets API](https://manifold.markets/). 
The library aims to provide a convenient way to interact with the endpoints offered by the Manifold API.

## Installation

To install the library, use `go get`:

```sh
go get github.com/e74000/manifold
```

## Usage

To use this library, you'll need to create a new `Client` instance and authenticate it with your API key.

### Example

```go
package main

import (
    "fmt"
    "log"
    "github.com/e74000/manifold"
)

func main() {
    client := manifold.NewClient("your-api-key")

    // Get information about the authenticated user
    user, err := client.User.Me()
    if err != nil {
        log.Fatalf("Failed to retrieve user info: %v", err)
    }
    fmt.Printf("Authenticated user: %s\n", user.Username)

    // Create a new binary market
    market, err := client.Market.CreateBinary("Will it rain tomorrow?", 50, nil, nil, nil, nil)
    if err != nil {
        log.Fatalf("Failed to create market: %v", err)
    }
    fmt.Printf("Created market: %s\n", market.Question)
}
```

## Contributing

Please feel free to contribute!

## License

This library is licensed under the MIT License. See the `LICENSE` file for details.
