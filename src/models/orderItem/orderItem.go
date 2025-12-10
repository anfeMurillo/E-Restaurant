package orderitem

type OrderItem struct {
	OrderItemId   int     `db:"order_item_id" json:"order_item_id"`
	OrderId       int     `db:"order_id" json:"order_id"`
	DishId        int     `db:"dish_id" json:"dish_id"`
	Quantity      int     `db:"quantity" json:"quantity"`
	PriceAtMoment float64 `db:"price_at_moment" json:"price_at_moment"`
}
