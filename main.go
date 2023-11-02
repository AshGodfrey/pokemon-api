// Example showing both pokemon/api and /testing-playground-sdk
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ashgodfrey/pokemonapi/api"
	"github.com/speakeasy-sdks/testing-playground-sdk"
	"github.com/speakeasy-sdks/testing-playground-sdk/pkg/models/operations"
	"io/ioutil"
	"log"
)

// CustomResponse represents the fields to be included in the JSON output.
type CustomResponse struct {
	StatusCode  int    `json:"statusCode"`
	ContentType string `json:"contentType"`
	RawResponse []byte `json:"rawResponse,omitempty"`
}

func main() {
	// Initialize the AshTesting client
	s := testingplaygroundsdk.New()

	ctx := context.Background()

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

	statData, err := api.GetStat(ctx, api.StatHP)
	if err != nil {
		log.Fatalf("Error fetching stat: %v", err)
	}
	fmt.Printf("Stat Data: %+v\n", statData)

	

}
