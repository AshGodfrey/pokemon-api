//go:build e2e
// +build e2e

package pokemon

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetPokemonE2E(t *testing.T) {
	t.Run("Valid Pokemon", func(t *testing.T) {
		pokemonName := "pikachu"
		pokemonLocation := EncountersData("https://pokeapi.co/api/v2/pokemon/25/encounters")
		pokemon, err := GetPokemon(context.Background(), GetPokemonOpts{Name: pokemonName})

		require.NoError(t, err)
		require.NotNil(t, pokemon)
		require.Equal(t, pokemonName, pokemon.Name)
		require.Equal(t, pokemonLocation, pokemon.LocationAreaEncounters)
	})

	t.Run("Valid Pokemon with Location", func(t *testing.T) {
		pokemonName := "pikachu"
		includeLocation := true

		pokemonLocation := EncountersData("https://pokeapi.co/api/v2/pokemon/25/encounters")
		pokemon, err := GetPokemon(context.Background(), GetPokemonOpts{Name: pokemonName, IncludeLocation: includeLocation})

		require.NoError(t, err)
		require.NotNil(t, pokemon)
		require.Equal(t, pokemonName, pokemon.Name)
		require.NotEqual(t, pokemonLocation, pokemon.LocationAreaEncounters)
	})
	t.Run("Valid Pokemon by ID", func(t *testing.T) {
		pokemonID := 25
		pokemon, err := GetPokemon(context.Background(), GetPokemonOpts{ID: pokemonID})

		require.NoError(t, err)
		require.NotNil(t, pokemon)
		require.Equal(t, pokemonID, pokemon.ID)
	})

	t.Run("Invalid Pokemon", func(t *testing.T) {
		_, err := GetPokemon(context.Background(), GetPokemonOpts{Name: "invalid-pokemon-name"})

		require.Error(t, err)
	})
}

func TestGetNatureE2E(t *testing.T) {
	t.Run("Valid Nature", func(t *testing.T) {
		natureName := "adamant"
		nature, err := GetNature(context.Background(), GetNatureOpts{
			Name: natureName,
		})

		require.NoError(t, err)
		require.NotNil(t, nature)
		require.Equal(t, natureName, nature.Name)
	})

	t.Run("Invalid Nature", func(t *testing.T) {
		_, err := GetNature(context.Background(), GetNatureOpts{
			Name: "invalid-nature-name",
		})

		require.Error(t, err)
	})
}

func TestGetStatE2E(t *testing.T) {
	t.Run("Valid Stat", func(t *testing.T) {
		statName := "speed"
		stat, err := GetStat(context.Background(), GetStatOpts{
			Name: statName,
		})

		require.NoError(t, err)
		require.NotNil(t, stat)
		require.Equal(t, statName, stat.Name)
	})

	t.Run("Invalid Stat", func(t *testing.T) {
		_, err := GetStat(context.Background(), GetStatOpts{
			Name: "invalid-stat-name",
		})

		require.Error(t, err)
	})
}
