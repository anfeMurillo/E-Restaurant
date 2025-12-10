package restaurant

type Restaurant struct {
	RestaurantId   int    `db:"restaurant_id" json:"restaurant_id"`
	RestaurantName string `db:"restaurant_name" json:"restaurant_name"`
	Address        string `db:"address" json:"address"`
	IsActive       bool   `db:"is_active" json:"is_active"`
	IsOpen         bool   `db:"is_open" json:"is_open"`
}
