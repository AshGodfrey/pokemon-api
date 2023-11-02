package api

import (
	"context"
	"testing"
	"time"
)

func setup() *Client {
	return NewClient()
}

func TestGetPokemonIntegration(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	testCases := []struct {
		name    string
		wantErr bool
	}{
		{"pikachu", false},
		{"nonexistentpokemon", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pokemon, err := client.GetPokemon(ctx, tc.name)

			if tc.wantErr {
				if err == nil {
					t.Errorf("Expected an error for %s but got none", tc.name)
				}
				return
			}

			if err != nil {
				t.Fatalf("Error fetching Pokémon '%s': %v", tc.name, err)
			}

			if pokemon.Name != tc.name {
				t.Errorf("Expected name to be %s, got %s", tc.name, pokemon.Name)
			}

		})
	}
}

func TestGetNatureIntegration(t *testing.T) {
	client := setup()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	testCases := []struct {
		name    string
		wantErr bool
	}{
		{"hardy", false},
		{"nonexistentnature", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			nature, err := client.GetNature(ctx, tc.name)

			if tc.wantErr {
				if err == nil {
					t.Errorf("Expected an error for %s but got none", tc.name)
				}
				return
			}

			if err != nil {
				t.Fatalf("Error fetching Nature '%s': %v", tc.name, err)
			}

			if nature.Name != tc.name {
				t.Errorf("Expected name to be %s, got %s", tc.name, nature.Name)
			}
		})
	}
}

func TestGetStatIntegration(t *testing.T) {
	client := setup()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	testCases := []struct {
		name    string
		wantErr bool
	}{
		{"speed", false},
		{"nonexistentstat", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stat, err := client.GetStat(ctx, tc.name)

			if tc.wantErr {
				if err == nil {
					t.Errorf("Expected an error for %s but got none", tc.name)
				}
				return
			}

			if err != nil {
				t.Fatalf("Error fetching Stat '%s': %v", tc.name, err)
			}

			if stat.Name != tc.name {
				t.Errorf("Expected name to be %s, got %s", tc.name, stat.Name)
			}
		})
	}
}
