package builder

import "backend/services/main_service"

type Services struct {
	MainService main_service.Service
}

func BuildServices(repositories *Repositories) *Services {
	mainService := &main_service.SimpleService{
		FruitStarRepository: repositories.FruitStarRepository,
		PeapolisRepoistory:  repositories.PeapolisRepository,
		SeedRepository:      repositories.SeedRepository,
	}

	return &Services{
		MainService: mainService,
	}
}
