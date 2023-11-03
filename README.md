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

## Usage

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


**Accepts:**
- `ID`: An integer representing the Pokémon ID (do not use if `Name` is provided).
- `Name`: A string representing the Pokémon name (do not use if `ID` is provided).
- `IncludeLocation` (optional): A boolean that when set to `true`, includes an array of location data in the response.

**Returns:**
- A `Pokemon` object containing the requested Pokémon details.
- An error object if the call fails.

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
**Accepts:**
- `ID`: An integer representing the Nature ID (do not use if `Name` is provided).
- `Name`: A string representing the Nature name (do not use if `ID` is provided).

**Returns:**
- A `Nature` object containing the requested Nature details.
- An error object if the call fails.


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

**Accepts:**
- `ID`: An integer representing the Stat ID (do not use if `Name` is provided).
- `Name`: A string representing the Stat name (do not use if `ID` is provided).

**Returns:**
- A `Stat` object containing the requested Stat details.
- An error object if the call fails.


## Testing

* `make` or `make test` to run all tests.
* `make mock-test` to run only the mocked response tests.
* `make e2e-test` to run the end-to-end tests.
* `make run-main` to run the main.go file.
*  `make get-pokemon-{pokemon name}` to run the CLI command for {pokemon name}.
*  `make get-location-{pokemon name}` to run the CLI command for {pokemon name} w/ `location=true` and return `pokemon.LocationData`.


## Decisions and Notes
* The GetPokemon function is designed to return comprehensive location data if `IncludeLocation` is set to `true`, otherwise it will return a URL.
* Basic data normalization accepts Pokémon names in any case, enhancing usability by abstracting case sensitivity concerns.
* I explored adding a list of constants for the `names` and `ID`s. I experimented with usingitg `iota` and explored the option of `go generate`. Ultimately this was not implemented due to the possible dynamic nature of the underlying Pokémon data.
* The choice to keep types within their current files, rather than a separate model file, was made to favor ease of development, as they are not expected to be shared across different packages.

## Tools Used

- [Go](https://golang.org/): The Go programming language is an open source project to make programmers more productive.

- [Cobra](https://github.com/spf13/cobra): Cobra is both a library for creating powerful modern CLI applications as well as a program to generate applications and command files.

