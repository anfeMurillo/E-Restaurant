package modelinventory

import (
	enummeasure "e-restaurant/models/enums/enumMeasure"
	"time"
)

type Inventory struct {
	InventoryId    int                 `db:"inventory_id" json:"inventory_id"`
	RestaurantId   int                 `db:"restaurant_id" json:"restaurant_id"`
	IngredientId   string              `db:"ingredient_id" json:"ingredient_id"`
	Stock          uint                `db:"stock" json:"stock"`
	Unit           enummeasure.Measure `db:"unit" json:"unit"`
	ExpirationDate time.Time           `db:"expiration_date" json:"expiration_date"`
}
