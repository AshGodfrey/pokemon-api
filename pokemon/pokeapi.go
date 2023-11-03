package pokemon

import (
	"context"
	"net/http"
)

var DefaultClient = &Client{
	HTTPClient: http.DefaultClient,
	Endpoint:   "https://pokeapi.co/api/v2/",
}

// GetPokemon retrieves a Pokemon by its ID or name.
func GetPokemon(ctx context.Context, opts GetPokemonOpts) (Pokemon, error) {
	return DefaultClient.GetPokemon(ctx, opts)
}

// GetNature retrieves a Nature by its ID or name.
func GetNature(ctx context.Context, opts GetNatureOpts) (Nature, error) {
	return DefaultClient.GetNature(ctx, opts)
}

// GetStat retrieves a Stat by its ID or name.
func GetStat(ctx context.Context, opts GetStatOpts) (Stat, error) {
	return DefaultClient.GetStat(ctx, opts)
}
