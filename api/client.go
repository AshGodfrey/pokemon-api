package api

import (
	"context"
	"encoding/json"
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
	ID                         int                          `json:"id"`
	Name                       string                       `json:"name"`
	DecreasedStat              NamedURL                     `json:"decreased_stat"`
	IncreasedStat              NamedURL                     `json:"increased_stat"`
	LikesFlavor                NamedURL                     `json:"likes_flavor"`
	HatesFlavor                NamedURL                     `json:"hates_flavor"`
	PokeathlonStatChanges      []PokeathlonStatChange       `json:"pokeathlon_stat_changes"`
	MoveBattleStylePreferences []MoveBattleStylePreference  `json:"move_battle_style_preferences"`
	Names                      []NatureName                 `json:"names"`
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
	ID               int                 `json:"id"`
	Name             string              `json:"name"`
	GameIndex        int                 `json:"game_index"`
	IsBattleOnly     bool                `json:"is_battle_only"`
	AffectingMoves   AffectingMoves      `json:"affecting_moves"`
	AffectingNatures AffectingNatures    `json:"affecting_natures"`
	Characteristics  []Characteristic    `json:"characteristics"`
	MoveDamageClass  NamedURL            `json:"move_damage_class"`
	Names            []NatureName        `json:"names"`
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
	LocationAreaEncounters string         `json:"location_area_encounters"`
	Move                   NamedURL       `json:"move"`
	Species                NamedURL       `json:"species"`
	StatDetails            []StatDetails  `json:"stats"`
	Type                   NamedURL       `json:"type"`
	Generation             NamedURL       `json:"generation"`
}

// GetPokemonOpts contains options for GetPokemon function.
type GetPokemonOpts struct {
    IDOrName string // The ID or name of the Pokemon to retrieve.
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

func (c *Client) GetPokemon(ctx context.Context, idOrName string) (Pokemon, error) {
	var pokemon Pokemon
	err := fetchAndUnmarshal(c, "/pokemon/"+idOrName, &pokemon)
	return pokemon, err
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

