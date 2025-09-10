package services

import (
	"context"

	"github.com/h3th-IV/aml_test/internal/models"
)

func (us *UserService) GetuserById(ctx context.Context, id int) (*models.User, error) {
	var user models.User
	stmt, err := us.db.DB.Prepare("select * from users where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, id).Scan(&user.Id, &user.Name, &user.Email, &user.Gender, &user.Dob, &user.Address)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
