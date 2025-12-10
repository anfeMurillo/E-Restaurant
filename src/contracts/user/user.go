package user

import (
	"context" // Esto es util para la DB y poder cerrar
	userapp "e-restaurant/models/userApp"
)

type Repository interface {
	Create(ctx context.Context, user *userapp.User) (*userapp.User, error)

	GetById(ctx context.Context, userId int) (*userapp.User, error)

	GetRestaurantStaff(ctx context.Context, role string, restaurantId int) ([]*userapp.User, error)

	UpdateName(ctx context.Context, userId int, new string) error

	UpdateEmail(ctx context.Context, userId int, new string) error

	UpdateCountryCode(ctx context.Context, userId int, new string) error

	UpdatePassword(ctx context.Context, userId int, newHash string) error

	ToggleActiveState(ctx context.Context, userId int) error

	AssignRole(ctx context.Context, userId int, role string) error
}
