package modelinventoryandlost

import "time"

type Measure string

const (
	MeasureU  Measure = "u"
	MeasureKg Measure = "kg"
	MeasureG  Measure = "g"
	MeasureL  Measure = "l"
	MeasureMl Measure = "ml"
)

type Inventory struct {
	InventoryId    int       `db:"inventory_id" json:"inventory_id"`
	RestaurantId   int       `db:"restaurant_id" json:"restaurant_id"`
	IngredientId   string    `db:"ingredient_id" json:"ingredient_id"`
	Stock          uint      `db:"stock" json:"stock"`
	Unit           Measure   `db:"unit" json:"unit"`
	Price          float64   `db:"price" json:"price"`
	ExpirationDate time.Time `db:"expiration_date" json:"expiration_date"`
}

type LostInventory struct {
	LostInventory int       `db:"lost_inventory" json:"lost_inventory"`
	InventoryID   int       `db:"inventory_id" json:"inventory_id"`
	Quantity      float64   `db:"quantity" json:"quantity"`
	Unit          uint      `db:"unit" json:"unit"`
	UnitPrice     float64   `db:"unit_price" json:"unit_price"`
	DateEvent     time.Time `db:"date_time" json:"date_time"`
}
