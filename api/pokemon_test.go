package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockPokemonServer initializes a new httptest.Server with a handler that
// fakes the GetPokemon response.
func mockPokemonServer() *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := Pokemon{
			ID:                     25,
			Name:                   "pikachu",
			BaseExperience:         112,
			Height:                 4,
			IsDefault:              true,
			Order:                  1,
			Weight:                 60,
			Ability:                NamedURL{Name: "static", URL: "/ability/static"},
			Form:                   NamedURL{Name: "pikachu", URL: "/pokemon-form/pikachu"},
			Version:                NamedURL{Name: "red", URL: "/version/red"},
			Item:                   NamedURL{Name: "light-ball", URL: "/item/light-ball"},
			LocationAreaEncounters: "URL to location area encounters",
			Move:                   NamedURL{Name: "thunder-shock", URL: "/move/thunder-shock"},
			Species:                NamedURL{Name: "pikachu", URL: "/pokemon-species/pikachu"},
			StatDetails: []StatDetails{
				{
					BaseStat: 35,
					Effort:   0,
					StatInfo: NamedURL{Name: "speed", URL: "/stat/6"},
				},
			},
			Type:       NamedURL{Name: "electric", URL: "/type/electric"},
			Generation: NamedURL{Name: "generation-i", URL: "/generation/i"},
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	})

	return httptest.NewServer(handler)
}

func TestAPIGetPokemon(t *testing.T) {
	server := mockPokemonServer()
	defer server.Close()

	DefaultClient.HTTPClient = server.Client()
	DefaultClient.Endpoint = server.URL

	expectedPokemon := Pokemon{
		ID:                     25,
		Name:                   "pikachu",
		BaseExperience:         112,
		Height:                 4,
		IsDefault:              true,
		Order:                  1,
		Weight:                 60,
		Ability:                NamedURL{Name: "static", URL: "/ability/static"},
		Form:                   NamedURL{Name: "pikachu", URL: "/pokemon-form/pikachu"},
		Version:                NamedURL{Name: "red", URL: "/version/red"},
		Item:                   NamedURL{Name: "light-ball", URL: "/item/light-ball"},
		LocationAreaEncounters: "URL to location area encounters",
		Move:                   NamedURL{Name: "thunder-shock", URL: "/move/thunder-shock"},
		Species:                NamedURL{Name: "pikachu", URL: "/pokemon-species/pikachu"},
		StatDetails: []StatDetails{
			{
				BaseStat: 35,
				Effort:   0,
				StatInfo: NamedURL{Name: "speed", URL: "/stat/6"},
			},
		},
		Type:       NamedURL{Name: "electric", URL: "/type/electric"},
		Generation: NamedURL{Name: "generation-i", URL: "/generation/i"},
	}

	pokemon, err := GetPokemon(context.Background(), GetPokemonOpts{
		Name: "Pikachu",
	})
	require.NoError(t, err)
	require.Equal(t, expectedPokemon, pokemon)
}

// mockNatureServer initializes a new httptest.Server with a handler that
// fakes the GetNature response.
func mockNatureServer() *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := Nature{
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
		}
		// Serialize the response to JSON
		responseData, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Write the response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseData)
	})

	return httptest.NewServer(handler)
}

func TestAPIGetNature(t *testing.T) {
	server := mockNatureServer()
	defer server.Close()

	DefaultClient.HTTPClient = server.Client()
	DefaultClient.Endpoint = server.URL

	expectedNature := Nature{
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
	}

	nature, err := GetNature(context.Background(), GetNatureOpts{Name: "hardy"})

	require.NoError(t, err)
	require.Equal(t, expectedNature, nature)
}

// mockStatServer initializes a new httptest.Server with a handler that
// fakes the GetStat response.
func mockStatServer() *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := Stat{
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
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	})

	return httptest.NewServer(handler)
}

func TestAPIGetStat(t *testing.T) {
	server := mockStatServer()
	defer server.Close()

	DefaultClient.HTTPClient = server.Client()
	DefaultClient.Endpoint = server.URL

	expectedStat := Stat{
		ID:           6,
		Name:         "speed",
		GameIndex:    1,
		IsBattleOnly: false,
		AffectingMoves: AffectingMoves{
			Increase: []MoveEffect{
				{
					Change: 1,
					Move: NamedURL{
						Name: "agility",
						URL:  "/move/agility",
					},
				},
			},
			Decrease: []MoveEffect{
				{
					Change: -1,
					Move: NamedURL{
						Name: "paralyze",
						URL:  "/move/paralyze",
					},
				},
			},
		},
		AffectingNatures: AffectingNatures{
			Increase: []NamedURL{
				{
					Name: "jolly",
					URL:  "/nature/jolly",
				},
			},
			Decrease: []NamedURL{
				{
					Name: "relaxed",
					URL:  "/nature/relaxed",
				},
			},
		},
		Characteristics: []Characteristic{
			{
				URL: "/characteristic/high-speed",
			},
		},
		MoveDamageClass: NamedURL{
			Name: "physical",
			URL:  "/move-damage-class/physical",
		},
		Names: []NatureName{
			{
				Name: "Speed",
				Language: NamedURL{
					Name: "en",
					URL:  "/language/en",
				},
			},
		},
	}

	stat, err := GetStat(context.Background(), GetStatOpts{Name: "speed"})

	require.NoError(t, err)
	require.Equal(t, expectedStat, stat)
}
