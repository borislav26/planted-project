package builder

import (
	"backend/repo/fruitstar"
	"backend/repo/peapolis"
	"backend/repo/seed"
)

type Repositories struct {
	FruitStarRepository fruitstar.Repository
	PeapolisRepository  peapolis.Repository
	SeedRepository      seed.Repository
}

func BuildRepositories() *Repositories {
	return &Repositories{
		FruitStarRepository: &fruitstar.SimpleRepository{},
		PeapolisRepository:  &peapolis.SimpleRepository{},
		SeedRepository:      &seed.SimpleRepository{},
	}
}
