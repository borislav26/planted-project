package main_service

type (
	UserType struct {
		Combination        string          `json:"name"`
		FirstItemElements  map[string]bool `json:"firstItemElements"`
		SecondItemElements map[string]bool `json:"secondItemElements"`
		NumberOfUsers      int64           `json:"numberOfUsers"`
		Percentage         float64         `json:"percentage"`
	}

	UserTypeWithNumberOfElements struct {
		Name             string `json:"name"`
		NumberOfElements int    `json:"numberOfElements"`
	}

	Zipcode struct {
		Key   string `json:"key"`
		Value int64  `json:"value"`
	}

	ZipcodeArr []Zipcode
)

func (z ZipcodeArr) Len() int {
	return len(z)
}

func (z ZipcodeArr) Swap(i, j int) {
	z[i], z[j] = z[j], z[i]
}

func (z ZipcodeArr) Less(i, j int) bool {
	return z[i].Value > z[j].Value
}
