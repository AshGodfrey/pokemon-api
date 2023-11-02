package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	HTTPClient *http.Client
	Endpoint   string
}

type NamedURL struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Structs related to Pokemon Nature
type PokeathlonStatChange struct {
	MaxChange      int      `json:"max_change"`
	PokeathlonStat NamedURL `json:"pokeathlon_stat"`
}

type MoveBattleStylePreference struct {
	LowHPPreference  int      `json:"low_hp_preference"`
	HighHPPreference int      `json:"high_hp_preference"`
	MoveBattleStyle  NamedURL `json:"move_battle_style"`
}

type NatureName struct {
	Name     string   `json:"name"`
	Language NamedURL `json:"language"`
}

type Nature struct {
	ID                         int                         `json:"id"`
	Name                       string                      `json:"name"`
	DecreasedStat              NamedURL                    `json:"decreased_stat"`
	IncreasedStat              NamedURL                    `json:"increased_stat"`
	LikesFlavor                NamedURL                    `json:"likes_flavor"`
	HatesFlavor                NamedURL                    `json:"hates_flavor"`
	PokeathlonStatChanges      []PokeathlonStatChange      `json:"pokeathlon_stat_changes"`
	MoveBattleStylePreferences []MoveBattleStylePreference `json:"move_battle_style_preferences"`
	Names                      []NatureName                `json:"names"`
}

// Structs related to Stat
type MoveEffect struct {
	Change int      `json:"change"`
	Move   NamedURL `json:"move"`
}

type AffectingMoves struct {
	Increase []MoveEffect `json:"increase"`
	Decrease []MoveEffect `json:"decrease"`
}

type AffectingNatures struct {
	Increase []NamedURL `json:"increase"`
	Decrease []NamedURL `json:"decrease"`
}

type Characteristic struct {
	URL string `json:"url"`
}

// Stat represents the details of a specific stat for a Pokémon as defined by the Pokémon API.
type Stat struct {
	ID               int              `json:"id"`
	Name             string           `json:"name"`
	GameIndex        int              `json:"game_index"`
	IsBattleOnly     bool             `json:"is_battle_only"`
	AffectingMoves   AffectingMoves   `json:"affecting_moves"`
	AffectingNatures AffectingNatures `json:"affecting_natures"`
	Characteristics  []Characteristic `json:"characteristics"`
	MoveDamageClass  NamedURL         `json:"move_damage_class"`
	Names            []NatureName     `json:"names"`
}

// Struct for Pokemon details
type StatDetails struct {
	BaseStat int      `json:"base_stat"`
	Effort   int      `json:"effort"`
	StatInfo NamedURL `json:"stat_info"`
}

// Pokemon represents the details of a Pokémon as defined by the Pokémon API.
type Pokemon struct {
	ID                     int            `json:"id"`
	Name                   string         `json:"name"`
	BaseExperience         int            `json:"base_experience"`
	Height                 int            `json:"height"`
	IsDefault              bool           `json:"is_default"`
	Order                  int            `json:"order"`
	Weight                 int            `json:"weight"`
	IsHidden               bool           `json:"is_hidden"`
	Slot                   int            `json:"slot"`
	Ability                NamedURL       `json:"ability"`
	Form                   NamedURL       `json:"form"`
	Version                NamedURL       `json:"version"`
	Item                   NamedURL       `json:"item"`
	LocationAreaEncounters EncountersData `json:"location_area_encounters"`
	Move                   NamedURL       `json:"move"`
	Species                NamedURL       `json:"species"`
	StatDetails            []StatDetails  `json:"stats"`
	Type                   NamedURL       `json:"type"`
	Generation             NamedURL       `json:"generation"`
}

type LocationAreaEncounter struct {
	LocationArea   NamedURL        `json:"location_area"`
	VersionDetails []VersionDetail `json:"version_details"`
}

