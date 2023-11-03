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

* `IncludeLocation` is an optional boolean, if set to `true` then `location_area_encounters` will be an array of locations. Otherwise it will return a URL string as the value.

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
## Testing

* `make` or `make test` to run all tests.
* `make mock-test` to run only the mocked response tests.
* `make e2e-test` to run the end-to-end tests.
* `make run-main` to run the main.go file.
*  `make get-pokemon-{pokemon name}` to run the CLI command for {pokemon name}.
*  `make get-pokemon-location-{pokemon name}` to run the CLI command for {pokemon name} w/ `location=true` and return `pokemon.LocationData`.


## Decisions and Notes
* I streamlined the `GetPokemon` call to include the array of location details if a user sets `IncludeLocation` to `true`. Otherwise it will return the URL. 
* I did some very basic data normalization allowing `Pikachu` or `pikachu` to be sent. Theoretically, I could have left this to the user of the SDK. It's just a small quality-of-life improvement for users to not have to worry about casing.
* I heavily explored adding a list of constants for the `names` and `ID`s. I experimented with using `iota` and explored the option of `go generate`. In the end, because of the breadth of these fields and me not owning the underlying information, I opted to skip this.
* There's an argument to move the types into their model file for cleanlines. I decided against that because I don't necessairly expect these to be shared across packages.
* In a real world scenario Iw

## Tools Used

- [Go](https://golang.org/): The Go programming language is an open source project to make programmers more productive.

- [Cobra](https://github.com/spf13/cobra): Cobra is both a library for creating powerful modern CLI applications as well as a program to generate applications and command files.

