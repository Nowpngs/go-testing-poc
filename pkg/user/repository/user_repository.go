package repository

import (
	"context"
	"database/sql"
	modal "go-testing-poc/pkg/user"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *modal.User) error
	GetUserList(ctx context.Context, role *modal.Role) ([]*modal.User, error)
	GetUserById(ctx context.Context, id string) (*modal.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *modal.User) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO users (username, email) VALUES ($1, $2)", user.Username, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetUserList(ctx context.Context, role *modal.Role) ([]*modal.User, error) {
	query := "SELECT id, username, email, created_at, updated_at, valid FROM users"
	var args []interface{}

	if role != nil {
		query += " WHERE role = $1"
		args = append(args, *role)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*modal.User, 0)
	for rows.Next() {
		user := new(modal.User)
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.Valid)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) GetUserById(ctx context.Context, id string) (*modal.User, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, username, email, created_at, updated_at, valid FROM users WHERE id = $1", id)

	user := new(modal.User)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.Valid)
	if err != nil {
		return nil, err
	}

	return user, nil
}