type EncounterMethod struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type ConditionValue struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type EncounterDetail struct {
	MinLevel        int              `json:"min_level"`
	MaxLevel        int              `json:"max_level"`
	ConditionValues []ConditionValue `json:"condition_values"`
	Chance          int              `json:"chance"`
	Method          EncounterMethod  `json:"method"`
}

type VersionDetail struct {
	MaxChance        int               `json:"max_chance"`
	EncounterDetails []EncounterDetail `json:"encounter_details"`
	Version          NamedURL          `json:"version"`
}

// GetPokemonOpts contains options for GetPokemon function.
type GetPokemonOpts struct {
	// IDorName is the name or ID of the Pokemon to Retrieve.
	IDOrName string
	// IncludeLocation is an optional value to get the location area encounters data.
	IncludeLocation bool
}

// GetNatureOpts contains options for GetNature function.
type GetNatureOpts struct {
	IDOrName string // The ID or name of the Nature to retrieve.
}

// GetStatOpts contains options for GetStat function.
type GetStatOpts struct {
	IDOrName string // The ID or name of the Stat to retrieve.
}

func NewClient() *Client {
	return &Client{
		HTTPClient: http.DefaultClient,
		Endpoint:   "https://pokeapi.co/api/v2",
	}
}

type EncountersData string

func (e *EncountersData) UnmarshalJSON(data []byte) error {
	// First, try to unmarshal the data into a string
	var encounterURL string
	if err := json.Unmarshal(data, &encounterURL); err == nil {
		*e = EncountersData(encounterURL)
		return nil
	}

	// Next, try to unmarshal the data into a slice of LocationAreaEncounter
	var encounterDetails []LocationAreaEncounter
	if err := json.Unmarshal(data, &encounterDetails); err == nil {
		detailsJSON, err := json.Marshal(encounterDetails)
		if err != nil {
			return err
		}
		*e = EncountersData(detailsJSON)
		return nil
	}

	// If neither unmarshalling is successful, return an error
	return fmt.Errorf("location_area_encounters contains neither a URL string nor an array")
}

// In the GetPokemon function, make sure to convert EncountersData to string when necessary:
func (c *Client) GetPokemon(ctx context.Context, opts GetPokemonOpts) (Pokemon, error) {
	var pokemon Pokemon
	err := fetchAndUnmarshal(c, "/pokemon/"+opts.IDOrName, &pokemon)
	if err != nil {
		return pokemon, err
	}

	// If IncludeLocation is true, make an additional API call for location details
	if opts.IncludeLocation {
		var locationDetails []LocationAreaEncounter
		err = fetchAndUnmarshal(c, fmt.Sprintf("/pokemon/%s/encounters", opts.IDOrName), &locationDetails)
		if err != nil {
			return pokemon, err
		}
		// Since the locationDetails is already a slice, we can marshal it directly to the EncountersData type
		locationDetailsJSON, err := json.Marshal(locationDetails)
		if err != nil {
			return pokemon, fmt.Errorf("error marshalling encounters details: %w", err)
		}
		pokemon.LocationAreaEncounters = EncountersData(locationDetailsJSON)
	}

	// No need to handle the else block because the UnmarshalJSON method already does the necessary conversion

	return pokemon, nil
}

func (c *Client) GetNature(ctx context.Context, idOrName string) (Nature, error) {
	var nature Nature
	err := fetchAndUnmarshal(c, "/nature/"+idOrName, &nature)
	return nature, err
}

func (c *Client) GetStat(ctx context.Context, idOrName string) (Stat, error) {
	var stat Stat
	err := fetchAndUnmarshal(c, "/stat/"+idOrName, &stat)
	return stat, err
}

func (c *Client) fetchData(endpoint string) ([]byte, error) {
	resp, err := c.HTTPClient.Get(c.Endpoint + endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// Generic fetch function
func fetchAndUnmarshal[T any](c *Client, endpoint string, dest *T) error {
	body, err := c.fetchData(endpoint)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, dest)
}
