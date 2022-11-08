package main_service

import (
	"backend/repo"
	"backend/utils"
	"fmt"
)

func (s SimpleService) findOverlapBetweenCategories(elements map[string]map[string]bool) []*UserType {
	totalOverlaps := make([]*UserType, 0)
	for key1, elements1 := range elements {
		for key2, elements2 := range elements {
			if key1 == key2 {
				continue
			}
			user := &UserType{
				Combination:        fmt.Sprintf("%s->%s", key1, key2),
				FirstItemElements:  elements1,
				SecondItemElements: elements2,
				NumberOfUsers:      0,
			}
			totalOverlaps = append(totalOverlaps, user)
		}
	}

	for _, elem := range totalOverlaps {
		for item, _ := range elem.FirstItemElements {
			if _, ok := elem.SecondItemElements[item]; ok {
				elem.NumberOfUsers = elem.NumberOfUsers + 1
			}
		}
	}

	return totalOverlaps
}

func (s SimpleService) loadData() ([]*repo.Seed, []*repo.Peapolis, []*repo.FruitStar, error) {
	seeds, err := s.SeedRepository.LoadData("seeds.jsonl")
	if err != nil {
		return nil, nil, nil, utils.NewServiceError(err)
	}

	peapolis, err := s.PeapolisRepoistory.LoadData("peapolis.jsonl")
	if err != nil {
		return nil, nil, nil, utils.NewServiceError(err)
	}

	frutitStars, err := s.FruitStarRepository.LoadData("fruitstar.jsonl")
	if err != nil {
		return nil, nil, nil, utils.NewServiceError(err)
	}

	return seeds, peapolis, frutitStars, nil
}
