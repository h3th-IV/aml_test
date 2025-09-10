package services

import (
	"context"
	"errors"

	"github.com/h3th-IV/aml_test/internal/config"
	"github.com/h3th-IV/aml_test/internal/models"
)

type UserService struct {
	db *config.Database
}

func NewService(db *config.Database) *UserService {
	return &UserService{
		db: db,
	}
}

func (us *UserService) CreateNewUser(ctx context.Context, name, email, gender, address, dob string) (*models.User, error) {
	stmt, err := us.db.DB.Prepare("insert into users(name, email, gender, dob, address) values (?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, name, email, gender, dob, address)
	if err != nil {
		return nil, err
	}

	lid, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	if lid <= 0 {
		return nil, errors.New("invalid last insert id")
	}

	return &models.User{
		Id:      int(lid),
		Name:    name,
		Email:   email,
		Gender:  gender,
		Dob:     dob,
		Address: address,
	}, nil
}
