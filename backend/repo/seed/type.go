package seed

import "backend/repo"

type (
	Repository interface {
		repo.GenericRepository[repo.Seed]
	}

	SimpleRepository struct {
		repo.GenericSimpleRepository[repo.Seed]
	}
)
