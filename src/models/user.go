package models

import (
	"context"
	"studi_kasus_xyz/entities"
)

func RegisterUser(ctx context.Context, data entities.UserRegister) error {
	hashedPassword, err := hashPassword(data.Password)
	if err != nil {
		return err
	}

	_, err = dbPool.Query("INSERT INTO user (password, username) VALUES (?, ?)", hashedPassword, data.Username)
	if err != nil {
		return err
	}

	return nil
}
