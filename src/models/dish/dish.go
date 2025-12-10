package dish

type Dish struct {
	DishId       int     `db:"dish_id" json:"dish_id"`
	RestaurantId int     `db:"restaurant_id" json:"restaurant_id"`
	DishName     string  `db:"dish_name" json:"dish_name"`
	Price        float64 `db:"price" json:"price"`
}
