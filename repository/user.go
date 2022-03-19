package repository

import (
	"context"
	"database/sql"

	"github.com/cecepsprd/grpc-gateway-boilerplate/model"
)

type UserRepository interface {
	Read(context.Context) ([]model.User, error)
}

type mysqlUserRepository struct {
	db *sql.DB
}

var (
	readUsersQry = `SELECT id, name, email, password, phone, address, created_at, updated_at FROM user`
)

func NewUserRepository(db *sql.DB) UserRepository {
	return &mysqlUserRepository{
		db: db,
	}
}

func (repo *mysqlUserRepository) Read(ctx context.Context) (response []model.User, err error) {
	rows, err := repo.db.QueryContext(ctx, readUsersQry)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.Phone,
			&user.Address,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		response = append(response, user)
	}

	return response, nil
}
