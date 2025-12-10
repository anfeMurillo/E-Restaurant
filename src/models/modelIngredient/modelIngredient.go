package main

type Ingredient struct {
	IngredientId   string `db:"ingredient_id" json:"ingredient_id"`
	IngredientName string `db:"ingredient_name" json:"ingredient_name"`
}
