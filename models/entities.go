package models

type FoodData struct {
	Foods []Food `json:"data"`
}

type Food struct {
	ID       uint16   `json:"id"`
	NAME     string   `json:"name"`
	NUTRIENT Nutrient `json:"nutrients"`
	CATEGORY Category `json:"category"`
}

type Nutrient struct {
	KCAL          uint32  `json:"kcal"`
	KJ            uint32  `json:"kj"`
	PROTEIN       float32 `json:"protein"`
	CARBOHYDRATES float32 `json:"carbohydrates"`
}

type Category struct {
	ID   uint16 `json:"id"`
	NAME string `json:"name"`
}
