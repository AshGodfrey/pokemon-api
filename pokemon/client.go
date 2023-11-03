package pokemon

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"net/url"
)

type Client struct {
	HTTPClient *http.Client
	Endpoint   string
}

type NamedURL struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

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

// Nature represents the details of a specific nature as defined by the Pokémon pokemon.
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

// Stat represents the details of a specific stat for a Pokémon as defined by the Pokémon pokemon.
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

type StatDetails struct {
	BaseStat int      `json:"base_stat"`
	Effort   int      `json:"effort"`
	StatInfo NamedURL `json:"stat_info"`
}

// Pokemon represents the details of a Pokémon as defined by the Pokémon pokemon.
type Pokemon struct {
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	BaseExperience int      `json:"base_experience"`
	Height         int      `json:"height"`
	IsDefault      bool     `json:"is_default"`
	Order          int      `json:"order"`
	Weight         int      `json:"weight"`
	IsHidden       bool     `json:"is_hidden"`
	Slot           int      `json:"slot"`
	Ability        NamedURL `json:"ability"`
	Form           NamedURL `json:"form"`
	Version        NamedURL `json:"version"`
	Item           NamedURL `json:"item"`
	// LocationAreaEncounters will return an array if IncludeLocation is true, otherwise it will return a string URL.
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
	// ID is the ID of the Pokemon to retrieve. Only name or ID needs to be included.
	ID int
	// Name is the name of the Pokemon to retrieve.
	Name string
	// IncludeLocation is an optional value to get the location area encounters data.
	IncludeLocation bool
}

// GetNatureOpts contains options for GetNature function.
type GetNatureOpts struct {
	// ID is the ID of the Nature to retrieve. Only name or ID needs to be included.
	ID int
	// Name is the name of the Nature to retrieve.
	Name string
}

// GetStatOpts contains options for GetStat function.
type GetStatOpts struct {
	// ID is the ID of the Stat to retrieve. Only name or ID needs to be included.
	ID int
	// Name is the name of the Pokemon to retrieve.
	Name string
}

func NewClient() *Client {
	return &Client{
		HTTPClient: http.DefaultClient,
		Endpoint:   "https://pokeapi.co/api/v2",
	}
}

type EncountersData string

// helper function to determine the lookup value
func getLookupValue(id int, name string) (string, error) {
	if id != 0 && name != "" {
		return "", errors.New("you must provide either an ID or a Name, not both")
	}
	if id != 0 {
		return strconv.Itoa(id), nil
	}
	if name != "" {
		return strings.ToLower(name), nil
	}
	return "", errors.New("you must provide either an ID or a Name")
}

// GetPokemon gets a Pokemon by ID or Name.
func (c *Client) GetPokemon(ctx context.Context, opts GetPokemonOpts) (Pokemon, error) {
	var pokemon Pokemon
	lookupValue, err := getLookupValue(opts.ID, opts.Name)
	if err != nil {
		return pokemon, err
	}

	err = fetchAndUnmarshal(c, "pokemon/"+lookupValue, &pokemon)
	if err != nil {
		return pokemon, err
	}

	// If IncludeLocation is true, make an additional API call for location details
	if opts.IncludeLocation {
		var locationDetails []LocationAreaEncounter
		err = fetchAndUnmarshal(c, fmt.Sprintf("pokemon/%s/encounters", lookupValue), &locationDetails)
		if err != nil {
			return pokemon, err
		}
		locationDetailsJSON, err := json.Marshal(locationDetails)
		if err != nil {
			return pokemon, fmt.Errorf("error marshalling encounters details: %w", err)
		}
		pokemon.LocationAreaEncounters = EncountersData(locationDetailsJSON)
	}
	return pokemon, nil
}

// GetNature gets a nature by ID or Name.
func (c *Client) GetNature(ctx context.Context, opts GetNatureOpts) (Nature, error) {
	var nature Nature
	lookupValue, err := getLookupValue(opts.ID, opts.Name)
	if err != nil {
		return nature, err
	}
	err = fetchAndUnmarshal(c, "nature/"+lookupValue, &nature)
	return nature, err
}

// GetStat gets a stat by ID or Name.
func (c *Client) GetStat(ctx context.Context, opts GetStatOpts) (Stat, error) {
	var stat Stat
	lookupValue, err := getLookupValue(opts.ID, opts.Name)
	if err != nil {
		return stat, err
	}
	err = fetchAndUnmarshal(c, "stat/"+lookupValue, &stat)
	return stat, err
}
func fetchAndUnmarshal[T any](c *Client, parameters string, dest *T) error {
    // Parse the base URL and resolve the parameters
    finalURL, err := url.Parse(c.Endpoint)
    if err != nil {
        return err
    }

    finalURL, err = finalURL.Parse(parameters)
    if err != nil {
        return err
    }

    // Make the HTTP GET request
    resp, err := c.HTTPClient.Get(finalURL.String())
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusNotFound {
        return errors.New("not found")
    }

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return err
    }

    return json.Unmarshal(body, dest)
}
