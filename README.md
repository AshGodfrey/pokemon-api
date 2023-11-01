# SDK Readme

## Introduction

This is an SDK (Software Development Kit) for interacting with a Pokemon-related API. It provides convenient functions for retrieving information about Pokemon, Nature, and Stats. The SDK is designed to simplify the process of making API requests and handling responses in your Go applications.

## Installation

To use this SDK in your Go project, you can import it like this:

```go
go get -u "github.com/ashgodfrey/pokemon-api"
```

## Available Functions

The SDK provides the following functions for interacting with the Pokemon-related API:

### `GetPokemon`

The `GetPokemon` function retrieves a Pokemon by its ID or name.

```go
pokemon, err := api.GetPokemon(context.Background(), "bulbasaur")
if err != nil {
    // Handle error
}
// Use the 'pokemon' object
```
