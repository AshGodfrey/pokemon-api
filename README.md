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
	pokemon, err := api.GetPokemon(context.Background(), GetPokemonOpts {
        ID: 1,
        IncludeLocation: true,
    })
	if err != nil {
		// Handle error
	}
	// Use the 'pokemon' object
}
```

### `GetNature`

Retrieve a Nature by its ID or name.

```go
package main

import (
	"context"
	"github.com/ashgodfrey/pokemonapi/api"
)

func main() {
	pokemon, err := api.GetNature(context.Background(), GetNatureOpts{
        ID: 1,
    })
	if err != nil {
		// Handle error
	}
	// Use the 'nature' object
}
```


### `GetStat`

Retrieve a Stat by its ID or name.

```go
package main

import (
	"context"
	"github.com/ashgodfrey/pokemonapi/api"
)

func main() {
	stat, err := api.GetStat(context.Background(), GetStatOpts{
        Name: "speed",
    })
	if err != nil {
		// Handle error
	}
	// Use the 'stat' object
}
```


## Decisions and Notes

