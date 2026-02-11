package pokeapi

import (
	"io"
	"net/http"
	"time"

	"github.com/realdnchka/pokedexcli/internal/pokecache"
)

type Pagination struct {
	Next string
	Prev string
}

//URLs
const baseUrl = "https://pokeapi.co/api/v2/"

const locationAreaUrl = baseUrl + "/location-area/"
const locationAriaBaseOffset = "?offset=0&limit=20"

const pokemonUrl = baseUrl + "/pokemon/"

//Cache init
var CacheInterval = 10 * time.Minute
var cache = pokecache.NewCache(CacheInterval)

//Pagination init
var MapPagination = Pagination {
	Next: locationAreaUrl + locationAriaBaseOffset,
	Prev: locationAreaUrl + locationAriaBaseOffset,
}

/*
Get list areas as AreasListDTO from Pagination.Next or Paginator.Prev URLs
*/
func GetAreas(forward bool) (areas AreasListDTO, err error) {
	var url string

	//Checks if user needs next or previous maps page
	if forward {
		url = MapPagination.Next
	} else {
		url = MapPagination.Prev
	}

	//Checks if url as key in cache
	ce, err := get(url)
	if err != nil {
		return areas, err
	}
	
	//Convert to AreasListDTO
	areas, err = bytesToDTO[AreasListDTO](ce)
	if err != nil {
		return areas, err
	}
	
	setPagination(areas)
	return areas, nil
}

/*
Get list of Pokemons as PokemonsInAreaDTO from locationAreaUrl + <area_name> URL
*/
func GetAreaPokemons(a string) (poks PokemonsInAreaDTO, err error) {
	url := locationAreaUrl + a

	ce, err := get(url)
	if err != nil {
		return poks, err
	}

	poks, err = bytesToDTO[PokemonsInAreaDTO](ce)
	if err != nil {
		return poks, err
	}
	return poks, nil
}

func GetPokemon(a string) (poks PokemonDTO, err error) {
	url := pokemonUrl + a

	ce, err := get(url)
	if err != nil {
		return poks, err
	}

	poks, err = bytesToDTO[PokemonDTO](ce)
	if err != nil {
		return poks, err
	}
	return poks, nil
}
/*
Wraps 'net/http' library and creating cache entry
Returns JSON compatible []byte
*/
func get(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	cache.Add(url, data)
	return data, nil
}

/*
Sets page pagination for GetAreas()
*/
func setPagination(dto AreasListDTO) {
	MapPagination.Next = dto.Next
	MapPagination.Prev = dto.Prev
}