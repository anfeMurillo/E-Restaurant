package lostinventory

import (
	"context"
	"database/sql"
	lostinventory "e-restaurant/models/lostInventory"
	"time"
)

type LostInventoryRepository struct {
	db *sql.DB
}

func NewLostInventoryRepository(db *sql.DB) *LostInventoryRepository {
	return &LostInventoryRepository{db: db}
}

func (l *LostInventoryRepository) Create(ctx context.Context, lInvenotory *lostinventory.LostInventory) (*lostinventory.LostInventory, error) {
	query := `
	INSERT INTO lost_inventories (
	inventory_id,
	quantity,
	unit,
	unit_price,
	date_event
	)
	VALUES ($1,$2,$3,$4,$5)
	RETURNING lost_id,inventory_id
	`
	err := l.db.QueryRowContext(ctx, query,
		lInvenotory.InventoryID,
		lInvenotory.Quantity,
		lInvenotory.Unit,
		lInvenotory.UnitPrice,
		lInvenotory.DateEvent,
	).Scan(
		&lInvenotory.LostInventoryID,
		&lInvenotory.DateEvent,
	)

	if err != nil {
		return nil, err
	}

	return lInvenotory, nil

}

func (l *LostInventoryRepository) GetByRestaurant(ctx context.Context, restaurantId int, limit int, since time.Time) ([]*lostinventory.LostInventory, error) {
	query := `
	SELECT * FROM lost_inventories
	WHERE restaurant_id = $1 AND date_event > $2
	LIMIT $3
	`
	lostItemList := []*lostinventory.LostInventory{}

	rows, err := l.db.QueryContext(ctx, query, restaurantId, since, limit)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		lostItem := &lostinventory.LostInventory{}
		err := rows.Scan(
			&lostItem.LostInventoryID,
			&lostItem.InventoryID,
			&lostItem.Quantity,
			&lostItem.Unit,
			&lostItem.UnitPrice,
			&lostItem.DateEvent,
		)

		if err != nil {
			return nil, err
		}

		lostItemList = append(lostItemList, lostItem)
	}

	return lostItemList, nil
}
