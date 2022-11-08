package test

import (
	"backend/builder"
	"log"
	"testing"
)

func TestAllCombinagionsOverlap(t *testing.T) {
	t.Setenv("WORKING_DIR", "../../../data/")
	repositories := builder.BuildRepositories()
	services := builder.BuildServices(repositories)

	_, err := services.MainService.AllCombinationsOverlap()
	if err != nil {
		log.Println(err)
		t.Error("There is doesn't need to be failure")
	}
}
