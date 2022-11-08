package peapolis

import "backend/repo"

type (
	Repository interface {
		repo.GenericRepository[repo.Peapolis]
	}

	SimpleRepository struct {
		repo.GenericSimpleRepository[repo.Peapolis]
	}
)
