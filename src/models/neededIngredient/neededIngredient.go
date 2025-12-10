package neededingredient

import "e-restaurant/models/enums/measure"

type NeededIngredient struct {
	DishId         int             `db:"dish_id" json:"dish_id"`
	IngredientId   string          `db:"ingredient_id" json:"ingredient_id"`
	NeededQuantity float64         `db:"needed_quantity" json:"needed_quantity"`
	Unit           measure.Measure `db:"unit" json:"unit"`
	IsOptional     bool            `db:"is_optional" json:"is_optional"`
}
