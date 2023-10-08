package parse

import (
	"encoding/json"
	"example/go-taco/models"
	"io"
	"os"
)

func GetFoodData(f *os.File) (data models.FoodData, err error) {
	defer f.Close()

	byteValue, err := io.ReadAll(f)
	if err != nil {
		return data, err
	}

	if err := json.Unmarshal(byteValue, &data); err != nil {
		return data, err
	}

	return data, nil
}
