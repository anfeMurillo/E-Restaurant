package lostinventory

import (
	"e-restaurant/models/enums/measure"
	"time"
)

type LostInventory struct {
	LostInventoryID int             `db:"lost_inventory_id" json:"lost_inventory_id"`
	InventoryID     int             `db:"inventory_id" json:"inventory_id"`
	Quantity        float64         `db:"quantity" json:"quantity"`
	Unit            measure.Measure `db:"unit" json:"unit"`
	UnitPrice       float64         `db:"unit_price" json:"unit_price"`
	DateEvent       time.Time       `db:"date_time" json:"date_time"`
}
