// Example Usage

package main

import (
	"fmt"
	"log"
	"github.com/ashgodfrey/pokemonapi/api"
	"context"
)

func main() {
	ctx := context.Background()

	pokemonData, err := api.GetPokemon(ctx, "pikachu")
	if err != nil {
		log.Fatalf("Error fetching Pokemon: %v", err)
	}
	fmt.Printf("Pokemon Data: %+v\n", pokemonData)

	statData, err := api.GetStat(ctx, "hp")
	if err != nil {
		log.Fatalf("Error fetching stat: %v", err)
	}
	fmt.Printf("Stat Data: %+v\n", statData)

	natureData, err := api.GetNature(ctx, "bold")
	if err != nil {
		log.Fatalf("Error fetching nature: %v", err)
	}
	fmt.Printf("Nature Data: %+v\n", natureData)
}
