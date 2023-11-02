package api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"fmt"
)

// Constants for stats
// Stat names as string identifiers
const (
	StatHP              = "hp"
	StatAttack          = "attack"
	StatDefense         = "defense"
	StatSpecialAttack   = "special-attack"
	StatSpecialDefense  = "special-defense"
	StatSpeed           = "speed"
)

// Stat IDs mapped to names
var statsByID = map[string]string{
	"1": StatHP,
	"2": StatAttack,
	"3": StatDefense,
	"4": StatSpecialAttack,
	"5": StatSpecialDefense,
	"6": StatSpeed,
}

// Nature names as string identifiers
const (
	Hardy    = "hardy"
	Lonely   = "lonely"
	Brave    = "brave"
	Adamant  = "adamant"
	Naughty  = "naughty"
	Bold     = "bold"
	Docile   = "docile"
	Relaxed  = "relaxed"
	Impish   = "impish"
	Lax      = "lax"
	Timid    = "timid"
	Hasty    = "hasty"
	Serious  = "serious"
	Jolly    = "jolly"
	Naive    = "naive"
	Modest   = "modest"
	Mild     = "mild"
	Quiet    = "quiet"
	Bashful  = "bashful"
	Rash     = "rash"
	Calm     = "calm"
	Gentle   = "gentle"
	Sassy    = "sassy"
	Careful  = "careful"
	Quirky   = "quirky"
)

// Nature IDs mapped to nature names
var naturesByID = map[string]string{
	"1":  Hardy,
	"2":  Lonely,
	"3":  Brave,
	"4":  Adamant,
	"5":  Naughty,
	"6":  Bold,
	"7":  Docile,
	"8":  Relaxed,
	"9":  Impish,
	"10": Lax,
	"11": Timid,
	"12": Hasty,
	"13": Serious,
	"14": Jolly,
	"15": Naive,
	"16": Modest,
	"17": Mild,
	"18": Quiet,
	"19": Bashful,
	"20": Rash,
	"21": Calm,
	"22": Gentle,
	"23": Sassy,
	"24": Careful,
	"25": Quirky,
}


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
	ID                     int
	Name                   string
	BaseExperience         int
	Height                 int
	IsDefault              bool
	Order                  int
	Weight                 int
	IsHidden               bool
	Slot                   int
	Ability                NamedURL
	Form                   NamedURL
	Version                NamedURL
	Item                   NamedURL
	LocationAreaEncounters string
	Move                   NamedURL
	Species                NamedURL
	StatDetails            StatDetails
	Type                   NamedURL
	Generation             NamedURL
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

	// Check if idOrName is a valid nature identifier
	if _, ok := naturesByID[idOrName]; !ok {
		// Attempt to find a corresponding ID if a name was provided
		for id, name := range naturesByID {
			if name == idOrName {
				idOrName = id
				ok = true
				break
			}
		}
		// If idOrName is not valid, return an error
		if !ok {
			return nature, fmt.Errorf("invalid nature identifier: %s", idOrName)
		}
	}

	// If idOrName is valid, continue to fetch and unmarshal
	endpoint := "/nature/" + idOrName
	err := fetchAndUnmarshal(c, endpoint, &nature)
	return nature, err
}

func (c *Client) GetStat(ctx context.Context, idOrName string) (Stat, error) {
	var stat Stat

	// Check if the idOrName is a valid stat identifier by name
	if _, ok := statsByID[idOrName]; !ok {
		// If not a valid ID, check if it's a valid name
		valid := false
		for _, name := range []string{StatHP, StatAttack, StatDefense, StatSpecialAttack, StatSpecialDefense, StatSpeed} {
			if idOrName == name {
				valid = true
				break
			}
		}
		if !valid {
			return stat, fmt.Errorf("invalid stat identifier: %s", idOrName)
		}
	} else {
		// If it's a valid ID, get the corresponding name
		idOrName = statsByID[idOrName]
	}

	endpoint := "/stat/" + idOrName
	err := fetchAndUnmarshal(c, endpoint, &stat)
	if err != nil {
		return stat, fmt.Errorf("error fetching stat %s: %v", idOrName, err)
	}
	return stat, nil

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
