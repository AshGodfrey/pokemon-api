package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPokemon(t *testing.T) {
	tests := []struct {
		scenario    string
		client      *Client
		pokemonName string
		expected    Pokemon
		err         bool
	}{
		{
			scenario:    "Successful retrieval of Pikachu",
			client:      &Client{},
			pokemonName: "pikachu",
			expected: Pokemon{
				ID:                     25,
				Name:                   "pikachu",
				BaseExperience:         112,
				Height:                 4,
				IsDefault:              true,
				Order:                  1,
				Weight:                 60,
				Ability:                NamedURL{Name: "static", URL: "/ability/7"},
				Form:                   NamedURL{Name: "pikachu", URL: "/pokemon-form/25"},
				Version:                NamedURL{Name: "red", URL: "/version/1"},
				Item:                   NamedURL{Name: "light-ball", URL: "/item/213"},
				LocationAreaEncounters: "URL to location area encounters",
				Move:                   NamedURL{Name: "thunder-shock", URL: "/move/84"},
				Species:                NamedURL{Name: "pikachu", URL: "/pokemon-species/25"},
				StatDetails: []StatDetails{
					{
						BaseStat: 35,
						Effort:   0,
						StatInfo: NamedURL{Name: "speed", URL: "/stat/6"},
					},
				},
				Type:       NamedURL{Name: "electric", URL: "/type/13"},
				Generation: NamedURL{Name: "generation-i", URL: "/generation/1"},
			},
			err: false,
		},
		{
			scenario:    "Failed retrieval, Pokemon does not exist",
			client:      &Client{},
			pokemonName: "unknown",
			expected:    Pokemon{},
			err:         true,
		},
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if test.err {
					http.Error(w, "Pokemon not found", http.StatusNotFound)
					return
				}
				json.NewEncoder(w).Encode(test.expected)
			}))
			defer server.Close()

			client := test.client
			client.Endpoint = server.URL
			client.HTTPClient = server.Client()

			pokemon, err := client.GetPokemon(context.Background(), GetPokemonOpts{
				Name: test.pokemonName,
			})
			if test.err {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, test.expected, pokemon)
		})
	}
}

func TestGetNature(t *testing.T) {
	tests := []struct {
		scenario   string
		natureName string
		expected   Nature
		err        bool
	}{
		{
			scenario:   "Successful retrieval of a nature",
			natureName: "adamant",
			expected: Nature{
				ID:            1,
				Name:          "hardy",
				DecreasedStat: NamedURL{Name: "speed", URL: "/stat/speed"},
				IncreasedStat: NamedURL{Name: "attack", URL: "/stat/attack"},
				LikesFlavor:   NamedURL{Name: "spicy", URL: "/flavor/spicy"},
				HatesFlavor:   NamedURL{Name: "sweet", URL: "/flavor/sweet"},
				PokeathlonStatChanges: []PokeathlonStatChange{
					{
						MaxChange:      2,
						PokeathlonStat: NamedURL{Name: "speed", URL: "/pokeathlon/speed"},
					},
				},
				MoveBattleStylePreferences: []MoveBattleStylePreference{
					{
						LowHPPreference:  5,
						HighHPPreference: 10,
						MoveBattleStyle:  NamedURL{Name: "aggressive", URL: "/style/aggressive"},
					},
				},
				Names: []NatureName{
					{
						Name:     "Hardy",
						Language: NamedURL{Name: "en", URL: "/language/en"},
					},
				},
			},
			err: false,
		},
		{
			scenario:   "Nature not found",
			natureName: "unknown",
			expected:   Nature{},
			err:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.scenario, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if tt.err {
					http.Error(w, "Nature not found", http.StatusNotFound)
				} else {
					json.NewEncoder(w).Encode(tt.expected)
				}
			}))
			defer server.Close()

			client := &Client{
				HTTPClient: server.Client(),
				Endpoint:   server.URL,
			}

			nature, err := client.GetNature(context.Background(), GetNatureOpts{
				Name: tt.natureName,
			})
			if tt.err {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, nature)
			}
		})
	}
}

func TestGetStat(t *testing.T) {
	tests := []struct {
		scenario string
		statName string
		expected Stat
		err      bool
	}{
		{
			scenario: "Successful retrieval of a stat",
			statName: "speed",
			expected: Stat{
				ID:           6,
				Name:         "speed",
				GameIndex:    1,
				IsBattleOnly: false,
				AffectingMoves: AffectingMoves{
					Increase: []MoveEffect{
						{
							Change: 1,
							Move:   NamedURL{Name: "agility", URL: "/move/agility"},
						},
					},
					Decrease: []MoveEffect{
						{
							Change: -1,
							Move:   NamedURL{Name: "paralyze", URL: "/move/paralyze"},
						},
					},
				},
				AffectingNatures: AffectingNatures{
					Increase: []NamedURL{
						{Name: "jolly", URL: "/nature/jolly"},
					},
					Decrease: []NamedURL{
						{Name: "relaxed", URL: "/nature/relaxed"},
					},
				},
				Characteristics: []Characteristic{
					{URL: "/characteristic/high-speed"},
				},
				MoveDamageClass: NamedURL{Name: "physical", URL: "/move-damage-class/physical"},
				Names: []NatureName{
					{Name: "Speed", Language: NamedURL{Name: "en", URL: "/language/en"}},
				},
			},
			err: false,
		},
		{
			scenario: "Stat not found",
			statName: "unknown",
			expected: Stat{},
			err:      true,
		},
		// Additional test cases...
	}

	for _, tt := range tests {
		t.Run(tt.scenario, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if tt.err {
					http.Error(w, "Stat not found", http.StatusNotFound)
				} else {
					json.NewEncoder(w).Encode(tt.expected)
				}
			}))
			defer server.Close()

			client := &Client{
				HTTPClient: server.Client(),
				Endpoint:   server.URL,
			}

			stat, err := client.GetStat(context.Background(), GetStatOpts{
				Name: tt.statName})
			if tt.err {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, stat)
			}
		})
	}
}
