package pokeapi

import (
	"encoding/json"
	"fmt"
)

type AreasListDTO struct {
	Areas []AreaDTO `json:"results"`
	Next string `json:"next"`
	Prev string `json:"previous"`
}

type AreaDTO struct {
	Name string `json:"name"`
}

type PokemonsInAreaDTO struct {
	PokemonEncounters []struct {
		Pokemon NameDTO `json:"pokemon"`	
	} `json:"pokemon_encounters"`
}

type NameDTO struct {
	Name string `json:"name"`
}

type PokemonDTO struct {
	NameDTO
	BaseEXP int `json:"base_experience"`
	ID int `json:"id"`
	Height int `json:"height"`
	Weight int `json:"weight"`
	PokemonStatsDTO
	PokemonTypesDTO
}

type PokemonStatsDTO struct {
	Stats []PokemonStatDTO `json:"stats"`
}

type PokemonTypesDTO struct {
	Types []PokemonTypeDTO `json:"types"`
}

type PokemonStatDTO struct {
	Value int `json:"base_stat"`
	NameDTO `json:"stat"`
}

type PokemonTypeDTO struct {
	NameDTO `json:"type"`
}

/*
Unmarshals JSON-compatible []byte to DTO
*/
func bytesToDTO[T any](b []byte) (T, error) {
	var dto T
	if err := json.Unmarshal(b, &dto); err != nil {
		return dto, fmt.Errorf("unmarshall problem")
	}
	return dto, nil
}