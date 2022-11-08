package main_service

import (
	"backend/repo/fruitstar"
	"backend/repo/peapolis"
	"backend/repo/seed"
)

type (
	Service interface {
		TotalOverlap() ([]string, *float64, error)
		MostCommonZipcodes() ([]Zipcode, error)
		AllCombinationsOverlap() ([]*UserType, error)
		AllUsersWithNumberOfElements() ([]*UserTypeWithNumberOfElements, error)
	}

	SimpleService struct {
		FruitStarRepository fruitstar.Repository
		PeapolisRepoistory  peapolis.Repository
		SeedRepository      seed.Repository
	}
)
