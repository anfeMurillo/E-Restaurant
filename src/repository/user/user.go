package user

import (
	"context"
	"database/sql"
	userapp "e-restaurant/models/userApp"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) Create(ctx context.Context, user *userapp.User) (*userapp.User, error) {

	query := `INSERT INTO users 
	(
	user_name, 
	email, 
	first_name, 
	last_name, 
	country_code, 
	user_role, 
	password_hash
	) VALUES ($1, $2, $3, $4, $5, $6, $7) 
	 RETURNING user_id, created_at, is_active`
	err := u.db.QueryRowContext(ctx, query, user.UserName, user.Email, user.FirstName, user.LastName, user.CountryCode, user.UserRole, user.PasswordHash).Scan(&user.UserId, &user.CreatedAt, &user.IsActive)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) GetById(ctx context.Context, userId int) (*userapp.User, error) {
	query := `SELECT 
	user_id, 
	user_name, 
	email, 
	first_name, 
	last_name, 
	country_code, 
	user_role, 
	created_at, 
	is_active, 
	FROM users WHERE user_id = $1`
	user := &userapp.User{}
	err := u.db.QueryRowContext(ctx, query, userId).Scan(&user.UserId, &user.UserName, &user.Email, &user.FirstName, &user.LastName, &user.CountryCode, &user.UserRole, &user.CreatedAt, &user.IsActive, &user.PasswordHash)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) GetRestaurantStaff(ctx context.Context, restaurantId int) ([]*userapp.User, error) {
	query := `SELECT 
	u.user_id, 
	u.user_name, 
	u.email, 
	u.first_name, 
	u.last_name, 
	u.country_code, 
	u.user_role, 
	u.created_at, 
	u.is_active
	FROM users u 
	INNER JOIN restaurant_staff rs USING (user_id)`
	rows, err := u.db.QueryContext(ctx, query, restaurantId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*userapp.User
	for rows.Next() {
		user := &userapp.User{}
		err := rows.Scan(&user.UserId, &user.UserName, &user.Email, &user.FirstName, &user.LastName, &user.CountryCode, &user.UserRole, &user.CreatedAt, &user.IsActive)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *UserRepository) UpdateName(ctx context.Context, userId int, new string) error {
	query := `UPDATE users SET user_name = $1 WHERE user_id = $2`
	_, err := u.db.ExecContext(ctx, query, new, userId)
	return err
}
func (u *UserRepository) UpdateFirstName(ctx context.Context, userId int, new string) error {
	query := `UPDATE users SET first_name = $1 WHERE user_id = $2`
	_, err := u.db.ExecContext(ctx, query, new, userId)
	return err
}
func (u *UserRepository) UpdateLastName(ctx context.Context, userId int, new string) error {
	query := `UPDATE users SET last_name = $1 WHERE user_id = $2`
	_, err := u.db.ExecContext(ctx, query, new, userId)
	return err
}

func (u *UserRepository) UpdateEmail(ctx context.Context, userId int, new string) error {
	query := `UPDATE users SET email = $1 WHERE user_id = $2`
	_, err := u.db.ExecContext(ctx, query, new, userId)
	return err
}

func (u *UserRepository) UpdateCountryCode(ctx context.Context, userId int, new string) error {
	query := `UPDATE users SET country_code = $1 WHERE user_id = $2`
	_, err := u.db.ExecContext(ctx, query, new, userId)
	return err
}

func (u *UserRepository) UpdatePassword(ctx context.Context, userId int, newHash string) error {
	query := `UPDATE users SET password_hash = $1 WHERE user_id = $2`
	_, err := u.db.ExecContext(ctx, query, newHash, userId)
	return err
}

func (u *UserRepository) ToggleActiveState(ctx context.Context, userId int) error {
	query := `UPDATE users SET is_active = NOT is_active WHERE user_id = $1`
	_, err := u.db.ExecContext(ctx, query, userId)
	return err
}

func (u *UserRepository) AssignRole(ctx context.Context, userId int, role string) error {
	query := `UPDATE users SET user_role = $1 WHERE user_id = $2`
	_, err := u.db.ExecContext(ctx, query, role, userId)
	return err
}

func (u *UserRepository) Delete(ctx context.Context, userID int) error {
	query := `
	DELETE FROM users
	WHERE user_id = $1
	`
	_, err := u.db.ExecContext(ctx, query, userID)

	return err
}
