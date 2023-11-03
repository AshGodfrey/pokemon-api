// Example showing both pokemonapi/pokemon and /testing-playground-sdk
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ashgodfrey/pokemonapi/pokemon"
	"github.com/speakeasy-sdks/testing-playground-sdk"
	"github.com/speakeasy-sdks/testing-playground-sdk/pkg/models/operations"
"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	)

// CustomResponse represents the fields to be included in the JSON output.
type CustomResponse struct {
	StatusCode  int    `json:"statusCode"`
	ContentType string `json:"contentType"`
	RawResponse []byte `json:"rawResponse,omitempty"`
}

func main() {

	ctx := context.Background()

	// Initialize the AshTesting client
	s := testingplaygroundsdk.New()

	// Example 1: GetNatureIDOrName from testing-playground-sdk
	res, err := s.GetNatureIDOrName(ctx, operations.GetNatureIDOrNameRequest{
		IDOrName: "hardy",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create a custom response object with the desired fields.
	customRes := CustomResponse{
		StatusCode:  res.StatusCode,
		ContentType: res.ContentType,
	}

	// Access the underlying HTTP response to read the response body
	if res.RawResponse != nil {
		defer res.RawResponse.Body.Close()
		rawBody, err := ioutil.ReadAll(res.RawResponse.Body)
		if err != nil {
			log.Fatal(err)
		}
		customRes.RawResponse = rawBody
	}

	// Marshal the custom response to JSON
	jsonData, err := json.Marshal(customRes)
	if err != nil {
		log.Fatal(err)
	}

	// Print the JSON data
	log.Println(string(jsonData))

	// Display the fields in a human-readable format
	log.Printf("StatusCode: %d\n", customRes.StatusCode)
	log.Printf("ContentType: %s\n", customRes.ContentType)
	log.Printf("RawResponse: %s\n", string(customRes.RawResponse))

	// Example 2: Get data from ashgodfrey/pokemonapi

	pokemonData, err := pokemon.GetPokemon(ctx, pokemon.GetPokemonOpts{
		Name: "Pikachu",
	})
	if err != nil {
		log.Fatalf("Error fetching pokemon: %v", err)
	}
	fmt.Printf("Pokemon Data: %+v\n", pokemonData.LocationAreaEncounters)

	// Option 3: Pokemon CLI 

	var rootCmd = &cobra.Command{Use: "pokecli"}

	var cmdGetPokemon = &cobra.Command{
		Use:   "pokemon [name]",
		Short: "Get information about a Pokémon",
		Args:  cobra.ExactArgs(1),
		RunE:  GetPokemonCmd,
	}

	var cmdGetLocation = &cobra.Command{
		Use:   "location [name]",
		Short: "Get location information for a Pokémon",
		Args:  cobra.ExactArgs(1),
		RunE:  GetLocationCmd,
	}

	rootCmd.AddCommand(cmdGetPokemon, cmdGetLocation)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func GetPokemonCmd(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	pokemonName := args[0]
	pokemon, err := pokemon.GetPokemon(ctx, pokemon.GetPokemonOpts{
		Name: pokemonName,})
	if err != nil {
		fmt.Printf("Error getting Pokémon: %v\n", err)
		return err
	}

	fmt.Printf("Retrieved Pokémon: %+v\n", pokemon)
	return nil
}

func GetLocationCmd(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	pokemonName := args[0]
	location, err := pokemon.GetPokemon(ctx, pokemon.GetPokemonOpts{
		Name: pokemonName,
		IncludeLocation: true,})
	if err != nil {
		fmt.Printf("Error getting Pokémon location: %v\n", err)
		return err
	}

	fmt.Printf("Retrieved Pokémon Location: %+v\n", location.LocationAreaEncounters)
	return nil
}
