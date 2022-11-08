package main_service

import (
	"backend/utils"
	"sort"
)

func (s SimpleService) TotalOverlap() ([]string, *float64, error) {
	seeds, peapolis, fruitStars, err := s.loadData()
	if err != nil {
		return nil, nil, utils.NewServiceError(err)
	}

	seedsMap := make(map[string]bool)
	peapolisMap := make(map[string]bool)
	fruitStarsMap := make(map[string]bool)

	for _, seed := range seeds {
		seedsMap[seed.Email] = true
	}

	for _, peapol := range peapolis {
		peapolisMap[peapol.Email] = true
	}

	for _, fruitStar := range fruitStars {
		fruitStarsMap[fruitStar.Email] = true
	}

	totalOverlapUsers := make([]string, 0)
	totalOverlapCounter := 0

	for key, _ := range seedsMap {
		if _, peapolisOK := peapolisMap[key]; peapolisOK {
			if _, fruitStarOK := fruitStarsMap[key]; fruitStarOK {
				totalOverlapUsers = append(totalOverlapUsers, key)
				totalOverlapCounter++
			}
		}
	}

	totalUsers := len(seeds) + len(fruitStars) + len(peapolis)
	overlap := (float64(totalOverlapCounter) / float64(totalUsers)) * 100

	return totalOverlapUsers, &overlap, nil
}

func (s SimpleService) AllCombinationsOverlap() ([]*UserType, error) {
	seeds, peapolis, fruitStars, err := s.loadData()
	if err != nil {
		return nil, utils.NewServiceError(err)
	}

	seedsEmails := make(map[string]bool, 0)
	peapolisEmails := make(map[string]bool, 0)
	fruitStarsEmails := make(map[string]bool, 0)

	for _, seed := range seeds {
		seedsEmails[seed.Email] = true
	}

	for _, pea := range peapolis {
		peapolisEmails[pea.Email] = true
	}

	for _, fruit := range fruitStars {
		fruitStarsEmails[fruit.Email] = true
	}

	elements := map[string]map[string]bool{
		"Seeds":      seedsEmails,
		"Peapolis":   peapolisEmails,
		"FruitStars": fruitStarsEmails,
	}

	return s.findOverlapBetweenCategories(elements), nil
}

func (s SimpleService) MostCommonZipcodes() ([]Zipcode, error) {
	seeds, peapolis, fruitStars, err := s.loadData()
	if err != nil {
		return nil, err
	}

	mostCommonZipcodes := make(map[string]int64)

	for _, seed := range seeds {
		if _, ok := mostCommonZipcodes[seed.Postcode]; !ok {
			mostCommonZipcodes[seed.Postcode] = 1
		} else {
			mostCommonZipcodes[seed.Postcode] = mostCommonZipcodes[seed.Postcode] + 1
		}
	}

	for _, peapol := range peapolis {
		if _, ok := mostCommonZipcodes[peapol.Address.Zip]; !ok {
			mostCommonZipcodes[peapol.Address.Zip] = 1
		} else {
			mostCommonZipcodes[peapol.Address.Zip] = mostCommonZipcodes[peapol.Address.Zip] + 1
		}
	}

	for _, fruitStar := range fruitStars {
		if _, ok := mostCommonZipcodes[fruitStar.Address.Zip]; !ok {
			mostCommonZipcodes[fruitStar.Address.Zip] = 1
		} else {
			mostCommonZipcodes[fruitStar.Address.Zip] = mostCommonZipcodes[fruitStar.Address.Zip] + 1
		}
	}

	zipcodeKeyValues := make(ZipcodeArr, 0)

	for key, value := range mostCommonZipcodes {
		zip := Zipcode{
			Key:   key,
			Value: value,
		}
		zipcodeKeyValues = append(zipcodeKeyValues, zip)
	}

	sort.Sort(zipcodeKeyValues)

	return zipcodeKeyValues, nil
}

func (s SimpleService) AllUsersWithNumberOfElements() ([]*UserTypeWithNumberOfElements, error) {
	seeds, peapolis, fruitStars, err := s.loadData()
	if err != nil {
		return nil, utils.NewServiceError(err)
	}
	elements := []*UserTypeWithNumberOfElements{
		{Name: "Seeds", NumberOfElements: len(seeds)},
		{Name: "Peapolis", NumberOfElements: len(peapolis)},
		{Name: "FruitStars", NumberOfElements: len(fruitStars)},
	}

	return elements, nil
}
