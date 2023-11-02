
package api

import (
	"context"
	"net/http"

)

var DefaultClient = &Client{
	HTTPClient: http.DefaultClient,
	Endpoint: "https://pokeapi.co/api/v2",
}

// GetPokemon retrieves a Pokemon by its ID or name.
func GetPokemon(ctx context.Context, idOrName string) (Pokemon, error) {
	return DefaultClient.GetPokemon(ctx, idOrName)
}

// GetNature retrieves a Nature by its ID or name.
func GetNature(ctx context.Context, idOrName string) (Nature, error) {
	return DefaultClient.GetNature(ctx, idOrName)
}

// GetStat retrieves a Stat by its ID or name.
func GetStat(ctx context.Context, idOrName string) (Stat, error) {
	return DefaultClient.GetStat(ctx, idOrName)
}
