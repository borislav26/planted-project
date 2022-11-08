package test

import (
	"backend/builder"
	"log"
	"testing"
)

func TestAllUsersWithNumberOfElements(t *testing.T) {
	t.Setenv("WORKING_DIR", "../../../data/")
	repositories := builder.BuildRepositories()
	services := builder.BuildServices(repositories)

	_, err := services.MainService.AllUsersWithNumberOfElements()
	if err != nil {
		log.Println(err)
		t.Error("There is doesn't need to be failure")
	}
}
