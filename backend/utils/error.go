package utils

type (
	PlantedError struct {
		Message string `json:"message"`
	}
)

func (p *PlantedError) Error() string {
	return p.Message
}

func NewServiceError(err error) *PlantedError {
	return &PlantedError{
		Message: "[ServiceError]" + err.Error(),
	}
}

func NewRepositoryError(err error) *PlantedError {
	return &PlantedError{
		Message: "[RepositoryError]" + err.Error(),
	}
}
