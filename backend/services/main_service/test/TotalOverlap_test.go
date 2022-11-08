package test

import (
	"backend/builder"
	"log"
	"testing"
)

func TestTotalOverlap(t *testing.T) {
	t.Setenv("WORKING_DIR", "../../../data/")
	repositories := builder.BuildRepositories()
	services := builder.BuildServices(repositories)

	_, _, err := services.MainService.TotalOverlap()
	if err != nil {
		log.Println(err)
		t.Error("There is doesn't need to be failure")
	}
}
