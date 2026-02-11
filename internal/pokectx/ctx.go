package pokectx

type Pokemon struct {
	ID int
	Name string
	BaseEXP int
	Stats []PokemonStat
	Types []PokemonType
	Weight int
	Height int
}

type PokemonType struct {
	Name string
}

type PokemonStat struct {
	Name string
	Value int
}

type Context struct {
	Database *DB
}