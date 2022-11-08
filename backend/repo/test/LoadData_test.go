package test

import (
	"backend/builder"
	"log"
	"testing"
)

func TestLoadData(t *testing.T) {
	t.Setenv("WORKING_DIR", "../../data/")
	repositories := builder.BuildRepositories()
	_, err := repositories.PeapolisRepository.LoadData("peapolis.jsonl")
	if err != nil {
		log.Println(err)
		t.Error("There is doesn't need to be failure")
	}

	_, err = repositories.FruitStarRepository.LoadData("fruitstar.jsonl")
	if err != nil {
		t.Error("There is doesn't need to be failure")
	}

	_, err = repositories.SeedRepository.LoadData("seeds.jsonl")
	if err != nil {
		t.Error("There is doesn't need to be failure")
	}

	_, err = repositories.SeedRepository.LoadData("badfile.json")
	if err == nil {
		t.Error("There is need to be failure")
	}
}
