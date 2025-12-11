package orderitem

import (
	"context"
	orderitem "e-restaurant/models/orderItem"
)

type Repository interface {
	Create(ctx context.Context, oItem *orderitem.OrderItem) (*orderitem.OrderItem, error)

	GetById(ctx context.Context, oItemId int) (*orderitem.OrderItem, error)

	GetAllByOrder(ctx context.Context, orderId int) ([]orderitem.OrderItem, error)

	Delete(ctx context.Context, oItemId int)

	UpdateQuantity(ctx context.Context, oItemId int, quantity int)
}
