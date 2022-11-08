package fruitstar

import "backend/repo"

type (
	Repository interface {
		repo.GenericRepository[repo.FruitStar]
	}

	SimpleRepository struct {
		repo.GenericSimpleRepository[repo.FruitStar]
	}
)
