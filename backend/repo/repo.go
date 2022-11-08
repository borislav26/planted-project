package repo

import (
	"backend/utils"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type (
	PlantedRepository interface {
		TypeName() string
	}

	GenericRepository[T PlantedRepository] interface {
		LoadData(fileName string) ([]*T, error)
	}

	GenericSimpleRepository[T PlantedRepository] struct {
		GenericRepository[T]
	}
)

func (r GenericSimpleRepository[T]) LoadData(fileName string) (elements []*T, err error) {
	workingDir := os.Getenv("WORKING_DIR")
	bytes, err := os.Open(fmt.Sprintf("%s%s", workingDir, fileName))
	if err != nil {
		return nil, utils.NewRepositoryError(err)
	}
	defer bytes.Close()

	fileScanner := bufio.NewScanner(bytes)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		var element *T
		err = json.Unmarshal(fileScanner.Bytes(), &element)
		if err != nil {
			return nil, utils.NewRepositoryError(err)
		}
		elements = append(elements, element)
	}

	return
}

func (f FruitStar) TypeName() string {
	return "FruitStar"
}

func (f Peapolis) TypeName() string {
	return "Peapolis"
}

func (f Seed) TypeName() string {
	return "Seed"
}
