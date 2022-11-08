package repo

type (
	FruitStar struct {
		UUID string `json:"uuid"`
		FirstAndLastName
		User
		Address Address `json:"address"`
	}

	FirstAndLastName struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	BaseAddress struct {
		Street  string `json:"street"`
		Country string `json:"country"`
	}

	Address struct {
		BaseAddress
		HouseNumber string `json:"house_number"`
		City        string `json:"city"`
		Zip         string `json:"zip"`
	}

	User struct {
		Email     string `json:"email"`
		LastLogin string `json:"last_login"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		DeletedAt string `json:"deleted_at"`
	}

	Peapolis struct {
		id   string `json:"id"`
		Name string `json:"name"`
		User
		Address Address `json:"address"`
	}

	Seed struct {
		id       int64  `json:"id"`
		Postcode string `json:"postcode"`
		BaseAddress
		User
		FirstAndLastName
	}
)
