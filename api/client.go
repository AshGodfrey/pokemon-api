package api

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"context"
)

type Client struct {
	HTTPClient *http.Client
	Endpoint string
}



type NamedURL struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Structs related to Pokemon Nature
type PokeathlonStatChange struct {
	MaxChange      int       `json:"max_change"`
	PokeathlonStat NamedURL  `json:"pokeathlon_stat"`
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
	ID                         int
	Name                       string
	DecreasedStat              NamedURL
	IncreasedStat              NamedURL
	LikesFlavor                NamedURL
	HatesFlavor                NamedURL
	PokeathlonStatChanges      []PokeathlonStatChange
	MoveBattleStylePreferences []MoveBattleStylePreference
	Names                      []NatureName
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

type Stat struct {
	ID               int
	Name             string
	GameIndex        int
	IsBattleOnly     bool
	AffectingMoves   AffectingMoves
	AffectingNatures AffectingNatures
	Characteristics  []Characteristic
	MoveDamageClass  NamedURL
	Names            []NatureName
}

// Struct for Pokemon details
type StatDetails struct {
	BaseStat int      `json:"base_stat"`
	Effort   int      `json:"effort"`
	StatInfo NamedURL `json:"stat_info"`
}

type Pokemon struct {
	ID                    int
	Name                  string
	BaseExperience        int
	Height                int
	IsDefault             bool
	Order                 int
	Weight                int
	IsHidden              bool
	Slot                  int
	Ability               NamedURL
	Form                  NamedURL
	Version               NamedURL
	Item                  NamedURL
	LocationAreaEncounters string
	Move                  NamedURL
	Species               NamedURL
	StatDetails           StatDetails
	Type                  NamedURL
	Generation            NamedURL
}

func NewClient() *Client {
	return &Client{
		HTTPClient: http.DefaultClient,
		Endpoint: "https://pokeapi.co/api/v2",
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

