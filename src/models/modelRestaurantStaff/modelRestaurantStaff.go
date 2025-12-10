package modelrestaurantstaff

type RestaurantStaff struct {
	RestaurantId int `db:"restaurant_id" json:"restaurant_id"`
	UserId       int `db:"user_id" json:"user_id"`
}
