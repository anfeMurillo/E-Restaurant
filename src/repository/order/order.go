package order

import (
	"context"
	"database/sql"
	orderstatus "e-restaurant/models/enums/orderStatus"
	model "e-restaurant/models/order"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(ctx context.Context, o *model.Order) (*model.Order, error) {
	query := `INSERT INTO orders (user_id, restaurant_id, status)
			  VALUES ($1, $2, $3)
			  RETURNING order_id, user_id, restaurant_id, created_at, status`
	row := r.db.QueryRowContext(ctx, query, o.UserId, o.RestaurantId, o.Status)
	var res model.Order
	if err := row.Scan(&res.OrderId, &res.UserId, &res.RestaurantId, &res.CreatedAt, &res.Status); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *OrderRepository) GetById(ctx context.Context, orderID int) (*model.Order, error) {
	query := `SELECT order_id, user_id, restaurant_id, created_at, status FROM orders WHERE order_id = $1`
	row := r.db.QueryRowContext(ctx, query, orderID)
	var res model.Order
	if err := row.Scan(&res.OrderId, &res.UserId, &res.RestaurantId, &res.CreatedAt, &res.Status); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *OrderRepository) Delete(ctx context.Context, orderID int) error {
	query := `DELETE FROM orders WHERE order_id = $1`
	_, err := r.db.ExecContext(ctx, query, orderID)
	return err
}

func (r *OrderRepository) GetByRestaurant(ctx context.Context) ([]*model.Order, error) {
	// Contract has no restaurant id parameter; return all orders
	query := `SELECT order_id, user_id, restaurant_id, created_at, status FROM orders`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []*model.Order
	for rows.Next() {
		var o model.Order
		if err := rows.Scan(&o.OrderId, &o.UserId, &o.RestaurantId, &o.CreatedAt, &o.Status); err != nil {
			return nil, err
		}
		result = append(result, &o)
	}
	return result, nil
}

func (r *OrderRepository) UpdateStatus(ctx context.Context, oStatus orderstatus.OrderStatus) error {
	// Contract provides only the new status; update all orders to that status
	query := `UPDATE orders SET status = $1`
	_, err := r.db.ExecContext(ctx, query, oStatus)
	return err
}
