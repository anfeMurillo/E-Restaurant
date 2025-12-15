package dish

import (
	"context"
	"database/sql"
	"e-restaurant/models/dish"
)

type DishRepository struct {
	db *sql.DB
}

func NewDishRepository(db *sql.DB) *DishRepository {
	return &DishRepository{db: db}
}

func (d *DishRepository) Create(ctx context.Context, dish *dish.Dish) (*dish.Dish, error) {
	query := `
	INSER INTO dishes (restaurant_id, dish_name, price) 
	VALUES ($1,$2,$3)
	RETURNING dish_id,dish_name
	;
	`
	err := d.db.QueryRowContext(ctx, query, dish.RestaurantId, dish.DishName, dish.Price).Scan(&dish.DishId, &dish.DishName)
	if err != nil {
		return nil, err
	}
	return dish, nil
}

func (d *DishRepository) GetById(ctx context.Context, dishId int) (*dish.Dish, error) {
	query := `
	SELECT * FROM dishes
	WHERE dish_id = $1
	`
	dish := &dish.Dish{}

	err := d.db.QueryRowContext(ctx, query, dishId).Scan(&dish.DishId, &dish.RestaurantId, &dish.DishName, &dish.Price)
	if err != nil {
		return nil, err
	}
	return dish, nil
}

func (d *DishRepository) GetByRestaurant(ctx context.Context, restaurantId int) (*dish.Dish, error) {
	query := `
	SELECT
	
	d.dish_id,
	d.restaurant_id,
	d.dish_name,
	d.price
	
	FROM dishes d
	INNER JOIN restaurants USING (restaurant_id)
	WHERE dish_id = $1
	`
	dish := &dish.Dish{}

	err := d.db.QueryRowContext(ctx, query, restaurantId).Scan(&dish.DishId, &dish.RestaurantId, &dish.DishName, &dish.Price)
	if err != nil {
		return nil, err
	}
	return dish, nil
}

func (d *DishRepository) GetAll(ctx context.Context) ([]*dish.Dish, error) {

	query := `
	SELECT * FROM dishes
	`
	dishList := []*dish.Dish{}

	rows, err := d.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		dish := &dish.Dish{}

		err := rows.Scan(&dish.DishId, &dish.RestaurantId, &dish.DishName, &dish.Price)
		if err != nil {
			return nil, err
		}

		dishList = append(dishList, dish)
	}

	return dishList, nil

}

func (d *DishRepository) UpdateName(ctx context.Context, dishId int, new string) error {

	query := `
	UPDATE dishes SET dish_name = $1
	WHERE dish_id = $2
	`

	_, err := d.db.ExecContext(ctx, query, new, dishId)

	return err
}

func (d *DishRepository) UpdatePrice(ctx context.Context, dishId int, new int) error {
	query := `
	UPDATE dishes SET dish_price = $1
	WHERE dish_id = $2
	`
	_, err := d.db.ExecContext(ctx, query, new, dishId)

	return err
}

func (d *DishRepository) Delete(ctx context.Context, dishID int) error {
	query := `
	DELETE FROM dishes
	WHERE dish_id = $1
	`
	_, err := d.db.ExecContext(ctx, query, dishID)

	return err
}
