# SDK Readme

## Introduction

This is an SDK (Software Development Kit) for interacting with a Pokemon-related API. It provides convenient functions for retrieving information about Pokemon, Nature, and Stats. The SDK is designed to simplify the process of making API requests and handling responses in your Go applications.

## Installation

To use this SDK in your Go project, you can import it like this:

```go
go get -u "github.com/ashgodfrey/pokemon-api"
```

## SpeakEasy + Main.go 

In `main.go`, you will find two different (live) examples. One uses calls from a [managed SpeakEasy SDK](https://www.speakeasyapi.dev/docs/create-client-sdks). The other uses this SDK.  

You can find the "Testing Playground SDK" [here](https://github.com/speakeasy-sdks/testing-playground-sdk).

## Available Functions

The SDK provides the following functions for interacting with the Pokemon-related API:

### `GetPokemon`

Retrieve a Pokemon by its ID or name.

```go
package main

import (
	"context"
	"github.com/ashgodfrey/pokemonapi/api"
)

func main() {
	pokemon, err := api.GetPokemon(context.Background(), "bulbasaur")
	if err != nil {
		// Handle error
	}
	// Use the 'pokemon' object
}
```

#### Parameters

- `ctx`: `context.Context` - The context to control the lifetime of the request.
- `identifier`: `string` - The ID or name of the Pokemon to retrieve.

#### Returns

- `pokemon`: `pokemonapi.Pokemon` - A Pokemon object containing the details of the requested Pokemon.
- `err`: `error` - An error object that will be non-nil if there was an issue fetching the Pokemon.

## Decisions and Notes

* I opted to pass in the identifiers as strings instead of using an opts struct. I made this decision simply because the implemented methods all only take 1 string paramater. If I was implementing an endpoint that needed multiple parameters, I believe it would be best to use the opts struct for consistency.
